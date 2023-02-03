package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/moganxumerle/study-go-testing/words"
)

var (
	file   string
	filter string
)

func init() {
	flag.StringVar(&file, "file", "", "arquivo com texto")
	flag.StringVar(&filter, "filter", "", "filtra pela palavra")
	flag.Parse()
}

func main() {
	if file == "" {
		panic("missing file")
	}

	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	fmt.Println("Total: ", words.Count(string(content), filter))

}
