---
date: "2009-10-29T00:47:51Z"
tags:
- imported
- from-twitter
---
I wrote a hacky ruby library to make numbers dimensional!

```ruby
p Amount.new 42,(Meter + Kilogram - Second - Second)
=> 42 Newtons
```
