package message_manager

type MessageManager struct {
	storage storage
}

func New(storage storage) *MessageManager {
	return &MessageManager{storage: storage}
}
