package balance

import (
	"testing"
)

// func TestUse(t *testing.T) {
// 	chatServer := NewChatServerGin("localhost", 8080, 300)
// 	usersstorage := NewUserStorageMemory(new(PasswordHasherSha1))
// 	chatServer.Use(sessionStorage, usersstorage, messageStorage)
// 	if chatServer.router == nil {
// 		t.Errorf("router не сконфигурирован  %v", chatServer.router)
// 	}

// }

func TestRun(t *testing.T) {
	chatServer := NewChatServerGin("localhost", 8080, 300)
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
