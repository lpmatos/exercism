#!/usr/bin/env bash

# Option to exit with any command fails
set -e

function main(){
  echo "Hello, World!"
}

# Just call main function using special param to get all positional arguments.
# When we put this special param in double quotes, we prevent whitespace.
main "$@"
