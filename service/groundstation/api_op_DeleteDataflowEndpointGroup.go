// Code generated by smithy-go-codegen DO NOT EDIT.

package groundstation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a dataflow endpoint group.
func (c *Client) DeleteDataflowEndpointGroup(ctx context.Context, params *DeleteDataflowEndpointGroupInput, optFns ...func(*Options)) (*DeleteDataflowEndpointGroupOutput, error) {
	if params == nil {
		params = &DeleteDataflowEndpointGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDataflowEndpointGroup", params, optFns, c.addOperationDeleteDataflowEndpointGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDataflowEndpointGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DeleteDataflowEndpointGroupInput struct {

	// UUID of a dataflow endpoint group.
	//
	// This member is required.
	DataflowEndpointGroupId *string
}

//
type DeleteDataflowEndpointGroupOutput struct {

	// UUID of a dataflow endpoint group.
	DataflowEndpointGroupId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDeleteDataflowEndpointGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteDataflowEndpointGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteDataflowEndpointGroup{}, middleware.After)
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
	if err = addOpDeleteDataflowEndpointGroupValidationMiddleware(stack); err != nil {
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
