package users

import (
  "database/sql"
  "fmt"
  "github.com/ckbball/quik/common"
)

type UserModel struct {
  ID        int    `json:"id"`
  Firstname string `json:"first"`
  Lastname  string `json:"last"`
  Email     string `json:"email"`
  Hash      string `json:"pass"`
  HasInfo   bool   `json:"info"`
  // Info      UserInfo `json:"info"`  should this be in UserModel or should it just be a separate table that I also grab
}

type UserInfo struct {
  Status     string   // this is going to be searching, perusing, locked
  Level      string   // this is going to be entry, mid, senior
  Roles      []string // this is going to be backend, frontend, full stack, mobile,
  Frameworks []string // this is going to be all frameworks, front and back, that user knows meaning they built a project with it
  DB         []string // this is all dbs that a user has built a project with
  Front      []string // languages and methods for front end user has built a project with
  Back       []string // languages and methods for backend user has built a project with
  Extra      []string // dont know what should be here
  DevOps     []string // CI/CD tools, other things idk
  Cloud      []string // which cloud tools and platforms user has made a project with
  UserId     int      // the id that this info block belongs to
}

// -------- HELPER FUNCTIONS BEGIN --------------------------------------

func UserValidatorToModel(newUser UserModelValidator) (UserModel, error) {
  u := UserModel{}

  u.ID = newUser.User.ID
  u.Firstname = newUser.User.Firstname
  u.Lastname = newUser.User.Lastname
  u.Email = newUser.User.Email
  u.Hash = newUser.User.Hash
  u.HasInfo = newUser.User.HasInfo

  return u, nil
}

func scanUsers(s *sql.Rows) (*UserModel, error) {
  var (
    ID        int
    Firstname sql.NullString
    Lastname  sql.NullString
    Email     sql.NullString
    Hash      sql.NullString
    HasInfo   sql.NullBool
  )

  if err := s.Scan(&ID, &Firstname, &Lastname, &Email, &Hash, &HasInfo); err != nil {
    return nil, err
  }

  id := int(ID)

  user := &UserModel{
    ID:        id,
    Firstname: Firstname.String,
    Lastname:  Lastname.String,
    Email:     Email.String,
    Hash:      Hash.String,
    HasInfo:   HasInfo.String,
  }

  return user, nil
}

func scanUser(s *sql.Row) (*UserModel, error) {
  var (
    ID        int
    Firstname sql.NullString
    Lastname  sql.NullString
    Email     sql.NullString
    Hash      sql.NullString
    HasInfo   sql.NullBool
  )

  if err := s.Scan(&ID, &Firstname, &Lastname, &Email, &Hash, &HasInfo); err != nil {
    return nil, err
  }

  id := int(ID)

  user := &UserModel{
    ID:        id,
    Firstname: Firstname.String,
    Lastname:  Lastname.String,
    Email:     Email.String,
    Hash:      Hash.String,
    HasInfo:   HasInfo.String,
  }

  return user, nil
}

func scanUserInfos(s *sql.Rows) (*UserModel, error) {
  Status     string   // this is going to be searching, perusing, locked
  Level      string   // this is going to be entry, mid, senior
  Roles      []string // this is going to be backend, frontend, full stack, mobile,
  Frameworks []string // this is going to be all frameworks, front and back, that user knows meaning they built a project with it
  DB         []string // this is all dbs that a user has built a project with
  Front      []string // languages and methods for front end user has built a project with
  Back       []string // languages and methods for backend user has built a project with
  Extra      []string // dont know what should be here
  DevOps     []string // CI/CD tools, other things idk
  Cloud      []string // which cloud tools and platforms user has made a project with
  UserId     int      // the id that this info block belongs to
  var (
    UserID        int
    Status sql.NullString
    Level  sql.NullString
    Email     sql.NullString
    Hash      sql.NullString
    HasInfo   sql.NullBool
  )

  if err := s.Scan(&ID, &Firstname, &Lastname, &Email, &Hash, &HasInfo); err != nil {
    return nil, err
  }

  id := int(ID)

  user := &UserModel{
    ID:        id,
    Firstname: Firstname.String,
    Lastname:  Lastname.String,
    Email:     Email.String,
    Hash:      Hash.String,
    HasInfo:   HasInfo.String,
  }

  return user, nil
}

func scanUserInfo(s *sql.Row) (*UserModel, error) {
  var (
    ID        int
    Firstname sql.NullString
    Lastname  sql.NullString
    Email     sql.NullString
    Hash      sql.NullString
    HasInfo   sql.NullBool
  )

  if err := s.Scan(&ID, &Firstname, &Lastname, &Email, &Hash, &HasInfo); err != nil {
    return nil, err
  }

  id := int(ID)

  user := &UserModel{
    ID:        id,
    Firstname: Firstname.String,
    Lastname:  Lastname.String,
    Email:     Email.String,
    Hash:      Hash.String,
    HasInfo:   HasInfo.String,
  }

  return user, nil
}

// ------------- HELPER FUNCTIONS END ----------------------------

// --------------- DB FUNCTIONS BEGIN -------------------------------------

// ----------------- DB FUNCTIONS END ------------------------------
