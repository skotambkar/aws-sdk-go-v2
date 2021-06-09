// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudsearch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about the search domains owned by this account. Can be limited
// to specific domains. Shows all domains by default. To get the number of
// searchable documents in a domain, use the console or submit a matchall request
// to your domain's search endpoint: q=matchall&q.parser=structured&size=0. For
// more information, see Getting Information about a Search Domain
// (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/getting-domain-info.html)
// in the Amazon CloudSearch Developer Guide.
func (c *Client) DescribeDomains(ctx context.Context, params *DescribeDomainsInput, optFns ...func(*Options)) (*DescribeDomainsOutput, error) {
	if params == nil {
		params = &DescribeDomainsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDomains", params, optFns, c.addOperationDescribeDomainsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDomainsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DescribeDomains operation. By default shows
// the status of all domains. To restrict the response to particular domains,
// specify the names of the domains you want to describe.
type DescribeDomainsInput struct {

	// The names of the domains you want to include in the response.
	DomainNames []string
}

// The result of a DescribeDomains request. Contains the status of the domains
// specified in the request or all domains owned by the account.
type DescribeDomainsOutput struct {

	// A list that contains the status of each requested domain.
	//
	// This member is required.
	DomainStatusList []types.DomainStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeDomainsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDomains{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDomains{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDomains(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDomains(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudsearch",
		OperationName: "DescribeDomains",
	}
}
