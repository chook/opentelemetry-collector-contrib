# Span Context

The Span Context is a Context implementation for [pdata Spans](https://github.com/open-telemetry/opentelemetry-collector/tree/main/pdata/ptrace), the collector's internal representation for OTLP span data.  This Context should be used when interacted with OTLP spans.

## Paths
In general, the Span Context supports accessing pdata using the field names from the [traces proto](https://github.com/open-telemetry/opentelemetry-proto/blob/main/opentelemetry/proto/trace/v1/trace.proto).  All integers are returned and set via `int64`.  All doubles are returned and set via `float64`.

All TraceIDs and SpanIDs are returned as pdata [SpanID](https://github.com/open-telemetry/opentelemetry-collector/blob/main/pdata/pcommon/spanid.go) and [TraceID](https://github.com/open-telemetry/opentelemetry-collector/blob/main/pdata/pcommon/traceid.go) types.  Use the [SpanID function](https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/pkg/ottl/ottlfuncs/README.md#spanid) and [TraceID function](https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/pkg/ottl/ottlfuncs/README.md#traceid) when interacting with pdata representations of SpanID and TraceID.  When checking for nil, instead check against an empty byte slice (`SpanID(0x0000000000000000)` and `TraceID(0x00000000000000000000000000000000)`).  

The following paths are supported.

| path                                           | field accessed                                                                                                                                     | type                                                                    |
|------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------|
| cache                                          | the value of the current transform context's temporary cache. cache can be used as a temporary placeholder for data during complex transformations | pcommon.Map                                                             |
| cache\[""\]                                    | the value of an item in cache. Supports multiple indexes to access nested fields.                                                                  | string, bool, int64, float64, pcommon.Map, pcommon.Slice, []byte or nil |
| resource                                       | resource of the span being processed                                                                                                               | pcommon.Resource                                                        |
| resource.attributes                            | resource attributes of the span being processed                                                                                                    | pcommon.Map                                                             |
| resource.attributes\[""\]                      | the value of the resource attribute of the span being processed. Supports multiple indexes to access nested fields.                                | string, bool, int64, float64, pcommon.Map, pcommon.Slice, []byte or nil |
| resource.dropped_attributes_count              | number of dropped attributes of the resource of the span being processed                                                                           | int64                                                                   |
| instrumentation_scope                          | instrumentation scope of the span being processed                                                                                                  | pcommon.InstrumentationScope                                            |
| instrumentation_scope.name                     | name of the instrumentation scope of the span being processed                                                                                      | string                                                                  |
| instrumentation_scope.version                  | version of the instrumentation scope of the span being processed                                                                                   | string                                                                  |
| instrumentation_scope.dropped_attributes_count | number of dropped attributes of the instrumentation scope of the span being processed                                                              | int64                                                                   |
| instrumentation_scope.attributes               | instrumentation scope attributes of the span being processed                                                                                       | pcommon.Map                                                             |
| instrumentation_scope.attributes\[""\]         | the value of the instrumentation scope attribute of the span being processed. Supports multiple indexes to access nested fields.                   | string, bool, int64, float64, pcommon.Map, pcommon.Slice, []byte or nil |
| attributes                                     | attributes of the span being processed                                                                                                             | pcommon.Map                                                             |
| attributes\[""\]                               | the value of the attribute of the span being processed. Supports multiple indexes to access nested fields.                                         | string, bool, int64, float64, pcommon.Map, pcommon.Slice, []byte or nil |
| trace_id                                       | a byte slice representation of the trace id                                                                                                        | pcommon.TraceID                                                         |
| trace_id.string                                | a string representation of the trace id                                                                                                            | string                                                                  |
| span_id                                        | a byte slice representation of the span id                                                                                                         | pcommon.SpanID                                                          |
| span_id.string                                 | a string representation of the span id                                                                                                             | string                                                                  |
| parent_span_id                                 | a byte slice representation of the parent span id                                                                                                  | pcommon.SpanID                                                          |
| parent_span_id.string                          | a string representation of the parent span id                                                                                                      | string                                                                  |
| trace_state                                    | the trace state of the current span                                                                                                                | string                                                                  |
| trace_state\[""\]                              | an individual entry in the trace state                                                                                                             | string                                                                  |
| status                                         | the status of the span being processed                                                                                                             | ptrace.Status                                                           |
| status.code                                    | the status code of the span being processed                                                                                                        | int64                                                                   |
| status.message                                 | the status message of the span being processed                                                                                                     | string                                                                  |
| name                                           | the name of the span                                                                                                                               | string                                                                  |
| kind                                           | the kind of the span                                                                                                                               | int64                                                                   |
| start_time_unix_nano                           | the start time in unix nano of the span                                                                                                            | int64                                                                   |
| end_time_unix_nano                             | the end time in unix nano of the span                                                                                                              | int64                                                                   |
| dropped_attributes_count                       | the dropped attributes count of the span                                                                                                           | int64                                                                   |
| events                                         | the events of the span                                                                                                                             | ptrace.SpanEventSlice                                                   |
| dropped_events_count                           | the dropped events count of the span                                                                                                               | int64                                                                   |
| links                                          | the links of the span                                                                                                                              | ptrace.SpanLinkSlice                                                    |
| dropped_links_count                            | the dropped links count of the span                                                                                                                | int64                                                                   |


## Enums

The Span Context supports the enum names from the [traces proto](https://github.com/open-telemetry/opentelemetry-proto/blob/main/opentelemetry/proto/trace/v1/trace.proto).

| Enum Symbol           | Value |
|-----------------------|-------|
| SPAN_KIND_UNSPECIFIED | 0     |
| SPAN_KIND_INTERNAL    | 1     |
| SPAN_KIND_SERVER      | 2     |
| SPAN_KIND_CLIENT      | 3     |
| 	SPAN_KIND_PRODUCER   | 4     |
| 	SPAN_KIND_CONSUMER   | 5     |
| 	STATUS_CODE_UNSET    | 0     |
| 	STATUS_CODE_OK       | 1     |
| 	STATUS_CODE_ERROR    | 2     |
