package users

import (
  "errors"
  "fmt"
  "github.com/ckbball/quik/common"
  "github.com/gin-gonic/gin"
  "net/http"
  "strconv"
)

// No-Auth Routes
func UsersRegister(router *gin.RouterGroup) {
  router.POST("/", UsersRegistration)
  router.POST("/login", UsersLogin)
  router.GET("/:id", UserGet) // have to move this to authenticated routes later
}

// Auth Routes
func Register(router *gin.RouterGroup) {
  // router.GET("/:id", UserGet) // to be used by another person looking at someone else's profile or by applications of a job view
  router.POST("", UserUpdate)
  router.GET("/profiles/:id", UserInfoGet)
  router.POST("/profiles", UserInfoCreate)
}

func UsersRegistration(c *gin.Context) {
  user := NewUserModelValidator()
  if err := user.Bind(c); err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("validator", err))
    return
  }

  fmt.Println("check if validator validated user: ", user.userModel)

  if err := SaveOne(&user.userModel); err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
    return
  }
  c.Set("my_user_model", user.userModel)
  serializer := UserSerializer{c}
  c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
}

func UsersLogin(c *gin.Context) {
  login := NewLoginValidator()
  if err := login.Bind(c); err != nil {
    c.JSON(http.StatusForbidden, common.NewError("login", errors.New("Validator error")))
    return
  }

  fmt.Println("Checking login validator binding: --> ", login)

  user, err := FindOneUser(&UserModel{Email: login.userModel.Email}) //models.go function

  // sending error with token and pass
  if err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("login", errors.New("DB: Email not registered or invalid password")))
    return
  }

  if user.checkPassword(login.User.Hash) != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("login", errors.New("Check: Email not registered  or invalid password")))
    return
  }
  UpdateContextUserModel(c, user.ID)
  serializer := UserSerializer{c}
  c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}

// Retrieves a user instance by user_id
func UserGet(c *gin.Context) {
  id := c.Param("id")
  Id, err := strconv.Atoi(id)
  if err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("user get", errors.New("DB: Invalid Id")))
    return
  }

  user, err := FindOneUser(&UserModel{ID: Id}) // models.go function

  if err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("user get", errors.New("DB: Invalid Id")))
    return
  }

  UpdateContextUserModel(c, user.ID)
  serializer := UnthSerializer{c} // serializer struct for working with a different user's data.
  c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}

// Receives the new user struct and updates the user in the db
// PUT -
func UserUpdate(c *gin.Context) {
  myUser := c.MustGet("my_user_model").(UserModel)
  user := NewUserModelValidatorFillWith(myUser)
  if err := user.Bind(c); err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("validator", err))
    return
  }

  fmt.Println("check if validator validated user: ", user.userModel)

  user.userModel.ID = myUser.ID
  if err := myUser.Update(user.userModel); err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
    return
  }
  UpdateContextUserModel(c, myUser.ID)
  serializer := UserSerializer{c}
  c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
}

// Get a specific user's info
func UserInfoGet(c *gin.Context) {
  id := c.Param("id")
  Id, err := strconv.Atoi(id)
  if err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("profile get", errors.New("DB: Invalid Id")))
    fmt.Println(err)
    return
  }

  user, err := FindOneProfile(&Profile{UserID: Id}) // models.go function

  if err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("profile get", errors.New("DB: Invalid Id")))
    fmt.Println(err)
    return
  }

  UpdateContextProfile(c, user.ID)
  serializer := ProfileSerializer{c} // serializer struct for working with a user's profile.
  c.JSON(http.StatusOK, gin.H{"profile": serializer.Response()})
}

func UserInfoCreate(c *gin.Context) {
  // validator
  var profile = NewProfileValidator()

  //bind
  if err := profile.Bind(c); err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("validator", err))
    return
  }
  /*
    user, err := c.MustGet("my_user_model").(UserModel)
    if !err {
      fmt.Println("POST - /profile : error in getting user, ", err)
    }
    fmt.Println("got user")*/

  // check if user already has profile. if yes return error. if no continue
  /*
     if _, err := FindProfileByUser(user); err == nil {
       c.JSON(http.StatusBadRequest, gin.H{"msg": "Profile already exists"})
     }*/

  fmt.Println("users/routers/147 - check if validator validated profile: ", profile.profileModel)

  if err := SaveOne(&profile.profileModel); err != nil {
    c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
    return
  }
  c.Set("my_profile_model", profile.profileModel)
  serializer := ProfileSerializer{c}
  c.JSON(http.StatusCreated, gin.H{"profile": serializer.Response()})
}
