package ticker

import "time"

type TimeRange struct {
	Start time.Time `json:"start" form:"start" binding:"required" time_utc:"1"`
	End   time.Time `json:"end" form:"end" binding:"required" time_utc:"1"`
}

type DateValidator struct {
	Date   time.Time `uri:"date" binding:"required" time_utc:"1"`
	Symbol string    `uri:"symbol" binding:"required"`
}
