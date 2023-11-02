---
title: "{{ replace .Name "-" " " | title }}"
media:
- "{{ .Name }}.webp"
date: {{ .Date }}
tags:
syndications:
- syndicate:instagram:{{ now.UnixNano | md5 }}
- syndicate:pixelfed:{{ now.UnixNano | md5 }}
- syndicate:bluesky:{{ now.UnixNano | md5 }}
---