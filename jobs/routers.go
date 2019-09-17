package jobs

import (
  "github.com/ckbball/quik/common"
  "github.com/gin-gonic/gin"
  "net/http"
)

// Auth Routes
func Register(router *gin.RouterGroup) {
  // create a job
  router.POST("", JobCreate)
  router.POST("/", JobUpdate)
  router.DELETE("/:id", JobDelete)
}

func JobCreate(c *gin.Context) {
  // validate job object in http body
  job := NewJobModelValidator()

  // Bind to input
  if err := job.Bind(c); err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("validator", err))
    return
  }

  // Create new job in db
  if err := SaveOne(&job.jobModel); err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
    return
  }
  // Create and return response
  serializer := JobSerializer{c, &job.jobModel}
  c.JSON(http.StatusCreated, gin.H{"job": serializer.Response()})
}

func JobUpdate(c *gin.Context) {
  // validate job object in http body
  job := NewJobModelValidator()

  if err := job.Bind(c); err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("validator", err))
    return
  }
  // Check that job exists. if it does continue. if not return job doesn't exist
  if db_job, err := FindOneJob(&JobModel{ID: job.ID}); err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
    return
  }

  // save job object into db
  if err := db_job.Update(job.jobModel); err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
    return
  }
  // serialize job object into response format
  serializer := JobSerializer{c, &job.jobModel}
  c.JSON(http.StatusCreated, gin.H{"job": serializer.Response()})
  // send response
}

func JobUpdate(c *gin.Context) {
  // validate job object in http body
  job := NewJobModelValidator()

  if err := job.Bind(c); err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("validator", err))
    return
  }
  // Check that job exists. if it does continue. if not return job doesn't exist
  if db_job, err := FindOneJob(&JobModel{ID: job.ID}); err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
    return
  }

  // save job object into db
  if err := db_job.Update(job.jobModel); err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
    return
  }
  // serialize job object into response format
  serializer := JobSerializer{c, &job.jobModel}
  c.JSON(http.StatusCreated, gin.H{"job": serializer.Response()})
  // send response
}
