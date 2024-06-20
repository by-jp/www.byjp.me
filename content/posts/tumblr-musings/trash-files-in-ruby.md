---
title: Trash files in Ruby
date: 2009-03-25T22:42:00+00:00
draft: false
emoji: 🗑️
summary: A Ruby gem I wrote that sends files to Operating System specific Trash/Recycle bins.
tags:
  - ruby
  - code
  - github
  - from-tumblr
---
[github.com](https://github.com/jphastings/trash)

I’ve written a ruby `File` class extension - `File.trash( filename )` - which allows you to send files to the Trash or Recycle bin (depending on your OS).

Mac support is ready, Windows support is touch-and-go (read the [wiki](https://github.com/jphastings/trash/wiki) for more info) and its been so long since I used Linux (ahh, the [Arch](https://www.archlinux.org/) days) I can’t remember how the bins work over there. Anyone want to help?
