package main

import (
	"fmt"
	"strings"
)

type Hotel struct {
	Name  string
	IsVeg string
}

var hotels []Hotel

func Init() {
	h1 := Hotel{"A", "false"}
	h2 := Hotel{"B", "true"}
	hotels = []Hotel{h1, h2}

}
func GetRestaurant(name string, isVeg string) Hotel {
	nameMatch := []Hotel{}
	for _, h := range hotels {
		if strings.Contains(name, h.Name) {
			nameMatch = append(nameMatch, h)
		}
	}
	for _, v := range nameMatch {
		fmt.Printf("Check veg or non neg")
		if strings.EqualFold(v.IsVeg, isVeg) {
			fmt.Printf("found")
			return v
		}
	}
	return Hotel{}
}
