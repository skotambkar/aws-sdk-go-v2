// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/fsx/enums"
)

// The request object for DeleteFileSystem operation.
type DeleteFileSystemInput struct {
	_ struct{} `type:"structure"`

	// (Optional) A string of up to 64 ASCII characters that Amazon FSx uses to
	// ensure idempotent deletion. This is automatically filled on your behalf when
	// using the AWS CLI or SDK.
	ClientRequestToken *string `min:"1" type:"string" idempotencyToken:"true"`

	// The ID of the file system you want to delete.
	//
	// FileSystemId is a required field
	FileSystemId *string `min:"11" type:"string" required:"true"`

	// The configuration object for the Microsoft Windows file system used in the
	// DeleteFileSystem operation.
	WindowsConfiguration *DeleteFileSystemWindowsConfiguration `type:"structure"`
}

// String returns the string representation
func (s DeleteFileSystemInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteFileSystemInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteFileSystemInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}

	if s.FileSystemId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FileSystemId"))
	}
	if s.FileSystemId != nil && len(*s.FileSystemId) < 11 {
		invalidParams.Add(aws.NewErrParamMinLen("FileSystemId", 11))
	}
	if s.WindowsConfiguration != nil {
		if err := s.WindowsConfiguration.Validate(); err != nil {
			invalidParams.AddNested("WindowsConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The response object for the DeleteFileSystem operation.
type DeleteFileSystemOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the file system being deleted.
	FileSystemId *string `min:"11" type:"string"`

	// The file system lifecycle for the deletion request. Should be DELETING.
	Lifecycle enums.FileSystemLifecycle `type:"string" enum:"true"`

	// The response object for the Microsoft Windows file system used in the DeleteFileSystem
	// operation.
	WindowsResponse *DeleteFileSystemWindowsResponse `type:"structure"`
}

// String returns the string representation
func (s DeleteFileSystemOutput) String() string {
	return awsutil.Prettify(s)
}