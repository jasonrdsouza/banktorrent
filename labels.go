package banktorrent

import (
  "github.com/russross/meddler"
)


type Label struct {
  Id    int     `meddler:"id,pk"`
  Name  string  `meddler:"name"`
}


func GetLabelByName(db meddler.DB, name string) (*Label, error) {
  label := new(Label)
  err := meddler.QueryRow(db, label, "select * from labels where name = ?", name)
  if err != nil {
    return nil, err
  }
  return label, nil
}