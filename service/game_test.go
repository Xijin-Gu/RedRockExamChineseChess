package service

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestKingMove(t *testing.T) {
	got := KingMove(10,10,"11")
	assert.Equal(t,got,false)
}





