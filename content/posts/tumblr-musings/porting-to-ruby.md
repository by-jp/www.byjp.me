---
title: Porting to Ruby
date: 2009-05-05T04:08:00+01:00
draft: false
emoji: ⬇️
tags:
  - ruby
  - code
  - dlc
  - from-tumblr
---
I had some time to kill today, so I rewrote the [DLC api](https://www.jdownloader.org/news/blog/x20090501-002-dlcapi-oop) in Ruby.

If you use [JDownloader](https://jdownloader.org) you’ll know how useful the DLCs are, essentially they allow you to contain all the links from a variety of online hosting sources in one place with any passwords that might be needed. All this data is also kept encrypted to prevent people from poking around inside and stealing the links!

Now, should you wish to create one of these DLCs there’s a [handy ruby API](https://github.com/jphastings/ruby-DLC) I’ve built thats so simple you can even use it from `irb`! Enjoy.
