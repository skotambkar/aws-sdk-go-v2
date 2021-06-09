// Code generated by smithy-go-codegen DO NOT EDIT.

package iotanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the contents of a data set as presigned URIs.
func (c *Client) GetDatasetContent(ctx context.Context, params *GetDatasetContentInput, optFns ...func(*Options)) (*GetDatasetContentOutput, error) {
	if params == nil {
		params = &GetDatasetContentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDatasetContent", params, optFns, c.addOperationGetDatasetContentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDatasetContentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDatasetContentInput struct {

	// The name of the data set whose contents are retrieved.
	//
	// This member is required.
	DatasetName *string

	// The version of the data set whose contents are retrieved. You can also use the
	// strings "$LATEST" or "$LATEST_SUCCEEDED" to retrieve the contents of the latest
	// or latest successfully completed data set. If not specified, "$LATEST_SUCCEEDED"
	// is the default.
	VersionId *string
}

type GetDatasetContentOutput struct {

	// A list of DatasetEntry objects.
	Entries []types.DatasetEntry

	// The status of the data set content.
	Status *types.DatasetContentStatus

	// The time when the request was made.
	Timestamp *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationGetDatasetContentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDatasetContent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDatasetContent{}, middleware.After)
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
	if err = addOpGetDatasetContentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDatasetContent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDatasetContent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotanalytics",
		OperationName: "GetDatasetContent",
	}
}
