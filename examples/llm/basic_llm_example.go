package main

import (
	"fmt"
	"log"

	"github.com/twwch/langchain-go/llms/openai"
)

func main() {
	llm, err := openai.New()
	if err != nil {
		log.Fatal(err)
	}
	completion, err := llm.Call("What would be a good company name a company that makes colorful socks?")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(completion)
}
