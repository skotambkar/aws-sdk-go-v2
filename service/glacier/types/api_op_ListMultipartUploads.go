// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Provides options for retrieving list of in-progress multipart uploads for
// an Amazon Glacier vault.
type ListMultipartUploadsInput struct {
	_ struct{} `type:"structure"`

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen),
	// in which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// Specifies the maximum number of uploads returned in the response body. If
	// this value is not specified, the List Uploads operation returns up to 50
	// uploads.
	Limit *string `location:"querystring" locationName:"limit" type:"string"`

	// An opaque string used for pagination. This value specifies the upload at
	// which the listing of uploads should begin. Get the marker value from a previous
	// List Uploads response. You need only include the marker if you are continuing
	// the pagination of results started in a previous List Uploads request.
	Marker *string `location:"querystring" locationName:"marker" type:"string"`

	// The name of the vault.
	//
	// VaultName is a required field
	VaultName *string `location:"uri" locationName:"vaultName" type:"string" required:"true"`
}

// String returns the string representation
func (s ListMultipartUploadsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListMultipartUploadsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListMultipartUploadsInput"}

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

// Contains the Amazon S3 Glacier response to your request.
type ListMultipartUploadsOutput struct {
	_ struct{} `type:"structure"`

	// An opaque string that represents where to continue pagination of the results.
	// You use the marker in a new List Multipart Uploads request to obtain more
	// uploads in the list. If there are no more uploads, this value is null.
	Marker *string `type:"string"`

	// A list of in-progress multipart uploads.
	UploadsList []UploadListElement `type:"list"`
}

// String returns the string representation
func (s ListMultipartUploadsOutput) String() string {
	return awsutil.Prettify(s)
}