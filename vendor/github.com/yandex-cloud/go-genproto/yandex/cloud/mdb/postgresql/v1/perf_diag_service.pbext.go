// Code generated by protoc-gen-goext. DO NOT EDIT.

package postgresql

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *ListRawStatementsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListRawStatementsRequest) SetFromTime(v *timestamppb.Timestamp) {
	m.FromTime = v
}

func (m *ListRawStatementsRequest) SetToTime(v *timestamppb.Timestamp) {
	m.ToTime = v
}

func (m *ListRawStatementsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListRawStatementsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListRawSessionStatesRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListRawSessionStatesRequest) SetFromTime(v *timestamppb.Timestamp) {
	m.FromTime = v
}

func (m *ListRawSessionStatesRequest) SetToTime(v *timestamppb.Timestamp) {
	m.ToTime = v
}

func (m *ListRawSessionStatesRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListRawSessionStatesRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListRawSessionStatesResponse) SetSessionStates(v []*SessionState) {
	m.SessionStates = v
}

func (m *ListRawSessionStatesResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ListRawStatementsResponse) SetStatements(v []*QueryStatement) {
	m.Statements = v
}

func (m *ListRawStatementsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
