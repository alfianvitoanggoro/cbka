package dto

type RequestUserReconcile struct {
	RequestID   *string `json:"request_id,omitempty"`
	Timestamp   *int64  `json:"timestamp,omitempty"`
	TriggeredBy string  `json:"triggered_by"`
}
