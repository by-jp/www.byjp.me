---
title: "ChefGPT"
emoji: 🤖
date: 2024-03-15T19:13:34.952Z
summary: I asked ChatGPT to give me a recipe for the ingredients I had at home, cooked its suggestion blind and… was unsurprised. I still got a lot from it though!
tags:
- food
- recipes
- mela
- LLMs
- "AI"
topics:
- Creativity
- Cooking
syndications:
- https://hachyderm.io/@byjp/112102305952063134
- https://bsky.app/profile/byjp.me/post/3knrhqsj6ac23
---
A few weeks ago I started my year with [Zoe](https://zoe.com/), the science-backed diet learning programme. I've loved seeing my glucose levels respond to the foods I eat in realtime, and the blood, fat, and gut microbiome reports that arrived a few days ago have been _fascinating_. I'm trying to shift my diet to meet what I've learned there (more on this below), but this evening I found myself with lots of fresh ingredients not long for this world, and no immediate ideas on how to combine them.

## Zoe

A quick aside to outline the changes I'm making to my diet because of what I've learned from Zoe. None of these will be news to those of you who've spent time learning about the dietary needs of humans this decade (especially if you've read anything by [Tim Spector](https://en.wikipedia.org/wiki/Tim_Spector), or have a copy of [The Glucose Goddess Method](https://openlibrary.org/works/OL34953020W/Glucose_Goddess_Method), like {{< friend yvette >}} does) but, at least, these ideas have been made newly concrete for me:

- My blood sugar control is "poor". To balance this I'm trying to wildly reduce refined sugars in my diet, reduce carbohydrates in general, and try to sequence them after fats and proteins to reduce how much they make my my blood-glucose levels swing.
- My blood fat control is surprisingly good. This means I can lean into gaining energy from fats if the reduction in carbohydrates leaves me with a deficit. I need to remain careful though, it still takes ~6 hours for fat to clear my bloodstream, so one of lunch or dinner should be lower in fat.
- My microbiome score is middling to poor. This is unsurprising (a family trait), but this means I'm trying to lean further into high-fibre greens (like broccoli, spinach, and kale), tofu, mushrooms, and [nuts](https://www.reddit.com/user/pfobwpfo/comments/18ohqi2/nuts/).

{{< figure src="gut-boosters.webp" title="I now have a large selection of Zoe's recommendated foods at home. But do they include nuts, or just legumes? I'd best call David Mitchell." alt="A list of Zoe's \"updated gut booster\" foods, for me. Inlcuding almonds, broccoli, spring greens, hazelnuts, and more." >}}

And so the scene is set for my ChefGPT challenge…

## sudo make me a ~~sandwich~~ meal

I'm _extremely_ skeptical about Large Language Models (or LLMs)—things like [ChatGPT](https://chatgpt.openai.com), or what the world is calling "AI" this decade. They're eerily lifelike at responding to human questions, and can give answers that are in the right ballpark most of the time, but their energy consumption (both when being "trained" and used), their ethics (selective application of copyright), and their "intelligence" are all _highly_ questionable. When I'm skeptical about something, I like to learn about how it works, its limtations, and how to recognise what it produces; these are all things I think will be particularly useful for LLMs in the months and years ahead.

So, with that in mind and my rapidly wilting vegetables in hand, I decided to ask ChatGPT to write a recipe for me to cook. I'm not the first to try this, and I won't be the last to write about it, but I want to document tis little adventure, the results and my tasting notes, so I can look back and laugh in 25 years' time.

I wrote a prompt (and re-wrote it a few times, fiddling with the output) so that I'd end up with a recipe I could use in my favourite recipe app, [Mela](https://mela.recipes). In my prompt I included my general request (including my Zoe-taught guidelines), the full list of ingredients I had to hand, and some details of [Mela's (awesome) file format](https://mela.recipes/fileformat/index.html), so it'd be easy to import and use.

{{% details "If you like reading prompts meant for LLMs you can peek in here." %}}
> Give me a tasty recipe, that keeps glucose levels low, which I can make with any of the following ingredients I have available:
>
> Smoked salmon\
> Gnocchi\
> Broccoli\
> Brussel Sprouts\
> Onions\
> Sweet Potato\
> Mushrooms\
> Butter beans\
> Almonds\
> Sunflower seeds\
> Hazelnuts\
> Sweetcorn\
> Chickpeas\
> Tuna\
> Tomatoes
>
> Give your answer in JSON format with the keys: `id` (a UUID), `title`, `text` (the description of the recipe), `images` (an empty array), `categories` (an empty array), `yield` (the yield as a string, eg. "4 people"), `prepTime` (the prep time, in the format 3h 30m), `cookTime` (the time to cook, in the format 3h 30m), `totalTime` (the total time to cook, in the format 3h 30m), `ingredients` (a list of ingredients, in metric measurements, separated with `\n`), `instructions`(one paragraph per step, separated by `\n\n`. It's very important you don't use numbers), `notes` (written in markdown, containing any other notes if any are needed), and `nutrition` (an empty string).
{{%/ details %}}

## Importing into Mela

ChatGPT _is_ impressive, in that it did produce a recipe for "Smoked Salmon and Vegetable Gnocchi Stir-Fry". It's definitely a low-glucose recipe, and in exactly the right structure to import into Mela. I copied it's response, put it in a text editor, saved it with a `.melarecipe` extension and openned it in Mela.

{{< figure src="recipe.webp" title="Mela's gorgeous user interface, inviting me to start a-cookin'. If you'd like to try the same recipe (🫣) then you can use the mela recipe data below, or [grab the PDF](recipe.pdf)." alt="A simple and elegant screenshot of the Mela recipe app showing the recipe generated by ChatGPT. You can read the recipe in the PDF linked to in this image's caption.">}}

{{% details "If you're the kind of person who likes reading JSON recipe files (hi friend!), then get a load of this!" %}}

```json
{
  "id": "bf30de65-8497-4d8a-93d1-fdcb9fc7b8ad",
  "title": "Smoked Salmon and Vegetable Gnocchi Stir-Fry",
  "text": "This recipe offers a delicious stir-fry featuring smoked salmon, gnocchi, and an array of low-glycemic vegetables. The combination of flavors and textures makes it a satisfying and balanced meal that won't spike glucose levels.",
  "images": [],
  "categories": [],
  "yield": "4 servings",
  "prepTime": "15 minutes",
  "cookTime": "20 minutes",
  "totalTime": "35 minutes",
  "ingredients": "200g smoked salmon\n400g gnocchi\n1 cup broccoli florets\n1 cup brussel sprouts, halved\n1 onion, sliced\n1 medium sweet potato, cubed\n200g mushrooms, sliced\n1/4 cup almonds\n2 tbsp sunflower seeds\n2 tbsp hazelnuts\n1/2 cup sweetcorn\n1/2 cup chickpeas\n150g butter beans\n1 can tuna, drained\n2 tomatoes, diced\n2 tbsp olive oil\nSalt and pepper to taste",
  "instructions": "Begin by boiling gnocchi in a large pot of salted water until they float, then drain and set aside. In a large skillet, heat olive oil over medium heat. Add sliced onions and cook until translucent. Next, add sweet potatoes, broccoli, and brussel sprouts, cooking until slightly tender. Stir in mushrooms, sweetcorn, chickpeas, and butter beans. Once the vegetables are cooked to your liking, add the cooked gnocchi and smoked salmon, stirring gently to combine. Toss in diced tomatoes and tuna, then sprinkle with almonds, sunflower seeds, and hazelnuts. Season with salt and pepper, and cook for an additional 2-3 minutes until everything is heated through. Serve hot and enjoy!",
  "notes": "* Feel free to adjust the vegetables based on personal preference or availability.\n* This recipe is versatile - you can add additional herbs or spices for extra flavor.\n* Ensure to check the label when purchasing smoked salmon to ensure it doesn't contain added sugars or sauces.",
  "nutrition": ""
}
```

{{% /details %}}

So, the fact that it looks so good is _definitely_ Mela's fault (did I mention how much I like that app?), but nothing here looked like it'd kill me so I dug out my large skillet and got to work.

## The proof is in the eating

As I followed the recipe (I made sure not to read ahead too closely, to keep the sense of suspense alive) I began to realise ~~ChefGPT's first mistake~~ why a statistical word-picker won't be making your most loved recipes any time soon. It turns out my prompt "with any of the following ingredients I have available" ended up being interpreted more like "using exactly all of these ingredients".

So I got a _Smoked salmon, Gnocchi, Broccoli, Brussel Sprout, Onion, Sweet Potato, Mushroom, and Butter bean_ stir fry with _Almonds, Sunflower seeds, Hazelnuts, Sweetcorn, Chickpeas, Tuna, and Tomatoes_ on top. 🤷

I kept going. For science.

I appreciated that the instructions were trivial to follow, and there was room for my expression as a sous-chef too. I decided to add a little water in the middle of step 3 to help the diced sweet potato soften up, and I ended up taking the advice in the notes and dropping the chickpeas and sweetcorn, to keep the carbohydrate count lower. All the while my apprehension was growing.

Finally, just over 35 minutes later (good timing guesses!), I plated up and got out my notepad.

{{< figure src="plating.webp" title="" alt="A photo of a serving of ChefGPT's work, smartly plated up in a pasta bowl, knife and fork either side." >}}

It… wasn't bad? I mean, there aren't many ways you can screw up a stir fry, so a safe choice there for sure. It was definitely edible, and even enjoyable perhaps?… to a point.

Where it started to fall down is the exact same place I see LLMs fall down when I've used them as a "co-author" of software. A statistical word-picker isn't _thinking_, so has no sense of composition or balance.

This was very clear as I smelled my creation — _wow_, definitely tuna there. It wasn't gross by any means, but there's no part of that tuna which was matched or blended into other aromas to make something interesting and new. Even the "pepper to taste" didn't keep that initial dominating smell at bay.

As I took my first mouthful I noticed that the butter beans, slightly undercooked, combined with the various nuts made for a very aggressive texture mismatch; soft, silty beans with hard, crunchy almonds and hazelnuts was definitely a novel experience.

But, after these few eyebrow-raising moments, I realised I could stop cringing, it was actually pretty tasty; the flavours were all good, even if they weren't exactly jamming together.

{{< figure src="joey-trifle-good.webp" alt="An animated clip from the TV show Friends. Joey is describing the layers of an uncoventionally constructed trifle as he eats it: \"Custard, good. Jam, good. Meat… gooooood!\"" >}}

Despite the lingering dominance of the tuna, the total absence of the smoked salmon, and those textural WTF moments, I really enjoyed this adventure, and might even get ChatGPT to offer ideas in the future—in emergencies, and with some heavy editing.

This recipe did give me some interesting ideas around adding hazelnuts into other dishes; the mellow flavour that came from them spending a little time steaming next to the veg was really enjoyable. I'd probably toast and crush them though, to minimize that jarring texture clash.

It's not a recipe that'll be staying in my library, but it did become a member of the [clean plate club](https://www.youtube.com/playlist?list=PLdAiyt6EtZP4wkMoQIHp2h00Ta9PyBtdi), and when I asked Zoe to check it out, it gave a commendable score of 65/100 for matching my dietary needs. Zoe helpfully proclaims this as being squarely in the "Enjoy regularly" category. Thanks, but no thanks.
