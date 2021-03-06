// Code generated by protoc-gen-grpc-gateway
// source: examples/examplepb/a_bit_of_everything.proto
// DO NOT EDIT!

/*
Package examplepb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package examplepb

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gengo/grpc-gateway/examples/sub"
	"github.com/gengo/grpc-gateway/runtime"
	"github.com/gengo/grpc-gateway/utilities"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var _ codes.Code
var _ io.Reader
var _ = runtime.String
var _ = json.Marshal
var _ = utilities.PascalFromSnake

var (
	filter_ABitOfEverythingService_Create_0 = &utilities.DoubleArray{Encoding: map[string]int{"float_value": 0, "double_value": 1, "int64_value": 2, "uint64_value": 3, "int32_value": 4, "fixed64_value": 5, "fixed32_value": 6, "bool_value": 7, "string_value": 8, "uint32_value": 9, "sfixed32_value": 10, "sfixed64_value": 11, "sint32_value": 12, "sint64_value": 13}, Base: []int{1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Check: []int{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}}
)

func request_ABitOfEverythingService_Create_0(ctx context.Context, client ABitOfEverythingServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	var protoReq ABitOfEverything

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["float_value"]
	if !ok {
		return nil, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "float_value")
	}

	protoReq.FloatValue, err = runtime.Float32(val)

	if err != nil {
		return nil, err
	}

	val, ok = pathParams["double_value"]
	if !ok {
		return nil, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "double_value")
	}

	protoReq.DoubleValue, err = runtime.Float64(val)

	if err != nil {
		return nil, err
	}

	val, ok = pathParams["int64_value"]
	if !ok {
		return nil, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "int64_value")
	}

	protoReq.Int64Value, err = runtime.Int64(val)

	if err != nil {
		return nil, err
	}

	val, ok = pathParams["uint64_value"]
	if !ok {
		return nil, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "uint64_value")
	}

	protoReq.Uint64Value, err = runtime.Uint64(val)

	if err != nil {
		return nil, err
	}

	val, ok = pathParams["int32_value"]
	if !ok {
		return nil, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "int32_value")
	}

	protoReq.Int32Value, err = runtime.Int32(val)

	if err != nil {
		return nil, err
	}

	val, ok = pathParams["fixed64_value"]
	if !ok {
		return nil, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "fixed64_value")
	}

	protoReq.Fixed64Value, err = runtime.Uint64(val)

	if err != nil {
		return nil, err
	}

	val, ok = pathParams["fixed32_value"]
	if !ok {
		return nil, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "fixed32_value")
	}

	protoReq.Fixed32Value, err = runtime.Uint32(val)

	if err != nil {
		return nil, err
	}

	val, ok = pathParams["bool_value"]
	if !ok {
		return nil, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "bool_value")
	}

	protoReq.BoolValue, err = runtime.Bool(val)

	if err != nil {
		return nil, err
	}

	val, ok = pathParams["string_value"]
	if !ok {
		return nil, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "string_value")
	}

	protoReq.StringValue, err = runtime.String(val)

	if err != nil {
		return nil, err
	}

	val, ok = pathParams["uint32_value"]
	if !ok {
		return nil, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "uint32_value")
	}

	protoReq.Uint32Value, err = runtime.Uint32(val)

	if err != nil {
		return nil, err
	}

	val, ok = pathParams["sfixed32_value"]
	if !ok {
		return nil, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "sfixed32_value")
	}

	protoReq.Sfixed32Value, err = runtime.Int32(val)

	if err != nil {
		return nil, err
	}

	val, ok = pathParams["sfixed64_value"]
	if !ok {
		return nil, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "sfixed64_value")
	}

	protoReq.Sfixed64Value, err = runtime.Int64(val)

	if err != nil {
		return nil, err
	}

	val, ok = pathParams["sint32_value"]
	if !ok {
		return nil, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "sint32_value")
	}

	protoReq.Sint32Value, err = runtime.Int32(val)

	if err != nil {
		return nil, err
	}

	val, ok = pathParams["sint64_value"]
	if !ok {
		return nil, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "sint64_value")
	}

	protoReq.Sint64Value, err = runtime.Int64(val)

	if err != nil {
		return nil, err
	}

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_ABitOfEverythingService_Create_0); err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	return client.Create(ctx, &protoReq)
}

func request_ABitOfEverythingService_CreateBody_0(ctx context.Context, client ABitOfEverythingServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	var protoReq ABitOfEverything

	if err := json.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	return client.CreateBody(ctx, &protoReq)
}

func request_ABitOfEverythingService_BulkCreate_0(ctx context.Context, client ABitOfEverythingServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	stream, err := client.BulkCreate(ctx)
	if err != nil {
		glog.Errorf("Failed to start streaming: %v", err)
		return nil, err
	}
	dec := json.NewDecoder(req.Body)
	for {
		var protoReq ABitOfEverything
		err = dec.Decode(&protoReq)
		if err == io.EOF {
			break
		}
		if err != nil {
			glog.Errorf("Failed to decode request: %v", err)
			return nil, grpc.Errorf(codes.InvalidArgument, "%v", err)
		}
		if err = stream.Send(&protoReq); err != nil {
			glog.Errorf("Failed to send request: %v", err)
			return nil, err
		}
	}

	return stream.CloseAndRecv()

}

func request_ABitOfEverythingService_Lookup_0(ctx context.Context, client ABitOfEverythingServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	var protoReq IdMessage

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["uuid"]
	if !ok {
		return nil, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "uuid")
	}

	protoReq.Uuid, err = runtime.String(val)

	if err != nil {
		return nil, err
	}

	return client.Lookup(ctx, &protoReq)
}

func request_ABitOfEverythingService_List_0(ctx context.Context, client ABitOfEverythingServiceClient, req *http.Request, pathParams map[string]string) (ABitOfEverythingService_ListClient, error) {
	var protoReq EmptyMessage

	return client.List(ctx, &protoReq)
}

func request_ABitOfEverythingService_Update_0(ctx context.Context, client ABitOfEverythingServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	var protoReq ABitOfEverything

	if err := json.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["uuid"]
	if !ok {
		return nil, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "uuid")
	}

	protoReq.Uuid, err = runtime.String(val)

	if err != nil {
		return nil, err
	}

	return client.Update(ctx, &protoReq)
}

func request_ABitOfEverythingService_Delete_0(ctx context.Context, client ABitOfEverythingServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	var protoReq IdMessage

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["uuid"]
	if !ok {
		return nil, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "uuid")
	}

	protoReq.Uuid, err = runtime.String(val)

	if err != nil {
		return nil, err
	}

	return client.Delete(ctx, &protoReq)
}

func request_ABitOfEverythingService_Echo_0(ctx context.Context, client ABitOfEverythingServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	var protoReq sub.StringMessage

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["value"]
	if !ok {
		return nil, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "value")
	}

	protoReq.Value, err = runtime.StringP(val)

	if err != nil {
		return nil, err
	}

	return client.Echo(ctx, &protoReq)
}

func request_ABitOfEverythingService_Echo_1(ctx context.Context, client ABitOfEverythingServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	var protoReq sub.StringMessage

	if err := json.NewDecoder(req.Body).Decode(&protoReq.Value); err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	return client.Echo(ctx, &protoReq)
}

var (
	filter_ABitOfEverythingService_Echo_2 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_ABitOfEverythingService_Echo_2(ctx context.Context, client ABitOfEverythingServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, error) {
	var protoReq sub.StringMessage

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_ABitOfEverythingService_Echo_2); err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	return client.Echo(ctx, &protoReq)
}

func request_ABitOfEverythingService_BulkEcho_0(ctx context.Context, client ABitOfEverythingServiceClient, req *http.Request, pathParams map[string]string) (ABitOfEverythingService_BulkEchoClient, error) {
	stream, err := client.BulkEcho(ctx)
	if err != nil {
		glog.Errorf("Failed to start streaming: %v", err)
		return nil, err
	}
	dec := json.NewDecoder(req.Body)
	for {
		var protoReq sub.StringMessage
		err = dec.Decode(&protoReq)
		if err == io.EOF {
			break
		}
		if err != nil {
			glog.Errorf("Failed to decode request: %v", err)
			return nil, grpc.Errorf(codes.InvalidArgument, "%v", err)
		}
		if err = stream.Send(&protoReq); err != nil {
			glog.Errorf("Failed to send request: %v", err)
			return nil, err
		}
	}

	if err = stream.CloseSend(); err != nil {
		glog.Errorf("Failed to terminate client stream: %v", err)
		return nil, err
	}
	return stream, nil

}

// RegisterABitOfEverythingServiceHandlerFromEndpoint is same as RegisterABitOfEverythingServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterABitOfEverythingServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string) (err error) {
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				glog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				glog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterABitOfEverythingServiceHandler(ctx, mux, conn)
}

// RegisterABitOfEverythingServiceHandler registers the http handlers for service ABitOfEverythingService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterABitOfEverythingServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	client := NewABitOfEverythingServiceClient(conn)

	mux.Handle("POST", pattern_ABitOfEverythingService_Create_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		resp, err := request_ABitOfEverythingService_Create_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, w, err)
			return
		}

		forward_ABitOfEverythingService_Create_0(ctx, w, req, resp)

	})

	mux.Handle("POST", pattern_ABitOfEverythingService_CreateBody_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		resp, err := request_ABitOfEverythingService_CreateBody_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, w, err)
			return
		}

		forward_ABitOfEverythingService_CreateBody_0(ctx, w, req, resp)

	})

	mux.Handle("POST", pattern_ABitOfEverythingService_BulkCreate_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		resp, err := request_ABitOfEverythingService_BulkCreate_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, w, err)
			return
		}

		forward_ABitOfEverythingService_BulkCreate_0(ctx, w, req, resp)

	})

	mux.Handle("GET", pattern_ABitOfEverythingService_Lookup_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		resp, err := request_ABitOfEverythingService_Lookup_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, w, err)
			return
		}

		forward_ABitOfEverythingService_Lookup_0(ctx, w, req, resp)

	})

	mux.Handle("GET", pattern_ABitOfEverythingService_List_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		resp, err := request_ABitOfEverythingService_List_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, w, err)
			return
		}

		forward_ABitOfEverythingService_List_0(ctx, w, req, func() (proto.Message, error) { return resp.Recv() })

	})

	mux.Handle("PUT", pattern_ABitOfEverythingService_Update_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		resp, err := request_ABitOfEverythingService_Update_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, w, err)
			return
		}

		forward_ABitOfEverythingService_Update_0(ctx, w, req, resp)

	})

	mux.Handle("DELETE", pattern_ABitOfEverythingService_Delete_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		resp, err := request_ABitOfEverythingService_Delete_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, w, err)
			return
		}

		forward_ABitOfEverythingService_Delete_0(ctx, w, req, resp)

	})

	mux.Handle("GET", pattern_ABitOfEverythingService_Echo_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		resp, err := request_ABitOfEverythingService_Echo_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, w, err)
			return
		}

		forward_ABitOfEverythingService_Echo_0(ctx, w, req, resp)

	})

	mux.Handle("POST", pattern_ABitOfEverythingService_Echo_1, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		resp, err := request_ABitOfEverythingService_Echo_1(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, w, err)
			return
		}

		forward_ABitOfEverythingService_Echo_1(ctx, w, req, resp)

	})

	mux.Handle("GET", pattern_ABitOfEverythingService_Echo_2, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		resp, err := request_ABitOfEverythingService_Echo_2(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, w, err)
			return
		}

		forward_ABitOfEverythingService_Echo_2(ctx, w, req, resp)

	})

	mux.Handle("POST", pattern_ABitOfEverythingService_BulkEcho_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		resp, err := request_ABitOfEverythingService_BulkEcho_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		if err != nil {
			runtime.HTTPError(ctx, w, err)
			return
		}

		forward_ABitOfEverythingService_BulkEcho_0(ctx, w, req, func() (proto.Message, error) { return resp.Recv() })

	})

	return nil
}

var (
	pattern_ABitOfEverythingService_Create_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3, 1, 0, 4, 1, 5, 4, 1, 0, 4, 1, 5, 5, 2, 6, 1, 0, 4, 1, 5, 7, 1, 0, 4, 1, 5, 8, 1, 0, 4, 1, 5, 9, 1, 0, 4, 1, 5, 10, 1, 0, 4, 1, 5, 11, 2, 12, 1, 0, 4, 2, 5, 13, 1, 0, 4, 1, 5, 14, 1, 0, 4, 1, 5, 15, 1, 0, 4, 1, 5, 16, 1, 0, 4, 1, 5, 17, 1, 0, 4, 1, 5, 18}, []string{"v1", "example", "a_bit_of_everything", "float_value", "double_value", "int64_value", "separator", "uint64_value", "int32_value", "fixed64_value", "fixed32_value", "bool_value", "strprefix", "string_value", "uint32_value", "sfixed32_value", "sfixed64_value", "sint32_value", "sint64_value"}, ""))

	pattern_ABitOfEverythingService_CreateBody_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "example", "a_bit_of_everything"}, ""))

	pattern_ABitOfEverythingService_BulkCreate_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v1", "example", "a_bit_of_everything", "bulk"}, ""))

	pattern_ABitOfEverythingService_Lookup_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3}, []string{"v1", "example", "a_bit_of_everything", "uuid"}, ""))

	pattern_ABitOfEverythingService_List_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "example", "a_bit_of_everything"}, ""))

	pattern_ABitOfEverythingService_Update_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3}, []string{"v1", "example", "a_bit_of_everything", "uuid"}, ""))

	pattern_ABitOfEverythingService_Delete_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3}, []string{"v1", "example", "a_bit_of_everything", "uuid"}, ""))

	pattern_ABitOfEverythingService_Echo_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 1, 0, 4, 1, 5, 4}, []string{"v1", "example", "a_bit_of_everything", "echo", "value"}, ""))

	pattern_ABitOfEverythingService_Echo_1 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v2", "example", "echo"}, ""))

	pattern_ABitOfEverythingService_Echo_2 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v2", "example", "echo"}, ""))

	pattern_ABitOfEverythingService_BulkEcho_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v1", "example", "a_bit_of_everything", "echo"}, ""))
)

var (
	forward_ABitOfEverythingService_Create_0 = runtime.ForwardResponseMessage

	forward_ABitOfEverythingService_CreateBody_0 = runtime.ForwardResponseMessage

	forward_ABitOfEverythingService_BulkCreate_0 = runtime.ForwardResponseMessage

	forward_ABitOfEverythingService_Lookup_0 = runtime.ForwardResponseMessage

	forward_ABitOfEverythingService_List_0 = runtime.ForwardResponseStream

	forward_ABitOfEverythingService_Update_0 = runtime.ForwardResponseMessage

	forward_ABitOfEverythingService_Delete_0 = runtime.ForwardResponseMessage

	forward_ABitOfEverythingService_Echo_0 = runtime.ForwardResponseMessage

	forward_ABitOfEverythingService_Echo_1 = runtime.ForwardResponseMessage

	forward_ABitOfEverythingService_Echo_2 = runtime.ForwardResponseMessage

	forward_ABitOfEverythingService_BulkEcho_0 = runtime.ForwardResponseStream
)
