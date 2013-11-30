package main

import (
  "fmt"
  "os"
  "log"
  "time"
  "bytes"
  "strings"
  "text/template"
  "github.com/jasonrdsouza/banktorrent"
  "database/sql"
  "github.com/codegangsta/cli"
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


// Currently, this only works with 2 users... need to add the ability
// to add split expenses with multiple users, meaning taking a list of
// users with the lender flag.
func main() {
  db, err := banktorrent.Connect(banktorrent.PROD_DB)
  if err != nil {
    log.Fatalln(err)
  }
  defer db.Close()

  app := cli.NewApp()
  app.Name = "BankTorrent"
  app.Usage = "Easy, distributed debt and payment tracking system"

  command_add := cli.Command{
      Name:       "add",
      Usage:      "add an expense",
      Flags:      []cli.Flag {
                    cli.StringFlag{"label", "misc", "The label for the expense(s)"},
                    cli.StringFlag{"lender, l", "bob", "The lender for the expense"},
                    cli.StringFlag{"debtor, d", "alice", "The debtor for the expense"},
                    cli.StringFlag{"date", "2013-09-31", "The date the expense took place"},
                    cli.StringFlag{"amount, a", "0.00", "The amount of the expense"},
                    cli.StringFlag{"comment, c", "", "Comments associated with the expense"},
                    cli.StringFlag{"type, t", "simple", "The type of expense this is"},
                  },
      Action:     func(c *cli.Context) {
                    valid, err := contextToValidParams(c, db)
                    if err != nil {
                      log.Fatalln("Could not add expense due to malformed paramaters. Error: ", err)
                    }
                    fmt.Println("Validated params: ", valid)
                    addExpense(db, valid)
                  },
  }
  command_dry_run := cli.Command{
      Name:       "dry-run",
      Usage:      "dry run of adding an expense",
      Flags:      []cli.Flag {
                    cli.StringFlag{"label", "misc", "The label for the expense(s)"},
                    cli.StringFlag{"lender, l", "bob", "The lender for the expense"},
                    cli.StringFlag{"debtor, d", "alice", "The debtor for the expense"},
                    cli.StringFlag{"date", "2013-09-31", "The date the expense took place"},
                    cli.StringFlag{"amount, a", "0.00", "The amount of the expense"},
                    cli.StringFlag{"comment, c", "", "Comments associated with the expense"},
                    cli.StringFlag{"type, t", "simple", "The type of expense this is"},
                  },
      Action:     func(c *cli.Context) {
                    valid, err := contextToValidParams(c, db)
                    if err != nil {
                      log.Fatalln("Could not add expense due to malformed paramaters. Error: ", err)
                    }
                    fmt.Println("Validated params: ", valid)
                    //fake_add_expense
                  },
  }

  app.Commands = []cli.Command{
    command_add,
    command_dry_run,
    //backup
    //edit
    //rm
    //show
  }

  app.Run(os.Args)
}

func contextToValidParams(c *cli.Context, db *sql.DB) (*ValidParams, error) {
  raw := &RawParams{Label: c.String("label"),
                    Lender: c.String("lender"),
                    Debtor: c.String("debtor"),
                    Date: c.String("date"),
                    Amount: c.String("amount"),
                    Comment: c.String("comment"),
                    Etype: c.String("type")}
  fmt.Println("Input params: ", raw)
  valid, err := validateParams(db, raw)
  if err != nil {
    return nil, err
  }
  return valid, nil
}

func addExpense(db *sql.DB, valid *ValidParams) {
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
    log.Println("Error validating label: ", raw.Label)
    return nil, err
  }
  lender, err := banktorrent.GetUserByName(db, strings.TrimSpace(raw.Lender))
  if err != nil {
    log.Println("Error validating lender: ", raw.Lender)
    return nil, err
  }
  debtor, err := banktorrent.GetUserByName(db, strings.TrimSpace(raw.Debtor))
  if err != nil {
    log.Println("Error validating debtor: ", raw.Debtor)
    return nil, err
  }
  date, err := time.Parse(banktorrent.DATE_FMT, strings.TrimSpace(raw.Date))
  if err != nil {
    log.Println("Error validating date: ", raw.Date)
    return nil, err
  }
  amount, err := banktorrent.StringToMoney(raw.Amount)
  if err != nil {
    log.Println("Error validating amount: ", raw.Amount)
    return nil, err
  }
  comment := strings.TrimSpace(raw.Comment)
  etype, err := banktorrent.StringToExpenseType(raw.Etype)
  if err != nil {
    log.Println("Error validating expense type: ", raw.Etype)
    return nil, err
  }

  return &ValidParams{Label: label, Lender: lender, Debtor: debtor, Date: date, Amount: amount, Comment: comment, Etype: etype}, nil
}
