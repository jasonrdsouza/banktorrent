package banktorrent

import (
  "github.com/russross/meddler"
)


type Label struct {
  Id    int     `meddler:"id,pk"`
  Name  string  `meddler:"name"`
}


func GetLabelById(db meddler.DB, id int) (*Label, error) {
  label := new(Label)
  err := meddler.Load(db, "labels", label, id)
  return label, err
}

func GetLabelByName(db meddler.DB, name string) (*Label, error) {
  label := new(Label)
  err := meddler.QueryRow(db, label, "select * from labels where name = ?", name)
  return label, err
}