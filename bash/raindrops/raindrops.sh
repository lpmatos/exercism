#!/usr/bin/env bash

# Make script exit when there is an error
set -o errexit

# Main function
function main(){
  local THREE=$(expr $1 % 3) FIVE=$(expr $1 % 5) SEVEN=$(expr $1 % 7) RESULT=""
  [ $THREE -eq 0 ] && RESULT+="Pling"
  [ $FIVE -eq 0 ] && RESULT+="Plang"
  [ $SEVEN -eq 0 ] && RESULT+="Plong"
  [ -z ${RESULT} ] && echo $1 || echo $RESULT
}

main "$@"
