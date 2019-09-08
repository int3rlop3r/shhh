package main

import "time"

type Group struct {
	Name      string
	Records   []Record
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Record struct {
	Title     string
	User      string
	Pass      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (g *Group) AddRec(r *Record) {
	g.Records = append(g.Records, *r)
}
