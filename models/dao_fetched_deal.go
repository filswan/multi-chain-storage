package models

type DaoFetchedDeal struct {
	ID       int64 `json:"id"`
	DealId   int64 `json:"deal_id"`
	CreateAt int64 `json:"create_at"`
}
