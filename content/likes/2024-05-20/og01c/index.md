---
date: 2024-05-20T19:12:08.626+01:00
publishDate: 2024-05-20T19:12:08.626+01:00
likeOf: https://gabrielsieben.tech/2024/05/17/thinking-out-loud-2nd-gen-email/
references:
  https://gabrielsiebenTech/2024/05/17/thinkingOutLoud2NdGenEmail/:
    url: https://gabrielsieben.tech/2024/05/17/thinking-out-loud-2nd-gen-email/
    children:
      - type: entry
        name: Thinking out loud about 2nd-gen Email
        author:
          children:
            - type: card
              url: https://gabrielsieben.tech/author/gjsman/
              name: Gabriel Sieben
            - type: card
              url: https://gabrielsieben.tech/author/gjsman/
              name: Gabriel Sieben
        url:
          - https://gabrielsieben.tech/2024/05/17/thinking-out-loud-2nd-gen-email/
          - https://gabrielsieben.tech/2024/05/17/thinking-out-loud-2nd-gen-email/
        updated:
          - 2024-05-18T10:56:24-06:00
          - 2024-05-18T10:56:24-06:00
        content:
          html: |-
            <p><em>Note: This is just me, thinking out loud; you absolutely do not need to think that I have carefully thought this through, or that this is a good idea. With expectations set as low as possible, let’s continue.</em></p>



            <p>There are many old pieces of tech still in use, but there’s one that grinds my gears every time I try to use it: Email.</p>



            <p>For users, email works pretty well. Sometimes it sends too many emails to Junk, but Email is old, reliable, easy to understand, and relatively easy to search. It’s a good system, and I’m not eager to replace it with Slack anytime soon.</p>



            <p>However… the backend for email, is a mess. In escalating order (and “we” is used in a very imprecise, broad hand-waving sense for technologists):</p>



            <ul>
            <li>Many things in Email have no spec; even basic things. For example: When you reply, are you replying at the top of the message, or the bottom? It might even be a political question, <a href="https://web.archive.org/web/20041224190115/http://lists.suse.com/archive/suse-linux-e/2002-Oct/1698.html">depending on who you ask</a>. This has been worked around by email clients basically guessing the order and rarely even showing the email’s original text, putting layers between the user and the actual message.</li>



            <li>What HTML are you allowed to put in an email? Well… it depends. When there’s no spec, and Microsoft Outlook <a href="https://www.hteumeuleu.com/2020/outlook-rendering-engine/">abuses the Microsoft Word HTML renderer, it gets ugly</a>. There’s no guarantee the receiver even has an HTML renderer, and then it’s even more ugly.</li>



            <li>How do you encrypt an email, from the provider itself? <a href="https://moxie.org/2015/02/24/gpg-and-me.html">Well, we invented this thing called OpenPGP, but almost nobody used it. Then it turned out to have major flaws.</a></li>



            <li><a href="https://dmarcian.com/history-of-spf/">We couldn’t always make sure emails were authentic. So we invented SPF.</a></li>



            <li><a href="https://postmarkapp.com/glossary/dkim-domainkeys-identified-mail">Then it turned out SPF didn’t fix everything. So we invented DKIM.</a></li>



            <li><a href="https://dmarc.org/overview/">Then it turned out DKIM didn’t fix everything. So we invented DMARC.</a></li>



            <li><a href="https://www.zone.eu/blog/2024/05/17/bimi-and-dmarc-cant-save-you/">And then it turns out DKIM itself has major flaws that also bypass DMARC by extension.</a> </li>



            <li><a href="https://postmarkapp.com/blog/what-the-heck-is-bimi">Also, because we threw on yet another layer called BIMI, which itself relies on DMARC, when DMARC relies on DKIM, and DKIM has flaws, then it’s exceedingly bad for users.</a></li>



            <li><a href="https://dmarc.org/stats/dmarc/">Even if your sender has DMARC, 68.2% of records have their policy set to <code>p=none</code>. This tells DMARC to basically… do nothing!</a> (Sure, it technically speaking notifies the domain owner, but it’s clear the domain owners don’t care.)</li>



            <li>Did I mention all of the above, plus aggressive anti-spam policies, makes self-hosting email insanely difficult?</li>



            <li>Last, but not least, there’s the inane juggling of IP reputation. Some IP addresses are “cleaner” than other addresses, especially on shared systems like SendGrid or AWS SES. This makes signing up for a mass-mailing account, for whatever reason, messy; and causes countless surprise instances of legitimate emails going to Junk. Combine that with IP Address depletion, and the number of mostly clean addresses is shrinking over time.</li>
            </ul>



            <p>My gut reaction to the above is that we’ve got a lousy spec, with decades of cruft and unofficial spec, and we aren’t that great at securing it, or making sure messages are authentic. So… could we do better?</p>



            <p>Thus the hypothetical: <em>2nd-gen email</em>.</p>



            <p>Your initial reaction might be: That would be pointless, because not everyone would opt into it, and it would break compatibility all over the place. My thought is… that’s not necessarily a given. Imagine this:</p>



            <ul>
            <li>We create a new DNS record, called <code>MX2</code>. Most email services, then, would have an <code>MX2</code> and <code>MX</code> record. Older services only have <code>MX</code>.</li>



            <li>If an ancient, 20 year old email client, tries to send a message – it finds the <code>MX</code> record and sends the message just like normal. A modern client sees the <code>MX2</code> and sends the message there if it exists; otherwise, it falls back to <code>MX</code>.</li>



            <li>From there, the email services which implement <code>MX2</code> would publish a public date, on which all messages sent to them by the old <code>MX</code> record, will be automatically sent to Junk. If just Microsoft and Google alone agreed on such a date, that would be 40% of global email traffic.</li>
            </ul>



            <p>If the above looks slightly familiar, it’s because this strategy already worked, in a sense, with the transition from HTTP to HTTPS. We threw away a multi-decade-old protocol, for a new and more secure one. We set browsers to automatically upgrade the connection wherever possible, and now warn users about insecure connections when accessing HTTP (especially on login pages). Nevertheless – users can still visit HTTP pages, ancient browsers still work on HTTP, but most websites have gotten the memo and upgraded to HTTPS anyway. </p>



            <p>The incentive to upgrading to <code>MX2</code> would be simple: Your messages, while they still would arrive, would go to Junk automatically past the publicly posted date. No business wants that, even if users are already trained to expect, and act, like that can happen. Thus, the incentive to upgrade without truly breaking any day-to-day compatibility.</p>



            <p>Personally, I think that such a transition could go even faster than the HTTP to HTTPS transition. Self-hosted email is not very popular in part because of the complexity of the current email system, so between Microsoft, Google, Amazon, Zoho, GoDaddy, Gandhi, Wix, Squarespace, MailChimp, SparkPost, and SendGrid – you have most of the email market covered for the US; anyone not in the above list would quickly fold. The relative centralization of email, ironically, makes a mass upgrade to email much more achievable.</p>



            <p>What would a 2nd-gen email prioritize then? Everyone has different priorities, but I’d personally suggest the following which would hopefully win a broad enough consensus if this idea goes anywhere (though experts, of which I am not one, would have plenty of their own ideas):</p>



            <ol>
            <li>A standardized HTML specification for email; complete with a test suite for conformance. Or, maybe we just declare a version of the HTML5 spec to be officially binding and that’s the end of it.<br><br></li>



            <li>Headers for email chain preferences, or other email-specific preferences (i.e. Is this email chain a top-reply chain, or a bottom-reply chain? The client shouldn’t need to guess, or worse, ignore it.)<br><br></li>



            <li>If an email has a rich, HTML view; it should be required to come with a text-only, non-HTML copy of the body as well; for accessibility, compatibility, and privacy reasons.<br><br></li>



            <li>All MX2 records must have a public key embedded in the record. To send an email from the domain:<br><br>– A hash of the email content, and all headers, is created.<br>– This hash is then encrypted with the private key, corresponding to the record’s public key.<br>– This header is then added to the email, as the only permitted untrusted header.<br>– When an email is received, the header containing the hash is decrypted with the DNS public key, and the rest of the email is checked against the hash for integrity and authenticity. <br><br></li>



            <li>Point #4 is a lot like DKIM and DMARC right now, except:<br><br>– There would always be an automatic reject policy (<code>p=reject</code>) . Currently, only 19.6% of email services which even have DKIM are this stringent.<br>– If headers do need to be added to an email, the spec can carefully define carve-outs for where untrusted data can go (i.e. if the spam filter wanted to add a header).<br>– There also could be standardized carve-outs for, say, appending untrusted data from the receiving server to a message body (i.e. your business could add data to the body’s top or bottom indicating that the message from an external recipient and you have legal obligations, but your email client can also clearly show that this was not part of the original message and is not signed).<br>– As such, the signing would not need to work around email compatibility to such an extent as DKIM, reducing the likelihood of critical flaws. <br><br></li>



            <li>By simplifying the stack to the above, eliminating SPF, DKIM, and DMARC (and their respective configuration options), and standardizing on one record (MX2) for the future, running your own self-hosted email stack would become much easier. Additionally, the additional authenticity verifications would hopefully allow spam filters to be significantly less aggressive by authenticating against domains instead of IPs.<br><br></li>



            <li>Point #6 is the biggest change – we’re no longer authenticating, or caring, about the IP Address that’s sending the email. Every email can and always would be verified against the domain using MX2 records and the public keys in them. Send a fake spam email? It doesn’t have a signature, so it gets tossed without any heuristics. Send a real spam email? Block that domain when there’s complaints. Go after the registrar (or treat domains belonging to that registrar as suspicious) if needed. This would mostly eliminate the need for IP reputation by replacing it with domain reputation – which, at least to me, is a far superior standard with more understandable and controllable outcomes(1).<br><br></li>



            <li>Clients which implement MX2 can, optionally, have an updated encryption scheme to replace OpenPGP. Something like Apple’s Contact Key Verification. Hopefully there would be forward secrecy this time.</li>
            </ol>



            <p>If you have got great counterarguments, let me hear them.<br><br>(1) This would, perhaps, be the one and only “new feature” we could advertise to users. Not getting emails? You can just type in the name of the website, and always receive the emails.</p>



            <p><em>Edit 1, for clarification: For bulk senders, there would be multiple MX2 records on the domain, each containing a public key for every authorized sender. One of those records would have a marker indicating it as suitable for incoming mail.</em></p>



            <hr class="wp-block-separator has-alpha-channel-opacity">



            <p>Edit 2: This article has had a very large discussion on <a href="https://news.ycombinator.com/item?id=40392709">Hacker News</a>. While the discussion winds down, I have some additional thoughts from there:</p>



            <ul>
            <li>If there is an MX2 (ever), a sane way to share large files (like hundreds of megabytes, or even gigabytes) would be great. Designing the protocol wouldn’t be easy especially due to spam concerns, but if I had a nickel for how many links are shared to avoid email size limits… this is a real-world problem.</li>



            <li>MX2 will literally never happen if Google and Microsoft don’t join in. They would also, of course, have considerable control on the outcome. <em>However, </em>if even open-source communities and developers adopted MX2 because it was easy to implement and open source… you never know what grassroots can do. </li>



            <li>Part of me wonders what would happen if MX2 threw out SMTP for HTTP with a standardized REST API and JSON bodies. Sure – it would add a mountain of HTTP overhead and be more complex. However, it would sure as heck make implementing MX2 into a project quite easy in most programming languages, as it would just be a web server running on a custom port answering endpoints. REST APIs are also, despite their complexity, a well-documented system including for preventing spam (it’s not like Stripe or S3 lets people spam their APIs with garbage). I don’t know enough about SMTP to know if that’s a good idea – but I do know that SMTP is sub-optimal enough that Microsoft and Google don’t use it when exchanging messages with each other.</li>



            <li>There has been interesting commentary about being a <em>pull</em> protocol instead of a <em>push</em> protocol (i.e. instead of <em>sending</em> a message from X to Y; X sends Y a tiny standardized note saying to pick up a message from X). The most popular proposal of this was <a href="https://cr.yp.to/im2000.html">DJB’s Internet Mail 2000</a>.</li>



            <li>The idea of plain-text alternatives to HTML is probably impossible to enforce.</li>



            <li>As some commenters have pointed out, if the public key is always on the DNS, and every MX2 implementation is required to have that public key, a sending-server-to-receiving-server email encryption becomes possible.</li>



            <li>The idea of using HTML, at all, is controversial. Email was originally never designed for HTML, and the security risks of processing it are quite large. Using a superset of Markdown with style directives, or a customized XML schema, or even a new simple markup language all-together (Modern Mail Markup Language – M3L, claiming it now) might be an interesting thought experiment.</li>



            <li>A consistent point that came up was that standards drift – people don’t always implement the spec right, mistakes are made. I answer that, being a new standard, this is a chance to rigidly enforce the rules from the beginning. For example, we could put it in spec that any incoming message that’s not signed, despite that being immediately and easily verifiable by the sender, causes a 1 hour IP ban for laziness. Just an example.   </li>
            </ul>
          text: |-
            Note: This is just me, thinking out loud; you absolutely do not need to think that I have carefully thought this through, or that this is a good idea. With expectations set as low as possible, let’s continue.



            There are many old pieces of tech still in use, but there’s one that grinds my gears every time I try to use it: Email.



            For users, email works pretty well. Sometimes it sends too many emails to Junk, but Email is old, reliable, easy to understand, and relatively easy to search. It’s a good system, and I’m not eager to replace it with Slack anytime soon.



            However… the backend for email, is a mess. In escalating order (and “we” is used in a very imprecise, broad hand-waving sense for technologists):




            Many things in Email have no spec; even basic things. For example: When you reply, are you replying at the top of the message, or the bottom? It might even be a political question, depending on who you ask. This has been worked around by email clients basically guessing the order and rarely even showing the email’s original text, putting layers between the user and the actual message.



            What HTML are you allowed to put in an email? Well… it depends. When there’s no spec, and Microsoft Outlook abuses the Microsoft Word HTML renderer, it gets ugly. There’s no guarantee the receiver even has an HTML renderer, and then it’s even more ugly.



            How do you encrypt an email, from the provider itself? Well, we invented this thing called OpenPGP, but almost nobody used it. Then it turned out to have major flaws.



            We couldn’t always make sure emails were authentic. So we invented SPF.



            Then it turned out SPF didn’t fix everything. So we invented DKIM.



            Then it turned out DKIM didn’t fix everything. So we invented DMARC.



            And then it turns out DKIM itself has major flaws that also bypass DMARC by extension. 



            Also, because we threw on yet another layer called BIMI, which itself relies on DMARC, when DMARC relies on DKIM, and DKIM has flaws, then it’s exceedingly bad for users.



            Even if your sender has DMARC, 68.2% of records have their policy set to p=none. This tells DMARC to basically… do nothing! (Sure, it technically speaking notifies the domain owner, but it’s clear the domain owners don’t care.)



            Did I mention all of the above, plus aggressive anti-spam policies, makes self-hosting email insanely difficult?



            Last, but not least, there’s the inane juggling of IP reputation. Some IP addresses are “cleaner” than other addresses, especially on shared systems like SendGrid or AWS SES. This makes signing up for a mass-mailing account, for whatever reason, messy; and causes countless surprise instances of legitimate emails going to Junk. Combine that with IP Address depletion, and the number of mostly clean addresses is shrinking over time.




            My gut reaction to the above is that we’ve got a lousy spec, with decades of cruft and unofficial spec, and we aren’t that great at securing it, or making sure messages are authentic. So… could we do better?



            Thus the hypothetical: 2nd-gen email.



            Your initial reaction might be: That would be pointless, because not everyone would opt into it, and it would break compatibility all over the place. My thought is… that’s not necessarily a given. Imagine this:




            We create a new DNS record, called MX2. Most email services, then, would have an MX2 and MX record. Older services only have MX.



            If an ancient, 20 year old email client, tries to send a message – it finds the MX record and sends the message just like normal. A modern client sees the MX2 and sends the message there if it exists; otherwise, it falls back to MX.



            From there, the email services which implement MX2 would publish a public date, on which all messages sent to them by the old MX record, will be automatically sent to Junk. If just Microsoft and Google alone agreed on such a date, that would be 40% of global email traffic.




            If the above looks slightly familiar, it’s because this strategy already worked, in a sense, with the transition from HTTP to HTTPS. We threw away a multi-decade-old protocol, for a new and more secure one. We set browsers to automatically upgrade the connection wherever possible, and now warn users about insecure connections when accessing HTTP (especially on login pages). Nevertheless – users can still visit HTTP pages, ancient browsers still work on HTTP, but most websites have gotten the memo and upgraded to HTTPS anyway. 



            The incentive to upgrading to MX2 would be simple: Your messages, while they still would arrive, would go to Junk automatically past the publicly posted date. No business wants that, even if users are already trained to expect, and act, like that can happen. Thus, the incentive to upgrade without truly breaking any day-to-day compatibility.



            Personally, I think that such a transition could go even faster than the HTTP to HTTPS transition. Self-hosted email is not very popular in part because of the complexity of the current email system, so between Microsoft, Google, Amazon, Zoho, GoDaddy, Gandhi, Wix, Squarespace, MailChimp, SparkPost, and SendGrid – you have most of the email market covered for the US; anyone not in the above list would quickly fold. The relative centralization of email, ironically, makes a mass upgrade to email much more achievable.



            What would a 2nd-gen email prioritize then? Everyone has different priorities, but I’d personally suggest the following which would hopefully win a broad enough consensus if this idea goes anywhere (though experts, of which I am not one, would have plenty of their own ideas):




            A standardized HTML specification for email; complete with a test suite for conformance. Or, maybe we just declare a version of the HTML5 spec to be officially binding and that’s the end of it.



            Headers for email chain preferences, or other email-specific preferences (i.e. Is this email chain a top-reply chain, or a bottom-reply chain? The client shouldn’t need to guess, or worse, ignore it.)



            If an email has a rich, HTML view; it should be required to come with a text-only, non-HTML copy of the body as well; for accessibility, compatibility, and privacy reasons.



            All MX2 records must have a public key embedded in the record. To send an email from the domain:– A hash of the email content, and all headers, is created.– This hash is then encrypted with the private key, corresponding to the record’s public key.– This header is then added to the email, as the only permitted untrusted header.– When an email is received, the header containing the hash is decrypted with the DNS public key, and the rest of the email is checked against the hash for integrity and authenticity. 



            Point #4 is a lot like DKIM and DMARC right now, except:– There would always be an automatic reject policy (p=reject) . Currently, only 19.6% of email services which even have DKIM are this stringent.– If headers do need to be added to an email, the spec can carefully define carve-outs for where untrusted data can go (i.e. if the spam filter wanted to add a header).– There also could be standardized carve-outs for, say, appending untrusted data from the receiving server to a message body (i.e. your business could add data to the body’s top or bottom indicating that the message from an external recipient and you have legal obligations, but your email client can also clearly show that this was not part of the original message and is not signed).– As such, the signing would not need to work around email compatibility to such an extent as DKIM, reducing the likelihood of critical flaws. 



            By simplifying the stack to the above, eliminating SPF, DKIM, and DMARC (and their respective configuration options), and standardizing on one record (MX2) for the future, running your own self-hosted email stack would become much easier. Additionally, the additional authenticity verifications would hopefully allow spam filters to be significantly less aggressive by authenticating against domains instead of IPs.



            Point #6 is the biggest change – we’re no longer authenticating, or caring, about the IP Address that’s sending the email. Every email can and always would be verified against the domain using MX2 records and the public keys in them. Send a fake spam email? It doesn’t have a signature, so it gets tossed without any heuristics. Send a real spam email? Block that domain when there’s complaints. Go after the registrar (or treat domains belonging to that registrar as suspicious) if needed. This would mostly eliminate the need for IP reputation by replacing it with domain reputation – which, at least to me, is a far superior standard with more understandable and controllable outcomes(1).



            Clients which implement MX2 can, optionally, have an updated encryption scheme to replace OpenPGP. Something like Apple’s Contact Key Verification. Hopefully there would be forward secrecy this time.




            If you have got great counterarguments, let me hear them.(1) This would, perhaps, be the one and only “new feature” we could advertise to users. Not getting emails? You can just type in the name of the website, and always receive the emails.



            Edit 1, for clarification: For bulk senders, there would be multiple MX2 records on the domain, each containing a public key for every authorized sender. One of those records would have a marker indicating it as suitable for incoming mail.







            Edit 2: This article has had a very large discussion on Hacker News. While the discussion winds down, I have some additional thoughts from there:




            If there is an MX2 (ever), a sane way to share large files (like hundreds of megabytes, or even gigabytes) would be great. Designing the protocol wouldn’t be easy especially due to spam concerns, but if I had a nickel for how many links are shared to avoid email size limits… this is a real-world problem.



            MX2 will literally never happen if Google and Microsoft don’t join in. They would also, of course, have considerable control on the outcome. However, if even open-source communities and developers adopted MX2 because it was easy to implement and open source… you never know what grassroots can do. 



            Part of me wonders what would happen if MX2 threw out SMTP for HTTP with a standardized REST API and JSON bodies. Sure – it would add a mountain of HTTP overhead and be more complex. However, it would sure as heck make implementing MX2 into a project quite easy in most programming languages, as it would just be a web server running on a custom port answering endpoints. REST APIs are also, despite their complexity, a well-documented system including for preventing spam (it’s not like Stripe or S3 lets people spam their APIs with garbage). I don’t know enough about SMTP to know if that’s a good idea – but I do know that SMTP is sub-optimal enough that Microsoft and Google don’t use it when exchanging messages with each other.



            There has been interesting commentary about being a pull protocol instead of a push protocol (i.e. instead of sending a message from X to Y; X sends Y a tiny standardized note saying to pick up a message from X). The most popular proposal of this was DJB’s Internet Mail 2000.



            The idea of plain-text alternatives to HTML is probably impossible to enforce.



            As some commenters have pointed out, if the public key is always on the DNS, and every MX2 implementation is required to have that public key, a sending-server-to-receiving-server email encryption becomes possible.



            The idea of using HTML, at all, is controversial. Email was originally never designed for HTML, and the security risks of processing it are quite large. Using a superset of Markdown with style directives, or a customized XML schema, or even a new simple markup language all-together (Modern Mail Markup Language – M3L, claiming it now) might be an interesting thought experiment.



            A consistent point that came up was that standards drift – people don’t always implement the spec right, mistakes are made. I answer that, being a new standard, this is a chance to rigidly enforce the rules from the beginning. For example, we could put it in spec that any incoming message that’s not signed, despite that being immediately and easily verifiable by the sender, causes a 1 hour IP ban for laziness. Just an example.
        category: Uncategorized
      - type: card
        photo: https://secure.gravatar.com/avatar/dd3bb3c1bb48472b5aaf67916a3aae01?s=60&d=mm&r=g
      - type: card
        photo: https://secure.gravatar.com/avatar/df12a26f75dd1c440a2724d94b467f94?s=60&d=mm&r=g
      - type: card
        photo: https://secure.gravatar.com/avatar/41ac83e5f1de1d3155bd01e0d4367605?s=60&d=mm&r=g
      - type: card
        photo: https://secure.gravatar.com/avatar/b3934925e9c46e6f47a40531cb2e9cce?s=60&d=mm&r=g
      - type: card
        photo: https://secure.gravatar.com/avatar/c5ae818e7555d41886c33fdfdf89e66f?s=60&d=mm&r=g
      - type: card
        photo: https://secure.gravatar.com/avatar/5d67ef9ef33aadee2d44bcd23fba35f6?s=60&d=mm&r=g
      - type: card
        photo: https://secure.gravatar.com/avatar/dd3bb3c1bb48472b5aaf67916a3aae01?s=60&d=mm&r=g
        name: Brian Walker
      - type: card
        photo: https://secure.gravatar.com/avatar/df12a26f75dd1c440a2724d94b467f94?s=60&d=mm&r=g
        name: Kevin
      - type: card
        photo: https://secure.gravatar.com/avatar/41ac83e5f1de1d3155bd01e0d4367605?s=60&d=mm&r=g
        name: dan
      - type: card
        photo: https://secure.gravatar.com/avatar/b3934925e9c46e6f47a40531cb2e9cce?s=60&d=mm&r=g
        name: ricky712
      - type: card
        url: https://sangupta.com
        photo: https://secure.gravatar.com/avatar/c5ae818e7555d41886c33fdfdf89e66f?s=60&d=mm&r=g
        name: Sandeep Gupta
      - type: card
        url: https://zomglol.wtf/@jamie
        photo: https://secure.gravatar.com/avatar/4a9c33ccd38d1ab465fe83f16715c3b2?s=60&d=mm&r=g
        name: Jamie Gaskins
      - type: card
        photo: https://secure.gravatar.com/avatar/5d67ef9ef33aadee2d44bcd23fba35f6?s=60&d=mm&r=g
        name: Luis Fabiani
      - type: card
        url: https://alexeystar.com
        photo: https://secure.gravatar.com/avatar/23fceadcf7b580a481653868b2e2e86c?s=60&d=mm&r=g
        name: Alexey
slug: og01c
---
