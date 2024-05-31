package types

type Project struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CreatorID   string `json:"creatorId"`
	Description string `json:"description"`
}
type ProjectCreateRequest struct {
	Name        string `json:"name"`
	CreatorID   string `json:"creatorId"`
	Description string `json:"description"`
}

type ProjectResponse struct {
	ID string `json:"id"`
}
