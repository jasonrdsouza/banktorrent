package banktorrent

import (
  "testing"
)


func Test_GetTransactionById(t *testing.T) {
  db, err := Connect(TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()

  transaction, err := GetTransactionById(db, 1)
  handle_error(t, err)
  if transaction.Amount != 2500 {
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

func Test_AddRemoveTransaction(t *testing.T) {
  db, err := Connect(TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()

  // Add and check transaction
  amount := 200 // $2
  lender, err := GetUserById(db, 1)
  handle_error(t, err)
  old_lender_balance := lender.Balance
  debtor, err := GetUserById(db, 2)
  handle_error(t, err)
  old_debtor_balance := debtor.Balance
  expense, err := GetExpenseById(db, 2)
  handle_error(t, err)
  transaction, err := addTransaction(db, lender, debtor, amount, expense)
  handle_error(t, err)
  t.Log("Created transaction with id: ", transaction.Id)
  if lender.Balance != (old_lender_balance + amount) {
    t.Error("Lender balance not updated correctly on transaction creation: ", lender)
  }
  if debtor.Balance != (old_debtor_balance - amount) {
    t.Error("Debtor balance not updated correctly on transaction creation: ", debtor)
  }

  // Remove transaction
  err = transaction.remove(db)
  handle_error(t, err)
  err = lender.Reload(db)
  handle_error(t, err)
  err = debtor.Reload(db)
  handle_error(t, err)

  if lender.Balance != old_lender_balance {
    t.Error("Lender balance not updated correctly on transaction removal: ", lender)
  }
  if debtor.Balance != old_debtor_balance {
    t.Error("Debtor balance not updated correctly on transaction removal: ", debtor)
  }
}
