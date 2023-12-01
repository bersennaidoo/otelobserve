package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/bersennaidoo/otelobserve/physical/otrace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
	"go.opentelemetry.io/otel/trace"
)

func main() {

	serviceName := "otelobserve"
	serviceVersion := "0.1.0"
	otelShutdown, err := otrace.SetupOTelSDK(context.Background(), serviceName, serviceVersion)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	tr := otel.GetTracerProvider().Tracer("main package")
	ctx, sp := tr.Start(context.Background(), "visit store")
	defer sp.End()
	browse(ctx)
	sp = trace.SpanFromContext(ctx)
	sp.SetAttributes(
		attribute.String("browse store", "successful"),
		attribute.String("add items to cart", "successful"),
	)
}

func browse(ctx context.Context) {

	fmt.Println("visiting the grocery store")
	tr := otel.GetTracerProvider().Tracer("main package1")
	ctx, sp1 := tr.Start(ctx, "browse")
	sp1.SetAttributes(
		attribute.KeyValue(semconv.HTTPMethod("GET")),
		attribute.KeyValue(semconv.NetProtocolVersion("1.1")),
		attribute.KeyValue(semconv.NetPeerName("example.com")),
		attribute.KeyValue(semconv.NetSockPeerAddr("101.10.9.5")),
	)
	defer sp1.End()
	addItemToCart(ctx, "orange", "5")
}

func addItemToCart(ctx context.Context, item, quantity string) {
	tr := otel.GetTracerProvider().Tracer("main package2")
	_, sp2 := tr.Start(ctx, "add item to cart")
	defer sp2.End()
	sp2.SetAttributes(
		attribute.String("item", item),
		attribute.String("quantity", quantity),
	)
	fmt.Printf("add %s to cart\n", item)
}
