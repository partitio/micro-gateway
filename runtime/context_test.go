package runtime_test

import (
	"context"
	"encoding/base64"
	"github.com/micro/go-micro/metadata"
	"github.com/partitio/grpc-gateway/runtime"
	"net/http"
	"reflect"
	"testing"
	"time"
)

const (
	emptyForwardMetaCount = 1
)

func TestAnnotateContext_WorksWithEmpty(t *testing.T) {
	ctx := context.Background()

	request, err := http.NewRequest("GET", "http://www.example.com", nil)
	if err != nil {
		t.Fatalf("http.NewRequest(%q, %q, nil) failed with %v; want success", "GET", "http://www.example.com", err)
	}
	request.Header.Add("Some-Irrelevant-Header", "some value")
	annotated, err := runtime.AnnotateContext(ctx, runtime.NewServeMux(), request)
	if err != nil {
		t.Errorf("runtime.AnnotateContext(ctx, %#v) failed with %v; want success", request, err)
		return
	}
	md, ok := metadata.FromContext(annotated)
	if !ok || len(md) != emptyForwardMetaCount {
		t.Errorf("Expected %d metadata items in context; got %v", emptyForwardMetaCount, md)
	}
}

func TestAnnotateContext_ForwardsGrpcMetadata(t *testing.T) {
	ctx := context.Background()
	request, err := http.NewRequest("GET", "http://www.example.com", nil)
	if err != nil {
		t.Fatalf("http.NewRequest(%q, %q, nil) failed with %v; want success", "GET", "http://www.example.com", err)
	}
	request.Header.Add("Some-Irrelevant-Header", "some value")
	request.Header.Add("Grpc-Metadata-FooBar", "Value1")
	request.Header.Add("Grpc-Metadata-Foo-BAZ", "Value2")
	request.Header.Add("Grpc-Metadata-foo-bAz", "Value3")
	request.Header.Add("Authorization", "Token 1234567890")
	annotated, err := runtime.AnnotateContext(ctx, runtime.NewServeMux(), request)
	if err != nil {
		t.Errorf("runtime.AnnotateContext(ctx, %#v) failed with %v; want success", request, err)
		return
	}
	md, ok := metadata.FromContext(annotated)
	if got, want := len(md), emptyForwardMetaCount+3; !ok || got != want {
		t.Errorf("metadata items in context = %d want %d: %v", got, want, md)
	}
	if got, want := md["Foobar"], "Value1"; !reflect.DeepEqual(got, want) {
		t.Errorf(`md["Foobar"] = %q; want %q`, got, want)
	}
	if got, want := md["Foo-Baz"], "Value2,Value3"; !reflect.DeepEqual(got, want) {
		t.Errorf(`md["Foo-Baz"] = %q want %q`, got, want)
	}
	if got, want := md["Authorization"], "Token 1234567890"; !reflect.DeepEqual(got, want) {
		t.Errorf(`md["Authorization"] = %q want %q`, got, want)
	}
}

func TestAnnotateContext_ForwardGrpcBinaryMetadata(t *testing.T) {
	ctx := context.Background()
	request, err := http.NewRequest("GET", "http://www.example.com", nil)
	if err != nil {
		t.Fatalf("http.NewRequest(%q, %q, nil) failed with %v; want success", "GET", "http://www.example.com", err)
	}

	binData := []byte("\x00test-binary-data")
	request.Header.Add("Grpc-Metadata-Test-Bin", base64.StdEncoding.EncodeToString(binData))

	annotated, err := runtime.AnnotateContext(ctx, runtime.NewServeMux(), request)
	if err != nil {
		t.Errorf("runtime.AnnotateContext(ctx, %#v) failed with %v; want success", request, err)
		return
	}
	md, ok := metadata.FromContext(annotated)
	if !ok || len(md) != emptyForwardMetaCount+1 {
		t.Errorf("Expected %d metadata items in context; got %v", emptyForwardMetaCount+1, md)
	}
	if got, want := md["Test-Bin"], string(binData); !reflect.DeepEqual(got, want) {
		t.Errorf(`md["test-bin"] = %q want %q`, got, want)
	}
}

