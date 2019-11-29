// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request structure for get App request.
type GetAppInput struct {
	_ struct{} `type:"structure"`

	// Unique Id for an Amplify App.
	//
	// AppId is a required field
	AppId *string `location:"uri" locationName:"appId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetAppInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAppInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetAppInput"}

	if s.AppId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AppId"))
	}
	if s.AppId != nil && len(*s.AppId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AppId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetAppOutput struct {
	_ struct{} `type:"structure"`

	// Amplify App represents different branches of a repository for building, deploying,
	// and hosting.
	//
	// App is a required field
	App *App `locationName:"app" type:"structure" required:"true"`
}

// String returns the string representation
func (s GetAppOutput) String() string {
	return awsutil.Prettify(s)
}