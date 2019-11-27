// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package enums

type ApplicationStatus string

// Enum values for ApplicationStatus
const (
	ApplicationStatusNotStarted ApplicationStatus = "NOT_STARTED"
	ApplicationStatusInProgress ApplicationStatus = "IN_PROGRESS"
	ApplicationStatusCompleted  ApplicationStatus = "COMPLETED"
)

func (enum ApplicationStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ApplicationStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ResourceAttributeType string

// Enum values for ResourceAttributeType
const (
	ResourceAttributeTypeIpv4Address              ResourceAttributeType = "IPV4_ADDRESS"
	ResourceAttributeTypeIpv6Address              ResourceAttributeType = "IPV6_ADDRESS"
	ResourceAttributeTypeMacAddress               ResourceAttributeType = "MAC_ADDRESS"
	ResourceAttributeTypeFqdn                     ResourceAttributeType = "FQDN"
	ResourceAttributeTypeVmManagerId              ResourceAttributeType = "VM_MANAGER_ID"
	ResourceAttributeTypeVmManagedObjectReference ResourceAttributeType = "VM_MANAGED_OBJECT_REFERENCE"
	ResourceAttributeTypeVmName                   ResourceAttributeType = "VM_NAME"
	ResourceAttributeTypeVmPath                   ResourceAttributeType = "VM_PATH"
	ResourceAttributeTypeBiosId                   ResourceAttributeType = "BIOS_ID"
	ResourceAttributeTypeMotherboardSerialNumber  ResourceAttributeType = "MOTHERBOARD_SERIAL_NUMBER"
)

func (enum ResourceAttributeType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ResourceAttributeType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Status string

// Enum values for Status
const (
	StatusNotStarted Status = "NOT_STARTED"
	StatusInProgress Status = "IN_PROGRESS"
	StatusFailed     Status = "FAILED"
	StatusCompleted  Status = "COMPLETED"
)

func (enum Status) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Status) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
