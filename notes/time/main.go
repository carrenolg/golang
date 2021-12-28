package main

import (
	"fmt"
	"time"
)

func main() {
	//datetime string
	dts := "2021-09-17T16:03:18+0000"
	//layaout RFC3339 with -0700 (Numeric Time Zone)
	layout := "2006-01-02T15:04:05-0700"
	t, err := time.Parse(layout, dts)
	fmt.Println(t.UTC().String(), err)

	// Example 2
	dts = "2021-12-27T15:55:39Z"
	t, err = time.Parse(time.RFC3339, dts)
	fmt.Println(t.UTC().String(), err)

	// Custom layout (Ubuntu datetime)
	customLayout := "Mon 02 Jan 2006 15:04:05 -07"
	dts = t.Format(customLayout)
	fmt.Println(dts)
}
