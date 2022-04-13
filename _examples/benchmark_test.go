package main

import (
	"context"
	"testing"

	"github.com/gabrielsoaressantos/go-mediator/pkg/mediator"
)

func BenchmarkMediator(b *testing.B) {
	m, _ := mediator.New(mediator.WithHandler(&FakeCommand{}, NewFakeCommandHandler()))

	cmd := &FakeCommand{Name: "Emre"}
	ctx := context.TODO()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Send(ctx, cmd)
	}
}

func BenchmarkHandler(b *testing.B) {
	handler := NewFakeCommandHandler()
	cmd := &FakeCommand{Name: "Emre"}
	ctx := context.TODO()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		handler.Handle(ctx, cmd)
	}
}
