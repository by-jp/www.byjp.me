package shared

import (
	"regexp"
	"strings"
)

type SyndicationID struct {
	Source string
	ID     string
}

func (sid SyndicationID) String() string {
	return "syndicate:" + sid.Source + ":" + sid.ID
}

func (sid SyndicationID) Bytes() []byte {
	return []byte(sid.String())
}

func TagMatcher(serviceTags []string) (*regexp.Regexp, error) {
	return regexp.Compile("syndicate:(" + strings.Join(serviceTags, "|") + `):(\S+)`)
}
