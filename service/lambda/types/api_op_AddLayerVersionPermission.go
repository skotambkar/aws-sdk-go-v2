// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AddLayerVersionPermissionInput struct {
	_ struct{} `type:"structure"`

	// The API action that grants access to the layer. For example, lambda:GetLayerVersion.
	//
	// Action is a required field
	Action *string `type:"string" required:"true"`

	// The name or Amazon Resource Name (ARN) of the layer.
	//
	// LayerName is a required field
	LayerName *string `location:"uri" locationName:"LayerName" min:"1" type:"string" required:"true"`

	// With the principal set to *, grant permission to all accounts in the specified
	// organization.
	OrganizationId *string `type:"string"`

	// An account ID, or * to grant permission to all AWS accounts.
	//
	// Principal is a required field
	Principal *string `type:"string" required:"true"`

	// Only update the policy if the revision ID matches the ID specified. Use this
	// option to avoid modifying a policy that has changed since you last read it.
	RevisionId *string `location:"querystring" locationName:"RevisionId" type:"string"`

	// An identifier that distinguishes the policy from others on the same layer
	// version.
	//
	// StatementId is a required field
	StatementId *string `min:"1" type:"string" required:"true"`

	// The version number.
	//
	// VersionNumber is a required field
	VersionNumber *int64 `location:"uri" locationName:"VersionNumber" type:"long" required:"true"`
}

// String returns the string representation
func (s AddLayerVersionPermissionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddLayerVersionPermissionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddLayerVersionPermissionInput"}

	if s.Action == nil {
		invalidParams.Add(aws.NewErrParamRequired("Action"))
	}

	if s.LayerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LayerName"))
	}
	if s.LayerName != nil && len(*s.LayerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LayerName", 1))
	}

	if s.Principal == nil {
		invalidParams.Add(aws.NewErrParamRequired("Principal"))
	}

	if s.StatementId == nil {
		invalidParams.Add(aws.NewErrParamRequired("StatementId"))
	}
	if s.StatementId != nil && len(*s.StatementId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StatementId", 1))
	}

	if s.VersionNumber == nil {
		invalidParams.Add(aws.NewErrParamRequired("VersionNumber"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AddLayerVersionPermissionOutput struct {
	_ struct{} `type:"structure"`

	// A unique identifier for the current revision of the policy.
	RevisionId *string `type:"string"`

	// The permission statement.
	Statement *string `type:"string"`
}

// String returns the string representation
func (s AddLayerVersionPermissionOutput) String() string {
	return awsutil.Prettify(s)
}