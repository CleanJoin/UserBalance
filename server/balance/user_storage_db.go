package balance

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UserStorageDB struct {
	Users          []UserModel
	passwordHasher IPasswordHasher
	connect        *pgxpool.Pool
}

type IUserStorage interface {
	Create(username string, password string) (UserModel, error)
	GetByName(name string) (UserModel, error)
	GetById(id uint) (UserModel, error)
}

func NewUserStorageDB(passwordHasher IPasswordHasher, iConnectDB IConnectDB) *UserStorageDB {
	udb := new(UserStorageDB)
	udb.connect = iConnectDB.Use()
	udb.passwordHasher = passwordHasher
	return udb
}

func (userStorageDB *UserStorageDB) Create(username string, password string) (UserModel, error) {
	var id uint
	newPasswordHash := userStorageDB.passwordHasher.CalculateHash(password)
	query := `INSERT INTO "avito".users (username,"password") VALUES($1, $2) RETURNING id;`
	row := userStorageDB.connect.QueryRow(context.Background(), query, username, newPasswordHash)
	err := row.Scan(&id)
	if err != nil {
		return UserModel{}, fmt.Errorf(err.Error())
	}
	return UserModel{id, username, newPasswordHash, 0}, nil
}

func (userStorageDB *UserStorageDB) GetByName(name string) (UserModel, error) {
	userModel := new(UserModel)
	query := `select * from "avito".users u where username like $1`
	commandTag := userStorageDB.connect.QueryRow(context.Background(), query, name)
	err := commandTag.Scan(&userModel.ID, &userModel.Username, &userModel.PasswordHash, &userModel.Money)
	if err != nil {
		return *userModel, fmt.Errorf("пользователь не найден")
	}
	return *userModel, nil
}

func (userStorageDB *UserStorageDB) GetById(id uint) (UserModel, error) {
	userModel := new(UserModel)
	query := `select * from "avito".users u where id=$1`
	commandTag := userStorageDB.connect.QueryRow(context.Background(), query, id)
	err := commandTag.Scan(&userModel.ID, &userModel.Username, &userModel.PasswordHash, &userModel.Money)
	if err != nil {
		return *userModel, fmt.Errorf(err.Error())
	}

	return *userModel, nil
}
