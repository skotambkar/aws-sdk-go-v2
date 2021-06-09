// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a quota increase request to your quota request template.
func (c *Client) PutServiceQuotaIncreaseRequestIntoTemplate(ctx context.Context, params *PutServiceQuotaIncreaseRequestIntoTemplateInput, optFns ...func(*Options)) (*PutServiceQuotaIncreaseRequestIntoTemplateOutput, error) {
	if params == nil {
		params = &PutServiceQuotaIncreaseRequestIntoTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutServiceQuotaIncreaseRequestIntoTemplate", params, optFns, c.addOperationPutServiceQuotaIncreaseRequestIntoTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutServiceQuotaIncreaseRequestIntoTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutServiceQuotaIncreaseRequestIntoTemplateInput struct {

	// The AWS Region.
	//
	// This member is required.
	AwsRegion *string

	// The new, increased value for the quota.
	//
	// This member is required.
	DesiredValue *float64

	// The quota identifier.
	//
	// This member is required.
	QuotaCode *string

	// The service identifier.
	//
	// This member is required.
	ServiceCode *string
}

type PutServiceQuotaIncreaseRequestIntoTemplateOutput struct {

	// Information about the quota increase request.
	ServiceQuotaIncreaseRequestInTemplate *types.ServiceQuotaIncreaseRequestInTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationPutServiceQuotaIncreaseRequestIntoTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutServiceQuotaIncreaseRequestIntoTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutServiceQuotaIncreaseRequestIntoTemplate{}, middleware.After)
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
	if err = addOpPutServiceQuotaIncreaseRequestIntoTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutServiceQuotaIncreaseRequestIntoTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutServiceQuotaIncreaseRequestIntoTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "PutServiceQuotaIncreaseRequestIntoTemplate",
	}
}
