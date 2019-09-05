package users

import (
  // "errors"
  "fmt"
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
    c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
    return
  }

  fmt.Println("check if validator validated user: ", user.userModel)

  if err := SaveOne(&user.userModel); err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
    return
  }
  c.Set("my_user_model", user.userModel)
  serializer := UserSerializer{c}
  c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
}

func UsersLogin(c *gin.Context) {
  login := NewLoginValidator()
  if err := login.Bind(c); err != nil {
    c.JSON(http.StatusForbidden, common.NewError("login", errors.New("Validator error")))
    return
  }

  user, err := FindOneUser(&UserModel{Email: login.userModel.Email})

  if err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("login", errors.New("Email not registered  or invalid password")))
    return
  }

  if user.checkPassword(login.User.Hash) != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("login", errors.New("Email not registered  or invalid password")))
    return
  }
  UpdateContextUserModel(c, user.ID)
  serializer := UserSerializer{c}
  c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}
