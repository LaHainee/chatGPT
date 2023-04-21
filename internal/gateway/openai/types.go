package openai

const (
	defaultModel = "gpt-3.5-turbo"
	defaultRole  = "user"
)

type chatRequest struct {
	Model    string        `json:"model"`
	Messages []chatMessage `json:"messages"`
}

type chatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type chatResponse struct {
	Choices []chatChoice `json:"choices"`
}

type chatChoice struct {
	Message chatChoiceMessage `json:"message"`
}

type chatChoiceMessage struct {
	Content string `json:"content"`
}
