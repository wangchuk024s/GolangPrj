package model

import (
	"math/rand"
	"strconv"
	"time"
)

// URLMapping represents a URL mapping
type URLMapping struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ShortURL  string `json:"shortUrl"`
	Details   string `json:"details"`
	Timestamp int64  `json:"timestamp"`
}

// GenerateID generates a random ID for the URL mapping
func GenerateID() string {
	timestamp := time.Now().UnixNano()
	randomNumber := rand.Intn(10000)
	return time.Unix(0, timestamp).Format("20060102") + "-" + strconv.Itoa(randomNumber)
}
