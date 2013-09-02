package banktorrent

import (
  "testing"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)


func Test_GetTransactionById(t *testing.T) {
  db, err := sql.Open("sqlite3", TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()
  
  transaction, err := GetTransactionById(db, 1)
  test_error_helper(t, err)
  if transaction.Amount != 1000 {
    t.Error("Fetched transaction has the wrong amount: ", transaction)
  }
  if transaction.Date != "2013-07-01" {
    t.Error("Fetched transaction has wrong date: ", transaction)
  }
  if transaction.LenderId != 1 || transaction.DebtorId != 2 {
    t.Error("Fetched transaction has incorrect lender/ debtor: ", transaction)
  }
  if transaction.ExpenseId != 1 {
    t.Error("Fetched transaction has incorrect expense association: ", transaction)
  } 
}

func Test_AddTransaction(t *testing.T) {
  db, err := sql.Open("sqlite3", TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()

  t.Error("Unimplemented test")
}

func Test_RemoveTransaction(t *testing.T) {
  db, err := sql.Open("sqlite3", TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()

  t.Error("Unimplemented test")
}