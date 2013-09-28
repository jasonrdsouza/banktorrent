package main

import (
  "flag"
  "fmt"
  "log"
  "time"
  "bytes"
  "strings"
  "text/template"
  "github.com/jasonrdsouza/banktorrent"
  "database/sql"
)


const (
  DATE_FMT = "2006-01-02"
  tmpl = `
  Parameters
  ----------
  - Date: {{.Date}}
  - Label:  {{.Label}}
  - Lender: {{.Lender}}
  - Debtor: {{.Debtor}}
  - Amount: {{.Amount}}
  - Kind: {{.Kind}}
  - Comment:  {{.Comment}}
  `
)

type RawParams struct {
  Label string
  Lender string
  Debtor string
  Date string
  Amount string
  Comment string
  Kind string
}

type ValidParams struct {
  Label *banktorrent.Label
  Lender *banktorrent.User
  Debtor *banktorrent.User
  Date time.Time
  Amount *banktorrent.MoneyAmount
  Comment string
  Kind banktorrent.ExpenseType
}

var raw = new(RawParams)
func init() {
  flag.StringVar(&raw.Label, "label", "misc", "The label for this expense")
  flag.StringVar(&raw.Lender, "lender", "bob", "The lender for this expense")
  flag.StringVar(&raw.Debtor, "debtor", "alice", "The debtor for this expense")
  flag.StringVar(&raw.Date, "date", "2013-09-31", "The date that the expense took place")
  flag.StringVar(&raw.Amount, "amount", "25.00", "The amount of this expense")
  flag.StringVar(&raw.Comment, "comment", "sample cli expense", "Comments associated with this expense")
  flag.StringVar(&raw.Kind, "kind", "simple", "What kind of expense this is")
}

func main() {
  flag.Parse()
  fmt.Println("Input params: ", raw)

  db, err := banktorrent.Connect(banktorrent.TEST_DB)
  if err != nil {
    log.Fatalln(err)
  }
  defer db.Close()

  valid, err := validateParams(db, raw)
  if err != nil {
    log.Fatalln("Failed to validate params: ", err)
  }
  fmt.Println("Validated params: ", valid)
  fmt.Println("Adding Expense")
}

func (r *RawParams) String() (string) {
  t := template.Must(template.New("raw_tmpl").Parse(tmpl))
  var buf = new(bytes.Buffer)
  err := t.Execute(buf, r)
  if err != nil {
    return ""
  }
  return buf.String()
}

func (v *ValidParams) String() (string) {
  t := template.Must(template.New("valid_tmpl").Parse(tmpl))
  var buf = new(bytes.Buffer)
  err := t.Execute(buf, v)
  if err != nil {
    return ""
  }
  return buf.String()
}

func validateParams(db *sql.DB, raw *RawParams) (*ValidParams, error) {
  label, err := banktorrent.GetLabelByName(db, strings.TrimSpace(raw.Label))
  if err != nil {
    log.Println("Error validating label", raw.Label)
    return nil, err
  }
  lender, err := banktorrent.GetUserByName(db, strings.TrimSpace(raw.Lender))
  if err != nil {
    log.Println("Error validating lender ", raw.Lender)
    return nil, err
  }
  debtor, err := banktorrent.GetUserByName(db, strings.TrimSpace(raw.Debtor))
  if err != nil {
    log.Println("Error validating debtor ", raw.Debtor)
    return nil, err
  }
  date, err := time.Parse(DATE_FMT, strings.TrimSpace(raw.Date))
  if err != nil {
    log.Println("Error validating date ", raw.Date)
    return nil, err
  }
  amount, err := banktorrent.StringToMoney(raw.Amount)
  if err != nil {
    log.Println("Error validating amount ", raw.Amount)
    return nil, err
  }
  comment := strings.TrimSpace(raw.Comment)
  // fix the kind
  kind := banktorrent.SIMPLE_EXPENSE

  return &ValidParams{Label: label, Lender: lender, Debtor: debtor, Date: date, Amount: amount, Comment: comment, Kind: kind}, nil
}
