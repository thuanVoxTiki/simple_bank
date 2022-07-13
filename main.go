package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thuanVoxTiki/simple_bank/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "postgresql://postgres:123456@localhost:5432/simple_bank"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db = db.Debug()
	runService(db)
	if err != nil {
		fmt.Println("DB has Error")
		panic(gorm.ErrInvalidDB)
	}
	fmt.Println("Hello world")
}

func runService(db *gorm.DB) {
	r := gin.Default()
	r.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("createAccount", func(ctx *gin.Context) {
		var account models.Account
		if err := ctx.ShouldBind(&account); err != nil {
			fmt.Printf("Account %v", account)
			panic(err)
		}
		account.Create_at = time.Now()
		// var account = models.Account{
		// 		id: 1,
		// 		owner: "Phuc",
		// 		balance: 100000.0,
		// 		currency: "hihi",
		// }
		fmt.Printf("Account %v", account)

		if err := db.Create(&account).Error; err != nil {
			fmt.Printf("Error: %v", err)
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err,
			})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Okila",
		})

	})

	r.PUT("accounts/", func(ctx *gin.Context) {
		var account models.Account
		if err := ctx.ShouldBind(&account); err != nil {
			fmt.Printf("Account %v", account)
			panic(err)
		}

		if err := db.Where("id = ?",account.Id).Updates(account).Error; err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Update successfully",
		})

	})
	r.Run()
}
