package cli

import (
	"flag"
	"fmt"
)

func Run() {
	data := flag.String("data", "", "some data...")
	flag.Parse()

	if *data != "" {
		fmt.Printf("data:%s\n", *data)
	} else {
		flag.Usage()
	}
}
