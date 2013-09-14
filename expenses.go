package banktorrent

import (
  "github.com/russross/meddler"
  "errors"
)


type MoneyAmount struct {
  Dollars int
  Cents int
}

type Expense struct {
  Id      int     `meddler:"id,pk"`
  Amount  int     `meddler:"amount"`
  LabelId int     `meddler:"label_id"`
  Comment string  `meddler:"comment"`
  Date    string  `meddler:"date"`
}


func GetExpenseById(db meddler.DB, id int) (*Expense, error) {
  expense := new(Expense)
  err := meddler.Load(db, "expenses", expense, id)
  return expense, err
}

func GetExpensesByLabel(db meddler.DB, label *Label) ([]*Expense, error) {
  var expenses []*Expense
  err := meddler.QueryAll(db, &expenses, "select * from expenses where label_id = ?", label.Id)
  return expenses, err
}

// Creates the expense in the db and gives it an ID
func CreateExpense(db meddler.DB, amount int, label *Label, comment string, date string) (*Expense, error) {
  expense := new(Expense)
  expense.Amount = amount
  expense.LabelId = label.Id
  expense.Comment = comment
  expense.Date = date

  err := meddler.Insert(db, "expenses", expense)
  return expense, err
}

// Adds a single transaction. The entire expense amount is credited to the lender and charged to the debtor
func (e *Expense) AddSimple(db meddler.DB, lender *User, debtor *User) (*Transaction, error) {
  return addTransaction(db, lender, debtor, e.Amount, e)
}

// Adds 1 transaction per debtor. The amount is split evenly between everyone (penny errors can occur).
// This type of transaction is useful even with just 2 people (a debtor and a lender) because it behaves
// differently than a SimpleExpense. Namely, it splits the expense.Amount evenly between everyone, including
// the lender, allowing for "shared expenses" to be added easily
func (e *Expense) AddSplit(db meddler.DB, lender *User, debtors []*User) ([]*Transaction, error) {
  if len(debtors) < 1 {
    return nil, errors.New("Debtors list must have at least 1 element")
  }
  split_amount := int(float32(e.Amount) / float32(len(debtors) + 1)) //# of debtors + 1 lender
  transactions := make([]*Transaction, len(debtors))
  for index, debtor := range(debtors) {
    transaction, err := addTransaction(db, lender, debtor, split_amount, e)
    if err != nil {
      return transactions, err
    }
    transactions[index] = transaction
  }
  return transactions, nil
}

func (e *Expense) Transactions(db meddler.DB) ([]*Transaction, error) {
  var transactions []*Transaction
  err := meddler.QueryAll(db, &transactions, "select * from transactions where expense_id = ?", e.Id)
  return transactions, err
}

// Removes the expense by reversing all transactions associated with it,
// and deleting the expense from the db
func (e *Expense) Remove(db meddler.DB) (error) {
  transactions, err := e.Transactions(db)
  if err != nil {
    return err
  }
  for _, transaction := range(transactions) {
    err := transaction.remove(db)
    if err != nil {
      return err
    }
  }
  // remove the expense from the db
  _, err = db.Exec("DELETE FROM expenses WHERE id = ?", e.Id)
  if err != nil {
    return err
  }
  e = nil
  return nil
}
