package task

type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Concluded   bool   `json:"concluded"`
}
