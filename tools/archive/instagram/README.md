# Instagram archive

This tool will take an [Instagram data archive](https://help.instagram.com/181231772500920) and turn it into a series of timestamped posts for your Hugo blog.

```sh
go run . <path/to/archive/username_yyyymmdd.zip> <path/to/hugo/root/>
```

It will create one folder per post in the `content/photos` directory. You can change how those posts look by creating a `layouts/instgram-posts/single.html` template.
