// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: tag/v1/service_internal.proto

package tagconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/core-pb/tag/tag/v1"
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
	// InternalName is the fully-qualified name of the Internal service.
	InternalName = "tag.v1.Internal"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// InternalGetTagIDTreeizeProcedure is the fully-qualified name of the Internal's GetTagIDTreeize
	// RPC.
	InternalGetTagIDTreeizeProcedure = "/tag.v1.Internal/GetTagIDTreeize"
	// InternalBindRelationProcedure is the fully-qualified name of the Internal's BindRelation RPC.
	InternalBindRelationProcedure = "/tag.v1.Internal/BindRelation"
	// InternalUnbindRelationProcedure is the fully-qualified name of the Internal's UnbindRelation RPC.
	InternalUnbindRelationProcedure = "/tag.v1.Internal/UnbindRelation"
	// InternalGetAllByModuleProcedure is the fully-qualified name of the Internal's GetAllByModule RPC.
	InternalGetAllByModuleProcedure = "/tag.v1.Internal/GetAllByModule"
	// InternalRegisterModuleProcedure is the fully-qualified name of the Internal's RegisterModule RPC.
	InternalRegisterModuleProcedure = "/tag.v1.Internal/RegisterModule"
	// InternalRegisterTagProcedure is the fully-qualified name of the Internal's RegisterTag RPC.
	InternalRegisterTagProcedure = "/tag.v1.Internal/RegisterTag"
	// InternalSetTypeWithModuleProcedure is the fully-qualified name of the Internal's
	// SetTypeWithModule RPC.
	InternalSetTypeWithModuleProcedure = "/tag.v1.Internal/SetTypeWithModule"
	// InternalDeleteTypeWithModuleProcedure is the fully-qualified name of the Internal's
	// DeleteTypeWithModule RPC.
	InternalDeleteTypeWithModuleProcedure = "/tag.v1.Internal/DeleteTypeWithModule"
	// InternalSetTagWithModuleProcedure is the fully-qualified name of the Internal's SetTagWithModule
	// RPC.
	InternalSetTagWithModuleProcedure = "/tag.v1.Internal/SetTagWithModule"
	// InternalDeleteTagWithModuleProcedure is the fully-qualified name of the Internal's
	// DeleteTagWithModule RPC.
	InternalDeleteTagWithModuleProcedure = "/tag.v1.Internal/DeleteTagWithModule"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	internalServiceDescriptor                    = v1.File_tag_v1_service_internal_proto.Services().ByName("Internal")
	internalGetTagIDTreeizeMethodDescriptor      = internalServiceDescriptor.Methods().ByName("GetTagIDTreeize")
	internalBindRelationMethodDescriptor         = internalServiceDescriptor.Methods().ByName("BindRelation")
	internalUnbindRelationMethodDescriptor       = internalServiceDescriptor.Methods().ByName("UnbindRelation")
	internalGetAllByModuleMethodDescriptor       = internalServiceDescriptor.Methods().ByName("GetAllByModule")
	internalRegisterModuleMethodDescriptor       = internalServiceDescriptor.Methods().ByName("RegisterModule")
	internalRegisterTagMethodDescriptor          = internalServiceDescriptor.Methods().ByName("RegisterTag")
	internalSetTypeWithModuleMethodDescriptor    = internalServiceDescriptor.Methods().ByName("SetTypeWithModule")
	internalDeleteTypeWithModuleMethodDescriptor = internalServiceDescriptor.Methods().ByName("DeleteTypeWithModule")
	internalSetTagWithModuleMethodDescriptor     = internalServiceDescriptor.Methods().ByName("SetTagWithModule")
	internalDeleteTagWithModuleMethodDescriptor  = internalServiceDescriptor.Methods().ByName("DeleteTagWithModule")
)

