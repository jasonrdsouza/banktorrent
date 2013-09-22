package main

import (
  "flag"
  "fmt"
  "log"
  "time"
  "bytes"
  "github.com/jasonrdsouza/banktorrent"
)


const (
  DATE_FMT = "2006-01-02"
  const tmpl = `
  Parameters
  ----------
  - Date:\t{{.Date}}
  - Label:\t{{.Label}}
  - Lender:\t{{.Lender}}
  - Debtor:\t{{.Debtor}}
  - Amount:\t{{.Amount}}
  - Kind:\t{{.Kind}}
  - Comment:\t{{.Comment}}
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

raw = new(RawParams)
func init() {
  flag.StringVar(&raw.Label, "Label", "miscellaneous", "The label for this expense")
  flag.StringVar(&raw.Lender, "Lender", "", "The lender for this expense")
  flag.StringVar(&raw.Debtor, "Debtor", "", "The debtor for this expense")
  flag.StringVar(&raw.Date, "Date", "", "The date that the expense took place")
  flag.IntVar(&raw.Amount, "Amount", 0, "The amount of this expense")
  flag.StringVar(&raw.Comment, "Comment", "", "Comments associated with this expense")
  flag.StringVar(&raw.Kind, "Expense Kind", "", "What kind of expense this is")
}

func main() {
  flag.Parse()
  fmt.Println("Input params: ", raw)
  valid, err := validateParams(raw)
  if err != nil {
    log.Fatalln("Failed to validate params: ", err)
  }
  fmt.Println("Validated params: ", valid)
  fmt.Println("Adding Expense")
}

func (r *RawParams) String() (string) {
  t := template.Must(template.New("raw_tmpl").Parse(tmpl))
  var buf = bytes.Buffer
  err := t.Execute(buf, r)
  if err != nil {
    return ""
  }
  return buf.String()
}

func (v *ValidParams) String() (string) {
  t := template.Must(template.New("valid_tmpl").Parse(tmpl))
  var buf = bytes.Buffer
  err := t.Execute(buf, v)
  if err != nil {
    return ""
  }
  return buf.String()
}

func validateParams(raw RawParams) (ValidParams, error) {

}
