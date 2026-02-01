---
title: Twitter links in Tumblr
date: 2009-05-20T21:43:00+01:00
draft: false
emoji: 🐤
images:
tags:
  - javascript
  - code
  - tumblr
  - twitter
  - from-tumblr
atUri: "at://did:plc:ephkzpinhaqcabtkugtbzrwu/site.standard.document/3mdryr2m6mh2b"
---
When writing [a post](/posts/how-developers-brains-work) the other day[^1], I wanted to be able to use the tumblr chat where the participants were twitter users. Tumblr doesn’t support HTML in chat usernames, so I wrote a tiny javascript chunk to automatically troll any elements you give it and replace `@twitter-username` with a link to their account. (You can see it in action on the aforementioned post[^2]) You can find the code in [this github gist](https://gist.github.com/jphastings/115054), all you need to do is link in [mootools](https://web.archive.org/web/20230812014508/https://mootools.net/) (you only need to have the selectors part available, for now) and this script in the head of your template like this:

```html
<script type="text/javascript" src="mootools.js"></script>
<script type="text/javascript" src="tumblrtweets.js"></script>
```

Then add a small javascript section at the end of your template (inside the **body**):

```html
<script type="text/javascipt"> tumblrTweets("span.label"); </script>
```

Where `span.label` is the [CSS Selector](https://css.maxdesign.com.au/selectutorial/) for the elements you wish to scan. You can link to these files from my web server if you like[^3] (check the source of any of these pages), but if you have your own hosting _please_ use that instead, I only have limited bandwidth!

[^1]: Import note: this post was originally written on Tumblr, which feels like a pretty crucial piece of information for this post to make sense!
[^2]: Sadly you can't! See the footnote above.
[^3]: I doubt that you'd still want to, but that server doesn't exist any more. You can always load directly from the [raw gist](https://gist.githubusercontent.com/jphastings/115054/raw/3a20eb86caf64ae5101d0447bed4ed04f2da0e57/tumblrtweet.js) on Github.
