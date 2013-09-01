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


func GetUserByName(db meddler.DB, name string) (*User, error) {
  user := new(User)
  err := meddler.QueryRow(db, user, "select * from users where name = ?", name)
  if err != nil {
    return nil, err
  }
  return user, nil
}

func (u *User) UpdateBalance(delta int) (error) {
  u.balance += delta
  err := meddler.Update(db, "users", u)
  if err != nil {
    return err
  }
  return nil
}

//func AddUser(db meddler.DB, )
