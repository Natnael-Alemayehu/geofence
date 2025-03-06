package web

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const (
	writerKey = iota + 1
	TraceID
)

func setTraceID(ctx context.Context, traceIDkey uuid.UUID) context.Context {
	return context.WithValue(ctx, TraceID, traceIDkey)
}

// GetTraceID returns the traceID for the request.
func GetTraceID(ctx context.Context) http.ResponseWriter {
	v, ok := ctx.Value(TraceID).(http.ResponseWriter)
	if !ok {
		return nil
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
