// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input to ResetDBClusterParameterGroup.
type ResetDBClusterParameterGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the DB cluster parameter group to reset.
	//
	// DBClusterParameterGroupName is a required field
	DBClusterParameterGroupName *string `type:"string" required:"true"`

	// A list of parameter names in the DB cluster parameter group to reset to the
	// default values. You can't use this parameter if the ResetAllParameters parameter
	// is set to true.
	Parameters []Parameter `locationNameList:"Parameter" type:"list"`

	// A value that is set to true to reset all parameters in the DB cluster parameter
	// group to their default values, and false otherwise. You can't use this parameter
	// if there is a list of parameter names specified for the Parameters parameter.
	ResetAllParameters *bool `type:"boolean"`
}

// String returns the string representation
func (s ResetDBClusterParameterGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResetDBClusterParameterGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResetDBClusterParameterGroupInput"}

	if s.DBClusterParameterGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterParameterGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the name of a DB cluster parameter group.
type ResetDBClusterParameterGroupOutput struct {
	_ struct{} `type:"structure"`

	// The name of a DB cluster parameter group.
	//
	// Constraints:
	//
	//    * Must be from 1 to 255 letters or numbers.
	//
	//    * The first character must be a letter.
	//
	//    * Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// This value is stored as a lowercase string.
	DBClusterParameterGroupName *string `type:"string"`
}

// String returns the string representation
func (s ResetDBClusterParameterGroupOutput) String() string {
	return awsutil.Prettify(s)
}
