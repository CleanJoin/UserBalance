package balance

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert"
	"github.com/joho/godotenv"
)

func TestUse(t *testing.T) {
	chatServer := NewServerGin("localhost", 8000)
	usersstorage := NewUserStorageDB(new(PasswordHasherSha1), NewConnectDB(5432))
	transactions := NewBalanceStorageDB(NewConnectDB(5432))
	chatServer.Use(usersstorage, transactions)
	if chatServer.router == nil {
		t.Errorf("router не сконфигурирован  %v", chatServer.router)
	}

}

func TestRun(t *testing.T) {
	chatServer := NewServerGin("localhost", 8000)
	chatServer.Run()
	if chatServer.router != nil {
		t.Errorf("router сконфигурирован  %v", chatServer.router)
	}
}

func TestValidatenUserName(t *testing.T) {
	userName := "Andrey"
	if !validatenUserName(userName) {
		t.Errorf("не корректный userName")
	}
}

func TestValidatenUserNameMore(t *testing.T) {
	userName := "Andreydsfdsfsdfdsfsdfdsfdsfdsfdsfdsfsdfdsfdsfdsfdsfdsfdsfdsfdsfdsfdsfds"
	if validatenUserName(userName) {
		t.Errorf("Длина userName меньше 16 символов")
	}
}

func TestValidatenUserNameNotCorrect(t *testing.T) {
	userName := "Andreyd!"
	if validatenUserName(userName) {
		t.Errorf("userName корректный")
	}
}

func TestValidatePasswordLess(t *testing.T) {
	password := "Jdsfsdfdsfsdfdsfdsfsdfs"
	if !validatePassword(password) {
		t.Errorf("Длина пароля больше 32 символов")
	}
}

func TestValidatePasswordMore(t *testing.T) {
	password := "Jdsfsdfdsfsdfdsfdsfsdfsrertyerutuyertyeruitrey"
	if validatePassword(password) {
		t.Errorf("Длина меньше 32 символов")
	}
}
func TestUserHandler(t *testing.T) {
	godotenv.Load(".env")
	usersStorage := NewUserStorageDB(new(PasswordHasherSha1), NewConnectDB(5432))

	chatServer := NewServerGin("localhost", 8000)
	balanceStorage := NewBalanceStorageDB(NewConnectDB(5432))
	chatServer.Use(usersStorage, balanceStorage)

	values := map[string]string{"username": "Andrey5", "password": "fghfghfghfgh"}
	jsonValue, _ := json.Marshal(values)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/user", bytes.NewBuffer(jsonValue))

	chatServer.router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

}
func TestGet(t *testing.T) {
	godotenv.Load(".env")
	usersStorage := NewUserStorageDB(new(PasswordHasherSha1), NewConnectDB(5432))

	chatServer := NewServerGin("localhost", 8000)
	balanceStorage := NewBalanceStorageDB(NewConnectDB(5432))
	chatServer.Use(usersStorage, balanceStorage)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/health", nil)

	chatServer.router.ServeHTTP(w, req)
	d := w.Body.String()
	fmt.Println(d)
	assert.Equal(t, http.StatusOK, w.Code)

}
