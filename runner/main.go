package main

import (
  "fmt"
  "github.com/jasonrdsouza/banktorrent"
  "log"
)


func main() {
  fmt.Println("hello")

  a := banktorrent.User{"jason", "jason@gmail.com", banktorrent.Amount{1,0}}
  b := banktorrent.User{"justin", "justin@gmail.com", banktorrent.Amount{3,5}}
  c := banktorrent.User{"jun", "jun@gmail.com", banktorrent.Amount{6,2}}
  d := banktorrent.User{"steve", "steve@gmail.com", banktorrent.Amount{7,99}}

  l := banktorrent.UserMap {
    a.Email: a, 
    b.Email: b,
    c.Email: c,
    d.Email: d,
  }
  err := banktorrent.StoreUsers(banktorrent.USER_STORE, l)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("loading them")

  users, err := banktorrent.LoadUsers(banktorrent.USERSTORE)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(users)
}