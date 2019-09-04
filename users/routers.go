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

func UsersRegistration(c *gin.Context) {
  user := NewUserModelValidator()
  if err := user.Bind(c); err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
    return
  }

  if err := SaveOne(&user.userModel); err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
    return
  }
  c.Set("my_user_model", user.userModel)
  serializer := UserSerializer{c}
  c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
}
