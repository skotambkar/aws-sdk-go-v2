// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Provides options for initiating a multipart upload to an Amazon S3 Glacier
// vault.
type InitiateMultipartUploadInput struct {
	_ struct{} `type:"structure"`

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen),
	// in which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The archive description that you are uploading in parts.
	//
	// The part size must be a megabyte (1024 KB) multiplied by a power of 2, for
	// example 1048576 (1 MB), 2097152 (2 MB), 4194304 (4 MB), 8388608 (8 MB), and
	// so on. The minimum allowable part size is 1 MB, and the maximum is 4 GB (4096
	// MB).
	ArchiveDescription *string `location:"header" locationName:"x-amz-archive-description" type:"string"`

	// The size of each part except the last, in bytes. The last part can be smaller
	// than this part size.
	PartSize *string `location:"header" locationName:"x-amz-part-size" type:"string"`

	// The name of the vault.
	//
	// VaultName is a required field
	VaultName *string `location:"uri" locationName:"vaultName" type:"string" required:"true"`
}

// String returns the string representation
func (s InitiateMultipartUploadInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *InitiateMultipartUploadInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "InitiateMultipartUploadInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.VaultName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VaultName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The Amazon S3 Glacier response to your request.
type InitiateMultipartUploadOutput struct {
	_ struct{} `type:"structure"`

	// The relative URI path of the multipart upload ID Amazon S3 Glacier created.
	Location *string `location:"header" locationName:"Location" type:"string"`

	// The ID of the multipart upload. This value is also included as part of the
	// location.
	UploadId *string `location:"header" locationName:"x-amz-multipart-upload-id" type:"string"`
}

// String returns the string representation
func (s InitiateMultipartUploadOutput) String() string {
	return awsutil.Prettify(s)
}
