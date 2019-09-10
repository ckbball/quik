package users

import (
  "errors"
  //"fmt"
  "github.com/ckbball/quik/common"
  //"github.com/jinzhu/gorm"
  "golang.org/x/crypto/bcrypt"
)

type UserModel struct {
  ID        int     ` gorm:"primary_key"`
  Firstname string  ` gorm:"column:firstname"`
  Lastname  string  `gorm:"column:lastname"`
  Email     string  `gorm:"column:email;unique_index"`
  Hash      string  ` gorm:"column:password;not null"`
  HasInfo   bool    ` gorm:"column:hasinfo"`
  Status    string  `gorm:"column:status"` // this is going to be searching, perusing, locked
  Level     string  `gorm:"column:level"`  // this is going to be entry, mid, senior
  Title     string  `gorm:"column:title"`  // this is going to be frontend, backend, etc.
  Profile   Profile `gorm:"foreignkey:UserID"`
}

type Profile struct {
  Roles      []Role      `gorm:"foreignkey:InfoID;column:roles"`      // this is going to be backend, frontend, full stack, mobile,
  Frameworks []Framework `gorm:"foreignkey:InfoID;column:frameworks"` // this is going to be all frameworks, front and back, that user knows meaning they built a project with it
  DB         []DB        `gorm:"foreignkey:InfoID;column:db"`         // this is all dbs that a user has built a project with
  Front      []Front     `gorm:"foreignkey:InfoID;column:front"`      // languages and methods for front end user has built a project with
  Back       []Back      `gorm:"foreignkey:InfoID;column:back"`       // languages and methods for backend user has built a project with
  Extra      []Extra     `gorm:"foreignkey:InfoID;column:extra"`      // dont know what should be here
  DevOps     []Devops    `gorm:"foreignkey:InfoID;column:devops"`     // CI/CD tools, other things idk
  Cloud      []Cloud     `gorm:"foreignkey:InfoID;column:cloud"`      // which cloud tools and platforms user has made a project with
  ID         int         `gorm:"primary_key`
  UserID     int         `gorm:"column:userid"`
}

type Role struct {
  ID     int    `gorm:"primary_key"`
  Name   string `json: "name" gorm:"column:name"`
  Years  int    `json: "years" gorm:"column:years"`
  InfoID int
}

func (model *Profile) setRoles(roles []Role) error {
  db := common.GetDB()
  var roleList []Role
  for _, role := range roles {
    var roleModel Role
    err := db.FirstOrCreate(&roleModel, Role{Name: role.Name, Years: role.Years})
    if err != nil {
      return err
    }
    roleList = append(roleList, roleModel)
  }
  model.Roles = roleList
  return nil
}

type Framework struct {
  ID     int    ` gorm:"primary_key"`
  Name   string `json: "name" gorm:"column:name"`
  Years  int    `json: "years" gorm:"column:years"`
  InfoID int    `gorm:"column:infoid"`
}

type DB struct {
  ID     int    ` gorm:"primary_key"`
  Name   string `json: "name" gorm:"column:name"`
  Years  int    `json: "years" gorm:"column:years"`
  InfoID int    `gorm:"column:infoid"`
}

type Front struct {
  ID     int    ` gorm:"primary_key"`
  Name   string `json: "name" gorm:"column:name"`
  Years  int    `json: "years" gorm:"column:years"`
  InfoID int    `gorm:"column:infoid"`
}

type Back struct {
  ID     int    ` gorm:"primary_key"`
  Name   string `json: "name" gorm:"column:name"`
  Years  int    `json: "years" gorm:"column:years"`
  InfoID int    `gorm:"column:infoid"`
}

type Extra struct {
  ID     int    ` gorm:"primary_key"`
  Name   string `json: "name" gorm:"column:name"`
  Years  int    `json: "years" gorm:"column:years"`
  InfoID int    `gorm:"column:infoid"`
}

type Devops struct {
  ID     int    ` gorm:"primary_key"`
  Name   string `json: "name" gorm:"column:name"`
  Years  int    `json: "years" gorm:"column:years"`
  InfoID int    `gorm:"column:infoid"`
}

type Cloud struct {
  ID     int    ` gorm:"primary_key"`
  Name   string `json: "name" gorm:"column:name"`
  Years  int    `json: "years" gorm:"column:years"`
  InfoID int    `gorm:"column:infoid"`
}

func AutoMigrate() {
  db := common.GetDB()

  db.AutoMigrate(&UserModel{})
  db.AutoMigrate(&Profile{})
  db.AutoMigrate(&Role{})
  db.AutoMigrate(&Framework{})
  db.AutoMigrate(&DB{})
  db.AutoMigrate(&Front{})
  db.AutoMigrate(&Back{})
  db.AutoMigrate(&Extra{})
  db.AutoMigrate(&Devops{})
  db.AutoMigrate(&Cloud{})
}

// There will be multiple relations table with id, user_id, <a field of Profile>_id,

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

func FindOneProfile(condition interface{}) (Profile, error) {
  db := common.GetDB()
  var model Profile
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

// func Find
// Methods to do
// get Profile
// create Profile
// update Profile
