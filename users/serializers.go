package users

import (
  "github.com/ckbball/quik/common"
  "github.com/gin-gonic/gin"
)

//-----  for working with auth'd user's own data
type UserSerializer struct {
  c *gin.Context
}

type UserResponse struct {
  Firstname string `json:"firstname"`
  Lastname  string `json:"lastname"`
  Email     string `json:"email"`
  Status    string `json:"status"`
  Level     string `json:"level"`
  Title     string `json:"title"`
  Token     string `json:"token"`
}

func (self *UserSerializer) Response() UserResponse {
  user := self.c.MustGet("my_user_model").(UserModel)
  out := UserResponse{
    Firstname: user.Firstname,
    Lastname:  user.Lastname,
    Email:     user.Email,
    Status:    user.Status,
    Level:     user.Level,
    Title:     user.Title,
    Token:     common.GenToken(user.ID),
  }
  return out
}

//-------  END ------------------------------

//------- For working with user's data from another account -------------

type UnthSerializer struct {
  c *gin.Context
}

type UnthResponse struct {
  Firstname string `json:"firstname"`
  Lastname  string `json:"lastname"`
  Email     string `json:"email"`
  Status    string `json:"status"`
  Level     string `json:"level"`
  Title     string `json:"title"`
}

func (self *UnthSerializer) Response() UnthResponse {
  user := self.c.MustGet("my_user_model").(UserModel)
  out := UnthResponse{
    Firstname: user.Firstname,
    Lastname:  user.Lastname,
    Email:     user.Email,
    Status:    user.Status,
    Level:     user.Level,
    Title:     user.Title,
  }
  return out
}
