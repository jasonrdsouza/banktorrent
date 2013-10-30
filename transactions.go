package banktorrent

import (
  "github.com/russross/meddler"
  "fmt"
)


type Transaction struct {
  Id        int64     `meddler:"id,pk"`
  LenderId  int64     `meddler:"lender_id"`
  DebtorId  int64     `meddler:"debtor_id"`
  Amount    int       `meddler:"amount"`
  Date      string    `meddler:"date"`
  ExpenseId int64     `meddler:"expense_id"`
}


func GetTransactionById(db meddler.DB, id int64) (*Transaction, error) {
  trans := new(Transaction)
  err := meddler.Load(db, "transactions", trans, id)
  return trans, err
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

  lender.UpdateBalance(db, amount)
  debtor.UpdateBalance(db, -amount)

  return trans, nil
}

func (t *Transaction) String() (string) {
  return fmt.Sprintf("Transaction #%v [LenderID: %v, DebtorID: %v, Amount: %v, Date: %v, ExpenseID: %v]", t.Id, t.LenderId, t.DebtorId, t.Amount, t.Date, t.ExpenseId)
}

// Reverses a transaction, removes it from the DB, and frees the struct
func (t *Transaction) remove(db meddler.DB) error {
  lender, err := GetUserById(db, t.LenderId)
  if err != nil {
    return err
  }
  debtor, err := GetUserById(db, t.DebtorId)
  if err != nil {
    return err
  }

  // reverse the balance updates due to this transaction
  lender.UpdateBalance(db, -(t.Amount))
  debtor.UpdateBalance(db, t.Amount)

  // remove the transaction from the db
  _, err = db.Exec("DELETE FROM transactions WHERE id = ?", t.Id)
  if err != nil {
    return err
  }
  t = nil

  return nil
}
