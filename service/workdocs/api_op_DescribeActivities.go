// Code generated by smithy-go-codegen DO NOT EDIT.

package workdocs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes the user activities in a specified time period.
func (c *Client) DescribeActivities(ctx context.Context, params *DescribeActivitiesInput, optFns ...func(*Options)) (*DescribeActivitiesOutput, error) {
	if params == nil {
		params = &DescribeActivitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeActivities", params, optFns, addOperationDescribeActivitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeActivitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeActivitiesInput struct {

	// Specifies which activity types to include in the response. If this field is left
	// empty, all activity types are returned.
	ActivityTypes *string

	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string

	// The timestamp that determines the end time of the activities. The response
	// includes the activities performed before the specified timestamp.
	EndTime *time.Time

	// Includes indirect activities. An indirect activity results from a direct
	// activity performed on a parent resource. For example, sharing a parent folder
	// (the direct activity) shares all of the subfolders and documents within the
	// parent folder (the indirect activity).
	IncludeIndirectActivities bool

	// The maximum number of items to return.
	Limit *int32

	// The marker for the next set of results.
	Marker *string

	// The ID of the organization. This is a mandatory parameter when using
	// administrative API (SigV4) requests.
	OrganizationId *string

	// The document or folder ID for which to describe activity types.
	ResourceId *string

	// The timestamp that determines the starting time of the activities. The response
	// includes the activities performed after the specified timestamp.
	StartTime *time.Time

	// The ID of the user who performed the action. The response includes activities
	// pertaining to this user. This is an optional parameter and is only applicable
	// for administrative API (SigV4) requests.
	UserId *string
}

type DescribeActivitiesOutput struct {

	// The marker for the next set of results.
	Marker *string

	// The list of activities for the specified user and time period.
	UserActivities []types.Activity

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeActivitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeActivities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeActivities{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeActivities(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeActivities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workdocs",
		OperationName: "DescribeActivities",
	}
}
