package main

import (
	"fmt"
	"os"

	"github.com/sushil/jsoncompare"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("usage: %s  <%s>  <%s>\n", os.Args[0], "first file path", "second file path")
		return
	}

	result, err := jsoncompare.CompareFiles(os.Args[1], os.Args[2])

	if err != nil {
		panic(err)
	}

	if result.IsEqual {
		fmt.Println("both are same")
	} else {
		fmt.Printf("\n%s\n%s\n%s",
			"given json contents are not same\n\n -- first:\n\n%s \n\n -- second:\n\n%s \n\n",
			result.FirstNodePaths,
			result.SecondNodePaths)
	}
}
