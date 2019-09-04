package users

import (
  "errors"
  "github.com/ckbball/quik/common"
  "github.com/gin-gonic/gin"
  "net/http"
)

func UsersRegister(router *gin.RouterGroup) {
  router.POST("/", UsersRegistration)
  router.POST("/login", UsersLogin)
}
