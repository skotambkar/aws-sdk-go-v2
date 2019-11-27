// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// CreateLocationNfsRequest
type CreateLocationNfsInput struct {
	_ struct{} `type:"structure"`

	// The NFS mount options that DataSync can use to mount your NFS share.
	MountOptions *NfsMountOptions `type:"structure"`

	// Contains a list of Amazon Resource Names (ARNs) of agents that are used to
	// connect to an NFS server.
	//
	// OnPremConfig is a required field
	OnPremConfig *OnPremConfig `type:"structure" required:"true"`

	// The name of the NFS server. This value is the IP address or Domain Name Service
	// (DNS) name of the NFS server. An agent that is installed on-premises uses
	// this host name to mount the NFS server in a network.
	//
	// This name must either be DNS-compliant or must be an IP version 4 (IPv4)
	// address.
	//
	// ServerHostname is a required field
	ServerHostname *string `type:"string" required:"true"`

	// The subdirectory in the NFS file system that is used to read data from the
	// NFS source location or write data to the NFS destination. The NFS path should
	// be a path that's exported by the NFS server, or a subdirectory of that path.
	// The path should be such that it can be mounted by other NFS clients in your
	// network.
	//
	// To see all the paths exported by your NFS server. run "showmount -e nfs-server-name"
	// from an NFS client that has access to your server. You can specify any directory
	// that appears in the results, and any subdirectory of that directory. Ensure
	// that the NFS export is accessible without Kerberos authentication.
	//
	// To transfer all the data in the folder you specified, DataSync needs to have
	// permissions to read all the data. To ensure this, either configure the NFS
	// export with no_root_squash, or ensure that the permissions for all of the
	// files that you want DataSync allow read access for all users. Doing either
	// enables the agent to read the files. For the agent to access directories,
	// you must additionally enable all execute access.
	//
	// For information about NFS export configuration, see 18.7. The /etc/exports
	// Configuration File in the Red Hat Enterprise Linux documentation.
	//
	// Subdirectory is a required field
	Subdirectory *string `type:"string" required:"true"`

	// The key-value pair that represents the tag that you want to add to the location.
	// The value can be an empty string. We recommend using tags to name your resources.
	Tags []TagListEntry `type:"list"`
}

// String returns the string representation
func (s CreateLocationNfsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateLocationNfsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateLocationNfsInput"}

	if s.OnPremConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("OnPremConfig"))
	}

	if s.ServerHostname == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServerHostname"))
	}

	if s.Subdirectory == nil {
		invalidParams.Add(aws.NewErrParamRequired("Subdirectory"))
	}
	if s.OnPremConfig != nil {
		if err := s.OnPremConfig.Validate(); err != nil {
			invalidParams.AddNested("OnPremConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// CreateLocationNfsResponse
type CreateLocationNfsOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the source NFS file system location that
	// is created.
	LocationArn *string `type:"string"`
}

// String returns the string representation
func (s CreateLocationNfsOutput) String() string {
	return awsutil.Prettify(s)
}
