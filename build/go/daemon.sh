#! /bin/sh

trap processUserSig SIGUSR1
processUserSig() {
  echo "doing stuff"
}

while true; do
  sleep 1000
done
