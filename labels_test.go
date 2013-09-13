package banktorrent

import (
  "testing"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)


func Test_GetLabelById(t *testing.T) {
  db, err := sql.Open("sqlite3", TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()

  label, err := GetLabelById(db, 1)
  handle_error(t, err)
  if label.Name != "groceries" {
    t.Error("Fetched label has wrong name: ", label)
  }
}

func Test_GetLabelByName(t *testing.T) {
  db, err := sql.Open("sqlite3", TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()
  
  label, err := GetLabelByName(db, "groceries")
  handle_error(t, err)
  if label.Id != 1 {
    t.Error("Fetched label has the wrong id: ", label)
  }
}