package banktorrent

import (
  "errors"
  "fmt"
  "strings"
  "strconv"
)


type MoneyAmount struct {
  Dollars int
  Cents int
}

func IntToMoney(amount int) (*MoneyAmount) {
  dollars := amount / 100
  cents := amount % 100
  return &MoneyAmount{Dollars: dollars, Cents: cents}
}

func StringToMoney(amount string) (*MoneyAmount, error) {
  raw_amount := strings.Split(strings.TrimLeft(strings.TrimSpace(amount), "$"), ".")
  ma := new(MoneyAmount)
  dollars, err := strconv.Atoi(raw_amount[0])
  if err != nil {
    return nil, err
  }
  if len(raw_amount) < 2 { // no cents given
    raw_amount = append(raw_amount, "00")
  }
  cents, err := strconv.Atoi(raw_amount[1])
  if err != nil {
    return nil, err
  }
  if cents > 99 {
    return nil, errors.New("Invalid cents amount... cannot be greater than 99")
  }
  ma.Dollars = dollars
  ma.Cents = cents
  return ma, nil
}

func (m *MoneyAmount) String() (string) {
  return fmt.Sprintf("$%d.%d", m.Dollars, m.Cents)
}

// Return the money amount as a regular int
// useful for saving to the db until I figure out how to use Meddler to do it
func (m *MoneyAmount) Int() (int) {
  return (m.Dollars * 100 + m.Cents)
}
