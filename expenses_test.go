package banktorrent

import (
  "testing"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)


func Test_GetExpenseById(t *testing.T) {
  db, err := sql.Open("sqlite3", TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()
  
  expense, err := GetExpenseById(db, 1)
  test_error_helper(t, err)
  if expense.Amount != 2000 {
    t.Error("Fetched expense has the wrong amount: ", expense)
  }
  if expense.Date != "2013-07-01" {
    t.Error("Fetched expense has wrong date: ", expense)
  }
  if expense.LabelId != 1 {
    t.Error("Fetched expense has the wrong label association: ", expense)
  }
  if expense.Comment != "grocery test expense" {
    t.Error("Fetched expense has the wrong comment")
  }
}

func Test_GetExpensesByLabel(t *testing.T) {
  db, err := sql.Open("sqlite3", TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()
  
  t.Error("Unimplemented test")
}

func Test_CreateExpense(t *testing.T) {
  db, err := sql.Open("sqlite3", TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()

  t.Error("Unimplemented test")
}

func Test_AddSimpleExpense(t *testing.T) {
  db, err := sql.Open("sqlite3", TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()

  expense_amount := 100
  user1, err := GetUserByName(db, "User One")
  test_error_helper(t, err)
  user2, err := GetUserByName(db, "User Two")
  test_error_helper(t, err)
  label, err := GetLabelByName(db, "miscellaneous")
  test_error_helper(t, err)
  expense, err := CreateExpense(db, expense_amount, label, "test miscellaneous simple expense", "2013-08-01")

  user1_old_balance := user1.Balance
  user2_old_balance := user2.Balance

  transaction, err := AddSimpleExpense(db, user1, user2, expense)
  test_error_helper(t, err)
  
  if transaction.LenderId != user1.Id {
    t.Error("Transaction has the incorrect lender. Should be: ", user1.Id, ", but got: ", transaction.LenderId)
  }
  if transaction.DebtorId != user2.Id {
    t.Error("Transaction has the incorrect debtor. Should be: ", user2.Id, ", but got: ", transaction.DebtorId)
  }
  if transaction.Date != expense.Date {
    t.Error("Transaction has the incorrect date. Should be: ", expense.Date, ", but got: ", transaction.Date)
  } 
  if transaction.Amount != expense.Amount {
    t.Error("Transaction has the incorrect amount. Should be: ", expense.Amount, ", but got: ", transaction.Amount)
  }
  if transaction.ExpenseId != expense.Id {
    t.Error("Transaction has the wrong expense association. Should be: ", expense.Id, ", but got: ", transaction.ExpenseId)
  }

  if user1.Balance != user1_old_balance + expense_amount {
    t.Error("Lender balance not updated correctly. Should be: ", (user1_old_balance + expense_amount), " but got: ", user1.Balance, "instead.")
  }
  if user2.Balance != user2_old_balance - expense_amount {
    t.Error("Debtor balance not updated correctly. Should be: ", (user2_old_balance - expense_amount), " but got: ", user2.Balance, "instead.")
  }

  t.Error("Need to remove the expense at the end of the test")
}

func Test_AddSplitExpense(t *testing.T) {
  db, err := sql.Open("sqlite3", TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()

  t.Error("Unimplemented test")
}

func Test_ExpenseTransactions(t *testing.T) {
  db, err := sql.Open("sqlite3", TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()

  t.Error("Unimplemented test")
}

