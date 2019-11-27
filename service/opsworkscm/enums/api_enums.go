// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package enums

type BackupStatus string

// Enum values for BackupStatus
const (
	BackupStatusInProgress BackupStatus = "IN_PROGRESS"
	BackupStatusOk         BackupStatus = "OK"
	BackupStatusFailed     BackupStatus = "FAILED"
	BackupStatusDeleting   BackupStatus = "DELETING"
)

func (enum BackupStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum BackupStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type BackupType string

// Enum values for BackupType
const (
	BackupTypeAutomated BackupType = "AUTOMATED"
	BackupTypeManual    BackupType = "MANUAL"
)

func (enum BackupType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum BackupType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type MaintenanceStatus string

// Enum values for MaintenanceStatus
const (
	MaintenanceStatusSuccess MaintenanceStatus = "SUCCESS"
	MaintenanceStatusFailed  MaintenanceStatus = "FAILED"
)

func (enum MaintenanceStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum MaintenanceStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The status of the association or disassociation request.
//
// Possible values:
//
//    * SUCCESS: The association or disassociation succeeded.
//
//    * FAILED: The association or disassociation failed.
//
//    * IN_PROGRESS: The association or disassociation is still in progress.
type NodeAssociationStatus string

// Enum values for NodeAssociationStatus
const (
	NodeAssociationStatusSuccess    NodeAssociationStatus = "SUCCESS"
	NodeAssociationStatusFailed     NodeAssociationStatus = "FAILED"
	NodeAssociationStatusInProgress NodeAssociationStatus = "IN_PROGRESS"
)

func (enum NodeAssociationStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum NodeAssociationStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ServerStatus string

// Enum values for ServerStatus
const (
	ServerStatusBackingUp        ServerStatus = "BACKING_UP"
	ServerStatusConnectionLost   ServerStatus = "CONNECTION_LOST"
	ServerStatusCreating         ServerStatus = "CREATING"
	ServerStatusDeleting         ServerStatus = "DELETING"
	ServerStatusModifying        ServerStatus = "MODIFYING"
	ServerStatusFailed           ServerStatus = "FAILED"
	ServerStatusHealthy          ServerStatus = "HEALTHY"
	ServerStatusRunning          ServerStatus = "RUNNING"
	ServerStatusRestoring        ServerStatus = "RESTORING"
	ServerStatusSetup            ServerStatus = "SETUP"
	ServerStatusUnderMaintenance ServerStatus = "UNDER_MAINTENANCE"
	ServerStatusUnhealthy        ServerStatus = "UNHEALTHY"
	ServerStatusTerminated       ServerStatus = "TERMINATED"
)

func (enum ServerStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ServerStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
