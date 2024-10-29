---
date: 2024-10-29T10:35:52.127Z
publishDate: 2024-10-29T10:35:52.127Z
tags:
  - overmind
  - rails
---

I just threw together a script for easily attaching a debugger to any of your [overmind](/tags/overmind) run processes. let me know if it's useful!

I'm using it for attaching to my [rails](/tags/rails) processes, as I don't like how `overmind connect web` echoes _everything_ the web process outputs.

```sh bin/debug
#!/usr/bin/env sh

if [[ "$1" == "--help" ]]; then
  echo "Usage: $0 [web/worker/etc…]"
  echo
  echo "Attach an rdbg debugger to the named overmind process. (default: web)"
  exit 0
fi

PROCESS=${1:-web}
PROCESS_PID=$(overmind ps | grep -E "^${PROCESS}" | tr -s ' '  | cut -d' ' -f2)
[[ -z $PROCESS_PID ]] && { echo "There is no running overmind process called '$PROCESS'" >&2; exit 1; }
DEBUGGER_PID=$(pgrep -P $PROCESS_PID)
[[ -z $DEBUGGER_PID ]] && { echo "The $PROCESS process doesn't seem to have any child processes" >&2; exit 1; }

bundle exec rdbg -An "rdbg-$DEBUGGER_PID"
```
