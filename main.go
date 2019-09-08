package main

import "fmt"

func main() {
	group := NewGroup("Test group")
	record1 := NewRecord("Primary email", "a@x.y", "pass1")
	group.AddRec(record1)
	record2 := NewRecord("Secondary email", "b@x.y", "pass1")
	group.AddRec(record2)

	fmt.Println("Printing entries for group:", group.Name)
	for _, r := range group.Records {
		fmt.Println("Title:", r.Title)
	}
}
