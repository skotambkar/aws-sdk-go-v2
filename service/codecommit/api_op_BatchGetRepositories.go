// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about one or more repositories. The description field for a
// repository accepts all HTML characters and all valid Unicode characters.
// Applications that do not HTML-encode the description and display it in a webpage
// can expose users to potentially malicious code. Make sure that you HTML-encode
// the description field in any application that uses this API to display the
// repository description on a webpage.
func (c *Client) BatchGetRepositories(ctx context.Context, params *BatchGetRepositoriesInput, optFns ...func(*Options)) (*BatchGetRepositoriesOutput, error) {
	if params == nil {
		params = &BatchGetRepositoriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetRepositories", params, optFns, c.addOperationBatchGetRepositoriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetRepositoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a batch get repositories operation.
type BatchGetRepositoriesInput struct {

	// The names of the repositories to get information about. The length constraint
	// limit is for each string in the array. The array itself can be empty.
	//
	// This member is required.
	RepositoryNames []string
}

// Represents the output of a batch get repositories operation.
type BatchGetRepositoriesOutput struct {

	// A list of repositories returned by the batch get repositories operation.
	Repositories []types.RepositoryMetadata

	// Returns a list of repository names for which information could not be found.
	RepositoriesNotFound []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationBatchGetRepositoriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetRepositories{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetRepositories{}, middleware.After)
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
	if err = addOpBatchGetRepositoriesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetRepositories(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetRepositories(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "BatchGetRepositories",
	}
}
