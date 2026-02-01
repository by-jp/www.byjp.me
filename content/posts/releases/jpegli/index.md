---
title: "Easy image minification"
emoji: 🌇
date: 2024-10-11T10:24:33+01:00
summary: I wrote the `jpegli` CLI tool to quickly reduce the pixel- and byte-size of images for my blog, using JPEGli compression.
tags:
- JPEGli
- CLI
- code
atUri: "at://did:plc:ephkzpinhaqcabtkugtbzrwu/site.standard.document/3mdryqll5zg2o"
---
I've not written anything substantial here in _some time_, and sadly that's not going to change today!

I do have a mini-announcement about a tool I built for image posting here though — I've released the [`jpegli`](https://github.com/jphastings/jpegli) command line tool which:

- Ensures the image is no larger than 2048×1920px (maintaining aspect ratio)
- Compresses the image with the impressive [JPEGli](https://opensource.googleblog.com/2024/04/introducing-jpegli-new-jpeg-coding-library.html) algorithm

JPEGli is a (relatively) new way of producing JPEG images but at significantly smaller file sizes (with no real noticeable difference). The end up _roughly_ the size of WebP images, but (being JPEG images) they're readily understood _everywhere_.

If this sounds useful then you can download for mac/linux/windows [from github](https://github.com/jphastings/jpegli/releases) or you can do any of the following

```bash
# With Homebrew installed
brew install jphastings/tools/jpegli

# With Go installed
go install github.com/jphastings/jpegli@latest
```

Do [let me know](/standing-invitation) if you use it, or would like to see some improvements/changes to how it works. It's very basic at the moment!
