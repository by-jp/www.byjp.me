#!/bin/bash
# Usage:
#  1. Build www.byjp.me
#  2. Clone and enter a local copy of the byjp.fyi repo
#  3. `cat www.byjp.me/public/index.redirects | www.byjp.me/tools/redirects/update-shortlinks.bash``
#  4. `git commit -am "Update blog redirects" && git push`

# This script assumes it is being executed at the root of the byjp.fyi repo
gitRoot=$(git rev-parse --show-toplevel) || { echo "Please run inside the byjp.fyi git repo"; exit 1; }
if [[ "$(git remote get-url origin)" != "https://github.com/by-jp/byjp.fyi" ]]; then
  echo "This repo doesn't have the byjp.fyi repo as the origin"
  git remote -v
  exit 2
fi
cd $gitRoot

# Prepare existing files for new _redirects info
mv public/_redirects public/_redirects.previous
find public -type f -iname "_redirects" -delete

# Create local _redirects files
echo "Sorting shortlinks into different directory _redirects files…"

while IFS=$'\n' read -r line
do
  # Extract the absolute file path
  outdir=$(dirname $(echo "$line" | cut -d " " -f1) | cut -d "/" -f2)

  if [[ -n "$outdir" ]]; then
    [[ ! -d $outdir ]] && mkdir -p "public/$outdir"
    echo $line | cut -c $((${#outdir}+2))- >> "public/$outdir/_redirects"
  else
    echo $line  >> "public/_redirects"
  fi
done < /dev/stdin

# Combine new (root) _redirects and _redirects.previous
echo "Merging with existing links…"

start_line=$(awk '/^# From blog/{print NR; exit}' public/_redirects.previous)
end_line=$(awk 'NR > '"$start_line"' && /^#/{print NR; exit}' public/_redirects.previous)

sed "$((start_line + 1)),$((end_line - 2))d" public/_redirects.previous > old_removed.txt
awk 'NR=='"$((start_line + 1))"'{system("cat public/_redirects")} 1' old_removed.txt > public/_redirects.new

# Remove temporary files
mv public/_redirects.new public/_redirects
rm public/_redirects.previous
rm old_removed.txt

# Push changes
echo "Pushing to github…"

git add -A
git commit -m "Update blog shortlinks"
git push -q

echo "All files cleaned up; update complete!"
