package banktorrent

import (
  "testing"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)


func TestGetLabelByName(t *testing.T) {
  db, err := sql.Open("sqlite3", TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()
  
  label, err := GetLabelByName(db, "groceries")
  test_error_helper(t, err)
  if label.Id != 1 {
    t.Error("Fetched label has the wrong id: ", label)
  }
}