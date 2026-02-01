---
title: SlyPIs out in the wild
date: 2009-05-24T20:52:00+01:00
draft: false
emoji: 🕵️
images:
tags:
  - ruby
  - code
  - slypi
  - gem
  - screen scraping
  - from-tumblr
atUri: "at://did:plc:ephkzpinhaqcabtkugtbzrwu/site.standard.document/3mdryr2iv4w2o"
---
I get really frustrated when websites don't have APIs. I know its kinda mean to screen scrape, and potentially prevented by ToCs etc, but I built this anyway. I have finals exams to revise for and not enough to procrastinate with!

So, what is all this SlyPI nonsense about then?

A SlyPI is a small ([YAML](https://en.wikipedia.org/wiki/YAML) based) settings file that describes how to mechanically traverse a website in order to get a variety of different bits of information. For example, my [tv.com example](https://github.com/jphastings/SlyPI-examples/blob/d6513b87c08b356c2877f546477bdd11b2df5362/slypis/tv.com.slypi) will allow you to search their database for shows by name, then use their 'showid' to find out about that show's episodes, its air times, genre and so on.

These settings files on their own aren't particularly helpful. Lucky for you I built a [ruby gem](https://github.com/jphastings/SlyPI) that interprets them and creates a class with methods pertaining to that slypi, so you just do:

```ruby
require 'slypi'
s = SlyPI.new("tv.com.slypi")
s.SearchShows(:q => "Terminator")
```

And you'll get a list of shows that have the word _Terminator_ in them. Cool ne?

So if a SlyPI exists for the site you want to use this is simplicity itself, but why bother making one for a new site? When you screen scrape you run the risk that the site layout will change and everything will be broken. With SlyPI even if it doesyou only need to edit the SlyPI settings file and order is restored. You could feasibly write an application that depended on screen scraping that auto-updated when a newer version of the slypi was available.

Naturally, people get frustrated at screen scraping, so please remember that you use this at your own risk, and definitely don't attempt to download the entire of a website though it...
