// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing Amazon Kendra data source.
func (c *Client) UpdateDataSource(ctx context.Context, params *UpdateDataSourceInput, optFns ...func(*Options)) (*UpdateDataSourceOutput, error) {
	if params == nil {
		params = &UpdateDataSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDataSource", params, optFns, addOperationUpdateDataSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDataSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDataSourceInput struct {

	// The unique identifier of the data source to update.
	//
	// This member is required.
	Id *string

	// The identifier of the index that contains the data source to update.
	//
	// This member is required.
	IndexId *string

	// Configuration information for a Amazon Kendra data source.
	Configuration *types.DataSourceConfiguration

	// The new description for the data source.
	Description *string

	// The name of the data source to update. The name of the data source can't be
	// updated. To rename a data source you must delete the data source and re-create
	// it.
	Name *string

	// The Amazon Resource Name (ARN) of the new role to use when the data source is
	// accessing resources on your behalf.
	RoleArn *string

	// The new update schedule for the data source.
	Schedule *string
}

type UpdateDataSourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateDataSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDataSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDataSource{}, middleware.After)
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
	if err = addOpUpdateDataSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDataSource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDataSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "UpdateDataSource",
	}
}
