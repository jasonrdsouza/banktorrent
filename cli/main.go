package main

import (
  "flag"
  "fmt"
  //"log"
  "time"
  "bytes"
  "text/template"
  "github.com/jasonrdsouza/banktorrent"
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
  Amount int
  Comment string
  Kind string
}

type ValidParams struct {
  Label banktorrent.Label
  Lender banktorrent.User
  Debtor banktorrent.User
  Date time.Time
  Amount banktorrent.MoneyAmount
  Comment string
  Kind banktorrent.ExpenseType
}

var raw = new(RawParams)
func init() {
  flag.StringVar(&raw.Label, "label", "misc", "The label for this expense")
  flag.StringVar(&raw.Lender, "lender", "bob", "The lender for this expense")
  flag.StringVar(&raw.Debtor, "debtor", "alice", "The debtor for this expense")
  flag.StringVar(&raw.Date, "date", "2013-09-31", "The date that the expense took place")
  flag.IntVar(&raw.Amount, "amount", 0, "The amount of this expense")
  flag.StringVar(&raw.Comment, "comment", "sample cli expense", "Comments associated with this expense")
  flag.StringVar(&raw.Kind, "kind", "simple", "What kind of expense this is")
}

func main() {
  flag.Parse()
  fmt.Println("Input params: ", raw)
  // valid, err := validateParams(raw)
  // if err != nil {
  //   log.Fatalln("Failed to validate params: ", err)
  // }
  // fmt.Println("Validated params: ", valid)
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

func validateParams(raw RawParams) (*ValidParams, error) {
  return nil, nil
}
