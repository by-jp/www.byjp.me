---
title: "Bye bye Twitter"
emoji: 💀
date: 2024-03-27T19:43:01.595Z
type: timeless
outputs:
- html
summary: I link here instead of linking to Twitter now.
---

As Twitter is no longer a part of the open web I don't want to link to it directly from pages elsewhere on my site any more. Particularly as you can't see much of that site without an account now, and owning an account increasingly seems to be funding and supporting [Dilbert Stark](https://bsky.app/profile/dilbert-stark.bsky.social).

<style>
.if-from-link { display: none }
.if-from-link p { margin-inline: auto; width: fit-content }
</style>
<div class="if-from-link">
<hr />

If you're _really_ keen on seeing that link, you should be able to get to it below, if the account still exists!

<a href="#" target="_blank" rel="nofollow noreferrer external"></a>
</div>
<script>
const explain = document.querySelector('.if-from-link')
const explainLink = explain.querySelector('a')
if (window.location.hash) {
  const xURL = new URL(window.location.hash.substring(1), 'https://x.com')
  const [_, user, extra] = xURL.pathname.split('/')
  explainLink.textContent = (extra) ? `A tweet from @${user}` : `@${user}'s twitter account`;
  explainLink.href = xURL.toString();
  explain.style.display = 'block';
}
</script>
