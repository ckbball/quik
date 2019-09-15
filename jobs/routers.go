package jobs

import (
  "github.com/ckbball/quik/common"
  "github.com/gin-gonic/gin"
)

// Auth Routes
func Register(router *gin.RouterGroup) {
  // create a job
  router.POST("", JobCreate)
}

func JobCreate(c *gin.Context) {
  // validate job object in http body
  job := NewJobModelValidator()
  if err := job.Bind(c); err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("validator", err))
    return
  }

  if err := SaveOne(&job.jobModel); err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
    return
  }
  serializer := JobSerializer{c, job.jobModel}
  c.JSON(http.StatusCreated, gin.H{"job": serializer.Response()})

  // save job object into db

  // serialize job object into response format

  // send response
}

func JobUpdate(c *gin.Context) {
  // validate job object in http body

  // Check that job exists. if it does continue. if not return job doesn't exist

  // save job object into db

  // serialize job object into response format

  // send response
}
