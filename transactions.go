package banktorrent

import (
  "time"
  "io/ioutil"
  "encoding/json"
)


const GROUP_STORE = "groups.banktorrent"

type GroupMap map[GroupKey]TransactionGroup

type TransactionKey int
type TransactionSessionKey int
type GroupKey int

type Amount struct {
  Dollars int
  Cents int
}

type Transaction struct {
  Id TransactionKey
  From UserKey
  To UserKey
  Amount Amount
}

type TransactionSession struct {
  Id TransactionSessionKey
  Transactions []Transaction
  Note string
  Timestamp time.Time
}

type TransactionGroup struct {
  Id GroupKey
  Name string
  Balances map[UserKey]Amount
  Transactions []TransactionSession
}


func CreateGroup(name string, users []UserKey) (GroupKey, error) {

}

func (t TransactionGroup) AddUser(user UserKey) (error) {

}

func (t TransactionGroup) Users() ([]UserKey, error) {
  
}

// Adding a transaction really means adding a "transaction session". 
// a transaction session is a discrete unit potentially comprised of multiple 
// transactions. 
// For example, if you bought something to share amongst the members of your
// transaction group, you could add that purchase in a single transaction session,
// which translates to one transaction per person.
func (t TransactionGroup) AddTransaction(transaction TransactionSession) error {

}

// Removes erroneous transactions from the transaction group history and
// readjusts the group members balances accordingly
func (t TransactionGroup) DeleteTransaction(transaction_session_id int) error {

}

// Loads the transaction groups from the given file on disk
func LoadGroups(transactionstore string) (GroupMap, error) {
  f, err := ioutil.ReadFile(transactionstore)
  if err != nil { 
    return nil, err
  }
  var groups GroupMap
  err = json.Unmarshal(f, &groups)
  if err != nil {
    return nil, err
  }
  return groups, nil
}

// Writes the given transaction groups to disk
func StoreGroups(transactionstore string, groups GroupMap) (error) {
  b, err := json.Marshal(groups)
  if err != nil {
    return err
  }
  err = ioutil.WriteFile(transactionstore, b, 0644)
  if err != nil {
    return err
  }
  return nil
}