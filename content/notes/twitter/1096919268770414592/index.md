---
date: "2019-02-16T23:48:28Z"
tags:
- imported
- from-twitter
- golang
inReplyTo: /notes/twitter/1096919221345357824
---
I decided to generate the gif in [golang](/tags/golang) too. Palette quantisation was fun to learn about; how you take a 32 bit colour image and decide which colours you should use when downsampling into 256 colours.

I used the "Median Cut" method \([https://github.com/ericpauley/go-quantize/quantize](https://github.com/ericpauley/go-quantize/quantize)) {{< imgorvid src="media-1.mp4" >}}
