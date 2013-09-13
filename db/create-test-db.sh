#!/bin/bash

TEST_DB="/tmp/banktorrent.test.db"
SCHEMAS="/home/jason/gocode/src/github.com/jasonrdsouza/banktorrent/db/schemas.sql"
DATA="/home/jason/gocode/src/github.com/jasonrdsouza/banktorrent/db/test-data.sql"

echo "Removing the old test db"
rm $TEST_DB

echo "Generating fresh test db"
sqlite3 $TEST_DB < $SCHEMAS
sqlite3 $TEST_DB < $DATA

echo "Finished creating new test db at $TEST_DB"