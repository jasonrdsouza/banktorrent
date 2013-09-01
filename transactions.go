package banktorrent

import (
  "time"
  "github.com/russross/meddler"
)



type MoneyAmount struct {
  Dollars int
  Cents int
}

type Transaction struct {
  Id        int     `meddler:"id,pk"`
  LenderId  int     `meddler:"lender_id"`
  DebtorId  int     `meddler:"debtor_id"`
  Amount    int     `meddler:"amount"`
  LabelId   int     `meddler:"label_id"`
  Comment   string  `meddler:"comment"`
  Date      string  `meddler:"date"`
}


func GetTransactionById(db meddler.DB, id int) (*Transaction, error) {
  trans := new(Transaction)
  err := meddler.QueryRow(db, trans, "select * from transactions where id = ?", id)
  if err != nil {
    return nil, err
  }
  return trans, nil
}

func GetTransactionsByLabel(db meddler.DB, label_id int) ([]*Transaction, error) {
  var transactions []*Transaction
  err := meddler.QueryAll(db, &transactions, "select * from transactions where label_id = ?", label_id)
  if err != nil {
    return nil, err
  }
  return transactions, nil
}


