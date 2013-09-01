package banktorrent

import (
  "testing"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)


func TestGetUserByName(t *testing.T) {
  db, err := sql.Open("sqlite3", TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()
  
  user_jason, err := GetUserByName(db, "Jason Dsouza")
  test_error_helper(t, err)
  if user_jason.Id != 1 {
    t.Error("Fetched user has wrong ID", user_jason)
  }
  user_rachel, err := GetUserByName(db, "Rachel Repard")
  test_error_helper(t, err)
  if user_rachel.Id != 2 {
    t.Error("Fetched user has wrong ID: ", user_rachel)
  }
}

func TestSample(t *testing.T) {
  t.Log("Hello")
  //t.Error("testing failures")
}
