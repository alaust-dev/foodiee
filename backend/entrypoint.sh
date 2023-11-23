#!/bin/sh

if [[ ! -f $DB_FILE_PATH ]]; then
    echo "Sqlite DB does not exist, creating one..."
    sqlite3 $DB_FILE_PATH < schema.sql
fi

exec ./foodiee-backend
