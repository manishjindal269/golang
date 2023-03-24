package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/manishjindal269/golang/models"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

// get all users
func (uc *UserController) FindUsers(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var users []models.User
	results := uc.DB.Limit(intLimit).Offset(offset).Find(&users)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(users), "data": users})
}

func (uc *UserController) FindUsersTest(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	type User struct {
		Id   uint   `json:"id" gorm:"primaryKey"`
		Name string `json:"name"`
	}
	var users []User
	results := uc.DB.Limit(intLimit).Offset(offset).Find(&users)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(users), "data": users})
}

// create new user
func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.UserCreate
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	if err := uc.DB.Create(&user).Error; err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
}

// update user
func (uc *UserController) UpdateUser(ctx *gin.Context) {

	userId := ctx.Param("userId")
	var user models.UserUpdate
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}
	var updatedUser models.User
	result := uc.DB.First(&updatedUser, "id = ?", userId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No user with that title exists"})
		return
	}

	uc.DB.Model(&updatedUser).Updates(user)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedUser})
}

// get all users
func (uc *UserController) GetProducts(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var products []models.ProductWithImgae
	results := uc.DB.Limit(intLimit).Offset(offset).Find(&products)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}
	var image models.Image
	var returnData []models.ProductWithImgae
	for _, product := range products {
		var data models.ProductWithImgae
		err := json.Unmarshal([]byte(product.Image), &image)
		if err != nil {
			panic(err)
		}
		data.ID = product.ID
		data.Title = product.Title
		data.Vendor = product.Vendor
		data.Image = image.Src
		returnData = append(returnData, data)
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": returnData})
}
