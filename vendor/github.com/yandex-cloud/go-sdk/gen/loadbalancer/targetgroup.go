// Code generated by sdkgen. DO NOT EDIT.

// nolint
package loadbalancer

import (
	"context"

	"google.golang.org/grpc"

	loadbalancer "github.com/yandex-cloud/go-genproto/yandex/cloud/loadbalancer/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// TargetGroupServiceClient is a loadbalancer.TargetGroupServiceClient with
// lazy GRPC connection initialization.
type TargetGroupServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// AddTargets implements loadbalancer.TargetGroupServiceClient
func (c *TargetGroupServiceClient) AddTargets(ctx context.Context, in *loadbalancer.AddTargetsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return loadbalancer.NewTargetGroupServiceClient(conn).AddTargets(ctx, in, opts...)
}

// Create implements loadbalancer.TargetGroupServiceClient
func (c *TargetGroupServiceClient) Create(ctx context.Context, in *loadbalancer.CreateTargetGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return loadbalancer.NewTargetGroupServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements loadbalancer.TargetGroupServiceClient
func (c *TargetGroupServiceClient) Delete(ctx context.Context, in *loadbalancer.DeleteTargetGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return loadbalancer.NewTargetGroupServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements loadbalancer.TargetGroupServiceClient
func (c *TargetGroupServiceClient) Get(ctx context.Context, in *loadbalancer.GetTargetGroupRequest, opts ...grpc.CallOption) (*loadbalancer.TargetGroup, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return loadbalancer.NewTargetGroupServiceClient(conn).Get(ctx, in, opts...)
}

// List implements loadbalancer.TargetGroupServiceClient
func (c *TargetGroupServiceClient) List(ctx context.Context, in *loadbalancer.ListTargetGroupsRequest, opts ...grpc.CallOption) (*loadbalancer.ListTargetGroupsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return loadbalancer.NewTargetGroupServiceClient(conn).List(ctx, in, opts...)
}

type TargetGroupIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *TargetGroupServiceClient
	request *loadbalancer.ListTargetGroupsRequest

	items []*loadbalancer.TargetGroup
}

func (c *TargetGroupServiceClient) TargetGroupIterator(ctx context.Context, req *loadbalancer.ListTargetGroupsRequest, opts ...grpc.CallOption) *TargetGroupIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &TargetGroupIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *TargetGroupIterator) Next() bool {
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

	it.items = response.TargetGroups
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *TargetGroupIterator) Take(size int64) ([]*loadbalancer.TargetGroup, error) {
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

	var result []*loadbalancer.TargetGroup

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *TargetGroupIterator) TakeAll() ([]*loadbalancer.TargetGroup, error) {
	return it.Take(0)
}

func (it *TargetGroupIterator) Value() *loadbalancer.TargetGroup {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *TargetGroupIterator) Error() error {
	return it.err
}

// ListOperations implements loadbalancer.TargetGroupServiceClient
func (c *TargetGroupServiceClient) ListOperations(ctx context.Context, in *loadbalancer.ListTargetGroupOperationsRequest, opts ...grpc.CallOption) (*loadbalancer.ListTargetGroupOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return loadbalancer.NewTargetGroupServiceClient(conn).ListOperations(ctx, in, opts...)
}

type TargetGroupOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *TargetGroupServiceClient
	request *loadbalancer.ListTargetGroupOperationsRequest

	items []*operation.Operation
}

func (c *TargetGroupServiceClient) TargetGroupOperationsIterator(ctx context.Context, req *loadbalancer.ListTargetGroupOperationsRequest, opts ...grpc.CallOption) *TargetGroupOperationsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &TargetGroupOperationsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *TargetGroupOperationsIterator) Next() bool {
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

	response, err := it.client.ListOperations(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Operations
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *TargetGroupOperationsIterator) Take(size int64) ([]*operation.Operation, error) {
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

	var result []*operation.Operation

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *TargetGroupOperationsIterator) TakeAll() ([]*operation.Operation, error) {
	return it.Take(0)
}

func (it *TargetGroupOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *TargetGroupOperationsIterator) Error() error {
	return it.err
}

// RemoveTargets implements loadbalancer.TargetGroupServiceClient
func (c *TargetGroupServiceClient) RemoveTargets(ctx context.Context, in *loadbalancer.RemoveTargetsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return loadbalancer.NewTargetGroupServiceClient(conn).RemoveTargets(ctx, in, opts...)
}

// Update implements loadbalancer.TargetGroupServiceClient
func (c *TargetGroupServiceClient) Update(ctx context.Context, in *loadbalancer.UpdateTargetGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return loadbalancer.NewTargetGroupServiceClient(conn).Update(ctx, in, opts...)
}
