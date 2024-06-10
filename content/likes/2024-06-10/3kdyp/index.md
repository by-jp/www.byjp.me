---
date: 2024-06-10T15:01:38.546+01:00
publishDate: 2024-06-10T15:01:38.546+01:00
likeOf: https://rscottjones.com/analytics-for-the-personal-web/
references:
  https://rscottjonesCom/analyticsForThePersonalWeb/:
    url: https://rscottjones.com/analytics-for-the-personal-web/
    type: feed
    children:
      - type: entry
        name: Analytics for the personal web
        content:
          html: |-
            <p>If<sup data-fn="3ec72dbf-448a-4115-afc1-24c6c9775902" class="fn"><a id="3ec72dbf-448a-4115-afc1-24c6c9775902-link" href="#3ec72dbf-448a-4115-afc1-24c6c9775902">1</a></sup> you have a website, you probably have some form of site analytics. That might be something like Jetpack for WordPress, or Google Analytics (ick, stop that), or perhaps a more privacy-focused alternative like Plausible, or even a lightweight indie version like Tinylytics.</p>



            <p>It can be useful to see how much traffic your posts are getting, especially if they’re important to your personal income.</p>



            <p>But for many of us, our sites aren’t businesses—they’re forms of self expression and vehicles for connecting with others. If you’re not in it for the clicks, then none of these alternatives really prioritize what you might care most about: <strong>connecting with the readers that showed up to read what you wrote.</strong></p>



            <p>I’m a layperson who doesn’t know how all of this stuff actually works under the hood, so I don’t know what’s technically feasible or not. But if I <em>could</em> design my own analytics service, here are some thoughts on what it might look like.</p>



            <h2 id="referrers-aka-new-friends" class="wp-block-heading">Referrers (aka, new friends)</h2>



            <p>First and foremost, the dashboard would be focused on which webpages have linked to you.</p>



            <p>Links are the heart of the personal web, so knowing that someone else found your post interesting and/or useful enough to link to is pure <em>personal web fuel</em>—the feedback and motivation we all adore. Links can also be the start of an asynchronous conversation about the topic, so it’s useful to know when that’s happening.</p>



            <p>That means that referring sites would be front and center. No, I don’t mean Google and other search engines—I don’t need to know which search queries got them to me, either<sup data-fn="85ed3e68-12ae-46d1-9aa0-a0a73e8f924e" class="fn"><a href="#85ed3e68-12ae-46d1-9aa0-a0a73e8f924e" id="85ed3e68-12ae-46d1-9aa0-a0a73e8f924e-link">2</a></sup>. Same with aggregator sites like Reddit or Hacker News.</p>



            <p>And I don’t need to know <em>how much</em> traffic came from which social media platform, though I’d love if it could direct me to the actual social media posts that contain the link—that would allow me to add more information, answer any questions, or otherwise participate in the conversation.</p>



            <p>Ideally, I’d love to see <em>new</em> referrers highlighted first. If this is the first time that someone has clicked on a link from a particular webpage and landed on my site, that’s something I’d like to notice quickly.</p>



            <p>And I’d like to be able to hide—from this dashboard, at least—certain referrers too. Perhaps that happens automatically after a certain amount of time (a week? a month?). If I already know that Johnny has linked to me, and that people often click on that link and find themselves on my site, then I don’t need this info to be displayed day after day after day. I can tell you right now what all my top referrers are because they are <em>always the same</em> due to a few popular sites that link to mine. Don’t make me scroll down every single time to look for the new stuff.</p>



            <h2 id="kudos" class="wp-block-heading">Kudos</h2>



            <p>I signed up for tinylytics earlier this year almost exclusively for its kudos feature. It seemed like a fun “like this post” button, displayed as an emoji after the content of each post. A public counter of sorts for how many folks had appreciated the post.</p>



            <p>So I’d love some sort of low-friction way for someone to express the sentiment of “hello, I saw this,” or “I appreciated reading this,” or “I found this useful.” I’d want a listing of the pages that had received “new” kudos, alongside the overall total for those pages. And I’d also like to be able to delve into which pages have the most overall kudos—aka, the ones that have been the most impactful. Kudos are one of the few metrics I’d pay attention to, as those tiny pats-on-the-back and thank-yous are what inspire me to keep blogging.</p>



            <h2 id="comments-and-replies" class="wp-block-heading">Comments and replies</h2>



            <p>If my site accepts visitor comments, or publishes its posts in a way that others can reply to the post and be displayed on the site (such as how micro.blog will retrieve and display replies sent by mastodon, bluesky, or via its own platform), or accepts pingbacks or webmentions, then those numbers should be displayed for each of the posts as well.</p>



            <p><strong>In short, I want to see all of the latest interactions my readers have had with my posts.</strong></p>



            <h2 id="recently-posted" class="wp-block-heading">Recently posted</h2>



            <p>Another section of the main dashboard should be devoted to posts I’ve recently published, perhaps the last 5 or 10 or so.</p>



            <p>This would show the post name, plus the number of kudos for that page, the number of page visitors, and the number of referrers linking to that page. And of course, add in those interactions mentioned above, too. A quick glance at this section would give me a sense of how my recent posts are resonating.</p>



            <h2 id="visits" class="wp-block-heading">Visits</h2>



            <p>I don’t find it useful to pour over detailed visitor stats, especially not for older posts. Personal websites are about writing about what you want to write about, not chasing views. </p>



            <p>While every analytics service should keep track of the basic visit stats, most of that data should be hidden away out of view, accessible only if you really need it. I don’t need to see which operating system or browser people are using (certainly not on a per day or month basis), but I might want to see an aggregate of that information every so often when I’m redesigning my site.</p>



            <p>Crucially, I don’t need a daily or weekly or monthly accounting of how my posts <em>“performed.”</em> In fact, I don’t <em>want</em> to see a bar graph comparing the number of hits I got this month with how many I got last month. Even if I cared—and I do not—there’s really no actionable information in those bar graphs, anyway. But it’s usually the most prominent feature of any analytics app.</p>



            <p>That said, it might be fun to have a small alert if something is really blowing up. And I think it’d be fun if an analytics system occasionally posted your all-time views when you passed some round milestones: “Congrats, Scott, 50,000 people have visited your site to read your stuff – that’s cool! Keep posting, my friend.”</p>



            <p>Besides, I hope that many people are reading my posts via RSS anyway! Speaking of that, I’d also want to see a rough approximation of how many people have subscribed to my RSS feed, <a href="https://docs.bearblog.dev/analytics/">as Bear provides</a>.</p>



            <h2 id="tracking-clicks" class="wp-block-heading">Tracking clicks</h2>



            <p>I’m not sure if there’s much value in seeing which links people clicked on in my posts—whether that’s to other posts on my own site, or “outgoing” links to other sites. Again, I’m not exactly sure how knowing this would improve a personal website. Same with “bounce rates.” I’m happy enough they read the page they came to.</p>



            <h2 id="other-fun-features" class="wp-block-heading">Other fun features</h2>



            <p>There are number of other fun features that could add some additional joy to the personal web for anyone using this imagined analytics service. These aren’t important for analytics purposes, per se, but are some additional ideas that could be integrated into such a service anyway.</p>



            <h3 id="reply-via" class="wp-block-heading">Reply via</h3>



            <p>One of the best ways to start a conversation about a blog post is to email the author. So adding a “Reply by email” button to the bottom of the post, perhaps next to the kudos button, is a great way to facilitate that—especially for sites that don’t offer comments.</p>



            <p>One the plugin features I like at micro.blog is that you can sign in to reply with&nbsp;<a href="https://micro.blog/account/comments/39243461/mb?url=https://rsjon.es/2024/06/09/wait-so-brigham.html">Micro.blog</a>,&nbsp;<a href="https://micro.blog/account/comments/39243461/mastodon?url=https://rsjon.es/2024/06/09/wait-so-brigham.html">Mastodon</a>, or&nbsp;<a href="https://micro.blog/account/comments/39243461/bluesky?url=https://rsjon.es/2024/06/09/wait-so-brigham.html">Bluesky</a>&nbsp;right below the post, like you might in a comment form. This seems to take quite a bit more magic behind the scenes, but seems like an excellent way to foster conversations as well.</p>



            <h3 id="country-flags" class="wp-block-heading">Country flags</h3>



            <p>I don’t really need to know the origin country of visitors to each post, but it is fun to see how the aggregate of your posts have been read by people from around the world. <a href="https://tinylytics.app/docs/showing_countries">Tinylytics shows </a>a fun grouping of the flag emojis showing all the countries your site visitors are from. That’d be a great feature to include.</p>



            <h3 id="hit-counters" class="wp-block-heading">Hit counters</h3>



            <p>Hit counters were all the rage on the early internet. I’m generally uninterested in displaying one on my own site, but I can see how it might be something fun to show, especially on specific pages.</p>



            <p>For instance, it could be fun to be able to integrate a visitor count on any pages you might have devoted to obscure topics: “Somehow, you’re the 1125th person to read this weird story about how I went door-to-door in Cawker City, Kansas looking for some twine I could add to the World’s Largest Ball of Twine and was confronted by a kid in tighty-whities and we just pretended that was a normal everyday sort of thing while he fetched this here stranger some twine and I’m just not really sure what that says about my own personal choices that day, or if I’m honest, about your own personal interests and how you chose to spend your time today.” I’m just saying, I think I could have some fun with it. It’d be great if this could be integrated within paragraph text, and not simply a graphic counter at the bottom of the page.</p>



            <h3 id="top-popular-lists" class="wp-block-heading">Top/Popular lists</h3>



            <p>Another fun feature would be providing a way to display your most popular posts (whether by views, kudos, comments, links, etc) on a public page on your site, with the ability to filter it by time frame. Just a simple bulleted list of links would be great: “Here are the most popular posts from this year,” or “Here are the 20 posts that were given the most number of kudos in 2023.” What might also be fun? Showing a listing of the least viewed pages, too: “these unloved posts are lonely and could use a hug.”</p>



            <h3 id="guestbook" class="wp-block-heading">Guestbook</h3>



            <p>Guestbooks were an old school way to leave a comment on a website. Sometimes they were simply fun messages (“Hey, I was here!”), but often they acted like a general comment for the entire site (“Hey, I always enjoy reading your posts, they’re always so inspiring. Keep it up!”). Some folks have been bringing them back, and I support that.</p>



            <h3 id="webrings" class="wp-block-heading">Webrings</h3>



            <p>Webrings were among the early “discovery” tools for finding fun websites, but eventually fell out of fashion. Luckily, they too have started making a comeback on the personal web.</p>



            <p>This may seem to be an odd feature for an analytics tool to offer. But if you’re designing a website service <em>specifically for the personal web</em>, then your customers probably all have personal websites. So a service like this is perfectly positioned to help run a webring—or several topical ones—of its customers.</p>



            <h3 id="discovery" class="wp-block-heading">Discovery</h3>



            <p>Similarly, an analytics service for the personal web could do a variety of things to help others find their customers who have a personal website. Perhaps that’s running something like <a href="https://bearblog.dev/discover/">Bear’s extremely clever upvoting (kudos!) discovery feed</a>. Or maybe it’s hosting an old fashioned web directory, where customers can add their site into appropriate categories. Or maybe it’s something else entirely.</p>



            <h2 id="the-tldr-version" class="wp-block-heading">The tl;dr version</h2>



            <p>If you were building an analytics service for the personal web, you’d prioritize an entirely different set of data. You’d emphasize the ways that readers were connecting with your posts—and the resulting conversations and relationships that could ensue—not the traffic, search performance, or conversions that were resulting. You’d focus on visitors as human beings who had taken the time to visit your site and read your posts, not as faceless datapoints you manipulate in spreadsheets. You’d focus on the people who did show up, not how you can game the system to increase site growth.</p>



            <h2 id="whats-missing" class="wp-block-heading">What’s missing?</h2>



            <p>Have an additional suggestion you think I missed? <a href="https://rscottjones.com/contact/">Let me know</a>.</p>



            <h2 id="someone-please-please-build-this-for-me" class="wp-block-heading">Someone please <em>please</em> build this for me!</h2>



            <p>I promise to be your customer.</p>



            <hr class="wp-block-separator has-alpha-channel-opacity is-style-dots">


            <ol class="wp-block-footnotes"><li id="3ec72dbf-448a-4115-afc1-24c6c9775902"><em>You do have a website</em>, don’t you? No?! <a href="https://rscottjones.com/why-you-should-have-a-website/">Oh, you definitely should</a>. Stop reading this and <a href="https://rscottjones.com/the-easiest-ways-to-start-your-own-personal-website/">go get one</a> instead. <a href="#3ec72dbf-448a-4115-afc1-24c6c9775902-link" aria-label="Jump to footnote reference 1">↩︎</a></li><li id="85ed3e68-12ae-46d1-9aa0-a0a73e8f924e">My most popular posts apparently rank high on Google and therefore the same set of queries sends people my way. I’m tired of seeing the same results over and over. I’m <em>not trying</em> to rank high for any of these, so there’s no need to monitor how things are “performing.” <a href="#85ed3e68-12ae-46d1-9aa0-a0a73e8f924e-link" aria-label="Jump to footnote reference 2">↩︎</a></li></ol>
          text: |-
            If1 you have a website, you probably have some form of site analytics. That might be something like Jetpack for WordPress, or Google Analytics (ick, stop that), or perhaps a more privacy-focused alternative like Plausible, or even a lightweight indie version like Tinylytics.



            It can be useful to see how much traffic your posts are getting, especially if they’re important to your personal income.



            But for many of us, our sites aren’t businesses—they’re forms of self expression and vehicles for connecting with others. If you’re not in it for the clicks, then none of these alternatives really prioritize what you might care most about: connecting with the readers that showed up to read what you wrote.



            I’m a layperson who doesn’t know how all of this stuff actually works under the hood, so I don’t know what’s technically feasible or not. But if I could design my own analytics service, here are some thoughts on what it might look like.



            Referrers (aka, new friends)



            First and foremost, the dashboard would be focused on which webpages have linked to you.



            Links are the heart of the personal web, so knowing that someone else found your post interesting and/or useful enough to link to is pure personal web fuel—the feedback and motivation we all adore. Links can also be the start of an asynchronous conversation about the topic, so it’s useful to know when that’s happening.



            That means that referring sites would be front and center. No, I don’t mean Google and other search engines—I don’t need to know which search queries got them to me, either2. Same with aggregator sites like Reddit or Hacker News.



            And I don’t need to know how much traffic came from which social media platform, though I’d love if it could direct me to the actual social media posts that contain the link—that would allow me to add more information, answer any questions, or otherwise participate in the conversation.



            Ideally, I’d love to see new referrers highlighted first. If this is the first time that someone has clicked on a link from a particular webpage and landed on my site, that’s something I’d like to notice quickly.



            And I’d like to be able to hide—from this dashboard, at least—certain referrers too. Perhaps that happens automatically after a certain amount of time (a week? a month?). If I already know that Johnny has linked to me, and that people often click on that link and find themselves on my site, then I don’t need this info to be displayed day after day after day. I can tell you right now what all my top referrers are because they are always the same due to a few popular sites that link to mine. Don’t make me scroll down every single time to look for the new stuff.



            Kudos



            I signed up for tinylytics earlier this year almost exclusively for its kudos feature. It seemed like a fun “like this post” button, displayed as an emoji after the content of each post. A public counter of sorts for how many folks had appreciated the post.



            So I’d love some sort of low-friction way for someone to express the sentiment of “hello, I saw this,” or “I appreciated reading this,” or “I found this useful.” I’d want a listing of the pages that had received “new” kudos, alongside the overall total for those pages. And I’d also like to be able to delve into which pages have the most overall kudos—aka, the ones that have been the most impactful. Kudos are one of the few metrics I’d pay attention to, as those tiny pats-on-the-back and thank-yous are what inspire me to keep blogging.



            Comments and replies



            If my site accepts visitor comments, or publishes its posts in a way that others can reply to the post and be displayed on the site (such as how micro.blog will retrieve and display replies sent by mastodon, bluesky, or via its own platform), or accepts pingbacks or webmentions, then those numbers should be displayed for each of the posts as well.



            In short, I want to see all of the latest interactions my readers have had with my posts.



            Recently posted



            Another section of the main dashboard should be devoted to posts I’ve recently published, perhaps the last 5 or 10 or so.



            This would show the post name, plus the number of kudos for that page, the number of page visitors, and the number of referrers linking to that page. And of course, add in those interactions mentioned above, too. A quick glance at this section would give me a sense of how my recent posts are resonating.



            Visits



            I don’t find it useful to pour over detailed visitor stats, especially not for older posts. Personal websites are about writing about what you want to write about, not chasing views. 



            While every analytics service should keep track of the basic visit stats, most of that data should be hidden away out of view, accessible only if you really need it. I don’t need to see which operating system or browser people are using (certainly not on a per day or month basis), but I might want to see an aggregate of that information every so often when I’m redesigning my site.



            Crucially, I don’t need a daily or weekly or monthly accounting of how my posts “performed.” In fact, I don’t want to see a bar graph comparing the number of hits I got this month with how many I got last month. Even if I cared—and I do not—there’s really no actionable information in those bar graphs, anyway. But it’s usually the most prominent feature of any analytics app.



            That said, it might be fun to have a small alert if something is really blowing up. And I think it’d be fun if an analytics system occasionally posted your all-time views when you passed some round milestones: “Congrats, Scott, 50,000 people have visited your site to read your stuff – that’s cool! Keep posting, my friend.”



            Besides, I hope that many people are reading my posts via RSS anyway! Speaking of that, I’d also want to see a rough approximation of how many people have subscribed to my RSS feed, as Bear provides.



            Tracking clicks



            I’m not sure if there’s much value in seeing which links people clicked on in my posts—whether that’s to other posts on my own site, or “outgoing” links to other sites. Again, I’m not exactly sure how knowing this would improve a personal website. Same with “bounce rates.” I’m happy enough they read the page they came to.



            Other fun features



            There are number of other fun features that could add some additional joy to the personal web for anyone using this imagined analytics service. These aren’t important for analytics purposes, per se, but are some additional ideas that could be integrated into such a service anyway.



            Reply via



            One of the best ways to start a conversation about a blog post is to email the author. So adding a “Reply by email” button to the bottom of the post, perhaps next to the kudos button, is a great way to facilitate that—especially for sites that don’t offer comments.



            One the plugin features I like at micro.blog is that you can sign in to reply with Micro.blog, Mastodon, or Bluesky right below the post, like you might in a comment form. This seems to take quite a bit more magic behind the scenes, but seems like an excellent way to foster conversations as well.



            Country flags



            I don’t really need to know the origin country of visitors to each post, but it is fun to see how the aggregate of your posts have been read by people from around the world. Tinylytics shows a fun grouping of the flag emojis showing all the countries your site visitors are from. That’d be a great feature to include.



            Hit counters



            Hit counters were all the rage on the early internet. I’m generally uninterested in displaying one on my own site, but I can see how it might be something fun to show, especially on specific pages.



            For instance, it could be fun to be able to integrate a visitor count on any pages you might have devoted to obscure topics: “Somehow, you’re the 1125th person to read this weird story about how I went door-to-door in Cawker City, Kansas looking for some twine I could add to the World’s Largest Ball of Twine and was confronted by a kid in tighty-whities and we just pretended that was a normal everyday sort of thing while he fetched this here stranger some twine and I’m just not really sure what that says about my own personal choices that day, or if I’m honest, about your own personal interests and how you chose to spend your time today.” I’m just saying, I think I could have some fun with it. It’d be great if this could be integrated within paragraph text, and not simply a graphic counter at the bottom of the page.



            Top/Popular lists



            Another fun feature would be providing a way to display your most popular posts (whether by views, kudos, comments, links, etc) on a public page on your site, with the ability to filter it by time frame. Just a simple bulleted list of links would be great: “Here are the most popular posts from this year,” or “Here are the 20 posts that were given the most number of kudos in 2023.” What might also be fun? Showing a listing of the least viewed pages, too: “these unloved posts are lonely and could use a hug.”



            Guestbook



            Guestbooks were an old school way to leave a comment on a website. Sometimes they were simply fun messages (“Hey, I was here!”), but often they acted like a general comment for the entire site (“Hey, I always enjoy reading your posts, they’re always so inspiring. Keep it up!”). Some folks have been bringing them back, and I support that.



            Webrings



            Webrings were among the early “discovery” tools for finding fun websites, but eventually fell out of fashion. Luckily, they too have started making a comeback on the personal web.



            This may seem to be an odd feature for an analytics tool to offer. But if you’re designing a website service specifically for the personal web, then your customers probably all have personal websites. So a service like this is perfectly positioned to help run a webring—or several topical ones—of its customers.



            Discovery



            Similarly, an analytics service for the personal web could do a variety of things to help others find their customers who have a personal website. Perhaps that’s running something like Bear’s extremely clever upvoting (kudos!) discovery feed. Or maybe it’s hosting an old fashioned web directory, where customers can add their site into appropriate categories. Or maybe it’s something else entirely.



            The tl;dr version



            If you were building an analytics service for the personal web, you’d prioritize an entirely different set of data. You’d emphasize the ways that readers were connecting with your posts—and the resulting conversations and relationships that could ensue—not the traffic, search performance, or conversions that were resulting. You’d focus on visitors as human beings who had taken the time to visit your site and read your posts, not as faceless datapoints you manipulate in spreadsheets. You’d focus on the people who did show up, not how you can game the system to increase site growth.



            What’s missing?



            Have an additional suggestion you think I missed? Let me know.



            Someone please please build this for me!



            I promise to be your customer.






            You do have a website, don’t you? No?! Oh, you definitely should. Stop reading this and go get one instead. ↩︎My most popular posts apparently rank high on Google and therefore the same set of queries sends people my way. I’m tired of seeing the same results over and over. I’m not trying to rank high for any of these, so there’s no need to monitor how things are “performing.” ↩︎
slug: 3kdyp
---
