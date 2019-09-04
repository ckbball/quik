package users

import (
  "errors"
  "fmt"
  "github.com/ckbball/quik/common"
  "github.com/jinzhu/gorm"
  "golang.org/x/crypto/bcrypt"
)

type UserModel struct {
  ID        int    `json:"id" gorm:"primary_key"`
  Firstname string `json:"first" gorm:"column:firstname"`
  Lastname  string `json:"last" gorm:"column:lastname"`
  Email     string `json:"email" gorm:"column:email;unique_index"`
  Hash      string `json:"pass" gorm:"column:password;not null"`
  HasInfo   bool   `json:"info" gorm:"column:hasinfo"`
  Status    string `gorm:"column:status"` // this is going to be searching, perusing, locked
  Level     string `gorm:"column:level"`  // this is going to be entry, mid, senior
  // Info      UserInfo `json:"info"`  should this be in UserModel or should it just be a separate table that I also grab
}

type UserInfo struct {
  Roles      []string // this is going to be backend, frontend, full stack, mobile,
  Frameworks []string // this is going to be all frameworks, front and back, that user knows meaning they built a project with it
  DB         []string // this is all dbs that a user has built a project with
  Front      []string // languages and methods for front end user has built a project with
  Back       []string // languages and methods for backend user has built a project with
  Extra      []string // dont know what should be here
  DevOps     []string // CI/CD tools, other things idk
  Cloud      []string // which cloud tools and platforms user has made a project with
  ID         int
}

func AutoMigrate() {
  db := common.GetDB()

  db.AutoMigrate(&UserModel{})
  db.AutoMigrate(&UserInfo{})
}

// There will be multiple relations table with id, user_id, <a field of userInfo>_id,

// -------- HELPER FUNCTIONS BEGIN --------------------------------------

func (u *UserModel) setPassword(password string) error {
  if len(password) == 0 {
    return errors.New("password cannot be empty")
  }

  bytePassword := []byte(password)
  passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
  u.Hash = string(passwordHash)
  return nil
}

func (u *UserModel) checkPassword(password string) error {
  bytePassword := []byte(password)
  byteHashedPassword := []byte(u.Hash)
  return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func FindOneUser(condition interface{}) (UserModel, error) {
  db := common.GetDB()
  var model UserModel
  err := db.Where(condition).First(&model).Error
  return model, err
}

func SaveOne(data interface{}) error {
  db := common.GetDB()
  err := db.Save(data).Error
  return err
}

func (model *UserModel) Update(data interface{}) error {
  db := common.GetDB()
  err := db.Model(model).Update(data).Error
  return err
}

// Methods to do
// get userInfo
// create userinfo
// update userinfo
