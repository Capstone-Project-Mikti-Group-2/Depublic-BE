package entity

type FilterEvent struct {
	Sort   SortQuery
	Filter FilterQuery
	Search string
}

type SortQuery struct {
	By    string
	Order string
}

type FilterQuery struct {
	Price     float64
	StartDate string
	EndDate   string
	Location  string
	Available bool
}
