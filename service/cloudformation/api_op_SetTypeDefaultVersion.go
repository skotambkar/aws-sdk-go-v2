// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Specify the default version of a type. The default version of a type will be
// used in CloudFormation operations.
func (c *Client) SetTypeDefaultVersion(ctx context.Context, params *SetTypeDefaultVersionInput, optFns ...func(*Options)) (*SetTypeDefaultVersionOutput, error) {
	if params == nil {
		params = &SetTypeDefaultVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetTypeDefaultVersion", params, optFns, addOperationSetTypeDefaultVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetTypeDefaultVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetTypeDefaultVersionInput struct {

	// The Amazon Resource Name (ARN) of the type for which you want version summary
	// information. Conditional: You must specify either TypeName and Type, or Arn.
	Arn *string

	// The kind of type. Conditional: You must specify either TypeName and Type, or
	// Arn.
	Type types.RegistryType

	// The name of the type. Conditional: You must specify either TypeName and Type, or
	// Arn.
	TypeName *string

	// The ID of a specific version of the type. The version ID is the value at the end
	// of the Amazon Resource Name (ARN) assigned to the type version when it is
	// registered.
	VersionId *string
}

type SetTypeDefaultVersionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSetTypeDefaultVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetTypeDefaultVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetTypeDefaultVersion{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetTypeDefaultVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetTypeDefaultVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "SetTypeDefaultVersion",
	}
}
