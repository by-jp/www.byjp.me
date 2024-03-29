package main

import (
	synd "github.com/by-jp/www.byjp.me/tools/syndicate/shared"
)

type InteractionFile struct {
	Interactions []synd.Interaction `json:"interactions"`
}
