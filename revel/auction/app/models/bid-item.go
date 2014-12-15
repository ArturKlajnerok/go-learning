package models

type BidItem struct {
	Id             int64   `db:"id" json:"id"`
	Name           string  `db:"name" json:"name"`
	Category       string  `db:"category" json:"category"`
	EstimatedValue float32 `db:"est_value" json:"est_value"`
	StartBid       float32 `db:"start_bid" json:"start_bid"`
	BidIncrement   float32 `db:"bid_incr" json:"bid_incr"`
	InstantBuy     float32 `db:"inst_buy" json:"inst_buy"`
}
