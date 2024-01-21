---
title: "Bookmarks With Omnivore"
emoji: 🔖
date: 2024-01-21T09:41:59Z
summary: I created an importer that brings articles I've read & enjoyed in Omnivore into my blog as IndieWeb bookmarks!
tags:
- omnivore
- bookmarks
- IndieWeb
- Micropub
- go
- code
- CosyWeb
syndications:
- https://hachyderm.io/@byjp/111793793066838551
---

Yesterday, recovering from a nasty (thankfully non-Covid) cold, I decided to build an import tool for pulling articles I've read & enjoyed in [Omnivore](https://omnivore.app) into this blog as [bookmarks](/booksmarks) (in the [IndieWeb style](https://indieweb.org/bookmark)).

Check out an example here with my bookmark of [The Tyrany of Obviousness](/bookmarks/the-tyrany-of-obviousness) by P.E. Moskowitz.

I promise I won't gush about Omnivore _too_ much, but there's a good reason I put the effort into building a custom integration here. Omnivore is brilliant. It's been the first read-it-later app I've found simple & powerful enough to actually _read_ from. The text-to-speech systems they have are _outstanding_ (use "Fable" if you like an English/SSB accent!), being able to subscribe to RSS feeds, email newsletters, and pull (even some paywalled) browser articles in make it useful everywhere, and their dedication to open source and commitment to [sustainable growth](https://docs.omnivore.app/about/pricing.html) make me love this tool even more.

So my import tool had a lot to live up to! I decided to start with a proof-of-concept [Go](/tags/go) importer (it's available [in this blog's repo](https://github.com/by-jp/www.byjp.me/tree/main/tools/import/omnivore) under an MIT licence). It works surprisingly well! It seeks out articles that have been completely read and have at least 1 annotation (as a signal for "I like it enough to share it"), then pulls my commentary, highlights, and the title/author/summary into the markdown frontmatter & content so I can display my bookmark.

## Limitations

I built this as a sick-day hack so I could figure out if it was worthwhile making something more permanent, so it definitely has its limitations!

- I have to manually run a `task import` to pull any new articles into my blog. (I may automate this, but the What's Next section, below, has a better idea)
- It re-processes every article every time it runs. Large lists will likely be slow.
- A limitation of Omnivore's API means that the `readAt` timestamp (which I use as the publish timestamp for my bookmark) gets updated _every_ time you read that article.
  - The Omnivore team have very kindly said they'll consider putting a `firstRead` column in the response for usecases like this! It's been absolutely delightful to have a human response from their team, and even moreso to have something like this considered, even if it never makes it to the top of the backlog. (Thanks Jackson!)
- I haven't yet tested if this will work articles that get pulled into Omnivore by email (another frankly awesome feature) but there's only one way to find out!

## What's next

I'm already enjoying this tool a bunch, so I think I'll develop it further. If others are interested I'll try to make it something easy to run as well (do get in touch on [mastodon](https://hachyderm.io/@byjp) or [send an email](https://www.byjp.me/standing-invitation/) if that sounds interesting!)

My thoughts are to turn this into a service/serverless function that will listen for Omnivore's excellent [webhooks](https://docs.omnivore.app/integrations/webhooks.html), (configurably) filter down to events that should result in "publish a new bookmark", and trigger the [Micropub]([/tags/micropub](https://indieweb.org/Micropub)) endpoint of your blog with all the same highlight/annotation information.

This should mean that you could publish your favourite Omnivore articles to your blog no matter what it looks like or where it lives! Lots of blogging tools [already support Micropub](https://indieweb.org/Micropub/Servers#CMS_Software) ([micro.blog](https://micro.blog) is a fave of mine), and tools like [IndieKit](https://getindiekit.com) (which is great, I use it here) can integrate with almost any statically generated blog too. This feels like a great way to make it available for all!

Long live the [cosy web]([/tags/cosyweb](https://maggieappleton.com/cozy-web)), filled with stories and articles recommended by friends and humans, not corporations and algorithms!
