import { getPostTemplate } from "./lib/post-template.js";
import { getPostTypes } from "./lib/post-types.js";

const defaults = {};

export default class HugoPreset {
  constructor(options = {}) {
    this.name = "byJP preset";
    this.options = { ...defaults, ...options };
  }

  get info() {
    return {
      name: "byJP",
    };
  }

  get prompts() {
    return [];
  }

  postTemplate(properties) {
    return getPostTemplate(properties);
  }

  init(Indiekit) {
    const { application } = Indiekit.config;
    this.postTypes = getPostTypes(application.postTypes);

    Indiekit.addPreset(this);
  }
}
