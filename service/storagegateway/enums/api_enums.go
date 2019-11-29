// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package enums

type ActiveDirectoryStatus string

// Enum values for ActiveDirectoryStatus
const (
	ActiveDirectoryStatusAccessDenied ActiveDirectoryStatus = "ACCESS_DENIED"
	ActiveDirectoryStatusDetached     ActiveDirectoryStatus = "DETACHED"
	ActiveDirectoryStatusJoined       ActiveDirectoryStatus = "JOINED"
	ActiveDirectoryStatusJoining      ActiveDirectoryStatus = "JOINING"
	ActiveDirectoryStatusNetworkError ActiveDirectoryStatus = "NETWORK_ERROR"
	ActiveDirectoryStatusTimeout      ActiveDirectoryStatus = "TIMEOUT"
	ActiveDirectoryStatusUnknownError ActiveDirectoryStatus = "UNKNOWN_ERROR"
)

func (enum ActiveDirectoryStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ActiveDirectoryStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type AvailabilityMonitorTestStatus string

// Enum values for AvailabilityMonitorTestStatus
const (
	AvailabilityMonitorTestStatusComplete AvailabilityMonitorTestStatus = "COMPLETE"
	AvailabilityMonitorTestStatusFailed   AvailabilityMonitorTestStatus = "FAILED"
	AvailabilityMonitorTestStatusPending  AvailabilityMonitorTestStatus = "PENDING"
)

func (enum AvailabilityMonitorTestStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AvailabilityMonitorTestStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ErrorCode string

// Enum values for ErrorCode
const (
	ErrorCodeActivationKeyExpired              ErrorCode = "ActivationKeyExpired"
	ErrorCodeActivationKeyInvalid              ErrorCode = "ActivationKeyInvalid"
	ErrorCodeActivationKeyNotFound             ErrorCode = "ActivationKeyNotFound"
	ErrorCodeGatewayInternalError              ErrorCode = "GatewayInternalError"
	ErrorCodeGatewayNotConnected               ErrorCode = "GatewayNotConnected"
	ErrorCodeGatewayNotFound                   ErrorCode = "GatewayNotFound"
	ErrorCodeGatewayProxyNetworkConnectionBusy ErrorCode = "GatewayProxyNetworkConnectionBusy"
	ErrorCodeAuthenticationFailure             ErrorCode = "AuthenticationFailure"
	ErrorCodeBandwidthThrottleScheduleNotFound ErrorCode = "BandwidthThrottleScheduleNotFound"
	ErrorCodeBlocked                           ErrorCode = "Blocked"
	ErrorCodeCannotExportSnapshot              ErrorCode = "CannotExportSnapshot"
	ErrorCodeChapCredentialNotFound            ErrorCode = "ChapCredentialNotFound"
	ErrorCodeDiskAlreadyAllocated              ErrorCode = "DiskAlreadyAllocated"
	ErrorCodeDiskDoesNotExist                  ErrorCode = "DiskDoesNotExist"
	ErrorCodeDiskSizeGreaterThanVolumeMaxSize  ErrorCode = "DiskSizeGreaterThanVolumeMaxSize"
	ErrorCodeDiskSizeLessThanVolumeSize        ErrorCode = "DiskSizeLessThanVolumeSize"
	ErrorCodeDiskSizeNotGigAligned             ErrorCode = "DiskSizeNotGigAligned"
	ErrorCodeDuplicateCertificateInfo          ErrorCode = "DuplicateCertificateInfo"
	ErrorCodeDuplicateSchedule                 ErrorCode = "DuplicateSchedule"
	ErrorCodeEndpointNotFound                  ErrorCode = "EndpointNotFound"
	ErrorCodeIamnotSupported                   ErrorCode = "IAMNotSupported"
	ErrorCodeInitiatorInvalid                  ErrorCode = "InitiatorInvalid"
	ErrorCodeInitiatorNotFound                 ErrorCode = "InitiatorNotFound"
	ErrorCodeInternalError                     ErrorCode = "InternalError"
	ErrorCodeInvalidGateway                    ErrorCode = "InvalidGateway"
	ErrorCodeInvalidEndpoint                   ErrorCode = "InvalidEndpoint"
	ErrorCodeInvalidParameters                 ErrorCode = "InvalidParameters"
	ErrorCodeInvalidSchedule                   ErrorCode = "InvalidSchedule"
	ErrorCodeLocalStorageLimitExceeded         ErrorCode = "LocalStorageLimitExceeded"
	ErrorCodeLunAlreadyAllocated               ErrorCode = "LunAlreadyAllocated "
	ErrorCodeLunInvalid                        ErrorCode = "LunInvalid"
	ErrorCodeJoinDomainInProgress              ErrorCode = "JoinDomainInProgress"
	ErrorCodeMaximumContentLengthExceeded      ErrorCode = "MaximumContentLengthExceeded"
	ErrorCodeMaximumTapeCartridgeCountExceeded ErrorCode = "MaximumTapeCartridgeCountExceeded"
	ErrorCodeMaximumVolumeCountExceeded        ErrorCode = "MaximumVolumeCountExceeded"
	ErrorCodeNetworkConfigurationChanged       ErrorCode = "NetworkConfigurationChanged"
	ErrorCodeNoDisksAvailable                  ErrorCode = "NoDisksAvailable"
	ErrorCodeNotImplemented                    ErrorCode = "NotImplemented"
	ErrorCodeNotSupported                      ErrorCode = "NotSupported"
	ErrorCodeOperationAborted                  ErrorCode = "OperationAborted"
	ErrorCodeOutdatedGateway                   ErrorCode = "OutdatedGateway"
	ErrorCodeParametersNotImplemented          ErrorCode = "ParametersNotImplemented"
	ErrorCodeRegionInvalid                     ErrorCode = "RegionInvalid"
	ErrorCodeRequestTimeout                    ErrorCode = "RequestTimeout"
	ErrorCodeServiceUnavailable                ErrorCode = "ServiceUnavailable"
	ErrorCodeSnapshotDeleted                   ErrorCode = "SnapshotDeleted"
	ErrorCodeSnapshotIdInvalid                 ErrorCode = "SnapshotIdInvalid"
	ErrorCodeSnapshotInProgress                ErrorCode = "SnapshotInProgress"
	ErrorCodeSnapshotNotFound                  ErrorCode = "SnapshotNotFound"
	ErrorCodeSnapshotScheduleNotFound          ErrorCode = "SnapshotScheduleNotFound"
	ErrorCodeStagingAreaFull                   ErrorCode = "StagingAreaFull"
	ErrorCodeStorageFailure                    ErrorCode = "StorageFailure"
	ErrorCodeTapeCartridgeNotFound             ErrorCode = "TapeCartridgeNotFound"
	ErrorCodeTargetAlreadyExists               ErrorCode = "TargetAlreadyExists"
	ErrorCodeTargetInvalid                     ErrorCode = "TargetInvalid"
	ErrorCodeTargetNotFound                    ErrorCode = "TargetNotFound"
	ErrorCodeUnauthorizedOperation             ErrorCode = "UnauthorizedOperation"
	ErrorCodeVolumeAlreadyExists               ErrorCode = "VolumeAlreadyExists"
	ErrorCodeVolumeIdInvalid                   ErrorCode = "VolumeIdInvalid"
	ErrorCodeVolumeInUse                       ErrorCode = "VolumeInUse"
	ErrorCodeVolumeNotFound                    ErrorCode = "VolumeNotFound"
	ErrorCodeVolumeNotReady                    ErrorCode = "VolumeNotReady"
)

func (enum ErrorCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ErrorCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The type of the file share.
type FileShareType string

// Enum values for FileShareType
const (
	FileShareTypeNfs FileShareType = "NFS"
	FileShareTypeSmb FileShareType = "SMB"
)

func (enum FileShareType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum FileShareType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type HostEnvironment string

// Enum values for HostEnvironment
const (
	HostEnvironmentVmware HostEnvironment = "VMWARE"
	HostEnvironmentHyperV HostEnvironment = "HYPER-V"
	HostEnvironmentEc2    HostEnvironment = "EC2"
	HostEnvironmentOther  HostEnvironment = "OTHER"
)

func (enum HostEnvironment) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum HostEnvironment) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// A value that sets the access control list permission for objects in the S3
// bucket that a file gateway puts objects into. The default value is "private".
type ObjectACL string

// Enum values for ObjectACL
const (
	ObjectACLPrivate                ObjectACL = "private"
	ObjectACLPublicRead             ObjectACL = "public-read"
	ObjectACLPublicReadWrite        ObjectACL = "public-read-write"
	ObjectACLAuthenticatedRead      ObjectACL = "authenticated-read"
	ObjectACLBucketOwnerRead        ObjectACL = "bucket-owner-read"
	ObjectACLBucketOwnerFullControl ObjectACL = "bucket-owner-full-control"
	ObjectACLAwsExecRead            ObjectACL = "aws-exec-read"
)

func (enum ObjectACL) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ObjectACL) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SMBSecurityStrategy string

// Enum values for SMBSecurityStrategy
const (
	SMBSecurityStrategyClientSpecified     SMBSecurityStrategy = "ClientSpecified"
	SMBSecurityStrategyMandatorySigning    SMBSecurityStrategy = "MandatorySigning"
	SMBSecurityStrategyMandatoryEncryption SMBSecurityStrategy = "MandatoryEncryption"
)

func (enum SMBSecurityStrategy) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SMBSecurityStrategy) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}