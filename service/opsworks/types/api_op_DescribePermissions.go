// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribePermissionsInput struct {
	_ struct{} `type:"structure"`

	// The user's IAM ARN. This can also be a federated user's ARN. For more information
	// about IAM ARNs, see Using Identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html).
	IamUserArn *string `type:"string"`

	// The stack ID.
	StackId *string `type:"string"`
}

// String returns the string representation
func (s DescribePermissionsInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the response to a DescribePermissions request.
type DescribePermissionsOutput struct {
	_ struct{} `type:"structure"`

	// An array of Permission objects that describe the stack permissions.
	//
	//    * If the request object contains only a stack ID, the array contains a
	//    Permission object with permissions for each of the stack IAM ARNs.
	//
	//    * If the request object contains only an IAM ARN, the array contains a
	//    Permission object with permissions for each of the user's stack IDs.
	//
	//    * If the request contains a stack ID and an IAM ARN, the array contains
	//    a single Permission object with permissions for the specified stack and
	//    IAM ARN.
	Permissions []Permission `type:"list"`
}

// String returns the string representation
func (s DescribePermissionsOutput) String() string {
	return awsutil.Prettify(s)
}
