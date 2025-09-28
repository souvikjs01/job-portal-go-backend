package controllers

// import (
// 	"job_portal/packages/config"
// 	"job_portal/packages/models"
// 	"job_portal/packages/store"
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt/v5"
// 	"golang.org/x/crypto/bcrypt"
// )

// var jwt_secret string

// func init() {
// 	config.LoadEnv()
// 	jwt_secret = config.GetEnv("JWT_SECRET")
// }

// func RegisterUser(c *gin.Context) {
// 	var user models.User
// 	err := c.ShouldBindJSON(&user)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Error registering user",
// 		})
// 		return
// 	}

// 	// hashed the password:
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "Error while encrypting the password",
// 		})
// 	}

// 	user.Password = string(hashedPassword)

// 	// generate jwt
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"user_id":  user.ID,
// 		"is_admin": user.IsAdmin,
// 		"exp":      time.Now().Add(time.Hour * 72).Unix(),
// 	})

// 	tokenStr, err := token.SignedString([]byte(jwt_secret)) // SignedString() expects a []byte when using SigningMethodHS256

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "Could not generate token",
// 		})
// 		return
// 	}

// 	if err := store.DB.Create(&user).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
// 		return
// 	}

// 	// Return token
// 	c.JSON(http.StatusCreated, gin.H{
// 		"message": "User registered successfully",
// 		"token":   tokenStr,
// 	})
// }

// var input struct {
// 	Email    string `json:"email" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// }

// func LogInUser(c *gin.Context) {
// 	err := c.ShouldBindBodyWithJSON(&input)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Invalid credentials",
// 		})
// 		return
// 	}

// 	var user models.User
// 	if err := store.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"error": "Invalid credentials",
// 		})
// 		return
// 	}

// 	// compare the hashed password
// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"error": "Invalid credentials",
// 		})
// 		return
// 	}

// 	// generate jwt token:
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"user_id":  user.ID,
// 		"is_admin": user.IsAdmin,
// 		"exp":      time.Now().Add(time.Hour * 72).Unix(),
// 	})

// 	tokenStr, err := token.SignedString([]byte(jwt_secret))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
// 		return
// 	}

// 	// Return token
// 	c.JSON(http.StatusCreated, gin.H{
// 		"message": "Login successfull",
// 		"token":   tokenStr,
// 	})
// }

// func GetAllUsers(c *gin.Context) {
// 	var users []models.User

// 	result := store.DB.Find(&users)
// 	if result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "failed to fetch the users",
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"users": users,
// 	})
// }

// func GetUserById(c *gin.Context) {
// 	userIdStr := c.Param("id")
// 	userId, err := strconv.Atoi(userIdStr)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"error": "Invalid user ID",
// 		})
// 		return
// 	}

// 	var user models.User
// 	if err := store.DB.First(&user, userId).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"error": "User not found",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"user": user,
// 	})
// }

// // delete user by its id only admin can do it
// func RemoveUserById(c *gin.Context) {
// 	userIdStr := c.Param("id")

// 	userId, err := strconv.Atoi(userIdStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Invalid user ID",
// 		})
// 		return
// 	}

// 	// check if the user is admin or not:
// 	isAdminVal, exist := c.Get("is_admin")
// 	if !exist {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"error": "Unauthorized",
// 		})
// 		return
// 	}

// 	isAdmin, ok := isAdminVal.(bool)
// 	if !ok || !isAdmin {
// 		c.JSON(http.StatusForbidden, gin.H{
// 			"error": "Only admin can delete jobs",
// 		})
// 		return
// 	}

// 	var user models.User
// 	result := store.DB.Delete(&user, userId)

// 	if result.RowsAffected == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"error": "Job not found",
// 		})
// 		return
// 	}

// 	if result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "Failed to delete user",
// 		})
// 		return
// 	}

// 	// success
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "User deleted successfully",
// 	})
// }
