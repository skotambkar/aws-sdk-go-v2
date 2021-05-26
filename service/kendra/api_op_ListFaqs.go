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

// Gets a list of FAQ lists associated with an index.
func (c *Client) ListFaqs(ctx context.Context, params *ListFaqsInput, optFns ...func(*Options)) (*ListFaqsOutput, error) {
	if params == nil {
		params = &ListFaqsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFaqs", params, optFns, c.addOperationListFaqsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFaqsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFaqsInput struct {

	// The index that contains the FAQ lists.
	//
	// This member is required.
	IndexId *string

	// The maximum number of FAQs to return in the response. If there are fewer results
	// in the list, this response contains only the actual results.
	MaxResults *int32

	// If the result of the previous request to ListFaqs was truncated, include the
	// NextToken to fetch the next set of FAQs.
	NextToken *string
}

type ListFaqsOutput struct {

	// information about the FAQs associated with the specified index.
	FaqSummaryItems []types.FaqSummary

	// The ListFaqs operation returns a page of FAQs at a time. The maximum size of the
	// page is set by the MaxResults parameter. If there are more jobs in the list than
	// the page size, Amazon Kendra returns the NextPage token. Include the token in
	// the next request to the ListFaqs operation to return the next page of FAQs.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListFaqsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListFaqs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListFaqs{}, middleware.After)
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
	if err = addOpListFaqsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFaqs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListFaqs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "ListFaqs",
	}
}
