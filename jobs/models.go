package jobs

import (
  //"errors"
  //"fmt"
  "github.com/ckbball/quik/common"
  "github.com/jinzhu/gorm"
  //"golang.org/x/crypto/bcrypt"
)

// unsure about this model
type JobModel struct {
  gorm.Model
  CompanyID        int    `gorm:"column:company_id"`
  Responsibilities string `gorm:"column:responsibilities"`
  Skills           string `gorm:"column:skills"`
  Location         string `gorm:"colum:location"` // city start as bay area only
}

func AutoMigrate() {
  db := common.GetDB()

  db.AutoMigrate(&JobModel{})
}

func SaveOne(data interface{}) error {
  db := common.GetDB()
  err := db.Save(data).Error
  return err
}

func FindOneJob(condition interface{}) (JobModel, error) {
  db := common.GetDB()
  var model JobModel
  err := db.Where(condition).First(&model).Error
  return model, err
}

func (model *JobModel) Update(data interface{}) error {
  db := common.GetDB()
  err := db.Model(model).Update(data).Error
  return err
}

func DeleteJobModel(condition interface{}) error {
  db := common.GetDB()
  err := db.Where(condition).Delete(JobModel{}).Error
  return err
}

func FilteredJobs(query, location, company, limit, offset string) ([]JobModel, int, error) {
  db := common.GetDB()
  var models []JobModel
  var count int

  offset_int, err := strconv.Atoi(offset)
  if err != nil {
    offset_int = 0
  }

  limit_int, err := strconv.Atoi(limit)
  if err != nil {
    limit_int = 20
  }

  tx := db.Begin()

}
