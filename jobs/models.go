package jobs

import (
//"errors"
//"fmt"
//"github.com/ckbball/quik/common"
//"github.com/jinzhu/gorm"
//"golang.org/x/crypto/bcrypt"
)

// unsure about this model
type JobModel struct {
  ID               int `gorm:"primary_key"`
  CompanyID        int
  Responsibilities string `gorm:"column:responsibilities"`
  Skills           string `gorm:"string"`
}
