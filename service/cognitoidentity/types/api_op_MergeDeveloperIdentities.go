// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Input to the MergeDeveloperIdentities action.
type MergeDeveloperIdentitiesInput struct {
	_ struct{} `type:"structure"`

	// User identifier for the destination user. The value should be a DeveloperUserIdentifier.
	//
	// DestinationUserIdentifier is a required field
	DestinationUserIdentifier *string `min:"1" type:"string" required:"true"`

	// The "domain" by which Cognito will refer to your users. This is a (pseudo)
	// domain name that you provide while creating an identity pool. This name acts
	// as a placeholder that allows your backend and the Cognito service to communicate
	// about the developer provider. For the DeveloperProviderName, you can use
	// letters as well as period (.), underscore (_), and dash (-).
	//
	// DeveloperProviderName is a required field
	DeveloperProviderName *string `min:"1" type:"string" required:"true"`

	// An identity pool ID in the format REGION:GUID.
	//
	// IdentityPoolId is a required field
	IdentityPoolId *string `min:"1" type:"string" required:"true"`

	// User identifier for the source user. The value should be a DeveloperUserIdentifier.
	//
	// SourceUserIdentifier is a required field
	SourceUserIdentifier *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s MergeDeveloperIdentitiesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MergeDeveloperIdentitiesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "MergeDeveloperIdentitiesInput"}

	if s.DestinationUserIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationUserIdentifier"))
	}
	if s.DestinationUserIdentifier != nil && len(*s.DestinationUserIdentifier) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DestinationUserIdentifier", 1))
	}

	if s.DeveloperProviderName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeveloperProviderName"))
	}
	if s.DeveloperProviderName != nil && len(*s.DeveloperProviderName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeveloperProviderName", 1))
	}

	if s.IdentityPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdentityPoolId"))
	}
	if s.IdentityPoolId != nil && len(*s.IdentityPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdentityPoolId", 1))
	}

	if s.SourceUserIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceUserIdentifier"))
	}
	if s.SourceUserIdentifier != nil && len(*s.SourceUserIdentifier) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SourceUserIdentifier", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Returned in response to a successful MergeDeveloperIdentities action.
type MergeDeveloperIdentitiesOutput struct {
	_ struct{} `type:"structure"`

	// A unique identifier in the format REGION:GUID.
	IdentityId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s MergeDeveloperIdentitiesOutput) String() string {
	return awsutil.Prettify(s)
}