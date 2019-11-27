// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyClusterParameterGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the parameter group to be modified.
	//
	// ParameterGroupName is a required field
	ParameterGroupName *string `type:"string" required:"true"`

	// An array of parameters to be modified. A maximum of 20 parameters can be
	// modified in a single request.
	//
	// For each parameter to be modified, you must supply at least the parameter
	// name and parameter value; other name-value pairs of the parameter are optional.
	//
	// For the workload management (WLM) configuration, you must supply all the
	// name-value pairs in the wlm_json_configuration parameter.
	//
	// Parameters is a required field
	Parameters []Parameter `locationNameList:"Parameter" type:"list" required:"true"`
}

// String returns the string representation
func (s ModifyClusterParameterGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyClusterParameterGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyClusterParameterGroupInput"}

	if s.ParameterGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParameterGroupName"))
	}

	if s.Parameters == nil {
		invalidParams.Add(aws.NewErrParamRequired("Parameters"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyClusterParameterGroupOutput struct {
	_ struct{} `type:"structure"`

	// The name of the cluster parameter group.
	ParameterGroupName *string `type:"string"`

	// The status of the parameter group. For example, if you made a change to a
	// parameter group name-value pair, then the change could be pending a reboot
	// of an associated cluster.
	ParameterGroupStatus *string `type:"string"`
}

// String returns the string representation
func (s ModifyClusterParameterGroupOutput) String() string {
	return awsutil.Prettify(s)
}
