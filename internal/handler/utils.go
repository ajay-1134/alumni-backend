package handler

import (
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type DateOnly struct {
	time.Time
}

func (d *DateOnly) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = s[1 : len(s)-1] // remove quotes
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return errors.New("invalid date format, use YYYY-MM-DD")
	}
	d.Time = t
	return nil
}

// Convert back to JSON string
func (d DateOnly) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.Time.Format("2006-01-02") + `"`), nil
}

func stringToUint(s string) (uint, error) {
	val64, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(val64), nil
}

func getId(c *gin.Context) (uint, error) {
	id := c.Param("id")

	if id != "" {
		id := c.Param("id")
		return stringToUint(id)
	}

	userId, exists := c.Get("userID")
	if !exists {
		return 0, errors.New("userID key does not exist")
	}

	return userId.(uint), nil
}
