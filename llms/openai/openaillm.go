package openai

import (
	"context"
	"errors"
	"os"

	"github.com/twwch/langchain-go/llms"
	"github.com/twwch/langchain-go/llms/openai/internal/openaiclient"
)

const tokenEnvVarName = "OPENAI_API_KEY"

var ErrEmptyResponse = errors.New("no response")
var ErrMissingToken = errors.New("missing the OpenAI API key, set it in the OPENAI_API_KEY environment variable")

type LLM struct {
	client *openaiclient.Client
}

var _ llms.LLM = (*LLM)(nil)

func (o *LLM) Call(prompt string) (string, error) {
	r, err := o.Generate([]string{prompt})
	if err != nil {
		return "", err
	}
	if len(r) == 0 {
		return "", ErrEmptyResponse
	}
	return r[0].Text, nil
}

func (o *LLM) Generate(prompts []string) ([]*llms.Generation, error) {
	result, err := o.client.CreateCompletion(context.TODO(), &openaiclient.CompletionRequest{
		Prompt: prompts[0],
	})
	if err != nil {
		return nil, err
	}
	return []*llms.Generation{
		{Text: result.Text},
	}, nil
}

func New() (*LLM, error) {
	token := os.Getenv(tokenEnvVarName)
	if token == "" {
		return nil, ErrMissingToken
	}
	c, err := openaiclient.New(token)
	return &LLM{
		client: c,
	}, err
}
