package banktorrent

import (
  "github.com/russross/meddler"
)


type User struct {
  Id      int       `meddler:"id,pk"`
  Name    string    `meddler:"name"`
  Email   string    `meddler:"email"`
  Balance int       `meddler:"balance"`
}


func GetUserById(db meddler.DB, id int) (*User, error) {
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

//func AddUser(db meddler.DB, )