// InternalClient is a client for the tag.v1.Internal service.
type InternalClient interface {
	GetTagIDTreeize(context.Context, *connect.Request[v1.GetTagIDTreeizeRequest]) (*connect.Response[v1.GetTagIDTreeizeResponse], error)
	BindRelation(context.Context, *connect.Request[v1.BindRelationRequest]) (*connect.Response[v1.BindRelationResponse], error)
	UnbindRelation(context.Context, *connect.Request[v1.UnbindRelationRequest]) (*connect.Response[v1.UnbindRelationResponse], error)
	GetAllByModule(context.Context, *connect.Request[v1.GetAllByModuleRequest]) (*connect.Response[v1.GetAllByModuleResponse], error)
	RegisterModule(context.Context, *connect.Request[v1.RegisterModuleRequest]) (*connect.Response[v1.RegisterModuleResponse], error)
	RegisterTag(context.Context, *connect.Request[v1.RegisterTagRequest]) (*connect.Response[v1.RegisterTagResponse], error)
	SetTypeWithModule(context.Context, *connect.Request[v1.SetTypeWithModuleRequest]) (*connect.Response[v1.SetTypeWithModuleResponse], error)
	DeleteTypeWithModule(context.Context, *connect.Request[v1.DeleteTypeWithModuleRequest]) (*connect.Response[v1.DeleteTypeWithModuleResponse], error)
	SetTagWithModule(context.Context, *connect.Request[v1.SetTagWithModuleRequest]) (*connect.Response[v1.SetTagWithModuleResponse], error)
	DeleteTagWithModule(context.Context, *connect.Request[v1.DeleteTagWithModuleRequest]) (*connect.Response[v1.DeleteTagWithModuleResponse], error)
}

