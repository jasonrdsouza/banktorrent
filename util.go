package banktorrent

import (
  "testing"
  "database/sql"
)


const (
  TEST_DB = "/tmp/banktorrent.test.db"
  PROD_DB = ""
)


func handle_error(t *testing.T, err error) {
  if err != nil {
    t.Error(err)
  }
}

func Connect(database string) (*sql.DB, error) {
  return sql.Open("sqlite3", database)
}
