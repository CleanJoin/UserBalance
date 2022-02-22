package balance

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type BalanceStorageDB struct {
	transaction []TransactionsModel
	connect     *pgxpool.Pool
	interUser   IUserStorage
}
type ItransactionsStorage interface {
	AddMoney(userId uint, money float64) (TransactionsModel, error)
	WriteOffMoney(userId uint, money float64) (TransactionsModel, error)
	TransferMoney(userIdFrom uint, userIdTo uint, money float64) (TransactionsModel, error)
}

func NewBalanceStorageDB(iConnectDB IConnectDB, interUser IUserStorage) *BalanceStorageDB {

	sdb := new(BalanceStorageDB)
	sdb.connect = iConnectDB.Use()
	sdb.interUser = interUser
	return sdb
}

func (balanceStorageDB *BalanceStorageDB) AddMoney(userId uint, money float64) (TransactionsModel, error) {
	var id uint
	userModel, err := balanceStorageDB.interUser.GetById(userId)
	if err != nil {
		fmt.Println(err)
	}
	query := `UPDATE "users" u set "money" = $money WHERE id=$id RETURNING id;`
	row := balanceStorageDB.connect.QueryRow(context.Background(), query, userId, userModel.Money+money)
	err = row.Scan(&id)
	if err != nil {
		return TransactionsModel{}, fmt.Errorf(err.Error())
	}
	return TransactionsModel{id, 0, userId, time.Now()}, nil
}

func (balanceStorageDB *BalanceStorageDB) WriteOffMoney(userId uint, money float64) (TransactionsModel, error) {
	var id uint
	userModel, err := balanceStorageDB.interUser.GetById(userId)
	if err != nil {
		fmt.Println(err)
	}
	if userModel.Money-money < 0 {
		return TransactionsModel{}, fmt.Errorf("недостаточно средств для списания")
	}
	query := `UPDATE "users" u set "money" = $money WHERE id=$id RETURNING id;`
	row := balanceStorageDB.connect.QueryRow(context.Background(), query, userId, userModel.Money-money)
	err = row.Scan(&id)
	if err != nil {
		return TransactionsModel{}, fmt.Errorf(err.Error())
	}
	return TransactionsModel{id, 0, userId, time.Now()}, nil
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
	query := `UPDATE "users" u set "money" = $money WHERE id=$id RETURNING id;`
	row := balanceStorageDB.connect.QueryRow(context.Background(), query, userIdFrom, userModel.Money-money)
	err = row.Scan(&id)
	if err != nil {
		return TransactionsModel{}, fmt.Errorf(err.Error())
	}
	query = `UPDATE "users" u set "money" = $money WHERE id=$id RETURNING id;`
	row = balanceStorageDB.connect.QueryRow(context.Background(), query, userIdTo, userModel.Money+money)

	return TransactionsModel{id, userIdFrom, userIdTo, time.Now()}, nil
}
