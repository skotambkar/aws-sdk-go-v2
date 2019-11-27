// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/appstream/enums"
)

// Describes an interface VPC endpoint (interface endpoint) that lets you create
// a private connection between the virtual private cloud (VPC) that you specify
// and AppStream 2.0. When you specify an interface endpoint for a stack, users
// of the stack can connect to AppStream 2.0 only through that endpoint. When
// you specify an interface endpoint for an image builder, administrators can
// connect to the image builder only through that endpoint.
type AccessEndpoint struct {
	_ struct{} `type:"structure"`

	// The type of interface endpoint.
	//
	// EndpointType is a required field
	EndpointType enums.AccessEndpointType `type:"string" required:"true" enum:"true"`

	// The identifier (ID) of the VPC in which the interface endpoint is used.
	VpceId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s AccessEndpoint) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AccessEndpoint) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AccessEndpoint"}
	if len(s.EndpointType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("EndpointType"))
	}
	if s.VpceId != nil && len(*s.VpceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VpceId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes an application in the application catalog.
type Application struct {
	_ struct{} `type:"structure"`

	// The application name to display.
	DisplayName *string `min:"1" type:"string"`

	// If there is a problem, the application can be disabled after image creation.
	Enabled *bool `type:"boolean"`

	// The URL for the application icon. This URL might be time-limited.
	IconURL *string `min:"1" type:"string"`

	// The arguments that are passed to the application at launch.
	LaunchParameters *string `min:"1" type:"string"`

	// The path to the application executable in the instance.
	LaunchPath *string `min:"1" type:"string"`

	// Additional attributes that describe the application.
	Metadata map[string]string `type:"map"`

	// The name of the application.
	Name *string `min:"1" type:"string"`
}

// String returns the string representation
func (s Application) String() string {
	return awsutil.Prettify(s)
}

// The persistent application settings for users of a stack.
type ApplicationSettings struct {
	_ struct{} `type:"structure"`

	// Enables or disables persistent application settings for users during their
	// streaming sessions.
	//
	// Enabled is a required field
	Enabled *bool `type:"boolean" required:"true"`

	// The path prefix for the S3 bucket where users’ persistent application settings
	// are stored. You can allow the same persistent application settings to be
	// used across multiple stacks by specifying the same settings group for each
	// stack.
	SettingsGroup *string `type:"string"`
}

// String returns the string representation
func (s ApplicationSettings) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ApplicationSettings) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ApplicationSettings"}

	if s.Enabled == nil {
		invalidParams.Add(aws.NewErrParamRequired("Enabled"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes the persistent application settings for users of a stack.
type ApplicationSettingsResponse struct {
	_ struct{} `type:"structure"`

	// Specifies whether persistent application settings are enabled for users during
	// their streaming sessions.
	Enabled *bool `type:"boolean"`

	// The S3 bucket where users’ persistent application settings are stored.
	// When persistent application settings are enabled for the first time for an
	// account in an AWS Region, an S3 bucket is created. The bucket is unique to
	// the AWS account and the Region.
	S3BucketName *string `min:"1" type:"string"`

	// The path prefix for the S3 bucket where users’ persistent application settings
	// are stored.
	SettingsGroup *string `type:"string"`
}

// String returns the string representation
func (s ApplicationSettingsResponse) String() string {
	return awsutil.Prettify(s)
}

// Describes the capacity for a fleet.
type ComputeCapacity struct {
	_ struct{} `type:"structure"`

	// The desired number of streaming instances.
	//
	// DesiredInstances is a required field
	DesiredInstances *int64 `type:"integer" required:"true"`
}

// String returns the string representation
func (s ComputeCapacity) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ComputeCapacity) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ComputeCapacity"}

	if s.DesiredInstances == nil {
		invalidParams.Add(aws.NewErrParamRequired("DesiredInstances"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes the capacity status for a fleet.
type ComputeCapacityStatus struct {
	_ struct{} `type:"structure"`

	// The number of currently available instances that can be used to stream sessions.
	Available *int64 `type:"integer"`

	// The desired number of streaming instances.
	//
	// Desired is a required field
	Desired *int64 `type:"integer" required:"true"`

	// The number of instances in use for streaming.
	InUse *int64 `type:"integer"`

	// The total number of simultaneous streaming instances that are running.
	Running *int64 `type:"integer"`
}

// String returns the string representation
func (s ComputeCapacityStatus) String() string {
	return awsutil.Prettify(s)
}

// Describes the configuration information required to join fleets and image
// builders to Microsoft Active Directory domains.
type DirectoryConfig struct {
	_ struct{} `type:"structure"`

	// The time the directory configuration was created.
	CreatedTime *time.Time `type:"timestamp"`

	// The fully qualified name of the directory (for example, corp.example.com).
	//
	// DirectoryName is a required field
	DirectoryName *string `type:"string" required:"true"`

	// The distinguished names of the organizational units for computer accounts.
	OrganizationalUnitDistinguishedNames []string `type:"list"`

	// The credentials for the service account used by the fleet or image builder
	// to connect to the directory.
	ServiceAccountCredentials *ServiceAccountCredentials `type:"structure"`
}

// String returns the string representation
func (s DirectoryConfig) String() string {
	return awsutil.Prettify(s)
}

// Describes the configuration information required to join fleets and image
// builders to Microsoft Active Directory domains.
type DomainJoinInfo struct {
	_ struct{} `type:"structure"`

	// The fully qualified name of the directory (for example, corp.example.com).
	DirectoryName *string `type:"string"`

	// The distinguished name of the organizational unit for computer accounts.
	OrganizationalUnitDistinguishedName *string `type:"string"`
}

// String returns the string representation
func (s DomainJoinInfo) String() string {
	return awsutil.Prettify(s)
}

// Describes a fleet.
type Fleet struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) for the fleet.
	//
	// Arn is a required field
	Arn *string `type:"string" required:"true"`

	// The capacity status for the fleet.
	//
	// ComputeCapacityStatus is a required field
	ComputeCapacityStatus *ComputeCapacityStatus `type:"structure" required:"true"`

	// The time the fleet was created.
	CreatedTime *time.Time `type:"timestamp"`

	// The description to display.
	Description *string `min:"1" type:"string"`

	// The amount of time that a streaming session remains active after users disconnect.
	// If they try to reconnect to the streaming session after a disconnection or
	// network interruption within this time interval, they are connected to their
	// previous session. Otherwise, they are connected to a new session with a new
	// streaming instance.
	//
	// Specify a value between 60 and 360000.
	DisconnectTimeoutInSeconds *int64 `type:"integer"`

	// The fleet name to display.
	DisplayName *string `min:"1" type:"string"`

	// The name of the directory and organizational unit (OU) to use to join the
	// fleet to a Microsoft Active Directory domain.
	DomainJoinInfo *DomainJoinInfo `type:"structure"`

	// Indicates whether default internet access is enabled for the fleet.
	EnableDefaultInternetAccess *bool `type:"boolean"`

	// The fleet errors.
	FleetErrors []FleetError `type:"list"`

	// The fleet type.
	//
	// ALWAYS_ON
	//
	// Provides users with instant-on access to their apps. You are charged for
	// all running instances in your fleet, even if no users are streaming apps.
	//
	// ON_DEMAND
	//
	// Provide users with access to applications after they connect, which takes
	// one to two minutes. You are charged for instance streaming when users are
	// connected and a small hourly fee for instances that are not streaming apps.
	FleetType enums.FleetType `type:"string" enum:"true"`

	// The ARN of the IAM role that is applied to the fleet. To assume a role, the
	// fleet instance calls the AWS Security Token Service (STS) AssumeRole API
	// operation and passes the ARN of the role to use. The operation creates a
	// new session with temporary credentials.
	IamRoleArn *string `type:"string"`

	// The amount of time that users can be idle (inactive) before they are disconnected
	// from their streaming session and the DisconnectTimeoutInSeconds time interval
	// begins. Users are notified before they are disconnected due to inactivity.
	// If users try to reconnect to the streaming session before the time interval
	// specified in DisconnectTimeoutInSeconds elapses, they are connected to their
	// previous session. Users are considered idle when they stop providing keyboard
	// or mouse input during their streaming session. File uploads and downloads,
	// audio in, audio out, and pixels changing do not qualify as user activity.
	// If users continue to be idle after the time interval in IdleDisconnectTimeoutInSeconds
	// elapses, they are disconnected.
	//
	// To prevent users from being disconnected due to inactivity, specify a value
	// of 0. Otherwise, specify a value between 60 and 3600. The default value is
	// 0.
	//
	// If you enable this feature, we recommend that you specify a value that corresponds
	// exactly to a whole number of minutes (for example, 60, 120, and 180). If
	// you don't do this, the value is rounded to the nearest minute. For example,
	// if you specify a value of 70, users are disconnected after 1 minute of inactivity.
	// If you specify a value that is at the midpoint between two different minutes,
	// the value is rounded up. For example, if you specify a value of 90, users
	// are disconnected after 2 minutes of inactivity.
	IdleDisconnectTimeoutInSeconds *int64 `type:"integer"`

	// The ARN for the public, private, or shared image.
	ImageArn *string `type:"string"`

	// The name of the image used to create the fleet.
	ImageName *string `min:"1" type:"string"`

	// The instance type to use when launching fleet instances.
	//
	// InstanceType is a required field
	InstanceType *string `min:"1" type:"string" required:"true"`

	// The maximum amount of time that a streaming session can remain active, in
	// seconds. If users are still connected to a streaming instance five minutes
	// before this limit is reached, they are prompted to save any open documents
	// before being disconnected. After this time elapses, the instance is terminated
	// and replaced by a new instance.
	//
	// Specify a value between 600 and 360000.
	MaxUserDurationInSeconds *int64 `type:"integer"`

	// The name of the fleet.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The current state for the fleet.
	//
	// State is a required field
	State enums.FleetState `type:"string" required:"true" enum:"true"`

	// The VPC configuration for the fleet.
	VpcConfig *VpcConfig `type:"structure"`
}

// String returns the string representation
func (s Fleet) String() string {
	return awsutil.Prettify(s)
}

// Describes a fleet error.
type FleetError struct {
	_ struct{} `type:"structure"`

	// The error code.
	ErrorCode enums.FleetErrorCode `type:"string" enum:"true"`

	// The error message.
	ErrorMessage *string `min:"1" type:"string"`
}

// String returns the string representation
func (s FleetError) String() string {
	return awsutil.Prettify(s)
}

// Describes an image.
type Image struct {
	_ struct{} `type:"structure"`

	// The applications associated with the image.
	Applications []Application `type:"list"`

	// The version of the AppStream 2.0 agent to use for instances that are launched
	// from this image.
	AppstreamAgentVersion *string `min:"1" type:"string"`

	// The ARN of the image.
	Arn *string `type:"string"`

	// The ARN of the image from which this image was created.
	BaseImageArn *string `type:"string"`

	// The time the image was created.
	CreatedTime *time.Time `type:"timestamp"`

	// The description to display.
	Description *string `min:"1" type:"string"`

	// The image name to display.
	DisplayName *string `min:"1" type:"string"`

	// The name of the image builder that was used to create the private image.
	// If the image is shared, this value is null.
	ImageBuilderName *string `min:"1" type:"string"`

	// Indicates whether an image builder can be launched from this image.
	ImageBuilderSupported *bool `type:"boolean"`

	// The permissions to provide to the destination AWS account for the specified
	// image.
	ImagePermissions *ImagePermissions `type:"structure"`

	// The name of the image.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The operating system platform of the image.
	Platform enums.PlatformType `type:"string" enum:"true"`

	// The release date of the public base image. For private images, this date
	// is the release date of the base image from which the image was created.
	PublicBaseImageReleasedDate *time.Time `type:"timestamp"`

	// The image starts in the PENDING state. If image creation succeeds, the state
	// is AVAILABLE. If image creation fails, the state is FAILED.
	State enums.ImageState `type:"string" enum:"true"`

	// The reason why the last state change occurred.
	StateChangeReason *ImageStateChangeReason `type:"structure"`

	// Indicates whether the image is public or private.
	Visibility enums.VisibilityType `type:"string" enum:"true"`
}

// String returns the string representation
func (s Image) String() string {
	return awsutil.Prettify(s)
}

// Describes a virtual machine that is used to create an image.
type ImageBuilder struct {
	_ struct{} `type:"structure"`

	// The list of virtual private cloud (VPC) interface endpoint objects. Administrators
	// can connect to the image builder only through the specified endpoints.
	AccessEndpoints []AccessEndpoint `min:"1" type:"list"`

	// The version of the AppStream 2.0 agent that is currently being used by the
	// image builder.
	AppstreamAgentVersion *string `min:"1" type:"string"`

	// The ARN for the image builder.
	Arn *string `type:"string"`

	// The time stamp when the image builder was created.
	CreatedTime *time.Time `type:"timestamp"`

	// The description to display.
	Description *string `min:"1" type:"string"`

	// The image builder name to display.
	DisplayName *string `min:"1" type:"string"`

	// The name of the directory and organizational unit (OU) to use to join the
	// image builder to a Microsoft Active Directory domain.
	DomainJoinInfo *DomainJoinInfo `type:"structure"`

	// Enables or disables default internet access for the image builder.
	EnableDefaultInternetAccess *bool `type:"boolean"`

	// The ARN of the IAM role that is applied to the image builder. To assume a
	// role, the image builder calls the AWS Security Token Service (STS) AssumeRole
	// API operation and passes the ARN of the role to use. The operation creates
	// a new session with temporary credentials.
	IamRoleArn *string `type:"string"`

	// The ARN of the image from which this builder was created.
	ImageArn *string `type:"string"`

	// The image builder errors.
	ImageBuilderErrors []ResourceError `type:"list"`

	// The instance type for the image builder.
	InstanceType *string `min:"1" type:"string"`

	// The name of the image builder.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// Describes the network details of the fleet or image builder instance.
	NetworkAccessConfiguration *NetworkAccessConfiguration `type:"structure"`

	// The operating system platform of the image builder.
	Platform enums.PlatformType `type:"string" enum:"true"`

	// The state of the image builder.
	State enums.ImageBuilderState `type:"string" enum:"true"`

	// The reason why the last state change occurred.
	StateChangeReason *ImageBuilderStateChangeReason `type:"structure"`

	// The VPC configuration of the image builder.
	VpcConfig *VpcConfig `type:"structure"`
}

// String returns the string representation
func (s ImageBuilder) String() string {
	return awsutil.Prettify(s)
}

// Describes the reason why the last image builder state change occurred.
type ImageBuilderStateChangeReason struct {
	_ struct{} `type:"structure"`

	// The state change reason code.
	Code enums.ImageBuilderStateChangeReasonCode `type:"string" enum:"true"`

	// The state change reason message.
	Message *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ImageBuilderStateChangeReason) String() string {
	return awsutil.Prettify(s)
}

// Describes the permissions for an image.
type ImagePermissions struct {
	_ struct{} `type:"structure"`

	// Indicates whether the image can be used for a fleet.
	AllowFleet *bool `locationName:"allowFleet" type:"boolean"`

	// Indicates whether the image can be used for an image builder.
	AllowImageBuilder *bool `locationName:"allowImageBuilder" type:"boolean"`
}

// String returns the string representation
func (s ImagePermissions) String() string {
	return awsutil.Prettify(s)
}

// Describes the reason why the last image state change occurred.
type ImageStateChangeReason struct {
	_ struct{} `type:"structure"`

	// The state change reason code.
	Code enums.ImageStateChangeReasonCode `type:"string" enum:"true"`

	// The state change reason message.
	Message *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ImageStateChangeReason) String() string {
	return awsutil.Prettify(s)
}

// Describes the error that is returned when a usage report can't be generated.
type LastReportGenerationExecutionError struct {
	_ struct{} `type:"structure"`

	// The error code for the error that is returned when a usage report can't be
	// generated.
	ErrorCode enums.UsageReportExecutionErrorCode `type:"string" enum:"true"`

	// The error message for the error that is returned when a usage report can't
	// be generated.
	ErrorMessage *string `min:"1" type:"string"`
}

// String returns the string representation
func (s LastReportGenerationExecutionError) String() string {
	return awsutil.Prettify(s)
}

// Describes the network details of the fleet or image builder instance.
type NetworkAccessConfiguration struct {
	_ struct{} `type:"structure"`

	// The resource identifier of the elastic network interface that is attached
	// to instances in your VPC. All network interfaces have the eni-xxxxxxxx resource
	// identifier.
	EniId *string `min:"1" type:"string"`

	// The private IP address of the elastic network interface that is attached
	// to instances in your VPC.
	EniPrivateIpAddress *string `min:"1" type:"string"`
}

// String returns the string representation
func (s NetworkAccessConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Describes a resource error.
type ResourceError struct {
	_ struct{} `type:"structure"`

	// The error code.
	ErrorCode enums.FleetErrorCode `type:"string" enum:"true"`

	// The error message.
	ErrorMessage *string `min:"1" type:"string"`

	// The time the error occurred.
	ErrorTimestamp *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s ResourceError) String() string {
	return awsutil.Prettify(s)
}

// Describes the credentials for the service account used by the fleet or image
// builder to connect to the directory.
type ServiceAccountCredentials struct {
	_ struct{} `type:"structure"`

	// The user name of the account. This account must have the following privileges:
	// create computer objects, join computers to the domain, and change/reset the
	// password on descendant computer objects for the organizational units specified.
	//
	// AccountName is a required field
	AccountName *string `min:"1" type:"string" required:"true" sensitive:"true"`

	// The password for the account.
	//
	// AccountPassword is a required field
	AccountPassword *string `min:"1" type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s ServiceAccountCredentials) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ServiceAccountCredentials) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ServiceAccountCredentials"}

	if s.AccountName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountName"))
	}
	if s.AccountName != nil && len(*s.AccountName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AccountName", 1))
	}

	if s.AccountPassword == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountPassword"))
	}
	if s.AccountPassword != nil && len(*s.AccountPassword) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AccountPassword", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes a streaming session.
type Session struct {
	_ struct{} `type:"structure"`

	// The authentication method. The user is authenticated using a streaming URL
	// (API) or SAML 2.0 federation (SAML).
	AuthenticationType enums.AuthenticationType `type:"string" enum:"true"`

	// Specifies whether a user is connected to the streaming session.
	ConnectionState enums.SessionConnectionState `type:"string" enum:"true"`

	// The name of the fleet for the streaming session.
	//
	// FleetName is a required field
	FleetName *string `min:"1" type:"string" required:"true"`

	// The identifier of the streaming session.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`

	// The time when the streaming session is set to expire. This time is based
	// on the MaxUserDurationinSeconds value, which determines the maximum length
	// of time that a streaming session can run. A streaming session might end earlier
	// than the time specified in SessionMaxExpirationTime, when the DisconnectTimeOutInSeconds
	// elapses or the user chooses to end his or her session. If the DisconnectTimeOutInSeconds
	// elapses, or the user chooses to end his or her session, the streaming instance
	// is terminated and the streaming session ends.
	MaxExpirationTime *time.Time `type:"timestamp"`

	// The network details for the streaming session.
	NetworkAccessConfiguration *NetworkAccessConfiguration `type:"structure"`

	// The name of the stack for the streaming session.
	//
	// StackName is a required field
	StackName *string `min:"1" type:"string" required:"true"`

	// The time when a streaming instance is dedicated for the user.
	StartTime *time.Time `type:"timestamp"`

	// The current state of the streaming session.
	//
	// State is a required field
	State enums.SessionState `type:"string" required:"true" enum:"true"`

	// The identifier of the user for whom the session was created.
	//
	// UserId is a required field
	UserId *string `min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s Session) String() string {
	return awsutil.Prettify(s)
}

// Describes the permissions that are available to the specified AWS account
// for a shared image.
type SharedImagePermissions struct {
	_ struct{} `type:"structure"`

	// Describes the permissions for a shared image.
	//
	// ImagePermissions is a required field
	ImagePermissions *ImagePermissions `locationName:"imagePermissions" type:"structure" required:"true"`

	// The 12-digit identifier of the AWS account with which the image is shared.
	//
	// SharedAccountId is a required field
	SharedAccountId *string `locationName:"sharedAccountId" type:"string" required:"true"`
}

// String returns the string representation
func (s SharedImagePermissions) String() string {
	return awsutil.Prettify(s)
}

// Describes a stack.
type Stack struct {
	_ struct{} `type:"structure"`

	// The list of virtual private cloud (VPC) interface endpoint objects. Users
	// of the stack can connect to AppStream 2.0 only through the specified endpoints.
	AccessEndpoints []AccessEndpoint `min:"1" type:"list"`

	// The persistent application settings for users of the stack.
	ApplicationSettings *ApplicationSettingsResponse `type:"structure"`

	// The ARN of the stack.
	Arn *string `type:"string"`

	// The time the stack was created.
	CreatedTime *time.Time `type:"timestamp"`

	// The description to display.
	Description *string `min:"1" type:"string"`

	// The stack name to display.
	DisplayName *string `min:"1" type:"string"`

	// The URL that users are redirected to after they click the Send Feedback link.
	// If no URL is specified, no Send Feedback link is displayed.
	FeedbackURL *string `type:"string"`

	// The name of the stack.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The URL that users are redirected to after their streaming session ends.
	RedirectURL *string `type:"string"`

	// The errors for the stack.
	StackErrors []StackError `type:"list"`

	// The storage connectors to enable.
	StorageConnectors []StorageConnector `type:"list"`

	// The actions that are enabled or disabled for users during their streaming
	// sessions. By default these actions are enabled.
	UserSettings []UserSetting `min:"1" type:"list"`
}

// String returns the string representation
func (s Stack) String() string {
	return awsutil.Prettify(s)
}

// Describes a stack error.
type StackError struct {
	_ struct{} `type:"structure"`

	// The error code.
	ErrorCode enums.StackErrorCode `type:"string" enum:"true"`

	// The error message.
	ErrorMessage *string `min:"1" type:"string"`
}

// String returns the string representation
func (s StackError) String() string {
	return awsutil.Prettify(s)
}

// Describes a connector that enables persistent storage for users.
type StorageConnector struct {
	_ struct{} `type:"structure"`

	// The type of storage connector.
	//
	// ConnectorType is a required field
	ConnectorType enums.StorageConnectorType `type:"string" required:"true" enum:"true"`

	// The names of the domains for the account.
	Domains []string `type:"list"`

	// The ARN of the storage connector.
	ResourceIdentifier *string `min:"1" type:"string"`
}

// String returns the string representation
func (s StorageConnector) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StorageConnector) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StorageConnector"}
	if len(s.ConnectorType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ConnectorType"))
	}
	if s.ResourceIdentifier != nil && len(*s.ResourceIdentifier) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceIdentifier", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes information about the usage report subscription.
type UsageReportSubscription struct {
	_ struct{} `type:"structure"`

	// The time when the last usage report was generated.
	LastGeneratedReportDate *time.Time `type:"timestamp"`

	// The Amazon S3 bucket where generated reports are stored.
	//
	// If you enabled on-instance session scripts and Amazon S3 logging for your
	// session script configuration, AppStream 2.0 created an S3 bucket to store
	// the script output. The bucket is unique to your account and Region. When
	// you enable usage reporting in this case, AppStream 2.0 uses the same bucket
	// to store your usage reports. If you haven't already enabled on-instance session
	// scripts, when you enable usage reports, AppStream 2.0 creates a new S3 bucket.
	S3BucketName *string `min:"1" type:"string"`

	// The schedule for generating usage reports.
	Schedule enums.UsageReportSchedule `type:"string" enum:"true"`

	// The errors that were returned if usage reports couldn't be generated.
	SubscriptionErrors []LastReportGenerationExecutionError `type:"list"`
}

// String returns the string representation
func (s UsageReportSubscription) String() string {
	return awsutil.Prettify(s)
}

// Describes a user in the user pool.
type User struct {
	_ struct{} `type:"structure"`

	// The ARN of the user.
	Arn *string `type:"string"`

	// The authentication type for the user.
	//
	// AuthenticationType is a required field
	AuthenticationType enums.AuthenticationType `type:"string" required:"true" enum:"true"`

	// The date and time the user was created in the user pool.
	CreatedTime *time.Time `type:"timestamp"`

	// Specifies whether the user in the user pool is enabled.
	Enabled *bool `type:"boolean"`

	// The first name, or given name, of the user.
	FirstName *string `type:"string" sensitive:"true"`

	// The last name, or surname, of the user.
	LastName *string `type:"string" sensitive:"true"`

	// The status of the user in the user pool. The status can be one of the following:
	//
	//    * UNCONFIRMED – The user is created but not confirmed.
	//
	//    * CONFIRMED – The user is confirmed.
	//
	//    * ARCHIVED – The user is no longer active.
	//
	//    * COMPROMISED – The user is disabled because of a potential security
	//    threat.
	//
	//    * UNKNOWN – The user status is not known.
	Status *string `min:"1" type:"string"`

	// The email address of the user.
	//
	// Users' email addresses are case-sensitive.
	UserName *string `min:"1" type:"string" sensitive:"true"`
}

// String returns the string representation
func (s User) String() string {
	return awsutil.Prettify(s)
}

// Describes an action and whether the action is enabled or disabled for users
// during their streaming sessions.
type UserSetting struct {
	_ struct{} `type:"structure"`

	// The action that is enabled or disabled.
	//
	// Action is a required field
	Action enums.Action `type:"string" required:"true" enum:"true"`

	// Indicates whether the action is enabled or disabled.
	//
	// Permission is a required field
	Permission enums.Permission `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s UserSetting) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UserSetting) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UserSetting"}
	if len(s.Action) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Action"))
	}
	if len(s.Permission) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Permission"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes a user in the user pool and the associated stack.
type UserStackAssociation struct {
	_ struct{} `type:"structure"`

	// The authentication type for the user.
	//
	// AuthenticationType is a required field
	AuthenticationType enums.AuthenticationType `type:"string" required:"true" enum:"true"`

	// Specifies whether a welcome email is sent to a user after the user is created
	// in the user pool.
	SendEmailNotification *bool `type:"boolean"`

	// The name of the stack that is associated with the user.
	//
	// StackName is a required field
	StackName *string `min:"1" type:"string" required:"true"`

	// The email address of the user who is associated with the stack.
	//
	// Users' email addresses are case-sensitive.
	//
	// UserName is a required field
	UserName *string `min:"1" type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s UserStackAssociation) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UserStackAssociation) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UserStackAssociation"}
	if len(s.AuthenticationType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("AuthenticationType"))
	}

	if s.StackName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackName"))
	}
	if s.StackName != nil && len(*s.StackName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StackName", 1))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes the error that is returned when a user can’t be associated with
// or disassociated from a stack.
type UserStackAssociationError struct {
	_ struct{} `type:"structure"`

	// The error code for the error that is returned when a user can’t be associated
	// with or disassociated from a stack.
	ErrorCode enums.UserStackAssociationErrorCode `type:"string" enum:"true"`

	// The error message for the error that is returned when a user can’t be associated
	// with or disassociated from a stack.
	ErrorMessage *string `min:"1" type:"string"`

	// Information about the user and associated stack.
	UserStackAssociation *UserStackAssociation `type:"structure"`
}

// String returns the string representation
func (s UserStackAssociationError) String() string {
	return awsutil.Prettify(s)
}

// Describes VPC configuration information for fleets and image builders.
type VpcConfig struct {
	_ struct{} `type:"structure"`

	// The identifiers of the security groups for the fleet or image builder.
	SecurityGroupIds []string `type:"list"`

	// The identifiers of the subnets to which a network interface is attached from
	// the fleet instance or image builder instance. Fleet instances use one or
	// more subnets. Image builder instances use one subnet.
	SubnetIds []string `type:"list"`
}

// String returns the string representation
func (s VpcConfig) String() string {
	return awsutil.Prettify(s)
}
