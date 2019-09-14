package companies

import (
  //"errors"
  //"fmt"
  "github.com/ckbball/quik/common"
  //"github.com/ckbball/quik/jobs"
  "golang.org/x/crypto/bcrypt"
)

type CompanyModel struct {
  ID      int    `gorm:"primary_key"`
  Name    string `gorm:"column:name"`
  Size    int    `gorm:"column:size:unique_index"`
  Mission string `gorm:"column:mission"`
  Hash    string `gorm:"column:pass"`
  Email   string `gorm:"column:email"`
  // maybe have a jobs model here
}

func AutoMigrate() {
  db := common.GetDB()

  db.AutoMigrate(&CompanyModel{})
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
  byteHashedPassword := []byte(u.Hash)
  return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
