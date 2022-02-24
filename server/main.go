package main

import (
	"os"
	"strconv"

	"userbalance/balance"
)

// @title           Swagger USERBALANCE
// @version         1.0
// @description     This is a sample server USERBALANCE
// @termsOfService  https://github.com/CleanJoin/USERBALANCE/
// @contact.name   Github.com
// @contact.url    https://github.com/CleanJoin/USERBALANCE/
// @host      localhost:8898
// @BasePath  /
func main() {
	serverPort, _ := strconv.Atoi(os.Getenv("SERVER_PORT"))

	serverGin := balance.NewServerGin("localhost", serverPort)

	usersStorage := balance.NewUserStorageDB(new(balance.PasswordHasherSha1), balance.NewConnectDB(5432))
	balanceStorage := balance.NewBalanceStorageDB(balance.NewConnectDB(5432))
	serverGin.Use(usersStorage, balanceStorage)
	serverGin.Run()

}
