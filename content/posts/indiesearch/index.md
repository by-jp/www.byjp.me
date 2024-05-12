---
title: "IndieSearch"
emoji: 🔍
date: 2024-05-12T21:27:36+01:00
summary: I built a prototype for client-side, fully distributed search for the IndieWeb — check out the demo & a little explainer.
tags:
- IndieWeb
- search
- decentralised
- demo
- IPFS
---

I've been exploring what totally decentralised and local-first compatible search might look like for the web. Here's a quick video demo of a prototype I've built called _IndieSearch_, powered by the (awesome) client-side search tool called [Pagefind][pagefind] (or read on, if videos aren't your thing).

{{< vimeo 945570914 >}}

Firstly, I should be clear that I designed this based on the premise that you want to search the sites you've _previously visited_. It's not intended to help with discovering new sites directly (but I do have some plans for how this might change in the future).

IndieSearch works as a browser extension (in my video I'm using [Arc](/posts/arc-new-browser/), a Chromium-based browser). That extension provides the search homepage, but its most important job is to check for a specific HTML `<link>` tag on the pages I visit & store what it finds. Specifically something like this:

```html
<link rel="search" type="application/pagefind" href="/search" title="byJP">
```

The `href` attribute is the location of the site's [Pagefind][pagefind] index, and the `title` is what we'll show to the person searching when they're managing their sites. There they can sort through supported sites they've visited, seeing the new ones, and flagging them to be included or excluded in future searches.

{{< figure src="config-popup.webp" alt="The IndieSearch config page, showing a newly visited site, three sites to be included in search results, and one excluded" height="600" >}}

Any time they visit the IndieSearch homepage (a page served from their browser extension, but linked to from _indiesearch.club_, for ease of access — though I've not built that yet!) they can now search all the sites they've visited and/or included (assuming they support IndieSearch).

{{< figure src="indiesearch.webp" alt="Using IndieSearch to query this blog; looking for 'Appreciate' and finding my recent post on Easy appreciation." >}}

## What's good

- The search is blazing fast as Pagefind method breaks up the index so clients only need to download tiny parts of the index, then the fragments of search hits. Usually this is ~30kB per search, per site.
- The search index is entirely static! No need for a special server; just some odd-looking files in a directory of your site.
- IndieSearch doesn't have many moving parts (awesome!), and its _very_ simple to [add Pagefind support to a site](https://pagefind.app/docs/), so adoption has fewer obstacles.
- [PageFindUI](https://pagefind.app/docs/ui-usage/)'s is powerful, pre-built, and have loads of great features that work out-of-the-box (like search facets).
- Index configuration is almost entirely done by the indexer (ie. the site owner), with presentation configurable by the UI (within IndieSearch's code) — this is entirely down to [CloudCannon](https://cloudcannon.com/)'s awesome work with Pagefind — have a raved enough about it yet?

## What's difficult

- Scalability is rough; an extra 30kB per site, per search could make searching all your most frequented sites might end up being quite bulky. Each query Pagefind sends uses suitable cache headers though, so this could be limited, particularly for sites that change infrequently.
- IndieSearch requires that everyone uses the exact same search index structure & platform. CloudCannon have done something awesome here, but there's not much by way of standards or industry adoption behind the format they use — I'd like to see more permanence in the index format, and Pagefind's contributor base.
- I don't know how this could work on mobile, or on browsers that don't have extension capabilities in the same way.

## What's next?

Firstly, if you like where I'm going with this — give me a clap, for motivation! 😄

{{< claps >}}

I know I want to get IndieSearch into the various extension stores so it's easy to install, even in this prototype stage. (The [code is on Github](https://github.com/jphastings/indiesearch), by the way, if you're feeling brave and want to try it yourself.) This is also my first foray into building browser extensions, so I've a lot to learn (especially how to build them for Firefox, Safari, and Chromium all at once).

I'd also be interested in adding some IndieWeb-themed features, like being able to sync your indexed sites with a blogroll, or seeing if I can convince the Pagefind team to include [microformats](https://microformats.org) context to their indexer.

Oh! And if you use Pagefind on your site already, add my provisional `<link>` tag to your site & let me know! It'd be great to have some other test sites out there.

It's tangentially related, but I'm still desperate to see the [IPFS](/tags/ipfs) folks allow DNSLink records travel offline, and be associated with their SSL certs. The `TXT` record at `_dnslink.www.byjp.me` that allows IPFS enabled browsers to see `/ipns/www.byjp.me` and resolve whichever `/ipfs/Qm…` CID is the current version of my site _also_ contains references to my entire Pagefind search index. If you've pinned my site on IPFS you've _also_ cached a local copy of my search index. If pinning `/ipns/www.byjp.me` also pinned an assertion that "at {timestamp} the root CID of www.byjp.me was `Qm…` as signed by that domain's SSL certificate" then IndieSearch could provide full, local-first search of IndieWeb sites on a local machine, or local network, totally detached from the internet.

I can dream 😁

Let me know your thoughts too — my Webmentions, [Mastodon](https://hachyderm.io/@byjp), [Bluesky](https://bsky.app/profile/byjp.me), and [email](/standing-invitation) are always open 😊

[pagefind]: https://pagefind.app "The Pagefind homepage & developer docs"
