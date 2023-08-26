---
title: It's been a while…
date: 2009-11-15T23:21:00+00:00
draft: false
emoji: 📆
images:
tags:
  - irotoku
  - code
  - ruby
  - github
  - poetry
  - iTunes
  - DAAP
  - Songbird
  - gem
  - from-tumblr
---

Well, its been a few months since I posted last - the genuine joy of interesting/difficult work in my final year as a Physics undergrad - but there's a fair amount I should probably put up here for posterity:

- I've become very interested in [Playdar](https://playdar.org) - an old friend of mine, [James Wheare](http://james.wheare.org/), is a ket developer - and the possibilities music resolution has to offer. As a consequence I released quite a few libraries and a program that helps some of playdar's features become a little more tangible.

  Voici [DAAPlaydar](https://github.com/jphastings/DAAPlaydar): a script that will resolve playlists of yours (using playdar) and make them available to you via DAAP share. A very cool concept (even if I say so myself!) but it falls short because iTunes is very picky about the DAAP servers it talks to, Songbird works well though :) I hope to be ablet o work a bit more on this in the near future - if you have any expertise with DAAP, you know who to get in touch with!

- In order to get DAAPlaydar working I had to build a few libraries, one to decode Apple's DMAP object encoding method (now a ruby gem on gemcutter called (drumroll please) [dmap](https://github.com/jphastings/dmap) - woop! There's also [PlaydARR](https://github.com/jphastings/PlaydARR), a ruby library for interacting with the Playdar server (also now a gem on gemcutter). So named because [pirates help combat global warming](https://en.wikipedia.org/wiki/Flying_Spaghetti_Monster#Pirates_and_global_warming). True story.
- My dashing alter ego [Cy Densham](https://www.byjp.me/poetry)[^1] has released some more poetry, your opinions are more than welcome.
- At some point some of the many photos I took while traveling Japan will arrive on Flickr (I'm fighting with FlickrUpload for Aperture at the moment) some of[^2] [the photos are already up](https://www.flickr.com/photos/jphastings/albums/72157622516311575), and of course they're geotagged ~so you can just [browse to Japan in Flickr](https://www.flickr.com/photos/jphastings/map?&fLat=35.5322&fLon=136.2963&zl=12&order_by=recent) to see them~[^3].
- I've also realized that one of the projects I enjoyed working on the most isn't mentioned on here! Shock horror. [irotoku](https://github.com/jphastings/irotoku) is a (very basic) way of hiding information in images - ie. Stenography - the implementation of the decoder I wrote in C here is quick enough that if you hide MP3 data in a (big!) image you can pipe directly from the image through irotoku to an MP3 player and listen to your heart's content. Good fun.
- So that's pretty much all for now, Cy has some interesting ideas so there may some more stuff up here soon, but he might be pushed down by this newly inspired Physicist and such tedious things as job applications. Yay.

[^1]: Import note: This was once a link to "Cy Densham"'s Twitter account, but here in 2023 Twitter is X and rapidly turning into a dungpile, so it's pointing back to poetry on this blog now.
[^2]: All of
[^3]: Apparently this no longer works?
