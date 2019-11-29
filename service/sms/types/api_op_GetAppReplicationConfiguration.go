// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetAppReplicationConfigurationInput struct {
	_ struct{} `type:"structure"`

	// ID of the application associated with the replication configuration.
	AppId *string `locationName:"appId" type:"string"`
}

// String returns the string representation
func (s GetAppReplicationConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

type GetAppReplicationConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// Replication configurations associated with server groups in this application.
	ServerGroupReplicationConfigurations []ServerGroupReplicationConfiguration `locationName:"serverGroupReplicationConfigurations" type:"list"`
}

// String returns the string representation
func (s GetAppReplicationConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}