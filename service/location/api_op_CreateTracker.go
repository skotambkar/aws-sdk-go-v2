// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a tracker resource in your AWS account, which lets you retrieve current
// and historical location of devices.
func (c *Client) CreateTracker(ctx context.Context, params *CreateTrackerInput, optFns ...func(*Options)) (*CreateTrackerOutput, error) {
	if params == nil {
		params = &CreateTrackerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTracker", params, optFns, addOperationCreateTrackerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTrackerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTrackerInput struct {

	// Specifies the pricing plan for your tracker resource. For additional details and
	// restrictions on each pricing plan option, see the Amazon Location Service
	// pricing page (https://aws.amazon.com/location/pricing/).
	//
	// This member is required.
	PricingPlan types.PricingPlan

	// The name for the tracker resource. Requirements:
	//
	// * Contain only alphanumeric
	// characters (A-Z, a-z, 0-9) , hyphens (-), periods (.), and underscores (_).
	//
	// *
	// Must be a unique tracker resource name.
	//
	// * No spaces allowed. For example,
	// ExampleTracker.
	//
	// This member is required.
	TrackerName *string

	// An optional description for the tracker resource.
	Description *string

	// Specifies the plan data source. Required if the Mobile Asset Tracking (MAT) or
	// the Mobile Asset Management (MAM) pricing plan is selected. Billing is
	// determined by the resource usage, the associated pricing plan, and data source
	// that was specified. For more information about each pricing plan option and
	// restrictions, see the Amazon Location Service pricing page
	// (https://aws.amazon.com/location/pricing/). Valid Values: Esri | Here
	PricingPlanDataSource *string
}

type CreateTrackerOutput struct {

	// The timestamp for when the tracker resource was created in  ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ.
	//
	// This member is required.
	CreateTime *time.Time

	// The Amazon Resource Name (ARN) for the tracker resource. Used when you need to
	// specify a resource across all AWS.
	//
	// This member is required.
	TrackerArn *string

	// The name of the tracker resource.
	//
	// This member is required.
	TrackerName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateTrackerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateTracker{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateTracker{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateTrackerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}
