-- SQL to populate the db with test data

-- Populate the users table
--                      id, name,     email,         balance
INSERT INTO users VALUES(1,'User One','user-one@test.com',4500);
INSERT INTO users VALUES(2,'User Two','user-two@test.com',-2500);
INSERT INTO users VALUES(3,'User Three','user-three@test.com',-2000)

-- Populate the labels table
--                       id, name
INSERT INTO labels VALUES(1,'groceries');
INSERT INTO labels VALUES(2,'internet');
INSERT INTO labels VALUES(3,'utilities');
INSERT INTO labels VALUES(4,'miscellaneous');

-- Populate the expenses table
--                         id, amount, label_id, date, comment
INSERT INTO expenses VALUES(1,2000,1,'2013-07-01','grocery test expense');
INSERT INTO expenses VALUES(2,500,4,'2013-07-02','simple miscellaneous expense');
INSERT INTO expenses VALUES(3,6000,2,'2013-07-03','internet split 3 ways');

-- Populate the transactions table
--          id, amount, lender_id, debtor_id, date, expense_id
INSERT INTO transactions VALUES(1,1000,1,2,'2013-07-01',1);
INSERT INTO transactions VALUES(2,500,2,1,'2013-07-02',2);
INSERT INTO transactions VALUES(3,2000,1,2,'2013-07-03',3);
INSERT INTO transactions VALUES(4,2000,1,3,'2013-07-03',3);