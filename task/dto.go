package task

type CreateTaskTypeRequest struct {
	Name      string `json:"name"`
	CreatedBy string `json:"created_by"`
}
