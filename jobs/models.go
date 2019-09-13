package jobs

import (
  "errors"
  "fmt"
  "github.com/ckbball/quik/common"
  //"github.com/jinzhu/gorm"
  "github.com/ckbball/quik/companies"
  "golang.org/x/crypto/bcrypt"
)

type JobModel struct {
  ID               int `gorm:"primary_key"`
  CompanyID        int
  Responsibilities string `gorm:"column:responsibilities"`
  Skills           string `gorm:"string"`
}
