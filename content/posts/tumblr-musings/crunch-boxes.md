---
title: Crunch boxes
date: 2010-10-20T23:18:00+01:00
draft: false
emoji: 💻
tags:
  - idea
  - computing
  - from-tumblr
atUri: "at://did:plc:ephkzpinhaqcabtkugtbzrwu/site.standard.document/3mdryqtrcrm2s"
---
After a quick support request from the Plex team I mentioned how much I'd like a linux version of the Plex Media Server so I could look to getting it working on a [Drobo FS](https://web.archive.org/web/20101012030603/http://www.drobo.com/products/drobo-fs.php).

As the anonymous Plexian noted, the major issue with this would be the lack of power in Drobo's CPU for transcoding — you just wouldn't be able to live-stream videos to mobile devices from PMS running on embedded systems.

With a bit of tinkering you can farm out data intensive tasks like this to more powerful processors, I suggested leveraging the power from a PMS on a traditional computer on the network, but it got me to thinking:

Wouldn't it be cool if you could buy or create a totally headless computer which you could hook up to power and the network and would provide number crunching power to any application which needed it?

My first thoughts would be to create/use a specific distro of linux. You'd create daemons for a few common place CPU intensive tasks (video transcoding for example) and, as the OS would be standard, any specific daemons could be pushed to the 'crunch box' to perform any other task.

The major benefit of these, as far as I can see, is being able to have the grunt force for processing sleeping on your network, using minimal power, until it's required. I'm pretty conscious of how much power my iMac would use if I left it on just to funnel video to my iPhone via Plex while I'm travelling.

So, for anyone reading this, do you think such a system could be created? Is there anything you would want included?
