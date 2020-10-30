// Code generated by smithy-go-codegen DO NOT EDIT.

package timestreamwrite

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a given Timestream database. This is an irreversible operation. After a
// database is deleted, the time series data from its tables cannot be recovered.
// All tables in the database must be deleted first, or a ValidationException error
// will be thrown.
func (c *Client) DeleteDatabase(ctx context.Context, params *DeleteDatabaseInput, optFns ...func(*Options)) (*DeleteDatabaseOutput, error) {
	if params == nil {
		params = &DeleteDatabaseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDatabase", params, optFns, addOperationDeleteDatabaseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDatabaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDatabaseInput struct {

	// The name of the Timestream database to be deleted.
	//
	// This member is required.
	DatabaseName *string
}

type DeleteDatabaseOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteDatabaseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteDatabase{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteDatabase{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDeleteDatabaseValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDatabase(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteDatabase(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "timestream",
		OperationName: "DeleteDatabase",
	}
}