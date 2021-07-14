// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Two or more batch entries in the request have the same Id.
type BatchEntryIdsNotDistinct struct {
	Message *string
}

func (e *BatchEntryIdsNotDistinct) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BatchEntryIdsNotDistinct) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BatchEntryIdsNotDistinct) ErrorCode() string {
	return "AWS.SimpleQueueService.BatchEntryIdsNotDistinct"
}
func (e *BatchEntryIdsNotDistinct) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The length of all the messages put together is more than the limit.
type BatchRequestTooLong struct {
	Message *string
}

func (e *BatchRequestTooLong) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BatchRequestTooLong) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BatchRequestTooLong) ErrorCode() string             { return "AWS.SimpleQueueService.BatchRequestTooLong" }
func (e *BatchRequestTooLong) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The batch request doesn't contain any entries.
type EmptyBatchRequest struct {
	Message *string
}

func (e *EmptyBatchRequest) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EmptyBatchRequest) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EmptyBatchRequest) ErrorCode() string             { return "AWS.SimpleQueueService.EmptyBatchRequest" }
func (e *EmptyBatchRequest) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified attribute doesn't exist.
type InvalidAttributeName struct {
	Message *string
}

func (e *InvalidAttributeName) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidAttributeName) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidAttributeName) ErrorCode() string             { return "InvalidAttributeName" }
func (e *InvalidAttributeName) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The Id of a batch entry in a batch request doesn't abide by the specification.
type InvalidBatchEntryId struct {
	Message *string
}

func (e *InvalidBatchEntryId) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidBatchEntryId) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidBatchEntryId) ErrorCode() string             { return "AWS.SimpleQueueService.InvalidBatchEntryId" }
func (e *InvalidBatchEntryId) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified receipt handle isn't valid for the current version.
type InvalidIdFormat struct {
	Message *string
}

func (e *InvalidIdFormat) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidIdFormat) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidIdFormat) ErrorCode() string             { return "InvalidIdFormat" }
func (e *InvalidIdFormat) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The message contains characters outside the allowed set.
type InvalidMessageContents struct {
	Message *string
}

func (e *InvalidMessageContents) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidMessageContents) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidMessageContents) ErrorCode() string             { return "InvalidMessageContents" }
func (e *InvalidMessageContents) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified message isn't in flight.
type MessageNotInflight struct {
	Message *string
}

func (e *MessageNotInflight) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MessageNotInflight) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MessageNotInflight) ErrorCode() string             { return "AWS.SimpleQueueService.MessageNotInflight" }
func (e *MessageNotInflight) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified action violates a limit. For example, ReceiveMessage returns this
// error if the maximum number of inflight messages is reached and AddPermission
// returns this error if the maximum number of permissions for the queue is
// reached.
type OverLimit struct {
	Message *string
}

func (e *OverLimit) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OverLimit) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OverLimit) ErrorCode() string             { return "OverLimit" }
func (e *OverLimit) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the specified queue previously received a PurgeQueue request
// within the last 60 seconds (the time it can take to delete the messages in the
// queue).
type PurgeQueueInProgress struct {
	Message *string
}

func (e *PurgeQueueInProgress) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PurgeQueueInProgress) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PurgeQueueInProgress) ErrorCode() string {
	return "AWS.SimpleQueueService.PurgeQueueInProgress"
}
func (e *PurgeQueueInProgress) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You must wait 60 seconds after deleting a queue before you can create another
// queue with the same name.
type QueueDeletedRecently struct {
	Message *string
}

func (e *QueueDeletedRecently) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *QueueDeletedRecently) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *QueueDeletedRecently) ErrorCode() string {
	return "AWS.SimpleQueueService.QueueDeletedRecently"
}
func (e *QueueDeletedRecently) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified queue doesn't exist.
type QueueDoesNotExist struct {
	Message *string
}

func (e *QueueDoesNotExist) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *QueueDoesNotExist) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *QueueDoesNotExist) ErrorCode() string             { return "AWS.SimpleQueueService.NonExistentQueue" }
func (e *QueueDoesNotExist) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A queue with this name already exists. Amazon SQS returns this error only if the
// request includes attributes whose values differ from those of the existing
// queue.
type QueueNameExists struct {
	Message *string
}

func (e *QueueNameExists) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *QueueNameExists) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *QueueNameExists) ErrorCode() string             { return "QueueAlreadyExists" }
func (e *QueueNameExists) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified receipt handle isn't valid.
type ReceiptHandleIsInvalid struct {
	Message *string
}

func (e *ReceiptHandleIsInvalid) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ReceiptHandleIsInvalid) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ReceiptHandleIsInvalid) ErrorCode() string             { return "ReceiptHandleIsInvalid" }
func (e *ReceiptHandleIsInvalid) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The batch request contains more entries than permissible.
type TooManyEntriesInBatchRequest struct {
	Message *string
}

func (e *TooManyEntriesInBatchRequest) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyEntriesInBatchRequest) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyEntriesInBatchRequest) ErrorCode() string {
	return "AWS.SimpleQueueService.TooManyEntriesInBatchRequest"
}
func (e *TooManyEntriesInBatchRequest) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Error code 400. Unsupported operation.
type UnsupportedOperation struct {
	Message *string
}

func (e *UnsupportedOperation) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedOperation) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedOperation) ErrorCode() string {
	return "AWS.SimpleQueueService.UnsupportedOperation"
}
func (e *UnsupportedOperation) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
