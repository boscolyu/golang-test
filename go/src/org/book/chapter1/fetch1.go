package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr,  "fetch : %v\n", err)
			os.Exit(1)
		}

		output, err := os.Open(url + ".txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "open fail  : %v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch : reading %s: %v\n", url, err)
			os.Exit(1)
		}

		_, err = output.Write(b)
		if err != nil {
			fmt.Fprintf(os.Stderr, "write fail : %v\n", err)
		}
		fmt.Printf("%s", b)
		output.Close()
	}
}