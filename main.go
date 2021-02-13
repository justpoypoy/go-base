package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/justpoypoy/go-base/infrastructure"
)

func main() {
	logger := infrastructure.NewLogger()
	infrastructure.Load(logger)

	sqlHandler, err := infrastructure.NewSQL()
	if err != nil {
		logger.LogError("%s", err)
	}

	infrastructure.Dispatch(logger, sqlHandler)
}
