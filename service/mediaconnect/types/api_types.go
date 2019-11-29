// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/enums"
)

// The output that you want to add to this flow.
type AddOutputRequest struct {
	_ struct{} `type:"structure"`

	// The range of IP addresses that should be allowed to initiate output requests
	// to this flow. These IP addresses should be in the form of a Classless Inter-Domain
	// Routing (CIDR) block; for example, 10.0.0.0/16.
	CidrAllowList []string `locationName:"cidrAllowList" type:"list"`

	// A description of the output. This description appears only on the AWS Elemental
	// MediaConnect console and will not be seen by the end user.
	Description *string `locationName:"description" type:"string"`

	// The IP address from which video will be sent to output destinations.
	Destination *string `locationName:"destination" type:"string"`

	// The type of key used for the encryption. If no keyType is provided, the service
	// will use the default setting (static-key).
	Encryption *Encryption `locationName:"encryption" type:"structure"`

	// The maximum latency in milliseconds for Zixi-based streams.
	MaxLatency *int64 `locationName:"maxLatency" type:"integer"`

	// The name of the output. This value must be unique within the current flow.
	Name *string `locationName:"name" type:"string"`

	// The port to use when content is distributed to this output.
	Port *int64 `locationName:"port" type:"integer"`

	// The protocol to use for the output.
	//
	// Protocol is a required field
	Protocol enums.Protocol `locationName:"protocol" type:"string" required:"true" enum:"true"`

	// The remote ID for the Zixi-pull output stream.
	RemoteId *string `locationName:"remoteId" type:"string"`

	// The smoothing latency in milliseconds for RIST, RTP, and RTP-FEC streams.
	SmoothingLatency *int64 `locationName:"smoothingLatency" type:"integer"`

	// The stream ID that you want to use for this transport. This parameter applies
	// only to Zixi-based streams.
	StreamId *string `locationName:"streamId" type:"string"`
}

