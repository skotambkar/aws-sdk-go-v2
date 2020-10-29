package customizations

import (
	"context"
	"fmt"
	"strings"

	"github.com/awslabs/smithy-go/middleware"
	"github.com/awslabs/smithy-go/transport/http"

	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
)

// processOutpostIDMiddleware is special customization middleware to be applied for operations
// CreateBucket, ListRegionalBuckets which must resolve endpoint to s3-outposts.{region}.amazonaws.com
// with region as client region and signed by s3-control if an outpost id is provided.
type processOutpostIDMiddleware struct {
	// GetOutpostID points to a function that processes an input and returns an outpostID as string ptr,
	// and bool indicating if outpostID is supported or set.
	GetOutpostID func(interface{}) (*string, bool)
}

// ID returns the middleware ID.
func (*processOutpostIDMiddleware) ID() string { return "S3Control:ProcessARNResourceMiddleware" }

// HandleSerialize adds a serialize step, this has to be before operation serializer and arn endpoint logic.
// Ideally this step will be ahead of ARN customization for CreateBucket, ListRegionalBucket operation.
func (m *processOutpostIDMiddleware) HandleSerialize(
	ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler,
) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	// if host name is immutable, skip this customization
	if http.GetHostnameImmutable(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	outpostID, ok := m.GetOutpostID(in)
	if !ok {
		return next.HandleSerialize(ctx, in)
	}

	// check if outpostID was not set or is empty
	if outpostID == nil || len(strings.TrimSpace(*outpostID)) == 0 {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*http.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown request type %T", req)
	}

	serviceEndpointLabel := "s3-outposts."
	requestRegion := awsmiddleware.GetRegion(ctx)

	// set request url
	req.URL.Host = serviceEndpointLabel + requestRegion + ".amazonaws.com"

	// redirect signer
	ctx = awsmiddleware.SetSigningName(ctx, "s3-outposts")
	ctx = awsmiddleware.SetSigningRegion(ctx, requestRegion)

	return next.HandleSerialize(ctx, in)
}
