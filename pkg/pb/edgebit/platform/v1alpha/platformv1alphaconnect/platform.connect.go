// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: edgebit/platform/v1alpha/platform.proto

package platformv1alphaconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1alpha "github.com/edgebitio/edgebit-cli/pkg/pb/edgebit/platform/v1alpha"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// EdgeBitPublicAPIServiceName is the fully-qualified name of the EdgeBitPublicAPIService service.
	EdgeBitPublicAPIServiceName = "edgebit.platform.v1alpha.EdgeBitPublicAPIService"
)

// EdgeBitPublicAPIServiceClient is a client for the
// edgebit.platform.v1alpha.EdgeBitPublicAPIService service.
type EdgeBitPublicAPIServiceClient interface {
	// Project Management (org-scoped)
	ListProjects(context.Context, *connect_go.Request[v1alpha.ListProjectsRequest]) (*connect_go.Response[v1alpha.ListProjectsResponse], error)
	// Agent Deployment Token Management (project-scoped)
	GenerateAgentDeployToken(context.Context, *connect_go.Request[v1alpha.GenerateAgentDeployTokenRequest]) (*connect_go.Response[v1alpha.GenerateAgentDeployTokenResponse], error)
	// Machine Management (project-scoped)
	ListMachines(context.Context, *connect_go.Request[v1alpha.ListMachinesRequest]) (*connect_go.Response[v1alpha.ListMachinesResponse], error)
	// Inventory Exploration (project-scoped)
	GetMachineInventory(context.Context, *connect_go.Request[v1alpha.GetMachineInventoryRequest]) (*connect_go.Response[v1alpha.GetMachineInventoryResponse], error)
	// Org (Project) Access Token Management (project-scoped)
	CreateOrgAccessToken(context.Context, *connect_go.Request[v1alpha.CreateOrgAccessTokenRequest]) (*connect_go.Response[v1alpha.CreateOrgAccessTokenResponse], error)
	ListOrgAccessTokens(context.Context, *connect_go.Request[v1alpha.ListOrgAccessTokensRequest]) (*connect_go.Response[v1alpha.ListOrgAccessTokensResponse], error)
	DeleteOrgAccessToken(context.Context, *connect_go.Request[v1alpha.DeleteOrgAccessTokenRequest]) (*connect_go.Response[v1alpha.DeleteOrgAccessTokenResponse], error)
	// SBOM Management (project-scoped)
	UploadSBOM(context.Context) *connect_go.ClientStreamForClient[v1alpha.UploadSBOMRequest, v1alpha.UploadSBOMResponse]
}

// NewEdgeBitPublicAPIServiceClient constructs a client for the
// edgebit.platform.v1alpha.EdgeBitPublicAPIService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewEdgeBitPublicAPIServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) EdgeBitPublicAPIServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &edgeBitPublicAPIServiceClient{
		listProjects: connect_go.NewClient[v1alpha.ListProjectsRequest, v1alpha.ListProjectsResponse](
			httpClient,
			baseURL+"/edgebit.platform.v1alpha.EdgeBitPublicAPIService/ListProjects",
			opts...,
		),
		generateAgentDeployToken: connect_go.NewClient[v1alpha.GenerateAgentDeployTokenRequest, v1alpha.GenerateAgentDeployTokenResponse](
			httpClient,
			baseURL+"/edgebit.platform.v1alpha.EdgeBitPublicAPIService/GenerateAgentDeployToken",
			opts...,
		),
		listMachines: connect_go.NewClient[v1alpha.ListMachinesRequest, v1alpha.ListMachinesResponse](
			httpClient,
			baseURL+"/edgebit.platform.v1alpha.EdgeBitPublicAPIService/ListMachines",
			opts...,
		),
		getMachineInventory: connect_go.NewClient[v1alpha.GetMachineInventoryRequest, v1alpha.GetMachineInventoryResponse](
			httpClient,
			baseURL+"/edgebit.platform.v1alpha.EdgeBitPublicAPIService/GetMachineInventory",
			opts...,
		),
		createOrgAccessToken: connect_go.NewClient[v1alpha.CreateOrgAccessTokenRequest, v1alpha.CreateOrgAccessTokenResponse](
			httpClient,
			baseURL+"/edgebit.platform.v1alpha.EdgeBitPublicAPIService/CreateOrgAccessToken",
			opts...,
		),
		listOrgAccessTokens: connect_go.NewClient[v1alpha.ListOrgAccessTokensRequest, v1alpha.ListOrgAccessTokensResponse](
			httpClient,
			baseURL+"/edgebit.platform.v1alpha.EdgeBitPublicAPIService/ListOrgAccessTokens",
			opts...,
		),
		deleteOrgAccessToken: connect_go.NewClient[v1alpha.DeleteOrgAccessTokenRequest, v1alpha.DeleteOrgAccessTokenResponse](
			httpClient,
			baseURL+"/edgebit.platform.v1alpha.EdgeBitPublicAPIService/DeleteOrgAccessToken",
			opts...,
		),
		uploadSBOM: connect_go.NewClient[v1alpha.UploadSBOMRequest, v1alpha.UploadSBOMResponse](
			httpClient,
			baseURL+"/edgebit.platform.v1alpha.EdgeBitPublicAPIService/UploadSBOM",
			opts...,
		),
	}
}

