package balance

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
)

func TestAddMoney(t *testing.T) {
	godotenv.Load(".env")
	var inter ItransactionsStorage = NewBalanceStorageDB(NewConnectDB(5432))
	transaction, err := inter.AddMoney(1, 33.3)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(transaction)
}
func TestWriteOffMoney(t *testing.T) {
	godotenv.Load(".env")
	var inter ItransactionsStorage = NewBalanceStorageDB(NewConnectDB(5432))
	transaction, err := inter.WriteOffMoney(1, 1.3)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(transaction)
}

func TestTransferMoney(t *testing.T) {
	godotenv.Load(".env")
	var inter ItransactionsStorage = NewBalanceStorageDB(NewConnectDB(5432))
	transaction, err := inter.TransferMoney(2, 1, 10.0)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(transaction)
}
func TestAddTransferMoney(t *testing.T) {
	godotenv.Load(".env")
	var balanceStorageDB = NewBalanceStorageDB(NewConnectDB(5432))
	transaction := TransactionsModel{}
	err := addTransferMoney(balanceStorageDB, transaction)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(transaction)
}
func TestListRecords(t *testing.T) {
	godotenv.Load(".env")
	var inter ItransactionsStorage = NewBalanceStorageDB(NewConnectDB(5432))
	_, err := inter.ListRecords(1, "asc", "desc", 1)
	if err != nil {
		t.Errorf(err.Error())
	}

}
