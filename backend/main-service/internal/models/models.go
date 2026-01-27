// Package models defines application data structures
package models

// Link represents shortened URL entity
type Link struct {
	OriginLink string `json:"origin_link"`
	PseudoLink string `json:"pseudo_link"`
}
