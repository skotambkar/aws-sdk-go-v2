package s3shared

import (
	"context"
	"fmt"

	"github.com/awslabs/smithy-go/middleware"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
)

// InitARNLookupMiddleware is the initial middleware that looks up if an arn is provided.
// This middleware is responsible for fetching ARN from a arnable field, and registering the ARN on
// middleware context. This middleware must be executed before input validation step or any other
// arn processing middleware.
type InitARNLookupMiddleware struct {

	// GetARNValue takes in a input interface and returns a ptr to string and a bool
	GetARNValue func(interface{}) (*string, bool)
}

func (m *InitARNLookupMiddleware) ID() string {
	return "S3Shared:InitARNLookupMiddleware"
}

func (m *InitARNLookupMiddleware) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	// check is input resource is an ARN; if not go to next
	v, ok := m.GetARNValue(in.Parameters)
	if !ok || v == nil || !arn.IsARN(*v) {
		return next.HandleInitialize(ctx, in)
	}

	// if ARN process ResourceRequest and put it on ctx
	av, err := arn.Parse(*v)
	if err != nil {
		return out, metadata, fmt.Errorf("error parsing arn: %w", err)
	}
	// set parsed arn on context
	ctx = setARNResourceOnContext(ctx, av)

	return next.HandleInitialize(ctx, in)
}

type arnResourceValue struct {
}

// SetARNResourceOnContext sets arn on context
func setARNResourceOnContext(ctx context.Context, value arn.ARN) context.Context {
	return context.WithValue(ctx, arnResourceValue{}, value)
}

// GetARNResourceFromContext returns an ARN from context and a bool indicating presence of ARN on ctx
func GetARNResourceFromContext(ctx context.Context) (arn.ARN, bool) {
	v, ok := ctx.Value(arnResourceValue{}).(arn.ARN)
	return v, ok
}
