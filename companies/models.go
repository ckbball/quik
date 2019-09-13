package companies

import (
  "errors"
  "fmt"
  "github.com/ckbball/quik/common"
  //"github.com/jinzhu/gorm"
  "github.com/ckbball/quik/jobs"
  "golang.org/x/crypto/bcrypt"
)

type CompanyModel struct {
  ID      int             `gorm:"primary_key"`
  Name    string          `gorm:"column:name"`
  Size    int             `gorm:"column:size:unique_index"`
  Mission string          `gorm:"column:mission"`
  Hash    string          `gorm:"cloumn:pass"`
  Jobs    []jobs.JobModel `gorm:"foreignkey:CompanyID;column:jobs"`
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

func FindOneCompany(data interface{}) (CompanyModel, error) {
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
