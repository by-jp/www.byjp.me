<?xml version="1.0" encoding="utf-8"?>
<xsl:stylesheet version="3.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:atom="http://www.w3.org/2005/Atom" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd">
  <xsl:output method="html" version="1.0" encoding="UTF-8" indent="yes"/>
  <xsl:template match="/">
    <html xmlns="http://www.w3.org/1999/xhtml" lang="en" dir="ltr">
      <head>
        <title><xsl:value-of select="/rss/channel/title"/> RSS Feed</title>
        <meta charset="UTF-8" />
        <meta http-equiv="x-ua-compatible" content="IE=edge,chrome=1" />
        <meta http-equiv="content-language" content="en_US" />
        <meta name="viewport" content="width=device-width,minimum-scale=1,initial-scale=1,shrink-to-fit=no" />
        <meta name="referrer" content="none" />
        <style type="text/css">
          body {
            color: #222;
            font-family: "Segoe UI", apple-system, BlinkMacSystemFont, Futura, Roboto, Arial, system-ui, sans-serif;
          }
          .container {
            align-item: center;
            display: flex;
            justify-content: center;
          }
          h3 {
            margin-block-end: 0.2em;
          }
          hgroup {
            margin-block-end: 1em;
          }
          time {
            display: block;
            font-style: italic;
            font-size: small;
          }
          .item {
            max-width: 768px;
          }
          a {
            color: #4166f5;
            text-decoration: none;
          }
          a:visited {
            color: #3f00ff;
          }
          a:hover {
            text-decoration: underline;
          }
        </style>
      </head>
      <body>
        <div class="container">
          <div class="item">
            <header>
              <h1>RSS Feed</h1>
              <h2>
                <xsl:value-of select="/rss/channel/title"/>
              </h2>
              <p>
                <xsl:value-of select="/rss/channel/description"/>
              </p>
              <a hreflang="en" target="_blank">
                <xsl:attribute name="href">
                  <xsl:value-of select="/rss/channel/link"/>
                </xsl:attribute>
                Visit Website &#x2192;
              </a>
            </header>
            <main>
              <h2>Recent Posts</h2>
              <xsl:for-each select="/rss/channel/item">
                <article>
                  <hgroup>
                    <h3>
                      <a hreflang="en" target="_blank">
                        <xsl:attribute name="href">
                          <xsl:value-of select="link"/>
                        </xsl:attribute>
                        <xsl:value-of select="title"/>
                      </a>
                    </h3>
                    <time>
                      <xsl:attribute name="datetime">
                        <xsl:value-of select="pubDate" /> 
                      </xsl:attribute>
                      <xsl:value-of select="substring(pubDate, 0, 17)" />
                    </time>
                  </hgroup>
                  <p><xsl:value-of select="description"/></p>
                </article>
              </xsl:for-each>
            </main>
          </div>
        </div>
      </body>
    </html>
  </xsl:template>
</xsl:stylesheet>