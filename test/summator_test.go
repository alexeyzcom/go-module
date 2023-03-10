package test

import (
	"github.com/alexeyzcom/go-module/internal/service"
	"testing"
)

func TestAdd(t *testing.T) {
	got := service.Add(3, 5)
	want := 8
	if got != want {
		t.Errorf("Got %d, but want %d", got, want)
	}
}
