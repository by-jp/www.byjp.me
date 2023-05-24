---
title: Geo Tumblr
date: 2009-08-31T15:45:00+01:00
draft: false
emoji: 🗺️
images:
tags:
  - code
  - geo
  - from-tumblr
---

Well, I'm back from Japan and while I was there I kept an online journal (Tumblr, you rock) to let everyone at home know what was going on. To make things even cooler I vowed to keep the posts tagged with where I was while I was writing them, of course tumblr has no functionality for this so I hacked up a geotagging solution using javascript.

After making a [formal request for geotagging features](https://getsatisfaction.com/tumblr/topics/geotagging_of_individual_posts)[^1] on their GetSatisfaction pages, someone asked me how I got it done, so I created a (very sloppy) [tutorial as a github gist](https://gist.github.com/jphastings/178487). If you feel like getting some rather hacky geotagging features in your tumblog, this is one way to do it with google maps and plenty of time!

The Cons:

- You have to find your lat and long in decimal format in order to tag your post with a `geo:51;-1`. This was very annoying when my iPhone doesn't really have this ability. (I used tweetie when it was offline, which posts the google map link of your position when you ask it to)
- Many of the posts I wrote in more than one place (long posts, written in the evening about a place I'd left, or about more than one place) which meant there was no obvious single location for the post.
- There's a painful javascript yank on any system visiting the site, due to the amount of computing that needs to be done to work around the tumblr system, which was never designed to do this. (Its not that bad really, but just irks me as it shouldn't be necessary!)

And yes, Japan was amazing.

[^1]: Import note: GetSatisfaction has [gone offline](https://en.wikipedia.org/wiki/Get_Satisfaction), so this forml request has succumbed to link rot.
