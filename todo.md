# byJP blog

Things I might work on within my personal blog's software.

### Bugs

- [ ] Fix val.town donations link ([deprecation notice](https://docs.val.town/api/run/))
- [ ] Now page isn't current
- [ ] Micropub Replies look weird in a bunch of places
  - [ ] When posted via Echofeed: They include the "in reply to" initial stuff; this is too cumbersome
  - [ ] When posted via Echofeed: They should probably be "Quiet public", or ideally posted as replies to a fedi post, if there's one referenced in the source as a syncidcation
  - [x] They get posted on the homepage, when they shouldn't

### Ideas

- [ ] Show (internal) replies to this note
- [ ] Import from Omnivore daily  
- [ ] Record & display cross-links within site?
- [ ] Customise IndieKit
  - [ ] Add tags from #hashtags
  - [ ] Add links to memex from §refs
- [ ] Switch Memex to single file (easier to post to)
- [ ] Remove http://gowal.la & https://4sq.com links and/or posts that are useless without them
- [ ] Automate Gemini deploy
- [ ] Fix `mention-of` imports tools/import/webmentionio/main.go:172
- [ ] Move to referenced SVGs in header, rather than inline — they take up 4kb on every ~16-20kb page 😱

### Working

- [ ] Remove lychee failing links
- [ ] Import Facebook posts

### Done ✓

- [x] Add Links to Bookmarks & Notes  
- [x] Add IndieSearch support (my own experiment!)
- [x] Fix ThisIsMyJam links (convert to spotify embeds & add context)
- [x] Import tool for webmentions
- [x] Redesign homepage
  - [x] Include recent photo
  - [x] Include summaries of recent posts (of all types)
- [x] Retire Twitter
- [x] Remove 'instagr.am' posts (they're already imported as cross-posts)
- [x] Complete Twitter import  
  - [x] Figure out missing multi-photo tweets (eg. 1031265304675053568)  
  - [x] Remove reply tweets with no parent (the one to mhudack is a good example)  
  - [x] Fix hashtags with poor accessibility multiwordhashtags  
  - [x] Fix goodreads links (point to blog)  
  - [x] Fix instagram links (~~point to blog~~ delete; they're cross-posts anyway)
- [x] Add `inReplyTo` support (like bookmarks)
- [x] Add calendar to front page
- [x] Facebook import tooling
- [x] Claps!
- [x] Gemini?!
  - [x] Tags pages
  - [x] Calendar page
  - [x] Include link, location, date on event page
  - [x] Handle `<details>` in posts (example: ChefGPT)
  - [x] Actually serve it up
- [x] Record new webmentions on publish
  - Done via EchoFeed!
- [x] Videos in "photos" don't have a thumbnail, and don't work in the list.
- [x] Reviews don't appear on homepage?
- [x] Spoiler open styling has regressed!
- [x] Shortlinks
  - [x] All posts have an auto-generated unqiue shortlink
  - [x] Allow for customisation of shortlinks
  - [x] Create suitable `_redirects` files
  - [x] Import pre-existing shortlinks
  - [x] Add shortlink as `<link rel="shortlink">` on all relevant pages
  - [x] Add additional logic to _combine_ the redirects files as needed
  - [x] Upload new redirects to byjp.fyi repo
  - [x] Remove automation 😭 as IPFS can't handle large redirects files