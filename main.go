package main

import (
	"fmt"
	"github.com/jcomo/go-nextbus/nextbus"
)

func main() {
	c, _ := nextbus.NewClient()
	resp, err := c.PredictionsForStopTag("sf-muni", "KT", "7166")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp)
}
