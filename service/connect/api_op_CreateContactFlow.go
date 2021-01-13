// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a contact flow for the specified Amazon Connect instance. You can also
// create and update contact flows using the Amazon Connect Flow language
// (https://docs.aws.amazon.com/connect/latest/adminguide/flow-language.html).
func (c *Client) CreateContactFlow(ctx context.Context, params *CreateContactFlowInput, optFns ...func(*Options)) (*CreateContactFlowOutput, error) {
	if params == nil {
		params = &CreateContactFlowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateContactFlow", params, optFns, addOperationCreateContactFlowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateContactFlowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateContactFlowInput struct {

	// The content of the contact flow.
	//
	// This member is required.
	Content *string

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// The name of the contact flow.
	//
	// This member is required.
	Name *string

	// The type of the contact flow. For descriptions of the available types, see
	// Choose a Contact Flow Type
	// (https://docs.aws.amazon.com/connect/latest/adminguide/create-contact-flow.html#contact-flow-types)
	// in the Amazon Connect Administrator Guide.
	//
	// This member is required.
	Type types.ContactFlowType

	// The description of the contact flow.
	Description *string

	// One or more tags.
	Tags map[string]string
}

type CreateContactFlowOutput struct {

	// The Amazon Resource Name (ARN) of the contact flow.
	ContactFlowArn *string

	// The identifier of the contact flow.
	ContactFlowId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateContactFlowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateContactFlow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateContactFlow{}, middleware.After)
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
	if err = addOpCreateContactFlowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateContactFlow(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateContactFlow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "CreateContactFlow",
	}
}
