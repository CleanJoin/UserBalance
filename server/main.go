package main

import (
	"fmt"
	"os"
	"strconv"

	"server/balance"
)

// @title           Swagger USERBALANCE
// @version         1.0
// @description     This is a sample server USERBALANCE
// @termsOfService  https://github.com/CleanJoin/USERBALANCE/

// @contact.name   Github.com
// @contact.url    https://github.com/CleanJoin/USERBALANCE/

// @host      localhost:8000
// @BasePath  /
func main() {
	serverPort, _ := strconv.Atoi(os.Getenv("SERVER_PORT"))
	maxMessagesNum, _ := strconv.Atoi(os.Getenv("SERVER_MAX_MESSAGES"))

	fmt.Println("Server port env:", serverPort)
	fmt.Println("Max messages num env:", maxMessagesNum)

	chatServerGin := balance.NewChatServerGin("localhost", serverPort, uint(maxMessagesNum))

	usersStorage := balance.NewUserStorageDB(new(balance.PasswordHasherSha1), balance.NewConnectDB(5432))
	messageStorage := balance.NewBalanceStorageDB(balance.NewConnectDB(5432), usersStorage)
	chatServerGin.Use(usersStorage, messageStorage)
	chatServerGin.Run()
}
