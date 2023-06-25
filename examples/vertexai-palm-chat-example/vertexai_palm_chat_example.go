package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms/vertexai"
	"github.com/tmc/langchaingo/schema"
)

func main() {
	llm, err := vertexai.NewChat()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	question := schema.HumanChatMessage{
		Text: "What would be a good company name a company that makes colorful socks?",
	}
	fmt.Println("---ASK---")
	fmt.Println(question.GetText())
	messages := []schema.ChatMessage{question}
	completion, err := llm.Call(ctx, messages)
	if err != nil {
		log.Fatal(err)
	}

	response := completion
	fmt.Println("---RESPONSE---")
	fmt.Println(response)

	// keep track of conversation
	messages = append(messages, &schema.AIChatMessage{
		Text: response,
	})

	question = schema.HumanChatMessage{
		Text: "Any other recommendation on how to get started with the company ?",
	}
	fmt.Println("---ASK---")
	fmt.Println(question.GetText())
	messages = append(messages, question)

	completion, err = llm.Call(ctx, messages)
	if err != nil {
		log.Fatal(err)
	}

	response = completion
	fmt.Println("---RESPONSE---")
	fmt.Println(response)
}
