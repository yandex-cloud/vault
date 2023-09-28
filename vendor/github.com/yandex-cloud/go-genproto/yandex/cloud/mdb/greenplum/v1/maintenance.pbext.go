// Code generated by protoc-gen-goext. DO NOT EDIT.

package greenplum

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type MaintenanceWindow_Policy = isMaintenanceWindow_Policy

func (m *MaintenanceWindow) SetPolicy(v MaintenanceWindow_Policy) {
	m.Policy = v
}

func (m *MaintenanceWindow) SetAnytime(v *AnytimeMaintenanceWindow) {
	m.Policy = &MaintenanceWindow_Anytime{
		Anytime: v,
	}
}

func (m *MaintenanceWindow) SetWeeklyMaintenanceWindow(v *WeeklyMaintenanceWindow) {
	m.Policy = &MaintenanceWindow_WeeklyMaintenanceWindow{
		WeeklyMaintenanceWindow: v,
	}
}

func (m *WeeklyMaintenanceWindow) SetDay(v WeeklyMaintenanceWindow_WeekDay) {
	m.Day = v
}

func (m *WeeklyMaintenanceWindow) SetHour(v int64) {
	m.Hour = v
}

func (m *MaintenanceOperation) SetInfo(v string) {
	m.Info = v
}

func (m *MaintenanceOperation) SetDelayedUntil(v *timestamppb.Timestamp) {
	m.DelayedUntil = v
}
