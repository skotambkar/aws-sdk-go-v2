// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// RefreshCacheInput
type RefreshCacheInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the file share you want to refresh.
	//
	// FileShareARN is a required field
	FileShareARN *string `min:"50" type:"string" required:"true"`

	// A comma-separated list of the paths of folders to refresh in the cache. The
	// default is ["/"]. The default refreshes objects and folders at the root of
	// the Amazon S3 bucket. If Recursive is set to "true", the entire S3 bucket
	// that the file share has access to is refreshed.
	FolderList []string `min:"1" type:"list"`

	// A value that specifies whether to recursively refresh folders in the cache.
	// The refresh includes folders that were in the cache the last time the gateway
	// listed the folder's contents. If this value set to "true", each folder that
	// is listed in FolderList is recursively updated. Otherwise, subfolders listed
	// in FolderList are not refreshed. Only objects that are in folders listed
	// directly under FolderList are found and used for the update. The default
	// is "true".
	Recursive *bool `type:"boolean"`
}

// String returns the string representation
func (s RefreshCacheInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RefreshCacheInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RefreshCacheInput"}

	if s.FileShareARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("FileShareARN"))
	}
	if s.FileShareARN != nil && len(*s.FileShareARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("FileShareARN", 50))
	}
	if s.FolderList != nil && len(s.FolderList) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FolderList", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// RefreshCacheOutput
type RefreshCacheOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the file share.
	FileShareARN *string `min:"50" type:"string"`

	// The randomly generated ID of the notification that was sent. This ID is in
	// UUID format.
	NotificationId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s RefreshCacheOutput) String() string {
	return awsutil.Prettify(s)
}