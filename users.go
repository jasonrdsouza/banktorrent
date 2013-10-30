package banktorrent

import (
  "fmt"
  "github.com/russross/meddler"
)


type User struct {
  Id      int64     `meddler:"id,pk"`
  Name    string    `meddler:"name"`
  Email   string    `meddler:"email"`
  Balance int       `meddler:"balance"`
}


func GetUserById(db meddler.DB, id int64) (*User, error) {
  user := new(User)
  err := meddler.Load(db, "users", user, id)
  return user, err
}

func GetUserByName(db meddler.DB, name string) (*User, error) {
  user := new(User)
  err := meddler.QueryRow(db, user, "select * from users where name = ?", name)
  return user, err
}

func (u *User) UpdateBalance(db meddler.DB, delta int) (error) {
  u.Balance += delta
  return meddler.Update(db, "users", u)
}

func (u *User) Reload(db meddler.DB) (error) {
  return meddler.Load(db, "users", u, u.Id)
}

func (u *User) String() (string) {
  return fmt.Sprintf("User #%v [Name: %v, Email: %v, Balance: %v]", u.Id, u.Name, u.Email, u.Balance)
}

func CreateUser(db meddler.DB, name string, email string, balance int) (*User, error) {
  user := new(User)
  user.Name = name
  user.Email = email
  user.Balance = balance

  err := meddler.Insert(db, "users", user)
  return user, err
}
