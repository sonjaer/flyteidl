// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: flyteidl/admin/execution.proto

package admin

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"

	core "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}

	_ = core.WorkflowExecution_Phase(0)
)

// Validate checks the field values on ExecutionCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ExecutionCreateRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Project

	// no validation rules for Domain

	// no validation rules for Name

	if v, ok := interface{}(m.GetSpec()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionCreateRequestValidationError{
				field:  "Spec",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetInputs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionCreateRequestValidationError{
				field:  "Inputs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ExecutionCreateRequestValidationError is the validation error returned by
// ExecutionCreateRequest.Validate if the designated constraints aren't met.
type ExecutionCreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExecutionCreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExecutionCreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExecutionCreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExecutionCreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExecutionCreateRequestValidationError) ErrorName() string {
	return "ExecutionCreateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ExecutionCreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExecutionCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExecutionCreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExecutionCreateRequestValidationError{}

// Validate checks the field values on ExecutionRelaunchRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ExecutionRelaunchRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionRelaunchRequestValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Name

	return nil
}

// ExecutionRelaunchRequestValidationError is the validation error returned by
// ExecutionRelaunchRequest.Validate if the designated constraints aren't met.
type ExecutionRelaunchRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExecutionRelaunchRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExecutionRelaunchRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExecutionRelaunchRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExecutionRelaunchRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExecutionRelaunchRequestValidationError) ErrorName() string {
	return "ExecutionRelaunchRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ExecutionRelaunchRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExecutionRelaunchRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExecutionRelaunchRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExecutionRelaunchRequestValidationError{}

// Validate checks the field values on ExecutionCreateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ExecutionCreateResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionCreateResponseValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ExecutionCreateResponseValidationError is the validation error returned by
// ExecutionCreateResponse.Validate if the designated constraints aren't met.
type ExecutionCreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExecutionCreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExecutionCreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExecutionCreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExecutionCreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExecutionCreateResponseValidationError) ErrorName() string {
	return "ExecutionCreateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ExecutionCreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExecutionCreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExecutionCreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExecutionCreateResponseValidationError{}

// Validate checks the field values on WorkflowExecutionGetRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *WorkflowExecutionGetRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkflowExecutionGetRequestValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// WorkflowExecutionGetRequestValidationError is the validation error returned
// by WorkflowExecutionGetRequest.Validate if the designated constraints
// aren't met.
type WorkflowExecutionGetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkflowExecutionGetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkflowExecutionGetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkflowExecutionGetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkflowExecutionGetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkflowExecutionGetRequestValidationError) ErrorName() string {
	return "WorkflowExecutionGetRequestValidationError"
}

// Error satisfies the builtin error interface
func (e WorkflowExecutionGetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkflowExecutionGetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkflowExecutionGetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkflowExecutionGetRequestValidationError{}

// Validate checks the field values on Execution with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Execution) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetSpec()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionValidationError{
				field:  "Spec",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetClosure()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionValidationError{
				field:  "Closure",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ExecutionValidationError is the validation error returned by
// Execution.Validate if the designated constraints aren't met.
type ExecutionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExecutionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExecutionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExecutionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExecutionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExecutionValidationError) ErrorName() string { return "ExecutionValidationError" }

// Error satisfies the builtin error interface
func (e ExecutionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExecution.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExecutionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExecutionValidationError{}

// Validate checks the field values on ExecutionList with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ExecutionList) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetExecutions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ExecutionListValidationError{
					field:  fmt.Sprintf("Executions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Token

	return nil
}

// ExecutionListValidationError is the validation error returned by
// ExecutionList.Validate if the designated constraints aren't met.
type ExecutionListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExecutionListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExecutionListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExecutionListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExecutionListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExecutionListValidationError) ErrorName() string { return "ExecutionListValidationError" }

// Error satisfies the builtin error interface
func (e ExecutionListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExecutionList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExecutionListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExecutionListValidationError{}

// Validate checks the field values on LiteralMapBlob with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *LiteralMapBlob) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Data.(type) {

	case *LiteralMapBlob_Values:

		if v, ok := interface{}(m.GetValues()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LiteralMapBlobValidationError{
					field:  "Values",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LiteralMapBlob_Uri:
		// no validation rules for Uri

	}

	return nil
}

// LiteralMapBlobValidationError is the validation error returned by
// LiteralMapBlob.Validate if the designated constraints aren't met.
type LiteralMapBlobValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LiteralMapBlobValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LiteralMapBlobValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LiteralMapBlobValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LiteralMapBlobValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LiteralMapBlobValidationError) ErrorName() string { return "LiteralMapBlobValidationError" }

// Error satisfies the builtin error interface
func (e LiteralMapBlobValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLiteralMapBlob.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LiteralMapBlobValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LiteralMapBlobValidationError{}

// Validate checks the field values on ExecutionClosure with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ExecutionClosure) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetComputedInputs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionClosureValidationError{
				field:  "ComputedInputs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Phase

	if v, ok := interface{}(m.GetStartedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionClosureValidationError{
				field:  "StartedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDuration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionClosureValidationError{
				field:  "Duration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionClosureValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionClosureValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetNotifications() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ExecutionClosureValidationError{
					field:  fmt.Sprintf("Notifications[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetWorkflowId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionClosureValidationError{
				field:  "WorkflowId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for InputsUri

	switch m.OutputResult.(type) {

	case *ExecutionClosure_Outputs:

		if v, ok := interface{}(m.GetOutputs()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ExecutionClosureValidationError{
					field:  "Outputs",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ExecutionClosure_Error:

		if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ExecutionClosureValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ExecutionClosure_AbortCause:
		// no validation rules for AbortCause

	}

	return nil
}

// ExecutionClosureValidationError is the validation error returned by
// ExecutionClosure.Validate if the designated constraints aren't met.
type ExecutionClosureValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExecutionClosureValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExecutionClosureValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExecutionClosureValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExecutionClosureValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExecutionClosureValidationError) ErrorName() string { return "ExecutionClosureValidationError" }

// Error satisfies the builtin error interface
func (e ExecutionClosureValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExecutionClosure.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExecutionClosureValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExecutionClosureValidationError{}

// Validate checks the field values on ExecutionMetadata with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ExecutionMetadata) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Mode

	// no validation rules for Principal

	// no validation rules for Nesting

	if v, ok := interface{}(m.GetScheduledAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionMetadataValidationError{
				field:  "ScheduledAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetParentNodeExecution()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionMetadataValidationError{
				field:  "ParentNodeExecution",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetReferenceExecution()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionMetadataValidationError{
				field:  "ReferenceExecution",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ExecutionMetadataValidationError is the validation error returned by
// ExecutionMetadata.Validate if the designated constraints aren't met.
type ExecutionMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExecutionMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExecutionMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExecutionMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExecutionMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExecutionMetadataValidationError) ErrorName() string {
	return "ExecutionMetadataValidationError"
}

// Error satisfies the builtin error interface
func (e ExecutionMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExecutionMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExecutionMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExecutionMetadataValidationError{}

// Validate checks the field values on NotificationList with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *NotificationList) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetNotifications() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NotificationListValidationError{
					field:  fmt.Sprintf("Notifications[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// NotificationListValidationError is the validation error returned by
// NotificationList.Validate if the designated constraints aren't met.
type NotificationListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NotificationListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NotificationListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NotificationListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NotificationListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NotificationListValidationError) ErrorName() string { return "NotificationListValidationError" }

// Error satisfies the builtin error interface
func (e NotificationListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNotificationList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NotificationListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NotificationListValidationError{}

// Validate checks the field values on ExecutionSpec with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ExecutionSpec) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetLaunchPlan()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionSpecValidationError{
				field:  "LaunchPlan",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetInputs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionSpecValidationError{
				field:  "Inputs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionSpecValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetLabels()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionSpecValidationError{
				field:  "Labels",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAnnotations()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionSpecValidationError{
				field:  "Annotations",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.NotificationOverrides.(type) {

	case *ExecutionSpec_Notifications:

		if v, ok := interface{}(m.GetNotifications()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ExecutionSpecValidationError{
					field:  "Notifications",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ExecutionSpec_DisableAll:
		// no validation rules for DisableAll

	}

	return nil
}

// ExecutionSpecValidationError is the validation error returned by
// ExecutionSpec.Validate if the designated constraints aren't met.
type ExecutionSpecValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExecutionSpecValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExecutionSpecValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExecutionSpecValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExecutionSpecValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExecutionSpecValidationError) ErrorName() string { return "ExecutionSpecValidationError" }

// Error satisfies the builtin error interface
func (e ExecutionSpecValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExecutionSpec.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExecutionSpecValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExecutionSpecValidationError{}

// Validate checks the field values on ExecutionTerminateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ExecutionTerminateRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExecutionTerminateRequestValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Cause

	return nil
}

// ExecutionTerminateRequestValidationError is the validation error returned by
// ExecutionTerminateRequest.Validate if the designated constraints aren't met.
type ExecutionTerminateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExecutionTerminateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExecutionTerminateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExecutionTerminateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExecutionTerminateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExecutionTerminateRequestValidationError) ErrorName() string {
	return "ExecutionTerminateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ExecutionTerminateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExecutionTerminateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExecutionTerminateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExecutionTerminateRequestValidationError{}

// Validate checks the field values on ExecutionTerminateResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ExecutionTerminateResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ExecutionTerminateResponseValidationError is the validation error returned
// by ExecutionTerminateResponse.Validate if the designated constraints aren't met.
type ExecutionTerminateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExecutionTerminateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExecutionTerminateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExecutionTerminateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExecutionTerminateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExecutionTerminateResponseValidationError) ErrorName() string {
	return "ExecutionTerminateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ExecutionTerminateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExecutionTerminateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExecutionTerminateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExecutionTerminateResponseValidationError{}

// Validate checks the field values on WorkflowExecutionGetDataRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *WorkflowExecutionGetDataRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkflowExecutionGetDataRequestValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// WorkflowExecutionGetDataRequestValidationError is the validation error
// returned by WorkflowExecutionGetDataRequest.Validate if the designated
// constraints aren't met.
type WorkflowExecutionGetDataRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkflowExecutionGetDataRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkflowExecutionGetDataRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkflowExecutionGetDataRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkflowExecutionGetDataRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkflowExecutionGetDataRequestValidationError) ErrorName() string {
	return "WorkflowExecutionGetDataRequestValidationError"
}

// Error satisfies the builtin error interface
func (e WorkflowExecutionGetDataRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkflowExecutionGetDataRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkflowExecutionGetDataRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkflowExecutionGetDataRequestValidationError{}

// Validate checks the field values on WorkflowExecutionGetDataResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *WorkflowExecutionGetDataResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetOutputs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkflowExecutionGetDataResponseValidationError{
				field:  "Outputs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetInputs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkflowExecutionGetDataResponseValidationError{
				field:  "Inputs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// WorkflowExecutionGetDataResponseValidationError is the validation error
// returned by WorkflowExecutionGetDataResponse.Validate if the designated
// constraints aren't met.
type WorkflowExecutionGetDataResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkflowExecutionGetDataResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkflowExecutionGetDataResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkflowExecutionGetDataResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkflowExecutionGetDataResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkflowExecutionGetDataResponseValidationError) ErrorName() string {
	return "WorkflowExecutionGetDataResponseValidationError"
}

// Error satisfies the builtin error interface
func (e WorkflowExecutionGetDataResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkflowExecutionGetDataResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkflowExecutionGetDataResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkflowExecutionGetDataResponseValidationError{}
