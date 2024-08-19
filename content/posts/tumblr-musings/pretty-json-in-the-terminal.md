---
title: Pretty JSON in the terminal
date: 2013-03-20T16:02:00+00:00
draft: false
emoji: ✨
tags:
  - terminal
  - tools
  - from-tumblr
---
[kaspertidemann.com](https://web.archive.org/web/20141024021149/http://www.kaspertidemann.com/pretty-printing-json-in-the-terminal/)

This was _exactly_ what I was looking for.

If you're investingating JSON returned from a web service, get the control you need with the terminal and [jsonpp](https://github.com/jmhodges/jsonpp):

```sh
brew install jsonpp
curl -u username:password https://awesome.com/stuff.json | jsonpp 
```

BOOM. As they say.
