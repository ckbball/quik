package users

import (
  "fmt"
  "github.com/ckbball/quik/common"
  "github.com/gin-gonic/gin"
)

type UserModelValidator struct {
  User struct {
    Firstname string  `json:"firstname" form:"firstname" binding:"exists,alphanum"`
    Lastname  string  `json:"lastname" form:"lastname" binding:"exists,alphanum"`
    Email     string  `json:"email" form:"email" binding:"exists,email"`
    Hash      string  `json:"pass" form:"password" binding:"exists,min=8,max=255"`
    HasInfo   bool    `json:"info" form:"hasinfo" binding:"exists"`
    Status    string  `json:"status" form:"status" binding:"alphanum"` // this is going to be searching, perusing, locked
    Level     string  `json:"level" form:"level" binding:"alphanum"`
    Title     string  `json:"title" form:"title" binding:"alphanum"`
    Profile   Profile `json:"profile" form:"profile"`
  } `json:"user"`
  userModel UserModel `json:"-"`
}

func (self *UserModelValidator) Bind(c *gin.Context) error {
  err := common.Bind(c, self)

  fmt.Println("/users/validators-26: Check if common.Bind() binded the body to User properly: ", self.User)
  if err != nil {
    return err
  }

  self.userModel.Firstname = self.User.Firstname
  self.userModel.Lastname = self.User.Lastname
  self.userModel.Email = self.User.Email
  self.userModel.HasInfo = self.User.HasInfo

  fmt.Println("/users/validators-36: Check if copying User to userModel correctly: ", self.userModel)

  if self.User.Hash != common.JWTSecretString {
    self.userModel.setPassword(self.User.Hash)
  }
  if self.User.Status != "" {
    self.userModel.Status = self.User.Status
  }
  if self.User.Level != "" {
    self.userModel.Level = self.User.Level
  }
  if self.User.Title != "" {
    self.userModel.Title = self.User.Title
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
  out.User.Title = user.Title

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

// profile validator goes down here
type ProfileValidator struct {
  Profile struct {
    Roles      []Role      `json:"roles" form:"roles"`
    Frameworks []Framework `json:"frameworks" form:"frameworks"`
    DB         []DB        `json:"db" form:"db"`
    Front      []Front     `json:"front" form:"front"`
    Back       []Back      `json:"back" form:"back"`
    Extra      []Extra     `json:"extra" form:"extra"`
    DevOps     []Devops    `json:"devops" form:"devops"`
    Cloud      []Cloud     `json:"cloud" form:"cloud"`
    ID         int         `json:"id" form:"id" `
    UserID     int         `json:"userid" form:"userid"`
  } `json: "profile"`
  profileModel Profile `json: "-"`
}

func (self *ProfileValidator) Bind(c *gin.Context) error {
  err := common.Bind(c, self)

  fmt.Println("/users/validators-116: Check if common.Bind() binded the body to Profile properly: ", self.Profile)
  if err != nil {
    return err
  }

  self.profileModel.Roles = self.Profile.Roles
  self.profileModel.Frameworks = self.Profile.Frameworks
  self.profileModel.DB = self.Profile.DB
  self.profileModel.Back = self.Profile.Back
  self.profileModel.Extra = self.Profile.Extra
  self.profileModel.DevOps = self.Profile.DevOps
  self.profileModel.Cloud = self.Profile.Cloud
  self.profileModel.Front = self.Profile.Front
  self.profileModel.ID = self.Profile.ID
  self.profileModel.UserID = self.Profile.UserID

  fmt.Println("/users/validators-126: Check if copying Profile to profileModel correctly: ", self.profileModel)
  return nil
}

func NewProfileValidator() ProfileValidator {
  profileValidator := ProfileValidator{}
  return profileValidator
}
