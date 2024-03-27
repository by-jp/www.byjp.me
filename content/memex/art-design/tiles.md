---
title: Tiles
emoji: 🔲
summary: Shapes that fit together to create an infinite pattern.
tags:
- monotile
- shapes
- tiling
- colour
references:
  chiral-aperiodic-monotile:
    url: https://cs.uwaterloo.ca/~csk/spectre/
    name: A chiral aperiodic monotile
    rel: +accurate
---

## Aperiodic monotiles

{{< figure src="../spectre.webp" caption="The spectre aperiodic monotile, with \"odd\" tiles shaded." >}}

2023 was a good year for tiling! The Einstein hat and Sprectre aperiodic monotiles were discovered and [documented](https://cs.uwaterloo.ca/~csk/spectre/). These are special because they've been proven to _never repeat_. There's something really quite wonderful about that!

I've designed 3D-printed cookie cutters that could make ceramic tiles with these, but I've not reached out to my local potters to ask for some time to make them — and I haven't figured out the 'rules' for placing a tile next to others, some placings definitely end up making tilings that can't be completed.

## Spectre + Tantrix

I wrote a little code a few weekends ago to try and figure out if I could make a Spectre tiling that also had lines that travelled across the tiles — something like a tiling game I used to play called [Tantrix](https://www.tantrix.com/).

Though there are some lines that close (you can see some here) I have a _hunch_ that there may be provably non-zero many that never close. I have no idea how to go about proving that though!

{{< figure src="../spectre.svg" title="My spectre & tantrix crossover tiling, each one with a #colour chosen for its angle of rotation." alt="A computer-generated tiling of 20 identical 14-sided irregular shapes. They vary in light pastel colours. Each of the 14 sides inside each shape is connected with a different side by a white line, making a big but pretty knot of white lines connected through each of the shapes.">}}

The code for making this is below; but it's definitely not polished!

<details>
<summary>Code for creating Spectre + Tantrix crossover</summary>

```ruby
require 'victor'
include Victor

class Spectre
  attr_reader :points

  # Angle 1 is the angle between side 1 and side 2 etc.
  ANGLES = [90, 60, -90, 60, 0, 60, 90, -60, 90, -60, 90, 60, -90, 60].freeze
  SIDE_LENGTH = 25
  LINE_WIDTH = 2.5
  CONTROL_LEN_MULT = 0.4

  SHOW_NAME = false

  ARCS = [
    [7, 9],
    [3, 11],
    [0, 13],
    [2, 8],
    [5, 10],
    [4, 6],
    [1, 12],
  ]

  def initialize(origin, corner_index, rotation, name: nil, show_corner_indeces: false)
    @show_corner_indeces = show_corner_indeces
    @hue = ANGLES[corner_index..].sum(-rotation) % 360
    @indexes = (0...ANGLES.length).cycle.take(corner_index + ANGLES.length)[corner_index..]
    @name = name || (@@i ||= 1).to_s

    previous_angle = -rotation
    @points = @indexes.each_with_object([[origin[0], origin[1]]]) do |idx, acc|
      previous_angle += ANGLES[idx]
      ang = previous_angle / 180.0 * Math::PI
      acc.push([
        acc.last[0] + SIDE_LENGTH*Math.sin(ang),
        acc.last[1] + SIDE_LENGTH*Math.cos(ang)
      ])
    end

    @background = "hsl(#{@hue}, 100%, 85%)"
    # @background = "hsl(#{(@hue / 90).floor * 90}, 100%, 85%)"
    # @background = "hsl(23, 100%, 72%)"

    @@i += 1
  end

  def point(idx) = @points[@indexes.index(idx)]

  def between(idx1, idx2) = [
    point(idx1)[0] + (point(idx2)[0] - point(idx1)[0])/2,
    point(idx1)[1] + (point(idx2)[1] - point(idx1)[1])/2
  ]

  def normal_point(idx1, idx2, n_dist)
    b = between(idx1, idx2)
    p1 = point(idx1)
    p2 = point(idx2)
    dx = p2[0] - p1[0]
    dy = p2[1] - p1[1]
    b_ang = Math.atan2(dy, dx)
    b_dist = Math.sqrt(dx*dx + dy*dy) / 2

    n_ang = Math.atan2(n_dist, b_dist)

    t_ang = b_ang - n_ang
    t_dist = Math.sqrt(n_dist*n_dist + b_dist*b_dist)
    n_x = Math.cos(t_ang) * t_dist
    n_y = Math.sin(t_ang) * t_dist
    [p1[0] + n_x, p1[1] + n_y]
  end


  def svg
    SVG.new.tap do |svg|
      svg.polygon points: @points, stroke: 'none', fill: @background
      
      ARCS.each do |(i1, i2)|
        i1next = (i1 + 1) % ANGLES.length
        i2next = (i2 + 1) % ANGLES.length
        p1 = between(i1, i1next)
        p2 = between(i2, i2next)
        dx = (p2[0] - p1[0])
        dy = (p2[1] - p1[1])
        mid_dist = Math.sqrt(dx*dx + dy*dy)
        c1 = normal_point(i1, i1next, mid_dist*CONTROL_LEN_MULT)
        c2 = normal_point(i2, i2next, mid_dist*CONTROL_LEN_MULT)

        show_control_points = false
        if show_control_points
          svg.circle cx: p1[0], cy: p1[1], r: 2, fill: 'green'
          svg.line x1: p1[0], y1: p1[1], x2: c1[0], y2: c1[1], stroke: 'green'
          svg.circle cx: c1[0], cy: c1[1], r: 2, fill: 'green'
          svg.circle cx: c2[0], cy: c2[1], r: 2, fill: 'green'
          svg.line x1: p2[0], y1: p2[1], x2: c2[0], y2: c2[1], stroke: 'green'
          svg.circle cx: p2[0], cy: p2[1], r: 2, fill: 'green'
        end

        svg.path d: "M#{p1[0]},#{p1[1]}C#{c1[0]} #{c1[1]},#{c2[0]} #{c2[1]},#{p2[0]} #{p2[1]}", stroke: @background, stroke_width: LINE_WIDTH*2, fill: 'none'
        svg.path d: "M#{p1[0]},#{p1[1]}C#{c1[0]} #{c1[1]},#{c2[0]} #{c2[1]},#{p2[0]} #{p2[1]}", stroke: 'white', stroke_width: LINE_WIDTH, fill: 'none'
      end

      centre = between(2, 7)

      if SHOW_NAME
        svg.text(
          @name,
          x: centre[0],
          y: centre[1],
          font_family: 'arial',
          font_size: 16,
          text_anchor: "middle",
          dominant_baseline: "middle"
        )
      end

      svg.polygon points: @points, stroke: 'black', fill: 'none'

      if @show_corner_indeces
        @points[0..-2].each.with_index do |p, idx|
          svg.circle cx: p[0], cy:p[1], r: 5, fill: "hsla(#{@hue}, 100%, 95%, 85%)"
          svg.text(
            @indexes[idx],
            x: p[0],
            y: p[1],
            font_family: 'arial',
            font_size: 6,
            text_anchor: "middle",
            dominant_baseline: "middle"
          )
        end
      end
    end
  end

  def into(svg)
    svg << self.svg
    self
  end
end


svg = Victor::SVG.new width: 475, height: 400

s1 = Spectre.new([100, 100], 0, 0).into(svg)
s2 = Spectre.new(s1.point(2), 8, 120).into(svg)
s3 = Spectre.new(s1.point(4), 8, 60).into(svg)
s4 = Spectre.new(s3.point(2), 10, 30).into(svg)
s5 = Spectre.new(s3.point(10), 8, 120).into(svg)
s6 = Spectre.new(s3.point(0), 2, 0).into(svg)
s7 = Spectre.new(s5.point(2), 0, 90).into(svg)
s8 = Spectre.new(s7.point(2), 10, 210).into(svg)
s9 = Spectre.new(s5.point(12), 0, -150).into(svg)
s10 = Spectre.new(s2.point(12), 6, -150).into(svg)
s11 = Spectre.new(s6.point(9), 3, 90).into(svg)
s12 = Spectre.new(s6.point(12), 0, 30).into(svg)
s13 = Spectre.new(s11.point(7), 3, 120).into(svg)
s14 = Spectre.new(s11.point(12), 6, 30).into(svg)
s15 = Spectre.new(s13.point(0), 6, 90).into(svg)
s16 = Spectre.new(s8.point(2), 0, 210).into(svg)
s17 = Spectre.new(s16.point(12), 6, 90).into(svg)
s18 = Spectre.new(s17.point(2), 0, 30).into(svg)
s19 = Spectre.new(s7.point(6), 0, 90).into(svg)
s20 = Spectre.new(s13.point(10), 10, 120).into(svg)
s21 = Spectre.new(s4.point(2), 10, 30).into(svg)
s22 = Spectre.new(s15.point(4), 8, 60).into(svg)

svg.save 'spectre'
```

</details>

## Mosaics

I've been toying with the idea of creating ceramic tiles with glazes specially chosen/mixed to have levels of reflectivity to sodium light at λ = ~589 nm that are different to their white/solar light grayscale mapping.

This would allow me to make a day/night split mosaic; in the day the full-colour mosaic would be of one picture, but by night the street lights' singular frequency lamps (if I can find anywhere that still uses sodium lamps!) would create a different grayscale image that would jump out!
