// Code generated by protoc-gen-micro-gateway. DO NOT EDIT.
// source: examples/proto/examplepb/wrappers.proto

/*
Package examplepb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package examplepb

import (
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/partitio/micro-gateway/runtime"
	"github.com/partitio/micro-gateway/utilities"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_WrappersService_Create_0(ctx context.Context, marshaler runtime.Marshaler, client WrappersService, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	var protoReq Wrappers

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Create(ctx, &protoReq)
	return msg, err

}

func request_WrappersService_CreateStringValue_0(ctx context.Context, marshaler runtime.Marshaler, client WrappersService, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	var protoReq wrappers.StringValue

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CreateStringValue(ctx, &protoReq)
	return msg, err

}

func request_WrappersService_CreateInt32Value_0(ctx context.Context, marshaler runtime.Marshaler, client WrappersService, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	var protoReq wrappers.Int32Value

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CreateInt32Value(ctx, &protoReq)
	return msg, err

}

func request_WrappersService_CreateInt64Value_0(ctx context.Context, marshaler runtime.Marshaler, client WrappersService, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	var protoReq wrappers.Int64Value

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CreateInt64Value(ctx, &protoReq)
	return msg, err

}

func request_WrappersService_CreateFloatValue_0(ctx context.Context, marshaler runtime.Marshaler, client WrappersService, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	var protoReq wrappers.FloatValue

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CreateFloatValue(ctx, &protoReq)
	return msg, err

}

func request_WrappersService_CreateDoubleValue_0(ctx context.Context, marshaler runtime.Marshaler, client WrappersService, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	var protoReq wrappers.DoubleValue

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CreateDoubleValue(ctx, &protoReq)
	return msg, err

}

func request_WrappersService_CreateBoolValue_0(ctx context.Context, marshaler runtime.Marshaler, client WrappersService, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	var protoReq wrappers.BoolValue

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CreateBoolValue(ctx, &protoReq)
	return msg, err

}

func request_WrappersService_CreateUInt32Value_0(ctx context.Context, marshaler runtime.Marshaler, client WrappersService, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	var protoReq wrappers.UInt32Value

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CreateUInt32Value(ctx, &protoReq)
	return msg, err

}

func request_WrappersService_CreateUInt64Value_0(ctx context.Context, marshaler runtime.Marshaler, client WrappersService, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	var protoReq wrappers.UInt64Value

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CreateUInt64Value(ctx, &protoReq)
	return msg, err

}

func request_WrappersService_CreateBytesValue_0(ctx context.Context, marshaler runtime.Marshaler, client WrappersService, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	var protoReq wrappers.BytesValue

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CreateBytesValue(ctx, &protoReq)
	return msg, err

}

func request_WrappersService_CreateEmpty_0(ctx context.Context, marshaler runtime.Marshaler, client WrappersService, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	var protoReq empty.Empty

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CreateEmpty(ctx, &protoReq)
	return msg, err

}

// RegisterWrappersServiceHandler registers the http handlers for service WrappersService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
// func RegisterWrappersServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
//  	return RegisterWrappersServiceHandlerClient(ctx, mux, NewWrappersServiceClient(conn))
//}

// RegisterWrappersServiceHandlerService registers the http handlers for service WrappersService
// to "mux". The handlers forward requests to the micro service endpoint over the given implementation of "WrappersService".
func RegisterWrappersServiceHandlerService(ctx context.Context, mux *runtime.ServeMux, client WrappersService) error {

	mux.Handle("POST", pattern_WrappersService_Create_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, err := request_WrappersService_Create_0(rctx, inboundMarshaler, client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_WrappersService_Create_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_WrappersService_CreateStringValue_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, err := request_WrappersService_CreateStringValue_0(rctx, inboundMarshaler, client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_WrappersService_CreateStringValue_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_WrappersService_CreateInt32Value_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, err := request_WrappersService_CreateInt32Value_0(rctx, inboundMarshaler, client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_WrappersService_CreateInt32Value_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_WrappersService_CreateInt64Value_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, err := request_WrappersService_CreateInt64Value_0(rctx, inboundMarshaler, client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_WrappersService_CreateInt64Value_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_WrappersService_CreateFloatValue_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, err := request_WrappersService_CreateFloatValue_0(rctx, inboundMarshaler, client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_WrappersService_CreateFloatValue_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_WrappersService_CreateDoubleValue_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, err := request_WrappersService_CreateDoubleValue_0(rctx, inboundMarshaler, client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_WrappersService_CreateDoubleValue_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_WrappersService_CreateBoolValue_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, err := request_WrappersService_CreateBoolValue_0(rctx, inboundMarshaler, client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_WrappersService_CreateBoolValue_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_WrappersService_CreateUInt32Value_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, err := request_WrappersService_CreateUInt32Value_0(rctx, inboundMarshaler, client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_WrappersService_CreateUInt32Value_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_WrappersService_CreateUInt64Value_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, err := request_WrappersService_CreateUInt64Value_0(rctx, inboundMarshaler, client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_WrappersService_CreateUInt64Value_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_WrappersService_CreateBytesValue_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, err := request_WrappersService_CreateBytesValue_0(rctx, inboundMarshaler, client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_WrappersService_CreateBytesValue_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_WrappersService_CreateEmpty_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, err := request_WrappersService_CreateEmpty_0(rctx, inboundMarshaler, client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_WrappersService_CreateEmpty_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_WrappersService_Create_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "example", "wrappers"}, ""))

	pattern_WrappersService_CreateStringValue_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "testString"}, ""))

	pattern_WrappersService_CreateInt32Value_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "testInt32"}, ""))

	pattern_WrappersService_CreateInt64Value_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "testInt64"}, ""))

	pattern_WrappersService_CreateFloatValue_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "testFloat"}, ""))

	pattern_WrappersService_CreateDoubleValue_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "testDouble"}, ""))

	pattern_WrappersService_CreateBoolValue_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "testBool"}, ""))

	pattern_WrappersService_CreateUInt32Value_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "testUint32"}, ""))

	pattern_WrappersService_CreateUInt64Value_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "testUint64"}, ""))

	pattern_WrappersService_CreateBytesValue_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "testBytes"}, ""))

	pattern_WrappersService_CreateEmpty_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "testEmpty"}, ""))
)

var (
	forward_WrappersService_Create_0 = runtime.ForwardResponseMessage

	forward_WrappersService_CreateStringValue_0 = runtime.ForwardResponseMessage

	forward_WrappersService_CreateInt32Value_0 = runtime.ForwardResponseMessage

	forward_WrappersService_CreateInt64Value_0 = runtime.ForwardResponseMessage

	forward_WrappersService_CreateFloatValue_0 = runtime.ForwardResponseMessage

	forward_WrappersService_CreateDoubleValue_0 = runtime.ForwardResponseMessage

	forward_WrappersService_CreateBoolValue_0 = runtime.ForwardResponseMessage

	forward_WrappersService_CreateUInt32Value_0 = runtime.ForwardResponseMessage

	forward_WrappersService_CreateUInt64Value_0 = runtime.ForwardResponseMessage

	forward_WrappersService_CreateBytesValue_0 = runtime.ForwardResponseMessage

	forward_WrappersService_CreateEmpty_0 = runtime.ForwardResponseMessage
)
