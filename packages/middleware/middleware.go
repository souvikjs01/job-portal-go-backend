package middleware

// import (
// 	"job_portal/packages/config"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt/v5"
// )

// var jwt_secret string

// func init() {
// 	config.LoadEnv()
// 	jwt_secret = config.GetEnv("JWT_SECRET")
// }

// func Authenticated(c *gin.Context) {
// 	authHeader := c.GetHeader("Authorization")
// 	if authHeader == "" {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"error": "Authorization token is required",
// 		})
// 		c.Abort()
// 		return
// 	}

// 	token, err := jwt.Parse(authHeader, func(t *jwt.Token) (interface{}, error) {
// 		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, jwt.ErrInvalidKey
// 		}
// 		return []byte(jwt_secret), nil
// 	})

// 	if err != nil || !token.Valid {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"error": "Invalid or expired token: " + err.Error(),
// 		})
// 		c.Abort()
// 		return
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if ok && token.Valid {
// 		c.Set("user_id", claims["user_id"])
// 		c.Set("is_admin", claims["is_admin"])
// 		c.Next()
// 	} else {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"error": "Invalid token claims",
// 		})
// 		c.Abort()
// 	}
// }
