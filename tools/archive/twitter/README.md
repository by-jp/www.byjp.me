# Twitter archive

This tool will take a Twitter archive (I'd link to how to get there, but I don't have an account any more, so can't!) and turn it into a series of timestamped notes for your Hugo blog.

```sh
go run . <path/to/archive.zip> <path/to/hugo/root/>
```

It will create one folder per post in the `content/notes/` directory. You can change how those posts look by creating a `layouts/notes/single.html` template.

## Hacks

This is super hacky code, please read through it before you run it. In particular, I've:

- hardcoded my own twitter ID (12721) to detect replies to myself
- made use of my own `{{< imgorvid >}}` shortcode to handle files that could be videos or images, rather than embed them 'properly'
- apparently Twitter (or one of the services I used to post) swapped the order of geographic coordinates from `[lon, lat]` to `[lat, lon]` at some point, with no way to determine the difference, so… ¯\_(ツ)_/¯
- Lots of short URLs and images went missing; I left the old URLs in there, in case they arise somewhere (very unliekly)
- Probably lots of other stuff.
