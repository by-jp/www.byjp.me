#!/bin/bash

rm **/_redirects

while IFS=$'\n' read -r line
do
  # Extract the absolute file path
  outdir=$(dirname $(echo "$line" | cut -d " " -f1) | cut -d "/" -f2)

  if [[ -n "$outdir" ]]; then
    [[ ! -d $outdir ]] && mkdir -p $outdir
    echo $line | cut -c $((${#outdir}+2))- >> "./$outdir/_redirects"
  else
    echo $line  >> "_redirects"
  fi
done < "${1:-/dev/stdin}"
