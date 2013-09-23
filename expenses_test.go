package banktorrent

import (
  "testing"
  _ "github.com/mattn/go-sqlite3"
)


func Test_GetExpenseById(t *testing.T) {
  db, err := Connect(TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()

  expense, err := GetExpenseById(db, 1)
  handle_error(t, err)
  if expense.Amount != 2500 {
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
  db, err := Connect(TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()

  label, err := GetLabelByName(db, "groceries")
  handle_error(t, err)
  expenses, err := GetExpensesByLabel(db, label)
  handle_error(t, err)
  if len(expenses) != 1 {
    t.Error("Expected 1 expense, but found ", len(expenses), " for the label: ", label.Name)
  }
  if expenses[0].Amount != 2500 {
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
  db, err := Connect(TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()

  expense, err := GetExpenseById(db, 1)
  handle_error(t, err)
  transactions, err := expense.Transactions(db)
  handle_error(t, err)
  if len(transactions) != 1 {
    t.Error("Expected 1 expense transaction, but got ", len(transactions))
  }
  if transactions[0].Id != 1 {
    t.Error("Expense transaction has wrong id: ", transactions[0])
  }
  if transactions[0].Amount != 2500 {
    t.Error("Expense transaction has wrong amount: ", transactions[0])
  }
}

func Test_CreateDeleteExpense(t *testing.T) {
  db, err := Connect(TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()

  label, err := GetLabelByName(db, "groceries")
  handle_error(t, err)
  expense, err := CreateExpense(db, 100, label, "test dollar expense", "2013-08-04")
  handle_error(t, err)
  t.Log("Created expense with id: ", expense.Id)
  err = expense.Remove(db)
  handle_error(t, err)
}

func Test_AddSimpleExpense(t *testing.T) {
  db, err := Connect(TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()

  expense_amount := 100
  user1, err := GetUserByName(db, "User One")
  handle_error(t, err)
  user2, err := GetUserByName(db, "User Two")
  handle_error(t, err)
  label, err := GetLabelByName(db, "miscellaneous")
  handle_error(t, err)
  expense, err := CreateExpense(db, expense_amount, label, "test miscellaneous simple expense", "2013-08-01")
  t.Log("Created expense with id: ", expense.Id)

  user1_old_balance := user1.Balance
  user2_old_balance := user2.Balance

  // Add expense
  transaction, err := expense.AddSimple(db, user1, user2)
  handle_error(t, err)
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
  handle_error(t, expense.Remove(db))
  handle_error(t, user1.Reload(db))
  handle_error(t, user2.Reload(db))
  if user1.Balance != user1_old_balance {
    t.Error("Lender balance incorrect after simple expense removal. Expected: ", user1_old_balance, ", but got: ", user1.Balance)
  }
  if user2.Balance != user2_old_balance {
    t.Error("Debtor balance incorrect after simple expense removal. Expected: ", user2_old_balance, ", but got: ", user2.Balance)
  }
}

func Test_AddSplitExpense(t *testing.T) {
  db, err := Connect(TEST_DB)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()

  lender, err := GetUserById(db, 1)
  handle_error(t, err)
  lender_starting_balance := lender.Balance
  debtor1, err := GetUserById(db, 2)
  handle_error(t, err)
  debtor1_starting_balance := debtor1.Balance
  debtor2, err := GetUserById(db, 3)
  handle_error(t, err)
  debtor2_starting_balance := debtor2.Balance
  expense, err := GetExpenseById(db, 3)
  handle_error(t, err)
  expense_amount := expense.Amount

  // Add expense
  transactions, err := expense.AddSplit(db, lender, []*User{debtor1, debtor2})
  handle_error(t, err)

  if len(transactions) != 2 {
    t.Error("Wrong number of transactions present. Expected: 2, but got: ", len(transactions))
  }
  if transactions[0].Amount != transactions[1].Amount {
    t.Error("Transaction amounts not equal: ", transactions[0].Amount, " and ", transactions[1].Amount)
  }
  if int(float32(transactions[0].Amount + transactions[1].Amount) * (3.0/2)) != expense.Amount {
    t.Error("Transaction amounts do not add up to expense amount.")
  }
  if lender.Balance != (lender_starting_balance + int((2.0/3) * float32(expense_amount))) {
    // 2/3 of the expense amount is the amount owed to the lender since he is also paying for the expense
    t.Error("Expense amount not split evenly between participants (lender)")
  }
  if debtor1.Balance != (debtor1_starting_balance - int((1.0/3) * float32(expense_amount))) {
    t.Error("Expense amount not split evenly between participants (debtor1)")
  }
  if debtor2.Balance != (debtor2_starting_balance - int((1.0/3) * float32(expense_amount))) {
    t.Error("Expense amount not split evenly between participants (debtor2)")
  }

  // Remove expense
  handle_error(t, expense.Remove(db))
  handle_error(t, lender.Reload(db))
  handle_error(t, debtor1.Reload(db))
  handle_error(t, debtor2.Reload(db))

  if lender.Balance != lender_starting_balance {
    t.Error("Lender balance incorrect after split expense removal. Expected: ", lender_starting_balance, ", but got: ", lender.Balance)
  }
  if debtor1.Balance != debtor1_starting_balance {
    t.Error("Debtor1 balance incorrect after split expense removal. Expected: ", debtor1_starting_balance, ", but got: ", debtor1.Balance)
  }
  if debtor2.Balance != debtor2_starting_balance {
    t.Error("Debtor2 balance incorrect after split expense removal. Expected: ", debtor2_starting_balance, ", but got: ", debtor2.Balance)
  }
}
