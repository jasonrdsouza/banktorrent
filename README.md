BankTorrent
-----------

BankTorrent is a distributed debt and payment tracking system. It allows
groups of people to keep track of shared expenses and how much they owe
eachother. 

Author: Jason D'Souza


Todo
----
- rename "test_error_helper"
- Setup foreign key relationships
  - add to db/schemas file
- Write ability
  - ability to add labels
  - ability to add users
  - make writing transactional
    - replace sql.Dx with sql.Tx
- Tests
  - make all tests pass
  - test coverage of all methods/ functions
- CLI 
  - ability to backup db
  - add transactions
  - add labels
  - add users
  - get statistics
    - upcoming bills
    - potentially missing transactions
- web frontend
  - build off of the cli
  - graphs
    - visualization of transaction history
    - breakdown by label


Workflow
--------
- easily add split/ simple transactions from command line
  - specify user/ label by name
  - enter amount in dollars and cents
- visual reassurance that the math works out
  - show users balance before and after
- easily correct a mistake
  - edit recently added expense?