// String returns the string representation
func (s AddOutputRequest) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddOutputRequest) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddOutputRequest"}
	if len(s.Protocol) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Protocol"))
	}
	if s.Encryption != nil {
		if err := s.Encryption.Validate(); err != nil {
			invalidParams.AddNested("Encryption", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about the encryption of the flow.
type Encryption struct {
	_ struct{} `type:"structure"`

	// The type of algorithm that is used for the encryption (such as aes128, aes192,
	// or aes256).
	//
	// Algorithm is a required field
	Algorithm enums.Algorithm `locationName:"algorithm" type:"string" required:"true" enum:"true"`

	// A 128-bit, 16-byte hex value represented by a 32-character string, to be
	// used with the key for encrypting content. This parameter is not valid for
	// static key encryption.
	ConstantInitializationVector *string `locationName:"constantInitializationVector" type:"string"`

	// The value of one of the devices that you configured with your digital rights
	// management (DRM) platform key provider. This parameter is required for SPEKE
	// encryption and is not valid for static key encryption.
	DeviceId *string `locationName:"deviceId" type:"string"`

	// The type of key that is used for the encryption. If no keyType is provided,
	// the service will use the default setting (static-key).
	KeyType enums.KeyType `locationName:"keyType" type:"string" enum:"true"`

	// The AWS Region that the API Gateway proxy endpoint was created in. This parameter
	// is required for SPEKE encryption and is not valid for static key encryption.
	Region *string `locationName:"region" type:"string"`

	// An identifier for the content. The service sends this value to the key server
	// to identify the current endpoint. The resource ID is also known as the content
	// ID. This parameter is required for SPEKE encryption and is not valid for
	// static key encryption.
	ResourceId *string `locationName:"resourceId" type:"string"`

	// The ARN of the role that you created during setup (when you set up AWS Elemental
	// MediaConnect as a trusted entity).
	//
	// RoleArn is a required field
	RoleArn *string `locationName:"roleArn" type:"string" required:"true"`

	// The ARN of the secret that you created in AWS Secrets Manager to store the
	// encryption key. This parameter is required for static key encryption and
	// is not valid for SPEKE encryption.
	SecretArn *string `locationName:"secretArn" type:"string"`

	// The URL from the API Gateway proxy that you set up to talk to your key server.
	// This parameter is required for SPEKE encryption and is not valid for static
	// key encryption.
	Url *string `locationName:"url" type:"string"`
}

// String returns the string representation
func (s Encryption) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Encryption) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Encryption"}
	if len(s.Algorithm) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Algorithm"))
	}

	if s.RoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The settings for a flow entitlement.
type Entitlement struct {
	_ struct{} `type:"structure"`

	// Percentage from 0-100 of the data transfer cost to be billed to the subscriber.
	DataTransferSubscriberFeePercent *int64 `locationName:"dataTransferSubscriberFeePercent" type:"integer"`

	// A description of the entitlement.
	Description *string `locationName:"description" type:"string"`

	// The type of encryption that will be used on the output that is associated
	// with this entitlement.
	Encryption *Encryption `locationName:"encryption" type:"structure"`

	// The ARN of the entitlement.
	//
	// EntitlementArn is a required field
	EntitlementArn *string `locationName:"entitlementArn" type:"string" required:"true"`

	// The name of the entitlement.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	// The AWS account IDs that you want to share your content with. The receiving
	// accounts (subscribers) will be allowed to create their own flow using your
	// content as the source.
	//
	// Subscribers is a required field
	Subscribers []string `locationName:"subscribers" type:"list" required:"true"`
}

// String returns the string representation
func (s Entitlement) String() string {
	return awsutil.Prettify(s)
}

// The settings for a flow, including its source, outputs, and entitlements.
type Flow struct {
	_ struct{} `type:"structure"`

	// The Availability Zone that you want to create the flow in. These options
	// are limited to the Availability Zones within the current AWS.
	//
	// AvailabilityZone is a required field
	AvailabilityZone *string `locationName:"availabilityZone" type:"string" required:"true"`

	// A description of the flow. This value is not used or seen outside of the
	// current AWS Elemental MediaConnect account.
	Description *string `locationName:"description" type:"string"`

	// The IP address from which video will be sent to output destinations.
	EgressIp *string `locationName:"egressIp" type:"string"`

	// The entitlements in this flow.
	//
	// Entitlements is a required field
	Entitlements []Entitlement `locationName:"entitlements" type:"list" required:"true"`

	// The Amazon Resource Name (ARN), a unique identifier for any AWS resource,
	// of the flow.
	//
	// FlowArn is a required field
	FlowArn *string `locationName:"flowArn" type:"string" required:"true"`

	// The name of the flow.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	// The outputs in this flow.
	//
	// Outputs is a required field
	Outputs []Output `locationName:"outputs" type:"list" required:"true"`

	// The settings for the source of the flow.
	//
	// Source is a required field
	Source *Source `locationName:"source" type:"structure" required:"true"`

	// The current status of the flow.
	//
	// Status is a required field
	Status enums.Status `locationName:"status" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s Flow) String() string {
	return awsutil.Prettify(s)
}

// The entitlements that you want to grant on a flow.
type GrantEntitlementRequest struct {
	_ struct{} `type:"structure"`

	// Percentage from 0-100 of the data transfer cost to be billed to the subscriber.
	DataTransferSubscriberFeePercent *int64 `locationName:"dataTransferSubscriberFeePercent" type:"integer"`

	// A description of the entitlement. This description appears only on the AWS
	// Elemental MediaConnect console and will not be seen by the subscriber or
	// end user.
	Description *string `locationName:"description" type:"string"`

	// The type of encryption that will be used on the output that is associated
	// with this entitlement.
	Encryption *Encryption `locationName:"encryption" type:"structure"`

	// The name of the entitlement. This value must be unique within the current
	// flow.
	Name *string `locationName:"name" type:"string"`

	// The AWS account IDs that you want to share your content with. The receiving
	// accounts (subscribers) will be allowed to create their own flows using your
	// content as the source.
	//
	// Subscribers is a required field
	Subscribers []string `locationName:"subscribers" type:"list" required:"true"`
}

// String returns the string representation
func (s GrantEntitlementRequest) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GrantEntitlementRequest) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GrantEntitlementRequest"}

	if s.Subscribers == nil {
		invalidParams.Add(aws.NewErrParamRequired("Subscribers"))
	}
	if s.Encryption != nil {
		if err := s.Encryption.Validate(); err != nil {
			invalidParams.AddNested("Encryption", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An entitlement that has been granted to you from other AWS accounts.
type ListedEntitlement struct {
	_ struct{} `type:"structure"`

	// Percentage from 0-100 of the data transfer cost to be billed to the subscriber.
	DataTransferSubscriberFeePercent *int64 `locationName:"dataTransferSubscriberFeePercent" type:"integer"`

	// The ARN of the entitlement.
	//
	// EntitlementArn is a required field
	EntitlementArn *string `locationName:"entitlementArn" type:"string" required:"true"`

	// The name of the entitlement.
	//
	// EntitlementName is a required field
	EntitlementName *string `locationName:"entitlementName" type:"string" required:"true"`
}

// String returns the string representation
func (s ListedEntitlement) String() string {
	return awsutil.Prettify(s)
}

// Provides a summary of a flow, including its ARN, Availability Zone, and source
// type.
type ListedFlow struct {
	_ struct{} `type:"structure"`

	// The Availability Zone that the flow was created in.
	//
	// AvailabilityZone is a required field
	AvailabilityZone *string `locationName:"availabilityZone" type:"string" required:"true"`

	// A description of the flow.
	//
	// Description is a required field
	Description *string `locationName:"description" type:"string" required:"true"`

	// The ARN of the flow.
	//
	// FlowArn is a required field
	FlowArn *string `locationName:"flowArn" type:"string" required:"true"`

	// The name of the flow.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	// The type of source. This value is either owned (originated somewhere other
	// than an AWS Elemental MediaConnect flow owned by another AWS account) or
	// entitled (originated at an AWS Elemental MediaConnect flow owned by another
	// AWS account).
	//
	// SourceType is a required field
	SourceType enums.SourceType `locationName:"sourceType" type:"string" required:"true" enum:"true"`

	// The current status of the flow.
	//
	// Status is a required field
	Status enums.Status `locationName:"status" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s ListedFlow) String() string {
	return awsutil.Prettify(s)
}

// Messages that provide the state of the flow.
type Messages struct {
	_ struct{} `type:"structure"`

	// A list of errors that might have been generated from processes on this flow.
	//
	// Errors is a required field
	Errors []string `locationName:"errors" type:"list" required:"true"`
}

// String returns the string representation
func (s Messages) String() string {
	return awsutil.Prettify(s)
}

// The settings for an output.
type Output struct {
	_ struct{} `type:"structure"`

	// Percentage from 0-100 of the data transfer cost to be billed to the subscriber.
	DataTransferSubscriberFeePercent *int64 `locationName:"dataTransferSubscriberFeePercent" type:"integer"`

	// A description of the output.
	Description *string `locationName:"description" type:"string"`

	// The address where you want to send the output.
	Destination *string `locationName:"destination" type:"string"`

	// The type of key used for the encryption. If no keyType is provided, the service
	// will use the default setting (static-key).
	Encryption *Encryption `locationName:"encryption" type:"structure"`

	// The ARN of the entitlement on the originator''s flow. This value is relevant
	// only on entitled flows.
	EntitlementArn *string `locationName:"entitlementArn" type:"string"`

	// The input ARN of the AWS Elemental MediaLive channel. This parameter is relevant
	// only for outputs that were added by creating a MediaLive input.
	MediaLiveInputArn *string `locationName:"mediaLiveInputArn" type:"string"`

	// The name of the output. This value must be unique within the current flow.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	// The ARN of the output.
	//
	// OutputArn is a required field
	OutputArn *string `locationName:"outputArn" type:"string" required:"true"`

	// The port to use when content is distributed to this output.
	Port *int64 `locationName:"port" type:"integer"`

	// Attributes related to the transport stream that are used in the output.
	Transport *Transport `locationName:"transport" type:"structure"`
}

// String returns the string representation
func (s Output) String() string {
	return awsutil.Prettify(s)
}

// The settings for the source of the flow.
type SetSourceRequest struct {
	_ struct{} `type:"structure"`

	// The type of encryption that is used on the content ingested from this source.
	Decryption *Encryption `locationName:"decryption" type:"structure"`

	// A description for the source. This value is not used or seen outside of the
	// current AWS Elemental MediaConnect account.
	Description *string `locationName:"description" type:"string"`

	// The ARN of the entitlement that allows you to subscribe to this flow. The
	// entitlement is set by the flow originator, and the ARN is generated as part
	// of the originator's flow.
	EntitlementArn *string `locationName:"entitlementArn" type:"string"`

	// The port that the flow will be listening on for incoming content.
	IngestPort *int64 `locationName:"ingestPort" type:"integer"`

	// The smoothing max bitrate for RIST, RTP, and RTP-FEC streams.
	MaxBitrate *int64 `locationName:"maxBitrate" type:"integer"`

	// The maximum latency in milliseconds. This parameter applies only to RIST-based
	// and Zixi-based streams.
	MaxLatency *int64 `locationName:"maxLatency" type:"integer"`

	// The name of the source.
	Name *string `locationName:"name" type:"string"`

	// The protocol that is used by the source.
	Protocol enums.Protocol `locationName:"protocol" type:"string" enum:"true"`

	// The stream ID that you want to use for this transport. This parameter applies
	// only to Zixi-based streams.
	StreamId *string `locationName:"streamId" type:"string"`

	// The range of IP addresses that should be allowed to contribute content to
	// your source. These IP addresses should be in the form of a Classless Inter-Domain
	// Routing (CIDR) block; for example, 10.0.0.0/16.
	WhitelistCidr *string `locationName:"whitelistCidr" type:"string"`
}

// String returns the string representation
func (s SetSourceRequest) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetSourceRequest) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetSourceRequest"}
	if s.Decryption != nil {
		if err := s.Decryption.Validate(); err != nil {
			invalidParams.AddNested("Decryption", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The settings for the source of the flow.
type Source struct {
	_ struct{} `type:"structure"`

	// Percentage from 0-100 of the data transfer cost to be billed to the subscriber.
	DataTransferSubscriberFeePercent *int64 `locationName:"dataTransferSubscriberFeePercent" type:"integer"`

	// The type of encryption that is used on the content ingested from this source.
	Decryption *Encryption `locationName:"decryption" type:"structure"`

	// A description for the source. This value is not used or seen outside of the
	// current AWS Elemental MediaConnect account.
	Description *string `locationName:"description" type:"string"`

	// The ARN of the entitlement that allows you to subscribe to content that comes
	// from another AWS account. The entitlement is set by the content originator
	// and the ARN is generated as part of the originator's flow.
	EntitlementArn *string `locationName:"entitlementArn" type:"string"`

	// The IP address that the flow will be listening on for incoming content.
	IngestIp *string `locationName:"ingestIp" type:"string"`

	// The port that the flow will be listening on for incoming content.
	IngestPort *int64 `locationName:"ingestPort" type:"integer"`

	// The name of the source.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	// The ARN of the source.
	//
	// SourceArn is a required field
	SourceArn *string `locationName:"sourceArn" type:"string" required:"true"`

	// Attributes related to the transport stream that are used in the source.
	Transport *Transport `locationName:"transport" type:"structure"`

	// The range of IP addresses that should be allowed to contribute content to
	// your source. These IP addresses should be in the form of a Classless Inter-Domain
	// Routing (CIDR) block; for example, 10.0.0.0/16.
	WhitelistCidr *string `locationName:"whitelistCidr" type:"string"`
}

// String returns the string representation
func (s Source) String() string {
	return awsutil.Prettify(s)
}

// Attributes related to the transport stream that are used in a source or output.
type Transport struct {
	_ struct{} `type:"structure"`

	// The range of IP addresses that should be allowed to initiate output requests
	// to this flow. These IP addresses should be in the form of a Classless Inter-Domain
	// Routing (CIDR) block; for example, 10.0.0.0/16.
	CidrAllowList []string `locationName:"cidrAllowList" type:"list"`

	// The smoothing max bitrate for RIST, RTP, and RTP-FEC streams.
	MaxBitrate *int64 `locationName:"maxBitrate" type:"integer"`

	// The maximum latency in milliseconds. This parameter applies only to RIST-based
	// and Zixi-based streams.
	MaxLatency *int64 `locationName:"maxLatency" type:"integer"`

	// The protocol that is used by the source or output.
	//
	// Protocol is a required field
	Protocol enums.Protocol `locationName:"protocol" type:"string" required:"true" enum:"true"`

	// The remote ID for the Zixi-pull stream.
	RemoteId *string `locationName:"remoteId" type:"string"`

	// The smoothing latency in milliseconds for RIST, RTP, and RTP-FEC streams.
	SmoothingLatency *int64 `locationName:"smoothingLatency" type:"integer"`

	// The stream ID that you want to use for this transport. This parameter applies
	// only to Zixi-based streams.
	StreamId *string `locationName:"streamId" type:"string"`
}

// String returns the string representation
func (s Transport) String() string {
	return awsutil.Prettify(s)
}

// Information about the encryption of the flow.
type UpdateEncryption struct {
	_ struct{} `type:"structure"`

	// The type of algorithm that is used for the encryption (such as aes128, aes192,
	// or aes256).
	Algorithm enums.Algorithm `locationName:"algorithm" type:"string" enum:"true"`

	// A 128-bit, 16-byte hex value represented by a 32-character string, to be
	// used with the key for encrypting content. This parameter is not valid for
	// static key encryption.
	ConstantInitializationVector *string `locationName:"constantInitializationVector" type:"string"`

	// The value of one of the devices that you configured with your digital rights
	// management (DRM) platform key provider. This parameter is required for SPEKE
	// encryption and is not valid for static key encryption.
	DeviceId *string `locationName:"deviceId" type:"string"`

	// The type of key that is used for the encryption. If no keyType is provided,
	// the service will use the default setting (static-key).
	KeyType enums.KeyType `locationName:"keyType" type:"string" enum:"true"`

	// The AWS Region that the API Gateway proxy endpoint was created in. This parameter
	// is required for SPEKE encryption and is not valid for static key encryption.
	Region *string `locationName:"region" type:"string"`

	// An identifier for the content. The service sends this value to the key server
	// to identify the current endpoint. The resource ID is also known as the content
	// ID. This parameter is required for SPEKE encryption and is not valid for
	// static key encryption.
	ResourceId *string `locationName:"resourceId" type:"string"`

	// The ARN of the role that you created during setup (when you set up AWS Elemental
	// MediaConnect as a trusted entity).
	RoleArn *string `locationName:"roleArn" type:"string"`

	// The ARN of the secret that you created in AWS Secrets Manager to store the
	// encryption key. This parameter is required for static key encryption and
	// is not valid for SPEKE encryption.
	SecretArn *string `locationName:"secretArn" type:"string"`

	// The URL from the API Gateway proxy that you set up to talk to your key server.
	// This parameter is required for SPEKE encryption and is not valid for static
	// key encryption.
	Url *string `locationName:"url" type:"string"`
}

// String returns the string representation
func (s UpdateEncryption) String() string {
	return awsutil.Prettify(s)
}