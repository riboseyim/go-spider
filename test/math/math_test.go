package main_test

import (
	"log"
	"testing"
)

func TestCountInt(t *testing.T) {
	aa := 100
	bb := 200
	cc := aa / bb
	log.Printf("aa / bb = %d", cc)
}

func TestCountFloat(t *testing.T) {
	aa := 100.0
	bb := 200.0
	cc := aa / bb
	log.Printf("aa / bb = %v", cc)
}
