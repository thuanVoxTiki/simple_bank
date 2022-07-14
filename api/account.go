package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"backend/simple_bank/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateAccount(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var account models.Account
		if err := ctx.ShouldBind(&account); err != nil {
			fmt.Printf("Account %v", account)
			panic(err)
		}
		account.Create_at = time.Now()
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
	}
}

func UpdateAccount(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var account models.Account
		// id := ctx.Param("id")
		if err := c.ShouldBind(&account); err != nil {
			fmt.Printf("Account %v", account)
			panic(err)
		}

		if err := db.Updates(account).Error; err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Update successfully",
		})
	}
}

func DeleteAccount(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(err)
		}

		if err := db.Delete(id).Error; err!=nil {
			fmt.Printf("Delete fail %v\n", err)
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Delete succeed",
		})
	}
}

func GetAccount(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var account models.Account

		if err := db.Where("id = ?", id).First(&account).Error; err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, account)
	}
}

func ListAccounts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var accounts []models.Account
		if err := db.Find(&accounts).Error; err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, accounts)
	}
}

