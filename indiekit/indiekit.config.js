const store = process.env.GITHUB_TOKEN ? "@indiekit/store-github" : "@indiekit/store-file-system"

export default {
  application: {
    name: "IndieKit byJP",
    mongodbUrl: process.env.MONGO_URL,
    port: process.env.PORT ?? 8808,
    timeZone: 'Europe/London',
  },
  plugins: [
    store,
    "indiekit-preset-byjp",
    "@indiekit/syndicator-mastodon",
    "@indiekit/endpoint-files",
    "@indiekit/endpoint-image",
    "@indiekit/endpoint-micropub",
    "@indiekit/endpoint-posts",
    "@indiekit/endpoint-share",
    "@indiekit/post-type-bookmark",
    "@indiekit/post-type-event",
    "@indiekit/post-type-like",
    "@indiekit/post-type-note",
    "@indiekit/post-type-photo",
    "@indiekit/post-type-jam",
    "@indiekit/post-type-reply"
  ],
  publication: {
    me: process.env.SITE,
    enrichPostData: true,
    postTypes: {
      event: {
        name: "Event",
        post: {
          path: "content/calendar/{yyyy}-{MM}-{dd}/{slug}/index.md",
          url: "calendar/{yyyy}-{MM}-{dd}/{slug}"
        }
      },
      like: {
        name: "Like",
        post: {
          path: "content/likes/{yyyy}-{MM}-{dd}/{slug}/index.md",
          url: "likes/{yyyy}-{MM}-{dd}/{slug}/"
        }
      },
      note: {
        name: "Note",
        post: {
          path: "content/notes/{yyyy}-{MM}-{dd}/{slug}/index.md",
          url: "notes/{yyyy}-{MM}-{dd}/{slug}/"
        }
      },
      reply: {
        name: "Reply",
        post: {
          path: "content/notes/{yyyy}-{MM}-{dd}/{slug}/index.md",
          url: "notes/{yyyy}-{MM}-{dd}/{slug}/"
        }
      },
      photo: {
        name: "Photo",
        post: {
          path: "content/photos/{yyyy}-{MM}-{dd}/{slug}/index.md",
          url: "photos/{yyyy}-{MM}-{dd}/{slug}/"
        },
        media: {
          path: "content/photos/{yyyy}-{MM}-{dd}/{filename}",
          url: "photos/{yyyy}-{MM}-{dd}/{filename}"
        }
      }
    }
  },
  "jphastings/indiekit-preset-byjp": {},
  "@indiekit/store-github": {
    user: "by-jp",
    repo: "www.byjp.me",
    branch: "main"
  },
  "@indiekit/post-type-bookmark": {
    name: "Bookmark"
  },
  "@indiekit/post-type-photo": {
    name: "Photo"
  },
  "@indiekit/post-type-jam": {
    name: "Jam"
  },
  "@indiekit/post-type-like": {
    name: "Like"
  },
  "@indiekit/post-type-note": {
    name: "Note"
  },
  "@indiekit/post-type-reply": {
    name: "Reply"
  },
  "@indiekit/endpoint-micropub": {},
  "@indiekit/endpoint-share": {},
  "@indiekit/endpoint-files": {},
  "@indiekit/endpoint-media": {
    "imageProcessing": {
      "resize": {
        "width": 1280,
        "height": 1280,
        "options": {
          "fit": "contain",
        },
      },
    }
  },
  "@indiekit/endpoint-posts": {},
  "@indiekit/store-file-system": {
    "directory": ".."
  }
}
