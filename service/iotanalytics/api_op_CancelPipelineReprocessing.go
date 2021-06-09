// Code generated by smithy-go-codegen DO NOT EDIT.

package iotanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Cancels the reprocessing of data through the pipeline.
func (c *Client) CancelPipelineReprocessing(ctx context.Context, params *CancelPipelineReprocessingInput, optFns ...func(*Options)) (*CancelPipelineReprocessingOutput, error) {
	if params == nil {
		params = &CancelPipelineReprocessingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelPipelineReprocessing", params, optFns, c.addOperationCancelPipelineReprocessingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelPipelineReprocessingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelPipelineReprocessingInput struct {

	// The name of pipeline for which data reprocessing is canceled.
	//
	// This member is required.
	PipelineName *string

	// The ID of the reprocessing task (returned by StartPipelineReprocessing).
	//
	// This member is required.
	ReprocessingId *string
}

type CancelPipelineReprocessingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCancelPipelineReprocessingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCancelPipelineReprocessing{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCancelPipelineReprocessing{}, middleware.After)
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
	if err = addOpCancelPipelineReprocessingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelPipelineReprocessing(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelPipelineReprocessing(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotanalytics",
		OperationName: "CancelPipelineReprocessing",
	}
}
