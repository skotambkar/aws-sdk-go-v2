// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the SAML provider resource objects defined in IAM in the account. IAM
// resource-listing operations return a subset of the available attributes for the
// resource. For example, this operation does not return tags, even though they are
// an attribute of the returned object. To view all of the information for a SAML
// provider, see GetSAMLProvider. This operation requires Signature Version 4
// (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).
func (c *Client) ListSAMLProviders(ctx context.Context, params *ListSAMLProvidersInput, optFns ...func(*Options)) (*ListSAMLProvidersOutput, error) {
	if params == nil {
		params = &ListSAMLProvidersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSAMLProviders", params, optFns, c.addOperationListSAMLProvidersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSAMLProvidersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSAMLProvidersInput struct {
}

// Contains the response to a successful ListSAMLProviders request.
type ListSAMLProvidersOutput struct {

	// The list of SAML provider resource objects defined in IAM for this AWS account.
	SAMLProviderList []types.SAMLProviderListEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListSAMLProvidersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListSAMLProviders{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListSAMLProviders{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSAMLProviders(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListSAMLProviders(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "ListSAMLProviders",
	}
}
