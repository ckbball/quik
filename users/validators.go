package users

import (
  "fmt"
  "github.com/ckbball/quik/common"
  "github.com/gin-gonic/gin"
)

type UserModelValidator struct {
  User struct {
    Firstname string `json:"firstname" form:"firstname" binding:"exists,alphanum"`
    Lastname  string `json:"lastname" form:"lastname" binding:"exists,alphanum"`
    Email     string `json:"email" form:"email" binding:"exists,email"`
    Hash      string `json:"pass" form:"password" binding:"exists,min=8,max=255"`
    HasInfo   bool   `json:"info" form:"hasinfo" binding:"exists"`
    Status    string `json:"status" form:"status" binding:"alphanum"` // this is going to be searching, perusing, locked
    Level     string `json:"level" form:"level" binding:"alphanum"`
  } `json:"user"`
  userModel UserModel `json:"-"`
}

func (self *UserModelValidator) Bind(c *gin.Context) error {
  err := common.Bind(c, self)

  fmt.Println("Check if common.Bind() binded the body to User properly: ", self.User)
  if err != nil {
    return nil
  }

  self.userModel.Firstname = self.User.Firstname
  self.userModel.Lastname = self.User.Lastname
  self.userModel.Email = self.User.Email
  self.userModel.HasInfo = self.User.HasInfo

  if self.User.Hash != common.JWTSecretString {
    self.userModel.setPassword(self.User.Hash)
  }
  if self.User.Status != "" {
    self.userModel.Status = self.User.Status
  }
  if self.User.Level != "" {
    self.userModel.Level = self.User.Level
  }
  return nil
}

func NewUserModelValidator() UserModelValidator {
  userModelValidator := UserModelValidator{}
  return userModelValidator
}

func NewUserModelValidatorFillWith(user UserModel) UserModelValidator {
  out := NewUserModelValidator()
  out.User.Firstname = user.Firstname
  out.User.Lastname = user.Lastname
  out.User.Email = user.Email
  out.User.HasInfo = user.HasInfo
  out.User.Hash = user.Hash
  out.User.Status = user.Status
  out.User.Level = user.Level

  return out
}

type LoginValidator struct {
  User struct {
    Email string `json:"email" form:"email" binding:"exists,email`
    Hash  string `json:"pass" form:"password" binding:"exists,min=8,max=255`
  } `json:"user"`
  userModel UserModel `json:"-"`
}

func (self *LoginValidator) Bind(c *gin.Context) error {
  err := common.Bind(c, self)
  if err != nil {
    return err
  }

  self.userModel.Email = self.User.Email
  return nil
}

func NewLoginValidator() LoginValidator {
  loginValidator := LoginValidator{}
  return loginValidator
}