// NewInternalClient constructs a client for the tag.v1.Internal service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewInternalClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) InternalClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &internalClient{
		getTagIDTreeize: connect.NewClient[v1.GetTagIDTreeizeRequest, v1.GetTagIDTreeizeResponse](
			httpClient,
			baseURL+InternalGetTagIDTreeizeProcedure,
			connect.WithSchema(internalGetTagIDTreeizeMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		bindRelation: connect.NewClient[v1.BindRelationRequest, v1.BindRelationResponse](
			httpClient,
			baseURL+InternalBindRelationProcedure,
			connect.WithSchema(internalBindRelationMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		unbindRelation: connect.NewClient[v1.UnbindRelationRequest, v1.UnbindRelationResponse](
			httpClient,
			baseURL+InternalUnbindRelationProcedure,
			connect.WithSchema(internalUnbindRelationMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getAllByModule: connect.NewClient[v1.GetAllByModuleRequest, v1.GetAllByModuleResponse](
			httpClient,
			baseURL+InternalGetAllByModuleProcedure,
			connect.WithSchema(internalGetAllByModuleMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		registerModule: connect.NewClient[v1.RegisterModuleRequest, v1.RegisterModuleResponse](
			httpClient,
			baseURL+InternalRegisterModuleProcedure,
			connect.WithSchema(internalRegisterModuleMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		registerTag: connect.NewClient[v1.RegisterTagRequest, v1.RegisterTagResponse](
			httpClient,
			baseURL+InternalRegisterTagProcedure,
			connect.WithSchema(internalRegisterTagMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		setTypeWithModule: connect.NewClient[v1.SetTypeWithModuleRequest, v1.SetTypeWithModuleResponse](
			httpClient,
			baseURL+InternalSetTypeWithModuleProcedure,
			connect.WithSchema(internalSetTypeWithModuleMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteTypeWithModule: connect.NewClient[v1.DeleteTypeWithModuleRequest, v1.DeleteTypeWithModuleResponse](
			httpClient,
			baseURL+InternalDeleteTypeWithModuleProcedure,
			connect.WithSchema(internalDeleteTypeWithModuleMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		setTagWithModule: connect.NewClient[v1.SetTagWithModuleRequest, v1.SetTagWithModuleResponse](
			httpClient,
			baseURL+InternalSetTagWithModuleProcedure,
			connect.WithSchema(internalSetTagWithModuleMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteTagWithModule: connect.NewClient[v1.DeleteTagWithModuleRequest, v1.DeleteTagWithModuleResponse](
			httpClient,
			baseURL+InternalDeleteTagWithModuleProcedure,
			connect.WithSchema(internalDeleteTagWithModuleMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// internalClient implements InternalClient.
type internalClient struct {
	getTagIDTreeize      *connect.Client[v1.GetTagIDTreeizeRequest, v1.GetTagIDTreeizeResponse]
	bindRelation         *connect.Client[v1.BindRelationRequest, v1.BindRelationResponse]
	unbindRelation       *connect.Client[v1.UnbindRelationRequest, v1.UnbindRelationResponse]
	getAllByModule       *connect.Client[v1.GetAllByModuleRequest, v1.GetAllByModuleResponse]
	registerModule       *connect.Client[v1.RegisterModuleRequest, v1.RegisterModuleResponse]
	registerTag          *connect.Client[v1.RegisterTagRequest, v1.RegisterTagResponse]
	setTypeWithModule    *connect.Client[v1.SetTypeWithModuleRequest, v1.SetTypeWithModuleResponse]
	deleteTypeWithModule *connect.Client[v1.DeleteTypeWithModuleRequest, v1.DeleteTypeWithModuleResponse]
	setTagWithModule     *connect.Client[v1.SetTagWithModuleRequest, v1.SetTagWithModuleResponse]
	deleteTagWithModule  *connect.Client[v1.DeleteTagWithModuleRequest, v1.DeleteTagWithModuleResponse]
}

// GetTagIDTreeize calls tag.v1.Internal.GetTagIDTreeize.
func (c *internalClient) GetTagIDTreeize(ctx context.Context, req *connect.Request[v1.GetTagIDTreeizeRequest]) (*connect.Response[v1.GetTagIDTreeizeResponse], error) {
	return c.getTagIDTreeize.CallUnary(ctx, req)
}

// BindRelation calls tag.v1.Internal.BindRelation.
func (c *internalClient) BindRelation(ctx context.Context, req *connect.Request[v1.BindRelationRequest]) (*connect.Response[v1.BindRelationResponse], error) {
	return c.bindRelation.CallUnary(ctx, req)
}

// UnbindRelation calls tag.v1.Internal.UnbindRelation.
func (c *internalClient) UnbindRelation(ctx context.Context, req *connect.Request[v1.UnbindRelationRequest]) (*connect.Response[v1.UnbindRelationResponse], error) {
	return c.unbindRelation.CallUnary(ctx, req)
}

// GetAllByModule calls tag.v1.Internal.GetAllByModule.
func (c *internalClient) GetAllByModule(ctx context.Context, req *connect.Request[v1.GetAllByModuleRequest]) (*connect.Response[v1.GetAllByModuleResponse], error) {
	return c.getAllByModule.CallUnary(ctx, req)
}

// RegisterModule calls tag.v1.Internal.RegisterModule.
func (c *internalClient) RegisterModule(ctx context.Context, req *connect.Request[v1.RegisterModuleRequest]) (*connect.Response[v1.RegisterModuleResponse], error) {
	return c.registerModule.CallUnary(ctx, req)
}

// RegisterTag calls tag.v1.Internal.RegisterTag.
func (c *internalClient) RegisterTag(ctx context.Context, req *connect.Request[v1.RegisterTagRequest]) (*connect.Response[v1.RegisterTagResponse], error) {
	return c.registerTag.CallUnary(ctx, req)
}

// SetTypeWithModule calls tag.v1.Internal.SetTypeWithModule.
func (c *internalClient) SetTypeWithModule(ctx context.Context, req *connect.Request[v1.SetTypeWithModuleRequest]) (*connect.Response[v1.SetTypeWithModuleResponse], error) {
	return c.setTypeWithModule.CallUnary(ctx, req)
}

// DeleteTypeWithModule calls tag.v1.Internal.DeleteTypeWithModule.
func (c *internalClient) DeleteTypeWithModule(ctx context.Context, req *connect.Request[v1.DeleteTypeWithModuleRequest]) (*connect.Response[v1.DeleteTypeWithModuleResponse], error) {
	return c.deleteTypeWithModule.CallUnary(ctx, req)
}

// SetTagWithModule calls tag.v1.Internal.SetTagWithModule.
func (c *internalClient) SetTagWithModule(ctx context.Context, req *connect.Request[v1.SetTagWithModuleRequest]) (*connect.Response[v1.SetTagWithModuleResponse], error) {
	return c.setTagWithModule.CallUnary(ctx, req)
}

// DeleteTagWithModule calls tag.v1.Internal.DeleteTagWithModule.
func (c *internalClient) DeleteTagWithModule(ctx context.Context, req *connect.Request[v1.DeleteTagWithModuleRequest]) (*connect.Response[v1.DeleteTagWithModuleResponse], error) {
	return c.deleteTagWithModule.CallUnary(ctx, req)
}

// InternalHandler is an implementation of the tag.v1.Internal service.
type InternalHandler interface {
	GetTagIDTreeize(context.Context, *connect.Request[v1.GetTagIDTreeizeRequest]) (*connect.Response[v1.GetTagIDTreeizeResponse], error)
	BindRelation(context.Context, *connect.Request[v1.BindRelationRequest]) (*connect.Response[v1.BindRelationResponse], error)
	UnbindRelation(context.Context, *connect.Request[v1.UnbindRelationRequest]) (*connect.Response[v1.UnbindRelationResponse], error)
	GetAllByModule(context.Context, *connect.Request[v1.GetAllByModuleRequest]) (*connect.Response[v1.GetAllByModuleResponse], error)
	RegisterModule(context.Context, *connect.Request[v1.RegisterModuleRequest]) (*connect.Response[v1.RegisterModuleResponse], error)
	RegisterTag(context.Context, *connect.Request[v1.RegisterTagRequest]) (*connect.Response[v1.RegisterTagResponse], error)
	SetTypeWithModule(context.Context, *connect.Request[v1.SetTypeWithModuleRequest]) (*connect.Response[v1.SetTypeWithModuleResponse], error)
	DeleteTypeWithModule(context.Context, *connect.Request[v1.DeleteTypeWithModuleRequest]) (*connect.Response[v1.DeleteTypeWithModuleResponse], error)
	SetTagWithModule(context.Context, *connect.Request[v1.SetTagWithModuleRequest]) (*connect.Response[v1.SetTagWithModuleResponse], error)
	DeleteTagWithModule(context.Context, *connect.Request[v1.DeleteTagWithModuleRequest]) (*connect.Response[v1.DeleteTagWithModuleResponse], error)
}

// NewInternalHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewInternalHandler(svc InternalHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	internalGetTagIDTreeizeHandler := connect.NewUnaryHandler(
		InternalGetTagIDTreeizeProcedure,
		svc.GetTagIDTreeize,
		connect.WithSchema(internalGetTagIDTreeizeMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	internalBindRelationHandler := connect.NewUnaryHandler(
		InternalBindRelationProcedure,
		svc.BindRelation,
		connect.WithSchema(internalBindRelationMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	internalUnbindRelationHandler := connect.NewUnaryHandler(
		InternalUnbindRelationProcedure,
		svc.UnbindRelation,
		connect.WithSchema(internalUnbindRelationMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	internalGetAllByModuleHandler := connect.NewUnaryHandler(
		InternalGetAllByModuleProcedure,
		svc.GetAllByModule,
		connect.WithSchema(internalGetAllByModuleMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	internalRegisterModuleHandler := connect.NewUnaryHandler(
		InternalRegisterModuleProcedure,
		svc.RegisterModule,
		connect.WithSchema(internalRegisterModuleMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	internalRegisterTagHandler := connect.NewUnaryHandler(
		InternalRegisterTagProcedure,
		svc.RegisterTag,
		connect.WithSchema(internalRegisterTagMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	internalSetTypeWithModuleHandler := connect.NewUnaryHandler(
		InternalSetTypeWithModuleProcedure,
		svc.SetTypeWithModule,
		connect.WithSchema(internalSetTypeWithModuleMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	internalDeleteTypeWithModuleHandler := connect.NewUnaryHandler(
		InternalDeleteTypeWithModuleProcedure,
		svc.DeleteTypeWithModule,
		connect.WithSchema(internalDeleteTypeWithModuleMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	internalSetTagWithModuleHandler := connect.NewUnaryHandler(
		InternalSetTagWithModuleProcedure,
		svc.SetTagWithModule,
		connect.WithSchema(internalSetTagWithModuleMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	internalDeleteTagWithModuleHandler := connect.NewUnaryHandler(
		InternalDeleteTagWithModuleProcedure,
		svc.DeleteTagWithModule,
		connect.WithSchema(internalDeleteTagWithModuleMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/tag.v1.Internal/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case InternalGetTagIDTreeizeProcedure:
			internalGetTagIDTreeizeHandler.ServeHTTP(w, r)
		case InternalBindRelationProcedure:
			internalBindRelationHandler.ServeHTTP(w, r)
		case InternalUnbindRelationProcedure:
			internalUnbindRelationHandler.ServeHTTP(w, r)
		case InternalGetAllByModuleProcedure:
			internalGetAllByModuleHandler.ServeHTTP(w, r)
		case InternalRegisterModuleProcedure:
			internalRegisterModuleHandler.ServeHTTP(w, r)
		case InternalRegisterTagProcedure:
			internalRegisterTagHandler.ServeHTTP(w, r)
		case InternalSetTypeWithModuleProcedure:
			internalSetTypeWithModuleHandler.ServeHTTP(w, r)
		case InternalDeleteTypeWithModuleProcedure:
			internalDeleteTypeWithModuleHandler.ServeHTTP(w, r)
		case InternalSetTagWithModuleProcedure:
			internalSetTagWithModuleHandler.ServeHTTP(w, r)
		case InternalDeleteTagWithModuleProcedure:
			internalDeleteTagWithModuleHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedInternalHandler returns CodeUnimplemented from all methods.
type UnimplementedInternalHandler struct{}

func (UnimplementedInternalHandler) GetTagIDTreeize(context.Context, *connect.Request[v1.GetTagIDTreeizeRequest]) (*connect.Response[v1.GetTagIDTreeizeResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("tag.v1.Internal.GetTagIDTreeize is not implemented"))
}

func (UnimplementedInternalHandler) BindRelation(context.Context, *connect.Request[v1.BindRelationRequest]) (*connect.Response[v1.BindRelationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("tag.v1.Internal.BindRelation is not implemented"))
}

func (UnimplementedInternalHandler) UnbindRelation(context.Context, *connect.Request[v1.UnbindRelationRequest]) (*connect.Response[v1.UnbindRelationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("tag.v1.Internal.UnbindRelation is not implemented"))
}

func (UnimplementedInternalHandler) GetAllByModule(context.Context, *connect.Request[v1.GetAllByModuleRequest]) (*connect.Response[v1.GetAllByModuleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("tag.v1.Internal.GetAllByModule is not implemented"))
}

func (UnimplementedInternalHandler) RegisterModule(context.Context, *connect.Request[v1.RegisterModuleRequest]) (*connect.Response[v1.RegisterModuleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("tag.v1.Internal.RegisterModule is not implemented"))
}

func (UnimplementedInternalHandler) RegisterTag(context.Context, *connect.Request[v1.RegisterTagRequest]) (*connect.Response[v1.RegisterTagResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("tag.v1.Internal.RegisterTag is not implemented"))
}

func (UnimplementedInternalHandler) SetTypeWithModule(context.Context, *connect.Request[v1.SetTypeWithModuleRequest]) (*connect.Response[v1.SetTypeWithModuleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("tag.v1.Internal.SetTypeWithModule is not implemented"))
}

func (UnimplementedInternalHandler) DeleteTypeWithModule(context.Context, *connect.Request[v1.DeleteTypeWithModuleRequest]) (*connect.Response[v1.DeleteTypeWithModuleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("tag.v1.Internal.DeleteTypeWithModule is not implemented"))
}

func (UnimplementedInternalHandler) SetTagWithModule(context.Context, *connect.Request[v1.SetTagWithModuleRequest]) (*connect.Response[v1.SetTagWithModuleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("tag.v1.Internal.SetTagWithModule is not implemented"))
}

func (UnimplementedInternalHandler) DeleteTagWithModule(context.Context, *connect.Request[v1.DeleteTagWithModuleRequest]) (*connect.Response[v1.DeleteTagWithModuleResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("tag.v1.Internal.DeleteTagWithModule is not implemented"))
}
