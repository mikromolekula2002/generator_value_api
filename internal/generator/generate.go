package generator

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type GenValue struct {
}

type Generate interface {
	RequestID() string
	NumericValue(length int) string
	AlphaNumericValue(length int) string
	StringValue(length int) string
}

func (g *GenValue) RequestID() string {
	id := uuid.New()

	fullID := id.String()

	shortUUID := fullID[:8]

	return shortUUID
}

func (g *GenValue) NumericValue(leng int) string {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	numbers := "0123456789"
	result := make([]byte, leng)

	for i := range result {
		result[i] = numbers[r.Intn(len(numbers))]
	}

	return string(result)
}

func (g *GenValue) AlphaNumericValue(leng int) string {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, leng)

	for i := range result {
		result[i] = letters[r.Intn(len(letters))]
	}

	return string(result)
}

func (g *GenValue) StringValue(leng int) string {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, leng)

	for i := range result {
		result[i] = letters[r.Intn(len(letters))]
	}

	return string(result)
}
