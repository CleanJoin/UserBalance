package balance

import (
	"context"
	"fmt"

	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type BalanceStorageDB struct {
	transaction TransactionsModel
	connect     *pgxpool.Pool
	interUser   IUserStorage
	limit       int
}
type ItransactionsStorage interface {
	AddMoney(userId uint, money float64) (TransactionsModel, error)
	WriteOffMoney(userId uint, money float64) (TransactionsModel, error)
	TransferMoney(userIdFrom uint, userIdTo uint, money float64) (TransactionsModel, error)
	ListRecords(page int, filtermoney string, filtertime string, userid int) ([]TransactionsModel, error)
}

func NewBalanceStorageDB(iConnectDB IConnectDB) *BalanceStorageDB {
	sdb := new(BalanceStorageDB)
	sdb.limit = 10
	sdb.connect = iConnectDB.Use()
	sdb.interUser = NewUserStorageDB(new(PasswordHasherSha1), iConnectDB)
	return sdb
}

func (balanceStorageDB *BalanceStorageDB) AddMoney(userId uint, money float64) (TransactionsModel, error) {
	var id int
	userModel, err := balanceStorageDB.interUser.GetById(userId)
	if err != nil {
		return TransactionsModel{}, fmt.Errorf(err.Error())
	}

	query := `UPDATE avito."users" u set "money" = $1 WHERE id=$2 RETURNING id;`
	row := balanceStorageDB.connect.QueryRow(context.Background(), query, userModel.Money+money, userId)
	err = row.Scan(&id)
	if err != nil {
		return TransactionsModel{}, fmt.Errorf(err.Error())
	}
	balanceStorageDB.transaction.UserIdFrom = 0
	balanceStorageDB.transaction.UserIdTo = userId
	balanceStorageDB.transaction.Money = money
	balanceStorageDB.transaction.Time = time.Now()
	balanceStorageDB.transaction.ID = uint(id)
	err = addTransferMoney(balanceStorageDB, balanceStorageDB.transaction)
	if err != nil {
		return TransactionsModel{}, fmt.Errorf(err.Error())
	}
	return balanceStorageDB.transaction, nil
}

func (balanceStorageDB *BalanceStorageDB) WriteOffMoney(userId uint, money float64) (TransactionsModel, error) {
	var id uint
	userModel, err := balanceStorageDB.interUser.GetById(userId)
	if err != nil {
		return TransactionsModel{}, fmt.Errorf(err.Error())
	}
	if userModel.Money-money < 0 {
		return TransactionsModel{}, fmt.Errorf("недостаточно средств для списания")
	}
	query := `UPDATE avito."users" u set "money" = $1 WHERE id=$2 RETURNING id;`
	row := balanceStorageDB.connect.QueryRow(context.Background(), query, userModel.Money-money, userId)
	err = row.Scan(&id)
	if err != nil {
		return TransactionsModel{}, fmt.Errorf(err.Error())
	}
	balanceStorageDB.transaction.UserIdTo = 0
	balanceStorageDB.transaction.UserIdFrom = userId
	balanceStorageDB.transaction.Time = time.Now()
	balanceStorageDB.transaction.Money = money
	balanceStorageDB.transaction.ID = id

	err = addTransferMoney(balanceStorageDB, balanceStorageDB.transaction)
	if err != nil {
		return TransactionsModel{}, fmt.Errorf(err.Error())
	}

	return balanceStorageDB.transaction, nil
}

func (balanceStorageDB *BalanceStorageDB) TransferMoney(userIdFrom uint, userIdTo uint, money float64) (TransactionsModel, error) {
	var id uint
	userModel, err := balanceStorageDB.interUser.GetById(userIdFrom)
	if err != nil {
		fmt.Println(err)
	}
	if userModel.Money-money < 0 {
		return TransactionsModel{}, fmt.Errorf("недостаточно средств для списания")
	}
	query := `UPDATE avito."users" u set "money" = $1 WHERE id=$2 RETURNING id;`
	row := balanceStorageDB.connect.QueryRow(context.Background(), query, userModel.Money-money, userIdFrom)
	err = row.Scan(&id)
	if err != nil {
		return TransactionsModel{}, fmt.Errorf(err.Error())
	}
	balanceStorageDB.transaction.UserIdTo = userIdTo
	balanceStorageDB.transaction.UserIdFrom = userIdFrom
	balanceStorageDB.transaction.Money = money
	balanceStorageDB.transaction.Time = time.Now()
	balanceStorageDB.transaction.ID = id

	err = addTransferMoney(balanceStorageDB, balanceStorageDB.transaction)
	if err != nil {
		return TransactionsModel{}, fmt.Errorf(err.Error())
	}
	userModel, err = balanceStorageDB.interUser.GetById(userIdTo)
	if err != nil {
		return TransactionsModel{}, fmt.Errorf(err.Error())
	}
	query = `UPDATE avito."users" u set "money" = $1 WHERE id=$2 RETURNING id;`
	row = balanceStorageDB.connect.QueryRow(context.Background(), query, userModel.Money+money, userIdTo)
	err = row.Scan(&id)
	if err != nil {
		return TransactionsModel{}, fmt.Errorf(err.Error())
	}
	balanceStorageDB.transaction.UserIdTo = userIdTo
	balanceStorageDB.transaction.UserIdFrom = userIdFrom
	balanceStorageDB.transaction.Money = money
	balanceStorageDB.transaction.Time = time.Now()
	balanceStorageDB.transaction.ID = id

	err = addTransferMoney(balanceStorageDB, balanceStorageDB.transaction)
	if err != nil {
		return TransactionsModel{}, fmt.Errorf(err.Error())
	}

	return balanceStorageDB.transaction, nil
}

func (balanceStorageDB *BalanceStorageDB) ListRecords(page int, filtermoney string, filtertime string, userid int) ([]TransactionsModel, error) {

	offset := balanceStorageDB.limit * (page - 1)

	pageTransaction := []TransactionsModel{}
	query := `SELECT * FROM "avito"."transaction" where useridfrom=$3  or  useridto=$3 ORDER BY "money"` + filtermoney + `,transaction_time ` + filtertime + ` LIMIT $2 OFFSET $1`

	commandTag, err := balanceStorageDB.connect.Query(context.Background(), query, offset, balanceStorageDB.limit, userid)
	if err != nil {
		return []TransactionsModel{}, fmt.Errorf(err.Error())
	}
	for commandTag.Next() {
		err := commandTag.Scan(&balanceStorageDB.transaction.ID, &balanceStorageDB.transaction.UserIdFrom, &balanceStorageDB.transaction.Money, &balanceStorageDB.transaction.Time, &balanceStorageDB.transaction.UserIdTo)
		pageTransaction = append(pageTransaction, balanceStorageDB.transaction)
		if err != nil {
			return []TransactionsModel{}, fmt.Errorf(err.Error())
		}
	}
	return pageTransaction, nil

}
func addTransferMoney(balanceStorageDB *BalanceStorageDB, transaction TransactionsModel) error {
	var id uint
	query := `INSERT INTO "avito"."transaction" (useridto,"money",useridfrom) VALUES($1,$2,$3) RETURNING id;`
	row := balanceStorageDB.connect.QueryRow(context.Background(), query, transaction.UserIdTo, transaction.Money, transaction.UserIdFrom)
	err := row.Scan(&id)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}
