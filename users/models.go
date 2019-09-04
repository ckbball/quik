package users

import (
  "errors"
  "github.com/jinzhu/gorm"
  "golang.org/x/crypto/bcrypt"
  "fmt"
  "github.com/ckbball/quik/common"
)

type UserModel struct {
  ID        int    `json:"id" gorm:"primary_key"`
  Firstname string `json:"first" gorm:"column:firstname"`
  Lastname  string `json:"last" gorm:"column:lastname"`
  Email     string `json:"email" gorm:"column:email;unique_index"`
  Hash      string `json:"pass" gorm:"column:password;not null"`
  HasInfo   bool   `json:"info" gorm:"column:hasinfo"`
  Status     string   `gorm:"column:status"`// this is going to be searching, perusing, locked
  Level      string   `gorm:"column:level"`// this is going to be entry, mid, senior
  // Info      UserInfo `json:"info"`  should this be in UserModel or should it just be a separate table that I also grab
}

type UserInfo struct {
  Roles      []string // this is going to be backend, frontend, full stack, mobile,
  Frameworks []string // this is going to be all frameworks, front and back, that user knows meaning they built a project with it
  DB         []string // this is all dbs that a user has built a project with
  Front      []string // languages and methods for front end user has built a project with
  Back       []string // languages and methods for backend user has built a project with
  Extra      []string // dont know what should be here
  DevOps     []string // CI/CD tools, other things idk
  Cloud      []string // which cloud tools and platforms user has made a project with
  ID     int      
}

func AutoMigrate() {
  db := common.GetDB()

  db.AutoMigrate(&UserModel{})
  db.AutoMigrate(&UserInfo{})
}

// There will be multiple relations table with id, user_id, <a field of userInfo>_id, 

// -------- HELPER FUNCTIONS BEGIN --------------------------------------

func UserValidatorToModel(newUser UserModelValidator) (UserModel, error) {
  u := UserModel{}

  u.ID = newUser.User.ID
  u.Firstname = newUser.User.Firstname
  u.Lastname = newUser.User.Lastname
  u.Email = newUser.User.Email
  u.Hash = newUser.User.Hash
  u.HasInfo = newUser.User.HasInfo
  u.Level = newUser.User.Level
  u.Status = newUser.User.Status

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
    Status    sql.NullString
    Level     sql.NullString
  )

  if err := s.Scan(&ID, &Firstname, &Lastname, &Email, &Hash, &HasInfo, &Status, &Level); err != nil {
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
    Level:     Level.String,
    Status:    Status.String
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
    Status    sql.NullString
    Level     sql.NullString
  )

  if err := s.Scan(&ID, &Firstname, &Lastname, &Email, &Hash, &HasInfo, &Status, &Level); err != nil {
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
    Level:      Level.String,
    Status:     Status.String
  }

  return user, nil
}

func scanUserInfos(s *sql.Rows) (*UserModel, error) {
  var (
    ID     int
    Roles      sql.JSON
    Frameworks sql.JSON
    DB         sql.JSON
    Front      sql.JSON
    Back       sql.JSON
    Extra      sql.JSON
    DevOps     sql.JSON
    Cloud      sql.JSON
  )

  if err := s.Scan(&ID, &Roles, &Frameworks, &DB, &Front, &Back, &Extra, &DevOps, &Cloud); err != nil {
    return nil, err
  }

  id := int(ID)

  user := &UserModel{
    ID:        id,
    Status: Status.String,
    Level:  Level.String,
    Roles:     Roles.Slice,
    Frameworks:      Frameworks.Slice,
    DB:   DB.String,
    Front: Front.

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

func GetAllUsers() ([]*UserModel, error) {
  rows, err := common.GetRows("users")
  if err != nil {
    return nil, err
  }

  defer rows.Close()

  var users []*UserModel
  for rows.Next() {
    user, err := scanUsers(rows)
    if err != nil {
      return nil, fmt.Errorf("ERROR ---> mysql: could not read row: %v", err)
    }

    users = append(users, user)
  }

  return users, nil
}

func SaveUser(u UserModel) error {

  r, err := common.ExecAffectingOneRow("insert", "user", u.ID, u.Firstname, u.Lastname, u.Email, u.Hash, u.HasInfo)
  if err != nil {
    return err
  }
  fmt.Println("Insert user operation results: %v", r)

  return nil
}

func GetUser(id int) (*UserModel, error) {
  row, err := common.GetRow("users", id)
  if err != nil {
    return nil, err
  }

  user, err := scanUser(row)
  if err != nil {
    return nil, fmt.Errorf("ERROR --->> mysql: could not read row: %v", err)
  }

  return flight, nil
}

// ----------------- DB FUNCTIONS END ------------------------------
