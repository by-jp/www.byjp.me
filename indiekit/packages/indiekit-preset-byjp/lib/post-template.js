import camelcaseKeys from "camelcase-keys";
import TOML from "@iarna/toml";
import YAML from "yaml";

/**
 * Get content
 * @access private
 * @param {object} properties - JF2 properties
 * @returns {string} Content
 */
const getContent = (properties) => {
  if (properties.content) {
    const content =
      properties.content.text || properties.content.html || properties.content;
    return `\n${content}\n`;
  } else {
    return "";
  }
};

/**
 * Get front matter
 * @access private
 * @param {object} properties - JF2 properties
 * @param {string} frontMatterFormat - Front matter format
 * @returns {string} Front matter in chosen format
 */
const getFrontMatter = (properties, frontMatterFormat) => {
  let delimiters;
  let frontMatter;

  /**
   * Go templates don’t accept hyphens in property names
   * and Hugo camelCases its predefined front matter keys
   * @see {@link https://gohugo.io/content-management/front-matter/}
   */
  properties = camelcaseKeys(properties, { deep: true });

  /**
   * Replace Microformat properties with Hugo equivalents
   * @see {@link https://gohugo.io/content-management/front-matter/}
   */
  properties = {
    date: properties.published,
    publishDate: properties.published,
    ...(properties.postStatus === "draft" && { draft: true }),
    ...(properties.updated && { lastmod: properties.updated }),
    ...(properties.deleted && { expiryDate: properties.deleted }),
    ...(properties.name && { title: properties.name }),
    ...(properties.photo && {
      media: properties.photo.map(({url, alt}) => ({
        url: `/${url}`,
        alt,
      })),
     }),
    ...properties,
  };

  delete properties.content; // Shown below front matter
  delete properties.deleted; // Use `expiryDate`
  delete properties.name; // Use `title`
  delete properties.postStatus; // Use `draft`
  delete properties.published; // Use `date`
  delete properties.slug; // File path dictates slug
  delete properties.type; // Not required
  delete properties.updated; // Use `lastmod`
  delete properties.url; // Not required
  delete properties.photo; // Was moved to media

  switch (frontMatterFormat) {
    case "json": {
      delimiters = ["", "\n"];
      frontMatter = JSON.stringify(properties, undefined, 2);
      break;
    }

    case "toml": {
      delimiters = ["+++\n", "+++\n"];
      frontMatter = TOML.stringify(properties);
      break;
    }

    default: {
      delimiters = ["---\n", "---\n"];
      frontMatter = YAML.stringify(properties, { lineWidth: 0 });
      break;
    }
  }

  return `${delimiters[0]}${frontMatter}${delimiters[1]}`;
};

/**
 * Get post template
 * @param {object} properties - JF2 properties
 * @param {string} [frontMatterFormat] - Front matter format
 * @returns {string} Rendered template
 */
export const getPostTemplate = (properties, frontMatterFormat = "yaml") => {
  const { content, tags } = replaceTags(getContent(properties));
  if (tags.length > 0) {
    properties.tags = tags;
  }
  const frontMatter = getFrontMatter(properties, frontMatterFormat);

  return frontMatter + content;
};

const hashtagFinder = /(?<=\s|^)#[a-z0-9]+\b/ig;
const peopletagFinder = /(?<=\s|^)@[a-z]+\b/ig;

/**
 * Replaces all #HashTags with [HashTags](/tags/hashtags) and returns a compact array of the tags used
 * @param {string} content - Post content
 * @returns {{content: string, tags: Array<string>}} The replaced content, and found tags
 */
const replaceTags = (content) => {
  let tags = [];
  content = content.replace(hashtagFinder, (tag) => {
    tag = tag.substring(1);
    tags.push(tag);
    return `[${tag}](/tags/${tag.toLowerCase()})`;
  }).replace(peopletagFinder, (tag) => {
    return `{{< friend "${tag.substring(1).toLowerCase()}" >}}`;
  });

  return { content, tags }
}