package common

import (
	"fmt"
	"log"
	"strings"
)

//ParseTime parses MAUDE dates to "YY-MM-DD"
func ParseDate(item string) string {
	//MAUDE dates look like: "YYMMDD"

	if len(item) < 1 {
		log.Fatal("Tried to parse empty string")
	}

	itemSlice := strings.Split(item, "")

	year := strings.Join(itemSlice[:4], "")
	month := strings.Join(itemSlice[4:6], "")
	day := strings.Join(itemSlice[6:8], "")

	date := fmt.Sprintf("%v-%v-%v", year, month, day)

	return date
}
