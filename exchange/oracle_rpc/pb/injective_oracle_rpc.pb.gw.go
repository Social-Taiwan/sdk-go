// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: injective_oracle_rpc.proto

/*
Package injective_oracle_rpcpb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package injective_oracle_rpcpb

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_InjectiveOracleRPC_OracleList_0(ctx context.Context, marshaler runtime.Marshaler, client InjectiveOracleRPCClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq OracleListRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.OracleList(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_InjectiveOracleRPC_OracleList_0(ctx context.Context, marshaler runtime.Marshaler, server InjectiveOracleRPCServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq OracleListRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.OracleList(ctx, &protoReq)
	return msg, metadata, err

}

func request_InjectiveOracleRPC_Price_0(ctx context.Context, marshaler runtime.Marshaler, client InjectiveOracleRPCClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq PriceRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Price(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_InjectiveOracleRPC_Price_0(ctx context.Context, marshaler runtime.Marshaler, server InjectiveOracleRPCServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq PriceRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.Price(ctx, &protoReq)
	return msg, metadata, err

}

func request_InjectiveOracleRPC_StreamPrices_0(ctx context.Context, marshaler runtime.Marshaler, client InjectiveOracleRPCClient, req *http.Request, pathParams map[string]string) (InjectiveOracleRPC_StreamPricesClient, runtime.ServerMetadata, error) {
	var protoReq StreamPricesRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	stream, err := client.StreamPrices(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

func request_InjectiveOracleRPC_StreamPricesByMarkets_0(ctx context.Context, marshaler runtime.Marshaler, client InjectiveOracleRPCClient, req *http.Request, pathParams map[string]string) (InjectiveOracleRPC_StreamPricesByMarketsClient, runtime.ServerMetadata, error) {
	var protoReq StreamPricesByMarketsRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	stream, err := client.StreamPricesByMarkets(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

// RegisterInjectiveOracleRPCHandlerServer registers the http handlers for service InjectiveOracleRPC to "mux".
// UnaryRPC     :call InjectiveOracleRPCServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterInjectiveOracleRPCHandlerFromEndpoint instead.
func RegisterInjectiveOracleRPCHandlerServer(ctx context.Context, mux *runtime.ServeMux, server InjectiveOracleRPCServer) error {

	mux.Handle("POST", pattern_InjectiveOracleRPC_OracleList_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/injective_oracle_rpc.InjectiveOracleRPC/OracleList", runtime.WithHTTPPathPattern("/injective_oracle_rpc.InjectiveOracleRPC/OracleList"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_InjectiveOracleRPC_OracleList_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_InjectiveOracleRPC_OracleList_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_InjectiveOracleRPC_Price_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/injective_oracle_rpc.InjectiveOracleRPC/Price", runtime.WithHTTPPathPattern("/injective_oracle_rpc.InjectiveOracleRPC/Price"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_InjectiveOracleRPC_Price_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_InjectiveOracleRPC_Price_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_InjectiveOracleRPC_StreamPrices_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	mux.Handle("POST", pattern_InjectiveOracleRPC_StreamPricesByMarkets_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	return nil
}

// RegisterInjectiveOracleRPCHandlerFromEndpoint is same as RegisterInjectiveOracleRPCHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterInjectiveOracleRPCHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption,logger *logrus.Logger) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				logger.Warnf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				logger.Warnf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterInjectiveOracleRPCHandler(ctx, mux, conn)
}

// RegisterInjectiveOracleRPCHandler registers the http handlers for service InjectiveOracleRPC to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterInjectiveOracleRPCHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterInjectiveOracleRPCHandlerClient(ctx, mux, NewInjectiveOracleRPCClient(conn))
}

// RegisterInjectiveOracleRPCHandlerClient registers the http handlers for service InjectiveOracleRPC
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "InjectiveOracleRPCClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "InjectiveOracleRPCClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "InjectiveOracleRPCClient" to call the correct interceptors.
func RegisterInjectiveOracleRPCHandlerClient(ctx context.Context, mux *runtime.ServeMux, client InjectiveOracleRPCClient) error {

	mux.Handle("POST", pattern_InjectiveOracleRPC_OracleList_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/injective_oracle_rpc.InjectiveOracleRPC/OracleList", runtime.WithHTTPPathPattern("/injective_oracle_rpc.InjectiveOracleRPC/OracleList"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_InjectiveOracleRPC_OracleList_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_InjectiveOracleRPC_OracleList_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_InjectiveOracleRPC_Price_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/injective_oracle_rpc.InjectiveOracleRPC/Price", runtime.WithHTTPPathPattern("/injective_oracle_rpc.InjectiveOracleRPC/Price"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_InjectiveOracleRPC_Price_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_InjectiveOracleRPC_Price_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_InjectiveOracleRPC_StreamPrices_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/injective_oracle_rpc.InjectiveOracleRPC/StreamPrices", runtime.WithHTTPPathPattern("/injective_oracle_rpc.InjectiveOracleRPC/StreamPrices"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_InjectiveOracleRPC_StreamPrices_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_InjectiveOracleRPC_StreamPrices_0(annotatedContext, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_InjectiveOracleRPC_StreamPricesByMarkets_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/injective_oracle_rpc.InjectiveOracleRPC/StreamPricesByMarkets", runtime.WithHTTPPathPattern("/injective_oracle_rpc.InjectiveOracleRPC/StreamPricesByMarkets"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_InjectiveOracleRPC_StreamPricesByMarkets_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_InjectiveOracleRPC_StreamPricesByMarkets_0(annotatedContext, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_InjectiveOracleRPC_OracleList_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"injective_oracle_rpc.InjectiveOracleRPC", "OracleList"}, ""))

	pattern_InjectiveOracleRPC_Price_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"injective_oracle_rpc.InjectiveOracleRPC", "Price"}, ""))

	pattern_InjectiveOracleRPC_StreamPrices_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"injective_oracle_rpc.InjectiveOracleRPC", "StreamPrices"}, ""))

	pattern_InjectiveOracleRPC_StreamPricesByMarkets_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"injective_oracle_rpc.InjectiveOracleRPC", "StreamPricesByMarkets"}, ""))
)

var (
	forward_InjectiveOracleRPC_OracleList_0 = runtime.ForwardResponseMessage

	forward_InjectiveOracleRPC_Price_0 = runtime.ForwardResponseMessage

	forward_InjectiveOracleRPC_StreamPrices_0 = runtime.ForwardResponseStream

	forward_InjectiveOracleRPC_StreamPricesByMarkets_0 = runtime.ForwardResponseStream
)
