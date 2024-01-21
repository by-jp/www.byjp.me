# Omnivore importer

This (even more shoddily built) script pulls my annotated articles in from Omnivore and turns them into bookmarks.

You can run it by [generating an API key](https://omnivore.app/settings/api), then running it from this repo:

```bash
OMNIVORE_API_KEY=<your key> go run .
```

Currently it assumes a few things about being in my blog's repo (eg. [where the bookmarks repo is](https://github.com/by-jp/www.byjp.me/blob/a0cc72a65dda3cba8bf37c554b160714a204765d/tools/import/omnivore/main.go#L38)) — but if you're comfortable with messing about with code you should be able to wrangle it to your needs.

**Or** let me know you like the idea of this (links in the article below), and I can turn it into something proper & easy to configure/run!

I've written about the limitations and plans for this tool on my blog, see [Bookmarks with Omnivore](https://www.byjp.me/posts/bookmarks-with-omnivore) for more :)
