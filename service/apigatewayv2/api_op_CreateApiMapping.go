// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an API mapping.
func (c *Client) CreateApiMapping(ctx context.Context, params *CreateApiMappingInput, optFns ...func(*Options)) (*CreateApiMappingOutput, error) {
	if params == nil {
		params = &CreateApiMappingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateApiMapping", params, optFns, c.addOperationCreateApiMappingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateApiMappingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates a new ApiMapping resource to represent an API mapping.
type CreateApiMappingInput struct {

	// The API identifier.
	//
	// This member is required.
	ApiId *string

	// The domain name.
	//
	// This member is required.
	DomainName *string

	// The API stage.
	//
	// This member is required.
	Stage *string

	// The API mapping key.
	ApiMappingKey *string
}

type CreateApiMappingOutput struct {

	// The API identifier.
	ApiId *string

	// The API mapping identifier.
	ApiMappingId *string

	// The API mapping key.
	ApiMappingKey *string

	// The API stage.
	Stage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateApiMappingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateApiMapping{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateApiMapping{}, middleware.After)
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
	if err = addOpCreateApiMappingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateApiMapping(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateApiMapping(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "CreateApiMapping",
	}
}
