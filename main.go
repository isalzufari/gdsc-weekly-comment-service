package main

import (
	"commentservice/db"

	"github.com/gin-gonic/gin"
)

func main() {

	db := db.OpenDB(
		"root",
		"",
		"127.0.0.1",
		"3306",
		"micro_web_gdsc",
	)
	defer db.Close()

	//

	r := gin.Default()

	// routing

	port := "5001"
	r.Run(":" + port)
}
