// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutvision

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lookoutvision/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new version of a model within an an Amazon Lookout for Vision project.
// CreateModel is an asynchronous operation in which Amazon Lookout for Vision
// trains, tests, and evaluates a new version of a model. To get the current
// status, check the Status field returned in the response from DescribeModel. If
// the project has a single dataset, Amazon Lookout for Vision internally splits
// the dataset to create a training and a test dataset. If the project has a
// training and a test dataset, Lookout for Vision uses the respective datasets to
// train and test the model. After training completes, the evaluation metrics are
// stored at the location specified in OutputConfig. This operation requires
// permissions to perform the lookoutvision:CreateModel operation. If you want to
// tag your model, you also require permission to the lookoutvision:TagResource
// operation.
func (c *Client) CreateModel(ctx context.Context, params *CreateModelInput, optFns ...func(*Options)) (*CreateModelOutput, error) {
	if params == nil {
		params = &CreateModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateModel", params, optFns, c.addOperationCreateModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateModelInput struct {

	// The location where Amazon Lookout for Vision saves the training results.
	//
	// This member is required.
	OutputConfig *types.OutputConfig

	// The name of the project in which you want to create a model version.
	//
	// This member is required.
	ProjectName *string

	// ClientToken is an idempotency token that ensures a call to CreateModel completes
	// only once. You choose the value to pass. For example, An issue, such as an
	// network outage, might prevent you from getting a response from CreateModel. In
	// this case, safely retry your call to CreateModel by using the same ClientToken
	// parameter value. An error occurs if the other input parameters are not the same
	// as in the first request. Using a different value for ClientToken is considered a
	// new call to CreateModel. An idempotency token is active for 8 hours.
	ClientToken *string

	// A description for the version of the model.
	Description *string

	// The identifier for your AWS Key Management Service (AWS KMS) customer master key
	// (CMK). The key is used to encrypt training and test images copied into the
	// service for model training. Your source images are unaffected. If this parameter
	// is not specified, the copied images are encrypted by a key that AWS owns and
	// manages.
	KmsKeyId *string

	// A set of tags (key-value pairs) that you want to attach to the model.
	Tags []types.Tag
}

type CreateModelOutput struct {

	// The response from a call to CreateModel.
	ModelMetadata *types.ModelMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateModel{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateModelMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateModel(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateModel struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateModel) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateModel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateModelInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateModelInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateModelMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateModel{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutvision",
		OperationName: "CreateModel",
	}
}
