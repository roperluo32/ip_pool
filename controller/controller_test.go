package controller

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestCompute(t *testing.T) {
	r := Compute(1, 3)
	assert.Equal(t, 7, r)
}