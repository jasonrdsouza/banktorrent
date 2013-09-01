-- Expenses Table
CREATE TABLE "expenses" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, 
  "amount" INTEGER, 
  "label_id" INTEGER, 
  "date" VARCHAR(255), 
  "comment" VARCHAR(255)
);

-- Labels Table
CREATE TABLE "labels" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, 
  "name" VARCHAR(255)
);

-- Transactions Table
CREATE TABLE "transactions" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, 
  "amount" INTEGER, 
  "lender_id" INTEGER, 
  "debtor_id" INTEGER, 
  "date" VARCHAR(255), 
  "label_id" INTEGER, 
  "comment" VARCHAR(255)
);

-- Users Table
CREATE TABLE users (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, 
  "name" VARCHAR(255), 
  "email" VARCHAR(255), 
  "balance" INTEGER
);