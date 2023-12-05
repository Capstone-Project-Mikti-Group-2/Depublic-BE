package entity

type TopUp struct {
	ID      string  `json:"id"`
	UserID  int64   `json:"user_id"`
	Status  int8    `json:"status"`
	Amount  float64 `json:"amount"`
	SnapURL string  `json:"snap_url"`
}

type TopUpRequest struct {
	Amount float64 `json:"amount"`
	UserID int64   `json:"user_id"`
}

type TopUpResponse struct {
	SnapURL string `json:"snap_url"`
}
