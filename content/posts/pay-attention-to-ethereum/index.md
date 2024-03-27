---
title: "Why you should pay attention to Ethereum"
emoji: 🔷
date: 2017-05-21T18:53:19.558Z
draft: false
tags:
  - cryptocurrency
  - ethereum
  - money
summary: A brief explanation of cryptocurrencies and why they're useful I wrote for my non-techie friends.
syndications:
- https://medium.com/@jphastings/why-you-should-pay-attention-to-ethereum-96766a2c89a4
---

[Cryptocurrencies](https://en.wikipedia.org/wiki/Cryptocurrency) are [stealing](https://www.forbes.com/sites/laurashin/2017/05/20/as-cryptocurrency-markets-reach-new-highs-the-ethereal-summit-paints-a-rich-future) [the](https://fortune.com/2017/05/20/ethereum-ico-blockchain-token/) [limelight](https://www.ft.com/content/61cdc5c8-370e-11e7-bce4-9023f8c0fd2e) in the world of finance at the moment. They definitely deserve your attention but I wanted to voice some interesting implications of their success that bear levels of consideration that only getting thoughts down in words can provide!

The first section below goes through the basic concepts of digital currencies assuming as little prior knowledge as possible, and beneath is a more editorial look at why I think the complexities of systems like this are useful and bound for success.

## Foundational concepts

Before I start, I should give some background for readers who are less familiar with cryptocurrencies like Bitcoin and Ether and how they work. There are a host of articles you can find on every aspect I’ll cover, so I’ll try to keep it brief and let you dig deeper as you deem necessary.

### The premise

Cryptocurrencies are borne of the belief that some of the foundational issues that more traditional [fiat currencies](https://en.wikipedia.org/wiki/Fiat_money) struggle with can be solved with the application of suitably well thought-through technological processes.

There are a couple of concepts which we’ll need a basic understanding of to grasp the potential of these new financial mechanisms, but while they’re hideously complicated in the detail, they’re not too complex at the high level.

### Cryptography

As the name implies, cryptography is critical to cryptocurrencies’ function. Proving who you are is hard enough in physical space, let alone online and cryptography provides humanity’s best approach for this to date.

It’s a massive topic, as you might expect, but suffice to say here that, digitally, proof of identity is based on the principle that multiplying two extremely large numbers together is easy, but taking one extremely large number and finding the numbers that divide into it without a remainder is very hard (and takes a very long time, even for a computer).

To operate with cryptocurrencies, you start by generating two extremely large random ([prime](https://en.wikipedia.org/wiki/Prime_number)) numbers, or “keys”. The chances of you generating the same ones as someone else is vanishingly small as these numbers have many thousands of digits in them.

You publish one publicly, which forms your “address” (think: sort code and account number), and keep the other private as proof of ownership. For every financial action you take (e.g. “Take 5 from account X and put it in account Y”), you “sign” that instruction with your private key and append the instruction-signature pair to a shared public ledger called the _blockchain_.

The beauty of maths means that Joe Bloggs can use your public key and the attached signature to confirm that it was really you who ordered that transaction, but neither he nor anyone else without your private key can create a valid signature, and thus a valid instruction for that account’s money.

### Decentralisation

Having one public ledger for all these transactions seems well and good (as it prevents double spending) but if you imagine a physical book with people writing their transactions down in it, you rapidly see that there are location and [locking](https://en.wikipedia.org/wiki/Lock_%28computer_science%29) issues: you need to be where the book is, and only one person can write in it at once.

That book is the equivalent of late 20th century banking. Then we didn’t write in the ledger, we told our bank to write in the ledger for us. Each bank had its own book, and computers to write in them, each of which was very fast. This made centralised banking manageable, but coordinating a transaction from Alex in London to Brabra in Australia still took a long time as lots of banking authorities had to agree.

Decentralised banking, a core component of cryptocurrency construction, requires that _anyone_ can write in the ledger (even a thief or “bad actor”). A decentralised banking system is more like a series of noticeboards scattered around every city; when you want to make a transaction you just write it on a piece of paper and pin it up where anyone can look at it to see that it occurred. How is honesty and accuracy ensured? For that we require _the blockchain_.

### The Blockchain

Originally detailed by Satoshi Nakomoto, the still pseudonymous name chosen by the group or individual that penned the [whitepaper on Bitcoin](https://bitcoin.org/bitcoin.pdf), the blockchain is a way of structuring a ledger such that tampering with it or posting untruthfully is impossible by design, even when everyone has one and no-one owns a ‘master copy’.

The blockchain is an immutable list of parts of every transaction that has ever happened, organised in such a way that by reading it you can determine if the transaction you’re suspicious of is legitimate or not.

To continue the metaphor above, imagine a messenger that would regularly travel between every noticeboard, arrange all the new transactions in alphabetical order, and add the first letter of each transactional instruction to his notepad. Upon completing the last noticeboard he would add the first letter of every _notepad_ from each previous round-the-world trip he’d completed then add today’s notepad to the library and call that day’s job complete.

In this analogy the notepads are _blocks_ and the library of notepads he has created is _the blockchain_. In truth, rather than just taking the first letter of an instruction or a notepad, the blockchain uses a method called _[hashing](https://en.wikipedia.org/wiki/Hash_function)_ to turn arbitrary information into a smaller, fixed length piece of information. Hashing algorithms use every letter of an instruction as part of their calculation, so that nipping back to the noticeboard and adding an extra zero to the transfer order giving you your salary would completely change the _hash_ of your instruction, which would then differ from the messenger’s notepad, marking your forged transaction as provably incorrect.

In fact, there are many messengers performing this verification service and they corroborate their hashes against each other, to ensure consistency, before taking each notebook and adding it to the library that is the immutable ledger: the truth of the currency’s use.

On top of this, because the third notepad contains information from the second, and the second from the first, every subsequent round the messengers do further confirms and adds proof to all the transactions which came before. A concerned individual can wait for many rounds of proof before they agree that a transaction is truly verified and, of course, rather than happening once a day, these checks can occur in seconds.

This corroboration is the infrastructure that makes cryptocurrencies work and work requires incentive to be completed, so who does it and why do they bother?

### Mining

When people talk of “mining Bitcoin” or “mining Ether” what they mean is being the messenger of the analogy above. By inspecting transactions from around the world, hashing and arranging them in a specific way (such that completing that task is very hard, but proving that it has been completed is really easy), any individual can perform this task, and by performing it enough times a miner is rewarded with a currency coin. In fact, this is the _only_ way that new coins are introduced into circulation.

To control the rate of inflation some arbitrary complexity rules are imposed into the currency’s definition, for example: a miner might be required to order all the transactions in a block such that the last digits of the hash are 3832. This means that the miner must try millions of combinations of all of the transactions to find the order that produces a hash with the right four last digits (which takes time), but anyone can look at the result and know that the work was done at a glance. It is restrictions like these, combined with [Moore’s Law](https://en.wikipedia.org/wiki/Moore%27s_law) and some fancy maths, which ensure that the inflation of cryptocurrencies like Bitcoin and Ether are predictable, regular and limited.

---

There’s much more that could be covered in this background, but let’s get straight to the point.

## Why is all this complexity useful?

As I see it there are two foundational and relevant changes to the global post-millenium economy:

- Technology has progressed to the point where there are things consumers are willing to pay for that are just ones and zeroes
- The internet makes information transfer to any human effectively free

You may see where I’m going with this: trade is no longer bound by geography. When I use the internet to buy a Chinese-made light bulb, it still has to be shipped to me; if I want to buy a computer game on a physical disc, there needs to be a store for me to walk into; but if an engineer sells her software through the internet she could be on the moon and the only difference to her being in the next door room would be the 1.3 second delay in delivery time imposed by the speed of light. Clearly if you don’t need to rely on a physical store or distribution network your business is simpler and cheaper to run, I believe the same is true for non-physical currencies.

So what currency does our hypothetical lunar coder—let’s call her Ada—charge her customers in? She weighs up the benefits each currency provides in terms of spending power (money is useless to her unless she can spend it), stability (so that a week’s work isn’t worth only a loaf of bread by Saturday), and the features that currency offers (like fungibility, being [backed by gold](https://web.archive.org/web/20171123012558/http://www.bankofengland.co.uk:80/banknotes/Pages/about/faqs.aspx) or being [accepted in more countries](https://en.wikipedia.org/wiki/International_use_of_the_U.S._dollar) than any other).

So, if we assume cryptocurrencies have enough features that they gain popularity, and thus the _potential_ for spending power and stability, what features do they have that would make them a good choice for pure-digital traders?

### Cost

Though it’s not always visible to consumers, all financial transactions cost money to process. To pay for your beer with a bank note, there must be the infrastructure that prints those notes, that ensures forgeries are difficult and quickly identified and there must be staff that operate that infrastructure and _physically move the notes around_. If you think your bank is pushing [contactless payment](https://en.wikipedia.org/wiki/Contactless_payment) upon you, you’re probably right; removing the cost of moving and verifying bits of paper has been a huge boon to banking.

Compare that to the infrastructural overhead of a cryptocurrency. As described in the background sections above the infrastructural cost of these new currencies is borne by machines; the cost of the electricity to run them paid for with controlled, predefined inflation of the currency.

Today the cost of a single low-volume digital transaction in GBP might be [20p + 1.4%](https://stripe.com/gb/pricing) but with Ethereum, being decentralised, you _specify_ how much you’re willing to pay for others to verify your transaction. This makes the amount so small it’s measured in _[Wei](https://ethdocs.org/en/latest/ether.html#denominations)_ (the micro-micro-pence to Ether’s pound) and, as I write, the [average amount spent per transaction](https://etherscan.io/chart/gasprice) (the _average GasPrice_) is 22 billion Wei which is easiest to conceptualise as ~4,300 verifications costing just a penny.

### Speed

Transfers within the same traditional bank are fast, seconds usually, because they operate entirely within the same ledger. Between banks in one country a couple of hours is normal and between countries with good foreign relations a couple of days. Trying to get money out of Egypt? Or China? Or large amounts of money between any two countries? Weeks of delay are often advised.

By contrast the speed of a transaction in Ethereum is blazingly fast and predictable. While it is proportional to your offered GasPrice or fee, on average a receiving entity can expect confidence in a transfer [with the next block](https://etherscan.io/chart/blocktime), or after about 15 seconds.

### Proof of ownership

In 1997 [Nick Szabo wrote a paper](https://journals.uic.edu/ojs/index.php/fm/article/view/548/469) detailing a framework for “Smart Contracts”, which I’ll cover below, but it also included a description of the critically useful concept of digital _Bearer Certificates._

Just like a banknote is a certificate of your ownership of a sum of money, and a deed that of a plot of land, a single Ether coin is the certificate that you bear, or own, the value for some amount of work performed for the benefit of the Ethereum network (the coin was ‘mined’ and passed on to you for goods or services). Now the currency’s nomenclature becomes clearer: _Ethereum_ is a platform for bearer certificates; the _Ether_ currency is the core certificate of that platform (and the one in which the machines that maintain the infrastructure of that network are paid) but it is only one of an unlimited number of certificates which can be created and used by anyone.

If our lunar coder, Ada, has faith in Ethereum she can set up a type of bearer certificate that represents a license to use the software she builds. She can encode the ownership of one of these certificates into the Ethereum blockchain and thus ownership can be mechanically verified by anyone.

What’s more, she can have her software _check the blockchain_ when it starts up and ensure that the current user owns the private key that’s paired with the owner of the licence—i.e. that the user has paid their dues. Ethereum provides a platform for the trade of more than money.

### Arbitration

So far the features described can broadly be provided by all cryptocurrencies, but Ethereum (amongst some less notable others) has an ace up its sleeve: Smart Contracts. These are, in essence, a way of creating conditions on transfers.

At this point Ada can set up a digital shop on a server on the Moon which accepts payments in Ether, waits for verification, then issues a new licence bearer certificate for that customer, then encodes the ownership of that certificate back to the blockchain. However this process could be quite slow as the data goes back and forward between the Earth and the Moon, a customer might get bored and not complete their purchase in the 2.6 seconds for each back and forth stage of that process. This is where Smart Contracts step in.

Smart Contracts can be thought of as a program in their own right. Ada can put one together that says: “Contract, issue a ‘licence’ bearer certificate to anyone who deposits 0.5 ETH or more in this account.” When this contract is written to the blockchain it now exists on every machine that interacts with Ethereum, it is executed by the machines that are working to verify transactions (to hark back to my analogies above, it is executed by every messenger who picks up a payment to your account) and, even if the internet link to the Moon goes down, Ada will find 0.5 ETH more in her account for each licence that gets created when the connection returns and she updates her copy of the blockchain.

### Autonomous business

From these building blocks great things can be made. A collection of Smart Contracts operating together is called a _Decentralised Autonomous Organisation_—a DAO—and our talented protagonist, being a savvy entrepreneur, can decide to make great use of them.

She decides that she’s very happy for people to pass on their licences to friends, but she feels like she’s due a small portion of the proceeds if her software is sold on. Because all ownership transfers on the Ethereum platform are public (they’re in the blockchain) she can codify that a second-hand licence is only valid if 5% of the exchange fee is delivered to her; she can even write a Smart Contract to handle this accounting for her, so she can focus on fixing the bugs with her program and let her autonomous organisation handle the money, even when she’s offline.

### Law and Tax

Where this automation becomes _really_ interesting is when it takes the place of complex manual processes.

> ’Tis impossible to be sure of anything but Death and Taxes

An idiom which surely predates even [Christopher Bullock](https://en.wikipedia.org/wiki/Death_and_taxes_%28idiom%29), but one that will remain true for a long time yet. Ethereum, or a framework like it, _will be_ the platform upon which community finance is built in the not-too-distant future.

Imagine Ada’s world, writing software on the Moon. Living there has some physical infrastructure costs that need to be paid somehow. The lunar government decides that it will extract 20% of its citizens’ income for repairs and useful things like oxygen and food. Being a forward-thinking celestial nation they offer small companies financial assistance and support writing Smart Contracts, on the condition that subscribed companies update their DAOs so that as they pay dividends, 20% goes to Lunar Governance, and 80% to the board’s wallets.

It turns out that the support that Lunar Governance offers for digital enterprises is so excellent, that people _who don’t live on the Moon_ decide that they’re going to register their DAOs there and benefit from the nation’s digital infrastructure for the low price of 20% in tax. With the complexities of extracting tax outsourced to machines, the Moon’s thriving corporate tourism fuels its economy and we slowly realise that [Matthew 5:5](https://en.wikipedia.org/wiki/Matthew_5:5) had a typo, and it is indeed the geeks that inherit the Earth.

Writing software on the Moon is something of a ridiculous example but there already exist a [host of micronations](https://en.wikipedia.org/wiki/List_of_micronations) which could, and in my opinion _will_, provide the physical location for digital-only communities, for a price, while our law-making struggles to catch up with our technological progress. Ethereum is the prototype, if not the platform, upon which this will occur.

---

The features described above make Ethereum an incredibly powerful platform, and I’m positing that its success to date is based in the potential that its backers see in the advantages this technology provides, as I’ve tried to describe above. Many of these backers are not curious individuals; the _[Blockchain Alliance](https://www.blockchainalliance.org/)_ is a group, made from some of the [largest names in finance](https://uk.reuters.com/article/us-ethereum-enterprises-consortium-idUKKBN1662K7), with the aim of harnessing the potential of these technologies. These firms invest time and money in many things, but there’s a reason [FinTech](https://en.wikipedia.org/wiki/Financial_technology) is seeing more startups than any other area at the moment.

## Stability and spending power

Despite their successes, cryptocurrencies are highly volatile and not without the potential for serious problems.

{{< figure src="ethereum-candles.png" caption="The [Kraken ETH/USD](https://cryptowat.ch/kraken/ethusd) exchange demonstrates the volatility in cryptocurrency value today." alt="A financial candle chart of Ethereum's price from May 17th to 21st, 2017" >}}

This chart displays the dollar value of a single Ether coin in the past 5 days—from $90 to $140, with fluctuations of 2% in a day being commonplace. The [growth of the currency](https://cryptowat.ch/kraken/ethusd/3d), from around $5 in February this year, through to $140 only 4 months later, displays huge potential and international investment in these technologies fuels that growth, but huge plummets in value are commonplace as issues are discovered then recovered from.

The problems that are encountered are novel too, often complex enough that it can be hard to wrap heads around.

China operates a scarily large proportion of the world’s [Bitcoin mining operations](s//www.bbc.com/future/story/20160504-we-looked-inside-a-secret-chinese-bitcoin-mine) which is risky as a single entity with a [51% share of mining power](https://www.coindesk.com/51-attacks-real-threat-bitcoin/) can, in theory, selectively allow double spending. Ethereum took a massive hit last year when “The DAO”—the then-preeminent Ethereum based decentralised autonomous organisation which facilitated startup investment to the tune of $200m—was “hacked” or, to phrase [more correctly](https://www.cryptocompare.com/coins/guides/the-dao-the-hack-the-soft-fork-and-the-hard-fork/), someone exploited a bug in the Smart Contracts it was built upon to extract $70m in an unintended manner.

Ethereum recovered from that scandal (some say it [demonstrated its maturity](https://www.cryptocoinsnews.com/dao-best-thing-happen-ethereum-can-back/) as it did) and vigilance and self-interest seems to be keep the [51% problem Bitcoin faces](https://www.coindesk.com/ahead-bitcoin-halving-51-attack-risks-reappear/) in check; but new technology always brings new risk.

These are risks I’m happy taking, as I think the potential is far, far greater; I’m certainly polishing my Smart Contract writing skills.
