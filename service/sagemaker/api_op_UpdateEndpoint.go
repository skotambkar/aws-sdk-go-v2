// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deploys the new EndpointConfig specified in the request, switches to using newly
// created endpoint, and then deletes resources provisioned for the endpoint using
// the previous EndpointConfig (there is no availability loss). When Amazon
// SageMaker receives the request, it sets the endpoint status to Updating. After
// updating the endpoint, it sets the status to InService. To check the status of
// an endpoint, use the DescribeEndpoint API. You must not delete an EndpointConfig
// in use by an endpoint that is live or while the UpdateEndpoint or CreateEndpoint
// operations are being performed on the endpoint. To update an endpoint, you must
// create a new EndpointConfig. If you delete the EndpointConfig of an endpoint
// that is active or being created or updated you may lose visibility into the
// instance type the endpoint is using. The endpoint must be deleted in order to
// stop incurring charges.
func (c *Client) UpdateEndpoint(ctx context.Context, params *UpdateEndpointInput, optFns ...func(*Options)) (*UpdateEndpointOutput, error) {
	if params == nil {
		params = &UpdateEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEndpoint", params, optFns, c.addOperationUpdateEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEndpointInput struct {

	// The name of the new endpoint configuration.
	//
	// This member is required.
	EndpointConfigName *string

	// The name of the endpoint whose configuration you want to update.
	//
	// This member is required.
	EndpointName *string

	// The deployment configuration for the endpoint to be updated.
	DeploymentConfig *types.DeploymentConfig

	// When you are updating endpoint resources with
	// UpdateEndpointInput$RetainAllVariantProperties, whose value is set to true,
	// ExcludeRetainedVariantProperties specifies the list of type VariantProperty to
	// override with the values provided by EndpointConfig. If you don't specify a
	// value for ExcludeAllVariantProperties, no variant properties are overridden.
	ExcludeRetainedVariantProperties []types.VariantProperty

	// When updating endpoint resources, enables or disables the retention of variant
	// properties
	// (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_VariantProperty.html),
	// such as the instance count or the variant weight. To retain the variant
	// properties of an endpoint when updating it, set RetainAllVariantProperties to
	// true. To use the variant properties specified in a new EndpointConfig call when
	// updating an endpoint, set RetainAllVariantProperties to false. The default is
	// false.
	RetainAllVariantProperties bool
}

type UpdateEndpointOutput struct {

	// The Amazon Resource Name (ARN) of the endpoint.
	//
	// This member is required.
	EndpointArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationUpdateEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateEndpoint{}, middleware.After)
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
	if err = addOpUpdateEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "UpdateEndpoint",
	}
}
