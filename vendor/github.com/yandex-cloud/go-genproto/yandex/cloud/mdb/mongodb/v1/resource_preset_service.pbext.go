// Code generated by protoc-gen-goext. DO NOT EDIT.

package mongodb

func (m *GetResourcePresetRequest) SetResourcePresetId(v string) {
	m.ResourcePresetId = v
}

func (m *ListResourcePresetsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListResourcePresetsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListResourcePresetsResponse) SetResourcePresets(v []*ResourcePreset) {
	m.ResourcePresets = v
}

func (m *ListResourcePresetsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
