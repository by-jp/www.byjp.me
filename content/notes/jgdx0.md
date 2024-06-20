---
date: 2024-02-22T08:13:10.058Z
publishDate: 2024-02-22T08:13:10.058Z
---
Using [Mona](https://apps.apple.com/gb/app/mona-for-mastodon/id1659154653?uo=4) for mac & want to redirect web links? This kludge may work for you!

1. Install [Redirector](https://chromewebstore.google.com/detail/redirector/pajiegeliagebegjdhebejdlknciafen)
2. Redirect `^https://<your instance>/@([^/@] )@([^/@] )(/. )?$` to `mona://$2/@$1$3`
3. Redirect `^https://([^/] )/@([^/@] )(/. )?$` to `mona://$1/@$2$3`
4. Hope no-one else uses 'at' symbols like this? 😬
