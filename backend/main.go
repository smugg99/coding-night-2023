package main

import (
	"fmt"
	"os"

	"github.com/smugg99/coding-night-2023/config"
	"github.com/smugg99/coding-night-2023/database"
	"github.com/smugg99/coding-night-2023/router"
)

func main() {
	config.Load()

	db, err := database.Setup()
	if err != nil {
		fmt.Println(fmt.Errorf("error setting up database: %s", err.Error()))
		os.Exit(2)
	}
	r := router.Setup(db)
	r.Listen(fmt.Sprintf("0.0.0.0:%d", config.Conf.SrvConfig.Port))
}
