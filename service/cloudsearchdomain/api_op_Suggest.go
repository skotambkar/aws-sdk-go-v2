// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudsearchdomain

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearchdomain/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves autocomplete suggestions for a partial query string. You can use
// suggestions enable you to display likely matches before users finish typing. In
// Amazon CloudSearch, suggestions are based on the contents of a particular text
// field. When you request suggestions, Amazon CloudSearch finds all of the
// documents whose values in the suggester field start with the specified query
// string. The beginning of the field must match the query string to be considered
// a match. For more information about configuring suggesters and retrieving
// suggestions, see Getting Suggestions
// (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/getting-suggestions.html)
// in the Amazon CloudSearch Developer Guide. The endpoint for submitting Suggest
// requests is domain-specific. You submit suggest requests to a domain's search
// endpoint. To get the search endpoint for your domain, use the Amazon CloudSearch
// configuration service DescribeDomains action. A domain's endpoints are also
// displayed on the domain dashboard in the Amazon CloudSearch console.
func (c *Client) Suggest(ctx context.Context, params *SuggestInput, optFns ...func(*Options)) (*SuggestOutput, error) {
	if params == nil {
		params = &SuggestInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "Suggest", params, optFns, c.addOperationSuggestMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SuggestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the Suggest request.
type SuggestInput struct {

	// Specifies the string for which you want to get suggestions.
	//
	// This member is required.
	Query *string

	// Specifies the name of the suggester to use to find suggested matches.
	//
	// This member is required.
	Suggester *string

	// Specifies the maximum number of suggestions to return.
	Size int64
}

// Contains the response to a Suggest request.
type SuggestOutput struct {

	// The status of a SuggestRequest. Contains the resource ID (rid) and how long it
	// took to process the request (timems).
	Status *types.SuggestStatus

	// Container for the matching search suggestion information.
	Suggest *types.SuggestModel

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationSuggestMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSuggest{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSuggest{}, middleware.After)
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
	if err = addOpSuggestValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSuggest(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSuggest(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudsearch",
		OperationName: "Suggest",
	}
}
