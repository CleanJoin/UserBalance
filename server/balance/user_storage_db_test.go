package balance

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
)

func TestUserCreateDB(t *testing.T) {
	godotenv.Load(".env")
	connectDB := NewConnectDB(5432)
	var passwordHash IPasswordHasher = new(PasswordHasherSha1)
	var inter IUserStorage = NewUserStorageDB(passwordHash, connectDB)
	userModel, err := inter.Create("Andrey", "qweasd123")
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(userModel)
}

func TestGetByIdDB(t *testing.T) {
	godotenv.Load(".env")

	connectDB := NewConnectDB(5432)
	var passwordHash IPasswordHasher = new(PasswordHasherSha1)
	var inter IUserStorage = NewUserStorageDB(passwordHash, connectDB)
	userModel, err := inter.GetById(1)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(userModel)
}

func TestGetByNameDB(t *testing.T) {
	godotenv.Load(".env")
	connectDB := NewConnectDB(5432)
	var passwordHash IPasswordHasher = new(PasswordHasherSha1)
	var inter IUserStorage = NewUserStorageDB(passwordHash, connectDB)
	userModel, err := inter.GetByName("Andrey")
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(userModel)
}