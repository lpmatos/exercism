#!/usr/bin/env bash

# Make script exit when there is an error
set -o errexit

# Handling function
function handling(){
  echo "Usage: error_handling.sh <person>" && exit 1
}

# Main function
function main(){
  [ "$#" -ne 1 ] && handling || echo "Hello, $*"
}

# $# = number of arguments. Answer is 3
# $@ = what parameters were passed. Answer is 1 2 3
# $? = was last command successful. Answer is 0 which means 'yes'

main "$@"
