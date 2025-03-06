package web

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const (
	writerKey = iota + 1
	TraceIDKey
)

func setTraceID(ctx context.Context, traceIDkey uuid.UUID) context.Context {
	return context.WithValue(ctx, TraceIDKey, traceIDkey)
}

// GetTraceID returns the traceID for the request.
func GetTraceID(ctx context.Context) uuid.UUID {
	v, ok := ctx.Value(TraceIDKey).(uuid.UUID)
	if !ok {
		return uuid.Nil
	}

	return v
}

func setWriter(ctx context.Context, w http.ResponseWriter) context.Context {
	return context.WithValue(ctx, writerKey, w)
}

// GetWriter returns the underlying writer for the request.
func GetWriter(ctx context.Context) http.ResponseWriter {
	v, ok := ctx.Value(writerKey).(http.ResponseWriter)
	if !ok {
		return nil
	}

	return v
}
