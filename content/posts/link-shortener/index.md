---
title: "A linkrot-free URL shortener"
emoji: 🔗
date: 2023-07-08T11:08:45+01:00
draft: false
tags:
  - tech
  - web
  - ipfs
  - shortlinks
  - val.town
  - indieweb
---

Link shorteners are superbly useful, but _really bad_ for the longevity of the world wide web. Every time you use tinyurl.com, bit.ly, or similar you are taking something open and direct (the web address of the think you're pointing people to) and hiding it within the black box that is the company behind that site.

What happens if tinyurl.com closes its doors? Or if bit.ly is bought by a billionaire who charges people to see the long link behind the short one? This is the way the web "rots"; by the removal of the content that holds it together.

However short links are _incredibly_ useful; humans can't remember the long sequences of letters and numbers that some web addresses need to be, so what's to be done?

_Ugh. JP. **I get it**. Let me [skip to the cool part](#how-it-works)._

## What is IPFS?

A few days ago the InterPlanetary FileSystem (IPFS) released a new version of their flagship server with an exciting new feature — the ability to perform redirects (geekier folks can [read about it here](https://specs.ipfs.tech/ipips/ipip-0002/)). To explain why this is so valuable, I'll quickly talk about IPFS, and why it's already solving many similar problems.

[IPFS](https://ipfs.tech/) is a way of storing & requesting data on the internet by referring to it based on _its content_, rather than _its location_.

A web address today looks like `www.byjp.me/posts`. This contains the name of the server you should contact to get my site (`www.byjp.me`), and where on that server the page is that lists all my posts (`/posts`) — it's all _location_ based. If my server dies, then everyone has forever lost access to what's on my site — this is what's known as "[link rot](https://en.wikipedia.org/wiki/Link_rot)".

In the world of IPFS data is accessed, instead, by a Content ID ("CID") — a name that is derrived from, and unique to, the data it points to, rather than where it's located. The way this works is fairly technical, but the same data _always_ produces the same CID, and the CID unquely refers to that data[^1].

My site (as I write this) has the CID of `/ipfs/QmWSnMtmfh78EdKjztRMRLFufk4BMKN3nSvqZBR5g8aDWF`. You can see exactly it at [ipfs.io](https://ipfs.io/ipfs/QmWSnMtmfh78EdKjztRMRLFufk4BMKN3nSvqZBR5g8aDWF), at [Cloudflare](https://www.cloudflare-ipfs.com/ipfs/QmWSnMtmfh78EdKjztRMRLFufk4BMKN3nSvqZBR5g8aDWF), [dweb.link](https://dweb.link/ipfs/QmWSnMtmfh78EdKjztRMRLFufk4BMKN3nSvqZBR5g8aDWF) or any other IPFS server.

This means that if Jo Bloggs makes a copy of my site ("pins" that CID), and then my site goes down forever, everyone can still view my site, _exactly_ as it was meant to be, by getting it from Jo Bloggs instead of me. I'm glossing over some details here but, broadly, so long as someone cares about the data in question: link rot defeated.

[^1]: It's not _technically_ unique, there are just more different combinations of these names (Content IDs, or "CIDs") than there are atoms in the universe, so the likelihood of two being the same for different data is _vanishingly_ small.

## How it works

So, you may be asking, how does this help with URL shorteners? This new 'redirects' feature of the IPFS server (called [kubo](https://github.com/ipfs/kubo#readme)) means information on how to redirect people (from a shortlink to a long URL) can be stored in a file on IPFS and used to do redirections.

My domain `byjp.fyi` is configured to serve content from IPFS, in particular a `_redirects` file ([see it live here](https://byjp.fyi/_redirects)) that holds all of my shortlinks and their destinations. Anyone who pins `/ipns/byjp.fyi` on their IPFS server will _always_ be able to know what my shortlinks were pointing to, even if my site disappears.

For convenience's sake, I keep my `_redirects` file and other files associated to byjp.fyi [on github](https://github.com/by-jp/byjp.fyi), with a [deploy workflow](https://github.com/by-jp/byjp.fyi/blob/main/.github/workflows/deploy.yaml) which pins all the files on [Filebase](https://filebase.com/), and updates Cloudflare with the new DNS entry every time I publish a new shortlink.

I also use [val.town](https://val.town) (an _incredibly_ cool Javascript functions in-the-cloud service) to build an API for adding new shortlinks. Automate all the things! If you had one of my Github access tokens, you could add shortlinks to my domain by heading to [this URL](https://byjp-addshortlink.express.val.run/example?to=https://example.com). You can read my val & see how it works at [byjp.fyi/new](https://byjp.fyi/new).

## How do I do this too?

If you too want to run your own (free!) URL shortener that is resilient to link rot then follow the steps below — then [get in touch](/standing-invitation), I'll happily pin your shortened domain locally, so we've got some more protection against link rot.

These instructions assume you're confident with Github, Cloudflare & code, but you shoud be able to make it through even if you're not, and you can always reach out to me if you have questions!

1. Set up the prerequisites:
   - You'll need a domain you want to host your shortlinks on. I tend to buy mine on Cloudflare (hit "Register domains" in the dashboard)
   - You'll need to be using Cloudflare as your Nameserver for this domain to follow these instructions exactly (though other DNS providers that have APIs to update DNS records should work)
   - You'll need an account with [Filebase](https://filebase.com), or any other API-capable IPFS pinning service.
   - To follow my code exactly you'll need a bucket in Filebase called "microsites".
   - A Github account
2. Fork my [byjp.fyi](https://github.com/by-jp/byjp.fyi) repo, and rename your fork to be whatever domain you registered.
   - I'll refer to this as github.com/yourghname/your.tld from here on. You _can_ calle the repo something different to what your domain is, but you'll need to edit the deploy scripts.
3. Add some secrets to your repo (see your repo's Settings > Secrets and variables > Actions):
   - `CLOUDFLARE_ZONE_ID` needs to be the Zone ID for the domain you're using (find it on the side pane of the Overview page for your domain in Cloudflare)
   - `CLOUDFLARE_TOKEN` needs to be a [Cloudflare token](https://dash.cloudflare.com/profile/api-tokens) with `Zone.Web3 Hostnames:Edit` permissions for the Zone (domain) you're using.
   - `FILEBASE_KEY` needs to be the "Key" from your [Filebase access keys page](https://console.filebase.com/keys).
   - `FILEBASE_SECRET` needs to be the "Secret" from the same place.
4. Create a Web3 Gateway for your domain in Cloudflare:
   - There's a tab called "Web3" on the left when you're viewing your domain, click "Create Gateway"
   - The "Hostname" should be the same as your repo name, and the domain you've bought
   - "Gateway Type" should be "IPFS DNSLink"
   - You can leave the DNSLink field empty (we'll be replacing it via the API shortly)
5. Nearly there! In Github, manually edit the `public/_redirects` file in your new repo
   - Perhaps change the `/me` link to point to your own homepage
   - You probably want to remove all my other shortlinks, but keep the bottom one (starting `/*`), as it manages shortcodes that don't exist.
   - Commit your changes
6. A few minutes after you committed you should see a green tick in Github, and your shortlinks site should be active! Visit your domain at a random path to see the 404 page, and at any shortlink you added to see it work.

If you'd also like an API to add a new shortlink, you'll need a [val.town](https://val.town) account:

1. Go to my val and fork it: [byjp.fyi/new](https://byjp.fyi/new)
2. Create another val with `export const shortlink_domain = 'your.tld'` in it, referring to your domain.
3. And another again with `export const shortlink_repo = 'yourghname/your.tld'` in it, referring to your github repo.
4. You can now send a query like below to have a new shortlink added to your repo!

    ```sh
    curl "https://yourvtname-addshortlink.express.val.run/example?to=http://example.com" \
         -H 'Authorization: Bearer abc123
    ```

   - The host refers to your _val.town_ username, so here `yourvtname-addshortlink.express.val.run` would only be accurate if your val.town account name was `yourvtname`.
   - The path refers to the shortlink you want to make, so here `/example` means you'd be hitting `your.tld/example` when complete.
   - The `to` query parameter is the destination, here that's `http://example.com`. It'll fail if you try and redirect to something on `your.tld` 😉
   - The Bearer token (`abc123` above) needs to be a [fine grained Github personal access token](https://github.com/settings/tokens?type=beta) (or the organisation equivalent) with at least "Read and Write access to code" for the `yourghname/your.tld` repo.

## Caveats & thoughts

- I'd like to build a [Raycast](raycast.com) extension for making new shortlinks. It's stopped raining now, so maybe later!
- If you try to post new shortlinks with the val.town API faster than the deploy time for the repo then your API call will be rejected, and you'll have to try again. I'd like to get some auto-retries in there.
- With val town and Raycast being drivable by Typescript I'd like to release this as a package, so it's mega easy to use. Perhaps it'll even have tests 😲
- I'd _really_ like to be able to publish this blog with "other things I'd like to be included in the root hash", so I could say "resolve `/ipns/byjp.fyi/` and include that IPFS data here", so this blog always includes a 'backup' of the shortlinks. This feels like it's the tip of a bigger feature IPFS might include around deeper archival support.
- Removing shortlinks is still a manual edit of the Github repo away. I don't see much value in allowing that to be automated any time soon.
- I do need to trust val.town with write access to my shortlinks repo — if something leaked out of that part of the flow then back folks could hijack links folks may trust are mine isn't great. I should probably check the shortlinks every now and again.
