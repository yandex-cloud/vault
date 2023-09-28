// Code generated by sdkgen. DO NOT EDIT.

// nolint
package iam

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	iam "github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// ServiceAccountServiceClient is a iam.ServiceAccountServiceClient with
// lazy GRPC connection initialization.
type ServiceAccountServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements iam.ServiceAccountServiceClient
func (c *ServiceAccountServiceClient) Create(ctx context.Context, in *iam.CreateServiceAccountRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewServiceAccountServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements iam.ServiceAccountServiceClient
func (c *ServiceAccountServiceClient) Delete(ctx context.Context, in *iam.DeleteServiceAccountRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewServiceAccountServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements iam.ServiceAccountServiceClient
func (c *ServiceAccountServiceClient) Get(ctx context.Context, in *iam.GetServiceAccountRequest, opts ...grpc.CallOption) (*iam.ServiceAccount, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewServiceAccountServiceClient(conn).Get(ctx, in, opts...)
}

// List implements iam.ServiceAccountServiceClient
func (c *ServiceAccountServiceClient) List(ctx context.Context, in *iam.ListServiceAccountsRequest, opts ...grpc.CallOption) (*iam.ListServiceAccountsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewServiceAccountServiceClient(conn).List(ctx, in, opts...)
}

type ServiceAccountIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ServiceAccountServiceClient
	request *iam.ListServiceAccountsRequest

	items []*iam.ServiceAccount
}

func (c *ServiceAccountServiceClient) ServiceAccountIterator(ctx context.Context, req *iam.ListServiceAccountsRequest, opts ...grpc.CallOption) *ServiceAccountIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ServiceAccountIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ServiceAccountIterator) Next() bool {
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

	it.items = response.ServiceAccounts
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ServiceAccountIterator) Take(size int64) ([]*iam.ServiceAccount, error) {
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

	var result []*iam.ServiceAccount

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *ServiceAccountIterator) TakeAll() ([]*iam.ServiceAccount, error) {
	return it.Take(0)
}

func (it *ServiceAccountIterator) Value() *iam.ServiceAccount {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ServiceAccountIterator) Error() error {
	return it.err
}

// ListAccessBindings implements iam.ServiceAccountServiceClient
func (c *ServiceAccountServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewServiceAccountServiceClient(conn).ListAccessBindings(ctx, in, opts...)
}

type ServiceAccountAccessBindingsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ServiceAccountServiceClient
	request *access.ListAccessBindingsRequest

	items []*access.AccessBinding
}

func (c *ServiceAccountServiceClient) ServiceAccountAccessBindingsIterator(ctx context.Context, req *access.ListAccessBindingsRequest, opts ...grpc.CallOption) *ServiceAccountAccessBindingsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ServiceAccountAccessBindingsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ServiceAccountAccessBindingsIterator) Next() bool {
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

	response, err := it.client.ListAccessBindings(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.AccessBindings
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ServiceAccountAccessBindingsIterator) Take(size int64) ([]*access.AccessBinding, error) {
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

	var result []*access.AccessBinding

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *ServiceAccountAccessBindingsIterator) TakeAll() ([]*access.AccessBinding, error) {
	return it.Take(0)
}

func (it *ServiceAccountAccessBindingsIterator) Value() *access.AccessBinding {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ServiceAccountAccessBindingsIterator) Error() error {
	return it.err
}

// ListOperations implements iam.ServiceAccountServiceClient
func (c *ServiceAccountServiceClient) ListOperations(ctx context.Context, in *iam.ListServiceAccountOperationsRequest, opts ...grpc.CallOption) (*iam.ListServiceAccountOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewServiceAccountServiceClient(conn).ListOperations(ctx, in, opts...)
}

type ServiceAccountOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ServiceAccountServiceClient
	request *iam.ListServiceAccountOperationsRequest

	items []*operation.Operation
}

func (c *ServiceAccountServiceClient) ServiceAccountOperationsIterator(ctx context.Context, req *iam.ListServiceAccountOperationsRequest, opts ...grpc.CallOption) *ServiceAccountOperationsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ServiceAccountOperationsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ServiceAccountOperationsIterator) Next() bool {
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

func (it *ServiceAccountOperationsIterator) Take(size int64) ([]*operation.Operation, error) {
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

func (it *ServiceAccountOperationsIterator) TakeAll() ([]*operation.Operation, error) {
	return it.Take(0)
}

func (it *ServiceAccountOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ServiceAccountOperationsIterator) Error() error {
	return it.err
}

// SetAccessBindings implements iam.ServiceAccountServiceClient
func (c *ServiceAccountServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewServiceAccountServiceClient(conn).SetAccessBindings(ctx, in, opts...)
}

// Update implements iam.ServiceAccountServiceClient
func (c *ServiceAccountServiceClient) Update(ctx context.Context, in *iam.UpdateServiceAccountRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewServiceAccountServiceClient(conn).Update(ctx, in, opts...)
}

// UpdateAccessBindings implements iam.ServiceAccountServiceClient
func (c *ServiceAccountServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return iam.NewServiceAccountServiceClient(conn).UpdateAccessBindings(ctx, in, opts...)
}
