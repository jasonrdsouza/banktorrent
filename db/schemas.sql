-- Schemas for all the tables that banktorrent uses

-- Users Table
CREATE TABLE users (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, 
  "name" VARCHAR(255), 
  "email" VARCHAR(255), 
  "balance" INTEGER
);

-- Labels Table
CREATE TABLE "labels" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, 
  "name" VARCHAR(255)
);

-- Expenses Table
CREATE TABLE "expenses" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, 
  "amount" INTEGER, 
  "label_id" INTEGER, 
  "date" VARCHAR(255), 
  "comment" VARCHAR(255),
  FOREIGN KEY(label_id) REFERENCES labels(id)
);

-- Transactions Table
CREATE TABLE "transactions" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, 
  "amount" INTEGER,
  "lender_id" INTEGER,
  "debtor_id" INTEGER,
  "date" VARCHAR(255), 
  "expense_id" INTEGER,
  FOREIGN KEY(lender_id) REFERENCES users(id),
  FOREIGN KEY(debtor_id) REFERENCES users(id),
  FOREIGN KEY(expense_id) REFERENCES expenses(id)
);