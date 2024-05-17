---
title: "git download-subpath"
emoji: 🤓
date: 2024-05-14T13:08:10+01:00
summary: I made a git helper tool that lets you retrieve files from a subpath of git without downloading the whole of the (possibly huge) repo.
tags:
- git
- computers
- tools
topics:
- Technology
syndications:
- https://gist.github.com/jphastings/6560bb173399fc3a155913b33e5f0f0c
---

My blog has ~250MB of [photos](/photos) in it (as I [archived my Instagram here](/posts/archiving-instagram-posts/) — _a decade_ of photos) which means that every time I want some files from [my blog's repo](https://github.com/by-jp/www.byjp.me) I've needed to pull down all that data to get at the few files I want.

This isn't a problem on my laptop where I work on my blog (as it's already cloned), but I also keep my [IndieKit](https://getindiekit.com) config in there, which I need to copy to the server it runs from. I don't want to have to download hundreds of MB (and growing!) of photo data every time I want to update it.

So I built `git-download-subpath`, which is a bash script around [git's partial clone](https://git-scm.com/docs/partial-clone) functionality.

```bash
$ git download-subpath
# usage: git-download-subpath <repo> <subpath> [destination]

$ git download-subpath https://github.com/by-jp/www.byjp.me indiekit
# Successfully downloaded to ./indiekit
```

This script (which you can find & download below) completes these steps:

1. Clones a "no tree" copy of the repo to a temporary directory (just references to commits & the latest files, not the data itself)
2. Completes a "no cone" sparse checkout of the subpath desired
3. Moves that desired subpath over the working directory (or the destination, if specified)

I'm no bash expert, so there may be subtle bugs here; let me know if you spot them!

_Copy the following bash to `git-download-subpath` somewhere in your `$PATH` (I keep it in `/usr/local/bin`), and mark as executable with `chmod +x git-download-subpath`:_

```bash /usr/local/bin/git-download-subpath
#!/bin/bash
set -e

repo="$1"
subpath="$2"
if [[ -z "${repo}" || -z "${subpath}" ]]; then
  echo "usage: git-download-subpath <repo> <subpath> [destination]"
  exit 1
fi

if [[ -z "$3" ]]; then
  dest=$(pwd)
else
  dest=$(realpath "$3")
fi
if [[ ! -d "${dest}" ]]; then
  >&2 echo "${dest} is not a directory that exists"
  exit 3
fi
if [[ -d "${dest}/${subpath}" ]]; then
  >&2 echo "${dest}/${subpath} already exists, exiting to avoid replacing"
  exit 3
fi

# Cross-platform create temporary dir: https://unix.stackexchange.com/a/84980/
tmpdir=$(mktemp -d 2>/dev/null || mktemp -d -t 'git-download-subpath')
cd "${tmpdir}"

git clone -n --depth=1 --filter=tree:0 "${repo}" repo > /dev/null 2>&1
cd repo
git sparse-checkout set --no-cone "${subpath}" > /dev/null 2>&1
git checkout > /dev/null 2>&1
if [[ ! -d "${subpath}" ]]; then
  >&2 echo "The directory '${subpath}' doesn't exist in this repo."
  exit 5
fi
cp -r "${subpath}" "${dest}/"

echo "Successfully downloaded to ${dest}/${subpath}"
```
