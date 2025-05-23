// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: tweet/v1/tweet.proto

package tweetv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/AI1411/m_app/gen/tweet/v1"
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
	// TweetServiceName is the fully-qualified name of the TweetService service.
	TweetServiceName = "tweet.v1.TweetService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TweetServiceCreateTweetProcedure is the fully-qualified name of the TweetService's CreateTweet
	// RPC.
	TweetServiceCreateTweetProcedure = "/tweet.v1.TweetService/CreateTweet"
	// TweetServiceGetTweetProcedure is the fully-qualified name of the TweetService's GetTweet RPC.
	TweetServiceGetTweetProcedure = "/tweet.v1.TweetService/GetTweet"
	// TweetServiceListTweetsProcedure is the fully-qualified name of the TweetService's ListTweets RPC.
	TweetServiceListTweetsProcedure = "/tweet.v1.TweetService/ListTweets"
	// TweetServiceUpdateTweetProcedure is the fully-qualified name of the TweetService's UpdateTweet
	// RPC.
	TweetServiceUpdateTweetProcedure = "/tweet.v1.TweetService/UpdateTweet"
	// TweetServiceDeleteTweetProcedure is the fully-qualified name of the TweetService's DeleteTweet
	// RPC.
	TweetServiceDeleteTweetProcedure = "/tweet.v1.TweetService/DeleteTweet"
)

// TweetServiceClient is a client for the tweet.v1.TweetService service.
type TweetServiceClient interface {
	// つぶやき作成
	CreateTweet(context.Context, *connect.Request[v1.CreateTweetRequest]) (*connect.Response[v1.CreateTweetResponse], error)
	// つぶやき取得
	GetTweet(context.Context, *connect.Request[v1.GetTweetRequest]) (*connect.Response[v1.GetTweetResponse], error)
	// つぶやき一覧取得
	ListTweets(context.Context, *connect.Request[v1.ListTweetsRequest]) (*connect.Response[v1.ListTweetsResponse], error)
	// つぶやき更新
	UpdateTweet(context.Context, *connect.Request[v1.UpdateTweetRequest]) (*connect.Response[v1.UpdateTweetResponse], error)
	// つぶやき削除
	DeleteTweet(context.Context, *connect.Request[v1.DeleteTweetRequest]) (*connect.Response[v1.DeleteTweetResponse], error)
}

// NewTweetServiceClient constructs a client for the tweet.v1.TweetService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTweetServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TweetServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	tweetServiceMethods := v1.File_tweet_v1_tweet_proto.Services().ByName("TweetService").Methods()
	return &tweetServiceClient{
		createTweet: connect.NewClient[v1.CreateTweetRequest, v1.CreateTweetResponse](
			httpClient,
			baseURL+TweetServiceCreateTweetProcedure,
			connect.WithSchema(tweetServiceMethods.ByName("CreateTweet")),
			connect.WithClientOptions(opts...),
		),
		getTweet: connect.NewClient[v1.GetTweetRequest, v1.GetTweetResponse](
			httpClient,
			baseURL+TweetServiceGetTweetProcedure,
			connect.WithSchema(tweetServiceMethods.ByName("GetTweet")),
			connect.WithClientOptions(opts...),
		),
		listTweets: connect.NewClient[v1.ListTweetsRequest, v1.ListTweetsResponse](
			httpClient,
			baseURL+TweetServiceListTweetsProcedure,
			connect.WithSchema(tweetServiceMethods.ByName("ListTweets")),
			connect.WithClientOptions(opts...),
		),
		updateTweet: connect.NewClient[v1.UpdateTweetRequest, v1.UpdateTweetResponse](
			httpClient,
			baseURL+TweetServiceUpdateTweetProcedure,
			connect.WithSchema(tweetServiceMethods.ByName("UpdateTweet")),
			connect.WithClientOptions(opts...),
		),
		deleteTweet: connect.NewClient[v1.DeleteTweetRequest, v1.DeleteTweetResponse](
			httpClient,
			baseURL+TweetServiceDeleteTweetProcedure,
			connect.WithSchema(tweetServiceMethods.ByName("DeleteTweet")),
			connect.WithClientOptions(opts...),
		),
	}
}

// tweetServiceClient implements TweetServiceClient.
type tweetServiceClient struct {
	createTweet *connect.Client[v1.CreateTweetRequest, v1.CreateTweetResponse]
	getTweet    *connect.Client[v1.GetTweetRequest, v1.GetTweetResponse]
	listTweets  *connect.Client[v1.ListTweetsRequest, v1.ListTweetsResponse]
	updateTweet *connect.Client[v1.UpdateTweetRequest, v1.UpdateTweetResponse]
	deleteTweet *connect.Client[v1.DeleteTweetRequest, v1.DeleteTweetResponse]
}

// CreateTweet calls tweet.v1.TweetService.CreateTweet.
func (c *tweetServiceClient) CreateTweet(ctx context.Context, req *connect.Request[v1.CreateTweetRequest]) (*connect.Response[v1.CreateTweetResponse], error) {
	return c.createTweet.CallUnary(ctx, req)
}

// GetTweet calls tweet.v1.TweetService.GetTweet.
func (c *tweetServiceClient) GetTweet(ctx context.Context, req *connect.Request[v1.GetTweetRequest]) (*connect.Response[v1.GetTweetResponse], error) {
	return c.getTweet.CallUnary(ctx, req)
}

