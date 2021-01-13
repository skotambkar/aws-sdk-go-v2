// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds an existing external connection to a repository. One external connection is
// allowed per repository. A repository can have one or more upstream repositories,
// or an external connection.
func (c *Client) AssociateExternalConnection(ctx context.Context, params *AssociateExternalConnectionInput, optFns ...func(*Options)) (*AssociateExternalConnectionOutput, error) {
	if params == nil {
		params = &AssociateExternalConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateExternalConnection", params, optFns, addOperationAssociateExternalConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateExternalConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateExternalConnectionInput struct {

	// The name of the domain that contains the repository.
	//
	// This member is required.
	Domain *string

	// The name of the external connection to add to the repository. The following
	// values are supported:
	//
	// * public:npmjs - for the npm public repository.
	//
	// *
	// public:pypi - for the Python Package Index.
	//
	// * public:maven-central - for Maven
	// Central.
	//
	// * public:maven-googleandroid - for the Google Android repository.
	//
	// *
	// public:maven-gradleplugins - for the Gradle plugins repository.
	//
	// *
	// public:maven-commonsware - for the CommonsWare Android repository.
	//
	// *
	// public:nuget-org - for the NuGet Gallery.
	//
	// This member is required.
	ExternalConnection *string

	// The name of the repository to which the external connection is added.
	//
	// This member is required.
	Repository *string

	// The 12-digit account number of the AWS account that owns the domain. It does not
	// include dashes or spaces.
	DomainOwner *string
}

type AssociateExternalConnectionOutput struct {

	// Information about the connected repository after processing the request.
	Repository *types.RepositoryDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAssociateExternalConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateExternalConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateExternalConnection{}, middleware.After)
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
	if err = addOpAssociateExternalConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateExternalConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateExternalConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "AssociateExternalConnection",
	}
}
