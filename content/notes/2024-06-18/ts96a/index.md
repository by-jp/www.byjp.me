---
date: 2024-06-18T06:49:39.132+01:00
publishDate: 2024-06-18T06:49:39.132+01:00
inReplyTo: https://www.macchaffee.com/blog/2024/ddos-attacks/
references:
  https://wwwMacchaffeeCom/blog/2024/ddosAttacks/:
    url: https://www.macchaffee.com/blog/2024/ddos-attacks/
    type: entry
    name: DDoS attacks can threaten the independent Internet
    summary: Mac's Tech Blog
    featured: https://www.macchaffee.com/static/favicon.png
slug: ts96a
tags:
- IPFS
- web
---
It's not a complete solution by any means (it would only work for static sites) but I'm leaning hard into the use of the peer-to-peer [IPFS](/tags/ipfs) to take the teeth out of DDoS attacks.

Though today most IPFS sites are served from IPFS↔HTTP gateways (which will fall to a sufficient DDoS attack, like any other server), using the p2p IPFS protocol itself would mean an attacker could only ever take out specific peers. Though that's still unpleasant, the aim (taking down a given site) is unachievable, so it disincentivises site-specific attacks.

For example; you can read this comment over HTTP at https://www.byjp.me (protected by a mega-corp at the moment), or you could use IPFS and read it at /ipns/www.byjp.me (protected by any- and everyone using IPFS who's ever been to any part of my site).
