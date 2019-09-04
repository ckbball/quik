package users

import (
  "github.com/ckbball/quik/common"
  "github.com/gin-gonic/gin"
)

type UserSerializer struct {
  c *gin.Context
}

type UserResponse struct {
  Firstname string `json:"firstname"`
  Lastname  string `json:"lastname"`
  Email     string `json:"email"`
  Status    string `json:"status"`
  Level     string `json:"level"`
}

func (self *UserSerializer) Response() UserResponse {
  user := self.c.MustGet("my_user_model").(UserModel)
  out := UserResponse{
    Firstname: user.Firstname,
    Lastname:  user.Lastname,
    Email:     user.Email,
    Status:    user.Status,
    Level:     user.Level,
    Token:     common.GenToken(user.ID),
  }

  return out
}
