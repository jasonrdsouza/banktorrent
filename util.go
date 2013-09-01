package banktorrent

import (
  "testing"
)


const (
  TEST_DB = "/tmp/banktorrent.test.db"
  PROD_DB = ""
)


func test_error_helper(t *testing.T, err error) {
  if err != nil {
    t.Error(err)
  }
}

