---
title: Facebook API Niggle
date: 2009-07-29T00:49:00+01:00
draft: false
emoji: ⌨️
images:
tags:
  - facebook
  - xfbml
  - code
  - from-tumblr
---

I like Facebook (despite its recent frustrating changes that make walls confusing) but I have a niggle as a developer using their APIs. If you choose to use Facebook's XFBML to put a commenting system on your blog (like I do here) then anyone posting that page inside facebook has a separate comments list. Try posting this page on facebook[^1], if you comment on that post in your feed the comments won't appear on this page and things you write here won't appear there.

Granted, sometimes this behaviour is preferable (popular sites might have thousands of comments on one page, and 99% will probably not be your friends or anyone you care about) but it would be good if when you initially 'share' a link on facebook, while the javascript is grabbing the pictures and headlines from the page, it would also look for XFBML and tie the commenting systems together if the site publisher had the relevant option toggled in their [developer settings](http://www.facebook.com/developers/). You could even have non-friends' posts hidden by default (or listed merely as '34 other posts').

I realise its hardly a pressing problem, but I really love Facebook's interconnectivity and I use off-site Facebook commenting in a couple of places[^2], it'd be nice to see it happen!

[^1]: Import note: This won't work any more, as this site no longer uses Facebook's XFBML comments. (I'm not even sure if that feature still exists!)
[^2]: Import note: There were three links to other blogs/websites of mine, all of which now exist within this one.
