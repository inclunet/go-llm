package main

import (
	"fmt"

	"github.com/inclunet/go-llm/openai"
)

func main() {
	apiToken := "Your-token-here"

	// Start a new client with your API key
	// Visit https://beta.openai.com/docs/api-reference/authentication to learn more about API keys
	client := openai.NewClient(apiToken, "")

	// Create a new completion request
	// Visit https://beta.openai.com/docs/api-reference/completions/create to learn more about completion requests
	completions := openai.NewCompletions()

	// Create a new message to be sent to the API
	// Visit https://beta.openai.com/docs/api-reference/completions/create#completionrequestmessage to learn more about messages
	message := openai.NewMessage("Hello chat GPT 3, how are you?")

	// Add a message to the completion request
	// Visit https://beta.openai.com/docs/api-reference/completions/create#completionrequestmessages to learn more about messages
	completions.AddMessage(message)

	// Send the completion request to the API
	// Visit https://beta.openai.com/docs/api-reference/completions/create to learn more about completion requests
	response, err := client.Completions(completions)

	// Handle errors
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s: %s\n", response.Choices[0].Message.Role, response.Choices[0].Message.Content)

	// The assistant message is automatically added to the completion request
	// if you want to add more messages, you can do so by calling completions.AddMessage(message) and then calling client.Completions(completions) again
	// all messages will be sent to the API including the assistant message that was added automatically by context

	message = openai.NewMessage("We can talk about anything you want. What do you want to talk about?")

	// Add a message to the completion request
	// Visit https://beta.openai.com/docs/api-reference/completions/create#completionrequestmessages to learn more about messages
	completions.AddMessage(message)

	// Send the completion request to the API
	// Visit https://beta.openai.com/docs/api-reference/completions/create to learn more about completion requests
	response, err = client.Completions(completions)

	// Handle errors
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s: %s\n", response.Choices[0].Message.Role, response.Choices[0].Message.Content)
}
