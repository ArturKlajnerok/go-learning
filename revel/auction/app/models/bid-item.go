package models

import (
	"github.com/revel/revel"
	"regexp"
)

type BidItem struct {
	Id             int64   `db:"id" json:"id"`
	Name           string  `db:"name" json:"name"`
	Category       string  `db:"category" json:"category"`
	EstimatedValue float32 `db:"est_value" json:"est_value"`
	StartBid       float32 `db:"start_bid" json:"start_bid"`
	BidIncrement   float32 `db:"bid_incr" json:"bid_incr"`
	InstantBuy     float32 `db:"inst_buy" json:"inst_buy"`
}

func (b *BidItem) Validate(v *revel.Validation) {

	v.Check(b.Name,
		revel.ValidRequired(),
		revel.ValidMaxSize(25))

	v.Check(b.Category,
		revel.ValidRequired(),
		revel.ValidMatch(
			regexp.MustCompile(
				"^(travel|leasure|sports|entertainment)$")))

	v.Check(b.EstimatedValue,
		revel.ValidRequired())

	v.Check(b.StartBid,
		revel.ValidRequired())

	v.Check(b.BidIncrement,
		revel.ValidRequired())
}
