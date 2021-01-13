// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Determines the dominant language of the input text for a batch of documents. For
// a list of languages that Amazon Comprehend can detect, see Amazon Comprehend
// Supported Languages
// (https://docs.aws.amazon.com/comprehend/latest/dg/how-languages.html).
func (c *Client) BatchDetectDominantLanguage(ctx context.Context, params *BatchDetectDominantLanguageInput, optFns ...func(*Options)) (*BatchDetectDominantLanguageOutput, error) {
	if params == nil {
		params = &BatchDetectDominantLanguageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDetectDominantLanguage", params, optFns, addOperationBatchDetectDominantLanguageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDetectDominantLanguageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDetectDominantLanguageInput struct {

	// A list containing the text of the input documents. The list can contain a
	// maximum of 25 documents. Each document should contain at least 20 characters and
	// must contain fewer than 5,000 bytes of UTF-8 encoded characters.
	//
	// This member is required.
	TextList []string
}

type BatchDetectDominantLanguageOutput struct {

	// A list containing one object for each document that contained an error. The
	// results are sorted in ascending order by the Index field and match the order of
	// the documents in the input list. If there are no errors in the batch, the
	// ErrorList is empty.
	//
	// This member is required.
	ErrorList []types.BatchItemError

	// A list of objects containing the results of the operation. The results are
	// sorted in ascending order by the Index field and match the order of the
	// documents in the input list. If all of the documents contain an error, the
	// ResultList is empty.
	//
	// This member is required.
	ResultList []types.BatchDetectDominantLanguageItemResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchDetectDominantLanguageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchDetectDominantLanguage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchDetectDominantLanguage{}, middleware.After)
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
	if err = addOpBatchDetectDominantLanguageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDetectDominantLanguage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchDetectDominantLanguage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "BatchDetectDominantLanguage",
	}
}
