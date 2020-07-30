#!/usr/bin/env bash

# Make script exit when there is an error
set -o errexit

# Main function
function main(){
  if ! [[ $1 =~ ^[0-9]+$ ]]; then echo "We need a number" && exit 1; fi
  local THREE=$(expr $1 % 3) FIVE=$(expr $1 % 5) SEVEN=$(expr $1 % 7) RESULT=""
  [ $THREE -eq 0 ] && RESULT+="Pling"
  [ $FIVE -eq 0 ] && RESULT+="Plang"
  [ $SEVEN -eq 0 ] && RESULT+="Plong"
  echo ${RESULT:-$1}
}

main "$@"
