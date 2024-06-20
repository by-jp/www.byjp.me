---
date: "2017-09-16T11:11:13Z"
tags:
- imported
- from-twitter
---
```ruby
def Hash.forever
  new { |h, k| h[k] = forever }
end

Hash.forever.tap { |h| h[:no]['need'][2] = 'init' }
# => {:no=>{"need"=>{2=>"init"}}}
```
