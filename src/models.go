package lighthouse

import (
	"time"
)

type analyticRecord struct {
	Address string
	time time.Time
}

type statModel struct {
	Count int64 `json:"count"`
}

type BadgeSvgData struct{
	Count string
}