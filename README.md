BankTorrent
-----------

BankTorrent is a distributed debt and payment tracking system. It allows
groups of people to keep track of shared expenses and how much they owe
eachother.

Author: Jason D'Souza


Todo
----
- Write ability
  - remove users/ labels
    - only if there are no expenses/ transactions/ associated with them
    - cant remove if it would leave db in inconsistent state
  - make writing transactional
    - replace sql.Dx with sql.Tx
- Meddler
  - use the MoneyAmount type instead of ints
  - change expense and transaction structs to use MoneyAmount
  - make a meddler that converts from the int to MoneyAmount
- CLI
  - ability to backup db
  - remove expenses
  - get list of expenses
  - add labels
  - add users
  - get statistics
    - upcoming bills
    - potentially missing transactions
  - predictive user input
    - best guess based on input?
  - tests
- web frontend
  - graphs
    - visualization of transaction history
    - breakdown by label
- Tests
  - test coverage of all methods/ functions
- Integrate terminal prettying
  - https://github.com/wsxiaoys/terminal


Workflow
--------
- easily add split/ simple transactions from command line
  - specify user/ label by name
  - enter amount in dollars and cents
- visual reassurance that the math works out
  - show users balance before and after
- easily correct a mistake
  - edit recently added expense?
