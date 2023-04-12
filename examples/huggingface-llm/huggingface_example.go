package main

import (
	"fmt"
	"log"

	"github.com/twwch/langchain-go/llms/huggingface"
)

func main() {
	llm, err := huggingface.New()
	if err != nil {
		log.Fatal(err)
	}
	completion, err := llm.Call("What would be a good company name be for name a company that makes colorful socks?")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(completion)
}
