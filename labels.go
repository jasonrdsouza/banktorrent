package banktorrent

import (
  "fmt"
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

func CreateLabel(db meddler.DB, name string) (*Label, error) {
  label := new(Label)
  label.Name = name

  err := meddler.Insert(db, "labels", label)
  return label, err
}

func (l *Label) String() (string) {
  return fmt.Sprintf("Label #%v [%v]", l.Id, l.Name)
}
