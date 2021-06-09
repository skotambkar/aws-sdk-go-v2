// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentity

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the IdentityID associated with a DeveloperUserIdentifier or the list
// of DeveloperUserIdentifier values associated with an IdentityId for an existing
// identity. Either IdentityID or DeveloperUserIdentifier must not be null. If you
// supply only one of these values, the other value will be searched in the
// database and returned as a part of the response. If you supply both,
// DeveloperUserIdentifier will be matched against IdentityID. If the values are
// verified against the database, the response returns both values and is the same
// as the request. Otherwise a ResourceConflictException is thrown.
// LookupDeveloperIdentity is intended for low-throughput control plane operations:
// for example, to enable customer service to locate an identity ID by username. If
// you are using it for higher-volume operations such as user authentication, your
// requests are likely to be throttled. GetOpenIdTokenForDeveloperIdentity is a
// better option for higher-volume operations for user authentication. You must use
// AWS Developer credentials to call this API.
func (c *Client) LookupDeveloperIdentity(ctx context.Context, params *LookupDeveloperIdentityInput, optFns ...func(*Options)) (*LookupDeveloperIdentityOutput, error) {
	if params == nil {
		params = &LookupDeveloperIdentityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "LookupDeveloperIdentity", params, optFns, c.addOperationLookupDeveloperIdentityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*LookupDeveloperIdentityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to the LookupDeveloperIdentityInput action.
type LookupDeveloperIdentityInput struct {

	// An identity pool ID in the format REGION:GUID.
	//
	// This member is required.
	IdentityPoolId *string

	// A unique ID used by your backend authentication process to identify a user.
	// Typically, a developer identity provider would issue many developer user
	// identifiers, in keeping with the number of users.
	DeveloperUserIdentifier *string

	// A unique identifier in the format REGION:GUID.
	IdentityId *string

	// The maximum number of identities to return.
	MaxResults int32

	// A pagination token. The first call you make will have NextToken set to null.
	// After that the service will return NextToken values as needed. For example,
	// let's say you make a request with MaxResults set to 10, and there are 20 matches
	// in the database. The service will return a pagination token as a part of the
	// response. This token can be used to call the API again and get results starting
	// from the 11th match.
	NextToken *string
}

// Returned in response to a successful LookupDeveloperIdentity action.
type LookupDeveloperIdentityOutput struct {

	// This is the list of developer user identifiers associated with an identity ID.
	// Cognito supports the association of multiple developer user identifiers with an
	// identity ID.
	DeveloperUserIdentifierList []string

	// A unique identifier in the format REGION:GUID.
	IdentityId *string

	// A pagination token. The first call you make will have NextToken set to null.
	// After that the service will return NextToken values as needed. For example,
	// let's say you make a request with MaxResults set to 10, and there are 20 matches
	// in the database. The service will return a pagination token as a part of the
	// response. This token can be used to call the API again and get results starting
	// from the 11th match.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationLookupDeveloperIdentityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpLookupDeveloperIdentity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpLookupDeveloperIdentity{}, middleware.After)
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
	if err = addOpLookupDeveloperIdentityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opLookupDeveloperIdentity(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opLookupDeveloperIdentity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-identity",
		OperationName: "LookupDeveloperIdentity",
	}
}
