package banktorrent

import (
  "testing"
  _ "github.com/mattn/go-sqlite3"
)


func Test_GetUserById(t *testing.T) {
  db, err := Connect(TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()

  user, err := GetUserById(db, 1)
  handle_error(t, err)
  if user.Name != "User One" {
    t.Error("Fetched user has wrong name: ", user)
  }
  if user.Email != "user-one@test.com" {
    t.Error("Fetched user has wrong email address: ", user)
  }
}

func Test_GetUserByName(t *testing.T) {
  db, err := Connect(TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()

  user1, err := GetUserByName(db, "User One")
  handle_error(t, err)
  if user1.Id != 1 {
    t.Error("Fetched user has wrong ID: ", user1)
  }
  user2, err := GetUserByName(db, "User Two")
  handle_error(t, err)
  if user2.Email != "user-two@test.com" {
    t.Error("Fetched user has wrong email address: ", user2)
  }
}

func Test_UserUpdateBalance(t *testing.T) {
  db, err := Connect(TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()

  amount := 100 // $1
  user, err := GetUserById(db, 1)
  handle_error(t, err)
  old_balance := user.Balance
  handle_error(t, user.UpdateBalance(db, amount))
  if user.Balance != (old_balance + amount) {
    t.Error("User balance update failed. Expected: ", (old_balance + amount), ", but got: ", user.Balance)
  }
  handle_error(t, user.UpdateBalance(db, -amount))
  if user.Balance != old_balance {
    t.Error("User balance negative update failed. Expected: ", old_balance, ", but got: ", user.Balance)
  }
}

func Test_Sample(t *testing.T) {
  t.Log("Hello")
  //t.Error("testing failures")
}