// edgeBitPublicAPIServiceClient implements EdgeBitPublicAPIServiceClient.
type edgeBitPublicAPIServiceClient struct {
	listProjects             *connect_go.Client[v1alpha.ListProjectsRequest, v1alpha.ListProjectsResponse]
	generateAgentDeployToken *connect_go.Client[v1alpha.GenerateAgentDeployTokenRequest, v1alpha.GenerateAgentDeployTokenResponse]
	listMachines             *connect_go.Client[v1alpha.ListMachinesRequest, v1alpha.ListMachinesResponse]
	getMachineInventory      *connect_go.Client[v1alpha.GetMachineInventoryRequest, v1alpha.GetMachineInventoryResponse]
	createOrgAccessToken     *connect_go.Client[v1alpha.CreateOrgAccessTokenRequest, v1alpha.CreateOrgAccessTokenResponse]
	listOrgAccessTokens      *connect_go.Client[v1alpha.ListOrgAccessTokensRequest, v1alpha.ListOrgAccessTokensResponse]
	deleteOrgAccessToken     *connect_go.Client[v1alpha.DeleteOrgAccessTokenRequest, v1alpha.DeleteOrgAccessTokenResponse]
	uploadSBOM               *connect_go.Client[v1alpha.UploadSBOMRequest, v1alpha.UploadSBOMResponse]
}

// ListProjects calls edgebit.platform.v1alpha.EdgeBitPublicAPIService.ListProjects.
func (c *edgeBitPublicAPIServiceClient) ListProjects(ctx context.Context, req *connect_go.Request[v1alpha.ListProjectsRequest]) (*connect_go.Response[v1alpha.ListProjectsResponse], error) {
	return c.listProjects.CallUnary(ctx, req)
}

// GenerateAgentDeployToken calls
// edgebit.platform.v1alpha.EdgeBitPublicAPIService.GenerateAgentDeployToken.
func (c *edgeBitPublicAPIServiceClient) GenerateAgentDeployToken(ctx context.Context, req *connect_go.Request[v1alpha.GenerateAgentDeployTokenRequest]) (*connect_go.Response[v1alpha.GenerateAgentDeployTokenResponse], error) {
	return c.generateAgentDeployToken.CallUnary(ctx, req)
}

// ListMachines calls edgebit.platform.v1alpha.EdgeBitPublicAPIService.ListMachines.
func (c *edgeBitPublicAPIServiceClient) ListMachines(ctx context.Context, req *connect_go.Request[v1alpha.ListMachinesRequest]) (*connect_go.Response[v1alpha.ListMachinesResponse], error) {
	return c.listMachines.CallUnary(ctx, req)
}

// GetMachineInventory calls edgebit.platform.v1alpha.EdgeBitPublicAPIService.GetMachineInventory.
func (c *edgeBitPublicAPIServiceClient) GetMachineInventory(ctx context.Context, req *connect_go.Request[v1alpha.GetMachineInventoryRequest]) (*connect_go.Response[v1alpha.GetMachineInventoryResponse], error) {
	return c.getMachineInventory.CallUnary(ctx, req)
}

// CreateOrgAccessToken calls edgebit.platform.v1alpha.EdgeBitPublicAPIService.CreateOrgAccessToken.
func (c *edgeBitPublicAPIServiceClient) CreateOrgAccessToken(ctx context.Context, req *connect_go.Request[v1alpha.CreateOrgAccessTokenRequest]) (*connect_go.Response[v1alpha.CreateOrgAccessTokenResponse], error) {
	return c.createOrgAccessToken.CallUnary(ctx, req)
}

