package entity

type TopUp struct {
	ID      string  `db:"id"`
	UserID  int64   `db:"user_id"`
	Status  int8    `db:"status"`
	Amount  float64 `db:"amount"`
	SnapURL string  `db:"snap_url"`
}

type TopUpRequest struct {
	Amount float64 `json:"amount"`
	UserID int64   `json:"user_id"`
}

type TopUpResponse struct {
	SnapURL string `json:"snap_url"`
}
