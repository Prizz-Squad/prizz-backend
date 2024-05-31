package types

type Message struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	UserId      string `json:"user_id"`
	TaskId      string `json:"task-id_id"`
	Contents    string `json:"contents"`
}

type MessageRequest struct {
	Description string `json:"description"`
	UserId      string `json:"user_id"`
	TaskId      string `json:"task-id_id"`
	Contents    string `json:"contents"`
}

type MessageResponse struct {
	Description string `json:"description"`
	UserId      string `json:"user_id"`
	TaskId      string `json:"task-id_id"`
	Contents    string `json:"contents"`
}
