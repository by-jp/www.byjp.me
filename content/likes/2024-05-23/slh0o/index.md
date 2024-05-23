---
date: 2024-05-23T11:30:24.744+01:00
publishDate: 2024-05-23T11:30:24.744+01:00
likeOf: https://seirdy.one/posts/2021/03/10/search-engines-with-own-indexes/
references:
  https://seirdyOne/posts/2021/03/10/searchEnginesWithOwnIndexes/:
    url: https://seirdy.one/posts/2021/03/10/search-engines-with-own-indexes/
    children:
      - type: entry
        name: A look at search engines with their own indexes
        author:
          type: card
          url: https://seirdy.one/
          uid: https://seirdy.one/
          photo: https://seirdy.one/favicon.1250396055.png
          name: Rohan “Seirdy” Kumar
          givenName: Rohan
          additionalName: Seirdy
          familyName: Kumar
        url: https://seirdy.one/posts/2021/03/10/search-engines-with-own-indexes/
        syndication: gemini://seirdy.one/posts/2021/03/10/search-engines-with-own-indexes/index.gmi
        updated: 2024-05-11T00:13:17Z
        content:
          html: |-
            <section role="doc-preface">
            					<h2 id="preface">Preface</h2>
            					<p>This is a cursory review of all the indexing search engines I have been able to find.</p>
            					<p>The three dominant English search engines with their own indexes<sup><a href="#fn:1" id="fnref:1" role="doc-noteref">note&nbsp;1</a></sup> are Google, Bing, and Yandex (<abbr title="Google, Bing, Yandex">GBY</abbr>). Many alternatives to GBY exist, but almost none of them have their own results; instead, they just source their results from GBY.</p>
            					<p>With that in mind, I decided to test and catalog all the different indexing search engines I could find. I prioritized breadth over depth, and encourage readers to try the engines out themselves if they’d like more information.</p>
            					<p>This page is a “living document” that I plan on updating indefinitely. Check for updates once in a while if you find this page interesting. Feel free to send me suggestions, updates, and corrections; I’d especially appreciate help from those who speak languages besides English and can evaluate a non-English indexing search engine. Contact info is in the article footer.</p>
            					<p>I plan on updating the engines in the top two categories with more info comparing the structured/linked data the engines leverage (RDFa vocabularies, microdata, microformats, JSON-LD, etc.) to help authors determine which formats to use.</p>
            				</section>
            				<details id="toc">
            					<summary>Toggle table of contents</summary>
            					<nav id="TableOfContents" role="doc-toc" aria-labelledby="toc-hd" tabindex="-1">
            						<h2 id="toc-hd">Table of Contents</h2>
            						<ol>
            							<li>
            								<a href="#about-the-list">About the list</a>
            							</li>
            							<li>
            								<a href="#general-indexing-search-engines">General indexing search-engines</a>
            								<ol>
            									<li>
            										<a href="#large-indexes-good-results">Large indexes, good results</a>
            									</li>
            									<li>
            										<a href="#smaller-indexes-or-less-relevant-results">Smaller indexes or less relevant results</a>
            									</li>
            									<li>
            										<a href="#smaller-indexes-hit-and-miss">Smaller indexes, hit-and-<wbr>miss</a>
            									</li>
            									<li>
            										<a href="#fledgling-engines">Fledgling engines</a>
            									</li>
            									<li>
            										<a href="#semi-independent-indexes">Semi-independent indexes</a>
            									</li>
            								</ol>
            							</li>
            							<li>
            								<a href="#non-generalist-search">Non-generalist search</a>
            								<ol>
            									<li>
            										<a href="#small-or-non-commercial-web">Small or non-commercial Web</a>
            									</li>
            									<li>
            										<a href="#site-finders">Site finders</a>
            									</li>
            									<li>
            										<a href="#other">Other</a>
            									</li>
            								</ol>
            							</li>
            							<li>
            								<a href="#other-languages">Other languages</a>
            								<ol>
            									<li>
            										<a href="#big-indexes">Big indexes</a>
            									</li>
            									<li>
            										<a href="#smaller-indexes">Smaller indexes</a>
            									</li>
            								</ol>
            							</li>
            							<li>
            								<a href="#almost-qualified">Almost qualified</a>
            							</li>
            							<li>
            								<a href="#misc">Misc</a>
            							</li>
            							<li>
            								<a href="#search-engines-without-a-web-interface">Search engines without a web interface</a>
            							</li>
            							<li>
            								<a href="#graveyard">Graveyard</a>
            							</li>
            							<li>
            								<a href="#exclusions">Exclusions</a>
            							</li>
            							<li>
            								<a href="#rationale">Rationale</a>
            								<ol>
            									<li>
            										<a href="#conflicts-of-interest">Conflicts of interest</a>
            									</li>
            									<li>
            										<a href="#information-diversity">Information diversity</a>
            									</li>
            								</ol>
            							</li>
            							<li>
            								<a href="#methodology">Method­ology</a>
            								<ol>
            									<li>
            										<a href="#discovery">Discovery</a>
            									</li>
            									<li>
            										<a href="#criteria-for-inclusion">Criteria for inclusion</a>
            									</li>
            									<li>
            										<a href="#evaluation">Evaluation</a>
            									</li>
            									<li>
            										<a href="#caveats">Caveats</a>
            									</li>
            								</ol>
            							</li>
            							<li>
            								<a href="#findings">Findings</a>
            							</li>
            							<li>
            								<a href="#acknowledgements">Acknow­ledgements</a>
            							</li>
            						</ol>
            					</nav>
            				</details>
            				<h2 id="about-the-list" tabindex="-1">About the list</h2>
            				<aside role="none">
            					<a href="#about-the-list" aria-labelledby="about-the-list-prefix about-the-list">
            						<span id="about-the-list-prefix">Permalink to section</span>
            					</a>
            				</aside>
            				<p>I discuss my motivation for making this page in the <a href="#rationale">Rationale section</a>.</p>
            				<p>I primarily evaluated English-speaking search engines because that’s my primary language. With some difficulty, I could probably evaluate a Spanish one; however, I wasn’t able to find many Spanish-language engines powered by their own crawlers.</p>
            				<p>I mention details like “allows site submissions” and structured data support where I can only to inform authors about their options, not as points in engines’ favor.</p>
            				<p>See the <a href="#methodology">Methodology section</a> at the bottom to learn how I evaluated each one.</p>
            				<h2 id="general-indexing-search-engines" tabindex="-1">General indexing search-engines</h2>
            				<aside role="none">
            					<a href="#general-indexing-search-engines" aria-labelledby="general-indexing-search-engines-prefix general-indexing-search-engines">
            						<span id="general-indexing-search-engines-prefix">Permalink to section</span>
            					</a>
            				</aside>
            				<h3 id="large-indexes-good-results" tabindex="-1">Large indexes, good results</h3>
            				<p>These are large engines that pass all my standard tests and more.</p>
            				<dl>
            					<dt>Google</dt>
            					<dd>The biggest index. Allows submitting pages and sitemaps for crawling, and <a href="https://developers.google.com/search/docs/advanced/sitemaps/build-sitemap#addsitemap">even supports WebSub</a> to automate the process. Powers a few other engines:
            <ul><li><p>A former version of <a href="https://www.startpage.com/">Startpage</a>, possibly the most popular Google proxy. Startpage now uses Bing<sup><a href="#fn:2" id="fnref:2" role="doc-noteref">note&nbsp;2</a></sup></p></li><li><p><a href="https://search.gmx.com/web">GMX Search</a>, run by a popular German email provider.</p></li><li><p>(discontinued) Runnaroo</p></li><li><p><a href="https://leta.mullvad.net/faq">Mullvad Leta</a></p></li><li><p><a href="https://www.sapo.pt/">SAPO</a> (Portuguese interface, can work with English results)</p></li><li><p><a href="https://www.dsearch.com/">DSearch</a></p></li><li><p><a href="https://www.13tabs.com/">13TABS</a></p></li><li><p><a href="https://zarebin.ir/">Zarebin</a> (Persian, can return English results)</p></li><li><p>A host of other engines using <a href="https://developers.google.com/custom-search/">Programmable Search Engine’s</a> client-side scripts.</p></li></ul>
            </dd>
            					<dt>Bing</dt>
            					<dd>The runner-up. Allows submitting pages and sitemaps for crawling without login using <a href="https://www.indexnow.org/">the IndexNow API</a>, sharing IndexNow page submissions with Yandex and Seznam. Its index powers many other engines:
            <ul><li>Yahoo (and its sibling engine, One­Search)</li><li>DuckDuck­Go<sup><a href="#fn:3" id="fnref:3" role="doc-noteref">note&nbsp;3</a></sup></li><li>AOL</li><li>Qwant (partial)<sup><a href="#fn:4" id="fnref:4" role="doc-noteref">note&nbsp;4</a></sup></li><li>Ecosia</li><li>Ekoru</li><li>Privado</li><li>Findx</li><li>Disconnect Search<sup><a href="#fn:5" id="fnref:5" role="doc-noteref">note&nbsp;5</a></sup></li><li>PrivacyWall</li><li>Lilo</li><li>Search­Scene</li><li>Peekier</li><li>Oscobo</li><li>Million Short</li><li>Yippy search<sup><a href="#fn:6" id="fnref:6" role="doc-noteref">note&nbsp;6</a></sup></li><li>Lycos</li><li>Givero</li><li>Swisscows</li><li>Fireball</li><li>Netzzappen</li><li>You.com<sup><a href="#fn:7" id="fnref:7" role="doc-noteref">note&nbsp;7</a></sup></li><li>Partially powers MetaGer by default; this can be turned off</li><li>At this point, I mostly stopped adding Bing-<wbr>based search engines. There are just too many.</li></ul>
            </dd>
            					<dt>Yandex</dt>
            					<dd>Originally a Russian search engine, it now has an English version. Some Russian results bleed into its English site. It allows submitting pages and sitemaps for crawling using the IndexNow API, sharing IndexNow page submissions with Bing and Seznam. Powers:
            <ul><li>Epic Search (went paid-only as of June 2021)</li><li>Occasionally powers DuckDuck­Go’s link results instead of Bing <ins cite="https://energycommerce.house.gov/committee-activity/hearings/hearing-on-holding-big-tech-accountable-legislation-to-protect-online">(update: DuckDuckGo has “paused” its partnership with Yandex, confirmed in <cite><a href="https://energycommerce.house.gov/committee-activity/hearings/hearing-on-holding-big-tech-accountable-legislation-to-protect-online">Hearing on “Holding Big Tech Accountable: Legislation to Protect Online Users”</a></cite></ins></li><li>Petal, for Russian users only.</li></ul>
            </dd>
            					<dt>
            						<a href="https://www.mojeek.com/">Mojeek</a>
            					</dt>
            					<dd>Seems privacy-oriented with a large index containing billions of pages. Quality isn’t at GBY’s level, but it’s not bad either. If I had to use Mojeek as my default general search engine, I’d live. Partially powers <a href="https://www.etools.ch/">eTools.ch</a>. At this moment, <em>I think that Mojeek is the best alternative to GBY</em> for general search.</dd>
            				</dl>
            				<p>Google, Bing, and Yandex support structured data such as microformats1, microdata, RDFa, Open Graph markup, and JSON-LD. Yandex’s support for microformats1 is limited; for instance, it can parse <code>h-card</code> metadata for organizations but not people. Open Graph and Schema.org are the only supported vocabularies I’m aware of. Mojeek is evaluating structured data; it’s interested in Open Graph and Schema.org vocabularies.</p>
            				<h3 id="smaller-indexes-or-less-relevant-results" tabindex="-1">Smaller indexes or less relevant results</h3>
            				<p>These engines pass most of the tests listed in the “methodology” section. All of them seem relatively privacy-friendly. I wouldn’t recommend using these engines to find specific answers; they’re better for learning about a topic by finding interesting pages related to a set of keywords.</p>
            				<dl>
            					<dt>
            						<a href="https://trystract.com/">Stract</a>
            					</dt>
            					<dd><strong>My favorite generalist engine on this page.</strong> Stract supports advanced ranking customization by allowing users ti import “optics” files, like a better version of Brave’s “goggles” feature. <a href="https://github.com/StractOrg/stract">Stract is fully open-source</a>, with code released under an AGPL-3.0 license. The index is isn’t massive but it’s big enough to be a useful supplement to more major engines. Stract started with the Common Crawl index, but now uses its own crawler. Plans to add contextual ads and a subscription option for ad-free search. Discovered in my access logs.</dd>
            					<dt>
            						<a href="https://rightdao.com">Right Dao</a>
            					</dt>
            					<dd>Very fast, good results. Passes the tests fairly well. It plans on including query-based ads if/when its user base grows.<sup><a href="#fn:8" id="fnref:8" role="doc-noteref">note&nbsp;8</a></sup> For the past few months, its index seems to have focused more on large, established sites rather than smaller, independent ones. It seems to be a bit lacking in more recent pages.</dd>
            					<dt>
            						<a href="https://www.alexandria.org/">Alexandria</a>
            					</dt>
            					<dd>A pretty new “non-profit, ad free” engine, with <a href="https://github.com/alexandria-org/alexandria">freely-licensed code</a>. Surprisingly good at finding recent pages. Its index is built from the Common Crawl; it isn’t as big as Gigablast or Right Dao but its ranking is great.</dd>
            					<dt>
            						<a href="https://yep.com/">Yep</a>
            					</dt>
            					<dd>An ambitious engine from Ahrefs, an SEO/backlink-finder company, that “shares ad profit with creators and protects your privacy”. Most engines show results that include keywords from or related to the query; Yep also shows results linked by pages containing the query. In other words, not all results contain relevant keywords. This makes it excellent for less precise searches and discovery of “related sites”, especially with its index of <em>hundreds of billions of pages.</em> It’s far worse at finding very specific information or recent events for now, but it will probably improve. It was known as “FairSearch” before its official launch.</dd>
            					<dt>
            						<a href="https://sese.yyj.moe/">SeSe Engine</a>
            					</dt>
            					<dd>Although it’s a Chinese engine, its index seems to have a large-enough proportion of English content to fit here. The engine is open-source; see the <a href="https://github.com/RimoChan/sese-engine">SeSe back-end Python code</a> and <a href="https://github.com/YunYouJun/sese-engine-ui">the SeSe-ui Vue-based front-end</a>. It has surprisingly good results for such a low-budget project. Each result is annotated with detailed ranking metadata such as keyword relevance and backlink weight. Discovered in my access logs.</dd>
            					<dt>
            						<a href="https://greppr.org/">greppr</a>
            					</dt>
            					<dd>Its tagline is “Search the Internet with no filters, no tracking, no ads.” At the time of writing, it has over 3 million pages indexed. It’s surprisingly good at finding interesting new results for broad short-tail queries, if you’re willing to scroll far enough down the page. It appears to be good at finding recent pages.</dd>
            				</dl>
            				<p>Yep supports Open Graph and some JSON-LD at the moment. A look through the source code for Alexandria and Gigablast didn’t seem to reveal the use of any structured data. The surprising quality of results from SeSe and Right Dao seems influenced by the crawlers’ high-quality starting location: Wikipedia.</p>
            				<h3 id="smaller-indexes-hit-and-miss" tabindex="-1">Smaller indexes, hit-and-<wbr>miss</h3>
            				<p>These engines fail badly at a few important tests. Otherwise, they seem to work well enough for users who’d like some more serendipity in less-specific searches.</p>
            				<dl>
            					<dt>
            						<a href="https://alpha.infotiger.com/">Infotiger</a>
            					</dt>
            					<dd>My favorite engine in this section. It offers advanced result filtering and sports a somewhat large index. It allows site submission for English and German pages. The fastest-improving engine in this section: I use it often to discover new sites, and look forward to the day it “graduates” to the previous section. <a href="http://infotiger4xywbfq45mvd5drh43jpqeurakg2ya7gqwvjf2bbwnixzqd.onion/">Infotier has a Tor hidden service</a>.</dd>
            					<dt>
            						<a href="http://www.seekport.com/">seekport</a>
            					</dt>
            					<dd>The interface is in German but it supports searching in English just fine. The default language is selected by your locale. It’s really good considering its small index; it hasn’t heard of less common terms. but it’s able to find relevant results in other tests. It’s the second-fastest-improving engines in this section.</dd>
            					<dt>
            						<a href="https://www.exalead.com/search/">Exalead</a>
            					</dt>
            					<dd>Slow, quality is hit-and-miss. Its indexer claims to crawl the DMOZ directory, which has since shut down and been replaced by the <a href="https://curlie.org">Curlie</a> directory. No relevant results for “Oppenheimer” and some other history-related queries. Allows submitting individual URLs for indexing, but requires solving a Google reCAPTCHA and entering an email address.</dd>
            					<dt>
            						<a href="https://www.exactseek.com/">ExactSeek</a>
            					</dt>
            					<dd>Small index, disproportionately dominated by big sites. Failed multiple tests. Allows submitting individual URLs for crawling, but requires entering an email address and receiving a newsletter. Webmaster tools seem to heavily push for paid <abbr title="search-engine optimization">SEO</abbr> options. It also powers SitesOnDisplay and <a href="https://www.blog-search.com">Blog-<wbr>search.com</a>.</dd>
            					<dt>
            						<a href="https://burf.co/">Burf.co</a>
            					</dt>
            					<dd>Very small index, but seems fine at ranking more relevant results higher. Allows site submission without any extra steps.</dd>
            					<dt>
            						<a href="https://siik.co/">Siik</a>
            					</dt>
            					<dd>Lacks contact info, and the ToS and Privacy Policy links are dead. Seems to have PHP errors in the backend for some of its instant-answer widgets. If you scroll past all that, it does have web results powered by what seems to be its own index. These results do tend to be somewhat relevant, but the index seems too small for more specific queries.</dd>
            					<dt>
            						<a href="https://www.chatnoir.eu/">ChatNoir</a>
            					</dt>
            					<dd>An experimental engine by researchers that uses the <a href="https://commoncrawl.org/">Common Crawl</a> index. The engine is <a href="https://github.com/chatnoir-eu">open source</a>. See the <a href="https://groups.google.com/g/common-crawl/c/3o2dOHpeRxo/m/H2Osqz9dAAAJ">announcement</a> on the Common Crawl mailing list (Google Groups).</dd>
            					<dt>
            						<a href="http://www.secretsearchenginelabs.com/">Secret Search Engine Labs</a>
            					</dt>
            					<dd>Very small index with very little SEO spam; it toes the line between a “search engine” and a “surf engine”. It’s best for reading about broad topics that would otherwise be dominated by SEO spam, thanks to its <a href="http://www.secretsearchenginelabs.com/tech/cashrank.php">CashRank algorithm</a>. Allows site submission.</dd>
            					<dt>
            						<a href="https://www.gabanza.com/">Gabanza</a>
            					</dt>
            					<dd>A search engine from a hosting company. I found few details abou the search engine itself, and the index was small, but it was suitable for discovering new pages related to short broad queries.</dd>
            					<dt>
            						<a href="https://jambot.com/">Jambo</a>
            					</dt>
            					<dd>Docs, blog posts, etc. have not been updated since around 2006 but the engine continues to crawl and index new pages. Discovered in my access logs. Has a bias towards older content.</dd>
            				</dl>
            				<h3 id="fledgling-engines" tabindex="-1">Fledgling engines</h3>
            				<p>Results from these search engines don’t seem particularly relevant; indexes in this category tend to be small.</p>
            				<dl>
            					<dt>
            						<a href="https://www.yessle.com/">Yessle</a>
            					</dt>
            					<dd>Seems new; allows page submission by pasting a page into the search box. Index is really small but it crawls new sites quickly. Claims to be private.</dd>
            					<dt>
            						<a href="https://search.aibull.io/">Bloopish</a>
            					</dt>
            					<dd>Extremely quick to update its index; site submissions show up in seconds. Unfortunately, its index only contains a few thousand documents (under 100 thousand at the time of writing). It’s growing fast: if you search for a term, it’ll start crawling related pages and grow its index.</dd>
            					<dt>YaCy</dt>
            					<dd>Community-made index; slow. Results are awful/irrelevant, but can be useful for intranet or custom search.</dd>
            					<dt>Scopia</dt>
            					<dd>only seems to be available via the <a href="https://metager.org">MetaGer</a> metasearch engine after turning off Bing and news results. Tiny index, very low-quality.</dd>
            					<dt>
            						<a href="https://www.artadosearch.com/">Artado Search</a>
            					</dt>
            					<dd>Primarily Turkish, but it also seems to support English results. Like Plumb, it uses client-side JS to fetch results from existing engines (Google, Bing, Yahoo, Petal, and others); like MetaGer, it has an option to use its own independent index. Results from its index are almost always empty. Very simple queries (“twitter”, “wikipedia”, “reddit”) give some answers. Supports site submission and crowdsourced instant answers.</dd>
            					<dt>
            						<a href="https://www.activesearchresults.com">Active Search Results</a>
            					</dt>
            					<dd>Very poor quality. Results seem highly biased towards commercial sites.</dd>
            					<dt>
            						<a href="https://www.crawlson.com">Crawlson</a>
            					</dt>
            					<dd>Young, slow. In this category because its index has a cap of 10 URLs per domain. I initially discovered Crawlson in the seirdy.one access logs. This is often down; if the current downtime persists, I’ll add it to the graveyard.</dd>
            					<dt>
            						<a href="https://www.anoox.com/">Anoox</a>
            					</dt>
            					<dd>Results are few and irrelevant; fails to find any results for basic terms. Allows site submission. It’s also a lightweight social network and claims to be powered by its users, letting members vote on listings to alter rankings.</dd>
            					<dt>
            						<a href="https://www.yioop.com">Yioop!</a>
            					</dt>
            					<dd>A FLOSS search engine that boasts a very impressive <a href="https://www.seekquarry.com/">feature-set</a>: it can parse sitemaps, feeds, and a variety of markup formats; it can import pre-curated data in forms such as access logs, Usenet posts, and WARC archives; it also supports feed-based news search. Despite the impressive feature set, Yioop’s results are few and irrelevant due to its small index. It allows submitting sites for crawling. Like Meorca, Yioop has social features such as blogs, wikis, and a chat bot API.</dd>
            					<dt>
            						<a href="https://spyda.dev/">Spyda</a>
            					</dt>
            					<dd><span class="h-cite" itemprop="mentions" itemscope="" itemtype="https://schema.org/BlogPosting">A small engine made by <span itemprop="author" itemscope="" itemtype="https://schema.org/Person" class="h-card vcard p-author"><a itemprop="url" href="https://www.prologic.blog/" class="u-url url"><span itemprop="name" class="p-name fn n"><span itemprop="givenName" class="p-given-name given-name">James</span>&nbsp;<span itemprop="familyName" class="p-family-name family-name">Mills</span></span></a></span>, described in <cite itemprop="name headline" class="p-name"><a class="u-url" itemprop="url" href="https://www.prologic.blog/2021/02/14/so-im-a.html">So I'm a Knucklehead eh?</a></cite></span>. It’s written in Go; check out its <a href="https://git.mills.io/prologic/spyda">MIT-licensed Spyda source code</a>.</dd>
            					<dt>
            						<a href="https://www.slzii.com/">Slzii.com</a>
            					</dt>
            					<dd>A new web portal with a search engine. Has a tiny index dominated by SEO spam. Discovered in the seirdy.one access logs.</dd>
            				</dl>
            				<h3 id="semi-independent-indexes" tabindex="-1">Semi-independent indexes</h3>
            				<p>Engines in this category fall back to GBY when their own indexes don’t have enough results. As their own indexes grow, some claim that this should happen less often.</p>
            				<dl>
            					<dt>
            						<a href="https://search.brave.com/">Brave Search</a>
            					</dt>
            					<dd>Many tests (including all the tests I listed in the “Methodology” section) resulted results identical to Google, revealed by a side-by-side comparison with Google, Startpage, and a Searx instance with only Google enabled. Brave claims that this is due to how Cliqz (the discontinued engine acquired by Brave) used query logs to build its page models and was optimized to match Google.<sup><a href="#fn:9" id="fnref:9" role="doc-noteref">note&nbsp;9</a></sup> The index is independent, but optimizing against Google resulted in too much similarity for the real benefit of an independent index to show. Furthermore, many queries have Bing results mixed in; users can click an “info” button to see the percentage of results that came from its own index. The independent percentage is typically quite high (often close to 100% independent) but can drop for advanced queries. <ins cite="https://brave.com/search-independence/" datetime="2023-08-15T20:39:00-07:00">Update 2023-08-15: Brave’s Bing contract appears to have expired as of April 2023.</ins>
            <p>I can’t in good conscience recommend using Brave Search, as the company runs cryptocurrency, has <a href="https://brave.com/rewards-update/">held payments to creators without disclosing that creators couldn’t receive rewards</a>, has made dangerously misleading claims about fingerprinting resistance,<sup><a href="#fn:10" id="fnref:10" role="doc-noteref">note&nbsp;10</a></sup> is run by a CEO who <a href="https://arstechnica.com/information-technology/2014/03/new-mozilla-ceo-issues-statement-expresses-sorrow-for-causing-pain/">spent thousands of dollars opposing gay marriage</a>, and <a href="https://www.pcmag.com/news/brave-browser-caught-redirecting-users-through-affiliate-links">has rewritten typed URLs with affiliate links</a>.</p>
            </dd>
            					<dt>
            						<a href="https://plumb.one/">Plumb</a>
            					</dt>
            					<dd>Almost all queries return no results; when this happens, it falls back to Google. It’s fairly transparent about the fallback process, but I’m concerned about <em>how</em> it does this: it loads Google’s Custom Search scripts from <code>cse.google.com</code> onto the page to do a client-side Google search. This can be mitigated by using a browser addon to block <code>cse.google.com</code> from loading any scripts. Plumb claims that this is a temporary measure while its index grows, and they’re planning on getting rid of this. Allows submitting URLs, but requires solving an hCaptcha. This engine is very new; hopefully as it improves, it could graduate from this section. Its Chief Product Officer <a href="https://archive.is/oVAre">previously founded</a> the Gibiru search engine which shares the same affiliates and (for now) the same index; the indexes will diverge with time.</dd>
            					<dt>
            						<a href="https://www.qwant.com">Qwant</a>
            					</dt>
            					<dd>Qwant claims to use its own index, but it still relies on Bing for most results. It seems to be in a position similar to Neeva. Try a side-by-side comparison to see if or how it compares with Bing.</dd>
            					<dt>
            						<a href="https://kagi.com/">Kagi Search</a>
            					</dt>
            					<dd>The most interesting entry in this category, IMO. Like Neeva, it requires an account and limits use without payment. It’s powered by its own Teclis index (Teclis can be used independently; see the <a href="#small-or-non-commercial-web">non-commercial section</a> below), and claims to also use results from Google and Bing. The result seems somewhat unique: I’m able to recognize some results from the Teclis index mixed in with the mainstream ones. In addition to Teclis, Kagi’s other products include the <a href="https://kagi.ai/">Kagi.ai</a> intelligent answer service and the <a href="https://tinygem.org/">TinyGem</a> social bookmarking service, both of which play a role in Kagi.com in the present or future. Unrelatedly: I’m concerned about the company’s biases, as it seems happy to <a href="https://kagifeedback.org/d/2808-reconsider-your-partnership-with-brave">use Brave’s commercial API</a> (allowing blatant homophobia in the comments) and <a href="https://kagifeedback.org/d/865-suicide-results-should-probably-have-a-dont-do-that-widget-like-google/50">allow its results to recommend suicide methods without intervention</a>. I reject the idea that avoiding an option that may seem politically biased is the same as being unbiased if such a decision has real political implications.</dd>
            					<dt>
            						<a href="https://svmetasearch.eu.org/">SVMetaSearch</a>
            					</dt>
            					<dd>A SearxNG metasearch engine that also includes results from its own index. All other sources can be turned off. Like most public Searx/SearxNG instances, reliability is very poor.</dd>
            				</dl>
            				<h2 id="non-generalist-search" tabindex="-1">Non-generalist search</h2>
            				<aside role="none">
            					<a href="#non-generalist-search" aria-labelledby="non-generalist-search-prefix non-generalist-search">
            						<span id="non-generalist-search-prefix">Permalink to section</span>
            					</a>
            				</aside>
            				<p>These indexing search engines don’t have a Google-like “ask me anything” endgame; they’re trying to do something different. You aren’t supposed to use these engines the same way you use GBY.</p>
            				<h3 id="small-or-non-commercial-web" tabindex="-1">Small or non-commercial Web</h3>
            				<dl>
            					<dt>
            						<a href="https://search.marginalia.nu/">Marginalia Search</a>
            					</dt>
            					<dd><em>My favorite entry on this page</em>. It has its own crawler but is strongly biased towards non-commercial, personal, and/or minimal sites. It’s a great response to the increasingly SEO-spam-filled SERPs of GBY. Partially powers Teclis, which in turn partially powers Kagi. <ins cite="https://memex.marginalia.nu/log/58-marginalia-open-source.gmi" datetime="2022-05-28T14:09:00-07:00">Update 2022-05-28: <a href="https://memex.marginalia.nu/log/58-marginalia-open-source.gmi">Marginalia.nu is now open source.</a></ins></dd>
            					<dt>
            						<a href="https://ichi.do/">Ichido</a>
            					</dt>
            					<dd>An engine that just rolled out its own independent index, with a lot of careful thought put into its ranking algorithm. Like Marginalia, it’s biased towards the non-commercial Web: it downranks ads, CAPTCHAs, trackers, SEO, and obfuscation. <a href="https://blog.ichi.do/post/2023/08/20/a-new-ichido/">More info about Ichido is in a blog post</a>.</dd>
            					<dt>
            						<a href="http://teclis.com/">Teclis</a>
            					</dt>
            					<dd>A project by the creator of Kagi search. Uses its own crawler that measures content blocked by uBlock Origin, and extracts content with the open-source article scrapers Trafilatura and Readability.js. This is quite an interesting approach: tracking blocked elements discourages tracking and advertising; using Trafilatura and Readability.js encourages the use of semantic HTML and Semantic Web standards such as <a href="https://microformats.org/">microformats</a>, <a href="https://html.spec.whatwg.org/multipage/microdata.html">microdata</a>, and <a href="https://www.w3.org/TR/rdfa-primer/">RDFa</a>. It claims to also use some results from Marginalia. <a href="https://kagifeedback.org/d/1838-teclis-is-broken/2">The Web interface has been shut down</a>, but its standalone API is still available for Kagi customers.</dd>
            					<dt>
            						<a href="https://clew.se/">Clew</a>
            					</dt>
            					<dd>a FOSS new engine with a small index of several thousand pages. It focuses on independent content and downranks ads and trackers; there seems to be a real focus on quality over quantity, which makes it excellent for short-tail searches (especially around technical concepts). Ranking is more egalitarian than other engines, making it better for discovery and surfing than research. It’s designed to be small and lightweight, with a compact index. Discovered in the seirdy.one access logs.</dd>
            				</dl>
            				<h3 id="site-finders" tabindex="-1">Site finders</h3>
            				<p>These engines try to find a website, typically at the domain-name level. They don’t focus on capturing particular pages within websites.</p>
            				<dl>
            					<dt>
            						<a href="https://kozmonavt.ml/">Kozmonavt</a>
            					</dt>
            					<dd>The best in this category. Has a small but growing index of over 8 million sites. If I want to find the website for a certain project, Kozmonavt works well (provided its index has crawled said website). It works poorly for learning things and finding general information. I cannot recommend it for anything serious since it lacks contact information, a privacy policy, or any other information about the org/people who made it. Discovered in the seirdy.one access logs.</dd>
            					<dt>
            						<a href="http://www.search.tl/">search.tl</a>
            					</dt>
            					<dd>Generalist search for one <abbr title="top-level domain">TLD</abbr> at a time (defaults to .com). I’m not sure why you’d want to always limit your searches to a single TLD, but now you can.<sup><a href="#fn:11" id="fnref:11" role="doc-noteref">note&nbsp;11</a></sup> There isn’t any visible UI for changing the TLD for available results; you need to add/change the <code>tld</code> URL parameter. For example, to search .org sites, append <code>&amp;tld=org</code> to the URL. It seems to be connected to <a href="http://www.amidalla.de/">Amidalla</a>. Amidalla allows users to manually add URLs to its index and directory; I have yet to see if doing so impacts search.tl results.</dd>
            					<dt>
            						<a href="https://search.thunderstone.com/">Thunderstone</a>
            					</dt>
            					<dd>A combined website catalog and search engine that focuses on categorization. Its <a href="https://search.thunderstone.com/texis/websearch19/about.html">about page</a> claims: <q cite="https://search.thunderstone.com/texis/websearch19/about.html">We continuously survey all primary COM, NET, and ORG web-servers and distill their contents to produce this database. This is an index of <em>sites</em> not pages. It is very good at finding companies and organizations by purpose, product, subject matter, or location. If you’re trying to finding things like <em>‘BillyBob’s personal beer can page on AOL’</em>, try Yahoo or Dogpile.</q> This seems to be the polar opposite of the engines in the <a href="#small-or-non-commercial-web">“small or non-commercial Web” category</a>.</dd>
            					<dt>
            						<a href="https://www.sengine.info/">sengine.info</a>
            					</dt>
            					<dd>Only shows domains, not individual pages. Developed by netEstate GmbH, which specializes in content extraction for inprints and job ads. Also has a German-only version available. Discovered in my access logs.</dd>
            					<dt>
            						<a href="https://www.gnomit.com/">Gnomit</a>
            					</dt>
            					<dd>Allows single-keyword queries and returns sites that seem to cover a related topic. I actually kind of enjoy using it; results are old (typically from 2009) and a bit random, but make for a nice way to discover something new. For instance, searching for “IRC” helped me discover new IRC networks I’d never heard of.</dd>
            				</dl>
            				<h3 id="other" tabindex="-1">Other</h3>
            				<dl>
            					<dt>
            						<a href="https://highbrow.se/">High Browse</a>
            					</dt>
            					<dd>Uses a non-traditional ranking algorithm which does an excellent job of introducing non-SEO-optimized serendipity into search results. One of my favorite “surf-engines”, as opposed to traditional “search-engines”.</dd>
            					<dt>
            						<a href="https://www.keybot.com/">Keybot</a>
            					</dt>
            					<dd>A must-have for anyone who does translation work. It crawls the web looking for multilingual websites. Translators who are unsure about how to translate a given word or phrase can see its usage in two given languages, to learn from other human translators. My parents are fluent English speakers but sometimes struggle to express a given Hindi idiom in English; something like this could be useful to them, since machine translation isn’t nuanced enough for every situation. Part of the <a href="https://www.ttn.ch/">TTN Translation Network</a>. Discovered in my access logs.</dd>
            					<dt>Quor</dt>
            					<dd>Seems to mainly index large news sites. Site is down as of June 2021; originally available at www dot quor dot com.</dd>
            					<dt>
            						<a href="https://www.semanticscholar.org/">Semantic Scholar</a>
            					</dt>
            					<dd>A search engine by the Allen Institute for AI focused on academic PDFs, with a couple hundred million papers indexed. Discovered in my access logs.</dd>
            					<dt>
            						<a href="https://bonzamate.com.au/">Bonzamate</a>
            					</dt>
            					<dd>A search engine specifically for Australian websites. Boyter wrote <a href="https://boyter.org/posts/abusing-aws-to-make-a-search-engine/">an interesting blog post about Bonzamate</a>.</dd>
            					<dt>
            						<a href="https://searchcode.com/">searchcode</a>
            					</dt>
            					<dd>A code-search engine by the developer of Bonzamate. Searches a hand-picked list of code forges for source code, supporting many search operators.</dd>
            					<dt>
            						<a href="https://search.lixialabs.com/">Lixia Labs Search</a>
            					</dt>
            					<dd>A new engine that focuses on indexing technical websites and blogs, with a minimal JavaScript-free front-end. Discovered in my access logs. Surprisingly good results for broad technical keyword queries.</dd>
            				</dl>
            				<h2 id="other-languages" tabindex="-1">Other languages</h2>
            				<aside role="none">
            					<a href="#other-languages" aria-labelledby="other-languages-prefix other-languages">
            						<span id="other-languages-prefix">Permalink to section</span>
            					</a>
            				</aside>
            				<p>I’m unable to evaluate these engines properly since I don’t speak the necessary languages. English searches on these are a hit-or-miss. I might have made a few mistakes in this category.</p>
            				<h3 id="big-indexes" tabindex="-1">Big indexes</h3>
            				<ul>
            					<li>
            						<p>Baidu: Chinese. Very large index; it’s a major engine alongside GBY. Offers webmaster tools for site submission.</p>
            					</li>
            					<li>
            						<p>Qihoo 360: Chinese. I’m not sure how independent this one is.</p>
            					</li>
            					<li>
            						<p>Toutiao: Chinese. Not sure how independent this one is either. Its index appears limited outside of its own content distribution platform.</p>
            					</li>
            					<li>
            						<p>Sogou: Chinese</p>
            					</li>
            					<li>
            						<p>Yisou: Chinese, by Yahoo. Appears defunct.</p>
            					</li>
            					<li>
            						<p><a href="https://search.naver.com">Naver</a>: Korean. Allows submitting sitemaps and feeds. Discovered via some Searx metasearch instances.</p>
            					</li>
            					<li>
            						<p><a href="https://www.daum.net/">Daum</a>: Korean. Also unsure about this one’s independence.</p>
            					</li>
            					<li>
            						<p><a href="https://www.seznam.cz/">Seznam</a>: Czech, seems relatively privacy-friendly. Discovered in the seirdy.one access logs. It allows site submission with webmaster tools. <a href="https://blog.seznam.cz/2022/03/mate-novy-obsah-dejte-o-nem-hned-vsem-vedet-pomoci-indexnow/">Seznam supports IndexNow</a>; it shares IndexNow-submitted pages with Bing and Yandex.</p>
            					</li>
            					<li>
            						<p><a href="https://coccoc.com/search">Cốc Cốc</a>: Vietnamese</p>
            					</li>
            					<li>
            						<p><a href="https://go.mail.ru/">go.mail.ru</a>: Russian</p>
            					</li>
            					<li>
            						<p><a href="https://letsearch.ru/">LetSearch.ru</a>: Russian. <a href="https://letsearch.ru/bots">Allows URL submission</a></p>
            					</li>
            				</ul>
            				<h3 id="smaller-indexes" tabindex="-1">Smaller indexes</h3>
            				<ul>
            					<li>
            						<p><a href="https://www.alibw.com/">ALibw.com</a>: Chinese, found in my access logs.</p>
            					</li>
            					<li>
            						<p><a href="https://www.vuhuv.com.tr/">Vuhuv</a>: Turkish. <a href="https://tr.vuhuv.com/">alt domain</a></p>
            					</li>
            					<li>
            						<p><a href="https://www.search.ch">search.ch</a>: Regional search engine for Switzerland; users can restrict searches to their local regions.</p>
            					</li>
            					<li>
            						<p><a href="https://www.fastbot.de/">fastbot</a>: German</p>
            					</li>
            					<li>
            						<p><a href="https://www.moose.at">Moose.at</a>: German (Austria-based)</p>
            					</li>
            					<li>
            						<p><a href="https://solofield.net">SOLOFIELD</a>: Japanese</p>
            					</li>
            					<li>
            						<p><a href="http://kaz.kz">kaz.kz</a>: Kazakh and Russian, with a focus on “Kazakhstan’s segment of the Internet”</p>
            					</li>
            				</ul>
            				<h2 id="almost-qualified" tabindex="-1">Almost qualified</h2>
            				<aside role="none">
            					<a href="#almost-qualified" aria-labelledby="almost-qualified-prefix almost-qualified">
            						<span id="almost-qualified-prefix">Permalink to section</span>
            					</a>
            				</aside>
            				<p>These engines come close enough to passing my inclusion criteria that I felt I had to mention them. They all display original organic results that you can’t find on other engines, and maintain their own indexes. Unfortunately, they don’t quite pass.</p>
            				<dl>
            					<dt>
            						<a href="https://wiby.me">wiby.me</a>
            					</dt>
            					<dt>
            						<a href="https://wiby.org">wiby.org</a>
            					</dt>
            					<dd>I love this one. It focuses on smaller independent sites that capture the spirit of the “early” web. It’s more focused on “discovering” new interesting pages that match a set of keywords than finding a specific resources. I like to think of Wiby as an engine for surfing, not searching. Runnaroo occasionally featured a hit from Wiby (Runnaroo has since shut down). If you have a small site or blog that isn’t very “commercial”, consider submitting it to the index. Does not qualify because it seems to be powered only by user-submitted sites; it doesn’t try to “crawl the Web”.</dd>
            					<dt>
            						<a href="https://mwmbl.org/">Mwmbl</a>
            					</dt>
            					<dd>Like YaCy, it’s an open-source engine whose crawling is community-driven. Users can install a Firefox addon to crawl pages in its backlog. Unfortunately, it doesn’t qualify because it only crawls pages linked by hand-picked sites (e.g. Wikipedia, GitHub, domains that rank well on Hacker News). The crawl-depth is “1”, so it doesn’t crawl the whole Web (yet).</dd>
            					<dt>
            						<a href="https://searchmysite.net">Search My Site</a>
            					</dt>
            					<dd>Similar to Marginalia and Teclis, but only indexes user-submitted personal and independent sites. It optionally supports IndieAuth. Its API powers this site’s search results; try it out using the search bar at the bottom of this page. Does not qualify because it’s limited to user-submitted and/or hand-picked sites.</dd>
            					<dt>
            						<a href="https://blogsurf.io/">Blog Surf</a>
            					</dt>
            					<dd>A search engine for blogs with RSS/Atom feeds. Does not qualify because all blogs submitted to the index require manual review, but it seems interesting. Its “MarketRank” algorithm seems to give it a bias towards sites popular on “Hacker” “News”.</dd>
            					<dt>
            						<a href="https://kukei.eu/">Kukei.eu</a>
            					</dt>
            					<dd>A curated search engine for web developers, which crawls <a href="https://github.com/Kukei-eu/spider/blob/914b8dfffc10cb3a948561aef2bf86937d3a0b2e/index-sources.js">a hand-picked list of sites</a>. As it does not index the whole Web, it doesn’t qualify. I still find it interesting.</dd>
            				</dl>
            				<h2 id="misc" tabindex="-1">Misc</h2>
            				<aside role="none">
            					<a href="#misc" aria-labelledby="misc-prefix misc">
            						<span id="misc-prefix">Permalink to section</span>
            					</a>
            				</aside>
            				<dl>
            					<dt>Ask.com</dt>
            					<dd>The site is back. They claim to outsource search results. The results seem similar to Google, Bing, and Yandex; however, I can’t pinpoint exactly where their results are coming from. Also, several sites from the “ask.com network” such as directhit.com, info.com, and kensaq.com have uniqe-looking results.</dd>
            					<dt>
            						<a href="https://infinitysearch.co">Infinity Search</a>
            					</dt>
            					<dd>Partially evaluated. Young, small index. It recently split into a paid offering with the main index and <a href="https://infinitydecentralized.com/">Infinity Decentralized</a>, the latter of which allows users to select from community-hosted crawlers. I managed to try it out before it became a paid offering, and it seemed decent; however, I wasn’t able to run the tests listed in the “Methodology” section. Allows submitting URLs and sitemaps into a text box, no other work required.</dd>
            				</dl>
            				<h2 id="search-engines-without-a-web-interface" tabindex="-1">Search engines without a web interface</h2>
            				<aside role="none">
            					<a href="#search-engines-without-a-web-interface" aria-labelledby="search-engines-without-a-web-interface-prefix search-engines-without-a-web-interface">
            						<span id="search-engines-without-a-web-interface-prefix">Permalink to section</span>
            					</a>
            				</aside>
            				<p>Some search engines are integrated into other appliances, but don’t have a web portal.</p>
            				<ul>
            					<li>
            						<p>Apple’s search engine is usable in the form of “Siri Suggested Websites”. Its index is built from the Applebot web crawler. If Apple already has working search engine, it’s not much of a stretch to say that they’ll make a web interface for it someday.</p>
            					</li>
            					<li>
            						<p>Amazon bought Alexa Internet (a web traffic analysis company, at the time unrelated to the Amazon Alexa virtual assistant) and discontinued its website ranking product. Amazon still runs the relevant crawlers, and also have <a href="https://developer.amazon.com/support/amazonbot">a bot called “Amazonbot”</a>. While Applebot powers the Siri personal assistant, Amazonbot powers the Alexa personal assistant <q cite="https://developer.amazon.com/support/amazonbot">to answer even more questions for customers</q>. Crawling the web to answer questions is the basis of a search engine.</p>
            					</li>
            				</ul>
            				<h2 id="graveyard" tabindex="-1">Graveyard</h2>
            				<aside role="none">
            					<a href="#graveyard" aria-labelledby="graveyard-prefix graveyard">
            						<span id="graveyard-prefix">Permalink to section</span>
            					</a>
            				</aside>
            				<p>These engines were originally included in the article, but have since been discontinued.</p>
            				<dl>
            					<dt>
            						<a href="https://web.archive.org/web/20230118225122/https://www.petalsearch.com/">Petal Search</a>
            					</dt>
            					<dd>A search engine by Huawei that recently switched from searching for Android apps to general search in order to reduce dependence on Western search providers. Despite its surprisingly good results, I wouldn’t recommend it due to privacy concerns: its privacy policy describes advanced fingerprinting metrics, and it doesn’t work without JavaScript. Requires an account to submit sites. I discovered this via my access logs. Be aware that in some jurisdictions, it doesn’t use its own index: in Russia and some EU regions it uses Yandex and Qwant, respectively. Inaccessible as of June 2023.</dd>
            					<dt>
            						<a href="https://web.archive.org/web/20230528051432/https://neeva.com/blog/may-announcement">Neeva</a>
            					</dt>
            					<dd>Formerly in <a href="#semi-independent-indexes">the “semi-independent” section</a>. Combined Bing results with results from its own index. Bing normally isn’t okay with this, but Neeva was one of few exceptions. Results were mostly identical to Bing, but original links not found by Bing frequently popped up. Long-tail and esoteric queries were less likely to feature original results. Required signing up with an email address or OAuth to use, and offered a paid tier with additional benefits. Acquired by Snowflake and announced its shut-down in May 2023.</dd>
            					<dt>
            						<a href="https://gigablast.com/">Gigablast</a>
            					</dt>
            					<dd>It’s been around for a while and also sports a classic web directory. Searches are a bit slow, and it charges to submit sites for crawling. It powers <a href="https://private.sh">Private.sh</a>. Gigablast was tied with Right Dao for quality. Shut down mid-2023.</dd>
            					<dt>
            						<a href="https://xangis.com/the-wbsrch-experiment/">wbsrch</a>
            					</dt>
            					<dd>In addition to its generalist search, it also had many other utilities related to domain name statistics. Failed multiple tests. Its index was a bit dated; it had an old backlog of sites it hadn’t finished indexing. It also had several dedicated per-language indexes.</dd>
            					<dt>
            						<a href="https://web.archive.org/web/20211226043304/https://www.gowiki.com/">Gowiki</a>
            					</dt>
            					<dd>Very young, small index, but showed promise. I discovered this in the seirdy.one access logs. It was only available in the US. Seems down as of early 2022.</dd>
            					<dt>
            						<a href="https://web.archive.org/web/20220429143153/https://www.meorca.com/search/">Meorca</a>
            					</dt>
            					<dd>A UK-based search engine that claims not to “index pornography or illegal content websites”. It also features an optional social network (“blog”). Discovered in the seirdy.one access logs.</dd>
            					<dt>
            						<a href="https://web.archive.org/web/20220624172257/https://ninfex.com/">Ninfex</a>
            					</dt>
            					<dd>A “people-powered” search engine that combines aspects of link aggregators and search. It lets users vote on submissions and it also displays links to forums about submissions.</dd>
            					<dt>
            						<a href="https://github.com/isovector/marlo">Marlo</a>
            					</dt>
            					<dd>Another FLOSS engine: Marlo is written in Haskell. Has a small index that’s good enough for surfing broad topics, but not good enough for specific research. Originally available at <code>marlo.sandymaguire.me</code>.</dd>
            					<dt>websearchengine.org</dt>
            					<dt>tuxdex.com</dt>
            					<dd>Both were run by the same people, powered by their inetdex.com index. Searches are fast, but crawls are a bit shallow. Claims to have an index of 10 million domains, and not to use cookies. The pages are currently down and the domains re-direct to porn sites; I’m not aware of any official notice.</dd>
            					<dt>
            						<a href="https://web.archive.org/web/20230810032916/https://entfer.com/">Entfer</a>
            					</dt>
            					<dd>a newcomer that let registered users upvote/downvote search results to customize ranking. Didn’t offer much information about who made it. Its index was small, but it did seem to return results related to the query.</dd>
            				</dl>
            				<p>Dead engines I don’t have an extended description for:</p>
            				<ul>
            					<li><a href="https://www.parsijoo.ir/">Parsijoo</a>: Persian search engine.</li>
            				</ul>
            				<h2 id="exclusions" tabindex="-1">Exclusions</h2>
            				<aside role="none">
            					<a href="#exclusions" aria-labelledby="exclusions-prefix exclusions">
            						<span id="exclusions-prefix">Permalink to section</span>
            					</a>
            				</aside>
            				<p>Two engines were excluded from this list for having a far-right focus.</p>
            				<p>One engine was excluded because it seems to be built using cryptocurrency in a way I’d rather not support.</p>
            				<p>Some fascinating little engines seem like hobbyist proofs-of-concept. I decided not to include them in this list, but watch them with interest to see if they can become something viable.</p>
            				<h2 id="rationale" tabindex="-1">Rationale</h2>
            				<aside role="none">
            					<a href="#rationale" aria-labelledby="rationale-prefix rationale">
            						<span id="rationale-prefix">Permalink to section</span>
            					</a>
            				</aside>
            				<p>Why bother using non-mainstream search engines?</p>
            				<h3 id="conflicts-of-interest" tabindex="-1">Conflicts of interest</h3>
            				<p>Google, Microsoft (the company behind Bing), and Yandex aren’t just search engine companies; they’re content and ad companies as well. For example, Google hosts video content on YouTube and Microsoft hosts social media content on LinkedIn. This gives these companies a powerful incentive to prioritize their own content. They are able to do so even if they claim that they treat their own content the same as any other: since they have complete access to their search engines’ inner workings, they can tailor their content pages to better fit their algorithms and tailor their algorithms to work well on their own content. They can also index their own content without limitations but throttle indexing for other crawlers.<sup><a href="#fn:12" id="fnref:12" role="doc-noteref">note&nbsp;12</a></sup></p>
            				<p>One way to avoid this conflict of interest is to <em>use search engines that aren’t linked to major content providers;</em> i.e., use engines with their own independent indexes.</p>
            				<h3 id="information-diversity" tabindex="-1">Information diversity</h3>
            				<p>There’s also a practical, less-ideological reason to try other engines: different providers have different results. Websites that are hard to find on one search engine might be easy to find on another, so using more indexes and ranking algorithms results in access to more content.</p>
            				<p>No search engine is truly unbiased. Most engines’ ranking algorithms incorporate a method similar to <a href="https://en.wikipedia.org/wiki/PageRank">PageRank</a>, which biases them towards sites with many backlinks. Search engines have to deal with unwanted results occupying the confusing overlap between SEO spam, shock content, and duplicate content. When this content’s manipulation of ranking algos causes it to rank high, engines have to address it through manual action or algorithm refinement. Choosing to address it through either option, or choosing to leave it there for popular queries after receiving user reports, reflects bias. The best solution is to mix different ranking algorithms and indexes instead of using one engine for everything.</p>
            				<h2 id="methodology" tabindex="-1">Method­ology</h2>
            				<aside role="none">
            					<a href="#methodology" aria-labelledby="methodology-prefix methodology">
            						<span id="methodology-prefix">Permalink to section</span>
            					</a>
            				</aside>
            				<h3 id="discovery" tabindex="-1">Discovery</h3>
            				<p>I find new engines by:</p>
            				<ul>
            					<li>Monitoring certain web directories for changes in their search engine listings.</li>
            					<li>Checking other curated lists of “good/bad bots” to spot search engines.</li>
            					<li>Using search engines to discover search engines: searching for the names of less-popular engines often pulls up similar lists.</li>
            					<li>Receiving suggestions from readers</li>
            					<li>Compiling a list of regular expressions for user-agent strings I’m familiar with. Before I delete my server access logs, I extract user-agents that don’t match that list along with the pages they request.</li>
            					<li>Checking the Searx and Searxng projects for new integrations.</li>
            				</ul>
            				<h3 id="criteria-for-inclusion" tabindex="-1">Criteria for inclusion</h3>
            				<p>Engines in this list should have their own indexes powered by web crawlers. Original results should not be limited to a set of websites hand-picked by the engine creators; indexes should be built from sites from across the Web. An engine should discover new interesting places around the Web.</p>
            				<p>Here’s an oversimplified example to illustrate what I’m looking for: imagine somone self-hosts their own personal or interest-specific website and happens to get some recognition. Could they get <em>automatically</em> discovered by your crawler, indexed, and included in the first page of results for a certain query?</p>
            				<p>I’m willing to make two exceptions:</p>
            				<ol>
            					<li>Engines in the “semi-independent” section may mix results that do meet the aforementioned criteria with results that do not.</li>
            					<li>Engines in the “almost qualified” section may use indexes primarily made of user-submitted or hand-picked sites, rather than focusing primarily on sites discovered organically through crawling.</li>
            				</ol>
            				<p>The reason the second exception exists is that while user submissions don’t represent automatic crawling, they do at least inform the engine of new interesting websites that it had not previously discovered; these websites can then be shown to other users. That’s fundamentally what an alternative web index needs to achieve.</p>
            				<p>I’m not usually willing to budge on my “no hand-picked websites” rule. Hand-picked sites will be ignored, whether your engine fetches content through their APIs or crawls and scrapes their content. It’s fine to use hand-picked websites as starting points for your crawler (Wikipedia is a popular option).</p>
            				<p>I only consider search engines that focus on link results for webpages. Image search engines are out of scope, though I <em>might</em> consider some other engines for non-generalist search (e.g., Semantic Scholar finds PDFs rather than webpages).</p>
            				<h3 id="evaluation" tabindex="-1">Evaluation</h3>
            				<p>I focused almost entirely on “organic results” (the classic link results), and didn’t focus too much on (often glaring) privacy issues, “enhanced” or “instant” results (e.g. Wikipedia sidebars, related searches, Stack Exchange answers), or other elements.</p>
            				<p>I compared results for esoteric queries side-by-side; if the first 20 results were (nearly) identical to another engine’s results (though perhaps in a slightly different order), they were likely sourced externally and not from an independent index.</p>
            				<p>I tried to pick queries that should have a good number of results and show variance between search engines. An incomplete selection of queries I tested:</p>
            				<ul>
            					<li>
            						<p>“vim”, “emacs”, “neovim”, and “nvimrc”: Search engines with relevant results for “nvimrc” typically have a big index. Finding relevant results for the text editors “vim” and “emacs” instead of other topics that share the name is a challenging task.</p>
            					</li>
            					<li>
            						<p>“vim cleaner”: should return results related to a <a href="https://en.wikipedia.org/wiki/Vim_%28cleaning_product%29">line of cleaning products</a> rather than the Correct Text Editor.</p>
            					</li>
            					<li>
            						<p>“Seirdy”: My site is relatively low-traffic, but my nickname is pretty unique and visible on several of the highest-traffic sites out there.</p>
            					</li>
            					<li>
            						<p>“Project London”: a small movie made with volunteers and <abbr title="Free, Libre, Open-Source Software">FLOSS</abbr> without much advertising. If links related to the movie show up, the engine’s really good.</p>
            					</li>
            					<li>
            						<p>“oppenheimer”: a name that could refer to many things. Without context, it should refer to the physicist who worked on the atomic bomb in Los Alamos. Other historical queries: “magna carta” (intermediate), “the prince” (very hard).</p>
            					</li>
            				</ul>
            				<p>Some less-mainstream engines have noticed this article, which is great! I’ve had excellent discussions with people who work on several of these engines. Unfortunately, this article’s visibility also incentivizes some engines to optimize specifically for any methodology I describe. I’ve addressed this by keeping a long list of test queries to myself. The simple queries above are a decent starting point for simple quick evaluations, but I also test for common search operators, keyword length, and types of domain-specific jargon. I also use queries designed to pull up specific pages with varying levels of popularity and recency to gauge the size, scope, and growth of an index.</p>
            				<p>Professional critics often work anonymously because personalization can damage the integrity of their reviews. For similar reasons, I attempt to try each engine anonymously at least once by using a VPN and/or my standard anonymous setup: an amnesiac Whonix VM with the Tor Browser. I also often test using a fresh profile when travelling, or via a Searx instance if it supports a given engine. When avoiding personalization, I use “varied” queries that I don’t repeat verbatim across search engines; this reduces the likelihood of identifying me. I also attempt to spread these tests out over time so admins won’t notice an unusual uptick in unpredictable and esoteric searches. This might seem overkill, but I already regularly employ similar methods for a variety of different scenarios.</p>
            				<h3 id="caveats" tabindex="-1">Caveats</h3>
            				<p>I didn’t try to avoid personalization when testing engines that require account creation. Entries in the “hit-and-miss” and “unusable” sections got less attention: I didn’t spend a lot of effort tracking results over time to see how new entries got added to them.</p>
            				<p>I avoided “natural language” queries like questions, focusing instead on keyword searches and search operators. I also mostly ignored infoboxes (also known as “instant answers”).</p>
            				<h2 id="findings" tabindex="-1">Findings</h2>
            				<aside role="none">
            					<a href="#findings" aria-labelledby="findings-prefix findings">
            						<span id="findings-prefix">Permalink to section</span>
            					</a>
            				</aside>
            				<p>What I learned by building this list has profoundly changed how I surf.</p>
            				<p>Using one engine for everything ignores the fact that different engines have different strengths. For example: while Google is focused on being an “answer engine”, other engines are better than Google at discovering new websites related to a broad topic. Fortunately, browsers like Chromium and Firefox make it easy to add many search engine shortcuts for easy switching.</p>
            				<p>When talking to search engine founders, I found that the biggest obstacle to growing an index is getting blocked by sites. Cloudflare was one of the worst offenders, although <a href="https://blog.cloudflare.com/friendly-bots">it has since launched a “verified bots” program to improve things</a>. Too many sites block perfectly well-behaved crawlers, only allowing major players like Googlebot, BingBot, and TwitterBot; this cements the current duopoly over English search and is harmful to the health of the Web as a whole.</p>
            				<p>Too many people optimize sites specifically for Google without considering the long-term consequences of their actions. One of many examples is how Google’s JavaScript support rendered the practice of testing a website without JavaScript or images “obsolete”: almost no non-GBY engines on this list are JavaScript-aware.</p>
            				<p>When building webpages, authors need to consider the barriers to entry for a new search engine. The best engines we can build today shouldn’t replace Google. They should try to be different. We want to see the Web that Google won’t show us, and search engine diversity is an important step in that direction.</p>
            				<p>Try a “bad” engine from lower in the list. It might show you utter crap. But every garbage heap has an undiscovered treasure. I’m sure that some hidden gems you’ll find will be worth your while. Let’s add some serendipity to the SEO-filled Web.</p>
            				<h2 id="acknowledgements" tabindex="-1">Acknow­ledgements</h2>
            				<aside role="none">
            					<a href="#acknowledgements" aria-labelledby="acknowledgements-prefix acknowledgements">
            						<span id="acknowledgements-prefix">Permalink to section</span>
            					</a>
            				</aside>
            				<p>Some of this content came from the <a href="https://www.searchenginemap.com/">Search Engine Map</a> and <a href="https://searchengine.party/">Search Engine Party</a>. A few web directories also proved useful.</p>
            				<p><span itemprop="mentions" itemscope="" itemtype="https://schema.org/Person" class="h-card vcard"><a itemprop="url" href="https://gigablast.com/bio.html" class="u-url url"><span itemprop="name" class="p-name fn n"><span itemprop="givenName" class="p-given-name given-name">Matt</span>&nbsp;<span itemprop="familyName" class="p-family-name family-name">Wells</span></span></a>
            from <span class="p-org org" itemprop="affiliation" itemscope="" itemtype="https://schema.org/Organization"><a itemprop="url" class="organization-name" href="https://gigablast.com/"><span itemprop="name">Gigablast</span></a></span></span> also gave me some helpful information about GBY which I included in the “Rationale” section. He’s written more about big tech in the <a href="https://gigablast.com/blog.html">Gigablast blog</a>.</p>
            				<p><span class="h-cite" itemprop="mentions" itemscope="" itemtype="https://schema.org/BlogPosting"><cite itemprop="name" class="p-name"><a class="u-url" itemprop="url" href="https://thenewleafjournal.com/a-2021-list-of-alternative-search-engines-and-search-resources/">A 2021 List of Alternative Search Engines and Search Resources</a></cite> by <span itemprop="author" itemscope="" itemtype="https://schema.org/Person" class="h-card vcard p-author"><a itemprop="url" href="https://emucafe.club/channel/naferrell" class="u-url url"><span itemprop="name" class="p-name fn n"><span itemprop="givenName" class="p-given-name given-name">Nicholas</span>&nbsp;<span itemprop="familyName" class="p-family-name family-name">Ferrell</span></span></a>
            from <span class="p-org org" itemprop="affiliation" itemscope="" itemtype="https://schema.org/Organization"><a itemprop="url" class="organization-name" href="https://thenewleafjournal.com/"><span itemprop="name">The New Leaf Journal</span></a></span></span></span> is a great post on alternative search engines. He also gave me some <a href="https://lists.sr.ht/~seirdy/seirdy.one-comments/%3C20210618031450.rb2twu4ypek6vvl3%40rkumarlappie.attlocal.net%3E">useful details</a> about Seznam, Naver, Baidu, and Goo.</p>
            				<hr>
            				<section role="doc-endnotes" aria-labelledby="note-hd">
            					<h2 id="note-hd">Notes</h2>
            					<ol>
            						<li id="fn:1" tabindex="-1">
            							<p>Yes, “indexes” is an acceptable plural form of the word “index”. The word “indices” sounds weird to me outside a math class.&nbsp;</p>
            							<a href="#fnref:1" aria-labelledby="bl1-1 bl2-1" role="doc-backlink">
            								<span id="bl1-1">Back</span>
            								<span id="bl2-1" hidden=""> to reference 1</span>
            							</a>
            						</li>
            						<li id="fn:2" tabindex="-1">
            							<p>Update: <a href="https://support.startpage.com/hc/en-us/articles/4522435533844-What-is-the-relationship-between-Startpage-and-your-search-partners-like-Google-and-Microsoft-Bing-">A Startpage support article</a> updated on <time>2023-03-08</time> claims that Startpage uses Microsoft (probably Bing) too. In my own tests, I still see Google results. I’ll update its placement if this changes.&nbsp;</p>
            							<a href="#fnref:2" aria-labelledby="bl1-2 bl2-2" role="doc-backlink">
            								<span id="bl1-2">Back</span>
            								<span id="bl2-2" hidden=""> to reference 2</span>
            							</a>
            						</li>
            						<li id="fn:3" tabindex="-1">
            							<p>DuckDuckGo has a crawler called DuckDuckBot. This crawler doesn’t impact the linked results displayed; it just grabs favicons and scrapes data for a few instant answers. DuckDuckGo’s help pages claim that the engine uses over 400 sources; my interpretation is that at least 398 sources don’t impact organic results. I don’t think DuckDuckGo is transparent enough about the fact that their organic results are proxied. Compare DuckDuckGo side-by-side with Bing and Yandex and you’ll see it’s sourcing organic results from one of them (probably Bing). <em>Update, March 2022:</em> DuckDuckGo <a href="https://web.archive.org/web/20220310222014/https://nitter.pussthecat.org/yegg/status/1501716484761997318">has the ability to downrank results on its own</a>; it was previously <a href="https://www.nytimes.com/2022/02/23/technology/duckduckgo-conspiracy-theories.html">working with Bing</a> to get Bing to remove misinformation and spam.&nbsp;</p>
            							<a href="#fnref:3" aria-labelledby="bl1-3 bl2-3" role="doc-backlink">
            								<span id="bl1-3">Back</span>
            								<span id="bl2-3" hidden=""> to reference 3</span>
            							</a>
            						</li>
            						<li id="fn:4" tabindex="-1"><p>Qwant claims to also use its own crawler for results, but it’s still mostly Bing in my experience. See the “semi-independent” section.&nbsp;</p><a href="#fnref:4" aria-labelledby="bl1-4 bl2-4" role="doc-backlink">↩︎</a>&nbsp;<a href="#fnref1:4" class="footnote-backref" role="doc-backlink"><span id="bl1-4">Back</span><span id="bl2-4" hidden=""> to reference 4</span></a></li>
            						<li id="fn:5" tabindex="-1">
            							<p>Disconnect Search allows users to have results proxied from Bing or Yahoo, but Yahoo sources its results from Bing.&nbsp;</p>
            							<a href="#fnref:5" aria-labelledby="bl1-5 bl2-5" role="doc-backlink">
            								<span id="bl1-5">Back</span>
            								<span id="bl2-5" hidden=""> to reference 5</span>
            							</a>
            						</li>
            						<li id="fn:6" tabindex="-1">
            							<p>Yippy claims to be powered by a certain IBM brand (a brand that could correspond to any number of products) and annotates results with the phrase “Yippy Index”, but a side-by-side comparison with Bing and other Bing-based engines revealed results to be nearly identical.&nbsp;</p>
            							<a href="#fnref:6" aria-labelledby="bl1-6 bl2-6" role="doc-backlink">
            								<span id="bl1-6">Back</span>
            								<span id="bl2-6" hidden=""> to reference 6</span>
            							</a>
            						</li>
            						<li id="fn:7" tabindex="-1">
            							<p>I’m in the process of re-evaluating You.com. It claims to operate a crawler and index. It seems very much like DuckDuckGo<sup><a href="#fn:4" id="fnref1:4" role="doc-noteref">note&nbsp;4</a></sup> to me: organic results look like they’re from Bing, while infoboxes (“apps”) seem to be scraped or queried from hand-picked websites. I’m not currently seeing results from “around the web” like the other engines that do pass my inclusion criteria. I might be wrong! I’m re-evaluating it to see if this isn’t actually the case.</p>
            							<p><ins datetime="2023-03-13T13:34:30-07:00">Update: You.com seems to source organic link results from Bing, and only interleaves those results with its own curated infoboxes</ins>&nbsp;</p>
            							<a href="#fnref:7" aria-labelledby="bl1-7 bl2-7" role="doc-backlink">
            								<span id="bl1-7">Back</span>
            								<span id="bl2-7" hidden=""> to reference 7</span>
            							</a>
            						</li>
            						<li id="fn:8" tabindex="-1">
            							<p>This is based on a statement Right Dao made in <a href="https://reddit.com/comments/k4clx1/_/ge9dwmh/?context=1">on Reddit</a> (<a href="https://web.archive.org/web/20210320042457/https://i.reddit.com/r/degoogle/comments/k4clx1/right_dao_a_new_independent_search_engine_that/ge9dwmh/?context=1">archived</a>).&nbsp;</p>
            							<a href="#fnref:8" aria-labelledby="bl1-8 bl2-8" role="doc-backlink">
            								<span id="bl1-8">Back</span>
            								<span id="bl2-8" hidden=""> to reference 8</span>
            							</a>
            						</li>
            						<li id="fn:9" tabindex="-1">
            							<p>More information can be found in <a href="https://news.ycombinator.com/item?id=27593801">this HN subthread</a> and some posts on the Cliqz tech blog (<a href="https://0x65.dev/blog/2019-12-06/building-a-search-engine-from-scratch.html">one</a>, <a href="https://0x65.dev/blog/2019-12-10/search-quality-at-cliqz.html">two</a>).&nbsp;</p>
            							<a href="#fnref:9" aria-labelledby="bl1-9 bl2-9" role="doc-backlink">
            								<span id="bl1-9">Back</span>
            								<span id="bl2-9" hidden=""> to reference 9</span>
            							</a>
            						</li>
            						<li id="fn:10" tabindex="-1">
            							<p>I will explain my thinking in another post later, and then edit this with a link to that post.&nbsp;</p>
            							<a href="#fnref:10" aria-labelledby="bl1-10 bl2-10" role="doc-backlink">
            								<span id="bl1-10">Back</span>
            								<span id="bl2-10" hidden=""> to reference 10</span>
            							</a>
            						</li>
            						<li id="fn:11" tabindex="-1">
            							<p>Some search engines support the <code>site:</code> search operator to limit searches to subpages or subdomains of a single site or TLD. <code>site:.one</code>, for instance, limits searches to websites with the “.one” TLD.&nbsp;</p>
            							<a href="#fnref:11" aria-labelledby="bl1-11 bl2-11" role="doc-backlink">
            								<span id="bl1-11">Back</span>
            								<span id="bl2-11" hidden=""> to reference 11</span>
            							</a>
            						</li>
            						<li id="fn:12" tabindex="-1">
            							<p>Matt from Gigablast told me that indexing YouTube or LinkedIn will get you blocked if you aren’t Google or Microsoft. I imagine that you could do so by getting special permission if you’re a megacorporation.&nbsp;</p>
            							<a href="#fnref:12" aria-labelledby="bl1-12 bl2-12" role="doc-backlink">
            								<span id="bl1-12">Back</span>
            								<span id="bl2-12" hidden=""> to reference 12</span>
            							</a>
            						</li>
            					</ol>
            				</section>
          text: |-
            Preface
            					This is a cursory review of all the indexing search engines I have been able to find.
            					The three dominant English search engines with their own indexesnote 1 are Google, Bing, and Yandex (GBY). Many alternatives to GBY exist, but almost none of them have their own results; instead, they just source their results from GBY.
            					With that in mind, I decided to test and catalog all the different indexing search engines I could find. I prioritized breadth over depth, and encourage readers to try the engines out themselves if they’d like more information.
            					This page is a “living document” that I plan on updating indefinitely. Check for updates once in a while if you find this page interesting. Feel free to send me suggestions, updates, and corrections; I’d especially appreciate help from those who speak languages besides English and can evaluate a non-English indexing search engine. Contact info is in the article footer.
            					I plan on updating the engines in the top two categories with more info comparing the structured/linked data the engines leverage (RDFa vocabularies, microdata, microformats, JSON-LD, etc.) to help authors determine which formats to use.
            				
            				
            					Toggle table of contents
            					
            						Table of Contents
            						
            							
            								About the list
            							
            							
            								General indexing search-engines
            								
            									
            										Large indexes, good results
            									
            									
            										Smaller indexes or less relevant results
            									
            									
            										Smaller indexes, hit-and-miss
            									
            									
            										Fledgling engines
            									
            									
            										Semi-independent indexes
            									
            								
            							
            							
            								Non-generalist search
            								
            									
            										Small or non-commercial Web
            									
            									
            										Site finders
            									
            									
            										Other
            									
            								
            							
            							
            								Other languages
            								
            									
            										Big indexes
            									
            									
            										Smaller indexes
            									
            								
            							
            							
            								Almost qualified
            							
            							
            								Misc
            							
            							
            								Search engines without a web interface
            							
            							
            								Graveyard
            							
            							
            								Exclusions
            							
            							
            								Rationale
            								
            									
            										Conflicts of interest
            									
            									
            										Information diversity
            									
            								
            							
            							
            								Method­ology
            								
            									
            										Discovery
            									
            									
            										Criteria for inclusion
            									
            									
            										Evaluation
            									
            									
            										Caveats
            									
            								
            							
            							
            								Findings
            							
            							
            								Acknow­ledgements
            							
            						
            					
            				
            				About the list
            				
            					
            						Permalink to section
            					
            				
            				I discuss my motivation for making this page in the Rationale section.
            				I primarily evaluated English-speaking search engines because that’s my primary language. With some difficulty, I could probably evaluate a Spanish one; however, I wasn’t able to find many Spanish-language engines powered by their own crawlers.
            				I mention details like “allows site submissions” and structured data support where I can only to inform authors about their options, not as points in engines’ favor.
            				See the Methodology section at the bottom to learn how I evaluated each one.
            				General indexing search-engines
            				
            					
            						Permalink to section
            					
            				
            				Large indexes, good results
            				These are large engines that pass all my standard tests and more.
            				
            					Google
            					The biggest index. Allows submitting pages and sitemaps for crawling, and even supports WebSub to automate the process. Powers a few other engines:
            A former version of Startpage, possibly the most popular Google proxy. Startpage now uses Bingnote 2GMX Search, run by a popular German email provider.(discontinued) RunnarooMullvad LetaSAPO (Portuguese interface, can work with English results)DSearch13TABSZarebin (Persian, can return English results)A host of other engines using Programmable Search Engine’s client-side scripts.

            					Bing
            					The runner-up. Allows submitting pages and sitemaps for crawling without login using the IndexNow API, sharing IndexNow page submissions with Yandex and Seznam. Its index powers many other engines:
            Yahoo (and its sibling engine, One­Search)DuckDuck­Gonote 3AOLQwant (partial)note 4EcosiaEkoruPrivadoFindxDisconnect Searchnote 5PrivacyWallLiloSearch­ScenePeekierOscoboMillion ShortYippy searchnote 6LycosGiveroSwisscowsFireballNetzzappenYou.comnote 7Partially powers MetaGer by default; this can be turned offAt this point, I mostly stopped adding Bing-based search engines. There are just too many.

            					Yandex
            					Originally a Russian search engine, it now has an English version. Some Russian results bleed into its English site. It allows submitting pages and sitemaps for crawling using the IndexNow API, sharing IndexNow page submissions with Bing and Seznam. Powers:
            Epic Search (went paid-only as of June 2021)Occasionally powers DuckDuck­Go’s link results instead of Bing (update: DuckDuckGo has “paused” its partnership with Yandex, confirmed in Hearing on “Holding Big Tech Accountable: Legislation to Protect Online Users”Petal, for Russian users only.

            					
            						Mojeek
            					
            					Seems privacy-oriented with a large index containing billions of pages. Quality isn’t at GBY’s level, but it’s not bad either. If I had to use Mojeek as my default general search engine, I’d live. Partially powers eTools.ch. At this moment, I think that Mojeek is the best alternative to GBY for general search.
            				
            				Google, Bing, and Yandex support structured data such as microformats1, microdata, RDFa, Open Graph markup, and JSON-LD. Yandex’s support for microformats1 is limited; for instance, it can parse h-card metadata for organizations but not people. Open Graph and Schema.org are the only supported vocabularies I’m aware of. Mojeek is evaluating structured data; it’s interested in Open Graph and Schema.org vocabularies.
            				Smaller indexes or less relevant results
            				These engines pass most of the tests listed in the “methodology” section. All of them seem relatively privacy-friendly. I wouldn’t recommend using these engines to find specific answers; they’re better for learning about a topic by finding interesting pages related to a set of keywords.
            				
            					
            						Stract
            					
            					My favorite generalist engine on this page. Stract supports advanced ranking customization by allowing users ti import “optics” files, like a better version of Brave’s “goggles” feature. Stract is fully open-source, with code released under an AGPL-3.0 license. The index is isn’t massive but it’s big enough to be a useful supplement to more major engines. Stract started with the Common Crawl index, but now uses its own crawler. Plans to add contextual ads and a subscription option for ad-free search. Discovered in my access logs.
            					
            						Right Dao
            					
            					Very fast, good results. Passes the tests fairly well. It plans on including query-based ads if/when its user base grows.note 8 For the past few months, its index seems to have focused more on large, established sites rather than smaller, independent ones. It seems to be a bit lacking in more recent pages.
            					
            						Alexandria
            					
            					A pretty new “non-profit, ad free” engine, with freely-licensed code. Surprisingly good at finding recent pages. Its index is built from the Common Crawl; it isn’t as big as Gigablast or Right Dao but its ranking is great.
            					
            						Yep
            					
            					An ambitious engine from Ahrefs, an SEO/backlink-finder company, that “shares ad profit with creators and protects your privacy”. Most engines show results that include keywords from or related to the query; Yep also shows results linked by pages containing the query. In other words, not all results contain relevant keywords. This makes it excellent for less precise searches and discovery of “related sites”, especially with its index of hundreds of billions of pages. It’s far worse at finding very specific information or recent events for now, but it will probably improve. It was known as “FairSearch” before its official launch.
            					
            						SeSe Engine
            					
            					Although it’s a Chinese engine, its index seems to have a large-enough proportion of English content to fit here. The engine is open-source; see the SeSe back-end Python code and the SeSe-ui Vue-based front-end. It has surprisingly good results for such a low-budget project. Each result is annotated with detailed ranking metadata such as keyword relevance and backlink weight. Discovered in my access logs.
            					
            						greppr
            					
            					Its tagline is “Search the Internet with no filters, no tracking, no ads.” At the time of writing, it has over 3 million pages indexed. It’s surprisingly good at finding interesting new results for broad short-tail queries, if you’re willing to scroll far enough down the page. It appears to be good at finding recent pages.
            				
            				Yep supports Open Graph and some JSON-LD at the moment. A look through the source code for Alexandria and Gigablast didn’t seem to reveal the use of any structured data. The surprising quality of results from SeSe and Right Dao seems influenced by the crawlers’ high-quality starting location: Wikipedia.
            				Smaller indexes, hit-and-miss
            				These engines fail badly at a few important tests. Otherwise, they seem to work well enough for users who’d like some more serendipity in less-specific searches.
            				
            					
            						Infotiger
            					
            					My favorite engine in this section. It offers advanced result filtering and sports a somewhat large index. It allows site submission for English and German pages. The fastest-improving engine in this section: I use it often to discover new sites, and look forward to the day it “graduates” to the previous section. Infotier has a Tor hidden service.
            					
            						seekport
            					
            					The interface is in German but it supports searching in English just fine. The default language is selected by your locale. It’s really good considering its small index; it hasn’t heard of less common terms. but it’s able to find relevant results in other tests. It’s the second-fastest-improving engines in this section.
            					
            						Exalead
            					
            					Slow, quality is hit-and-miss. Its indexer claims to crawl the DMOZ directory, which has since shut down and been replaced by the Curlie directory. No relevant results for “Oppenheimer” and some other history-related queries. Allows submitting individual URLs for indexing, but requires solving a Google reCAPTCHA and entering an email address.
            					
            						ExactSeek
            					
            					Small index, disproportionately dominated by big sites. Failed multiple tests. Allows submitting individual URLs for crawling, but requires entering an email address and receiving a newsletter. Webmaster tools seem to heavily push for paid SEO options. It also powers SitesOnDisplay and Blog-search.com.
            					
            						Burf.co
            					
            					Very small index, but seems fine at ranking more relevant results higher. Allows site submission without any extra steps.
            					
            						Siik
            					
            					Lacks contact info, and the ToS and Privacy Policy links are dead. Seems to have PHP errors in the backend for some of its instant-answer widgets. If you scroll past all that, it does have web results powered by what seems to be its own index. These results do tend to be somewhat relevant, but the index seems too small for more specific queries.
            					
            						ChatNoir
            					
            					An experimental engine by researchers that uses the Common Crawl index. The engine is open source. See the announcement on the Common Crawl mailing list (Google Groups).
            					
            						Secret Search Engine Labs
            					
            					Very small index with very little SEO spam; it toes the line between a “search engine” and a “surf engine”. It’s best for reading about broad topics that would otherwise be dominated by SEO spam, thanks to its CashRank algorithm. Allows site submission.
            					
            						Gabanza
            					
            					A search engine from a hosting company. I found few details abou the search engine itself, and the index was small, but it was suitable for discovering new pages related to short broad queries.
            					
            						Jambo
            					
            					Docs, blog posts, etc. have not been updated since around 2006 but the engine continues to crawl and index new pages. Discovered in my access logs. Has a bias towards older content.
            				
            				Fledgling engines
            				Results from these search engines don’t seem particularly relevant; indexes in this category tend to be small.
            				
            					
            						Yessle
            					
            					Seems new; allows page submission by pasting a page into the search box. Index is really small but it crawls new sites quickly. Claims to be private.
            					
            						Bloopish
            					
            					Extremely quick to update its index; site submissions show up in seconds. Unfortunately, its index only contains a few thousand documents (under 100 thousand at the time of writing). It’s growing fast: if you search for a term, it’ll start crawling related pages and grow its index.
            					YaCy
            					Community-made index; slow. Results are awful/irrelevant, but can be useful for intranet or custom search.
            					Scopia
            					only seems to be available via the MetaGer metasearch engine after turning off Bing and news results. Tiny index, very low-quality.
            					
            						Artado Search
            					
            					Primarily Turkish, but it also seems to support English results. Like Plumb, it uses client-side JS to fetch results from existing engines (Google, Bing, Yahoo, Petal, and others); like MetaGer, it has an option to use its own independent index. Results from its index are almost always empty. Very simple queries (“twitter”, “wikipedia”, “reddit”) give some answers. Supports site submission and crowdsourced instant answers.
            					
            						Active Search Results
            					
            					Very poor quality. Results seem highly biased towards commercial sites.
            					
            						Crawlson
            					
            					Young, slow. In this category because its index has a cap of 10 URLs per domain. I initially discovered Crawlson in the seirdy.one access logs. This is often down; if the current downtime persists, I’ll add it to the graveyard.
            					
            						Anoox
            					
            					Results are few and irrelevant; fails to find any results for basic terms. Allows site submission. It’s also a lightweight social network and claims to be powered by its users, letting members vote on listings to alter rankings.
            					
            						Yioop!
            					
            					A FLOSS search engine that boasts a very impressive feature-set: it can parse sitemaps, feeds, and a variety of markup formats; it can import pre-curated data in forms such as access logs, Usenet posts, and WARC archives; it also supports feed-based news search. Despite the impressive feature set, Yioop’s results are few and irrelevant due to its small index. It allows submitting sites for crawling. Like Meorca, Yioop has social features such as blogs, wikis, and a chat bot API.
            					
            						Spyda
            					
            					A small engine made by James Mills, described in So I'm a Knucklehead eh?. It’s written in Go; check out its MIT-licensed Spyda source code.
            					
            						Slzii.com
            					
            					A new web portal with a search engine. Has a tiny index dominated by SEO spam. Discovered in the seirdy.one access logs.
            				
            				Semi-independent indexes
            				Engines in this category fall back to GBY when their own indexes don’t have enough results. As their own indexes grow, some claim that this should happen less often.
            				
            					
            						Brave Search
            					
            					Many tests (including all the tests I listed in the “Methodology” section) resulted results identical to Google, revealed by a side-by-side comparison with Google, Startpage, and a Searx instance with only Google enabled. Brave claims that this is due to how Cliqz (the discontinued engine acquired by Brave) used query logs to build its page models and was optimized to match Google.note 9 The index is independent, but optimizing against Google resulted in too much similarity for the real benefit of an independent index to show. Furthermore, many queries have Bing results mixed in; users can click an “info” button to see the percentage of results that came from its own index. The independent percentage is typically quite high (often close to 100% independent) but can drop for advanced queries. Update 2023-08-15: Brave’s Bing contract appears to have expired as of April 2023.
            I can’t in good conscience recommend using Brave Search, as the company runs cryptocurrency, has held payments to creators without disclosing that creators couldn’t receive rewards, has made dangerously misleading claims about fingerprinting resistance,note 10 is run by a CEO who spent thousands of dollars opposing gay marriage, and has rewritten typed URLs with affiliate links.

            					
            						Plumb
            					
            					Almost all queries return no results; when this happens, it falls back to Google. It’s fairly transparent about the fallback process, but I’m concerned about how it does this: it loads Google’s Custom Search scripts from cse.google.com onto the page to do a client-side Google search. This can be mitigated by using a browser addon to block cse.google.com from loading any scripts. Plumb claims that this is a temporary measure while its index grows, and they’re planning on getting rid of this. Allows submitting URLs, but requires solving an hCaptcha. This engine is very new; hopefully as it improves, it could graduate from this section. Its Chief Product Officer previously founded the Gibiru search engine which shares the same affiliates and (for now) the same index; the indexes will diverge with time.
            					
            						Qwant
            					
            					Qwant claims to use its own index, but it still relies on Bing for most results. It seems to be in a position similar to Neeva. Try a side-by-side comparison to see if or how it compares with Bing.
            					
            						Kagi Search
            					
            					The most interesting entry in this category, IMO. Like Neeva, it requires an account and limits use without payment. It’s powered by its own Teclis index (Teclis can be used independently; see the non-commercial section below), and claims to also use results from Google and Bing. The result seems somewhat unique: I’m able to recognize some results from the Teclis index mixed in with the mainstream ones. In addition to Teclis, Kagi’s other products include the Kagi.ai intelligent answer service and the TinyGem social bookmarking service, both of which play a role in Kagi.com in the present or future. Unrelatedly: I’m concerned about the company’s biases, as it seems happy to use Brave’s commercial API (allowing blatant homophobia in the comments) and allow its results to recommend suicide methods without intervention. I reject the idea that avoiding an option that may seem politically biased is the same as being unbiased if such a decision has real political implications.
            					
            						SVMetaSearch
            					
            					A SearxNG metasearch engine that also includes results from its own index. All other sources can be turned off. Like most public Searx/SearxNG instances, reliability is very poor.
            				
            				Non-generalist search
            				
            					
            						Permalink to section
            					
            				
            				These indexing search engines don’t have a Google-like “ask me anything” endgame; they’re trying to do something different. You aren’t supposed to use these engines the same way you use GBY.
            				Small or non-commercial Web
            				
            					
            						Marginalia Search
            					
            					My favorite entry on this page. It has its own crawler but is strongly biased towards non-commercial, personal, and/or minimal sites. It’s a great response to the increasingly SEO-spam-filled SERPs of GBY. Partially powers Teclis, which in turn partially powers Kagi. Update 2022-05-28: Marginalia.nu is now open source.
            					
            						Ichido
            					
            					An engine that just rolled out its own independent index, with a lot of careful thought put into its ranking algorithm. Like Marginalia, it’s biased towards the non-commercial Web: it downranks ads, CAPTCHAs, trackers, SEO, and obfuscation. More info about Ichido is in a blog post.
            					
            						Teclis
            					
            					A project by the creator of Kagi search. Uses its own crawler that measures content blocked by uBlock Origin, and extracts content with the open-source article scrapers Trafilatura and Readability.js. This is quite an interesting approach: tracking blocked elements discourages tracking and advertising; using Trafilatura and Readability.js encourages the use of semantic HTML and Semantic Web standards such as microformats, microdata, and RDFa. It claims to also use some results from Marginalia. The Web interface has been shut down, but its standalone API is still available for Kagi customers.
            					
            						Clew
            					
            					a FOSS new engine with a small index of several thousand pages. It focuses on independent content and downranks ads and trackers; there seems to be a real focus on quality over quantity, which makes it excellent for short-tail searches (especially around technical concepts). Ranking is more egalitarian than other engines, making it better for discovery and surfing than research. It’s designed to be small and lightweight, with a compact index. Discovered in the seirdy.one access logs.
            				
            				Site finders
            				These engines try to find a website, typically at the domain-name level. They don’t focus on capturing particular pages within websites.
            				
            					
            						Kozmonavt
            					
            					The best in this category. Has a small but growing index of over 8 million sites. If I want to find the website for a certain project, Kozmonavt works well (provided its index has crawled said website). It works poorly for learning things and finding general information. I cannot recommend it for anything serious since it lacks contact information, a privacy policy, or any other information about the org/people who made it. Discovered in the seirdy.one access logs.
            					
            						search.tl
            					
            					Generalist search for one TLD at a time (defaults to .com). I’m not sure why you’d want to always limit your searches to a single TLD, but now you can.note 11 There isn’t any visible UI for changing the TLD for available results; you need to add/change the tld URL parameter. For example, to search .org sites, append &tld=org to the URL. It seems to be connected to Amidalla. Amidalla allows users to manually add URLs to its index and directory; I have yet to see if doing so impacts search.tl results.
            					
            						Thunderstone
            					
            					A combined website catalog and search engine that focuses on categorization. Its about page claims: We continuously survey all primary COM, NET, and ORG web-servers and distill their contents to produce this database. This is an index of sites not pages. It is very good at finding companies and organizations by purpose, product, subject matter, or location. If you’re trying to finding things like ‘BillyBob’s personal beer can page on AOL’, try Yahoo or Dogpile. This seems to be the polar opposite of the engines in the “small or non-commercial Web” category.
            					
            						sengine.info
            					
            					Only shows domains, not individual pages. Developed by netEstate GmbH, which specializes in content extraction for inprints and job ads. Also has a German-only version available. Discovered in my access logs.
            					
            						Gnomit
            					
            					Allows single-keyword queries and returns sites that seem to cover a related topic. I actually kind of enjoy using it; results are old (typically from 2009) and a bit random, but make for a nice way to discover something new. For instance, searching for “IRC” helped me discover new IRC networks I’d never heard of.
            				
            				Other
            				
            					
            						High Browse
            					
            					Uses a non-traditional ranking algorithm which does an excellent job of introducing non-SEO-optimized serendipity into search results. One of my favorite “surf-engines”, as opposed to traditional “search-engines”.
            					
            						Keybot
            					
            					A must-have for anyone who does translation work. It crawls the web looking for multilingual websites. Translators who are unsure about how to translate a given word or phrase can see its usage in two given languages, to learn from other human translators. My parents are fluent English speakers but sometimes struggle to express a given Hindi idiom in English; something like this could be useful to them, since machine translation isn’t nuanced enough for every situation. Part of the TTN Translation Network. Discovered in my access logs.
            					Quor
            					Seems to mainly index large news sites. Site is down as of June 2021; originally available at www dot quor dot com.
            					
            						Semantic Scholar
            					
            					A search engine by the Allen Institute for AI focused on academic PDFs, with a couple hundred million papers indexed. Discovered in my access logs.
            					
            						Bonzamate
            					
            					A search engine specifically for Australian websites. Boyter wrote an interesting blog post about Bonzamate.
            					
            						searchcode
            					
            					A code-search engine by the developer of Bonzamate. Searches a hand-picked list of code forges for source code, supporting many search operators.
            					
            						Lixia Labs Search
            					
            					A new engine that focuses on indexing technical websites and blogs, with a minimal JavaScript-free front-end. Discovered in my access logs. Surprisingly good results for broad technical keyword queries.
            				
            				Other languages
            				
            					
            						Permalink to section
            					
            				
            				I’m unable to evaluate these engines properly since I don’t speak the necessary languages. English searches on these are a hit-or-miss. I might have made a few mistakes in this category.
            				Big indexes
            				
            					
            						Baidu: Chinese. Very large index; it’s a major engine alongside GBY. Offers webmaster tools for site submission.
            					
            					
            						Qihoo 360: Chinese. I’m not sure how independent this one is.
            					
            					
            						Toutiao: Chinese. Not sure how independent this one is either. Its index appears limited outside of its own content distribution platform.
            					
            					
            						Sogou: Chinese
            					
            					
            						Yisou: Chinese, by Yahoo. Appears defunct.
            					
            					
            						Naver: Korean. Allows submitting sitemaps and feeds. Discovered via some Searx metasearch instances.
            					
            					
            						Daum: Korean. Also unsure about this one’s independence.
            					
            					
            						Seznam: Czech, seems relatively privacy-friendly. Discovered in the seirdy.one access logs. It allows site submission with webmaster tools. Seznam supports IndexNow; it shares IndexNow-submitted pages with Bing and Yandex.
            					
            					
            						Cốc Cốc: Vietnamese
            					
            					
            						go.mail.ru: Russian
            					
            					
            						LetSearch.ru: Russian. Allows URL submission
            					
            				
            				Smaller indexes
            				
            					
            						ALibw.com: Chinese, found in my access logs.
            					
            					
            						Vuhuv: Turkish. alt domain
            					
            					
            						search.ch: Regional search engine for Switzerland; users can restrict searches to their local regions.
            					
            					
            						fastbot: German
            					
            					
            						Moose.at: German (Austria-based)
            					
            					
            						SOLOFIELD: Japanese
            					
            					
            						kaz.kz: Kazakh and Russian, with a focus on “Kazakhstan’s segment of the Internet”
            					
            				
            				Almost qualified
            				
            					
            						Permalink to section
            					
            				
            				These engines come close enough to passing my inclusion criteria that I felt I had to mention them. They all display original organic results that you can’t find on other engines, and maintain their own indexes. Unfortunately, they don’t quite pass.
            				
            					
            						wiby.me
            					
            					
            						wiby.org
            					
            					I love this one. It focuses on smaller independent sites that capture the spirit of the “early” web. It’s more focused on “discovering” new interesting pages that match a set of keywords than finding a specific resources. I like to think of Wiby as an engine for surfing, not searching. Runnaroo occasionally featured a hit from Wiby (Runnaroo has since shut down). If you have a small site or blog that isn’t very “commercial”, consider submitting it to the index. Does not qualify because it seems to be powered only by user-submitted sites; it doesn’t try to “crawl the Web”.
            					
            						Mwmbl
            					
            					Like YaCy, it’s an open-source engine whose crawling is community-driven. Users can install a Firefox addon to crawl pages in its backlog. Unfortunately, it doesn’t qualify because it only crawls pages linked by hand-picked sites (e.g. Wikipedia, GitHub, domains that rank well on Hacker News). The crawl-depth is “1”, so it doesn’t crawl the whole Web (yet).
            					
            						Search My Site
            					
            					Similar to Marginalia and Teclis, but only indexes user-submitted personal and independent sites. It optionally supports IndieAuth. Its API powers this site’s search results; try it out using the search bar at the bottom of this page. Does not qualify because it’s limited to user-submitted and/or hand-picked sites.
            					
            						Blog Surf
            					
            					A search engine for blogs with RSS/Atom feeds. Does not qualify because all blogs submitted to the index require manual review, but it seems interesting. Its “MarketRank” algorithm seems to give it a bias towards sites popular on “Hacker” “News”.
            					
            						Kukei.eu
            					
            					A curated search engine for web developers, which crawls a hand-picked list of sites. As it does not index the whole Web, it doesn’t qualify. I still find it interesting.
            				
            				Misc
            				
            					
            						Permalink to section
            					
            				
            				
            					Ask.com
            					The site is back. They claim to outsource search results. The results seem similar to Google, Bing, and Yandex; however, I can’t pinpoint exactly where their results are coming from. Also, several sites from the “ask.com network” such as directhit.com, info.com, and kensaq.com have uniqe-looking results.
            					
            						Infinity Search
            					
            					Partially evaluated. Young, small index. It recently split into a paid offering with the main index and Infinity Decentralized, the latter of which allows users to select from community-hosted crawlers. I managed to try it out before it became a paid offering, and it seemed decent; however, I wasn’t able to run the tests listed in the “Methodology” section. Allows submitting URLs and sitemaps into a text box, no other work required.
            				
            				Search engines without a web interface
            				
            					
            						Permalink to section
            					
            				
            				Some search engines are integrated into other appliances, but don’t have a web portal.
            				
            					
            						Apple’s search engine is usable in the form of “Siri Suggested Websites”. Its index is built from the Applebot web crawler. If Apple already has working search engine, it’s not much of a stretch to say that they’ll make a web interface for it someday.
            					
            					
            						Amazon bought Alexa Internet (a web traffic analysis company, at the time unrelated to the Amazon Alexa virtual assistant) and discontinued its website ranking product. Amazon still runs the relevant crawlers, and also have a bot called “Amazonbot”. While Applebot powers the Siri personal assistant, Amazonbot powers the Alexa personal assistant to answer even more questions for customers. Crawling the web to answer questions is the basis of a search engine.
            					
            				
            				Graveyard
            				
            					
            						Permalink to section
            					
            				
            				These engines were originally included in the article, but have since been discontinued.
            				
            					
            						Petal Search
            					
            					A search engine by Huawei that recently switched from searching for Android apps to general search in order to reduce dependence on Western search providers. Despite its surprisingly good results, I wouldn’t recommend it due to privacy concerns: its privacy policy describes advanced fingerprinting metrics, and it doesn’t work without JavaScript. Requires an account to submit sites. I discovered this via my access logs. Be aware that in some jurisdictions, it doesn’t use its own index: in Russia and some EU regions it uses Yandex and Qwant, respectively. Inaccessible as of June 2023.
            					
            						Neeva
            					
            					Formerly in the “semi-independent” section. Combined Bing results with results from its own index. Bing normally isn’t okay with this, but Neeva was one of few exceptions. Results were mostly identical to Bing, but original links not found by Bing frequently popped up. Long-tail and esoteric queries were less likely to feature original results. Required signing up with an email address or OAuth to use, and offered a paid tier with additional benefits. Acquired by Snowflake and announced its shut-down in May 2023.
            					
            						Gigablast
            					
            					It’s been around for a while and also sports a classic web directory. Searches are a bit slow, and it charges to submit sites for crawling. It powers Private.sh. Gigablast was tied with Right Dao for quality. Shut down mid-2023.
            					
            						wbsrch
            					
            					In addition to its generalist search, it also had many other utilities related to domain name statistics. Failed multiple tests. Its index was a bit dated; it had an old backlog of sites it hadn’t finished indexing. It also had several dedicated per-language indexes.
            					
            						Gowiki
            					
            					Very young, small index, but showed promise. I discovered this in the seirdy.one access logs. It was only available in the US. Seems down as of early 2022.
            					
            						Meorca
            					
            					A UK-based search engine that claims not to “index pornography or illegal content websites”. It also features an optional social network (“blog”). Discovered in the seirdy.one access logs.
            					
            						Ninfex
            					
            					A “people-powered” search engine that combines aspects of link aggregators and search. It lets users vote on submissions and it also displays links to forums about submissions.
            					
            						Marlo
            					
            					Another FLOSS engine: Marlo is written in Haskell. Has a small index that’s good enough for surfing broad topics, but not good enough for specific research. Originally available at marlo.sandymaguire.me.
            					websearchengine.org
            					tuxdex.com
            					Both were run by the same people, powered by their inetdex.com index. Searches are fast, but crawls are a bit shallow. Claims to have an index of 10 million domains, and not to use cookies. The pages are currently down and the domains re-direct to porn sites; I’m not aware of any official notice.
            					
            						Entfer
            					
            					a newcomer that let registered users upvote/downvote search results to customize ranking. Didn’t offer much information about who made it. Its index was small, but it did seem to return results related to the query.
            				
            				Dead engines I don’t have an extended description for:
            				
            					Parsijoo: Persian search engine.
            				
            				Exclusions
            				
            					
            						Permalink to section
            					
            				
            				Two engines were excluded from this list for having a far-right focus.
            				One engine was excluded because it seems to be built using cryptocurrency in a way I’d rather not support.
            				Some fascinating little engines seem like hobbyist proofs-of-concept. I decided not to include them in this list, but watch them with interest to see if they can become something viable.
            				Rationale
            				
            					
            						Permalink to section
            					
            				
            				Why bother using non-mainstream search engines?
            				Conflicts of interest
            				Google, Microsoft (the company behind Bing), and Yandex aren’t just search engine companies; they’re content and ad companies as well. For example, Google hosts video content on YouTube and Microsoft hosts social media content on LinkedIn. This gives these companies a powerful incentive to prioritize their own content. They are able to do so even if they claim that they treat their own content the same as any other: since they have complete access to their search engines’ inner workings, they can tailor their content pages to better fit their algorithms and tailor their algorithms to work well on their own content. They can also index their own content without limitations but throttle indexing for other crawlers.note 12
            				One way to avoid this conflict of interest is to use search engines that aren’t linked to major content providers; i.e., use engines with their own independent indexes.
            				Information diversity
            				There’s also a practical, less-ideological reason to try other engines: different providers have different results. Websites that are hard to find on one search engine might be easy to find on another, so using more indexes and ranking algorithms results in access to more content.
            				No search engine is truly unbiased. Most engines’ ranking algorithms incorporate a method similar to PageRank, which biases them towards sites with many backlinks. Search engines have to deal with unwanted results occupying the confusing overlap between SEO spam, shock content, and duplicate content. When this content’s manipulation of ranking algos causes it to rank high, engines have to address it through manual action or algorithm refinement. Choosing to address it through either option, or choosing to leave it there for popular queries after receiving user reports, reflects bias. The best solution is to mix different ranking algorithms and indexes instead of using one engine for everything.
            				Method­ology
            				
            					
            						Permalink to section
            					
            				
            				Discovery
            				I find new engines by:
            				
            					Monitoring certain web directories for changes in their search engine listings.
            					Checking other curated lists of “good/bad bots” to spot search engines.
            					Using search engines to discover search engines: searching for the names of less-popular engines often pulls up similar lists.
            					Receiving suggestions from readers
            					Compiling a list of regular expressions for user-agent strings I’m familiar with. Before I delete my server access logs, I extract user-agents that don’t match that list along with the pages they request.
            					Checking the Searx and Searxng projects for new integrations.
            				
            				Criteria for inclusion
            				Engines in this list should have their own indexes powered by web crawlers. Original results should not be limited to a set of websites hand-picked by the engine creators; indexes should be built from sites from across the Web. An engine should discover new interesting places around the Web.
            				Here’s an oversimplified example to illustrate what I’m looking for: imagine somone self-hosts their own personal or interest-specific website and happens to get some recognition. Could they get automatically discovered by your crawler, indexed, and included in the first page of results for a certain query?
            				I’m willing to make two exceptions:
            				
            					Engines in the “semi-independent” section may mix results that do meet the aforementioned criteria with results that do not.
            					Engines in the “almost qualified” section may use indexes primarily made of user-submitted or hand-picked sites, rather than focusing primarily on sites discovered organically through crawling.
            				
            				The reason the second exception exists is that while user submissions don’t represent automatic crawling, they do at least inform the engine of new interesting websites that it had not previously discovered; these websites can then be shown to other users. That’s fundamentally what an alternative web index needs to achieve.
            				I’m not usually willing to budge on my “no hand-picked websites” rule. Hand-picked sites will be ignored, whether your engine fetches content through their APIs or crawls and scrapes their content. It’s fine to use hand-picked websites as starting points for your crawler (Wikipedia is a popular option).
            				I only consider search engines that focus on link results for webpages. Image search engines are out of scope, though I might consider some other engines for non-generalist search (e.g., Semantic Scholar finds PDFs rather than webpages).
            				Evaluation
            				I focused almost entirely on “organic results” (the classic link results), and didn’t focus too much on (often glaring) privacy issues, “enhanced” or “instant” results (e.g. Wikipedia sidebars, related searches, Stack Exchange answers), or other elements.
            				I compared results for esoteric queries side-by-side; if the first 20 results were (nearly) identical to another engine’s results (though perhaps in a slightly different order), they were likely sourced externally and not from an independent index.
            				I tried to pick queries that should have a good number of results and show variance between search engines. An incomplete selection of queries I tested:
            				
            					
            						“vim”, “emacs”, “neovim”, and “nvimrc”: Search engines with relevant results for “nvimrc” typically have a big index. Finding relevant results for the text editors “vim” and “emacs” instead of other topics that share the name is a challenging task.
            					
            					
            						“vim cleaner”: should return results related to a line of cleaning products rather than the Correct Text Editor.
            					
            					
            						“Seirdy”: My site is relatively low-traffic, but my nickname is pretty unique and visible on several of the highest-traffic sites out there.
            					
            					
            						“Project London”: a small movie made with volunteers and FLOSS without much advertising. If links related to the movie show up, the engine’s really good.
            					
            					
            						“oppenheimer”: a name that could refer to many things. Without context, it should refer to the physicist who worked on the atomic bomb in Los Alamos. Other historical queries: “magna carta” (intermediate), “the prince” (very hard).
            					
            				
            				Some less-mainstream engines have noticed this article, which is great! I’ve had excellent discussions with people who work on several of these engines. Unfortunately, this article’s visibility also incentivizes some engines to optimize specifically for any methodology I describe. I’ve addressed this by keeping a long list of test queries to myself. The simple queries above are a decent starting point for simple quick evaluations, but I also test for common search operators, keyword length, and types of domain-specific jargon. I also use queries designed to pull up specific pages with varying levels of popularity and recency to gauge the size, scope, and growth of an index.
            				Professional critics often work anonymously because personalization can damage the integrity of their reviews. For similar reasons, I attempt to try each engine anonymously at least once by using a VPN and/or my standard anonymous setup: an amnesiac Whonix VM with the Tor Browser. I also often test using a fresh profile when travelling, or via a Searx instance if it supports a given engine. When avoiding personalization, I use “varied” queries that I don’t repeat verbatim across search engines; this reduces the likelihood of identifying me. I also attempt to spread these tests out over time so admins won’t notice an unusual uptick in unpredictable and esoteric searches. This might seem overkill, but I already regularly employ similar methods for a variety of different scenarios.
            				Caveats
            				I didn’t try to avoid personalization when testing engines that require account creation. Entries in the “hit-and-miss” and “unusable” sections got less attention: I didn’t spend a lot of effort tracking results over time to see how new entries got added to them.
            				I avoided “natural language” queries like questions, focusing instead on keyword searches and search operators. I also mostly ignored infoboxes (also known as “instant answers”).
            				Findings
            				
            					
            						Permalink to section
            					
            				
            				What I learned by building this list has profoundly changed how I surf.
            				Using one engine for everything ignores the fact that different engines have different strengths. For example: while Google is focused on being an “answer engine”, other engines are better than Google at discovering new websites related to a broad topic. Fortunately, browsers like Chromium and Firefox make it easy to add many search engine shortcuts for easy switching.
            				When talking to search engine founders, I found that the biggest obstacle to growing an index is getting blocked by sites. Cloudflare was one of the worst offenders, although it has since launched a “verified bots” program to improve things. Too many sites block perfectly well-behaved crawlers, only allowing major players like Googlebot, BingBot, and TwitterBot; this cements the current duopoly over English search and is harmful to the health of the Web as a whole.
            				Too many people optimize sites specifically for Google without considering the long-term consequences of their actions. One of many examples is how Google’s JavaScript support rendered the practice of testing a website without JavaScript or images “obsolete”: almost no non-GBY engines on this list are JavaScript-aware.
            				When building webpages, authors need to consider the barriers to entry for a new search engine. The best engines we can build today shouldn’t replace Google. They should try to be different. We want to see the Web that Google won’t show us, and search engine diversity is an important step in that direction.
            				Try a “bad” engine from lower in the list. It might show you utter crap. But every garbage heap has an undiscovered treasure. I’m sure that some hidden gems you’ll find will be worth your while. Let’s add some serendipity to the SEO-filled Web.
            				Acknow­ledgements
            				
            					
            						Permalink to section
            					
            				
            				Some of this content came from the Search Engine Map and Search Engine Party. A few web directories also proved useful.
            				Matt Wells
            from Gigablast also gave me some helpful information about GBY which I included in the “Rationale” section. He’s written more about big tech in the Gigablast blog.
            				A 2021 List of Alternative Search Engines and Search Resources by Nicholas Ferrell
            from The New Leaf Journal is a great post on alternative search engines. He also gave me some useful details about Seznam, Naver, Baidu, and Goo.
            				
            				
            					Notes
            					
            						
            							Yes, “indexes” is an acceptable plural form of the word “index”. The word “indices” sounds weird to me outside a math class. 
            							
            								Back
            								 to reference 1
            							
            						
            						
            							Update: A Startpage support article updated on 2023-03-08 claims that Startpage uses Microsoft (probably Bing) too. In my own tests, I still see Google results. I’ll update its placement if this changes. 
            							
            								Back
            								 to reference 2
            							
            						
            						
            							DuckDuckGo has a crawler called DuckDuckBot. This crawler doesn’t impact the linked results displayed; it just grabs favicons and scrapes data for a few instant answers. DuckDuckGo’s help pages claim that the engine uses over 400 sources; my interpretation is that at least 398 sources don’t impact organic results. I don’t think DuckDuckGo is transparent enough about the fact that their organic results are proxied. Compare DuckDuckGo side-by-side with Bing and Yandex and you’ll see it’s sourcing organic results from one of them (probably Bing). Update, March 2022: DuckDuckGo has the ability to downrank results on its own; it was previously working with Bing to get Bing to remove misinformation and spam. 
            							
            								Back
            								 to reference 3
            							
            						
            						Qwant claims to also use its own crawler for results, but it’s still mostly Bing in my experience. See the “semi-independent” section. ↩︎ Back to reference 4
            						
            							Disconnect Search allows users to have results proxied from Bing or Yahoo, but Yahoo sources its results from Bing. 
            							
            								Back
            								 to reference 5
            							
            						
            						
            							Yippy claims to be powered by a certain IBM brand (a brand that could correspond to any number of products) and annotates results with the phrase “Yippy Index”, but a side-by-side comparison with Bing and other Bing-based engines revealed results to be nearly identical. 
            							
            								Back
            								 to reference 6
            							
            						
            						
            							I’m in the process of re-evaluating You.com. It claims to operate a crawler and index. It seems very much like DuckDuckGonote 4 to me: organic results look like they’re from Bing, while infoboxes (“apps”) seem to be scraped or queried from hand-picked websites. I’m not currently seeing results from “around the web” like the other engines that do pass my inclusion criteria. I might be wrong! I’m re-evaluating it to see if this isn’t actually the case.
            							Update: You.com seems to source organic link results from Bing, and only interleaves those results with its own curated infoboxes 
            							
            								Back
            								 to reference 7
            							
            						
            						
            							This is based on a statement Right Dao made in on Reddit (archived). 
            							
            								Back
            								 to reference 8
            							
            						
            						
            							More information can be found in this HN subthread and some posts on the Cliqz tech blog (one, two). 
            							
            								Back
            								 to reference 9
            							
            						
            						
            							I will explain my thinking in another post later, and then edit this with a link to that post. 
            							
            								Back
            								 to reference 10
            							
            						
            						
            							Some search engines support the site: search operator to limit searches to subpages or subdomains of a single site or TLD. site:.one, for instance, limits searches to websites with the “.one” TLD. 
            							
            								Back
            								 to reference 11
            							
            						
            						
            							Matt from Gigablast told me that indexing YouTube or LinkedIn will get you blocked if you aren’t Google or Microsoft. I imagine that you could do so by getting special permission if you’re a megacorporation. 
            							
            								Back
            								 to reference 12
        comment:
          children:
            - type: cite
              published: 2021-03-11 19:00:19Z
              url: https://littr.me/~Seirdy/152df9ce-b78c-4dba-a6aa-79985052685a
              name: activitypub
            - type: cite
              published: 2021-03-12 05:20:22Z
              url: https://tildes.net/~tech/vnw/a_look_at_search_engines_with_their_own_indexes
              name: A look at search engines with their own indexes - ~tech - Tildes
            - type: cite
              published: 2021-03-12 07:16:06Z
              url: https://news.ycombinator.com/item?id=26429942
              name: A look at search engines with their own indexes | Hacker News
            - type: cite
              published: 2021-03-12 14:43:27Z
              url: https://adhoc.systems/boosts/d2183854-1bb6-49bf-9476-c0b5bdb54e5d
              name: "Boosted: A look at search engines with their own indexes - Seirdy"
            - type: cite
              published: 2021-03-18 17:42:57Z
              url: http://th3core.com/talk/traffic/list-of-search-engines-with-their-own-indexes-march-2021/
              name: List of Search Engines with their own Indexes March 2021
            - type: cite
              published: 2021-03-19 10:56:03Z
              url: https://indieseek.xyz/2021/03/19/new-search-engines-added-march-2021/
              name: New Search Engines Added March 2021
              author:
                type: card
                name: Brad
              content: I added some new search engines to the directory here.  What makes these worth listing is they all have their own index.  This makes them unique even though the index may currently be small.GoWiki – Small, growing, privacy.Petal Search – New commercial search engine from Huawei.  I think their index is fairly large.  I don’t know if they are using another search engine (maybe Yandex?) for backfill. Assume it is not private.Plumb One – Index is small, growing.  Plumb uses Bing to …
            - type: cite
              published: 2021-03-29 22:45:14Z
              url: https://www.lealternative.net/2019/10/18/alternative-a-google-search/
              name: Alternative a Google Search
              content: Siete stanchi delle solite ricerche e state cercando delle alternative a Google Search? La quantità di tracker e la capacità di profilazione di Google potrebbe allarmare molti anche perché oltre ad un evidente problema etico e morale nell’essere costantemente seguiti e profilati c’è anche il problema della cosiddetta ‘bolla culturale’. In breve significa che conoscendovi così bene Google vi proporrà principalmente cose che, secondo lui, a voi interessano evitandovi così di ampl…
            - type: cite
              published: 2021-04-02 16:44:07Z
              url: https://searchnews.org/2021/04/02/a-look-at-search-engines-with-their-own-indexes/
              name: A look at search engines with their own indexes
              content: |-
                Rohan Kumar has a list of search engines with their own index.  



                It’s a really useful list and it breaks it into several categories. Such as:



                Engines with their own crawling and indexing.A list of engines that leverage other engines’ indexes. A list of engines that have their own smaller indexes.Big and small indexes for other languages.



                These lists are hard to put together. It’s challenging to unwind what’s happening behind the scenes, but this document covers a lot of groun…
            - type: cite
              published: 2021-06-17 20:28:05Z
              url: https://thenewleafjournal.com/a-2021-list-of-alternative-search-engines-and-search-resources/
              name: Alternative Search Engine Review (2021) · The New Leaf Journal
            - type: cite
              published: 2021-08-16 21:31:36Z
              url: https://malwaretips.com/threads/the-best-private-search-engine-what-are-you-using.109387
              name: Q&A - The Best Private Search Engine - What are you using? | MalwareTips Community
            - type: cite
              published: 2021-10-31 04:24:57Z
              url: https://chemistryinthecity.neocities.org/content/entry2110.html#282
              name: Chemistry in the city
            - type: cite
              published: 2022-01-30 20:57:40Z
              url: https://news.ycombinator.com/item?id=29600619
              name: "Show HN: Searchall – search all major indexes on one page (with iframes) | Hacker News"
            - type: cite
              published: 2022-01-31 04:46:45Z
              url: https://www.ghacks.net/2021/09/17/firefox-experiment-is-testing-bing-as-the-default-search-engine/
              name: Firefox Experiment is testing Bing as the default search engine - gHacks Tech News
            - type: cite
              published: 2022-01-31 04:51:03Z
              url: https://i.reddit.com/r/NoStupidQuestions/comments/sdvmpy/is_it_just_me_or_has_googles_search_result/hufk9kl/
              name: i.reddit.com/r/NoStupidQuestions/co …
            - type: cite
              published: 2022-02-08 20:11:30Z
              url: https://blog.mojeek.com/2022/02/search-choices-enable-freedom-to-seek.html
              name: Search Choices Enable Freedom to Seek | Official Mojeek Blog
            - type: cite
              published: 2022-02-09 07:27:13Z
              url: https://ramblinggit.com/2021/03/16/bookmarked-here-is.html
              name: Brad Enslen
              author:
                type: card
                name: Brad Enslen
              content: |-
                Bookmarked:  Here is a good rundown of search engines that have their own index, including some new ones to me.
                seirdy.one/2021/03/1…
            - type: cite
              published: 2022-02-09 07:41:24Z
              url: https://www.androidauthority.com/da-february-2-2022-3100832/
              name: "Daily Authority: 📈 Starlink, Pixel sales - Android Authority"
            - type: cite
              published: 2022-02-14 00:19:41Z
              url: https://brid.gy/comment/reddit/Seirdy/m9hesi/grnu7uq
              name: For completeness, at Neeva we're also investing heavily to build out our web index and search capabilities. Beyond the public …
              author:
                type: card
                name: ahvee_at_neeva
              content: For completeness, at Neeva we're also investing heavily to build out our web index and search capabilities. Beyond the public web we already index connected accounts like Dropbox to safely search across personal docs in one place. Keep an eye out for NeevaBot and more if you're a webmaster. ;)
            - type: cite
              published: 2022-02-14 00:27:48Z
              url: https://brid.gy/comment/reddit/Seirdy/m9hesi/grna51w
              name: It's interesting, somehow I (exclusively qualitatively) perceive that Qwant gives me better results than Bing/DDG, but after …
              author:
                type: card
                name: PepperJackson
              content: It's interesting, somehow I (exclusively qualitatively) perceive that Qwant gives me better results than Bing/DDG, but after reading this I put them side by side and it looks like they are similar than I thought! I guess I feel more comfortable using a search engine that's based out of the US though.
            - type: cite
              published: 2022-02-17 12:43:39Z
              url: https://pleroma.envs.net/notice/AGYgEgNubK6AhW7n5U
              name: brid.gy/like/mastodon/@Seirdy@plero …
              author:
                type: card
                name: Mojeek
            - type: cite
              published: 2022-02-23 03:46:13Z
              url: https://micro.blog/bradenslen/11175648
              name: "@odd I would really like for Apple to develop a search engine.  It might take a few years for them to make it really work but I …"
              author:
                type: card
                name: bradenslen
              content: |-
                @odd I would really like for Apple to develop a search engine.  It might take a few years for them to make it really work but I think it would be good for the web.

                Same here, I really like DuckDuckGo and Mojeek.com and use those two almost exclusively.
            - type: cite
              published: 2022-02-23 20:29:42Z
              url: https://brid.gy/post/twitter/seirdy/1431092779094982658
              name: This is a really good look at alternative search engines. Something I have been meaning to write for a while but never got …
              author:
                type: card
                name: boyter
              content: This is a really good look at alternative search engines. Something I have been meaning to write for a while but never got around to doing. seirdy.one/2021/03/10/sea…
            - type: cite
              published: 2022-02-23 20:39:27Z
              url: https://brid.gy/post/twitter/seirdy/1487481854005493763
              name: This is a good read. Im noticing changes on my search results.  A look at search engines with their own indexes - Seirdy seirdy.one/2021/03/10/sea…
              author:
                type: card
                name: Rae
              content: This is a good read. Im noticing changes on my search results.  A look at search engines with their own indexes - Seirdy seirdy.one/2021/03/10/sea…
            - type: cite
              published: 2022-02-26 04:52:17Z
              url: https://www.jayeless.net/wiki/search-engine.html
              name: search engine
              content: |-
                Search engines work by building a database of what pages exist on the web (which it does by sending out “spiders” to “crawl” pages, following all the links they find to other pages), analysing them in some way, and using an algorithm to produce what it thinks are the most relevant results for any search term you type in.
                For English-language sites, the two biggest search engines, which have the resources to do their own crawling, are Google and Microsoft’s Bing. These seem to be the…
            - type: cite
              published: 2022-02-27 19:24:33Z
              url: https://news.ycombinator.com/item?id=30352345
              name: I concur. And would add that on Safari and iOS, it suits Google and Microsoft to... | Hacker News
            - type: cite
              published: 2022-03-03 22:51:18Z
              url: https://brid.gy/post/twitter/seirdy/1499441860573642756
              name: |-
                Much more than a cursory review:
                seirdy.one/2021/03/10/sea…
              author:
                type: card
                name: SearchEngineMap
              content: |-
                Much more than a cursory review:
                seirdy.one/2021/03/10/sea…
            - type: cite
              published: 2022-03-11 00:22:53Z
              url: https://pleroma.envs.net/notice/AHHdovaz9VPSieLIiu
              name: "@Seirdy Surely it's possible to make a search engine that's only biased towards relevant content."
              author:
                type: card
                name: Hyolobrika
              content: Surely it's possible to make a search engine that's only biased towards relevant content.
            - type: cite
              published: 2022-03-15 22:02:52Z
              url: https://pleroma.envs.net/notice/AHRtzfm49UzpQIdRRY
              name: "Excellent read by @Seirdy A look at search engines with their own indexes: …"
              author:
                type: card
                name: "bbbhltz :debian: :thinkpad_tp:"
              content: |-
                Excellent read by @Seirdy A look at search engines with their own indexes: https://seirdy.one/2021/03/10/search-engines-with-own-indexes.htmlMake sure to take a moment to read through some of the references too
                A look at search engines with their own indexes
            - type: cite
              published: 2022-03-17 18:34:19Z
              url: https://swprs.org/advanced-online-media-use/
              name: Advanced Online Media Use
            - type: cite
              published: 2022-03-18 18:59:19Z
              url: https://en.wikipedia.org/wiki/Wikipedia_talk:WikiProject_Free_Software
              name: Wikipedia talk:WikiProject Software/Free and open-source software task force - Wikipedia
            - type: cite
              published: 2022-03-19 03:35:47Z
              url: https://lobste.rs/s/szkcf0
              name: A look at search engines with their own indexes | Lobsters
              author:
                type: card
                name: Seirdy
            - type: cite
              published: 2022-03-19 19:19:08Z
              url: https://brid.gy/post/twitter/seirdy/1505255507946311680
              name: "A look at search engines with their own indexes: seirdy.one/2021/03/10/sea…"
              author:
                type: card
                name: Thomas Steiner
              content: "A look at search engines with their own indexes: seirdy.one/2021/03/10/sea…"
            - type: cite
              published: 2022-03-19 19:19:11Z
              url: https://brid.gy/post/twitter/seirdy/1505237444635021315
              name: "Search engine survey: Comprehensively broad collection of search engines. There's only really 3 for English: Google, Bing, and …"
              author:
                type: card
                name: Nelson Minar
              content: |-
                Search engine survey: Comprehensively broad collection of search engines. There's only really 3 for English: Google, Bing, and Yandex.
                seirdy.one/2021/03/10/sea…
                google bing search via:lobsters tootme
            - type: cite
              published: 2022-03-22 02:19:56Z
              url: https://ronan.jouchet.fr/2022/03/20/reel
              name: Ronan Jouchet - weekly reel
            - type: cite
              published: 2022-03-24 03:34:57Z
              url: https://english.nuovosito.lealternative.net/2022/03/23/we-want-you/
              name: Alternatives to Google Search – Le Alternative | English
            - type: cite
              published: 2022-04-01 12:58:10Z
              url: https://brainbaking.com/post/2022/04/march-2022/
              name: March 2022 In Review
              author:
                type: card
                name: Wouter Groeneveld
              content: "March 2022 is no more. I’m starting to feel the same time anxiousness that Winnie Lim wrote about: it feels like yesterday writing “February 2022 is no more”, while this month has three more days than the previous one! Looking back at the March blog posts, they were mostly about music: creating your own streaming server, choosing the right codec, an introduction to hip-hop. I recently extended my self-hosted Docker images with Gitea (and moved a few low-impact and private repositories o…"
            - type: cite
              published: 2022-04-04 05:16:46Z
              url: https://forums.puri.sm/t/duckduckgo-censorship/16638/102
              name: Duckduckgo censorship - General security & privacy chat - Purism community
            - type: cite
              published: 2022-04-05 03:57:24Z
              url: https://ehret.me/news-from-last-month-202204-edition
              name: News from last month (2022/04 edition) [ehret.me]
            - type: cite
              published: 2022-04-17 15:47:33Z
              url: https://brid.gy/post/twitter/seirdy/1515716388752699397
              name: |-
                Did you know @Bing's index helps power these search engines?

                Yahoo! (& One­Search)
                DuckDuck­Go
                AOL
                Qwant (partial)
                Ecosia
                Ekoru …
              author:
                type: card
                name: Jon Henshaw
              content: |-
                Did you know @Bing's index helps power these search engines?

                Yahoo! (& One­Search)
                DuckDuck­Go
                AOL
                Qwant (partial)
                Ecosia
                Ekoru
                Privado
                Findx
                Disconnect Search
                PrivacyWall
                Lilo
                Search­Scene
                Peekier
                Oscobo
                Million Short
                Yippy search
                Lycos
                Givero
                & more...
                seirdy.one/2021/03/10/sea…
            - type: cite
              published: 2022-05-04 22:51:12Z
              url: https://forum.puppylinux.com/viewtopic.php?p=56606
              name: Regarding Search Engines - Puppy Linux Discussion Forum
            - type: cite
              published: 2022-05-04 23:05:19Z
              url: https://www.homesteadingtoday.com/threads/great-article-on-alternative-search-engines.621268/
              name: homesteadingtoday.com/threads/great …
            - type: cite
              published: 2022-05-16 14:53:42Z
              url: https://indieseek.xyz/2022/05/16/indieseek-xyz-directory-additions-may-2022/
              name: Indieseek.xyz Directory Additions May 2022
              author:
                type: card
                name: Brad
              content: "I have added quite a few new listings.  Some are submissions and some are editorial adds. Also added are some new categories:Amtrak: under Recreation > Travel  I added some useful guides for traveling on Amtrak rail.  Because we have to decrease our use of cars and airplanes for travel so be need to start figuring out how we can use Amtrak.Europe: under News  I found these useful and not overly commercial. I also added a couple of search engines to the existing category. Source for the…"
            - type: cite
              published: 2022-05-16 20:34:35Z
              url: https://news.ycombinator.com/item?id=31397912
              name: ">I noted that there had been multiple weeks where not one single real person had... | Hacker News"
            - type: cite
              published: 2022-05-20 03:02:44Z
              url: https://thenewleafjournal.com/a-2021-list-of-alternative-search-engines-and-search-resources/?replytocom=563
              name: Rohan Kumar mentioned this Post on seirdy.one.
              author:
                type: card
                name: Nicholas A. Ferrell
            - type: cite
              published: 2022-05-24 10:42:39Z
              url: https://brid.gy/post/twitter/seirdy/1528710772968275969
              name: Using one Search Engine (SE) for everything ignores the fact that different engines have different strengths. This report offers …
              author:
                type: card
                name: TRADECRAFT
              content: |-
                Using one Search Engine (SE) for everything ignores the fact that different engines have different strengths. This report offers a comprehensive overview of various SEs and the algorithms that power them. 
                #OSINT  #SearchEngines  #OnlineSearch

                seirdy.one/2021/03/10/sea…
            - type: cite
              published: 2022-05-28 11:28:24Z
              url: https://pleroma.envs.net/notice/AJu8Y1QyWLu9R0tO1w
              name: A large overview of search projects that build their own …
              author:
                type: card
                name: Humane Tech Now
              content: |-
                A large overview of search projects that build their own indexes..https://seirdy.one/2021/03/10/search-engines-with-own-indexes.html

                A look at search engines with their own indexes
            - type: cite
              published: 2022-05-28 23:26:42Z
              url: https://brid.gy/post/twitter/seirdy/1530661351030501376
              name: |-
                A whole review of search engines
                seirdy.one/posts/2021/03/…
              author:
                type: card
                name: ana.
              content: |-
                A whole review of search engines
                seirdy.one/posts/2021/03/…
            - type: cite
              published: 2022-06-03 04:22:20Z
              url: https://seirdy.one/notes/2022/06/02/duckduckgo-and-bing/
              name: DuckDuckGo and Bing
              author:
                type: card
                name: Rohan Kumar
              content: |-
                Reply to how would html.duckduckgo.com fit into this? by
                	@penryn@www.librepunk.club


                				I was referring to crawlers that build indexes for search engines to use. DuckDuckGo does have a crawler—DuckDuckBot—but it’s only used for fetching favicons and scraping certain sites for infoboxes (“instant answers”, the fancy widgets next to/above the classic link results).
                				DuckDuckGo and other engines that use Bing’s commercial API have contractual arrangements that typically inc…
            - type: cite
              published: 2022-06-14 14:38:25Z
              url: https://coywolf.social/@henshaw/108475945251308870
              name: 'mastodon:: "@Seirdy@pleroma.envs.net if this is the workaroun…" - Coywolf Social'
              author:
                type: card
                name: "Jon Henshaw :mastodon:"
              content: if this is the workaround that Neeva had to come up with that took them an inordinate amount of time and resources, they might need to hire all new engineers 😆> This forces startups to spend inordinate amounts of time and resources coming up with workarounds. For example, Neeva implements a policy of “crawling a site so long as the robots.txt allows GoogleBot and does not specifically disallow Neevabot.”
            - type: cite
              published: 2022-06-14 14:38:51Z
              url: https://social.coop/@cstanhope/108475963220607564
              name: '"Not that Google would go…" - social.coop'
              author:
                type: card
                name: Charles Stanhope
              content: "Not that Google would go for this, but maybe instead of search neutrality we need collective crawling? One co-op crawling body with neutral access terms for that data.And perhaps this could be coupled with a push model instead of a pull model. Websites that want to be included in search can submit new or changed URLs once a day for the crawler to pick up.Then search engines can optimize access to their pool of data without burdening sites.This has been my daily thonk. :thonking:"
            - type: cite
              published: 2022-06-21 13:40:09Z
              url: https://lobste.rs/s/rpyuic
              name: A look at search engines with their own indexes | Lobsters
              author:
                type: card
                name: h0p3fu1
            - type: cite
              published: 2022-06-21 14:10:32Z
              url: https://brid.gy/post/twitter/seirdy/1539243119937560577
              name: Starts off being a review of search engines, ends up almost a thesis. seirdy.one/posts/2021/03/…
              author:
                type: card
                name: "#WikiParty@michaelgraaf@campaign.openworlds.info"
              content: Starts off being a review of search engines, ends up almost a thesis. seirdy.one/posts/2021/03/…
            - type: cite
              published: 2022-06-21 19:50:20Z
              url: https://news.ycombinator.com/item?id=31820149
              name: A look at search engines with their own indexes (2021) | Hacker News
            - type: cite
              published: 2022-06-21 21:28:09Z
              url: https://arnoldit.com/wordpress/2022/06/21/search-the-web-maybe-find-a-nugget-or-two-for-intrepid-researchers/
              name: "Search the Web: Maybe Find a Nugget or Two for Intrepid Researchers? : Stephen E. Arnold @ Beyond Search"
            - type: cite
              published: 2022-06-23 16:41:19Z
              url: https://blog.gslin.org/archives/2022/06/24/10754/%e6%90%9c%e5%b0%8b%e5%bc%95%e6%93%8e%e7%9a%84%e6%9b%bf%e4%bb%a3%e6%96%b9%e6%a1%88%e6%b8%85%e5%96%ae/
              name: 搜尋引擎的替代方案清單
              author:
                type: card
                name: Gea-Suan Lin
              content: |-
                看到「A look at search engines with their own indexes」這篇在介紹各個搜尋引擎，作者設計了一套方法測試，另外在文章裡面也給了很多主觀的意見，算是很有參考價值的，可以試看看裡面提出來的建議。
                另外在 Hacker News 上也有討論可以參考：「A look at search engines with their own indexes (2021) (seirdy.one)」。
                在文章開頭的「General indexing search-engines」這個章節，先列出三大搜尋引擎 GBY (Google�…
            - type: cite
              published: 2022-06-25 05:44:49Z
              url: https://www.pointer.io/archives/5e342f0e7d/
              name: "Issue #328"
            - type: cite
              published: 2022-06-26 09:31:42Z
              url: https://the.dailywebthing.com/an-exhaustive-list-of-search-engines-out-there/
              name: an exhaustive list of search engines out there
              author:
                type: card
                name: joe jenett
            - type: cite
              published: 2022-06-27 00:23:13Z
              url: https://brid.gy/post/twitter/seirdy/1541213834215452672
              name: |-
                本文作者 Rohan Kumar 發現，搜尋引擎使用的索引很集中，以英語世界來說，大多來自「GBY」（Google, Bing 和 Yandex），他試圖整理各家搜尋引擎和索引，按照自行設計的研究方法去分類。

                seirdy.one/posts/2021/03/…
              author:
                type: card
                name: 三創育成 Star Rocket
              content: |-
                本文作者 Rohan Kumar 發現，搜尋引擎使用的索引很集中，以英語世界來說，大多來自「GBY」（Google, Bing 和 Yandex），他試圖整理各家搜尋引擎和索引，按照自行設計的研究方法去分類。

                seirdy.one/posts/2021/03/…
            - type: cite
              published: 2022-06-27 09:32:28Z
              url: https://brid.gy/post/twitter/seirdy/1541349873034960896
              name: Un article terriblement complet sur les moteurs de recherche actuellement existants. Ca donne un petit vertige ... seirdy.one/posts/2021/03/…
              author:
                type: card
                name: Nicolas Delsaux
              content: Un article terriblement complet sur les moteurs de recherche actuellement existants. Ca donne un petit vertige ... seirdy.one/posts/2021/03/…
            - type: cite
              published: 2022-06-28 13:24:14Z
              url: https://brid.gy/post/twitter/seirdy/1539151770622087169
              name: |-
                Different search engines and their primary searching index.
                - Good to be aware of different search indexes that exist.
                - Having …
              author:
                type: card
                name: Tech Article A Day
              content: |-
                Different search engines and their primary searching index. 
                - Good to be aware of different search indexes that exist.
                - Having diff search engines is helpful to discover different results on different criteria. 

                seirdy.one/posts/2021/03/…
            - type: cite
              published: 2022-06-29 18:01:13Z
              url: https://blog.revi.xyz/daum-naver-search-independence
              name: Daum and Naver | revi
            - type: cite
              published: 2022-06-30 16:28:34Z
              url: https://indieseek.xyz/2022/06/30/new-search-engines-added-june-2022/
              name: New Search Engines Added June 2022
              author:
                type: card
                name: Brad
              content: The elves that build search engines are really kicking into high gear, and I’m loving it.  If you have been following my writings for awhile you know that I want a Web with 5 to 8 major search engines and bunches of smaller search engines and hundreds of directories.  The Web is being harmed by having a Google and Bing duopoly controlling the gateways to the Web.  So I’m rooting for all search engines with their own indexes and crawlers and I’m always on the lookout for new contender…
            - type: cite
              published: 2022-07-08 02:57:56Z
              url: https://github.com/SerenityOS/serenity/pull/14504#issuecomment-1177510554
              name: "Browser: Add Brave and Mojeek to search engines by Xexxa · Pull Request #14504 · SerenityOS/serenity · GitHub"
            - type: cite
              published: 2022-07-12 03:09:16Z
              url: https://blog.rowan.earth/divorcing-google-part-2/
              name: Divorcing Google, Part 2 | It From Bit
            - type: cite
              published: 2022-07-25 01:12:49Z
              url: https://brid.gy/post/twitter/seirdy/1551370857322139648
              name: |-
                “What I learned by building this list has profoundly changed how I surf.

                “Using one engine for everything ignores the fact that …
              author:
                type: card
                name: delan
              content: |-
                “What I learned by building this list has profoundly changed how I surf.

                “Using one engine for everything ignores the fact that different engines have different strengths.”

                seirdy.one/posts/2021/03/…
            - type: cite
              published: 2022-08-05 06:47:31Z
              url: https://ramblinggit.com/2022/07/29/i-still-want.html
              name: Brad Enslen
              author:
                type: card
                name: Brad Enslen
              content: I still want at least 8 major English language web search engines with their own indexes.  At least there are contenders forming now. Seirdy has the best rundown.
            - type: cite
              published: 2022-08-06 19:33:10Z
              url: https://arstechnica.com/gadgets/2022/08/microsoft-trackers-run-afoul-of-duckduckgo-get-added-to-blocklist/?comments=1&post=41130314#comment-41130314
              name: Microsoft trackers run afoul of DuckDuckGo, get added to blocklist | Ars Technica
            - type: cite
              published: 2022-08-07 09:30:16Z
              url: https://thenewleafjournal.com/importance-of-bing-indexing-for-alt-search/
              name: Importance of Bing Indexing For Alt Search · The New Leaf Journal
            - type: cite
              published: 2022-08-07 19:50:49Z
              url: https://matienzo.org/2022/175/searchengines/
              name: 🔖 A look at search engines with their own indexes - Seirdy | M.A. Matienzo
              author:
                type: card
                name: M.A. Matienzo
            - type: cite
              published: 2022-08-10 19:24:44Z
              url: https://cyberzettel.com/522-2/
              name: cyberzettel.com/522-2
              author:
                type: card
                name: Brad
              content: |-
                Bookmark: A look at search engines with their own indexes by Seirdy.
                 
                The three dominant English search engines with their own indexesnote 1 are Google, Bing, and Yandex (GBY). Many alternatives to GBY exist, but almost none of them have their own results; instead, they just source their results from GBY.
                With that in mind, I decided to test and catalog all the different indexing search engines I could find. I prioritized breadth over depth, and encourage readers to try the engines out th…
            - type: cite
              published: 2022-09-08 19:27:50Z
              url: https://greycoder.com/mojeek-a-web-search-thats-private-and-independently-indexed/
              name: Mojeek — A Private Web Search That Uses An Independent Index
              author:
                type: card
                name: web master
              content: |-
                Mojeek is a privacy-friendly search engine based in the UK. Mojeek has created its own independent web search index. The index covers 5 billion pages (as of March 2022). This is impressive and makes Mojeek a unique search alternative because doesn’t base its search results on one of the three big search indexes (Google, Bing, and Yandex).



                Most alternative search engines get their organic results primarily from Bing. This includes search engines like DuckDuckGo, Qwant, Brave Search, Ecosi…
            - type: cite
              published: 2022-09-20 14:23:44Z
              url: https://mastodon.nz/@a/109019295322796955
              name: "Austin Huang ❤ You All: \"@freakazoid@retro.social Well, there aren't many …\" - Mastodon NZ"
              author:
                type: card
                name: Austin Huang ❤ You All
              content: "@freakazoid Well, there aren't many good search *indexes* available (given computation power required), so it's really just either Google or Microsoft. See https://seirdy.one/posts/2021/03/10/search-engines-with-own-indexes/ ."
            - type: cite
              published: 2022-09-24 17:29:24Z
              url: https://number1.co.za/alternate-search-engines/
              name: Alternate Search Engines
              content: |-
                Good

                gigablast.com – results aren’t from some generic source. Bitchute pages were in results. Still some pushing to specific directions.
                private.sh – Seems ok.
                neeva.com – Search Free From Corporate Influence. Decent results most times.

                OK

                wiby.me – only indexes old school websites, results not relevant at all.
                petalsearch – The chinese / huawei search engine – actually gives relevant results when you type in controversial search terms. Sites recommended include 4channel.org.…
            - type: cite
              published: 2022-10-02 00:45:37Z
              url: https://brid.gy/post/reddit/Seirdy/xswzzr
              name: A look at search engines with their own indexes
              author:
                type: card
                name: VapidCanary
            - type: cite
              published: 2022-10-02 00:46:05Z
              url: https://brid.gy/post/reddit/Seirdy/xsx3hq
              name: A look at search engines with their own indexes
              author:
                type: card
                name: VapidCanary
            - type: cite
              published: 2022-10-07 16:30:50Z
              url: https://discuss.privacyguides.org/t/a-look-at-search-engines-with-their-own-indexes/175
              name: '"A look at search engines with their own indexes" - Privacy - Privacy Guides'
            - type: cite
              published: 2022-10-10 16:01:48Z
              url: https://blog.starrocket.io/posts/newsletter-2022-06-23/
              name: "科技創業週報 #341：誰來告訴我，軟體工程師沒有受到詛咒？ | Star Rocket Blog"
            - type: cite
              published: 2022-10-11 10:23:24Z
              url: https://brid.gy/comment/twitter/seirdy/1579776851618455552/1579776851618455552
              name: |-
                De plus en plus, il faut tester les moteurs alternatifs à Google, Bing et Yandex :
                seirdy.one/posts/2021/03/…
                Dont Brave Search …
              author:
                type: card
                name: precisement.org
              content: |-
                De plus en plus, il faut tester les moteurs alternatifs à Google, Bing et Yandex : 
                seirdy.one/posts/2021/03/…
                Dont Brave Search : 
                precisement.org/blog/Brave-Sea…
            - type: cite
              published: 2022-11-02 20:39:11Z
              url: https://mdcdy.be/blog.html
              name: Liens
            - type: cite
              published: 2022-11-14 01:40:12Z
              url: https://stacker.news/items/37803
              name: Any opinions on mojeek.com? (privacy-focused search engine) \ stacker news
            - type: cite
              published: 2022-11-27 23:18:51Z
              url: https://brid.gy/post/twitter/seirdy/1596979417083174912
              name: |-
                A new mysearch user published an index on Yessle mysearch. break information bubble take a look. yessle.com/mysearch/frind…

                I …
              author:
                type: card
                name: Wadim Seminsky
              content: |-
                A new mysearch user published an index on Yessle mysearch. break information bubble take a look. yessle.com/mysearch/frind…

                I found this : seirdy.one/posts/2021/03/…
            - type: cite
              published: 2022-12-09 10:40:14Z
              url: https://brid.gy/comment/twitter/seirdy/1601133928990982144/1601133928990982144
              name: Thanks. Duck is pretty much all Bing despite what the comments say. seirdy.one/posts/2021/03/…
              author:
                type: card
                name: Colin Hayhurst
              content: Thanks. Duck is pretty much all Bing despite what the comments say. seirdy.one/posts/2021/03/…
            - type: cite
              published: 2022-12-13 19:13:08Z
              url: https://pleroma.envs.net/notice/AQZbDQyCpLcmowMUzI
              name: '@clovis Thanks - this is necessary. I am also building my own index, but have not made my instance public. Unlike "Linux Review" …'
              author:
                type: card
                name: radiocron
              content: '@clovis Thanks - this is necessary. I am also building my own index, but have not made my instance public. Unlike "Linux Review" I had good luck with searX as a front end, because now Yacy has improved somewhat, especially if you curate. When I consolidate my crawls, I will share a searX site that uses them - in addition to gigablast and others.'
            - type: cite
              published: 2022-12-13 19:13:14Z
              url: https://pleroma.envs.net/notice/AQZbDQyCpLcmowMUzI
              name: "@clovis THIS project. If you could combine this with drupal, it would rock! I don't think it would be terribly difficult. It …"
              author:
                type: card
                name: radiocron
              content: "@clovis THIS project. If you could combine this with drupal, it would rock! I don't think it would be terribly difficult. It would take some time, but would have users quickly - especially our news site. https://www.deepl.com/pro#developer"
            - type: cite
              published: 2022-12-16 16:58:45Z
              url: https://brid.gy/comment/twitter/seirdy/1603750994093182977/1603750994093182977
              name: |-
                I beg to differ on "real". But don't take it from me:
                seirdy.one/posts/2021/03/…
              author:
                type: card
                name: Colin Hayhurst
              content: |-
                I beg to differ on "real". But don't take it from me:
                seirdy.one/posts/2021/03/…
            - type: cite
              published: 2023-01-22 12:27:41Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/ARtTkcfEU5DOGPtL5U/ARtu2f7wf9SDvKaHlQ
              name: "@Seirdy wow, this was really thought out! thanks for the read, i didnt expect to recieve such nuanced replies when making the …"
              author:
                type: card
                name: jo!
              content: wow, this was really thought out! thanks for the read, i didnt expect to recieve such nuanced replies when making the pollwhat do you think about brave search's goggles? i didnt use brave search for long but i thought these were pretty neat for more specialised search (not having to add "rust lang" every time i look up some docs on smth is nice)one thing i tend to miss about large engines like google is the large amount of instant results like calculators which are often better/cleane…
            - type: cite
              published: 2023-01-23 19:18:20Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/ARwZ2xqpX4FKQplGaG/ARwZCdzQEPd1jEObHE
              name: "@Seirdy ooo cool! tysm for that list :3"
              author:
                type: card
                name: arkywarky
            - type: cite
              published: 2023-01-31 11:19:19Z
              url: https://thenewleafjournal.com/letter/120/
              name: Newsletter Leaf Journal CXX · The New Leaf Journal
            - type: cite
              published: 2023-02-08 14:18:46Z
              url: https://brid.gy/post/twitter/seirdy/1623230167794475008
              name: Looking forward to playing around with New Search™️ from the makers of Windows™️ and The World's Biggest Advertising Company™️. …
              author:
                type: card
                name: Richard Young
              content: Looking forward to playing around with New Search™️ from the makers of Windows™️ and The World's Biggest Advertising Company™️. But then someone pointed me at this, which is an interesting list seirdy.one/posts/2021/03/…
            - type: cite
              published: 2023-02-10 06:41:10Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/ASSlKqMHLv33OTSyzQ/ASSlKqMHLv33OTSyzQ
              name: "@aurynn @RichardYoung Most decent English generalist engines are just Bing with their own take on infoboxes: …"
              author:
                type: card
                name: Seirdy
              content: |-
                @aurynn @RichardYoung Most decent English generalist engines are just Bing with their own take on infoboxes: https://seirdy.one/posts/2021/03/10/search-engines-with-own-indexes/I like the idea of building multiple engines from a single common index such as the Common Crawl. Alexandria is one such example.
                A look at search engines with their own indexes
            - type: cite
              published: 2023-02-10 06:41:16Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/ASSlKqMHLv33OTSyzQ/ASSmK0d3mTbyucTvbk
              name: "@Seirdy @aurynn Mmmm - useful, thank you! DDG solved the Google enshittification problem for me, but better under-the-hood …"
              author:
                type: card
                name: Richard Young
              content: "@aurynn Mmmm - useful, thank you! DDG solved the Google enshittification problem for me, but better under-the-hood performance always welcome.(That said, I'm actually looking forward to experimenting with the AI Bing, even if my long-tem goal is just a nice, stripped back search experience that finds what I'm looking for... like Google did in the early days…)"
            - type: cite
              published: 2023-02-10 06:41:34Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/ASSlKqMHLv33OTSyzQ/AST8wbQlbDTz2wXW9w
              name: "@Seirdy @aurynn @RichardYoung very useful and enlightening list."
              author:
                type: card
                name: Firoozye
              content: "@aurynn @RichardYoung very useful and enlightening list."
            - type: cite
              published: 2023-02-10 09:35:13Z
              url: https://brid.gy/post/twitter/seirdy/1623977859092996100
              name: |-
                Search engines and the illusion of choice: A 🧵
                I've found this research project that I've been looking for since May. It raises …
              author:
                type: card
                name: Max
              content: |-
                Search engines and the illusion of choice: A 🧵
                I've found this research project that I've been looking for since May. It raises a number of EXTREMELY interesting and genuinely shocking points about search engines, specifically "alternative" ones.
                seirdy.one/posts/2021/03/…
                1/?
            - type: cite
              published: 2023-05-25 04:08:05Z
              url: https://community.mojeek.com/t/seach-index-map-updates/576/2?u=seirdy
              name: "Seach index map updates - #2 by Josh - The Mojeek Discourse"
            - type: cite
              published: 2023-07-01 19:12:16Z
              url: https://brid.gy/post/mastodon/@seirdy@pleroma.envs.net/AXFm7DA66Z9YDEueiu
              name: "@Skolliagh Qwant utilise Bing.Les moteurs utilisant leur propre index sont peu nombreux:- Google- Bing - Yandex- Petal Search …"
              author:
                type: card
                name: 9x0rg
              content: |-
                @Skolliagh Qwant utilise Bing.Les moteurs utilisant leur propre index sont peu nombreux:- Google- Bing - Yandex- Petal Search (Huawei)et...- MojeekPapier très détaillé à ce sujet:**A look at search engines with their own indexes**https://seirdy.one/posts/2021/03/10/search-engines-with-own-indexes/Thanks @Seirdy for your very detailed post BTW!
                A look at search engines with their own indexes
            - type: cite
              published: 2023-07-06 21:58:37Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AXGw4Grb4yupWXSKeG/AXGw4Grb4yupWXSKeG
              name: "@CobaltVelvet @haskal Ok, sampled from my blog post on the topic:Lixia Labs: brand new one that focuses on technical content and …"
              author:
                type: card
                name: Seirdy
              content: "@CobaltVelvet @haskal Ok, sampled from my blog post on the topic:Lixia Labs: brand new one that focuses on technical content and blogs.Wiby: mostly user-submitted Web-1.0 pages reminiscent of the “Old Web”.Marginalia Search: My favorite search engine. Penalizes pages with heavy JS, excessive harmful SEO practices, tracking, etc. and uses a weighted PageRank algorithm to prioritize better, simpler pages. Offers an optional ranking algorithm that prioritizes blogs.Kagi offers a “noncommer…"
            - type: cite
              published: 2023-07-15 19:17:30Z
              url: https://brid.gy/post/mastodon/@seirdy@pleroma.envs.net/AWYCNTo82yYDbO67f6
              name: wanna read about search engines with their own indexes? this one is a good one 😉 …
              author:
                type: card
                name: Mojeek
              content: |-
                wanna read about search engines with their own indexes? this one is a good one 😉 https://seirdy.one/posts/2021/03/10/search-engines-with-own-indexes/
                A look at search engines with their own indexes
            - type: cite
              published: 2023-07-15 19:17:31Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AWYCNTo82yYDbO67f6/AWaI5BxAg4rP3auTqK
              name: "@Mojeek Fascinating, thank you. One of the best articles Iʼve seen on the subject."
              author:
                type: card
                name: Jack Yan (甄爵恩)
              content: "@Mojeek Fascinating, thank you. One of the best articles Iʼve seen on the subject."
            - type: cite
              published: 2023-07-16 00:38:37Z
              url: https://old.reddit.com/r/searchengines/comments/11bz92i/is_it_true_googles_search_effectiveness_is/jcg7v2q/?context=3
              name: sillietechie comments on Is it true Google's search effectiveness is declining?
            - type: cite
              published: 2023-08-04 00:34:00Z
              url: https://brid.gy/post/mastodon/@seirdy@pleroma.envs.net/AYMAmK4ZSo5nQBGfmC
              name: An overview of search engines (mainly English language ones for …
              author:
                type: card
                name: Strypey
              content: |-
                An overview of search engines (mainly English language ones for now):https://seirdy.one/posts/2021/03/10/search-engines-with-own-indexes/#search #WebSearch #SearchEngines#HatTip to @indieterminacy for the link in a matrix room.
                hattip
                search
                searchengines
                websearch




                A look at search engines with their own indexes
            - type: cite
              published: 2023-08-08 01:05:50Z
              url: https://www.bookandsword.com/2023/03/18/state-of-web-search/
              name: Does DuckDuckGo Want To Search the Web?
            - type: cite
              published: 2023-08-08 01:09:14Z
              url: https://www.moonspeaker.ca/FoundSubjects/about.html#search
              name: The Moonspeaker
            - type: cite
              published: 2023-08-16 02:38:39Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AYm2vlvJsnls2T6WsC/AYm4EMWcda94iuyrdQ
              name: "@Seirdy Oh, PSA on Naver, here's a big thread that have done of experience with Naver's atrocious conduct with Zepeto (which …"
              author:
                type: card
                name: |-
                  Keep dreaming 'bout a better world
                  You keep wishing for some clarity
                  Always hoping that a lightning
              content: Oh, PSA on Naver, here's a big thread that have done of experience with Naver's atrocious conduct with Zepeto (which they run).https://toots.hwl.li/@jase/110888486598292767Just end of day, tldr, they're a corporation like any other only interested in spitting out corporate jargon and rather tone police than actually do shit.Not that any worse than any other search giant though tbf as they're all shit and wouldn't be any different any of them having control of Zepeto tbf. But yea, Nave…
            - type: cite
              published: 2023-08-16 06:33:08Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AYmNVuI0FpdRUeDpq4/AYmOnGgwJgxZqjjDm4
              name: Nice @Seirdy, Stract seems really cool!
              author:
                type: card
                name: McSinyx
            - type: cite
              published: 2023-08-16 06:33:10Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AYmNVuI0FpdRUeDpq4/AYmPJ3ZecfGYQuTjYu
              name: "@Seirdy optic is a great idea, i've wanted a search engine that has specific categories you can search in for a good while"
              author:
                type: card
                name: "'Tired of people's shit', aka Ember"
              content: optic is a great idea, i've wanted a search engine that has specific categories you can search in for a good while
            - type: cite
              published: 2023-08-16 06:33:10Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AYmNVuI0FpdRUeDpq4/AYmP0nkUinK48iqPIW
              name: "@Seirdy ooh, stract looks neat"
              author:
                type: card
                name: "'Tired of people's shit', aka Ember"
            - type: cite
              published: 2023-08-16 06:33:12Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AYmNVuI0FpdRUeDpq4/AYmPaYToy6HvD15tbc
              name: "@Seirdy \"Chat with an AI assistant that searches the web for you and cites its sources.\"blobfoxlurkglare​ that's sus tho"
              author:
                type: card
                name: "'Tired of people's shit', aka Ember"
              content: "\"Chat with an AI assistant that searches the web for you and cites its sources.\"blobfoxlurkglare​ that's sus tho"
            - type: cite
              published: 2023-08-16 06:33:13Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AYmNVuI0FpdRUeDpq4/AYmPiB79j0Cr3h7yRk
              name: "@Ember Kagi is somewhat similar but lacks detailed ranking metadata from third party sources, so personalized ranking …"
              author:
                type: card
                name: Seirdy
              content: "@Ember Kagi is somewhat similar but lacks detailed ranking metadata from third party sources, so personalized ranking adjustments just adjust a result’s position in a page rather than promoting results buried pages deep to the first page. Mojeek offers customization equivalent to adding a bunch of site: prefixes to your query, and might be working on something more that takes advantage of its own index.The homophobic cryptoshit one was leading in this area with its “goggles” feature off…"
            - type: cite
              published: 2023-08-16 06:54:27Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AYmNVuI0FpdRUeDpq4/AYmQBKhp8wE1ChBGyW
              name: "@Ember Marginalia has a few pre-set centrality algorithms centered around different types of pages, which is kind of similar. …"
              author:
                type: card
                name: Seirdy
              content: "@Ember Marginalia has a few pre-set centrality algorithms centered around different types of pages, which is kind of similar. You can select them in a drop-down menu when you search."
            - type: cite
              published: 2023-08-16 06:54:30Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AYmNVuI0FpdRUeDpq4/AYmQb7F3pAJayX1Ndg
              name: "@Seirdy aw, geez. I hadn’t heard about Gigablast. That’s too bad."
              author:
                type: card
                name: Jason Lefkowitz
              content: aw, geez. I hadn’t heard about Gigablast. That’s too bad.
            - type: cite
              published: 2023-08-16 07:30:20Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AYmNVuI0FpdRUeDpq4/AYmSGR4b64HKJ9OTho
              name: "@Seirdy It looks like https://www.slzii.com/about describes itself more as some sort of community/social network than a search …"
              author:
                type: card
                name: Tinyrabbit ✅
              content: It looks like https://www.slzii.com/about describes itself more as some sort of community/social network than a search engine. Does it really belong on the list?
            - type: cite
              published: 2023-08-16 08:27:39Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AYmNVuI0FpdRUeDpq4/AYma4Do3OcY7JtLIlk
              name: "@Seirdy slzii looks like yellow page, not search engine"
              author:
                type: card
                name: 𝟙𝔸
            - type: cite
              published: 2023-08-16 16:10:39Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AYmNVuI0FpdRUeDpq4/AYnD6ZeN351BrIvTVI
              name: "@iatendril It does appear to have a crawler and index for the included search engine so I included it for completeness."
              author:
                type: card
                name: Seirdy
              content: "@iatendril It does appear to have a crawler and index for the included search engine so I included it for completeness."
            - type: cite
              published: 2023-08-16 16:10:44Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AYmNVuI0FpdRUeDpq4/AYnEINBXJqvsd0d3rs
              name: "@Seirdy This is a great resource! Thank you."
              author:
                type: card
                name: jeremiah
            - type: cite
              published: 2023-09-03 00:46:38Z
              url: https://dolys.fr/forums/topic/searxng-ou-comment-avoir-son-propre-metamoteur-souverain/
              name: SearXNG où comment avoir son propre metamoteur souverain
              content: |-
                l’Almanet doLys Gnu/Linux – Open Source – Entreprises › Forums › L’almanet doLys Open Source › SearXNG où comment avoir son propre metamoteur souverain




                	Mots-clés : moteur de recherche
                	Ce sujet est vide.
                	
                	
                		Log In  Register  Lost Password 

                Affichage de 1 message (sur 1 au total)




                		



                	Auteur
                	Articles




                	
                		
                			


                	août 15, 2023 à 4:19  

                	
                	#12536

                	
                	
                	






                	
                	nam1962nam1962Maître des clés
                	
                	




                	
                	J’utilisais j…
            - type: cite
              published: 2023-09-07 21:38:15Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AZXIPhOQkkAKuQcqum/AZXIgdm9SJ22c3ncfI
              name: "@QuietMisdreavus @noracodes While HN has extremely “rich white neoliberal anglo techbro” biases, it does make for an excellent …"
              author:
                type: card
                name: Seirdy
              content: "@QuietMisdreavus @noracodes While HN has extremely “rich white neoliberal anglo techbro” biases, it does make for an excellent seed source for a search engine index (links to many places, easy to crawl, links are often to places that don’t do SEO so would be hard to find otherwise, etc.) so I sympathize with the decision to crawl it. But it shouldn’t receive so much disproportionate attention.I vastly prefer Marginalia’s approach of biasing its index around a few selected blogs and …"
            - type: cite
              published: 2023-09-13 21:27:24Z
              url: https://blog.adrianistan.eu/teletexto-012
              name: "Teletexto #012 - Adrianistán"
            - type: cite
              published: 2023-09-15 03:32:18Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AZmKp2Sgfh0oltE5om/AZmKuXdSCDEyu1t0RU
              name: Oh and anyone who markets their search engine as “unbiased” is lying. An unbiased engine is just biased towards whoever has the biggest SEO budget.
              author:
                type: card
                name: Seirdy
              content: Oh and anyone who markets their search engine as “unbiased” is lying. An unbiased engine is just biased towards whoever has the biggest SEO budget.
            - type: cite
              published: 2023-10-12 19:15:34Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AahQGqqvUi0qR4g0cy/AahQGqqvUi0qR4g0cy
              name: "@davidism i can personally recommend #kagi and have no experience but with ddg, brave, goohgle, but @Seirdy has a great list: …"
              author:
                type: card
                name: Nils Goroll 🤒
              content: |-
                @davidism i can personally recommend #kagi and have no experience but with ddg, brave, goohgle, but @Seirdy has a great list: https://seirdy.one/posts/2021/03/10/search-engines-with-own-indexes/
                kagi

                A look at search engines with their own indexes
            - type: cite
              published: 2023-10-20 15:54:52Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AaxxJRzWeZwusQ0DvU/AaxxJRzWeZwusQ0DvU
              name: "@lawlznet @Wolven @Seirdy did a writeup of different search engines with their own …"
              author:
                type: card
                name: "Pearlescent Ferret :flag_demigirl:"
              content: |-
                @lawlznet @Wolven @Seirdy did a writeup of different search engines with their own indexeshttps://seirdy.one/posts/2021/03/10/search-engines-with-own-indexes/
                A look at search engines with their own indexes
            - type: cite
              published: 2023-10-22 21:51:20Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/Ab2Z4cXVj9Dk8c5tvU/Ab2cJp9lOpeTTVvaNM
              name: "@Seirdy @cybertailor ty! Will be sure to check these out"
              author:
                type: card
                name: soweli Niko
            - type: cite
              published: 2023-11-14 02:23:11Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AbmZTasdxHRY8r2b8y/AbmaGlCYN0RCLFnXLU
              name: "@Seirdy I won’t be able to fit this into the piece I’m finalizing today but this is perfect for a follow-up that I’m writing …"
              author:
                type: card
                name: Mariya Delano
              content: I won’t be able to fit this into the piece I’m finalizing today but this is perfect for a follow-up that I’m writing later in the week!! Again thank you very much
            - type: cite
              published: 2023-11-22 01:38:18Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/Ac2vGtxl0CyF5691qi/Ac2vGtxl0CyF5691qi
              name: "@holly I used mojeek for about 10 minutes before switching back to DDG 😅@Seirdy did a writeup of search engines if you're …"
              author:
                type: card
                name: "Pearlescent Ferret :flag_demigirl:"
              content: |-
                @holly I used mojeek for about 10 minutes before switching back to DDG 😅@Seirdy did a writeup of search engines if you're interested.https://seirdy.one/posts/2021/03/10/search-engines-with-own-indexes/
                A look at search engines with their own indexes
            - type: cite
              published: 2023-11-22 01:38:21Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/Ac2vGtxl0CyF5691qi/Ac2xz7UIHKj5g4ZlyK
              name: "@PearlescentFerret @holly Kagi mixes Google, Mojeek, and Teclis (its own index) together so you can get Actually Different …"
              author:
                type: card
                name: Seirdy
              content: "@PearlescentFerret @holly Kagi mixes Google, Mojeek, and Teclis (its own index) together so you can get Actually Different results that are also Actually Good. But it costs money."
            - type: cite
              published: 2023-11-28 00:00:15Z
              url: https://www.marginalia.nu/marginalia-search/about/
              name: |-
                About Marginalia Search @
                marginalia.nu
            - type: cite
              published: 2023-12-08 07:58:24Z
              url: https://revi.blog/daum-naver-search-independence
              name: Cooking Pancakes · Daum and Naver
              content: |-
                Daum and Naver
                I've stumbled upon this article and found out...

                Daum: Korean. Also unsure about this one’s independence.

                While they are not really well known outside South Korea, it has Wikipedia article and is one of the first search engine in Korea. I don't have much insight on how their search index is run, but they run their own index with own webmaster tools, site registration site (Korean only), and robots.
                They once worked with Google (well, in 2000s) for their non-Korean index but…
            - type: cite
              published: 2023-12-13 01:38:16Z
              url: https://brid.gy/post/mastodon/@seirdy@pleroma.envs.net/AcbXV5Wg6aquwyDwfI
              name: I’m working on a “boutique search engine” write-up, so I got a little sources collection going.My favorite bookmark so far is …
              author:
                type: card
                name: Jason
              content: |-
                I’m working on a “boutique search engine” write-up, so I got a little sources collection going.My favorite bookmark so far is this list of search engines with their own indexes by Seirdy. The list is surprisingly comprehensive and wildly helpful. https://seirdy.one/posts/2021/03/10/search-engines-with-own-indexes/
                A look at search engines with their own indexes
            - type: cite
              published: 2023-12-13 09:36:17Z
              url: https://logicface.co.uk/rohan-kumar-on-search-engines-with-their-own-indexes/
              name: Rohan Kumar on search engines with their own indexes
              author:
                type: card
                name: lukealexdavis
              content: |-
                Google is the lord and emperor of search (regardless of what Microsoft Bing has going on with OpenAI). But its focus on commercially-slanted results means people looking for personal blogs or generally “rough around the edges” content will struggle to find it. That’s where alternative search engines come in, built on top of their own indexes.



                Rohan Kumar took a deep dive into a list of search engines with his own indexes including the main ones: Google, Bing, Yandex, as well as lesse…
            - type: cite
              published: 2023-12-13 09:37:16Z
              url: https://interregnum.ghost.io/are-you-still-googling-or-are-you-looking-for-more-a-call-for-anti-capitalist-language-use/
              name: Are you still Googling or are you looking for more? A call for anti-capitalist language use
            - type: cite
              published: 2024-02-08 14:08:25Z
              url: https://www.404media.co/this-guy-is-building-an-open-source-search-engine-in-real-time/
              name: 404media.co/this-guy-is-building-an …
            - type: cite
              published: 2024-02-13 03:34:08Z
              url: https://jackyan.com/blog/2024/02/google-fails-to-downrank-junk-sites-bing-and-mojeek-fare-better/
              name: Google fails to downrank junk sites; Bing and Mojeek fare better
              content: |-
                Understandably, I was bemused but then frustrated with the fake articles out there about the new Google algorithm update being created by me and named for me. (Even Perplexity has picked up this misinformation.) Of all those I found and reached out to, only one author—Shahid Jafar—had the decency to remove his article and apologize. 
                Shahid explained that he had found thousands of articles relating to this and based his on what he read. It’s easy to accept a sincere apology, which he ga…
            - type: cite
              published: 2024-02-22 05:31:02Z
              url: https://brid.gy/post/mastodon/@seirdy@pleroma.envs.net/Af7vy7ek5r9x6EGM3E
              name: "Today’s find: An amazing listing of existing search engines: those using #Google & Bing’s indexes & many more.“The 3 dominant …"
              author:
                type: card
                name: Jeri Dansky
              content: "Today’s find: An amazing listing of existing search engines: those using #Google & Bing’s indexes & many more.“The 3 dominant English search engines with their own indexes are Google, Bing, & Yandex (GBY). Many alternatives to GBY exist, but almost none of them have their own results; instead, they just source their results from GBY.“With that in mind, I decided to test & catalog all the different indexing search engines I could find.”https://seirdy.one/posts/2021/03/10/search-engin…"
            - type: cite
              published: 2024-02-28 05:12:20Z
              url: https://ask.metafilter.com/378482/Paid-search-engines-do-they-work-and-are-they-worth-it#5375246
              name: "Paid search engines: do they work and are they worth it? - google bing paidsearch | Ask MetaFilter"
            - type: cite
              published: 2024-02-28 05:48:06Z
              url: https://escapethealgorithm.substack.com/p/folk-search-engines
              name: Folk search engines - by Elan Ullendorff
            - type: cite
              published: 2024-02-28 06:03:27Z
              url: https://tybx.jp/2024/02/17/%e6%a4%9c%e7%b4%a2%e3%82%a8%e3%83%b3%e3%82%b8%e3%83%b3%e3%82%92%e3%82%81%e3%81%90%e3%82%8b%e5%86%92%e9%99%ba
              name: 検索エンジンをめぐる冒険
              author:
                type: card
                name: pinacol
            - type: cite
              published: 2024-02-28 20:40:52Z
              url: https://brainbaking.com/post/2024/01/how-to-search-the-internet/
              name: How To Search The Internet
              author:
                type: card
                name: Wouter Groeneveld
              content: |-
                Thanks to the multi billion dollar advertisement industry, searching for something on the internet has devolved from a joyous Altavista guess-the-keywords activity to a tiring chore where one has to wade through endless pools of generated SEO-optimized crap, hollow company blogs with more social media link embeds than actual content, and Reddit flame wars than ever before. In short: great stuff.
                Suppose you’re looking for a review of a video game. The first 100 hits will return the expected…
            - type: cite
              published: 2024-03-26 03:12:21Z
              url: https://github.com/webis-de/archive-query-log/issues/27
              name: "Add more search providers · Issue #27 · webis-de/archive-query-log · GitHub"
            - type: cite
              published: 2024-04-05 20:46:59Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AgZy6TRS3WKJYkOMjY/AgZy6TRS3WKJYkOMjY
              name: "@lori @omegaprobe @pluralistic there are small independent search engines  with their own indexes and some of them deprioritise …"
              author:
                type: card
                name: jarek 🇺🇦
              content: |-
                @lori @omegaprobe @pluralistic there are small independent search engines  with their own indexes and some of them deprioritise ad-driven and other shitty pages https://seirdy.one/posts/2021/03/10/search-engines-with-own-indexes/
                A look at search engines with their own indexes
            - type: cite
              published: 2024-04-05 20:46:59Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AgZy6TRS3WKJYkOMjY/AgZy6XTj30WS4xUEpk
              name: "@jarek @lori @omegaprobe @pluralistic Interesting.Might mean that I have to let the Ahrefs crawler back into my sites again, as …"
              author:
                type: card
                name: Angus McIntyre
              content: "@jarek @lori @omegaprobe @pluralistic Interesting.Might mean that I have to let the Ahrefs crawler back into my sites again, as it seems to be used to power the Yep search engine, which looks quite good.It's past time that we figured out a way to do an end-run around Google: they've just become too predatory and too useless. That's a bad combo."
            - type: cite
              published: 2024-04-05 20:47:00Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AgZy6TRS3WKJYkOMjY/AgZyvZqpz90GCsdN8C
              name: '@jarek I read this and was like "Oh damn I need to find this person and FOLLOW THEM" so I tracked them down and surprise! I already follow @Seirdy  🤣'
              author:
                type: card
                name: ResearchBuzz
              content: '@jarek I read this and was like "Oh damn I need to find this person and FOLLOW THEM" so I tracked them down and surprise! I already follow @Seirdy  🤣'
            - type: cite
              published: 2024-04-05 20:47:01Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AgZy6TRS3WKJYkOMjY/Aga4VlrLJewNe3wqrw
              name: "@Researchbuzz @Seirdy 😁👍🏾💪🏽"
              author:
                type: card
                name: jarek 🇺🇦
            - type: cite
              published: 2024-04-05 20:47:02Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AgZy6TRS3WKJYkOMjY/Aga4lDgU97TyHoEhwu
              name: "@researchbuzz @jarek Things I need to update when I get my laptop working again:Startpage uses Bing now, not Google.Some of …"
              author:
                type: card
                name: Seirdy
              content: "@researchbuzz @jarek Things I need to update when I get my laptop working again:Startpage uses Bing now, not Google.Some of these engines are dead.Mullvad has an engine for members that uses and aggressively caches Google’s commercial API."
            - type: cite
              published: 2024-04-05 21:55:28Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AgagnU0SEE4FkI5CO8/AgagnU0SEE4FkI5CO8
              name: "@rysiek @matt @Mojeek If you're interested in other options then @Seirdy did an excellent piece on the topic that pretty much …"
              author:
                type: card
                name: Jeremy Yap
              content: |-
                @rysiek @matt @Mojeek If you're interested in other options then @Seirdy did an excellent piece on the topic that pretty much covers what's out there for now:https://seirdy.one/posts/2021/03/10/search-engines-with-own-indexes/
                A look at search engines with their own indexes
            - type: cite
              published: 2024-04-05 21:55:28Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AgagnU0SEE4FkI5CO8/AgagtkKOSSYm6UObKq
              name: "@jeruyyap thank you, that is very useful!@matt @Mojeek @Seirdy"
              author:
                type: card
                name: Michał "rysiek" Woźniak · 🇺🇦
              content: "@jeruyyap thank you, that is very useful!@matt @Mojeek @Seirdy"
            - type: cite
              published: 2024-04-05 21:55:30Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AgagnU0SEE4FkI5CO8/Agagx6PYoEHbSkPfGa
              name: "@nus everyone gets to draw the line somewhere. I draw it there. You may draw it elsewhere."
              author:
                type: card
                name: Michał "rysiek" Woźniak · 🇺🇦
              content: "@nus everyone gets to draw the line somewhere. I draw it there. You may draw it elsewhere."
            - type: cite
              published: 2024-04-05 22:56:03Z
              url: https://proto.garden/blog/search_engines.html
              name: proto - A Brief Survey of Alternative Search Engines
            - type: cite
              published: 2024-04-06 15:29:08Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/Agaj8VbBhckIfrPzPM/Agaj8VbBhckIfrPzPM
              name: "@PersistentDreamer @crispius @Cal And, as luck would have it, @Seirdy  just updated their forever article A look at search …"
              author:
                type: card
                name: Amgine
              content: |-
                @PersistentDreamer @crispius @Cal And, as luck would have it, @Seirdy  just updated their forever article A look at search engines with their own indexes — very recommended.
                A look at search engines with their own indexes
            - type: cite
              published: 2024-04-06 15:29:10Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/Agaj8VbBhckIfrPzPM/Agb1bH6mSH7VJ13tx2
              name: "@Cal wow that's awful it's not only morally questionable but justbad"
              author:
                type: card
                name: cmake . -G halva
              content: "@Cal wow that's awful it's not only morally questionable but justbad"
            - type: cite
              published: 2024-04-06 15:29:13Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/Agaj8VbBhckIfrPzPM/AgcE1sB6e8AGPeUiWW
              name: "@amgine @PersistentDreamer @crispius @Cal @Seirdy Oh wow, this page is amazing. Thanks!"
              author:
                type: card
                name: Anthony Sorace
              content: "@amgine @PersistentDreamer @crispius @Cal @Seirdy Oh wow, this page is amazing. Thanks!"
            - type: cite
              published: 2024-04-09 05:00:34Z
              url: https://onethingnewsletter.substack.com/p/beyond-the-search-engine
              name: ⬜ Beyond the search engine - One Thing
            - type: cite
              published: 2024-04-14 23:22:18Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AgbvqyqphJCOkSfUS8/AgbvqyqphJCOkSfUS8
              name: I put this as a comment on the web page, but I'll repeat it here for people who aren't subscribed and might want this …
              author:
                type: card
                name: Anthony
              content: I put this as a comment on the web page, but I'll repeat it here for people who aren't subscribed and might want this information:Hate to say it but Kagi looks to me like it's going to full-on enshittify in a few years. Besides the fact that the Kagi people seem like libertarian techbros, this is a standard Silicon Valley startup. Their business model appears to be using your search queries to train their AI so they can sell it back to you (yes they take subscription fees but that's not where…
            - type: cite
              published: 2024-04-14 23:22:23Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AgbvqyqphJCOkSfUS8/Agcmvdx87Uwax3ipyy
              name: I can't thank you enough for sharing all these links! I kinda wished that Corey Doctorow did just a tad more research on the …
              author:
                type: card
                name: Robert Kingett, blind
              content: I can't thank you enough for sharing all these links! I kinda wished that Corey Doctorow did just a tad more research on the search engine before recommending it, but thank you for all these links! I bookmarked this so don't ever delete it! @abucci
            - type: cite
              published: 2024-04-14 23:23:42Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AgIX3c65uXixBzaioa/AgIX3c65uXixBzaioa
              name: |-
                @proto neat writeup! have you read seirdy's piece? https://seirdy.one/posts/2021/03/10/search-engines-with-own-indexes/
                A look …
              author:
                type: card
                name: "JJ :blobblackcat:"
              content: |-
                @proto neat writeup! have you read seirdy's piece? https://seirdy.one/posts/2021/03/10/search-engines-with-own-indexes/
                A look at search engines with their own indexes
            - type: cite
              published: 2024-04-15 09:56:30Z
              url: https://wiki.qunn.eu/books/b1d10/page/72b32#bkmrk-1.
              name: 综观自索引搜索引擎 | Grimoire
            - type: cite
              published: 2024-04-27 00:06:59Z
              url: https://buc.ci/abucci/p/1712157873.505318
              name: Anthony (@abucci@buc.ci)
            - type: cite
              published: 2024-05-01 15:39:34Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AhRBIKOBeEMQKJhpQG/AhRer9Q1KCkv6nWLGy
              name: "@strypey that is such a cool list! Good to know there are so many search engines out there, whatever their strengths and weaknesses 💪@Seirdy"
              author:
                type: card
                name: Hippo 🍉
              content: "@strypey that is such a cool list! Good to know there are so many search engines out there, whatever their strengths and weaknesses 💪@Seirdy"
            - type: cite
              published: 2024-05-01 15:39:34Z
              url: https://brid.gy/post/mastodon/@seirdy@pleroma.envs.net/AhRBIKOBeEMQKJhpQG
              name: '#TIL about #RightDao;"Independent, Uncensored, Private Search"... using its own index, not Goggle, Bing, or …'
              author:
                type: card
                name: Strypey
              content: |-
                #TIL about #RightDao;"Independent, Uncensored, Private Search"... using its own index, not Goggle, Bing, or Yandex:https://rightdao.com/#HatTip to @Seirdy for maintaining this page on web search engines;https://seirdy.one/posts/2021/03/10/search-engines-with-own-indexes/#search #WebSearch
                hattip
                rightdao
                search
                til
                websearch
            - type: cite
              published: 2024-05-01 15:40:10Z
              url: https://brid.gy/comment/mastodon/@seirdy@pleroma.envs.net/AhRBIKOBeEMQKJhpQG/AhRezRKScPiayOs9y4
              name: '@strypey "When talking to search engine founders, I found that the biggest obstacle to growing an index is getting blocked by …'
              author:
                type: card
                name: Hippo 🍉
              content: |-
                @strypey "When talking to search engine founders, I found that the biggest obstacle to growing an index is getting blocked by sites. Cloudflare is one of the worst offenders. Too many sites block perfectly well-behaved crawlers, only allowing major players like Googlebot, BingBot, and TwitterBot; this cements the current duopoly over English search and is harmful to the health of the Web as a whole."#Cloudfare is annoying! It sometimes blocks me too 🙄@Seirdy
                cloudfare
            - type: cite
              published: 2024-05-11 01:04:14Z
              url: https://raphael.computer/blog/openorb-curated-search-engine/
              name: OpenOrb - a curated RSS and Atom feed search engine - Raphael Kabo
            - type: cite
              published: 2024-05-21 23:33:19Z
              url: https://brid.gy/post/mastodon/@seirdy@pleroma.envs.net/Ai7QGjU8nRvzim8QZU
              name: "Talking about #SearchEngine, I can only recommend @Seirdy's synthesis work on « a look at search engines with their own indexes …"
              author:
                type: card
                name: DansLeRuSH ᴱᶰ
              content: |-
                Talking about #SearchEngine, I can only recommend @Seirdy's synthesis work on « a look at search engines with their own indexes » (Last updated 2024-05-11)› https://seirdy.one/posts/2021/03/10/search-engines-with-own-indexes/
                searchengine

                A look at search engines with their own indexes
        like:
          type: cite
          published: 2022-06-21 15:39:19Z
          author:
            type: card
            name: Tomáš Jakl
          url: https://nest.jakl.one/likes/2022-06-21-bzk8b/
        children:
          - type: cite
            author:
              type: card
              url: https://www.prologic.blog/
              name: James Mills
              givenName: James
              familyName: Mills
            name: So I'm a Knucklehead eh?
            url: https://www.prologic.blog/2021/02/14/so-im-a.html
          - type: card
            url: https://gigablast.com/bio.html
            name: Matt Wells
            givenName: Matt
            familyName: Wells
            org: Gigablast
          - type: cite
            name: A 2021 List of Alternative Search Engines and Search Resources
            url: https://thenewleafjournal.com/a-2021-list-of-alternative-search-engines-and-search-resources/
            author:
              type: card
              url: https://emucafe.club/channel/naferrell
              name: Nicholas Ferrell
              givenName: Nicholas
              familyName: Ferrell
              org: The New Leaf Journal
      - type: card
        url: https://seirdy.one/
        uid: https://seirdy.one/
        photo: https://seirdy.one/favicon.1250396055.png
        name: Rohan “Seirdy” Kumar
        givenName: Rohan
        additionalName: Seirdy
        familyName: Kumar
slug: slh0o
---
