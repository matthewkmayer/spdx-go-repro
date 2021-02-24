package main

import (
	"fmt"
	"os"

	"github.com/spdx/tools-golang/spdxlib"
	"github.com/spdx/tools-golang/tvloader"
)

func main() {
	// change to "./spdx-test-2.spdx" and see it work
	r, err := os.Open("./spdx-test-1.spdx")
	if err != nil {
		fmt.Printf("Error while opening file for reading: %v", err)
		return
	}
	defer r.Close()

	doc, err := tvloader.Load2_2(r)
	if err != nil {
		fmt.Printf("Error while parsing file: %v", err)
		return
	}

	p, err := spdxlib.GetDescribedPackageIDs2_2(doc)
	if err != nil {
		fmt.Printf("\nGot an error and should not have. Error: %v", err.Error())
		return
	}
	fmt.Printf("\ngot packages: %+v", p)
}
