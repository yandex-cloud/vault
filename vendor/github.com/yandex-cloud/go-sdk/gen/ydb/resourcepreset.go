// Code generated by sdkgen. DO NOT EDIT.

// nolint
package ydb

import (
	"context"

	"google.golang.org/grpc"

	ydb "github.com/yandex-cloud/go-genproto/yandex/cloud/ydb/v1"
)

//revive:disable

// ResourcePresetServiceClient is a ydb.ResourcePresetServiceClient with
// lazy GRPC connection initialization.
type ResourcePresetServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Get implements ydb.ResourcePresetServiceClient
func (c *ResourcePresetServiceClient) Get(ctx context.Context, in *ydb.GetResourcePresetRequest, opts ...grpc.CallOption) (*ydb.ResourcePreset, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return ydb.NewResourcePresetServiceClient(conn).Get(ctx, in, opts...)
}

// List implements ydb.ResourcePresetServiceClient
func (c *ResourcePresetServiceClient) List(ctx context.Context, in *ydb.ListResourcePresetsRequest, opts ...grpc.CallOption) (*ydb.ListResourcePresetsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return ydb.NewResourcePresetServiceClient(conn).List(ctx, in, opts...)
}

type ResourcePresetIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ResourcePresetServiceClient
	request *ydb.ListResourcePresetsRequest

	items []*ydb.ResourcePreset
}

func (c *ResourcePresetServiceClient) ResourcePresetIterator(ctx context.Context, req *ydb.ListResourcePresetsRequest, opts ...grpc.CallOption) *ResourcePresetIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ResourcePresetIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ResourcePresetIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.ResourcePresets
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ResourcePresetIterator) Take(size int64) ([]*ydb.ResourcePreset, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*ydb.ResourcePreset

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *ResourcePresetIterator) TakeAll() ([]*ydb.ResourcePreset, error) {
	return it.Take(0)
}

func (it *ResourcePresetIterator) Value() *ydb.ResourcePreset {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ResourcePresetIterator) Error() error {
	return it.err
}
