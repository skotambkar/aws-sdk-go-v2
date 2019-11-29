// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request structure for create a new deployment.
type CreateDeploymentInput struct {
	_ struct{} `type:"structure"`

	// Unique Id for an Amplify App.
	//
	// AppId is a required field
	AppId *string `location:"uri" locationName:"appId" min:"1" type:"string" required:"true"`

	// Name for the branch, for the Job.
	//
	// BranchName is a required field
	BranchName *string `location:"uri" locationName:"branchName" min:"1" type:"string" required:"true"`

	// Optional file map that contains file name as the key and file content md5
	// hash as the value. If this argument is provided, the service will generate
	// different upload url per file. Otherwise, the service will only generate
	// a single upload url for the zipped files.
	FileMap map[string]string `locationName:"fileMap" type:"map"`
}

// String returns the string representation
func (s CreateDeploymentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDeploymentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDeploymentInput"}

	if s.AppId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AppId"))
	}
	if s.AppId != nil && len(*s.AppId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AppId", 1))
	}

	if s.BranchName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BranchName"))
	}
	if s.BranchName != nil && len(*s.BranchName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BranchName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Result structure for create a new deployment.
type CreateDeploymentOutput struct {
	_ struct{} `type:"structure"`

	// When the fileMap argument is provided in the request, the fileUploadUrls
	// will contain a map of file names to upload url.
	//
	// FileUploadUrls is a required field
	FileUploadUrls map[string]string `locationName:"fileUploadUrls" type:"map" required:"true"`

	// The jobId for this deployment, will supply to start deployment api.
	JobId *string `locationName:"jobId" type:"string"`

	// When the fileMap argument is NOT provided. This zipUploadUrl will be returned.
	//
	// ZipUploadUrl is a required field
	ZipUploadUrl *string `locationName:"zipUploadUrl" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateDeploymentOutput) String() string {
	return awsutil.Prettify(s)
}
