package models

type DaoInfo struct {
	ID          int64  `json:"id"`
	DaoName     string `json:"dao_name"`
	DaoAddress  string `json:"dao_address"`
	OrderIndex  string `json:"order_index"`
	Description string `json:"description"`
	CreateAt    string `json:"create_at"`
}
