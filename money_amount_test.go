package banktorrent

import (
  "testing"
)


func Test_IntToMoneyAmount(t *testing.T) {
  samples := []int{100, 4400, 65, 0, 1}
  results := make([]*MoneyAmount, len(samples))
  for index, sample := range samples {
    res := IntToMoney(sample)
    results[index] = res
  }
  if results[0].Dollars != 1 && results[0].Cents != 0 {
    t.Error("Expected $1.00, but got: ", results[0])
  }
  if results[1].Dollars != 44 && results[1].Cents != 0 {
    t.Error("Expected $44.00, but got: ", results[1])
  }
  if results[2].Dollars != 0 && results[2].Cents != 65 {
    t.Error("Expected $0.65, but got: ", results[2])
  }
  if results[3].Dollars != 0 && results[3].Cents != 0 {
    t.Error("Expected $0.00, but got: ", results[3])
  }
  if results[4].Dollars != 0 && results[4].Cents != 1 {
    t.Error("Expected $0.01, but got: ", results[4])
  }
}

func Test_StringToMoneyAmount(t *testing.T) {
  samples := []string{"12.01",
                      "$13.20",
                      " 14.33 ",
                      "15",
                      "$16",
                      "17.4"}
  results := make([]*MoneyAmount, 7)
  for index, sample := range samples {
    res, err := StringToMoney(sample)
    handle_error(t, err)
    results[index] = res
  }
  if results[0].Dollars != 12 && results[0].Cents != 1 {
    t.Error("Expected $12.01, but got: ", results[0])
  }
  if results[1].Dollars != 13 && results[1].Cents != 20 {
    t.Error("Expected $13.20, but got: ", results[1])
  }
  if results[2].Dollars != 14 && results[2].Cents != 33 {
    t.Error("Expected $14.33, but got: ", results[2])
  }
  if results[3].Dollars != 15 && results[3].Cents != 0 {
    t.Error("Expected $15.00, but got: ", results[3])
  }
  if results[4].Dollars != 16 && results[4].Cents != 0 {
    t.Error("Expected $16.00, but got: ", results[4])
  }
  if results[5].Dollars != 17 && results[5].Cents != 4 {
    t.Error("Expected $17.04, but got: ", results[5])
  }
}

func Test_MoneyAmountToInt(t *testing.T) {
  ma1 := MoneyAmount{Dollars: 10, Cents: 20}
  ma2 := MoneyAmount{Dollars: 23, Cents: 4}
  ma3 := MoneyAmount{Dollars: 4, Cents: 0}
  ma4 := MoneyAmount{Dollars: 0, Cents: 98}

  if ma1.Int() != 1020 {
    t.Error("Expected 1020, but got: ", ma1.Int())
  }
  if ma2.Int() != 2304 {
    t.Error("Expected 2304, but got: ", ma2.Int())
  }
  if ma3.Int() != 400 {
    t.Error("Expected 400, but got: ", ma3.Int())
  }
  if ma4.Int() != 98 {
    t.Error("Expected 98, but got: ", ma4.Int())
  }
}
