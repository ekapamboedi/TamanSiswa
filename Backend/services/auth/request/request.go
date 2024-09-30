package request

type CreateRequest struct {
	PositionId   string           `json:"positionId"`
	PositionName string           `json:"positionName"`
	Properties   map[string]int64 `json:"properties"`
	Deductions   map[string]int64 `json:"deductions"`
}

type UpdateRequest struct {
	Properties map[string]int64 `json:"properties"`
	Deductions map[string]int64 `json:"deductions"`
}
