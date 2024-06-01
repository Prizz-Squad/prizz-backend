package types

type File struct {
	ID       string `json:"id"`
	File     []byte `json:"file"`
	TaskID   string `json:"taskId"`
	FileName string `json:"filename"`
}
type FileCreateRequest struct {
	File     []byte `json:"file"`
	TaskID   string `json:"taskId"`
	FileName string `json:"filename"`
}
type FileResponse struct {
	ID string `json:"id"`
}
