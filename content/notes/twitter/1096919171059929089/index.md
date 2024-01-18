---
date: "2019-02-16T23:48:04Z"
tags:
- imported
- from-twitter
- golang
in_reply_to: ../1096915446702489600
---
"How did I do it?", you ask?

WELL. I used ffmpeg to extract 5 second moments to PNG frames, then some custom [golang](/tags/golang) to compare the perceptual hashes of every frame and look for identical hashes that are more than a few frames apart. {{< imgorvid src="media-1.mp4" >}}
