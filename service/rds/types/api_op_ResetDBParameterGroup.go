// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ResetDBParameterGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the DB parameter group.
	//
	// Constraints:
	//
	//    * Must match the name of an existing DBParameterGroup.
	//
	// DBParameterGroupName is a required field
	DBParameterGroupName *string `type:"string" required:"true"`

	// To reset the entire DB parameter group, specify the DBParameterGroup name
	// and ResetAllParameters parameters. To reset specific parameters, provide
	// a list of the following: ParameterName and ApplyMethod. A maximum of 20 parameters
	// can be modified in a single request.
	//
	// MySQL
	//
	// Valid Values (for Apply method): immediate | pending-reboot
	//
	// You can use the immediate value with dynamic parameters only. You can use
	// the pending-reboot value for both dynamic and static parameters, and changes
	// are applied when DB instance reboots.
	//
	// MariaDB
	//
	// Valid Values (for Apply method): immediate | pending-reboot
	//
	// You can use the immediate value with dynamic parameters only. You can use
	// the pending-reboot value for both dynamic and static parameters, and changes
	// are applied when DB instance reboots.
	//
	// Oracle
	//
	// Valid Values (for Apply method): pending-reboot
	Parameters []Parameter `locationNameList:"Parameter" type:"list"`

	// A value that indicates whether to reset all parameters in the DB parameter
	// group to default values. By default, all parameters in the DB parameter group
	// are reset to default values.
	ResetAllParameters *bool `type:"boolean"`
}

// String returns the string representation
func (s ResetDBParameterGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResetDBParameterGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResetDBParameterGroupInput"}

	if s.DBParameterGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBParameterGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the result of a successful invocation of the ModifyDBParameterGroup
// or ResetDBParameterGroup action.
type ResetDBParameterGroupOutput struct {
	_ struct{} `type:"structure"`

	// Provides the name of the DB parameter group.
	DBParameterGroupName *string `type:"string"`
}

// String returns the string representation
func (s ResetDBParameterGroupOutput) String() string {
	return awsutil.Prettify(s)
}
