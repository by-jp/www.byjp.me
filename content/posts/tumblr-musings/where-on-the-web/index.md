---
title: Where on the Web
date: 2009-04-21T02:55:00+01:00
draft: false
emoji: 🌍
tags:
  - code
  - geo
  - gist
  - github
  - map
  - processing
  - ruby
  - whereontheweb
  - from-tumblr
---
{{< figure src="where-on-the-web.jpg" alt="A terrain map of the earth over the Atlantic, showing red dots moving between the UK and other places on earth" >}}

Where On The Web is aliiiiive!

I built it, and it appears to work averagely well. You can see the code in [the gist on github](https://gist.github.com/jphastings/98878) and soon there will be a mini-video online showing you how it all works. In the meanwhile, just enjoy the pretty colours :D

I’ve uploaded a snapshot[^1] of the whole window for you to ogle at too.

If you want to play with the code, get ruby installed and then just do:

```sh
gem install ruby-processing
```

or follow [the advice here](https://github.com/jashkenas/ruby-processing/wiki/getting-started).

Then download the code from [the gist](https://gist.github.com/98878), find an equirectangular map of the earth, call it `map.jpg` and then run:

```sh
rp5 run whereOnTheWeb.rb
```

(There are more details in the gist) - Enjoy!

[^1]: Import note: Sadly this image has been lost.
