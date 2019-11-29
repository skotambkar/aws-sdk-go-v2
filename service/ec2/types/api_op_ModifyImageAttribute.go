// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ec2/enums"
)

// Contains the parameters for ModifyImageAttribute.
type ModifyImageAttributeInput struct {
	_ struct{} `type:"structure"`

	// The name of the attribute to modify. The valid values are description, launchPermission,
	// and productCodes.
	Attribute *string `type:"string"`

	// A new description for the AMI.
	Description *AttributeValue `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the AMI.
	//
	// ImageId is a required field
	ImageId *string `type:"string" required:"true"`

	// A new launch permission for the AMI.
	LaunchPermission *LaunchPermissionModifications `type:"structure"`

	// The operation type. This parameter can be used only when the Attribute parameter
	// is launchPermission.
	OperationType enums.OperationType `type:"string" enum:"true"`

	// The DevPay product codes. After you add a product code to an AMI, it can't
	// be removed.
	ProductCodes []string `locationName:"ProductCode" locationNameList:"ProductCode" type:"list"`

	// The user groups. This parameter can be used only when the Attribute parameter
	// is launchPermission.
	UserGroups []string `locationName:"UserGroup" locationNameList:"UserGroup" type:"list"`

	// The AWS account IDs. This parameter can be used only when the Attribute parameter
	// is launchPermission.
	UserIds []string `locationName:"UserId" locationNameList:"UserId" type:"list"`

	// The value of the attribute being modified. This parameter can be used only
	// when the Attribute parameter is description or productCodes.
	Value *string `type:"string"`
}

// String returns the string representation
func (s ModifyImageAttributeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyImageAttributeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyImageAttributeInput"}

	if s.ImageId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ImageId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyImageAttributeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ModifyImageAttributeOutput) String() string {
	return awsutil.Prettify(s)
}
