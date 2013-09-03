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
  
  label, err := GetLabelByName(db, "groceries")
  test_error_helper(t, err)
  expenses, err := GetExpensesByLabel(db, label)
  test_error_helper(t, err)
  if len(expenses) != 1 {
    t.Error("Expected 1 expense, but found ", len(expenses), " for the label: ", label.Name)
  }
  if expenses[0].Amount != 2000 {
    t.Error("Expense amount incorrect: ", expenses[0])
  } 
  if expenses[0].Id != 1 {
    t.Error("Expense id incorrect: ", expenses[0])
  }
  if expenses[0].Date != "2013-07-01" {
    t.Error("Expense date incorrect: ", expenses[0])
  }
  if expenses[0].Comment != "grocery test expense" {
    t.Error("Expense comment incorrect: ", expenses[0])
  }
}

func Test_ExpenseTransactions(t *testing.T) {
  db, err := sql.Open("sqlite3", TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()

  expense, err := GetExpenseById(db, 1)
  test_error_helper(t, err)
  transactions, err := expense.Transactions(db)
  test_error_helper(t, err)
  if len(transactions) != 1 {
    t.Error("Expected 1 expense transaction, but got ", len(transactions))
  }
  if transactions[0].Id != 1 {
    t.Error("Expense transaction has wrong id: ", transactions[0])
  }
  if transactions[0].Amount != 1000 {
    t.Error("Expense transaction has wrong amount: ", transactions[0])
  }
}

func Test_CreateDeleteExpense(t *testing.T) {
  db, err := sql.Open("sqlite3", TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()

  label, err := GetLabelByName(db, "groceries")
  test_error_helper(t, err)
  expense, err := CreateExpense(db, 100, label, "test dollar expense", "2013-08-04")
  test_error_helper(t, err)
  t.Log("Created expense with id: ", expense.Id)
  err = expense.Remove(db)
  test_error_helper(t, err)
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
  t.Log("Created expense with id: ", expense.Id)

  user1_old_balance := user1.Balance
  user2_old_balance := user2.Balance

  // Add expense
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

  // Remove expense
  test_error_helper(t, expense.Remove(db))
  test_error_helper(t, user1.Reload(db))
  test_error_helper(t, user2.Reload(db))
  if user1.Balance != user1_old_balance {
    t.Error("Lender balance incorrect after simple expense removal. Expected: ", user1_old_balance, ", but got: ", user1.Balance)
  }
  if user2.Balance != user2_old_balance {
    t.Error("Debtor balance incorrect after simple expense removal. Expected: ", user2_old_balance, ", but got: ", user2.Balance)
  }
}

func Test_AddSplitExpense(t *testing.T) {
  db, err := sql.Open("sqlite3", TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()
  
  lender, err := GetUserById(db, 1)
  test_error_helper(t, err)
  lender_starting_balance := lender.Balance
  debtor1, err := GetUserById(db, 2)
  test_error_helper(t, err)
  debtor1_starting_balance := debtor1.Balance
  debtor2, err := GetUserById(db, 3)
  test_error_helper(t, err)
  debtor2_starting_balance := debtor2.Balance
  expense, err := GetExpenseById(3)
  test_error_helper(t, err)
  expense_amount := expense.Amount

  // Add expense
  transactions, err := AddSplitExpense(db, lender, []*User{debtor1, debtor2}, expense)
  test_error_helper(t, err)
  if len(transactions) != 2 {
    t.Error("Wrong number of transactions present. Expected: 2, but got: ", len(transactions))
  }
  if transactions[0].Amount != transactions[1].Amount {
    t.Error("Transaction amounts not equal: ", transactions[0].Amount, " and ", transactions[1].Amount)
  }
  if (transactions[0].Amount + transactions[1].Amount) != expense.Amount {
    t.Error("Transaction amounts do not add up to expense amount.")
  }
  if lender.Balance != (lender_starting_balance + (2.0/3) * expense_amount) {
    // 2/3 of the expense amount is the amount owed to the lender since he is also paying for the expense
    t.Error("Expense amount not split evenly between participants (lender)")
  }
  if debtor1.Balance != (debtor1_starting_balance + (1.0/3) * expense_amount) {
    t.Error("Expense amount not split evenly between participants (debtor1)")
  }
  
  // test user balances here

  // Remove expense
  // test user balances reverted


  t.Error("Unimplemented test")
}



