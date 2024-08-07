---
date: 2024-08-07T17:13:37.126+01:00
publishDate: 2024-08-07T17:13:37.126+01:00
slug: xwrgg
---

I've been toying with an idea for a partner tech to the #GeminiProtocol; a standard for dynamic pages you can derive locally (instead of needing a server).

A "page" that is a #WASM blob that accepts the Query part of the request to it, and returns a Gemini response (with status).

I really like the idea of being able to keep a static copy of a capsule, even if it has dynamic features, like search.

I'd store my capsule (including this new offline search functionality) on #IPFS, record the root CID in an [IPLD-encoded IPNS record](https://github.com/multiformats/multicodec/pull/312) signed with my capsule's TLS cert, bundle all that in a [CAR file](https://ipld.io/specs/transport/car/). A permanent, interactive, trustable, local-first copy of my (and anyone else who's keen) online space.
