package banktorrent

import (
  "testing"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)


func TestGetTransactionById(t *testing.T) {
  db, err := sql.Open("sqlite3", TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()
  
  transaction, err := GetTransactionById(db, 21)
  test_error_helper(t, err)
  if transaction.Amount != 81920 {
    t.Error("Fetched transaction has the wrong amount: ", transaction)
  }
  if transaction.Date != "2013-04-01" {
    t.Error("Fetched transaction has wrong date: ", transaction)
  }
}

func TestGetTransactionsByLabel(t *testing.T) {
  db, err := sql.Open("sqlite3", TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()

  transactions, err := GetTransactionsByLabel(db, 1)
  test_error_helper(t, err)
  if len(transactions) < 19 {
    t.Error("Didn't fetch all of the groceries transactions: ", transactions)
  }
  if len(transactions) > 0 && (transactions[0]).LabelId != 1 {
    t.Error("Transactions from the wrong label were found: ", transactions)
  }
}