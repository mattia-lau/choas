package ticker

import "time"

type TimeRange struct {
	Start time.Time `json:"start" form:"start" binding:"required" time_format:"2006-01-02" time_utc:"1"`
	End   time.Time `json:"end" form:"end" binding:"required" time_format:"2006-01-02" time_utc:"1"`
}
