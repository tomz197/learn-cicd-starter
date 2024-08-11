#!/bin/bash

if [ -f .env ]; then
    source .env
fi

cd sql/schema

echo "Print url for database"
echo $DATABASE_URL

goose turso $DATABASE_URL up
