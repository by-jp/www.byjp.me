package main

import (
	"fmt"
	"os"
	"path"
	"regexp"

	"github.com/by-jp/www.byjp.me/tools/shared"
)

type Config struct {
	Root string `default:"."`
}

var postsFileRE = regexp.MustCompile(
	`(facebook-\w+-\d{4}-\d{2}-\d{2}-\w{8})/your_facebook_activity/posts/your_posts__check_ins__photos_and_videos_\d+.json`,
)

func main() {
	var c Config
	shared.Check(shared.LoadConfig("facebook", &c), "Unable to load config")

	if len(os.Args) < 2 {
		shared.ExitMessage(
			"Usage: %s facebook-yourname-YYYY-MM-DD-xxxxxx-YYYYMMDDTHHMMSS-001.zip",
			path.Base(os.Args[0]),
		)
	}
	archives, err := findArchives(os.Args[1])
	shared.Check(err, "Couldn't find facebook archive zipfile")

	shared.Check(
		shared.ExtractPostsFromArchives(
			archives,
			postsFileRE,
			c.Root,
			postize,
		),
		"Unable to extract posts from archive files",
	)

	fmt.Println("Your Facebook export has been imported into your blog.")
}
