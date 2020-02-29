package controller

import (
	"ip_proxy/controller/play"
)

func Add(a, b int) int {
	return a + b
}

func Compute(a, b int) int {
	return Add(a, b) + play.Multi(a, b)
}
