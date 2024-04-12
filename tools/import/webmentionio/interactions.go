package main

import (
	synd "github.com/by-jp/www.byjp.me/tools/syndicate/shared"
)

type InteractionFile struct {
	Interactions []synd.Interaction `json:"interactions"`
}

type ByTimestamp []synd.Interaction

// Implement the sort.Interface - Len, Less, and Swap methods
func (a ByTimestamp) Len() int           { return len(a) }
func (a ByTimestamp) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTimestamp) Less(i, j int) bool { return a[i].Timestamp.Unix() < a[j].Timestamp.Unix() }
