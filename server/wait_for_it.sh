#!/bin/bash

set -e

host="$DB_HOST"
port="$DB_PORT"
cmd="$@"

>&2 echo "!!!!!!!! Check db for available !!!!!!!!"

until netcat -w 2 "$DB_HOST" "$DB_PORT"; do
  >&2 echo "db is unavailable - sleeping"
  sleep 1
done

>&2 echo "db is up - executing command"

exec $cmd