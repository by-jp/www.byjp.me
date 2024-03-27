import camelcaseKeys from "camelcase-keys";
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
 * @returns {string} Front matter in chosen format
 */
const getFrontMatter = (properties) => {
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
      images: properties.photo.map((image) => image.url),
    }),
    ...properties,
  };

  delete properties.content; // Shown below front matter
  delete properties.deleted; // Use `expiryDate`
  delete properties.name; // Use `title`
  delete properties.postStatus; // Use `draft`
  delete properties.published; // Use `date`
  delete properties.type; // Not required
  delete properties.updated; // Use `lastmod`
  delete properties.url; // Not required

  const frontMatter = YAML.stringify(properties, { lineWidth: 0 });
  return `---\n${frontMatter}---\n`;
};

/**
 * Get post template
 * @param {object} properties - JF2 properties
 * @param {string} [frontMatterFormat] - Front matter format
 * @returns {string} Rendered template
 */
export const getPostTemplate = (properties) => {
  const content = getContent(properties);
  const frontMatter = getFrontMatter(properties);

  return frontMatter + content;
};
