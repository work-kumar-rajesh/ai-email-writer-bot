package aiagent

type AIAgent interface {
	GenerateEmailReply(prompt string) (string, error)
}
