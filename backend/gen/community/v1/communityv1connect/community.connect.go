// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: community/v1/community.proto

package communityv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/AI1411/m_app/gen/community/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// CommunityServiceName is the fully-qualified name of the CommunityService service.
	CommunityServiceName = "community.v1.CommunityService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// CommunityServiceGetCommunityProcedure is the fully-qualified name of the CommunityService's
	// GetCommunity RPC.
	CommunityServiceGetCommunityProcedure = "/community.v1.CommunityService/GetCommunity"
	// CommunityServiceSearchCommunitiesProcedure is the fully-qualified name of the CommunityService's
	// SearchCommunities RPC.
	CommunityServiceSearchCommunitiesProcedure = "/community.v1.CommunityService/SearchCommunities"
	// CommunityServiceCreateCommunityProcedure is the fully-qualified name of the CommunityService's
	// CreateCommunity RPC.
	CommunityServiceCreateCommunityProcedure = "/community.v1.CommunityService/CreateCommunity"
	// CommunityServiceUpdateCommunityProcedure is the fully-qualified name of the CommunityService's
	// UpdateCommunity RPC.
	CommunityServiceUpdateCommunityProcedure = "/community.v1.CommunityService/UpdateCommunity"
	// CommunityServiceDeleteCommunityProcedure is the fully-qualified name of the CommunityService's
	// DeleteCommunity RPC.
	CommunityServiceDeleteCommunityProcedure = "/community.v1.CommunityService/DeleteCommunity"
)

// CommunityServiceClient is a client for the community.v1.CommunityService service.
type CommunityServiceClient interface {
	// コミュニティ情報取得
	GetCommunity(context.Context, *connect.Request[v1.GetCommunityRequest]) (*connect.Response[v1.GetCommunityResponse], error)
	// コミュニティ検索
	SearchCommunities(context.Context, *connect.Request[v1.SearchCommunitiesRequest]) (*connect.Response[v1.SearchCommunitiesResponse], error)
	// コミュニティ作成
	CreateCommunity(context.Context, *connect.Request[v1.CreateCommunityRequest]) (*connect.Response[v1.CreateCommunityResponse], error)
	// コミュニティ更新
	UpdateCommunity(context.Context, *connect.Request[v1.UpdateCommunityRequest]) (*connect.Response[v1.UpdateCommunityResponse], error)
	// コミュニティ削除
	DeleteCommunity(context.Context, *connect.Request[v1.DeleteCommunityRequest]) (*connect.Response[v1.DeleteCommunityResponse], error)
}

// NewCommunityServiceClient constructs a client for the community.v1.CommunityService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCommunityServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) CommunityServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	communityServiceMethods := v1.File_community_v1_community_proto.Services().ByName("CommunityService").Methods()
	return &communityServiceClient{
		getCommunity: connect.NewClient[v1.GetCommunityRequest, v1.GetCommunityResponse](
			httpClient,
			baseURL+CommunityServiceGetCommunityProcedure,
			connect.WithSchema(communityServiceMethods.ByName("GetCommunity")),
			connect.WithClientOptions(opts...),
		),
		searchCommunities: connect.NewClient[v1.SearchCommunitiesRequest, v1.SearchCommunitiesResponse](
			httpClient,
			baseURL+CommunityServiceSearchCommunitiesProcedure,
			connect.WithSchema(communityServiceMethods.ByName("SearchCommunities")),
			connect.WithClientOptions(opts...),
		),
		createCommunity: connect.NewClient[v1.CreateCommunityRequest, v1.CreateCommunityResponse](
			httpClient,
			baseURL+CommunityServiceCreateCommunityProcedure,
			connect.WithSchema(communityServiceMethods.ByName("CreateCommunity")),
			connect.WithClientOptions(opts...),
		),
		updateCommunity: connect.NewClient[v1.UpdateCommunityRequest, v1.UpdateCommunityResponse](
			httpClient,
			baseURL+CommunityServiceUpdateCommunityProcedure,
			connect.WithSchema(communityServiceMethods.ByName("UpdateCommunity")),
			connect.WithClientOptions(opts...),
		),
		deleteCommunity: connect.NewClient[v1.DeleteCommunityRequest, v1.DeleteCommunityResponse](
			httpClient,
			baseURL+CommunityServiceDeleteCommunityProcedure,
			connect.WithSchema(communityServiceMethods.ByName("DeleteCommunity")),
			connect.WithClientOptions(opts...),
		),
	}
}

