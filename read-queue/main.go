package main

import (
	"fmt"
	"readq/internal/config"
	"readq/internal/server"
	"readq/internal/utils"
	"readq/internal/utils/postgresql"
)

func main() {
	config := config.GetConfig()
	fmt.Println("ConnectionString : ", config.PostgresDB.Postgres_connectionstring)
	postgresql.InitDatabase(config.PostgresDB.Postgres_connectionstring)

	utils.InitRedis()
	server.GetDataFromStreamInfinite()
}
