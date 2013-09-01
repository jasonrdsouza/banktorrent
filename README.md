BankTorrent
-----------

BankTorrent is a distributed debt and payment tracking system. It allows
groups of people to keep track of shared expenses and how much they owe
eachother. 

Author: Jason D'Souza


Todo
----

- [ ] Read-only ability
  - [x] User model, helpers, tests
  - [x] Transactions model, helpers, tests
    - [x] migrate schema over
  - [x] Labels model, helpers, tests
    - [x] migrate schema over
- [ ] Write ability
  - [ ] expenses
    - [ ] split transaction
      - [ ] N person split
    - [ ] regular transaction
  - [ ] ability to add labels
  - [ ] ability to add users
- [ ] CLI 
  - [ ] ability to backup db
  - [ ] add transactions
  - [ ] add labels
  - [ ] add users
  - [ ] get statistics
    - [ ] upcoming bills
    - [ ] potentially missing transactions
- [ ] web frontend
  - [ ] build off of the cli
  - [ ] graphs
    - [ ] visualization of transaction history
    - [ ] breakdown by label

