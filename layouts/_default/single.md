{{- /*
  Raw-markdown output format.

  Emits the page's original front matter followed by its body with all
  shortcodes/includes expanded (via .RenderShortcodes) but the markdown
  otherwise left intact. This is what sequoia.pub consumes to build
  site.standard.document atproto records, so includes resolve to their
  replacement text rather than leaking `{{ "{{<" }} include >}}` notation.

  This output format is isHTML, so shortcodes without a dedicated template
  fall back to their .html version (valid inline HTML in markdown) rather
  than the gemini .gmi version. safeHTML keeps the html/template engine from
  escaping the front matter and body. To give a shortcode a cleaner markdown
  rendering, add a `name.markdown.md` shortcode template; nothing else needed.
*/ -}}
{{- with .File -}}
  {{- $raw := readFile (path.Join "content" .Path) -}}
  {{- $fm := findRE `(?s)^---\n.*?\n---` $raw 1 -}}
  {{- with $fm -}}{{- index . 0 | safeHTML -}}{{- "\n\n" -}}{{- end -}}
{{- end -}}
{{- .RenderShortcodes | safeHTML -}}
