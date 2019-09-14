package jobs

import ()

// Auth Routes
func Register(router *gin.RouterGroup) {
  // create a job
  router.POST("", JobCreate)
}

func JobCreate(c *gin.Context) {
  // validate job object in http body

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
