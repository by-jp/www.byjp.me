---
title: Fonts
emoji: 🔡
summary: Instructions for how to turn streams of data into legible writing.
draft: false
tags:
- ttf
- zwj
- emoji
- fonts
---

## TTF Fonts

I've a niggling mini-project to write code to insert new emoji into my mac's system emoji font file. For no good reason than to say that I can, I'm wondering whether I can (for example) show `🍝` + the [zero-width-joiner](#zero-width-joiner) + `🚲` as a newly inserted glyph of [Deliveroo's](https://deliveroo.co.uk) logo. Naturally this wouldn't be useful to anyone except myself (as no-one else would have this glyph), but I'm using it as a way to learn how TrueType Fonts are put together.

I've found the docs on the [OpenType spec](https://learn.microsoft.com/en-us/typography/opentype/spec/) at Microsoft, as well as the [TrueType reference](https://developer.apple.com/fonts/TrueType-Reference-Manual/) at Apple, but there is _a lot_. I think I need to start with making a trivial font with a tool like [Glyphs](https://glyphsapp.com) (which I used to make [Caspian](/posts/the-beauty-of-type)) or [FontForge](https://fontforge.org/en-US/), and seeing what's created. There is _definitely_ bytecode in the Apple Emoji TTF file…

## Zero Width Joiner

The ZWJ is a non-visible unicode character that can be used to link codepoints together into single glyphs or more accurately, as [wikipedia](https://en.wikipedia.org/wiki/Zero-width_joiner) tells me, indicates that the following grapheme is positioned relative to the previous one.

The classic example for me is the various family emoji, which look like this: 👨‍👩‍👦‍👦

That particular one is `👨` + ZWJ + `👩‍🦰` + ZWJ + `👦` + ZWJ + `👦` — so a full 7 multi-byte characters for a single glyph! Then of course you can add skin colour indicators for each of the members of the family, and you can rapidly get to _enormous_ single "characters".

I saw a post recently that called out the Welsh flag emoji 🏴󠁧󠁢󠁷󠁬󠁳󠁿 as the longest flag, being `🏴` + the letters `gbwls` + `␘`. Designing emoji fonts must be a multi-year effort at this point!
