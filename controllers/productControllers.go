package controllers

import (
	"Hacktiv10JWT/database"
	"Hacktiv10JWT/helpers"
	"Hacktiv10JWT/models"
	"Hacktiv10JWT/repositories"
	"Hacktiv10JWT/services"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateProduct(ctx *gin.Context) {
	db := database.GetDB()
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contenType := helpers.GetContentType(ctx)

	Product := models.Product{}
	userID := uint(userData["id"].(float64))

	if contenType == appJSON {
		ctx.ShouldBindJSON(&Product)
	} else {
		ctx.ShouldBind(&Product)
	}

	Product.UserID = userID
	fmt.Println("product : ", Product.UserID)
	err := productService.CreateProduct(&Product)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err":     "bad Request",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, Product)
}

func GetProducts(ctx *gin.Context) {
	db := database.GetDB()
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	results, err := productService.GetProductsByUserID(userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error message": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Product data : ": results,
	})

}

func GetAllProducts(ctx *gin.Context) {
	db := database.GetDB()
	productRepo := repositories.NewProductRepository(db)

	results, err := productRepo.GetAllProducts()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error message": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Product data : ": results,
	})

}

func GetProductbyID(ctx *gin.Context) {
	db := database.GetDB()
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productId := ctx.Param("ID")

	result, err := productService.GetProductByProductID(productId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error message": "Data Not Found",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"ID":          result.ID,
		"Title":       result.Title,
		"Description": result.Description,
		"CreatedAt":   result.CreatedAt,
		"UpdatedAt":   result.UpdatedAt,
	})

}

func UpdateProduct(ctx *gin.Context) {
	db := database.GetDB()
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	//userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)
	Product := models.Product{}
	//userId := uint(userData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Product)
	} else {
		ctx.ShouldBind(&Product)
	}

	productId, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "bad request",
			"message": "params invalid",
		})
	}
	Product.ID = uint(productId)

	err = productService.UpdateProduct(&Product)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad Request",
			"message": "invalid input",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"ID":          Product.ID,
		"CreatedAt":   Product.CreatedAt,
		"UpdatedAt":   Product.UpdatedAt,
		"Title":       Product.Title,
		"Description": Product.Description,
	})
}

func DeleteProduct(ctx *gin.Context) {
	db := database.GetDB()
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productId := ctx.Param("ID")

	rowsAffected, err := productService.DeleteProductByID(productId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error message": err.Error(),
		})
		return
	}
	if rowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error message :": fmt.Sprintf("Product with id %v not found", productId),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":       "product has been deleted successfully",
		"rows_affected": rowsAffected,
	})
}
