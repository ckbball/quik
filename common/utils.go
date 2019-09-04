package common

import (
  "fmt"
  "math/rand"
  "time"

  "github.com/dgrijalva/jwt-go"
  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/gin/binding"
)

const JWTSecretString = "hehehoohoo9wf"

func GenToken(id uint) string {
  jwt_token := jwt.New(jwt.GetSigningMethod("HS256"))

  jwt_token.Claims = jwt.MapClaims{
    "id":  id,
    "exp": time.Now().Add(time.Hour * 24).Unix(),
  }

  token, _ := jwt_token.SignedString([]byte(JWTSecretString))
  return token
}

func Bind(c *gin.Context, obj interface{}) error {
  b := binding.Default(c.Request.Method, c.ContentType())
  return c.ShouldBindWith(obj, b)
}
