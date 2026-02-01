---
title: mini_magic and Apache Passenger
date: 2013-01-23T11:26:00+00:00
draft: false
emoji: 🌇
tags:
  - ruby
  - apache
  - passenger
  - code
  - fix
  - imagemagick
  - homebrew
  - from-tumblr
atUri: "at://did:plc:ephkzpinhaqcabtkugtbzrwu/site.standard.document/3mdryqt646s2y"
---
A note out there to anyone else having the same problem as me:

If you're trying to use [mini_magick](https://github.com/minimagick/minimagick) in a ruby application being run via [Passenger](https://www.phusionpassenger.com/) on Mac OS X Mountain Lion's Apache2, when you've installed imagemagick via [Homebrew](https://brew.sh) - _mouthful_ - then you may notice that that you get a Server Error when you try to process an image.

This may be because the `PATH` variable Apache is supplying to your ruby instance doesn't have `/usr/local/bin` in it, which is where homebrew installs to.

It seems that the `_www` user doesn't check the `/etc/paths` file for additions to the PATH variable, so you get squat.

My solution is hacky, but it gets the job done. I have a file called `/etc/apache2/other/passenger.conf` which holds the details for the virtual hosts for my passenger apps. I added the `SetEnv` line at the top to force Apache to set the PATH variable correctly.

It now looks something like this:

```Tcl
SetEnv PATH /usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin
ServerName my.rubyapp.com
ServerAlias my.rubyapp.com
DocumentRoot /var/www/myapp/public
RackEnv production
Order allow,deny
Allow from all
ErrorLog "/var/log/apache2/my.rubyapp.com_error-log"
CustomLog "/var/log/apache2/my.rubyapp.com_access-log" common
```
