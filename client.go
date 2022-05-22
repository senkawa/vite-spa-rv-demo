package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/buger/jsonparser"
)

func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	uri, _ := url.Parse("http://localhost:8000/new")
	param := url.Values{}
	param.Add("body", string(output))
	uri.RawQuery = param.Encode()

	res, err := http.Get(uri.String())
	if err != nil {
		log.Fatalln("request failed")
	}

	body, _ := ioutil.ReadAll(res.Body)
	id, err := jsonparser.GetInt(body, "data", "ID")
	if err != nil {
		log.Fatalln("no id found", string(body))
	}

	fmt.Printf("http://localhost:8000/paste/%d\n", id)
}
