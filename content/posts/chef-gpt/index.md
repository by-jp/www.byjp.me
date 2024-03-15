---
title: "Chef GPT"
emoji: 🤖
date: 2024-03-15T19:13:34.952Z
summary: I asked ChatGPT to give me a recipe for ingredients I had, cooked it and was unsurprised.
tags:
- food
- recipes
- mela
- LLMs
- "AI"
---

A few weeks ago I started my year with [Zoe](https://zoe.com/). I've loved seeing my glucose response to the foods I eat, and the blood, fat, and gut microbiome reports that arrived a few days ago have been _fascinating_. I'm trying to shift my diet to meet what I've learned there (more on this below), but this evening I found myself with lots of fresh ingredients not long for this world, and no immediate ideas on how to combine them.

## Zoe

A quick aside to outline the changes I'm making to my diet because of what I've learned from Zoe. None of these will be news to those of you who've spent time reading about the dietary needs of humans this decade (especially if you've read anything by [Tim Spector](https://en.wikipedia.org/wiki/Tim_Spector), or have a copy of [The Glucose Goddess Method](https://openlibrary.org/works/OL34953020W/Glucose_Goddess_Method), like {{< friend yvette >}} does) but, at least, these were newly concrete for me:

- My blood sugar control is "poor". To balance this I'm trying to wildly reduce refined sugars in my diet, reduce carbohydrates in general, and try to sequence them after fats and proteins to reduce how much they make my my blood-glucose levels swing.
- My blood fat control is surprisingly good. This means I can lean into gaining energy from fats if the reduction in carbohydrates leaves me with a deficit. I need to remain careful though, it still takes ~6 hours for fat to clear my bloodstream, so one of lunch or dinner should be lower in fat.
- My microbiome score is middling to poor. This is unsurprising (a family trait), but this means I'm trying to lean further into high-fibre greens (like broccoli, spinach, and kale), tofu, mushrooms, and [nuts](https://www.reddit.com/user/pfobwpfo/comments/18ohqi2/nuts/).

Thus the scene is set for my ChefGPT challenge…

## `sudo make me a meal`

I'm _extremely_ skeptical about Large Language Models (LLMs, or what the world is calling "AI" tis decade). They're eerily lifelike at responding to human questions, and can get in the right ballpark for answers most of the time, but their energy consumption (both producing and using them), their ethics (selective application of copyright), and their "intelligence" are all _highly_ questionable. But leaning about how they work, their limtations, and ow to recognise what they produce are all things I think will be very useful in the months and years ahead.

So I decided to ask [ChatGPT](https://chatgpt.openai.com) to give me a recipe to cook. I'm not the first and I won't be the last, but I want to document the process, the results, and my tasting notes so that I can look back and laugh in 25 years.

I wrote a prompt (and re-wrote it ~10 times, fiddling with the output) so that I'd end up with a recipe I could use in my favourite recipe app, [Mela](https://mela.recipes). In my prompt I included my general request (including my Zoe-taught guidelines), the ingredients I had to hand, and some details of [Mela's (awesome) file format](https://mela.recipes/fileformat/index.html), so it'd be easy to import

> Give me a tasty recipe, that keeps glucose levels low, which I can make with any of the following ingredients I have:
>
> Smoked salmon
> Gnocchi
> Broccoli
> Brussel Sprouts
> Onions
> Sweet Potato
> Mushrooms
> Butter beans
> Almonds
> Sunflower seeds
> Hazelnuts
> Sweetcorn
> Chickpeas
> Tuna
> Tomatoes
>
> Give your answer in JSON format with the keys: `id` (a UUID), `title`, `text` (the description of the recipe), `images` (an empty array), `categories` (an empty array), `yield` (the yield as a string, eg. "4 people"), `prepTime` (the prep time, in the format 3h 30m), `cookTime` (the time to cook, in the format 3h 30m), `totalTime` (the total time to cook, in the format 3h 30m), `ingredients` (a list of ingredients, in metric measurements, separated with `\n`), `instructions`(one paragraph per step, separated by `\n\n`. It's very important you don't use numbers), `notes` (written in markdown, containing any other notes if any are needed), and `nutrition` (an empty string).

## Importing into Mela

