// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an alias
// (https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html) for a
// Lambda function version. Use aliases to provide clients with a function
// identifier that you can update to invoke a different version. You can also map
// an alias to split invocation requests between two versions. Use the
// RoutingConfig parameter to specify a second version and the percentage of
// invocation requests that it receives.
func (c *Client) CreateAlias(ctx context.Context, params *CreateAliasInput, optFns ...func(*Options)) (*CreateAliasOutput, error) {
	if params == nil {
		params = &CreateAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAlias", params, optFns, c.addOperationCreateAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAliasInput struct {

	// The name of the Lambda function. Name formats
	//
	// * Function name - MyFunction.
	//
	// *
	// Function ARN - arn:aws:lambda:us-west-2:123456789012:function:MyFunction.
	//
	// *
	// Partial ARN - 123456789012:function:MyFunction.
	//
	// The length constraint applies
	// only to the full ARN. If you specify only the function name, it is limited to 64
	// characters in length.
	//
	// This member is required.
	FunctionName *string

	// The function version that the alias invokes.
	//
	// This member is required.
	FunctionVersion *string

	// The name of the alias.
	//
	// This member is required.
	Name *string

	// A description of the alias.
	Description *string

	// The routing configuration
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-aliases.html#configuring-alias-routing)
	// of the alias.
	RoutingConfig *types.AliasRoutingConfiguration
}

// Provides configuration information about a Lambda function alias
// (https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html).
type CreateAliasOutput struct {

	// The Amazon Resource Name (ARN) of the alias.
	AliasArn *string

	// A description of the alias.
	Description *string

	// The function version that the alias invokes.
	FunctionVersion *string

	// The name of the alias.
	Name *string

	// A unique identifier that changes when you update the alias.
	RevisionId *string

	// The routing configuration
	// (https://docs.aws.amazon.com/lambda/latest/dg/lambda-traffic-shifting-using-aliases.html)
	// of the alias.
	RoutingConfig *types.AliasRoutingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAlias{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAlias{}, middleware.After)
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
	if err = addOpCreateAliasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAlias(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAlias(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "CreateAlias",
	}
}
