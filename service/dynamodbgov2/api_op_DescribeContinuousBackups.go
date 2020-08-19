// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodbgov2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbgov2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Checks the status of continuous backups and point in time recovery on the
// specified table. Continuous backups are ENABLED on all tables at table creation.
// If point in time recovery is enabled, PointInTimeRecoveryStatus will be set to
// ENABLED. After continuous backups and point in time recovery are enabled, you
// can restore to any point in time within EarliestRestorableDateTime and
// LatestRestorableDateTime. LatestRestorableDateTime is typically 5 minutes before
// the current time. You can restore your table to any point in time during the
// last 35 days. You can call DescribeContinuousBackups at a maximum rate of 10
// times per second.
func (c *Client) DescribeContinuousBackups(ctx context.Context, params *DescribeContinuousBackupsInput, optFns ...func(*Options)) (*DescribeContinuousBackupsOutput, error) {
	stack := middleware.NewStack("DescribeContinuousBackups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpDescribeContinuousBackupsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	v4.AddHTTPSignerMiddleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addOpDescribeContinuousBackupsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeContinuousBackups(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "DescribeContinuousBackups",
			Err:           err,
		}
	}
	out := result.(*DescribeContinuousBackupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeContinuousBackupsInput struct {
	// Name of the table for which the customer wants to check the continuous backups
	// and point in time recovery settings.
	TableName *string
}

type DescribeContinuousBackupsOutput struct {
	// Represents the continuous backups and point in time recovery settings on the
	// table.
	ContinuousBackupsDescription *types.ContinuousBackupsDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpDescribeContinuousBackupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeContinuousBackups{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeContinuousBackups{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeContinuousBackups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "DynamoDB GoV2",
		ServiceID:      "dynamodbgov2",
		EndpointPrefix: "dynamodbgov2",
		SigningName:    "dynamodb",
		OperationName:  "DescribeContinuousBackups",
	}
}
