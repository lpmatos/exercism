#!/usr/bin/env bash

# Make script exit when there is an error
set -o errexit

# Main function
function main(){
  # [ -z "${1}" ] && echo "One for you, one for me." || echo "One for $1, one for me."
  local NAME=$1 DEFAULT="you"
  echo "One for ${NAME:-$DEFAULT}, one for me."
}

main "$@"
