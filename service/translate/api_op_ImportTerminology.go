// Code generated by smithy-go-codegen DO NOT EDIT.

package translate

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/translate/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates a custom terminology, depending on whether or not one already
// exists for the given terminology name. Importing a terminology with the same
// name as an existing one will merge the terminologies based on the chosen merge
// strategy. Currently, the only supported merge strategy is OVERWRITE, and so the
// imported terminology will overwrite an existing terminology of the same name. If
// you import a terminology that overwrites an existing one, the new terminology
// take up to 10 minutes to fully propagate and be available for use in a
// translation due to cache policies with the DataPlane service that performs the
// translations.
func (c *Client) ImportTerminology(ctx context.Context, params *ImportTerminologyInput, optFns ...func(*Options)) (*ImportTerminologyOutput, error) {
	if params == nil {
		params = &ImportTerminologyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportTerminology", params, optFns, addOperationImportTerminologyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportTerminologyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportTerminologyInput struct {

	// The merge strategy of the custom terminology being imported. Currently, only the
	// OVERWRITE merge strategy is supported. In this case, the imported terminology
	// will overwrite an existing terminology of the same name.
	//
	// This member is required.
	MergeStrategy types.MergeStrategy

	// The name of the custom terminology being imported.
	//
	// This member is required.
	Name *string

	// The terminology data for the custom terminology being imported.
	//
	// This member is required.
	TerminologyData *types.TerminologyData

	// The description of the custom terminology being imported.
	Description *string

	// The encryption key for the custom terminology being imported.
	EncryptionKey *types.EncryptionKey
}

type ImportTerminologyOutput struct {

	// The properties of the custom terminology being imported.
	TerminologyProperties *types.TerminologyProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationImportTerminologyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpImportTerminology{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpImportTerminology{}, middleware.After)
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
	if err = addOpImportTerminologyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportTerminology(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opImportTerminology(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "translate",
		OperationName: "ImportTerminology",
	}
}
