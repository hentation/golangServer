#!/bin/sh
# wait-for-it.sh: Ждет, пока хост и порт станут доступными.

set -e

host="$1"
port="$2"
shift 2
cmd="$@"

until ncat -z "$host" "$port"; do
  >&2 echo "Ждем, пока $host:$port станет доступным..."
  sleep 1
done

>&2 echo "$host:$port доступен"
exec $cmd
