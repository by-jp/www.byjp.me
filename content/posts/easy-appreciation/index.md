---
title: Easy appreciation
emoji: 👏
date: 2024-04-06T23:48:55+01:00
summary: I've built a "clap" feature into my blog so you can show appreciation anonymously and easily, if you want.
tags:
- community
- ValTown
- ProgressiveEnhancement
- IndieWeb
syndications:
- https://hachyderm.io/@byjp/112229385641236069
- https://bsky.app/profile/byjp.me/post/3kpjvofkuah2f
---

Here in my little career break I'm spending a lot of time thinking about _community_. This blog hasn't had much of a need for community (there's such an ecclectic mix of stuff here that people stumble upon it rather than frequent it) but as I've been building in tools like [webmentions](https://indieweb.org/Webmention), and pulling comments from other sites (like [on this post](/posts/chef-gpt/#interactions)), I've noticed the absence of an easy "I appreciate this" mechanism for passers-through.

Yesterday I built a clap button, the one you can see below this paragraph, and on every page at the bottom of the post (though sometimes its a heart instead of a clap, for aesthetic reasons). If you click it it lets me know you appreciate my post, and keeps track of how many times that's happened (you can press it more than once if you really like something!)

<figure style="text-align:center">{{< claps >}}<br/><small style="font-style: italic">Press this to show appreciation!</small></figure>

I decided to make this button totally anonymous; this makes it much easier for you to use, but the lack of tracking also means that—for me—it's a feel good vibe rather than data I can get analytical over. If you like the sound of it, you can try it out 😜

## Going a little deeper

For those more technically inclined, there are some minor smarts behind the scenes here. My site tries to be as decentralised as possible, but recording & retrieving claps needed something centralised (distributed systems are complex!) so I ended up using [val.town](https://val.town). These folks offer an _awesome_ platform for simple lambda-like web functions; I _love_ that they can be totally public & transparent — you can see (and copy!) the one that [drives my claps here](https://www.val.town/v/byjp/claps).

That 'val' accepts `POST` requests (incrementing the number of claps), `GET` requests (telling you how many claps a post has), and the special `GET /` which lists the claps for every post. That last endpoint is useful for the code behind my static site builder ([hugo](https://gohugo.io)), it [pulls that data](https://github.com/by-jp/www.byjp.me/blob/38751361ff6b8730428d8227f98189312576a709/layouts/partials/claps.html#L8-L17) as the site is being built (at least once a day) so the clap counts are accurate reasonably quickly without needing any javascript.

Speaking of not needing javascript, the button is also a `<form>` and the 'val' also handles `text/html` requests, so you can use this button even if you have javascript turned off. ([Progressive enhancement](https://developer.mozilla.org/en-US/docs/Glossary/Progressive_Enhancement) for the win!)

…but of course javascript gives me more flexibility. I can record your clap without you leaving the page, and give the button an accented colour to give you some feedback. It also means that, when you click that button, I can recording your appreciation in your browser's local database. This means it'll remain accented for you ("you already showed your appreciation for this at least once") when you _next_ visit the same page, but while maintaining your anonymity and control. (If you ever 'clear browsing data' for this site, they'll stop being accented, but the clap still remains, anonymously, in my 'val').

If you find this interesting enough to want to reproduce it on your own blog, please [reach out](/standing-invitation)! I can definitely make it easier to use on other sites 😅

## What's appreciated?

I doubt this little tool will get much use — my site is a lovely wild garden for me, rather than a valuable resource for others — but to end with a little fun, here's a list of the six most appreciated posts across my site right now.

{{< topclaps 6 >}}
