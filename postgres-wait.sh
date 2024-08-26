#!/bin/sh
# postgres-wai.sh

set -e

host="$1"
shift
cmd="$@"

until PGPASSWORD=$DATABASE_PASSWORD psql -h "$host" -U "tonitaga" -d "url-shortener" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec $cmd