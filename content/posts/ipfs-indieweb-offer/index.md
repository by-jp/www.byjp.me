---
title: An IndieWeb IPFS offer
summary: "An offer: I'll pin your static IPFS-based IndieWeb site on my homelab! Get in touch :)"
emoji: 💾
date: 2023-08-13T15:44:14+01:00
draft: false
tags:
- ipfs
- indieweb
- hosting
- static-site
- web
- homelab
- admarus
syndications:
---

If you're hosting a static IndieWeb site and (want to) keep an up-to-date copy on [IPFS](https://ipfs.io), I'm offering to pin a replica of your site on my energy efficient [homelab](/tags/homelab) in London (in the UK), making it more highly available, as well as being searchable with [Admarus](https://admarus.net/).

To get started, just [send me an email](mailto:ipfs@byjp.me?subject=Re:%20An%20IndieWeb%20IPFS%20offer&body=Hi%20JP!%0A%0AWould%20you%20be%20interested%20in%20pinning%20my%20IndieWeb%20site%20on%20IPF%3F%20You%20can%20see%20it%20at%E2%80%A6) and I'll get it set up! (If I think your site isn't suitable for this offer—perhaps it violates one of the [Black Laws](https://terra-ignota.fandom.com/wiki/Black_Laws)—we can chat about it.)

If you'd like help in publishing your static blog to IPFS please reach out! I'm passionate about this tech (this blog is automatically co-hosted on IPFS, check it out [on dweb.link](https://www-byjp-me.ipns.dweb.link/)!), I'd love to help you out with understanding and automating your process where I can.

---

I run a power-efficient homelab at my home in London (a [Turing Pi 2](https://turingpi.com/product/turing-pi-2/) rig I should blog about soon), so I have a terrabyte or two of SSD storage with a roughly gigabit upload link, that runs at under 10W of total power consumption (at average load). Perfect for serving relatively low-volume static sites!

I would love to run it off-grid, like the _super_ cool [Low←Tech Magazine](https://solar.lowtechmagazine.com/about/the-solar-website/), but neither London's weather nor the orientation of my flat make that possible today — but maybe one day!

I'll be assuming you have your IPFS-based site set up with [DNSLink](https://docs.ipfs.tech/concepts/dnslink/) — ie. you have a `TXT` record set up for `_dnslink.your.domain.here` with the contents `dnslink=/ipfs/<CID to your site>`. I'll be polling for changes on that DNS record at roughly your DNS record's TTL, and swapping the pin out on my homelab if I see it change.

Initially I'll have my IPFS server configured so that you won't be able to use this as _primary_ hosting for your site, but I'll consider it if things go well. (That could include a webhook to request an update of your pin.)
