package entity

import (
	"fmt"
	"github.com/google/uuid"
	"strconv"
	"strings"
)

// GenerateID generates a unique ID that can be used as an identifier for an entity.
func GenerateID() string {
	return uuid.New().String()
}

type Ids []int64

func NewIds(ids []int64) Ids {
	return ids
}

func (a Ids) WhereIn(column string) string {
	var strs []string
	for i, num := range a {
		strs[i] = strconv.Itoa(int(num))
	}
	return fmt.Sprintf("WHERE %s IN (%s) ", column, strings.Join(strs, ", "))
}
