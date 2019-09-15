package jobs

import (
  //"fmt"
  "github.com/ckbball/quik/common"
  "github.com/gin-gonic/gin"
)

type JobModelValidator struct {
  Job struct {
    CompanyID        id     `json:"companyid" form:"companyid" binding:"exists"`
    Responsibilities string `json:"responsibilities" form:"responsibilities" binding:"exists"`
    Skills           string `json:"skills" form:"skills" binding:"exists"` //maybe give this a max
  } `json:"job"`
  jobModel JobModel `json:"-"`
}

func (self *JobModelValidator) Bind(c *gin.Context) error {
  err := common.Bind(c, self)

  if err != nil {
    return err
  }

  self.jobModel.CompanyID = self.Job.CompanyID
  self.jobModel.Responsibilities = self.Job.Responsibilities
  self.jobModel.Skills = self.Job.Skills
  return nil
}

func NewCompanyModelValidator() CompanyModelValidator {
  validator := CompanyModelValidator{}
  return validator
}

func NewCompanyModelValidatorFillWith(company CompanyModel) CompanyModelValidator {
  out := NewCompanyModelValidator()
  out.Company.Name = company.Name
  out.Company.Size = company.Size
  out.Company.Mission = company.Mission
  out.Company.Pass = company.Pass
  out.Company.Email = company.Email

  return out
}
