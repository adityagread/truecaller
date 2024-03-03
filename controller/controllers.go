package controller

import (
	"aditya-coding-task/config"
	"aditya-coding-task/helpers"
	"aditya-coding-task/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = config.DB

// marking the number spam
func MarkAsSpam(c *gin.Context) {

	var spamNumber models.SpamNumber
	if err := c.ShouldBindJSON(&spamNumber); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := db.Create(&spamNumber)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Already in spam"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Spam number reported successfully"})
}

// registering the user in db
func RegisterUser(c *gin.Context) {
	var user models.Registered
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// if user already exists update the token
	if flag := db.Where("phone_number = ?", user.PhoneNumber).First(&user).RowsAffected; flag > 0 {
		db.Model(&models.Registered{}).Where("ID = ?", user.ID).Update("token", token)
		c.JSON(http.StatusCreated, gin.H{"token": token, "update": "ok"})
		return
	}

	global := models.User{
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
	}
	user.Token = token
	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}
	GlobalUser(global)
	c.JSON(http.StatusCreated, gin.H{"token": token})
}

// inserting registered user into global db
func GlobalUser(user models.User) {
	result := db.Create(&user)
	if result.Error != nil {
		return
	}
}

// searching using number
func SearchByPhone(c *gin.Context) {

	var searchQuery struct {
		PhoneNumber string `json:"phone_number" binding:"required"`
	}

	if err := c.ShouldBindJSON(&searchQuery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var registeredUser models.Registered
	var spamResult models.SpamNumber
	markSpam := false
	if spam := db.Table("spam_numbers").Where("number = ?", searchQuery.PhoneNumber).First(&spamResult).RowsAffected; spam > 0 {
		markSpam = true
	}

	if result := db.Where("phone_number = ?", searchQuery.PhoneNumber).First(&registeredUser); result.Error == nil {
		c.JSON(http.StatusOK, gin.H{"result": "Registered User", "details": registeredUser, "spam": markSpam})
		return
	}

	var searchResults []models.User
	db.Select("name", "phone_number").Where("phone_number = ?", searchQuery.PhoneNumber).Find(&searchResults)

	c.JSON(http.StatusOK, gin.H{"result": "Non Registered User", "details": searchResults, "spam": markSpam})
}

// searching using name
func SearchByName(c *gin.Context) {

	var searchQuery struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&searchQuery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	condition := fmt.Sprintf("CASE WHEN name LIKE '%s' THEN 1 ELSE 2 END, name", "%"+searchQuery.Name+"%")

	var searchResults []models.User
	db.Where("name LIKE ?", "%"+searchQuery.Name+"%").Order(condition).Find(&searchResults)

	var finalResults []models.SearchByNameResult
	for _, val := range searchResults {
		var spamResult models.SpamNumber
		markSpam := false
		if spam := db.Table("spam_numbers").Where("number = ?", val.PhoneNumber).First(&spamResult).RowsAffected; spam > 0 {
			markSpam = true
		}
		finalResults = append(finalResults, models.SearchByNameResult{GlobalUser: val, MarkAsSpam: markSpam})
	}
	c.JSON(http.StatusOK, gin.H{"results": finalResults})
}
