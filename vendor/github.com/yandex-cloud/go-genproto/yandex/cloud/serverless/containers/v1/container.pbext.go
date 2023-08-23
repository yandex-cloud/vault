// Code generated by protoc-gen-goext. DO NOT EDIT.

package containers

import (
	v1 "github.com/yandex-cloud/go-genproto/yandex/cloud/logging/v1"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *Container) SetId(v string) {
	m.Id = v
}

func (m *Container) SetFolderId(v string) {
	m.FolderId = v
}

func (m *Container) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Container) SetName(v string) {
	m.Name = v
}

func (m *Container) SetDescription(v string) {
	m.Description = v
}

func (m *Container) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *Container) SetUrl(v string) {
	m.Url = v
}

func (m *Container) SetStatus(v Container_Status) {
	m.Status = v
}

func (m *Revision) SetId(v string) {
	m.Id = v
}

func (m *Revision) SetContainerId(v string) {
	m.ContainerId = v
}

func (m *Revision) SetDescription(v string) {
	m.Description = v
}

func (m *Revision) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Revision) SetImage(v *Image) {
	m.Image = v
}

func (m *Revision) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Revision) SetExecutionTimeout(v *durationpb.Duration) {
	m.ExecutionTimeout = v
}

func (m *Revision) SetConcurrency(v int64) {
	m.Concurrency = v
}

func (m *Revision) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *Revision) SetStatus(v Revision_Status) {
	m.Status = v
}

func (m *Revision) SetSecrets(v []*Secret) {
	m.Secrets = v
}

func (m *Revision) SetConnectivity(v *Connectivity) {
	m.Connectivity = v
}

func (m *Revision) SetProvisionPolicy(v *ProvisionPolicy) {
	m.ProvisionPolicy = v
}

func (m *Revision) SetScalingPolicy(v *ScalingPolicy) {
	m.ScalingPolicy = v
}

func (m *Revision) SetLogOptions(v *LogOptions) {
	m.LogOptions = v
}

func (m *Revision) SetS3Mounts(v []*S3Mount) {
	m.S3Mounts = v
}

func (m *Image) SetImageUrl(v string) {
	m.ImageUrl = v
}

func (m *Image) SetImageDigest(v string) {
	m.ImageDigest = v
}

func (m *Image) SetCommand(v *Command) {
	m.Command = v
}

func (m *Image) SetArgs(v *Args) {
	m.Args = v
}

func (m *Image) SetEnvironment(v map[string]string) {
	m.Environment = v
}

func (m *Image) SetWorkingDir(v string) {
	m.WorkingDir = v
}

func (m *Command) SetCommand(v []string) {
	m.Command = v
}

func (m *Args) SetArgs(v []string) {
	m.Args = v
}

func (m *Resources) SetMemory(v int64) {
	m.Memory = v
}

func (m *Resources) SetCores(v int64) {
	m.Cores = v
}

func (m *Resources) SetCoreFraction(v int64) {
	m.CoreFraction = v
}

func (m *ProvisionPolicy) SetMinInstances(v int64) {
	m.MinInstances = v
}

type Secret_Reference = isSecret_Reference

func (m *Secret) SetReference(v Secret_Reference) {
	m.Reference = v
}

func (m *Secret) SetId(v string) {
	m.Id = v
}

func (m *Secret) SetVersionId(v string) {
	m.VersionId = v
}

func (m *Secret) SetKey(v string) {
	m.Key = v
}

func (m *Secret) SetEnvironmentVariable(v string) {
	m.Reference = &Secret_EnvironmentVariable{
		EnvironmentVariable: v,
	}
}

func (m *Connectivity) SetNetworkId(v string) {
	m.NetworkId = v
}

func (m *Connectivity) SetSubnetIds(v []string) {
	m.SubnetIds = v
}

type LogOptions_Destination = isLogOptions_Destination

func (m *LogOptions) SetDestination(v LogOptions_Destination) {
	m.Destination = v
}

func (m *LogOptions) SetDisabled(v bool) {
	m.Disabled = v
}

func (m *LogOptions) SetLogGroupId(v string) {
	m.Destination = &LogOptions_LogGroupId{
		LogGroupId: v,
	}
}

func (m *LogOptions) SetFolderId(v string) {
	m.Destination = &LogOptions_FolderId{
		FolderId: v,
	}
}

func (m *LogOptions) SetMinLevel(v v1.LogLevel_Level) {
	m.MinLevel = v
}

func (m *ScalingPolicy) SetZoneInstancesLimit(v int64) {
	m.ZoneInstancesLimit = v
}

func (m *ScalingPolicy) SetZoneRequestsLimit(v int64) {
	m.ZoneRequestsLimit = v
}

func (m *S3Mount) SetBucketId(v string) {
	m.BucketId = v
}

func (m *S3Mount) SetPrefix(v string) {
	m.Prefix = v
}

func (m *S3Mount) SetMountPointName(v string) {
	m.MountPointName = v
}

func (m *S3Mount) SetReadOnly(v bool) {
	m.ReadOnly = v
}
