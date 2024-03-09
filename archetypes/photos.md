---
title: "{{ replace .Name "-" " " | title }}"
media:
- url: "{{ .Name }}.webp"
  alt: 
date: {{ .Date }}
tags:
syndications:
- syndicate:instagram:{{ now.UnixNano | md5 }}
- syndicate:pixelfed:{{ now.UnixNano | md5 }}
- syndicate:bluesky:{{ now.UnixNano | md5 }}
---