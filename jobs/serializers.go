package jobs

import (
  //"github.com/ckbball/quik/common"
  "github.com/gin-gonic/gin"
  "time"
)

type JobSerializer struct {
  c   *gin.Context
  job *JobModel
}

type JobResponse struct {
  CompanyID        int       `json:"companyid"`
  Responsibilities string    `json:"responsibilities"`
  Skills           string    `json:"skills"`
  CreatedAt        time.Time `json:"createdat"`
  // need to add job here later or maybe a companyprofileresponse
}

func (self *JobSerializer) Response() JobResponse {
  job := self.job
  out := JobResponse{
    CompanyID:        job.CompanyID,
    Responsibilities: job.Responsibilities,
    Skills:           job.Skills,
    CreatedAt:        job.CreatedAt,
  }
  return out
}
