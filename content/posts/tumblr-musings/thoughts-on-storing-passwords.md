---
title: Thoughts on storing passwords
date: 2013-03-20T14:53:00+00:00
draft: false
emoji: 🔐
tags:
  - passwords
  - code
  - security
  - database
  - from-tumblr
atUri: "at://did:plc:ephkzpinhaqcabtkugtbzrwu/site.standard.document/3mdryqsrb7d2a"
---
Yesterday at work I was discussing the design of password databases with my boss. I don't claim to have any deep knowledge in the area, but we came out with the following structure which seems to have good inherent security as well as the possibility to improve that as technology improves.

## Password

**Store:** Password

Without going over password databases 101, you obviously should never store just a plaintext against an email address/username, as in the eventuality where your database is compromised — let's say [Doctor Claw](https://en.wikipedia.org/wiki/Doctor_Claw) has stolen your database — anyone with that data can access anyone else's accounts. Bad times.

## Hashed password

The simple solution is hashing the passwords using a [one-way hash algorithm](https://en.wikipedia.org/wiki/Cryptographic_hash_function). You store `HASH(password)`, and when the user wants to authenticate, you compare `HASH(what_they_sent)` against the stored `HASH(password)` and if they match then the passwords are the same[^1].

**Store:** Password hash

The benefit here is that Doctor Claw now can't use the data he has to log in to anyone else's account. If he tries sending a hash from the stolen database to the site, the calculation performed will be `HASH(the_hash_doctor_claw_sent)` which will not match `HASH(the_real_user_password)`.

The hash algorithm used can be one of many, [MD5](https://en.wikipedia.org/wiki/MD5) is an old favourite, but it's considered [insecure these days](https://en.wikipedia.org/wiki/MD5#Security), and the same goes for [SHA1](https://en.wikipedia.org/wiki/SHA-1). The current advice is to use [SHA256](https://en.wikipedia.org/wiki/SHA-256), but this will undoubtedly change with time.

## Salted hashed password

There's a cloud in our perfect passworded sky; storage is cheap and people in the real world often don't use complex passwords. Doctor Claw has done some pre-work and he's set a home computer off calculating the hashes of all the most popular passwords — he'd put this in a lookup specialised database called a [rainbow table](https://en.wikipedia.org/wiki/Rainbow_table). He can now do a rapid lookup to see if any of the hashes in his stolen database match any in his rainbow table and will immediately know the original password for that user account.

**Store:** Salted password hash, Salt

The solution is simple, _[salt](https://en.wikipedia.org/wiki/Salt_(cryptography)) the password_. To do this, when you're writing a password to your database you follow these steps:

1. Generate a salt by creating a random string of characters, like `gZ7s(28d}"+`
2. Concatenate the salt and the password before you hash them (you could append the salt, prepend it, or anything else repeatable).
3. Write the result of `HASH(salt + password)` to your database along with the salt you used.

Now when a legitimate user sends you the right password, you simply add the salt to it in the same way, hash it, and compare against your database. Doctor Claw on the other hand has to generate a rainbow table for _each_ entry in the stolen database, and he doesn't necessarily know that the salt is prepended, appended or sprinkled around within the password. Doctor Claw's life grows harder…

## Repeated hashing & derrived keys

Any technologically savvy person can see that available computing power is by no means static and Doctor Claw could potentially have created or bought himself use of a [botnet](https://en.wikipedia.org/wiki/Botnet) to generate his rainbow tables, allowing him relatively speedy access to getting at your customer data yet again. So how can we tackle this one?

**Store:** Derrived key, salt, iterations

The one thing which is, within the forseeable technological future, always going to be true is that generating a hash takes _time_. Not necessarily much time, but it does (currently around the hundreds of microsecondsmark). The one thing which Doctor Claw here has always been interested in has been accessing accounts fairly quickly (on the order of days, not years or decades), so this is something we can use to our advantage.

Rather than storing just the hash of a salted password (1 time), we can store the hash of the salted hash of a salted password (2 times), or the hash of a salted hash of a salted hash of a salted password (3 times). With each additional salting and hashing we're increasing the amount of time it takes _us_ to check a legitimate user's credentials, but also increasing the time it takes Doctor Claw to either [brute force](https://en.wikipedia.org/wiki/Brute-force_attack) all possible passwords to their hashes, or generate a rainbow table.

If you salt and hash in the order of thousands of times for each password, it takes you around one second to authenticate that the password your legitimate user has sent you is correct or incorrect, but for Doctor Claw to reverse engineer or guess any user's password even from a complete password database will take in the order of millennia. Doctor Claw has been foiled again.

Your design could specify "we repeat hashing 1000 times", but this is a bit restrictive, as within 2 years computers may be quick enough that 1000 iterations may only take milliseconds rather than seconds, so it's worth considering storing the salt, the derrived key _and_ the number of repetitions of hashing that has taken place to create that hash. This means that if you want to upgrade your database to take account of that new quantum processor that's just been released, you only need to set a program iterating through your database, reading the info, hashing another N times, and replacing the data with the new derrived key and the incremented value for 'iterations'.

It's important to note that if you were to always put the salt in the same position next to the hash before performing the next hash, it would be simple(ish) for Doctor Claw to create a rainbow table of all possible salt + hash combinations and skip all the hard work you've done iterating a large number of times, so it's a good idea to have your hashing being a function of the previous hash, the salt _and_ the iteration number. A entirely unoptimised example of this would be:

```javascript
function hashIt(string,salt,iterations) {
  if (iterations <= 1) {
      var hashThis = string;
  } else {
      var hashThis = hashIt(string,salt,iterations - 1);
  }
  // Split the salt at a predictable, non regular point
  var saltParts = salt.splitAtPosition(iterations % salt.length);
  // Do the hash
  return SHA256(saltParts[0] + hashThis + saltParts[1]);
} 
```

There are formalised methods for achieving this, like [PBKDF2](https://en.wikipedia.org/wiki/PBKDF2), to prevent the need to worry about small but important details like this.[^2]

## Conclusion

Password databases require some serious thought (there are very likely to be angles of attack I haven't considered above too) but if you've avoided the traps I mentioned above then you're in reasonably good stead!

[^1]: There is a possibility that there would be a 'hash collision' - because the hash is not a 1:1 mapping, it is a certainty that there are values for `A` and `B` such that `HASH(A) == HASH(B)`, but by choosing the right hash algorithm this can be reduced to a statistical improbability.
[^2]: Import note: More than that, there are lots of subtle ways bugs can creep into complex key derrivation functions like this — use one that's been security audited!
