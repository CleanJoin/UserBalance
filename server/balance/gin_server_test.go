package balance

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
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
	r := rand.Intn(900000000)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/user", strings.NewReader(`{"username": "Andrey`+strconv.Itoa(r)+`", "password": "Andrey"}`))

	chatServer.router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

}
func TestGetHealth(t *testing.T) {
	godotenv.Load(".env")
	usersStorage := NewUserStorageDB(new(PasswordHasherSha1), NewConnectDB(5432))
	chatServer := NewServerGin("localhost", 8000)
	balanceStorage := NewBalanceStorageDB(NewConnectDB(5432))
	chatServer.Use(usersStorage, balanceStorage)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/health", nil)
	chatServer.router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestAddMoneyHandler(t *testing.T) {
	godotenv.Load(".env")
	usersStorage := NewUserStorageDB(new(PasswordHasherSha1), NewConnectDB(5432))

	chatServer := NewServerGin("localhost", 8000)
	balanceStorage := NewBalanceStorageDB(NewConnectDB(5432))
	chatServer.Use(usersStorage, balanceStorage)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/add", strings.NewReader(`{"userid": 1, "money": 66.1}`))
	d := w.Body.String()
	fmt.Println(d)
	chatServer.router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestReduceMoneyHandler(t *testing.T) {
	godotenv.Load(".env")
	usersStorage := NewUserStorageDB(new(PasswordHasherSha1), NewConnectDB(5432))

	chatServer := NewServerGin("localhost", 8000)
	balanceStorage := NewBalanceStorageDB(NewConnectDB(5432))
	chatServer.Use(usersStorage, balanceStorage)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/reduce", strings.NewReader(`{"userid": 1, "money": 33.0}`))
	d := w.Body.String()
	fmt.Println(d)
	chatServer.router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestTransferMoneyHandler(t *testing.T) {
	godotenv.Load(".env")
	usersStorage := NewUserStorageDB(new(PasswordHasherSha1), NewConnectDB(5432))

	chatServer := NewServerGin("localhost", 8000)
	balanceStorage := NewBalanceStorageDB(NewConnectDB(5432))
	chatServer.Use(usersStorage, balanceStorage)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/transfer", strings.NewReader(`{"Money": 1,"UserIdFrom": 2,"UserIdTo": 1}`))
	d := w.Body.String()
	fmt.Println(d)
	chatServer.router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetMoneyUserHadler(t *testing.T) {
	godotenv.Load(".env")
	usersStorage := NewUserStorageDB(new(PasswordHasherSha1), NewConnectDB(5432))

	chatServer := NewServerGin("localhost", 8000)
	balanceStorage := NewBalanceStorageDB(NewConnectDB(5432))
	chatServer.Use(usersStorage, balanceStorage)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/money?currency=USD", strings.NewReader(`{"userid": 1}`))
	d := w.Body.String()
	fmt.Println(d)
	chatServer.router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestGetLastTransactionHadler(t *testing.T) {
	godotenv.Load(".env")
	usersStorage := NewUserStorageDB(new(PasswordHasherSha1), NewConnectDB(5432))

	chatServer := NewServerGin("localhost", 8000)
	balanceStorage := NewBalanceStorageDB(NewConnectDB(5432))
	chatServer.Use(usersStorage, balanceStorage)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/getmovemoney?page=1&filtermoney=asc&filtertime=desc", strings.NewReader(`{"userid": 1}`))
	d := w.Body.String()
	fmt.Println(d)
	chatServer.router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}
