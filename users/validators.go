package users

import (
  "github.com/ckbball/quik/common"
  "github.com/gin-gonic/gin"
)

type UserModelValidator struct {
  User struct {
    Firstname string `json:"first" form:"firstname" binding:"exists,alphanum`
    Lastname  string `json:"last" form:"lastname" binding:"exists,alphanum`
    Email     string `json:"email" form:"email" binding:"exists,email`
    Hash      string `json:"pass" form:"password" binding:"exists,min=8,max=255`
    HasInfo   bool   `json:"info" form:"hasinfo" binding:"exists`
    Status    string `json:"status" form:"status" binding:"alphanum` // this is going to be searching, perusing, locked
    Level     string `json:"level" form:"level" binding:"alphanum`
  } `json:"user"`
  userModel UserModel `json:"-"`
}

func (self *UserModelValidator) Bind(c *gin.Context) error {
  err := common.Bind(c, self)
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
