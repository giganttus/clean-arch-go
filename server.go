package main

import (
	"os"

	"todoapp/helpers"
	"todoapp/routes"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(helpers.ErrEnvLoad)
	}

	if os.Getenv("CONN_STR") == "" {
		panic(helpers.ErrEnvLoad)
	}

	db, err := gorm.Open(mysql.Open(os.Getenv("CONN_STR")), &gorm.Config{})
	if err != nil {
		panic(helpers.ErrDbConn)
	}

	routes := routes.InitRoutes(db)

	routes.Run("0.0.0.0:1313")
}