// communityServiceClient implements CommunityServiceClient.
type communityServiceClient struct {
	getCommunity      *connect.Client[v1.GetCommunityRequest, v1.GetCommunityResponse]
	searchCommunities *connect.Client[v1.SearchCommunitiesRequest, v1.SearchCommunitiesResponse]
	createCommunity   *connect.Client[v1.CreateCommunityRequest, v1.CreateCommunityResponse]
	updateCommunity   *connect.Client[v1.UpdateCommunityRequest, v1.UpdateCommunityResponse]
	deleteCommunity   *connect.Client[v1.DeleteCommunityRequest, v1.DeleteCommunityResponse]
}

// GetCommunity calls community.v1.CommunityService.GetCommunity.
func (c *communityServiceClient) GetCommunity(ctx context.Context, req *connect.Request[v1.GetCommunityRequest]) (*connect.Response[v1.GetCommunityResponse], error) {
	return c.getCommunity.CallUnary(ctx, req)
}

// SearchCommunities calls community.v1.CommunityService.SearchCommunities.
func (c *communityServiceClient) SearchCommunities(ctx context.Context, req *connect.Request[v1.SearchCommunitiesRequest]) (*connect.Response[v1.SearchCommunitiesResponse], error) {
	return c.searchCommunities.CallUnary(ctx, req)
}

// CreateCommunity calls community.v1.CommunityService.CreateCommunity.
func (c *communityServiceClient) CreateCommunity(ctx context.Context, req *connect.Request[v1.CreateCommunityRequest]) (*connect.Response[v1.CreateCommunityResponse], error) {
	return c.createCommunity.CallUnary(ctx, req)
}

// UpdateCommunity calls community.v1.CommunityService.UpdateCommunity.
func (c *communityServiceClient) UpdateCommunity(ctx context.Context, req *connect.Request[v1.UpdateCommunityRequest]) (*connect.Response[v1.UpdateCommunityResponse], error) {
	return c.updateCommunity.CallUnary(ctx, req)
}

// DeleteCommunity calls community.v1.CommunityService.DeleteCommunity.
func (c *communityServiceClient) DeleteCommunity(ctx context.Context, req *connect.Request[v1.DeleteCommunityRequest]) (*connect.Response[v1.DeleteCommunityResponse], error) {
	return c.deleteCommunity.CallUnary(ctx, req)
}

// CommunityServiceHandler is an implementation of the community.v1.CommunityService service.
type CommunityServiceHandler interface {
	// コミュニティ情報取得
	GetCommunity(context.Context, *connect.Request[v1.GetCommunityRequest]) (*connect.Response[v1.GetCommunityResponse], error)
	// コミュニティ検索
	SearchCommunities(context.Context, *connect.Request[v1.SearchCommunitiesRequest]) (*connect.Response[v1.SearchCommunitiesResponse], error)
	// コミュニティ作成
	CreateCommunity(context.Context, *connect.Request[v1.CreateCommunityRequest]) (*connect.Response[v1.CreateCommunityResponse], error)
	// コミュニティ更新
	UpdateCommunity(context.Context, *connect.Request[v1.UpdateCommunityRequest]) (*connect.Response[v1.UpdateCommunityResponse], error)
	// コミュニティ削除
	DeleteCommunity(context.Context, *connect.Request[v1.DeleteCommunityRequest]) (*connect.Response[v1.DeleteCommunityResponse], error)
}

// NewCommunityServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCommunityServiceHandler(svc CommunityServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	communityServiceMethods := v1.File_community_v1_community_proto.Services().ByName("CommunityService").Methods()
	communityServiceGetCommunityHandler := connect.NewUnaryHandler(
		CommunityServiceGetCommunityProcedure,
		svc.GetCommunity,
		connect.WithSchema(communityServiceMethods.ByName("GetCommunity")),
		connect.WithHandlerOptions(opts...),
	)
	communityServiceSearchCommunitiesHandler := connect.NewUnaryHandler(
		CommunityServiceSearchCommunitiesProcedure,
		svc.SearchCommunities,
		connect.WithSchema(communityServiceMethods.ByName("SearchCommunities")),
		connect.WithHandlerOptions(opts...),
	)
	communityServiceCreateCommunityHandler := connect.NewUnaryHandler(
		CommunityServiceCreateCommunityProcedure,
		svc.CreateCommunity,
		connect.WithSchema(communityServiceMethods.ByName("CreateCommunity")),
		connect.WithHandlerOptions(opts...),
	)
	communityServiceUpdateCommunityHandler := connect.NewUnaryHandler(
		CommunityServiceUpdateCommunityProcedure,
		svc.UpdateCommunity,
		connect.WithSchema(communityServiceMethods.ByName("UpdateCommunity")),
		connect.WithHandlerOptions(opts...),
	)
	communityServiceDeleteCommunityHandler := connect.NewUnaryHandler(
		CommunityServiceDeleteCommunityProcedure,
		svc.DeleteCommunity,
		connect.WithSchema(communityServiceMethods.ByName("DeleteCommunity")),
		connect.WithHandlerOptions(opts...),
	)
	return "/community.v1.CommunityService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case CommunityServiceGetCommunityProcedure:
			communityServiceGetCommunityHandler.ServeHTTP(w, r)
		case CommunityServiceSearchCommunitiesProcedure:
			communityServiceSearchCommunitiesHandler.ServeHTTP(w, r)
		case CommunityServiceCreateCommunityProcedure:
			communityServiceCreateCommunityHandler.ServeHTTP(w, r)
		case CommunityServiceUpdateCommunityProcedure:
			communityServiceUpdateCommunityHandler.ServeHTTP(w, r)
		case CommunityServiceDeleteCommunityProcedure:
			communityServiceDeleteCommunityHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedCommunityServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedCommunityServiceHandler struct{}

func (UnimplementedCommunityServiceHandler) GetCommunity(context.Context, *connect.Request[v1.GetCommunityRequest]) (*connect.Response[v1.GetCommunityResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("community.v1.CommunityService.GetCommunity is not implemented"))
}

func (UnimplementedCommunityServiceHandler) SearchCommunities(context.Context, *connect.Request[v1.SearchCommunitiesRequest]) (*connect.Response[v1.SearchCommunitiesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("community.v1.CommunityService.SearchCommunities is not implemented"))
}

func (UnimplementedCommunityServiceHandler) CreateCommunity(context.Context, *connect.Request[v1.CreateCommunityRequest]) (*connect.Response[v1.CreateCommunityResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("community.v1.CommunityService.CreateCommunity is not implemented"))
}

func (UnimplementedCommunityServiceHandler) UpdateCommunity(context.Context, *connect.Request[v1.UpdateCommunityRequest]) (*connect.Response[v1.UpdateCommunityResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("community.v1.CommunityService.UpdateCommunity is not implemented"))
}

func (UnimplementedCommunityServiceHandler) DeleteCommunity(context.Context, *connect.Request[v1.DeleteCommunityRequest]) (*connect.Response[v1.DeleteCommunityResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("community.v1.CommunityService.DeleteCommunity is not implemented"))
}
