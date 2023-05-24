---
title: Podcasting your favourite tracks without copyright worries
date: 2009-11-16T16:31:00+00:00
draft: false
emoji: 🎙️
images:
tags:
  - idea
  - playdar
  - xspf
  - from-tumblr
---

> I had an idea last night, I'm going to sketch it out here it may turn into more than an idea at some point!

If you're a hobby podcaster or webradio DJ then you may find yourself limited to the songs you can use without having to hunt down a licence or permission. Despite the fact that its [easier all the time](http://www.ukpa.info/2008/08/22/the-new-mcps-prs-podcast-license/) it could be time consuming enough to make it not worth while.

Having become quite interested in [Playdar](https://playdar.org) I came up with an idea for making podcasting even the most protected of tracks (say the [Happy Birthday song](http://www.unhappybirthday.com/)[^1]) totally legal.

In brief I propose an extension of the [XSPF](https://www.xspf.org/) playlist format to include instructions for playing tracks at the same time. This will allow Joe Podcaster to record a voice track (chopped up into segments), upload them and include them in a playlist that includes these voice segments as well as defining resolvable names for the songs in the podcast, each with a description of when each track (speech or music) should be played relative to each other.

You podcast listener will retrieve the extended XSPF, download the unresolvable elements (the podcaster's voice) and resolve the copyrighted tracks from whatever sources are available, ready to play - because the songs are resolved locally there is no copyright infringement!

A program on the listener's computer can compile an audio file ready for use in any music player, or programs could be extended to support this format.

Below is an example extended XSPF[^2] I've sketched out based on version 0 of the XSPF file format (which is [available online](https://www.xspf.org/xspf-v0.html)).

I hope this makes vague sense, I've run out of time to write more! I'm bound to make massive changes to this idea as it progresses - what are your thoughts?

[^1]: Import note: Happy Birthday is now [in the public domain](https://en.wikipedia.org/wiki/Happy_Birthday_to_You#:~:text=%22Happy%20Birthday%20to%20You%22%20was%20in%20the%20public%20domain)!
[^2]: Except there isn't… this has been lost somewhere in Tumblr.
