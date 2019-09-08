package main

import "time"

func NewGroup(name string) *Group {
	return &Group{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewRecord(title, user, pass string) *Record {
	return &Record{
		Title:     title,
		User:      user,
		Pass:      pass,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