// ListTweets calls tweet.v1.TweetService.ListTweets.
func (c *tweetServiceClient) ListTweets(ctx context.Context, req *connect.Request[v1.ListTweetsRequest]) (*connect.Response[v1.ListTweetsResponse], error) {
	return c.listTweets.CallUnary(ctx, req)
}

// UpdateTweet calls tweet.v1.TweetService.UpdateTweet.
func (c *tweetServiceClient) UpdateTweet(ctx context.Context, req *connect.Request[v1.UpdateTweetRequest]) (*connect.Response[v1.UpdateTweetResponse], error) {
	return c.updateTweet.CallUnary(ctx, req)
}

// DeleteTweet calls tweet.v1.TweetService.DeleteTweet.
func (c *tweetServiceClient) DeleteTweet(ctx context.Context, req *connect.Request[v1.DeleteTweetRequest]) (*connect.Response[v1.DeleteTweetResponse], error) {
	return c.deleteTweet.CallUnary(ctx, req)
}

// TweetServiceHandler is an implementation of the tweet.v1.TweetService service.
type TweetServiceHandler interface {
	// つぶやき作成
	CreateTweet(context.Context, *connect.Request[v1.CreateTweetRequest]) (*connect.Response[v1.CreateTweetResponse], error)
	// つぶやき取得
	GetTweet(context.Context, *connect.Request[v1.GetTweetRequest]) (*connect.Response[v1.GetTweetResponse], error)
	// つぶやき一覧取得
	ListTweets(context.Context, *connect.Request[v1.ListTweetsRequest]) (*connect.Response[v1.ListTweetsResponse], error)
	// つぶやき更新
	UpdateTweet(context.Context, *connect.Request[v1.UpdateTweetRequest]) (*connect.Response[v1.UpdateTweetResponse], error)
	// つぶやき削除
	DeleteTweet(context.Context, *connect.Request[v1.DeleteTweetRequest]) (*connect.Response[v1.DeleteTweetResponse], error)
}

// NewTweetServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTweetServiceHandler(svc TweetServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	tweetServiceMethods := v1.File_tweet_v1_tweet_proto.Services().ByName("TweetService").Methods()
	tweetServiceCreateTweetHandler := connect.NewUnaryHandler(
		TweetServiceCreateTweetProcedure,
		svc.CreateTweet,
		connect.WithSchema(tweetServiceMethods.ByName("CreateTweet")),
		connect.WithHandlerOptions(opts...),
	)
	tweetServiceGetTweetHandler := connect.NewUnaryHandler(
		TweetServiceGetTweetProcedure,
		svc.GetTweet,
		connect.WithSchema(tweetServiceMethods.ByName("GetTweet")),
		connect.WithHandlerOptions(opts...),
	)
	tweetServiceListTweetsHandler := connect.NewUnaryHandler(
		TweetServiceListTweetsProcedure,
		svc.ListTweets,
		connect.WithSchema(tweetServiceMethods.ByName("ListTweets")),
		connect.WithHandlerOptions(opts...),
	)
	tweetServiceUpdateTweetHandler := connect.NewUnaryHandler(
		TweetServiceUpdateTweetProcedure,
		svc.UpdateTweet,
		connect.WithSchema(tweetServiceMethods.ByName("UpdateTweet")),
		connect.WithHandlerOptions(opts...),
	)
	tweetServiceDeleteTweetHandler := connect.NewUnaryHandler(
		TweetServiceDeleteTweetProcedure,
		svc.DeleteTweet,
		connect.WithSchema(tweetServiceMethods.ByName("DeleteTweet")),
		connect.WithHandlerOptions(opts...),
	)
	return "/tweet.v1.TweetService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TweetServiceCreateTweetProcedure:
			tweetServiceCreateTweetHandler.ServeHTTP(w, r)
		case TweetServiceGetTweetProcedure:
			tweetServiceGetTweetHandler.ServeHTTP(w, r)
		case TweetServiceListTweetsProcedure:
			tweetServiceListTweetsHandler.ServeHTTP(w, r)
		case TweetServiceUpdateTweetProcedure:
			tweetServiceUpdateTweetHandler.ServeHTTP(w, r)
		case TweetServiceDeleteTweetProcedure:
			tweetServiceDeleteTweetHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTweetServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTweetServiceHandler struct{}

func (UnimplementedTweetServiceHandler) CreateTweet(context.Context, *connect.Request[v1.CreateTweetRequest]) (*connect.Response[v1.CreateTweetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("tweet.v1.TweetService.CreateTweet is not implemented"))
}

func (UnimplementedTweetServiceHandler) GetTweet(context.Context, *connect.Request[v1.GetTweetRequest]) (*connect.Response[v1.GetTweetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("tweet.v1.TweetService.GetTweet is not implemented"))
}

func (UnimplementedTweetServiceHandler) ListTweets(context.Context, *connect.Request[v1.ListTweetsRequest]) (*connect.Response[v1.ListTweetsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("tweet.v1.TweetService.ListTweets is not implemented"))
}

func (UnimplementedTweetServiceHandler) UpdateTweet(context.Context, *connect.Request[v1.UpdateTweetRequest]) (*connect.Response[v1.UpdateTweetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("tweet.v1.TweetService.UpdateTweet is not implemented"))
}

func (UnimplementedTweetServiceHandler) DeleteTweet(context.Context, *connect.Request[v1.DeleteTweetRequest]) (*connect.Response[v1.DeleteTweetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("tweet.v1.TweetService.DeleteTweet is not implemented"))
}
