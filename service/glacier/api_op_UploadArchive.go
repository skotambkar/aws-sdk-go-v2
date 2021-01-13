// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
)

// This operation adds an archive to a vault. This is a synchronous operation, and
// for a successful upload, your data is durably persisted. Amazon S3 Glacier
// returns the archive ID in the x-amz-archive-id header of the response. You must
// use the archive ID to access your data in Amazon S3 Glacier. After you upload an
// archive, you should save the archive ID returned so that you can retrieve or
// delete the archive later. Besides saving the archive ID, you can also index it
// and give it a friendly name to allow for better searching. You can also use the
// optional archive description field to specify how the archive is referred to in
// an external index of archives, such as you might create in Amazon DynamoDB. You
// can also get the vault inventory to obtain a list of archive IDs in a vault. For
// more information, see InitiateJob. You must provide a SHA256 tree hash of the
// data you are uploading. For information about computing a SHA256 tree hash, see
// Computing Checksums
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/checksum-calculations.html).
// You can optionally specify an archive description of up to 1,024 printable ASCII
// characters. You can get the archive description when you either retrieve the
// archive or get the vault inventory. For more information, see InitiateJob.
// Amazon Glacier does not interpret the description in any way. An archive
// description does not need to be unique. You cannot use the description to
// retrieve or sort the archive list. Archives are immutable. After you upload an
// archive, you cannot edit the archive or its description. An AWS account has full
// permission to perform all operations (actions). However, AWS Identity and Access
// Management (IAM) users don't have any permissions by default. You must grant
// them explicit permission to perform specific actions. For more information, see
// Access Control Using AWS Identity and Access Management (IAM)
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html).
// For conceptual information and underlying REST API, see Uploading an Archive in
// Amazon Glacier
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/uploading-an-archive.html)
// and Upload Archive
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-archive-post.html) in
// the Amazon Glacier Developer Guide.
func (c *Client) UploadArchive(ctx context.Context, params *UploadArchiveInput, optFns ...func(*Options)) (*UploadArchiveOutput, error) {
	if params == nil {
		params = &UploadArchiveInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UploadArchive", params, optFns, addOperationUploadArchiveMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UploadArchiveOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides options to add an archive to a vault.
type UploadArchiveInput struct {

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen), in
	// which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// The name of the vault.
	//
	// This member is required.
	VaultName *string

	// The optional description of the archive you are uploading.
	ArchiveDescription *string

	// The data to upload.
	Body io.Reader

	// The SHA256 tree hash of the data being uploaded.
	Checksum *string
}

// Contains the Amazon S3 Glacier response to your request. For information about
// the underlying REST API, see Upload Archive
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-archive-post.html).
// For conceptual information, see Working with Archives in Amazon S3 Glacier
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/working-with-archives.html).
type UploadArchiveOutput struct {

	// The ID of the archive. This value is also included as part of the location.
	ArchiveId *string

	// The checksum of the archive computed by Amazon S3 Glacier.
	Checksum *string

	// The relative URI path of the newly added archive resource.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUploadArchiveMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUploadArchive{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUploadArchive{}, middleware.After)
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
	if err = addOpUploadArchiveValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUploadArchive(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddTreeHashMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion); err != nil {
		return err
	}
	if err = glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUploadArchive(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "UploadArchive",
	}
}
