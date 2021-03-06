package events

import (
	"context"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/lyft/flyteidl/clients/go/events/errors"
	"github.com/lyft/flyteidl/gen/pb-go/flyteidl/event"
	"github.com/lyft/flytestdlib/promutils"
	"github.com/lyft/flytestdlib/promutils/labeled"
)

type recordingMetrics struct {
	EventRecordingFailure           labeled.StopWatch
	EventRecordingSuccess           labeled.StopWatch
	EventRecordingAlreadyExists     labeled.Counter
	EventRecordingExecutionNotFound labeled.Counter
	EventRecordingResourceExhausted labeled.Counter
	EventRecordingEventSinkError    labeled.Counter
	EventRecordingInvalidArgument   labeled.Counter
}

// EventRecorder records workflow, node and task events to the eventSink it is configured with.
type eventRecorder struct {
	eventSink EventSink
	metrics   *recordingMetrics
}

func constructEventRecorder(eventSink EventSink, scope promutils.Scope) *eventRecorder {
	recordingScope := scope.NewSubScope("event_recording")
	return &eventRecorder{
		eventSink: eventSink,
		metrics: &recordingMetrics{
			EventRecordingFailure:           labeled.NewStopWatch("failure_duration", "The time it took the failed event recording to occur", time.Millisecond, recordingScope),
			EventRecordingSuccess:           labeled.NewStopWatch("success_duration", "The time it took for a successful event recording to occur", time.Millisecond, recordingScope),
			EventRecordingAlreadyExists:     labeled.NewCounter("already_exists", "The count that a recorded event already exists", recordingScope),
			EventRecordingResourceExhausted: labeled.NewCounter("resource_exhausted", "The count that recording events was throttled", recordingScope),
			EventRecordingInvalidArgument:   labeled.NewCounter("invalid_argument", "The count for invalid argument errors", recordingScope),
			EventRecordingEventSinkError:    labeled.NewCounter("unexpected_err", "The count of event recording failures for unexpected reasons", recordingScope),
		},
	}
}

func (r *eventRecorder) sinkEvent(ctx context.Context, event proto.Message) error {
	startTime := time.Now()

	err := r.eventSink.Sink(ctx, event)
	if errors.IsResourceExhausted(err) {
		r.metrics.EventRecordingResourceExhausted.Inc(ctx)
	}

	if err != nil {
		r.metrics.EventRecordingFailure.Observe(ctx, startTime, time.Now())
		return err
	}

	r.metrics.EventRecordingSuccess.Observe(ctx, startTime, time.Now())
	return nil
}

func (r *eventRecorder) RecordWorkflowEvent(ctx context.Context, event *event.WorkflowExecutionEvent) error {
	return r.sinkEvent(ctx, event)
}

func (r *eventRecorder) RecordNodeEvent(ctx context.Context, event *event.NodeExecutionEvent) error {
	return r.sinkEvent(ctx, event)
}

func (r *eventRecorder) RecordTaskEvent(ctx context.Context, event *event.TaskExecutionEvent) error {
	return r.sinkEvent(ctx, event)
}

// Construct a new Workflow Event Recorder
func NewWorkflowEventRecorder(eventSink EventSink, scope promutils.Scope) WorkflowEventRecorder {
	return constructEventRecorder(eventSink, scope)
}

// Construct a new Node Event Recorder
func NewNodeEventRecorder(eventSink EventSink, scope promutils.Scope) NodeEventRecorder {
	return constructEventRecorder(eventSink, scope)
}

// Construct a new Task Event Recorder
func NewTaskEventRecorder(eventSink EventSink, scope promutils.Scope) TaskEventRecorder {
	return constructEventRecorder(eventSink, scope)
}
