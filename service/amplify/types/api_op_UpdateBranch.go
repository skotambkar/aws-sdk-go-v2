// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/amplify/enums"
)

// Request structure for update branch request.
type UpdateBranchInput struct {
	_ struct{} `type:"structure"`

	// Unique Id for an Amplify App.
	//
	// AppId is a required field
	AppId *string `location:"uri" locationName:"appId" min:"1" type:"string" required:"true"`

	// Basic Authorization credentials for the branch.
	BasicAuthCredentials *string `locationName:"basicAuthCredentials" type:"string"`

	// Name for the branch.
	//
	// BranchName is a required field
	BranchName *string `location:"uri" locationName:"branchName" min:"1" type:"string" required:"true"`

	// BuildSpec for the branch.
	BuildSpec *string `locationName:"buildSpec" min:"1" type:"string"`

	// Description for the branch.
	Description *string `locationName:"description" type:"string"`

	// Display name for a branch, will use as the default domain prefix.
	DisplayName *string `locationName:"displayName" type:"string"`

	// Enables auto building for the branch.
	EnableAutoBuild *bool `locationName:"enableAutoBuild" type:"boolean"`

	// Enables Basic Auth for the branch.
	EnableBasicAuth *bool `locationName:"enableBasicAuth" type:"boolean"`

	// Enables notifications for the branch.
	EnableNotification *bool `locationName:"enableNotification" type:"boolean"`

	// Enables Pull Request Preview for this branch.
	EnablePullRequestPreview *bool `locationName:"enablePullRequestPreview" type:"boolean"`

	// Environment Variables for the branch.
	EnvironmentVariables map[string]string `locationName:"environmentVariables" type:"map"`

	// Framework for the branch.
	Framework *string `locationName:"framework" type:"string"`

	// Stage for the branch.
	Stage enums.Stage `locationName:"stage" type:"string" enum:"true"`

	// The content TTL for the website in seconds.
	Ttl *string `locationName:"ttl" type:"string"`
}

// String returns the string representation
func (s UpdateBranchInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateBranchInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateBranchInput"}

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
	if s.BuildSpec != nil && len(*s.BuildSpec) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BuildSpec", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Result structure for update branch request.
type UpdateBranchOutput struct {
	_ struct{} `type:"structure"`

	// Branch structure for an Amplify App.
	//
	// Branch is a required field
	Branch *Branch `locationName:"branch" type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateBranchOutput) String() string {
	return awsutil.Prettify(s)
}
