// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Saves a new query or updates an existing saved query. The QueryName must be
// unique for a single AWS account and a single AWS Region. You can create upto 300
// queries in a single AWS account and a single AWS Region.
func (c *Client) PutStoredQuery(ctx context.Context, params *PutStoredQueryInput, optFns ...func(*Options)) (*PutStoredQueryOutput, error) {
	if params == nil {
		params = &PutStoredQueryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutStoredQuery", params, optFns, c.addOperationPutStoredQueryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutStoredQueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutStoredQueryInput struct {

	// A list of StoredQuery objects. The mandatory fields are QueryName and
	// Expression. When you are creating a query, you must provide a query name and an
	// expression. When you are updating a query, you must provide a query name but
	// updating the description is optional.
	//
	// This member is required.
	StoredQuery *types.StoredQuery

	// A list of Tags object.
	Tags []types.Tag
}

type PutStoredQueryOutput struct {

	// Amazon Resource Name (ARN) of the query. For example,
	// arn:partition:service:region:account-id:resource-type/resource-name/resource-id.
	QueryArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationPutStoredQueryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutStoredQuery{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutStoredQuery{}, middleware.After)
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
	if err = addOpPutStoredQueryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutStoredQuery(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutStoredQuery(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "PutStoredQuery",
	}
}
