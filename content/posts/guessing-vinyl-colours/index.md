---
title: "Guessing Vinyl Colours"
emoji: 🎨
date: 2024-05-07T15:13:43+01:00
draft: false
summary: I used my Nix Mini colorimeter & a bit of maths to estimate transparency for some adhesive vinyl I bought.
tags:
- maths
- colour
- vinyl
- SilhouetteCameo
topics:
- Creativity
- Maths
math: true
atUri: "at://did:plc:ephkzpinhaqcabtkugtbzrwu/site.standard.document/3mdryqmempj2z"
---
I've been getting a lot more use from my [Silhouette Cameo 4](https://www.silhouetteamerica.com/featured-product/cameo) plotter/cutter recently, and I've been using a couple of different vinyls to make things. I found it annoying that I couldn't get a good read on what the colour was before I purchased them, so here's the readings I've taken for the vinyls I bought & where I bought them in case any one else finds them useful!

I took readings for their colour on both black and white card to try and estimate transparency — the process I followed is below. You can see the cheaper "IModeur" yellow vinyl is thinner, and transparent enough to make a difference!

| Vinyl colour                                                                                                      | Approx colour                     | On top of {{< swatch "239,236,230" >}} | On top of {{< swatch "64,63,66" >}} |
|-------------------------------------------------------------------------------------------------------------------|-----------------------------------|----------------------------------------|-------------------------------------|
| [Oracal 651 sample](https://www.gmcrafts.co.uk/product/sampler-pack-basic-oracal-651-for-silhouette/): Red        | {{< swatch "185,0,9,0.98" >}}     | {{< swatch "194,0,11" >}}              | {{< swatch "174,0,13" >}}           |
| [Oracal 651 sample](https://www.gmcrafts.co.uk/product/sampler-pack-basic-oracal-651-for-silhouette/): Yellow     | {{< swatch "254,192,0,0.99" >}}   | {{< swatch "255,192,0" >}}             | {{< swatch "251,191,0" >}}          |
| [IModeur Glossy Yellow](https://www.amazon.co.uk/dp/B09JG886SK)                                                   | {{< swatch "219,196,11,0.83" >}}  | {{< swatch "248,212,14" >}}            | {{< swatch "167,165,55" >}}         |
| [Oracal 651 sample](https://www.gmcrafts.co.uk/product/sampler-pack-basic-oracal-651-for-silhouette/): Lime Green | {{< swatch "100,164,43,0.99" >}}  | {{< swatch "101,164,44" >}}            | {{< swatch "100,163,44" >}}         |
| [IModeur Glossy Green](https://www.amazon.co.uk/dp/B09JG7Y1Y4)                                                    | {{< swatch "0,119,47,0.99" >}}    | {{< swatch "0,122,48" >}}              | {{< swatch "0,116,49" >}}           |
| [Oracal 651 sample](https://www.gmcrafts.co.uk/product/sampler-pack-basic-oracal-651-for-silhouette/): Azure Blue | {{< swatch "0,90,173,1" >}}       | {{< swatch "0,90,173" >}}              | {{< swatch "0,90,173" >}}           |
| [IModeur Glossy Blue](https://www.amazon.co.uk/dp/B09JGBNNCT)                                                     | {{< swatch "0,137,206,0.99" >}}   | {{< swatch "0,138,209" >}}             | {{< swatch "0,136,202" >}}          |
| [Oracal 651 sample](https://www.gmcrafts.co.uk/product/sampler-pack-basic-oracal-651-for-silhouette/): Ice Blue   | {{< swatch "59,159,211,1" >}}     | {{< swatch "58,159,212" >}}            | {{< swatch "60,158,209" >}}         |
| [Oracal 651 sample](https://www.gmcrafts.co.uk/product/sampler-pack-basic-oracal-651-for-silhouette/): Lilac      | {{< swatch "187,147,186,1" >}}    | {{< swatch "188,146,186" >}}           | {{< swatch "187,147,186" >}}        |
| [IModeur Glossy White](https://www.amazon.co.uk/dp/B09MHNX66S)                                                    | {{< swatch "238,236,237,0.95" >}} | {{< swatch "239,236,236" >}}           | {{< swatch "228,229,230" >}}        |
| [IModeur Glossy Black](https://www.amazon.co.uk/dp/B09MH8GL3N)                                                    | {{< swatch "5,0,8,1.0" >}}        | {{< swatch "4,0,8" >}}                 | {{< swatch "6,0,7" >}}              |

_I'll keep an up-to-date list of vinyl I buy and their colours in my [memex page on plotting](/memex/art-design/plotting/)._

### Colorimetric process

I measured these by putting two squares of the vinyl on black card and on white card, then using the difference between their measured colours (using my trusty [Nix Mini 2](https://www.nixsensor.com/mini-3-color-sensor/)) to deduce their colour _including_ transparency.

This process isn't perfect, in the real world a material can be differently transparent at different frequencies. My process was:

1. Measure the sRGB values for both the plain white ($R_c^w$, $G_c^w$, $B_c^w$) and black ($R_c^b$, $G_c^b$, $B_c^b$) card I'm using as backing paper (values in the table headers above).
2. Measure the sRGB values for the vinyl being tested on both black and white card (in the right-most two columns above).
3. Calculate the [colorimetric greyscale](https://en.wikipedia.org/wiki/Grayscale#Colorimetric_(perceptual_luminance-preserving)_conversion_to_grayscale) value for both the white and black card ($Y_c^w$, $Y_c^b$), and the vinyl on both black and white card ($Y_v^w$, $Y_v^b$), from their measured RGB values

    $$Y = 0.2126R + 0.7152G + 0.0722B$$

4. Calculate the alpha value for the vinyl from the greyscale values:

    $$A = 1 - \frac{Y_v^w - Y_v^b}{Y_c^w - Y_c^b}$$

5. Use the calculated alpha to find the 'true' red, green, and blue values for the vinyl, averaging the results from the black and white backing card:

   $$R = \frac{R_c^w + R_c^b}{2} + \frac{R_c^w + R_c^b - R_v^w - R_v^b}{2 A}$$

6. Clip the calculated sRGB values to be within [0, 255], for the oddities!
