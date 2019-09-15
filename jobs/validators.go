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

  fmt.Println("Check binding in jobs: ", self.Job)
  fmt.Println()

  self.jobModel.CompanyID = self.Job.CompanyID
  self.jobModel.Responsibilities = self.Job.Responsibilities
  self.jobModel.Skills = self.Job.Skills
  return nil
}

func NewJobModelValidator() JobModelValidator {
  validator := JobModelValidator{}
  return validator
}

func NewJobModelValidatorFillWith(job JobModel) JobModelValidator {
  out := NewJobModelValidator()
  out.Job.CompanyID = job.CompanyID
  out.Job.Responsibilities = job.Responsibilities
  out.Job.Skills = job.Skills

  return out
}
