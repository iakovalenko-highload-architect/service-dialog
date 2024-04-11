package models

type Message struct {
	ID       string
	DialogID string
	FromID   string
	ToID     string
	Text     string
}
