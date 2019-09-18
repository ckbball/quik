package jobs

import (
  "github.com/ckbball/quik/common"
  "github.com/gin-gonic/gin"
  "net/http"
)

// No Auth Routes
func JobsRegister(router *gin.RouterGroup) {
  router.GET("", JobFilter)
  router.GET("/:id", JobGetByID)
}

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

func JobDelete(c *gin.Context) {
  id := c.Param("id")
  Id, err := strconv.Atoi(id)
  if err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("job delete", errors.New("DB: Invalid Id")))
    fmt.Println(err)
    return
  }

  err := DeleteJobModel(&JobModel{ID: Id})
  if err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("job delete", errors.New("DB: Invalid Id")))
    fmt.Println(err)
    return
  }
  c.JSON(http.StatusOK, gin.H{"job": "Delete success"})
}

// query params:
// q = things to search for etc. python, backend, frontend, react
// l = location, city, state,
// c = company name
// limit = limit number of jobs
// offset = page offset
func JobFilter(c *gin.Context) {
  // need serializer and models funcs
  location := c.Query("location")
  limit := c.Query("limit")
  offset := c.Query("offset")

  j, jobCount, err := FilteredJobs(location, limit, offset)
  if err != nil {
    c.JSON(http.StatusNotFound, common.NewError("job filter", errors.New("DB: Invalid params")))
    return
  }
  serializer := JobsSerializer{c, j}
  c.JSON(http.StatusOK, gin.H{"jobs": serializer.Response(), "jobsCount": jobCount})
}
