---
title: Ruby and Jenkins
date: 2013-03-01T09:37:00+00:00
draft: false
emoji: ⚙️
images:
tags:
  - jenkins
  - code
  - ruby
  - from-tumblr
---

I was trying to integrate my Cucumber and RSpec tests with Jenkins earlier this week and came across a bunch of character encoding errors.

It took me a while to figure out the problem, essentially Jenkins loads a session without all your usual environment variables, so your `PATH` and `LANG` won't be set to the same as a terminal window, making the outcomes different.

My Jenkins executed shell script now looks like this, note the PATH and LANG exports!

```sh
export LANG=en_GB.UTF-8
export PATH=/Users/admin/.rvm/gems/ruby-1.9.3-p385/bin:/Users/admin/.rvm/gems/ruby-1.9.3-p385@global/bin:/Users/admin/.rvm/rubies/ruby-1.9.3-p385/bin:/Users/admin/.rvm/bin:/usr/local/bin:$PATH
rvm use 1.9.3 --install --binary --fuzzy
bundle install
ruby --version
gem --version
rake spec
rake features
```
