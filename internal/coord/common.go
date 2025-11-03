package coord

import "github.com/mimimc4/tchat/pkg/utils"

type ChatMessage string

type Node struct {
	msgBuffer *utils.CircularBuffer[ChatMessage]
}
