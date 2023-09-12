---
title: Search
emoji: 🕵️‍♂️
type: site-infra
_build:
  list: never
---

I quite enjoy this site _not_ having search, as it means I (and others!) need to explore it to find what's there. Having said that, I do find myself in need of finding something I _know_ is there, but which I can't put my finger on…

Sigh, it's time for search; but I'm sneaking this away where it's not too visible, so you'll have to have previously found this to know you can put `/search` into your browser to use it 😉

<link href="/search/pagefind-ui.css" rel="stylesheet">
<script src="/search/pagefind-ui.js" type="text/javascript"></script>
<div id="search"></div>
<script>
    window.addEventListener('DOMContentLoaded', (event) => {
        new PagefindUI({
          element: "#search",
          showEmptyFilters: false,
        });
    });
</script>
