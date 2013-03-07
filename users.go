package banktorrent

import (
  "io/ioutil"
  "encoding/json"
)


type User struct {
  Name string
  Email string
  Balance Amount
  Groups []int
}

type UserKey string

type UserMap map[UserKey]User

const USER_STORE = "users.banktorrent"

// Retrieves users from the given file on disk
func LoadUsers(userstore string) (UserMap, error) {
  f, err := ioutil.ReadFile(userstore)
  if err != nil { 
    return nil, err
  }
  var users UserMap
  err = json.Unmarshal(f, &users)
  if err != nil {
    return nil, err
  }
  return users, nil
}

// Writes the given users to disk
func StoreUsers(userstore string, users UserMap) (error) {
  b, err := json.Marshal(users)
  if err != nil {
    return err
  }
  err = ioutil.WriteFile(userstore, b, 0644)
  if err != nil {
    return err
  }
  return nil
}
