// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package enums

type ArtifactType string

// Enum values for ArtifactType
const (
	ArtifactTypeTest ArtifactType = "TEST"
)

func (enum ArtifactType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ArtifactType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DomainStatus string

// Enum values for DomainStatus
const (
	DomainStatusPendingVerification   DomainStatus = "PENDING_VERIFICATION"
	DomainStatusInProgress            DomainStatus = "IN_PROGRESS"
	DomainStatusAvailable             DomainStatus = "AVAILABLE"
	DomainStatusPendingDeployment     DomainStatus = "PENDING_DEPLOYMENT"
	DomainStatusFailed                DomainStatus = "FAILED"
	DomainStatusCreating              DomainStatus = "CREATING"
	DomainStatusRequestingCertificate DomainStatus = "REQUESTING_CERTIFICATE"
	DomainStatusUpdating              DomainStatus = "UPDATING"
)

func (enum DomainStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DomainStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type JobStatus string

// Enum values for JobStatus
const (
	JobStatusPending      JobStatus = "PENDING"
	JobStatusProvisioning JobStatus = "PROVISIONING"
	JobStatusRunning      JobStatus = "RUNNING"
	JobStatusFailed       JobStatus = "FAILED"
	JobStatusSucceed      JobStatus = "SUCCEED"
	JobStatusCancelling   JobStatus = "CANCELLING"
	JobStatusCancelled    JobStatus = "CANCELLED"
)

func (enum JobStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum JobStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type JobType string

// Enum values for JobType
const (
	JobTypeRelease JobType = "RELEASE"
	JobTypeRetry   JobType = "RETRY"
	JobTypeManual  JobType = "MANUAL"
	JobTypeWebHook JobType = "WEB_HOOK"
)

func (enum JobType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum JobType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Platform string

// Enum values for Platform
const (
	PlatformWeb Platform = "WEB"
)

func (enum Platform) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Platform) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Stage string

// Enum values for Stage
const (
	StageProduction   Stage = "PRODUCTION"
	StageBeta         Stage = "BETA"
	StageDevelopment  Stage = "DEVELOPMENT"
	StageExperimental Stage = "EXPERIMENTAL"
	StagePullRequest  Stage = "PULL_REQUEST"
)

func (enum Stage) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Stage) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
