// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Upgrades a published schema under a new minor version revision using the current
// contents of DevelopmentSchemaArn.
func (c *Client) UpgradePublishedSchema(ctx context.Context, params *UpgradePublishedSchemaInput, optFns ...func(*Options)) (*UpgradePublishedSchemaOutput, error) {
	if params == nil {
		params = &UpgradePublishedSchemaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpgradePublishedSchema", params, optFns, addOperationUpgradePublishedSchemaMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpgradePublishedSchemaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpgradePublishedSchemaInput struct {

	// The ARN of the development schema with the changes used for the upgrade.
	//
	// This member is required.
	DevelopmentSchemaArn *string

	// Identifies the minor version of the published schema that will be created. This
	// parameter is NOT optional.
	//
	// This member is required.
	MinorVersion *string

	// The ARN of the published schema to be upgraded.
	//
	// This member is required.
	PublishedSchemaArn *string

	// Used for testing whether the Development schema provided is backwards
	// compatible, or not, with the publish schema provided by the user to be upgraded.
	// If schema compatibility fails, an exception would be thrown else the call would
	// succeed. This parameter is optional and defaults to false.
	DryRun bool
}

type UpgradePublishedSchemaOutput struct {

	// The ARN of the upgraded schema that is returned as part of the response.
	UpgradedSchemaArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpgradePublishedSchemaMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpgradePublishedSchema{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpgradePublishedSchema{}, middleware.After)
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
	if err = addOpUpgradePublishedSchemaValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpgradePublishedSchema(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpgradePublishedSchema(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "UpgradePublishedSchema",
	}
}
