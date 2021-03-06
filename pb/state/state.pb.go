/*
Copyright The TestGrid Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: state.proto

package state

import (
	fmt "fmt"
	config "github.com/GoogleCloudPlatform/testgrid/pb/config"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A metric and its values for each test cycle.
type Metric struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Sparse encoding of values. Indices is a list of pairs of <index, count>
	// that details columns with metric values. So given:
	//   Indices: [0, 2, 6, 4]
	//   Values: [0.1,0.2,6.1,6.2,6.3,6.4]
	// Decoded 12-value equivalent is:
	// [0.1, 0.2, nil, nil, nil, nil, 6.1, 6.2, 6.3, 6.4, nil, nil, ...]
	Indices              []int32   `protobuf:"varint,2,rep,packed,name=indices,proto3" json:"indices,omitempty"`
	Values               []float64 `protobuf:"fixed64,3,rep,packed,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Metric) Reset()         { *m = Metric{} }
func (m *Metric) String() string { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()    {}
func (*Metric) Descriptor() ([]byte, []int) {
	return fileDescriptor_a888679467bb7853, []int{0}
}

func (m *Metric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metric.Unmarshal(m, b)
}
func (m *Metric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metric.Marshal(b, m, deterministic)
}
func (m *Metric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metric.Merge(m, src)
}
func (m *Metric) XXX_Size() int {
	return xxx_messageInfo_Metric.Size(m)
}
func (m *Metric) XXX_DiscardUnknown() {
	xxx_messageInfo_Metric.DiscardUnknown(m)
}

var xxx_messageInfo_Metric proto.InternalMessageInfo

func (m *Metric) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Metric) GetIndices() []int32 {
	if m != nil {
		return m.Indices
	}
	return nil
}

func (m *Metric) GetValues() []float64 {
	if m != nil {
		return m.Values
	}
	return nil
}

type UpdatePhaseData struct {
	// The name for a part of the update cycle.
	PhaseName string `protobuf:"bytes,1,opt,name=phase_name,json=phaseName,proto3" json:"phase_name,omitempty"`
	// Time taken for a part of the update cycle, in seconds.
	PhaseSeconds         float64  `protobuf:"fixed64,2,opt,name=phase_seconds,json=phaseSeconds,proto3" json:"phase_seconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePhaseData) Reset()         { *m = UpdatePhaseData{} }
func (m *UpdatePhaseData) String() string { return proto.CompactTextString(m) }
func (*UpdatePhaseData) ProtoMessage()    {}
func (*UpdatePhaseData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a888679467bb7853, []int{1}
}

func (m *UpdatePhaseData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePhaseData.Unmarshal(m, b)
}
func (m *UpdatePhaseData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePhaseData.Marshal(b, m, deterministic)
}
func (m *UpdatePhaseData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePhaseData.Merge(m, src)
}
func (m *UpdatePhaseData) XXX_Size() int {
	return xxx_messageInfo_UpdatePhaseData.Size(m)
}
func (m *UpdatePhaseData) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePhaseData.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePhaseData proto.InternalMessageInfo

func (m *UpdatePhaseData) GetPhaseName() string {
	if m != nil {
		return m.PhaseName
	}
	return ""
}

func (m *UpdatePhaseData) GetPhaseSeconds() float64 {
	if m != nil {
		return m.PhaseSeconds
	}
	return 0
}

// Info on time taken to update test results during the last update cycle.
type UpdateInfo struct {
	// Metrics for how long parts of the update cycle take.
	UpdatePhaseData      []*UpdatePhaseData `protobuf:"bytes,1,rep,name=update_phase_data,json=updatePhaseData,proto3" json:"update_phase_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UpdateInfo) Reset()         { *m = UpdateInfo{} }
func (m *UpdateInfo) String() string { return proto.CompactTextString(m) }
func (*UpdateInfo) ProtoMessage()    {}
func (*UpdateInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a888679467bb7853, []int{2}
}

func (m *UpdateInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateInfo.Unmarshal(m, b)
}
func (m *UpdateInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateInfo.Marshal(b, m, deterministic)
}
func (m *UpdateInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateInfo.Merge(m, src)
}
func (m *UpdateInfo) XXX_Size() int {
	return xxx_messageInfo_UpdateInfo.Size(m)
}
func (m *UpdateInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateInfo proto.InternalMessageInfo

func (m *UpdateInfo) GetUpdatePhaseData() []*UpdatePhaseData {
	if m != nil {
		return m.UpdatePhaseData
	}
	return nil
}

// Info on a failing test row about the failure.
type AlertInfo struct {
	// Number of results that have failed.
	FailCount int32 `protobuf:"varint,1,opt,name=fail_count,json=failCount,proto3" json:"fail_count,omitempty"`
	// The build ID the test first failed at.
	FailBuildId string `protobuf:"bytes,2,opt,name=fail_build_id,json=failBuildId,proto3" json:"fail_build_id,omitempty"`
	// The time the test first failed at.
	FailTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=fail_time,json=failTime,proto3" json:"fail_time,omitempty"`
	// The test ID for the first test failure.
	FailTestId string `protobuf:"bytes,4,opt,name=fail_test_id,json=failTestId,proto3" json:"fail_test_id,omitempty"`
	// The build ID the test last passed at.
	PassBuildId string `protobuf:"bytes,5,opt,name=pass_build_id,json=passBuildId,proto3" json:"pass_build_id,omitempty"`
	// The time the test last passed at.
	PassTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=pass_time,json=passTime,proto3" json:"pass_time,omitempty"`
	// A snippet explaining the failure.
	FailureMessage string `protobuf:"bytes,7,opt,name=failure_message,json=failureMessage,proto3" json:"failure_message,omitempty"`
	// Link to search for build changes, internally a code-search link.
	BuildLink string `protobuf:"bytes,8,opt,name=build_link,json=buildLink,proto3" json:"build_link,omitempty"`
	// Text for option to search for build changes.
	BuildLinkText string `protobuf:"bytes,9,opt,name=build_link_text,json=buildLinkText,proto3" json:"build_link_text,omitempty"`
	// Text to display for link to search for build changes.
	BuildUrlText string `protobuf:"bytes,10,opt,name=build_url_text,json=buildUrlText,proto3" json:"build_url_text,omitempty"`
	// The build ID for the latest test failure. (Does not indicate the failure is
	// 'over', just the latest test failure we found.)
	LatestFailBuildId string `protobuf:"bytes,11,opt,name=latest_fail_build_id,json=latestFailBuildId,proto3" json:"latest_fail_build_id,omitempty"`
	// The test ID for the latest test failure.
	LatestFailTestId string `protobuf:"bytes,14,opt,name=latest_fail_test_id,json=latestFailTestId,proto3" json:"latest_fail_test_id,omitempty"`
	// Maps (property name):(property value) for arbitrary alert properties.
	Properties map[string]string `protobuf:"bytes,12,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A list of IDs for issue hotlists related to this failure.
	HotlistIds []string `protobuf:"bytes,13,rep,name=hotlist_ids,json=hotlistIds,proto3" json:"hotlist_ids,omitempty"`
	// Dynamic email list, route email alerts to these instead of the configured defaults.
	EmailAddresses       []string `protobuf:"bytes,15,rep,name=email_addresses,json=emailAddresses,proto3" json:"email_addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlertInfo) Reset()         { *m = AlertInfo{} }
func (m *AlertInfo) String() string { return proto.CompactTextString(m) }
func (*AlertInfo) ProtoMessage()    {}
func (*AlertInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a888679467bb7853, []int{3}
}

func (m *AlertInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlertInfo.Unmarshal(m, b)
}
func (m *AlertInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlertInfo.Marshal(b, m, deterministic)
}
func (m *AlertInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlertInfo.Merge(m, src)
}
func (m *AlertInfo) XXX_Size() int {
	return xxx_messageInfo_AlertInfo.Size(m)
}
func (m *AlertInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AlertInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AlertInfo proto.InternalMessageInfo

func (m *AlertInfo) GetFailCount() int32 {
	if m != nil {
		return m.FailCount
	}
	return 0
}

func (m *AlertInfo) GetFailBuildId() string {
	if m != nil {
		return m.FailBuildId
	}
	return ""
}

func (m *AlertInfo) GetFailTime() *timestamp.Timestamp {
	if m != nil {
		return m.FailTime
	}
	return nil
}

func (m *AlertInfo) GetFailTestId() string {
	if m != nil {
		return m.FailTestId
	}
	return ""
}

func (m *AlertInfo) GetPassBuildId() string {
	if m != nil {
		return m.PassBuildId
	}
	return ""
}

func (m *AlertInfo) GetPassTime() *timestamp.Timestamp {
	if m != nil {
		return m.PassTime
	}
	return nil
}

func (m *AlertInfo) GetFailureMessage() string {
	if m != nil {
		return m.FailureMessage
	}
	return ""
}

func (m *AlertInfo) GetBuildLink() string {
	if m != nil {
		return m.BuildLink
	}
	return ""
}

func (m *AlertInfo) GetBuildLinkText() string {
	if m != nil {
		return m.BuildLinkText
	}
	return ""
}

func (m *AlertInfo) GetBuildUrlText() string {
	if m != nil {
		return m.BuildUrlText
	}
	return ""
}

func (m *AlertInfo) GetLatestFailBuildId() string {
	if m != nil {
		return m.LatestFailBuildId
	}
	return ""
}

func (m *AlertInfo) GetLatestFailTestId() string {
	if m != nil {
		return m.LatestFailTestId
	}
	return ""
}

func (m *AlertInfo) GetProperties() map[string]string {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *AlertInfo) GetHotlistIds() []string {
	if m != nil {
		return m.HotlistIds
	}
	return nil
}

func (m *AlertInfo) GetEmailAddresses() []string {
	if m != nil {
		return m.EmailAddresses
	}
	return nil
}

// Info on default test metadata for a dashboard tab.
type TestMetadata struct {
	// Name of the test with associated test metadata.
	TestName string `protobuf:"bytes,1,opt,name=test_name,json=testName,proto3" json:"test_name,omitempty"`
	// Default bug component.
	BugComponent int32 `protobuf:"varint,2,opt,name=bug_component,json=bugComponent,proto3" json:"bug_component,omitempty"`
	// Default owner.
	Owner string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	// Default list of cc's.
	Cc []string `protobuf:"bytes,4,rep,name=cc,proto3" json:"cc,omitempty"`
	// When present, only file a bug for failed tests with same error type.
	// Otherwise, always file a bug.
	ErrorType            string   `protobuf:"bytes,5,opt,name=error_type,json=errorType,proto3" json:"error_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestMetadata) Reset()         { *m = TestMetadata{} }
func (m *TestMetadata) String() string { return proto.CompactTextString(m) }
func (*TestMetadata) ProtoMessage()    {}
func (*TestMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_a888679467bb7853, []int{4}
}

func (m *TestMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestMetadata.Unmarshal(m, b)
}
func (m *TestMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestMetadata.Marshal(b, m, deterministic)
}
func (m *TestMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestMetadata.Merge(m, src)
}
func (m *TestMetadata) XXX_Size() int {
	return xxx_messageInfo_TestMetadata.Size(m)
}
func (m *TestMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TestMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TestMetadata proto.InternalMessageInfo

func (m *TestMetadata) GetTestName() string {
	if m != nil {
		return m.TestName
	}
	return ""
}

func (m *TestMetadata) GetBugComponent() int32 {
	if m != nil {
		return m.BugComponent
	}
	return 0
}

func (m *TestMetadata) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *TestMetadata) GetCc() []string {
	if m != nil {
		return m.Cc
	}
	return nil
}

func (m *TestMetadata) GetErrorType() string {
	if m != nil {
		return m.ErrorType
	}
	return ""
}

// TestGrid columns (also known as TestCycle).
type Column struct {
	// Unique instance of the job, typically BUILD_NUMBER from prow or a guid
	Build string `protobuf:"bytes,1,opt,name=build,proto3" json:"build,omitempty"`
	// Name associated with the column (such as the run/invocation ID).No two
	// columns should have the same build_id and name. The name field allows the
	// display of multiple columns with the same build_id.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Milliseconds since start of epoch (python time.time() * 1000)
	Started float64 `protobuf:"fixed64,3,opt,name=started,proto3" json:"started,omitempty"`
	// Additional custom headers like commit, image used, etc.
	Extra []string `protobuf:"bytes,4,rep,name=extra,proto3" json:"extra,omitempty"`
	// Custom hotlist ids.
	HotlistIds string `protobuf:"bytes,5,opt,name=hotlist_ids,json=hotlistIds,proto3" json:"hotlist_ids,omitempty"`
	// An optional hint for the updater.
	Hint string `protobuf:"bytes,6,opt,name=hint,proto3" json:"hint,omitempty"`
	// Dynamic email list, route email alerts to these instead of the configured defaults.
	EmailAddresses       []string `protobuf:"bytes,7,rep,name=email_addresses,json=emailAddresses,proto3" json:"email_addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Column) Reset()         { *m = Column{} }
func (m *Column) String() string { return proto.CompactTextString(m) }
func (*Column) ProtoMessage()    {}
func (*Column) Descriptor() ([]byte, []int) {
	return fileDescriptor_a888679467bb7853, []int{5}
}

func (m *Column) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Column.Unmarshal(m, b)
}
func (m *Column) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Column.Marshal(b, m, deterministic)
}
func (m *Column) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Column.Merge(m, src)
}
func (m *Column) XXX_Size() int {
	return xxx_messageInfo_Column.Size(m)
}
func (m *Column) XXX_DiscardUnknown() {
	xxx_messageInfo_Column.DiscardUnknown(m)
}

var xxx_messageInfo_Column proto.InternalMessageInfo

func (m *Column) GetBuild() string {
	if m != nil {
		return m.Build
	}
	return ""
}

func (m *Column) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Column) GetStarted() float64 {
	if m != nil {
		return m.Started
	}
	return 0
}

func (m *Column) GetExtra() []string {
	if m != nil {
		return m.Extra
	}
	return nil
}

func (m *Column) GetHotlistIds() string {
	if m != nil {
		return m.HotlistIds
	}
	return ""
}

func (m *Column) GetHint() string {
	if m != nil {
		return m.Hint
	}
	return ""
}

func (m *Column) GetEmailAddresses() []string {
	if m != nil {
		return m.EmailAddresses
	}
	return nil
}

// TestGrid rows (also known as TestRow)
type Row struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Results for this row, run-length encoded to reduce size/improve
	// performance. Thus (encoded -> decoded equivalent):
	//   [0, 3, 5, 4] -> [0, 0, 0, 5, 5, 5, 5]
	//   [5, 1] -> [5]
	//   [1, 5] -> [1, 1, 1, 1, 1]
	// The decoded values are Result enums
	Results []int32 `protobuf:"varint,3,rep,packed,name=results,proto3" json:"results,omitempty"`
	// Test IDs for each test result in this test case.
	// Must be present on every column, regardless of status.
	CellIds []string `protobuf:"bytes,4,rep,name=cell_ids,json=cellIds,proto3" json:"cell_ids,omitempty"`
	// Short description of the result, displayed on mouseover.
	// Present for any column with a non-empty status (not NO_RESULT).
	Messages []string `protobuf:"bytes,5,rep,name=messages,proto3" json:"messages,omitempty"`
	// Names of metrics associated with this test case. Stored separate from
	// metric info (which may be omitted).
	Metric  []string  `protobuf:"bytes,7,rep,name=metric,proto3" json:"metric,omitempty"`
	Metrics []*Metric `protobuf:"bytes,8,rep,name=metrics,proto3" json:"metrics,omitempty"`
	// Short string to place inside the cell (F for fail, etc)
	// Present for any column with a non-empty status (not NO_RESULT).
	Icons []string `protobuf:"bytes,9,rep,name=icons,proto3" json:"icons,omitempty"`
	// IDs for issues associated with this row.
	Issues []string `protobuf:"bytes,10,rep,name=issues,proto3" json:"issues,omitempty"`
	// An alert for the failure if there's a recent failure for this row.
	AlertInfo *AlertInfo `protobuf:"bytes,11,opt,name=alert_info,json=alertInfo,proto3" json:"alert_info,omitempty"`
	// Values of a user-defined property found in cells for this row.
	UserProperty         []string `protobuf:"bytes,12,rep,name=user_property,json=userProperty,proto3" json:"user_property,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Row) Reset()         { *m = Row{} }
func (m *Row) String() string { return proto.CompactTextString(m) }
func (*Row) ProtoMessage()    {}
func (*Row) Descriptor() ([]byte, []int) {
	return fileDescriptor_a888679467bb7853, []int{6}
}

func (m *Row) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Row.Unmarshal(m, b)
}
func (m *Row) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Row.Marshal(b, m, deterministic)
}
func (m *Row) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Row.Merge(m, src)
}
func (m *Row) XXX_Size() int {
	return xxx_messageInfo_Row.Size(m)
}
func (m *Row) XXX_DiscardUnknown() {
	xxx_messageInfo_Row.DiscardUnknown(m)
}

var xxx_messageInfo_Row proto.InternalMessageInfo

func (m *Row) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Row) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Row) GetResults() []int32 {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *Row) GetCellIds() []string {
	if m != nil {
		return m.CellIds
	}
	return nil
}

func (m *Row) GetMessages() []string {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *Row) GetMetric() []string {
	if m != nil {
		return m.Metric
	}
	return nil
}

func (m *Row) GetMetrics() []*Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *Row) GetIcons() []string {
	if m != nil {
		return m.Icons
	}
	return nil
}

func (m *Row) GetIssues() []string {
	if m != nil {
		return m.Issues
	}
	return nil
}

func (m *Row) GetAlertInfo() *AlertInfo {
	if m != nil {
		return m.AlertInfo
	}
	return nil
}

func (m *Row) GetUserProperty() []string {
	if m != nil {
		return m.UserProperty
	}
	return nil
}

// A single table of test results backing a dashboard tab.
type Grid struct {
	// A cycle of test results, not including the results. In the TestGrid client,
	// the cycles define the columns.
	Columns []*Column `protobuf:"bytes,1,rep,name=columns,proto3" json:"columns,omitempty"`
	// A test case with test results. In the TestGrid client, the cases define the
	// rows (and the results define the individual cells).
	Rows []*Row `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty"`
	// The latest configuration used to generate this test group.
	Config *config.TestGroup `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
	// Seconds since epoch for last time this cycle was updated.
	LastTimeUpdated float64 `protobuf:"fixed64,6,opt,name=last_time_updated,json=lastTimeUpdated,proto3" json:"last_time_updated,omitempty"`
	// Stored info on previous timing for parts of the update cycle.
	UpdateInfo []*UpdateInfo `protobuf:"bytes,8,rep,name=update_info,json=updateInfo,proto3" json:"update_info,omitempty"`
	// Stored info on default test metadata.
	TestMetadata []*TestMetadata `protobuf:"bytes,9,rep,name=test_metadata,json=testMetadata,proto3" json:"test_metadata,omitempty"`
	// Clusters of failures for a TestResultTable instance.
	Cluster []*Cluster `protobuf:"bytes,10,rep,name=cluster,proto3" json:"cluster,omitempty"`
	// Most recent timestamp that clusters have processed.
	MostRecentClusterTimestamp float64  `protobuf:"fixed64,11,opt,name=most_recent_cluster_timestamp,json=mostRecentClusterTimestamp,proto3" json:"most_recent_cluster_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *Grid) Reset()         { *m = Grid{} }
func (m *Grid) String() string { return proto.CompactTextString(m) }
func (*Grid) ProtoMessage()    {}
func (*Grid) Descriptor() ([]byte, []int) {
	return fileDescriptor_a888679467bb7853, []int{7}
}

func (m *Grid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Grid.Unmarshal(m, b)
}
func (m *Grid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Grid.Marshal(b, m, deterministic)
}
func (m *Grid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Grid.Merge(m, src)
}
func (m *Grid) XXX_Size() int {
	return xxx_messageInfo_Grid.Size(m)
}
func (m *Grid) XXX_DiscardUnknown() {
	xxx_messageInfo_Grid.DiscardUnknown(m)
}

var xxx_messageInfo_Grid proto.InternalMessageInfo

func (m *Grid) GetColumns() []*Column {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *Grid) GetRows() []*Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

func (m *Grid) GetConfig() *config.TestGroup {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *Grid) GetLastTimeUpdated() float64 {
	if m != nil {
		return m.LastTimeUpdated
	}
	return 0
}

func (m *Grid) GetUpdateInfo() []*UpdateInfo {
	if m != nil {
		return m.UpdateInfo
	}
	return nil
}

func (m *Grid) GetTestMetadata() []*TestMetadata {
	if m != nil {
		return m.TestMetadata
	}
	return nil
}

func (m *Grid) GetCluster() []*Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

func (m *Grid) GetMostRecentClusterTimestamp() float64 {
	if m != nil {
		return m.MostRecentClusterTimestamp
	}
	return 0
}

// A cluster of failures grouped by test status and message for a test results
// table.
type Cluster struct {
	// Test status cluster grouped by.
	TestStatus int32 `protobuf:"varint,1,opt,name=test_status,json=testStatus,proto3" json:"test_status,omitempty"`
	// Error message or testFailureClassification string cluster grouped by.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// ClusterRows that belong to this cluster.
	ClusterRow           []*ClusterRow `protobuf:"bytes,3,rep,name=cluster_row,json=clusterRow,proto3" json:"cluster_row,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Cluster) Reset()         { *m = Cluster{} }
func (m *Cluster) String() string { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()    {}
func (*Cluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_a888679467bb7853, []int{8}
}

func (m *Cluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cluster.Unmarshal(m, b)
}
func (m *Cluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cluster.Marshal(b, m, deterministic)
}
func (m *Cluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cluster.Merge(m, src)
}
func (m *Cluster) XXX_Size() int {
	return xxx_messageInfo_Cluster.Size(m)
}
func (m *Cluster) XXX_DiscardUnknown() {
	xxx_messageInfo_Cluster.DiscardUnknown(m)
}

var xxx_messageInfo_Cluster proto.InternalMessageInfo

func (m *Cluster) GetTestStatus() int32 {
	if m != nil {
		return m.TestStatus
	}
	return 0
}

func (m *Cluster) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Cluster) GetClusterRow() []*ClusterRow {
	if m != nil {
		return m.ClusterRow
	}
	return nil
}

// Cells in a TestRow that belong to a specific Cluster.
type ClusterRow struct {
	// Name of TestRow.
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Index within row that belongs to Cluster (refer to columns of the row).
	Index                []int32  `protobuf:"varint,2,rep,packed,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterRow) Reset()         { *m = ClusterRow{} }
func (m *ClusterRow) String() string { return proto.CompactTextString(m) }
func (*ClusterRow) ProtoMessage()    {}
func (*ClusterRow) Descriptor() ([]byte, []int) {
	return fileDescriptor_a888679467bb7853, []int{9}
}

func (m *ClusterRow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterRow.Unmarshal(m, b)
}
func (m *ClusterRow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterRow.Marshal(b, m, deterministic)
}
func (m *ClusterRow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterRow.Merge(m, src)
}
func (m *ClusterRow) XXX_Size() int {
	return xxx_messageInfo_ClusterRow.Size(m)
}
func (m *ClusterRow) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterRow.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterRow proto.InternalMessageInfo

func (m *ClusterRow) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *ClusterRow) GetIndex() []int32 {
	if m != nil {
		return m.Index
	}
	return nil
}

func init() {
	proto.RegisterType((*Metric)(nil), "Metric")
	proto.RegisterType((*UpdatePhaseData)(nil), "UpdatePhaseData")
	proto.RegisterType((*UpdateInfo)(nil), "UpdateInfo")
	proto.RegisterType((*AlertInfo)(nil), "AlertInfo")
	proto.RegisterMapType((map[string]string)(nil), "AlertInfo.PropertiesEntry")
	proto.RegisterType((*TestMetadata)(nil), "TestMetadata")
	proto.RegisterType((*Column)(nil), "Column")
	proto.RegisterType((*Row)(nil), "Row")
	proto.RegisterType((*Grid)(nil), "Grid")
	proto.RegisterType((*Cluster)(nil), "Cluster")
	proto.RegisterType((*ClusterRow)(nil), "ClusterRow")
}

func init() { proto.RegisterFile("state.proto", fileDescriptor_a888679467bb7853) }

var fileDescriptor_a888679467bb7853 = []byte{
	// 1100 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x56, 0xcf, 0x8e, 0xdb, 0x36,
	0x13, 0x87, 0x6c, 0xd9, 0x96, 0x46, 0xf6, 0xda, 0xe1, 0x17, 0x04, 0xfa, 0xb6, 0x08, 0xe2, 0x28,
	0x45, 0xbb, 0x2d, 0x5a, 0x2d, 0xe0, 0x1e, 0x5a, 0x04, 0xed, 0x21, 0xdd, 0xa6, 0xc1, 0x2e, 0x9a,
	0x20, 0x60, 0x36, 0x67, 0x41, 0x2b, 0x71, 0x1d, 0x21, 0xb2, 0x28, 0x90, 0x54, 0xbd, 0x7e, 0x90,
	0x3e, 0x4c, 0x6f, 0xed, 0x0b, 0xf5, 0x19, 0x8a, 0x19, 0x52, 0xb6, 0x13, 0x04, 0xe8, 0xc9, 0xfc,
	0xfd, 0x66, 0x34, 0x33, 0x9c, 0x7f, 0x34, 0x44, 0xda, 0xe4, 0x46, 0xa4, 0xad, 0x92, 0x46, 0x9e,
	0x3e, 0x5a, 0x4b, 0xb9, 0xae, 0xc5, 0x39, 0xa1, 0x9b, 0xee, 0xf6, 0xdc, 0x54, 0x1b, 0xa1, 0x4d,
	0xbe, 0x69, 0x9d, 0xc2, 0x83, 0xf6, 0xe6, 0xbc, 0x90, 0xcd, 0x6d, 0xb5, 0x76, 0x3f, 0x96, 0x4f,
	0x5e, 0xc1, 0xf8, 0xa5, 0x30, 0xaa, 0x2a, 0x18, 0x03, 0xbf, 0xc9, 0x37, 0x22, 0xf6, 0x96, 0xde,
	0x59, 0xc8, 0xe9, 0xcc, 0x62, 0x98, 0x54, 0x4d, 0x59, 0x15, 0x42, 0xc7, 0x83, 0xe5, 0xf0, 0x6c,
	0xc4, 0x7b, 0xc8, 0x1e, 0xc0, 0xf8, 0xf7, 0xbc, 0xee, 0x84, 0x8e, 0x87, 0xcb, 0xe1, 0x99, 0xc7,
	0x1d, 0x4a, 0xde, 0xc2, 0xfc, 0x6d, 0x5b, 0xe6, 0x46, 0xbc, 0x7e, 0x97, 0x6b, 0xf1, 0x4b, 0x6e,
	0x72, 0xf6, 0x10, 0xa0, 0x45, 0x90, 0x1d, 0x99, 0x0f, 0x89, 0x79, 0x85, 0x3e, 0x9e, 0xc0, 0xcc,
	0x8a, 0xb5, 0x28, 0x64, 0x53, 0xa2, 0x27, 0xef, 0xcc, 0xe3, 0x53, 0x22, 0xdf, 0x58, 0x2e, 0xb9,
	0x02, 0xb0, 0x66, 0x2f, 0x9b, 0x5b, 0xc9, 0x7e, 0x84, 0x7b, 0x1d, 0xa1, 0xcc, 0x7e, 0x59, 0xe6,
	0x26, 0x8f, 0xbd, 0xe5, 0xf0, 0x2c, 0x5a, 0x2d, 0xd2, 0x8f, 0xdc, 0xf3, 0x79, 0xf7, 0x21, 0x91,
	0xfc, 0x3d, 0x82, 0xf0, 0x59, 0x2d, 0x94, 0x21, 0x5b, 0x0f, 0x01, 0x6e, 0xf3, 0xaa, 0xce, 0x0a,
	0xd9, 0x35, 0x86, 0xa2, 0x1b, 0xf1, 0x10, 0x99, 0x0b, 0x24, 0x58, 0x02, 0x33, 0x12, 0xdf, 0x74,
	0x55, 0x5d, 0x66, 0x55, 0x49, 0xd1, 0x85, 0x3c, 0x42, 0xf2, 0x67, 0xe4, 0x2e, 0x4b, 0xf6, 0x3d,
	0xd0, 0x07, 0x19, 0xe6, 0x3c, 0x1e, 0x2e, 0xbd, 0xb3, 0x68, 0x75, 0x9a, 0xda, 0x82, 0xa4, 0x7d,
	0x41, 0xd2, 0xeb, 0xbe, 0x20, 0x3c, 0x40, 0x65, 0x84, 0x6c, 0x09, 0x53, 0xfb, 0xa1, 0xd0, 0x06,
	0x6d, 0xfb, 0x64, 0x9b, 0xe2, 0xb9, 0x16, 0xda, 0x5c, 0x96, 0xe8, 0xbe, 0xcd, 0xb5, 0x3e, 0xb8,
	0x1f, 0x59, 0xf7, 0x48, 0x1e, 0xb9, 0x27, 0x1d, 0x72, 0x3f, 0xfe, 0x6f, 0xf7, 0xa8, 0x4c, 0xee,
	0xbf, 0x84, 0x39, 0xba, 0xea, 0x94, 0xc8, 0x36, 0x42, 0xeb, 0x7c, 0x2d, 0xe2, 0x09, 0x99, 0x3f,
	0x71, 0xf4, 0x4b, 0xcb, 0x62, 0x8e, 0x6c, 0x00, 0x75, 0xd5, 0xbc, 0x8f, 0x03, 0x5b, 0x41, 0x62,
	0x7e, 0xab, 0x9a, 0xf7, 0xec, 0x0b, 0x98, 0x1f, 0xc4, 0x99, 0x11, 0x77, 0x26, 0x0e, 0x49, 0x67,
	0xb6, 0xd7, 0xb9, 0x16, 0x77, 0x86, 0x7d, 0x0e, 0x27, 0x56, 0xaf, 0x53, 0xb5, 0x55, 0x03, 0x52,
	0x9b, 0x12, 0xfb, 0x56, 0xd5, 0xa4, 0x75, 0x0e, 0xf7, 0xeb, 0x9c, 0x32, 0xf2, 0x61, 0xe2, 0x23,
	0xd2, 0xbd, 0x67, 0x65, 0xbf, 0x1e, 0xa5, 0xff, 0x5b, 0xf8, 0xdf, 0xf1, 0x07, 0x7d, 0x32, 0x4f,
	0x48, 0x7f, 0x71, 0xd0, 0x77, 0x29, 0x7d, 0x0a, 0xd0, 0x2a, 0xd9, 0x0a, 0x65, 0x2a, 0xa1, 0xe3,
	0x29, 0x75, 0xcd, 0x69, 0xba, 0x6f, 0x88, 0xf4, 0xf5, 0x5e, 0xf8, 0xbc, 0x31, 0x6a, 0xc7, 0x8f,
	0xb4, 0xd9, 0x23, 0x88, 0xde, 0x49, 0x53, 0x57, 0xe4, 0x41, 0xc7, 0xb3, 0xe5, 0x10, 0xeb, 0xe5,
	0xa8, 0xcb, 0x52, 0x63, 0x4a, 0xc5, 0x06, 0xa3, 0xc8, 0xcb, 0x52, 0x09, 0xad, 0x85, 0x8e, 0xe7,
	0xa4, 0x74, 0x42, 0xf4, 0xb3, 0x9e, 0x3d, 0xfd, 0x09, 0xe6, 0x1f, 0x39, 0x62, 0x0b, 0x18, 0xbe,
	0x17, 0x3b, 0x37, 0x20, 0x78, 0x64, 0xf7, 0x61, 0x44, 0x63, 0xe5, 0x9a, 0xce, 0x82, 0xa7, 0x83,
	0x1f, 0xbc, 0xe4, 0x0f, 0x0f, 0xa6, 0x78, 0x9f, 0x97, 0xc2, 0xe4, 0xd8, 0xfd, 0xec, 0x33, 0x08,
	0xe9, 0xe2, 0x47, 0x33, 0x16, 0x20, 0xd1, 0x8f, 0xd8, 0x4d, 0xb7, 0xce, 0x0a, 0xb9, 0x69, 0x65,
	0x23, 0x1a, 0x43, 0xf6, 0x46, 0x98, 0xf7, 0xf5, 0x45, 0xcf, 0xa1, 0x33, 0xb9, 0x6d, 0x84, 0xa2,
	0x0e, 0x0e, 0xb9, 0x05, 0xec, 0x04, 0x06, 0x45, 0x11, 0xfb, 0x74, 0x87, 0x41, 0x51, 0x60, 0x2b,
	0x08, 0xa5, 0xa4, 0xca, 0xcc, 0xae, 0x15, 0xae, 0x1b, 0x43, 0x62, 0xae, 0x77, 0xad, 0x48, 0xfe,
	0xf2, 0x60, 0x7c, 0x21, 0xeb, 0x6e, 0xd3, 0xa0, 0x3d, 0xaa, 0x9d, 0x8b, 0xc6, 0x82, 0xfd, 0x96,
	0x19, 0x7c, 0xb8, 0x65, 0xb4, 0xc9, 0x95, 0x11, 0x25, 0xf9, 0xf6, 0x78, 0x0f, 0xd1, 0x86, 0xb8,
	0x33, 0x2a, 0x77, 0x01, 0x58, 0xf0, 0x71, 0x15, 0x6c, 0x10, 0xc7, 0x55, 0x60, 0xe0, 0xbf, 0xab,
	0x1a, 0x43, 0xc3, 0x10, 0x72, 0x3a, 0x7f, 0xaa, 0x32, 0x93, 0x4f, 0x55, 0x26, 0xf9, 0x73, 0x00,
	0x43, 0x2e, 0xb7, 0x9f, 0xdc, 0x87, 0x27, 0x30, 0xd8, 0xaf, 0x80, 0x41, 0x55, 0x62, 0xe4, 0x4a,
	0xe8, 0xae, 0x36, 0x76, 0x0d, 0x8e, 0x78, 0x0f, 0xd9, 0xff, 0x21, 0x28, 0x44, 0x5d, 0x53, 0x80,
	0x36, 0xf8, 0x09, 0x62, 0x8c, 0xee, 0x14, 0x02, 0x37, 0x6e, 0x18, 0x3b, 0x8a, 0xf6, 0x18, 0xd7,
	0xea, 0x86, 0xd6, 0xb1, 0x0b, 0xce, 0x21, 0xf6, 0x18, 0x26, 0xf6, 0xa4, 0xe3, 0x80, 0x3a, 0x76,
	0x92, 0xda, 0xb5, 0xcd, 0x7b, 0x1e, 0x73, 0x55, 0x15, 0xb2, 0xd1, 0x71, 0x68, 0x73, 0x45, 0x00,
	0x0d, 0x56, 0x5a, 0xe3, 0x9e, 0x06, 0x6b, 0xd0, 0x22, 0xf6, 0x15, 0x40, 0x8e, 0x2d, 0x9f, 0x55,
	0xcd, 0xad, 0xa4, 0xd9, 0x8a, 0x56, 0x70, 0x98, 0x02, 0x1e, 0xe6, 0xfb, 0x0d, 0xf9, 0x04, 0x66,
	0x9d, 0x16, 0x2a, 0x73, 0x73, 0xb0, 0xa3, 0x99, 0x09, 0xf9, 0x14, 0x49, 0xd7, 0xc3, 0xbb, 0x2b,
	0x3f, 0x18, 0x2f, 0x26, 0xc9, 0x3f, 0x03, 0xf0, 0x5f, 0xa8, 0xaa, 0xc4, 0x78, 0x0b, 0x6a, 0x03,
	0xed, 0xf6, 0xf2, 0x24, 0xb5, 0x6d, 0xc1, 0x7b, 0x9e, 0xc5, 0xe0, 0x2b, 0xb9, 0xb5, 0x0f, 0x4b,
	0xb4, 0xf2, 0x53, 0x2e, 0xb7, 0x9c, 0x18, 0x96, 0xc0, 0xd8, 0xbe, 0x51, 0xb4, 0x10, 0x31, 0x2e,
	0x6c, 0xf5, 0x17, 0x4a, 0x76, 0x2d, 0x77, 0x12, 0xf6, 0x35, 0xdc, 0xab, 0x73, 0x6d, 0x68, 0xe9,
	0x65, 0x76, 0xc3, 0x97, 0x54, 0x6f, 0x8f, 0xcf, 0x51, 0x80, 0x0b, 0xce, 0xbe, 0x04, 0x25, 0xfb,
	0x06, 0x22, 0xf7, 0x5c, 0xd0, 0x65, 0x6d, 0x02, 0xa3, 0xf4, 0xf0, 0xa0, 0x70, 0xe8, 0x0e, 0x8f,
	0xcb, 0x0a, 0x66, 0x34, 0x49, 0x1b, 0x37, 0x5a, 0x94, 0xcf, 0x68, 0x35, 0x4b, 0x8f, 0xe7, 0x8d,
	0x4f, 0xcd, 0xf1, 0xf4, 0x25, 0x30, 0x29, 0xea, 0x4e, 0x1b, 0xa1, 0x28, 0xcd, 0xd1, 0x2a, 0x48,
	0x2f, 0x2c, 0xe6, 0xbd, 0x80, 0x3d, 0x83, 0x87, 0x1b, 0xa9, 0x4d, 0xa6, 0x44, 0x21, 0x1a, 0x93,
	0x39, 0x3a, 0xdb, 0x3f, 0xd4, 0x54, 0x04, 0x8f, 0x9f, 0xa2, 0x12, 0x27, 0x1d, 0x67, 0x62, 0xbf,
	0xba, 0xaf, 0xfc, 0x60, 0xb8, 0xf0, 0xaf, 0xfc, 0x60, 0xb4, 0x18, 0x5f, 0xf9, 0xc1, 0x64, 0x11,
	0x24, 0x0a, 0x26, 0x4e, 0x0b, 0xa7, 0x82, 0xe2, 0xc6, 0xbf, 0x05, 0x9d, 0x76, 0x2f, 0x19, 0x20,
	0xf5, 0x86, 0x18, 0x6c, 0xd6, 0x7e, 0xcd, 0xdb, 0x0e, 0xee, 0x21, 0x26, 0xa8, 0x0f, 0x47, 0xc9,
	0x2d, 0xb5, 0x32, 0x26, 0xa8, 0xbf, 0x82, 0xdc, 0x72, 0x28, 0xf6, 0xe7, 0xe4, 0x39, 0xc0, 0x41,
	0xc2, 0x1e, 0xc3, 0xb4, 0xac, 0x74, 0x5b, 0xe7, 0xbb, 0xe3, 0xdd, 0x13, 0x39, 0x8e, 0xd6, 0x0f,
	0x76, 0x66, 0x53, 0x8a, 0x3b, 0xf7, 0x1f, 0xc2, 0x82, 0x9b, 0x31, 0xbd, 0x4d, 0xdf, 0xfd, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0x23, 0xe8, 0x42, 0x8f, 0xc8, 0x08, 0x00, 0x00,
}
