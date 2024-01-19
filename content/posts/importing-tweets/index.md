---
title: "Importing Tweets"
emoji: 🐦
date: 2024-01-19T15:35:04Z
summary: More than a year after I closed my Twitter account, I've imported my tweets to my blog.
tags:
- twitter
- import
- code
- archive
- nostalgia
---

It wasn't until I started writing this post that I realised how long I was a user of Twitter. For 2 days shy of 13 years (from Nov 16, 2009 to Nov Nov 14, 2022) I posted [random thoughts](/notes/twitter/860849128724713472), experimented (with [geotagging](/notes/twitter/22679551), [OAuth](/notes/twitter/78268982) and more), and watched it change the internet with the invention of [hashtags](/notes/twitter/1412088964), retweets ([before](/notes/twitter/2460448329/) it was "via"), `@`tags ([before](/notes/twitter/74694)—while watching Ze Frank's The Show, [after](/notes/twitter/113686952)—when my friends all had handles that were just their names).

I even used it during its SMS-powered phase, where the 140 charcter limit came from; a tweet starting with your username ([as they did](/notes/twitter/70903) back then) allowed for 20 characters of username and 140 characters of thought was the 160 limit of SMS.

In late 2022 I became very worried about the (then new) management of Twitter, particularly the choices it was making in data security, so I decided to vote with my feet. I collected my data with their export functionality, checked it briefly, then put it on a backup drive and forgot about it. Until this weekend…

I threw together some (crappy) [Go](/tags/go) code that processes the data from a Twitter export `.zip` file and creates Hugo compatible markdown files for this blog. You can see the code _in_ this blog's repo (currently [on Github](https://github.com/by-jp/www.byjp.me/tree/main/tools/archive/twitter)).

This code also tries to rescue as much data from other silos as possible (like link shorteners and image hosters), but there's definitely plenty missing. If you find a `t.co` or `j.mp` shortener link still in there, then the data has already succumbed to linkrot, and I've left the (now dead) URL there in the unlikely event that I can pull info from an archive somewhere.

I _did_ manage to manually download my (few) [audioboo](/tags/audioboo) recordings from what's now [audioboom.com](https://audioboom.com), and serve them up here. If you want a few, short, mostly grumble-based, recordings of my voice from 15 years ago, you've come to the right place.

It's probably my age, and my [recent](/posts/mum) proximity to death, but I've really enjoyed travelling back into my past by reading through my early days of Twitter over-sharing. In fact, all the digital history work I've done around importing things into this blog has been thoroughly enjoyably nostalgic. I'd recommend it!

And on that note I'll leave you with a nostalgic track, the first I seem to have posted to Twitter, The Mistabishi Remix of The Temper Trap's _Science Of Fear_:

{{< spotify path="/track/0Z7CwV9ps38j0JvFmg43yr" artist="The Temper Trap" title="Science Of Fear (Mistabishi Remix)" >}}
