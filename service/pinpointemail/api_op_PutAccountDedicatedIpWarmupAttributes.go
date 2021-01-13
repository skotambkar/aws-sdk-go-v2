// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enable or disable the automatic warm-up feature for dedicated IP addresses.
func (c *Client) PutAccountDedicatedIpWarmupAttributes(ctx context.Context, params *PutAccountDedicatedIpWarmupAttributesInput, optFns ...func(*Options)) (*PutAccountDedicatedIpWarmupAttributesOutput, error) {
	if params == nil {
		params = &PutAccountDedicatedIpWarmupAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAccountDedicatedIpWarmupAttributes", params, optFns, addOperationPutAccountDedicatedIpWarmupAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAccountDedicatedIpWarmupAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to enable or disable the automatic IP address warm-up feature.
type PutAccountDedicatedIpWarmupAttributesInput struct {

	// Enables or disables the automatic warm-up feature for dedicated IP addresses
	// that are associated with your Amazon Pinpoint account in the current AWS Region.
	// Set to true to enable the automatic warm-up feature, or set to false to disable
	// it.
	AutoWarmupEnabled bool
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type PutAccountDedicatedIpWarmupAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutAccountDedicatedIpWarmupAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutAccountDedicatedIpWarmupAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutAccountDedicatedIpWarmupAttributes{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutAccountDedicatedIpWarmupAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutAccountDedicatedIpWarmupAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "PutAccountDedicatedIpWarmupAttributes",
	}
}
