package message_manager

type MessageManager struct {
	storage storage
	cache   cache
}

func New(storage storage, cache cache) *MessageManager {
	return &MessageManager{
		storage: storage,
		cache:   cache,
	}
}
