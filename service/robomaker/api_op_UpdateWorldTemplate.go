// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates a world template.
func (c *Client) UpdateWorldTemplate(ctx context.Context, params *UpdateWorldTemplateInput, optFns ...func(*Options)) (*UpdateWorldTemplateOutput, error) {
	if params == nil {
		params = &UpdateWorldTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateWorldTemplate", params, optFns, addOperationUpdateWorldTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateWorldTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateWorldTemplateInput struct {

	// The Amazon Resource Name (arn) of the world template to update.
	//
	// This member is required.
	Template *string

	// The name of the template.
	Name *string

	// The world template body.
	TemplateBody *string

	// The location of the world template.
	TemplateLocation *types.TemplateLocation
}

type UpdateWorldTemplateOutput struct {

	// The Amazon Resource Name (arn) of the world template.
	Arn *string

	// The time, in milliseconds since the epoch, when the world template was created.
	CreatedAt *time.Time

	// The time, in milliseconds since the epoch, when the world template was last
	// updated.
	LastUpdatedAt *time.Time

	// The name of the world template.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateWorldTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateWorldTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateWorldTemplate{}, middleware.After)
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
	if err = addOpUpdateWorldTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateWorldTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateWorldTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "UpdateWorldTemplate",
	}
}
