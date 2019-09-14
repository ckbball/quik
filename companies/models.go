package companies

import (
  "errors"
  //"fmt"
  "github.com/ckbball/quik/common"
  //"github.com/ckbball/quik/jobs"
  "golang.org/x/crypto/bcrypt"
)

type CompanyModel struct {
  ID      int    `gorm:"primary_key"`
  Name    string `gorm:"column:name"`
  Size    int    `gorm:"column:size"`
  Mission string `gorm:"column:mission"`
  Pass    string `gorm:"column:password;not null"`
  Email   string `gorm:"column:email:unique_index"`
  // maybe have a jobs model here
}

func AutoMigrate() {
  db := common.GetDB()

  db.AutoMigrate(&CompanyModel{})
}

func (u *CompanyModel) setPassword(password string) error {
  if len(password) == 0 {
    return errors.New("password cannot be empty")
  }

  bytePassword := []byte(password)
  passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
  u.Pass = string(passwordHash)
  return nil
}

func SaveOne(data interface{}) error {
  db := common.GetDB()
  err := db.Save(data).Error
  return err
}

func FindOneCompany(condition interface{}) (CompanyModel, error) {
  db := common.GetDB()
  var model CompanyModel
  err := db.Where(condition).First(&model).Error
  return model, err
}

func (u *CompanyModel) checkPassword(password string) error {
  bytePassword := []byte(password)
  byteHashedPassword := []byte(u.Pass)
  return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
