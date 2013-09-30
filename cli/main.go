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
  param_tmpl = `
  Parameters
  ----------
  - Date: {{.Date}}
  - Label:  {{.Label}}
  - Lender: {{.Lender}}
  - Debtor: {{.Debtor}}
  - Amount: {{.Amount}}
  - Type: {{.Etype}}
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
  Etype string
}

type ValidParams struct {
  Label *banktorrent.Label
  Lender *banktorrent.User
  Debtor *banktorrent.User
  Date time.Time
  Amount *banktorrent.MoneyAmount
  Comment string
  Etype banktorrent.ExpenseType
}

var raw = new(RawParams)
func init() {
  flag.StringVar(&raw.Label, "label", "misc", "The label for this expense")
  flag.StringVar(&raw.Lender, "lender", "bob", "The lender for this expense")
  flag.StringVar(&raw.Debtor, "debtor", "alice", "The debtor for this expense")
  flag.StringVar(&raw.Date, "date", "2013-09-31", "The date that the expense took place")
  flag.StringVar(&raw.Amount, "amount", "25.00", "The amount of this expense")
  flag.StringVar(&raw.Comment, "comment", "sample cli expense", "Comments associated with this expense")
  flag.StringVar(&raw.Etype, "type", "simple", "What type of expense this is")
}

// Currently, this only works with 2 users... need to add the ability
// to add split expenses with multiple users, meaning taking a list of
// users with the lender flag.
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
    fmt.Println("Could not add expense due to malformed paramaters :(")
    log.Fatalln("Failed to validate params: ", err)
  }
  fmt.Println("Validated params: ", valid)

  expense, err := banktorrent.CreateExpense(db, valid.Amount, valid.Label, valid.Comment, valid.Date)
  fmt.Printf("Adding %v as a %v expense\n", expense, valid.Etype)
  switch valid.Etype {
  case banktorrent.SIMPLE_EXPENSE:
    transaction, err := expense.AddSimple(db, valid.Lender, valid.Debtor)
    if err != nil {
      fmt.Println("Error adding simple expense")
      log.Fatalln(err)
    }
    fmt.Println("Expense transaction:\n", transaction)
  case banktorrent.SPLIT_EXPENSE:
    transactions, err := expense.AddSplit(db, valid.Lender, []*banktorrent.User{valid.Debtor})
    if err != nil {
      fmt.Println("Error adding split expense")
      log.Fatalln(err)
    }
    fmt.Println("Expense transactions:\n", transactions)
  }

  // reload the users to see their new balances
  err = valid.Lender.Reload(db)
  if err != nil {
    fmt.Println("Error reloading lender from DB")
    log.Fatalln(err)
  }
  err = valid.Debtor.Reload(db)
  if err != nil {
    fmt.Println("Error reloading debtor from DB")
    log.Fatalln(err)
  }
  fmt.Printf("New user balances:\n- Lender => %v\n- Debtor => %v\n", valid.Lender, valid.Debtor)
}

func (r *RawParams) String() (string) {
  t := template.Must(template.New("raw_tmpl").Parse(param_tmpl))
  var buf = new(bytes.Buffer)
  err := t.Execute(buf, r)
  if err != nil {
    return ""
  }
  return buf.String()
}

func (v *ValidParams) String() (string) {
  t := template.Must(template.New("valid_tmpl").Parse(param_tmpl))
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
  date, err := time.Parse(banktorrent.DATE_FMT, strings.TrimSpace(raw.Date))
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
  etype, err := banktorrent.StringToExpenseType(raw.Etype)
  if err != nil {
    log.Println("Error validating expense type ", raw.Etype)
    return nil, err
  }

  return &ValidParams{Label: label, Lender: lender, Debtor: debtor, Date: date, Amount: amount, Comment: comment, Etype: etype}, nil
}
