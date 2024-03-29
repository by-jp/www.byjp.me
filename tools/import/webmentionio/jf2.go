package main

type Mentions struct {
	Type     string
	Name     string
	Children []Mention
}

type Mention struct {
	Type   string
	Author Author
	URL    string
	// Fallback on WebmentionReceived as fallback for Published
	Published          string // (likes often don't have timestamps)
	WebmentionReceived string `json:"wm-received"`

	WebmentionID     uint   `json:"wm-id"`
	WebmentionTarget string `json:"wm-target"`
	Content          struct {
		Text string
	}
	Summary struct {
		ContentType string `json:"content-type"`
		Value       string
	}

	// The kind of action, eg. like-of, in-reply-to, repost-of
	WebmentionProperty string `json:"wm-property"`
	WebmentionPrivate  bool   `json:"wm-private"`
}

type Author struct {
	Type  string
	Name  string
	Photo string
	URL   string
}
