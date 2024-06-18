<?xml version="1.0" encoding="utf-8"?>
<xsl:stylesheet version="3.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform" xmlns:atom="http://www.w3.org/2005/Atom" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd">
  <xsl:output method="html" version="1.0" encoding="UTF-8" indent="yes"/>
  <xsl:template match="/">
    <html xmlns="http://www.w3.org/1999/xhtml" lang="en" dir="ltr">
      <head>
        <title><xsl:value-of select="/opml/head/title"/></title>
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
              <hgroup>
                <h1><xsl:value-of select="/opml/head/title"/></h1>
                <p>A collection of feeds from elsewhere on the internet that you may enjoy. Many feed readers will accept OPML feeds (like this one) as sources, allowing you to subscribe to the feeds below <em>and</em> any I add in the future.</p>
              </hgroup>
              <a hreflang="en" target="_blank">
                <xsl:attribute name="href">./</xsl:attribute>
                Visit this page on my site for more information &#x2192;
              </a>
            </header>
            <main>
              <xsl:for-each select="/opml/body/outline">
                <section>
                  <hgroup>
                    <h2>
                      <a hreflang="en" target="_blank">
                        <xsl:attribute name="href">
                          <xsl:value-of select="@htmlUrl"/>
                        </xsl:attribute>
                        <xsl:value-of select="@title"/>
                      </a>
                    </h2>
                  </hgroup>
                  <p>
                    <xsl:value-of select="@description"/>
                  </p>

                  <ul>
                    <xsl:for-each select="outline">
                      <li>
                        <a hreflang="en" target="_blank">
                          <xsl:attribute name="href">
                            <xsl:value-of select="@htmlUrl"/>
                          </xsl:attribute>
                          <xsl:value-of select="@title"/>
                        </a> — <xsl:value-of select="@description"/>
                      </li>
                    </xsl:for-each>
                  </ul>
                </section>
              </xsl:for-each>
            </main>
          </div>
        </div>
      </body>
    </html>
  </xsl:template>
</xsl:stylesheet>