// ListOrgAccessTokens calls edgebit.platform.v1alpha.EdgeBitPublicAPIService.ListOrgAccessTokens.
func (c *edgeBitPublicAPIServiceClient) ListOrgAccessTokens(ctx context.Context, req *connect_go.Request[v1alpha.ListOrgAccessTokensRequest]) (*connect_go.Response[v1alpha.ListOrgAccessTokensResponse], error) {
	return c.listOrgAccessTokens.CallUnary(ctx, req)
}

// DeleteOrgAccessToken calls edgebit.platform.v1alpha.EdgeBitPublicAPIService.DeleteOrgAccessToken.
func (c *edgeBitPublicAPIServiceClient) DeleteOrgAccessToken(ctx context.Context, req *connect_go.Request[v1alpha.DeleteOrgAccessTokenRequest]) (*connect_go.Response[v1alpha.DeleteOrgAccessTokenResponse], error) {
	return c.deleteOrgAccessToken.CallUnary(ctx, req)
}

// UploadSBOM calls edgebit.platform.v1alpha.EdgeBitPublicAPIService.UploadSBOM.
func (c *edgeBitPublicAPIServiceClient) UploadSBOM(ctx context.Context) *connect_go.ClientStreamForClient[v1alpha.UploadSBOMRequest, v1alpha.UploadSBOMResponse] {
	return c.uploadSBOM.CallClientStream(ctx)
}

// EdgeBitPublicAPIServiceHandler is an implementation of the
// edgebit.platform.v1alpha.EdgeBitPublicAPIService service.
type EdgeBitPublicAPIServiceHandler interface {
	// Project Management (org-scoped)
	ListProjects(context.Context, *connect_go.Request[v1alpha.ListProjectsRequest]) (*connect_go.Response[v1alpha.ListProjectsResponse], error)
	// Agent Deployment Token Management (project-scoped)
	GenerateAgentDeployToken(context.Context, *connect_go.Request[v1alpha.GenerateAgentDeployTokenRequest]) (*connect_go.Response[v1alpha.GenerateAgentDeployTokenResponse], error)
	// Machine Management (project-scoped)
	ListMachines(context.Context, *connect_go.Request[v1alpha.ListMachinesRequest]) (*connect_go.Response[v1alpha.ListMachinesResponse], error)
	// Inventory Exploration (project-scoped)
	GetMachineInventory(context.Context, *connect_go.Request[v1alpha.GetMachineInventoryRequest]) (*connect_go.Response[v1alpha.GetMachineInventoryResponse], error)
	// Org (Project) Access Token Management (project-scoped)
	CreateOrgAccessToken(context.Context, *connect_go.Request[v1alpha.CreateOrgAccessTokenRequest]) (*connect_go.Response[v1alpha.CreateOrgAccessTokenResponse], error)
	ListOrgAccessTokens(context.Context, *connect_go.Request[v1alpha.ListOrgAccessTokensRequest]) (*connect_go.Response[v1alpha.ListOrgAccessTokensResponse], error)
	DeleteOrgAccessToken(context.Context, *connect_go.Request[v1alpha.DeleteOrgAccessTokenRequest]) (*connect_go.Response[v1alpha.DeleteOrgAccessTokenResponse], error)
	// SBOM Management (project-scoped)
	UploadSBOM(context.Context, *connect_go.ClientStream[v1alpha.UploadSBOMRequest]) (*connect_go.Response[v1alpha.UploadSBOMResponse], error)
}

// NewEdgeBitPublicAPIServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewEdgeBitPublicAPIServiceHandler(svc EdgeBitPublicAPIServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/edgebit.platform.v1alpha.EdgeBitPublicAPIService/ListProjects", connect_go.NewUnaryHandler(
		"/edgebit.platform.v1alpha.EdgeBitPublicAPIService/ListProjects",
		svc.ListProjects,
		opts...,
	))
	mux.Handle("/edgebit.platform.v1alpha.EdgeBitPublicAPIService/GenerateAgentDeployToken", connect_go.NewUnaryHandler(
		"/edgebit.platform.v1alpha.EdgeBitPublicAPIService/GenerateAgentDeployToken",
		svc.GenerateAgentDeployToken,
		opts...,
	))
	mux.Handle("/edgebit.platform.v1alpha.EdgeBitPublicAPIService/ListMachines", connect_go.NewUnaryHandler(
		"/edgebit.platform.v1alpha.EdgeBitPublicAPIService/ListMachines",
		svc.ListMachines,
		opts...,
	))
	mux.Handle("/edgebit.platform.v1alpha.EdgeBitPublicAPIService/GetMachineInventory", connect_go.NewUnaryHandler(
		"/edgebit.platform.v1alpha.EdgeBitPublicAPIService/GetMachineInventory",
		svc.GetMachineInventory,
		opts...,
	))
	mux.Handle("/edgebit.platform.v1alpha.EdgeBitPublicAPIService/CreateOrgAccessToken", connect_go.NewUnaryHandler(
		"/edgebit.platform.v1alpha.EdgeBitPublicAPIService/CreateOrgAccessToken",
		svc.CreateOrgAccessToken,
		opts...,
	))
	mux.Handle("/edgebit.platform.v1alpha.EdgeBitPublicAPIService/ListOrgAccessTokens", connect_go.NewUnaryHandler(
		"/edgebit.platform.v1alpha.EdgeBitPublicAPIService/ListOrgAccessTokens",
		svc.ListOrgAccessTokens,
		opts...,
	))
	mux.Handle("/edgebit.platform.v1alpha.EdgeBitPublicAPIService/DeleteOrgAccessToken", connect_go.NewUnaryHandler(
		"/edgebit.platform.v1alpha.EdgeBitPublicAPIService/DeleteOrgAccessToken",
		svc.DeleteOrgAccessToken,
		opts...,
	))
	mux.Handle("/edgebit.platform.v1alpha.EdgeBitPublicAPIService/UploadSBOM", connect_go.NewClientStreamHandler(
		"/edgebit.platform.v1alpha.EdgeBitPublicAPIService/UploadSBOM",
		svc.UploadSBOM,
		opts...,
	))
	return "/edgebit.platform.v1alpha.EdgeBitPublicAPIService/", mux
}

// UnimplementedEdgeBitPublicAPIServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedEdgeBitPublicAPIServiceHandler struct{}

func (UnimplementedEdgeBitPublicAPIServiceHandler) ListProjects(context.Context, *connect_go.Request[v1alpha.ListProjectsRequest]) (*connect_go.Response[v1alpha.ListProjectsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("edgebit.platform.v1alpha.EdgeBitPublicAPIService.ListProjects is not implemented"))
}

func (UnimplementedEdgeBitPublicAPIServiceHandler) GenerateAgentDeployToken(context.Context, *connect_go.Request[v1alpha.GenerateAgentDeployTokenRequest]) (*connect_go.Response[v1alpha.GenerateAgentDeployTokenResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("edgebit.platform.v1alpha.EdgeBitPublicAPIService.GenerateAgentDeployToken is not implemented"))
}

func (UnimplementedEdgeBitPublicAPIServiceHandler) ListMachines(context.Context, *connect_go.Request[v1alpha.ListMachinesRequest]) (*connect_go.Response[v1alpha.ListMachinesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("edgebit.platform.v1alpha.EdgeBitPublicAPIService.ListMachines is not implemented"))
}

func (UnimplementedEdgeBitPublicAPIServiceHandler) GetMachineInventory(context.Context, *connect_go.Request[v1alpha.GetMachineInventoryRequest]) (*connect_go.Response[v1alpha.GetMachineInventoryResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("edgebit.platform.v1alpha.EdgeBitPublicAPIService.GetMachineInventory is not implemented"))
}

func (UnimplementedEdgeBitPublicAPIServiceHandler) CreateOrgAccessToken(context.Context, *connect_go.Request[v1alpha.CreateOrgAccessTokenRequest]) (*connect_go.Response[v1alpha.CreateOrgAccessTokenResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("edgebit.platform.v1alpha.EdgeBitPublicAPIService.CreateOrgAccessToken is not implemented"))
}

func (UnimplementedEdgeBitPublicAPIServiceHandler) ListOrgAccessTokens(context.Context, *connect_go.Request[v1alpha.ListOrgAccessTokensRequest]) (*connect_go.Response[v1alpha.ListOrgAccessTokensResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("edgebit.platform.v1alpha.EdgeBitPublicAPIService.ListOrgAccessTokens is not implemented"))
}

func (UnimplementedEdgeBitPublicAPIServiceHandler) DeleteOrgAccessToken(context.Context, *connect_go.Request[v1alpha.DeleteOrgAccessTokenRequest]) (*connect_go.Response[v1alpha.DeleteOrgAccessTokenResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("edgebit.platform.v1alpha.EdgeBitPublicAPIService.DeleteOrgAccessToken is not implemented"))
}

func (UnimplementedEdgeBitPublicAPIServiceHandler) UploadSBOM(context.Context, *connect_go.ClientStream[v1alpha.UploadSBOMRequest]) (*connect_go.Response[v1alpha.UploadSBOMResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("edgebit.platform.v1alpha.EdgeBitPublicAPIService.UploadSBOM is not implemented"))
}
