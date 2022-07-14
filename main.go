package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"backend/simple_bank/api"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "postgresql://postgres:123456@localhost:5432/simple_bank"
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DB has Error")
		panic(gorm.ErrInvalidDB)
	}

	db = db.Debug()
	fmt.Println("Run service on 8080")
	runService(db)
}

func runService(db *gorm.DB) {
	r := gin.Default()
	r.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	account:= r.Group("/accounts")
	{
		account.POST("", api.CreateAccount(db))
		account.PUT("", api.UpdateAccount(db))
		account.DELETE("/:id", api.DeleteAccount(db))
		account.GET("/:id", api.GetAccount(db))
		account.GET("", api.ListAccounts(db))
	}


	r.Run()
}
