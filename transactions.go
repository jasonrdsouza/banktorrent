package banktorrent

import (
  "github.com/russross/meddler"
)


type Transaction struct {
  Id        int     `meddler:"id,pk"`
  LenderId  int     `meddler:"lender_id"`
  DebtorId  int     `meddler:"debtor_id"`
  Amount    int     `meddler:"amount"`
  Date      string  `meddler:"date"`
  ExpenseId int     `meddler:"expense_id"`
}


func GetTransactionById(db meddler.DB, id int) (*Transaction, error) {
  trans := new(Transaction)
  err := meddler.QueryRow(db, trans, "select * from transactions where id = ?", id)
  if err != nil {
    return nil, err
  }
  return trans, nil
}

// Adds a transaction to the db, and updates the user's balances accordingly
func addTransaction(db meddler.DB, lender *User, debtor *User, amount int, expense *Expense) (*Transaction, error) {
  trans := new(Transaction)
  trans.LenderId = lender.Id
  trans.DebtorId = debtor.Id
  trans.Amount = amount
  trans.Date = expense.Date
  trans.ExpenseId = expense.Id

  err := meddler.Insert(db, "transactions", trans)
  if err != nil {
    return nil, err
  }

  lender.UpdateBalance(amount)
  debtor.UpdateBalance(-amount)

  return trans
}