func TestAnnotateContext_XForwardedFor(t *testing.T) {
	ctx := context.Background()
	request, err := http.NewRequest("GET", "http://bar.foo.example.com", nil)
	if err != nil {
		t.Fatalf("http.NewRequest(%q, %q, nil) failed with %v; want success", "GET", "http://bar.foo.example.com", err)
	}
	request.Header.Add("X-Forwarded-For", "192.0.2.100") // client
	request.RemoteAddr = "192.0.2.200:12345"             // proxy

	annotated, err := runtime.AnnotateContext(ctx, runtime.NewServeMux(), request)
	if err != nil {
		t.Errorf("runtime.AnnotateContext(ctx, %#v) failed with %v; want success", request, err)
		return
	}
	md, ok := metadata.FromContext(annotated)
	if !ok || len(md) != emptyForwardMetaCount+1 {
		t.Errorf("Expected %d metadata items in context; got %v", emptyForwardMetaCount+1, md)
	}
	if got, want := md["X-Forwarded-Host"], "bar.foo.example.com"; !reflect.DeepEqual(got, want) {
		t.Errorf(`md["host"] = %v; want %v`, got, want)
	}
	// Note: it must be in order client, proxy1, proxy2
	if got, want := md["X-Forwarded-For"], "192.0.2.100, 192.0.2.200"; !reflect.DeepEqual(got, want) {
		t.Errorf(`md["x-forwarded-for"] = %v want %v`, got, want)
	}
}

func TestAnnotateContext_SupportsTimeouts(t *testing.T) {
	ctx := context.Background()
	request, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		t.Fatalf(`http.NewRequest("GET", "http://example.com", nil failed with %v; want success`, err)
	}
	annotated, err := runtime.AnnotateContext(ctx, runtime.NewServeMux(), request)
	if err != nil {
		t.Errorf("runtime.AnnotateContext(ctx, %#v) failed with %v; want success", request, err)
		return
	}
	if _, ok := annotated.Deadline(); ok {
		// no deadline by default
		t.Errorf("annotated.Deadline() = _, true; want _, false")
	}

	const acceptableError = 50 * time.Millisecond
	runtime.DefaultContextTimeout = 10 * time.Second
	annotated, err = runtime.AnnotateContext(ctx, runtime.NewServeMux(), request)
	if err != nil {
		t.Errorf("runtime.AnnotateContext(ctx, %#v) failed with %v; want success", request, err)
		return
	}
	deadline, ok := annotated.Deadline()
	if !ok {
		t.Errorf("annotated.Deadline() = _, false; want _, true")
	}
	if got, want := deadline.Sub(time.Now()), runtime.DefaultContextTimeout; got-want > acceptableError || got-want < -acceptableError {
		t.Errorf("deadline.Sub(time.Now()) = %v; want %v; with error %v", got, want, acceptableError)
	}

	for _, spec := range []struct {
		timeout string
		want    time.Duration
	}{
		{
			timeout: "17H",
			want:    17 * time.Hour,
		},
		{
			timeout: "19M",
			want:    19 * time.Minute,
		},
		{
			timeout: "23S",
			want:    23 * time.Second,
		},
		{
			timeout: "1009m",
			want:    1009 * time.Millisecond,
		},
		{
			timeout: "1000003u",
			want:    1000003 * time.Microsecond,
		},
		{
			timeout: "100000007n",
			want:    100000007 * time.Nanosecond,
		},
	} {
		request.Header.Set("Grpc-Timeout", spec.timeout)
		annotated, err = runtime.AnnotateContext(ctx, runtime.NewServeMux(), request)
		if err != nil {
			t.Errorf("runtime.AnnotateContext(ctx, %#v) failed with %v; want success", request, err)
			return
		}
		deadline, ok := annotated.Deadline()
		if !ok {
			t.Errorf("annotated.Deadline() = _, false; want _, true; timeout = %q", spec.timeout)
		}
		if got, want := deadline.Sub(time.Now()), spec.want; got-want > acceptableError || got-want < -acceptableError {
			t.Errorf("deadline.Sub(time.Now()) = %v; want %v; with error %v; timeout= %q", got, want, acceptableError, spec.timeout)
		}
	}
}
func TestAnnotateContext_SupportsCustomAnnotators(t *testing.T) {
	md1 := func(context.Context, *http.Request) metadata.Metadata {
		return metadata.Metadata(map[string]string{"foo": "bar"})
	}
	md2 := func(context.Context, *http.Request) metadata.Metadata {
		return metadata.Metadata(map[string]string{"baz": "qux"})
	}
	expected := metadata.Metadata(map[string]string{"foo": "bar", "baz": "qux"})
	request, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		t.Fatalf(`http.NewRequest("GET", "http://example.com", nil failed with %v; want success`, err)
	}
	annotated, err := runtime.AnnotateContext(context.Background(), runtime.NewServeMux(runtime.WithMetadata(md1), runtime.WithMetadata(md2)), request)
	if err != nil {
		t.Errorf("runtime.AnnotateContext(ctx, %#v) failed with %v; want success", request, err)
		return
	}
	actual, _ := metadata.FromContext(annotated)
	for key, e := range expected {
		if a, ok := actual[key]; !ok || !reflect.DeepEqual(e, a) {
			t.Errorf("metadata.Metadata[%s] = %v; want %v", key, a, e)
		}
	}
}
