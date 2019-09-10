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

type ProfileSerializer struct {
  c *gin.Context
}

type ProfileResponse struct {
  Roles      []Role      `json: "roles"`      // this is going to be backend, frontend, full stack, mobile,
  Frameworks []Framework `json: "frameworks"` // this is going to be all frameworks, front and back, that user knows meaning they built a project with it
  DB         []DB        `json: "db"`         // this is all dbs that a user has built a project with
  Front      []Front     `json: "front"`      // languages and methods for front end user has built a project with
  Back       []Back      `json: "back"`       // languages and methods for backend user has built a project with
  Extra      []Extra     `json: "extra"`      // dont know what should be here
  DevOps     []Devops    `json: "devops"`     // CI/CD tools, other things idk
  Cloud      []Cloud     `json: "cloud"`      // which cloud tools and platforms user has made a project with
  ID         int         `json: "id"`
  UserID     int         `json: "userid"`
}

func (self *ProfileSerializer) Response() ProfileResponse {
  user := self.c.MustGet("my_profile_model").(Profile)
  out := ProfileResponse{
    Roles:      user.Roles,
    Frameworks: user.Frameworks,
    DB:         user.DB,
    Front:      user.Front,
    Back:       user.Back,
    Extra:      user.Extra,
    DevOps:     user.DevOps,
    Cloud:      user.Cloud,
    ID:         user.ID,
    UserID:     user.UserID,
  }

  return out
}
