---
title: "10 Days to Make, 30 Years to Remake"
date: 2025-02-28
lastmod: 2025-02-29
description: "This has to be shown below the title"
category: "JS"
---

> Roses are red, Violets are blue, How React's working, you ('ll probably) have no clue
> - Me

## A sample title

To my favorite screen readers, Prime and Theo <3 (sorted alphabetically)

**Programmer Check:** Please verify that you are a legit programmer by passing my self-invented, 100% guaranteed. Change your browser to **Microsoft Edge Windows** without buying a new PC. Why Edge on Windows, you ask? Because I don't have the mental capacity to manually change my UA each time I wanna open my own article.

**Warning:** This article is profanity-optional. You can opt-in to profanity if you'd like to, and the current version is considered SFW.

JS Programmers (and generally all programmers)... If they're doing things right, then you can't argue that the JS Ecosystem is not broken. If they're not doing things right, who is supposed to do so. Go kids? Fortran Grandpas?

It just doesn't make sense that millions of people gather around a fundamentally-flawed language, and do jaw-dropping shit to make it, well, actually work. Just think about it. A build system... For a language (theo named it, starting with C, now (langname)x) that transpiles to another language (lang name) that transpiles to another language (ts) transpiles to an interpreted language...



[[Attachments/10 days to make, 30 years to remake/e2dc9bd95e7c59387116282c2479af75_MD5.png|Open: Pasted image 20250224003231.png]]
![[Attachments/10 days to make, 30 years to remake/e2dc9bd95e7c59387116282c2479af75_MD5.png]]

https://www.wired.com/story/inside-the-cult-of-the-haskell-programmer/#:~:text=Great%20programming%20languages%20aren%E2%80%99t%20always%20great%20for%20programming.

If you think ESLint is so smart, think again. Grammarly is the world's most complex linter, period.

You know what's funny? People argue about JS, and write their blog posts in English. I'm wondering what was the Esperanto of programming languages?

Sometimes I wonder if JS devs even "program" at all. They use the Next.js framework to handle their frameworks, TS to prevent skill issues, Zod to prevent their skill issues' skill issues, Tailwind for styles, no sorry, they use [insert ui lib name] so they don't need Tailwind... Clerk for Auth, Turso or PlanetScale for DB, Prisma to fuck their DBs, Leftpad to get fucked by NPM (indirectly of course), PostHog to spy on their users, Then they even use Cursor to write that 50 lines of code they need to write to glue *the shits™️* together. Finally, host it on a Scam-as-a-Service, aka Vercel. "They're just here for the vibes, man [link to Alberta's JS is Vibes Short]" (there we go YouTuber pun™️!)

Convincing a JS dev to move to another language is arguably easier than trying to make them understand what does "Right Level Of Abstraction" mean. 

If you think that AI is going to take your job in the next 5 years and you've ever worked with Cursor, you're right.

I've watched and read many "JS dev tries X" contents, and they all have something in common: A pool (or lake, to be more accurate) of Skill Issues. 
- JS dev tries Go, and says: "It doesn't make me feel smart". Then what does make you feel smart? Making a S3 wrapper using a bunch of more wrappers?
- JS dev tries Rust: Rust sucks. Who needs performance when we have Bun?
- JS dev tries Rust: Googles "How to rewrite Remix App in Leptos?"
- JS dev tries C#: Okay this was too surreal.
Let's get a bit harsh: YOU ETHICALLY DON'T HAVE THE RIGHT TO EXPRESS YOUR OPINIONS ABOUT ANOTER LANGUAGE BEFORE YOU DROP YOUR JS BRAIN AND START WORKING WITH THE FUCKDAMN THING FOR AT LEAST A MONTH, THE WAY IT IS INTENDED TO BE USED. You are the one who's fucking carriers of numerous junior devs by teaching them plumbing and calling it "programming", not that Python nerd that doesn't teach R\*\*ct to their students. IMO "Juniors can't code anymore" isn't just because of AI. Maybe, just maybe we're not teaching them how to code.

> Cursor won't replace JetBrains. It will replace JSBrains.

One of the strongest arguments of JS programmers is that "JS is good because you can make things super fast with it, so you'll realize whether your idea is good enough or not without spending so much time and money making it.". Sounds fair, right? If you can get a working prototype ASAP and get feedback on it, why spend so much time actually _programming_ it? ATEOTDYCRIIRWYIS (at the end of the day, you can rewrite it in Rust whenever your idea succeeds (this sentence is so common among JS devs that has been abbreviated (This sections is the Lisp dev's dream))), right? Well, it turns out that this doesn't happen. Take Twitch, Netflix, Google, Microsoft, Me\*a, whoever talks about their codebases, doesn't say good things about it. Because as you succeed and grow, your codebase gets larger. The larger your codebase is, the harder it is to rewrite. That's why often when companies get big, instead or rewriting their whole codebase, they start doing the shitty band-air strategy, and replace only the parts that have most overhead with some external Rust thing. Let me be honest, it becomes dogshit after a few years, so hardly that a codebase that could be easily maintained by 3 devs, now has 20 people dedicated to it.

Starting with JS with this mindset isn't a smart move. It's deprovisioning. It's not "ship fast", it's "fail fast". It's like attaching jets to a bike. And if you ask me, I'll say one of the most significant reasons that 90% of startups fail, is this mindset: Planning for failure, instead of success. They choose to ship shit so if they fail, they've wasted less time and money; and that's the main reason tech startups fail. This idea is primarily backed by not "risking to take risks". Seems logical, but logic is a piece of assshit (don't quote me on this).

Go is simple, Python is easy, Java is Oracle C# (checkmate Linus), and JS is `undefined`. The language you use, with abstractions you use, is so simple, and has so many source code available that A. Anyone can quickly build things with it (I intentionally didn't use "learn") B. AI can also be one of those. That's why Bolt.new, v0 (V like Evil, like V\*\*\*\*l), etc. are focused on JS. Seems that JS's biggest strengths are now it's biggest weakness, and that's why Cursor sucks at building an ARM64 emulator with Assembly.

I don't mean spend two years just finding the best programming language for your niche (DreamBerd), but please don't go a path that doesn't work for 90% of people.

> In Utopia, JS devs use NPM (aka No Package Manager). But wait, it's not true. There's no JS in Utopia.