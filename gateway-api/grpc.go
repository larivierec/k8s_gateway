package gw_api

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
	gatewayapi_v1 "sigs.k8s.io/gateway-api/apis/v1"
	gatewayapi_v1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
	gatewayClient "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned"
)

type GRPCRouteV1 struct{}

func (v *GRPCRouteV1) NewLister(ctx context.Context, client gatewayClient.Interface, ns string) cache.ListFunc {
	return func(opts metav1.ListOptions) (runtime.Object, error) {
		return client.GatewayV1().GRPCRoutes(ns).List(ctx, opts)
	}
}

func (v *GRPCRouteV1) NewWatcher(ctx context.Context, client gatewayClient.Interface, ns string) cache.WatchFunc {
	return func(opts metav1.ListOptions) (watch.Interface, error) {
		return client.GatewayV1().GRPCRoutes(ns).Watch(ctx, opts)
	}
}

func (v *GRPCRouteV1) NewObject() runtime.Object {
	return &gatewayapi_v1.GRPCRoute{}
}

type GRPCRouteV1alpha2 struct{}

func (v *GRPCRouteV1alpha2) NewLister(ctx context.Context, client gatewayClient.Interface, ns string) cache.ListFunc {
	return func(opts metav1.ListOptions) (runtime.Object, error) {
		return client.GatewayV1().GRPCRoutes(ns).List(ctx, opts)
	}
}

func (v *GRPCRouteV1alpha2) NewWatcher(ctx context.Context, client gatewayClient.Interface, ns string) cache.WatchFunc {
	return func(opts metav1.ListOptions) (watch.Interface, error) {
		return client.GatewayV1alpha2().GRPCRoutes(ns).Watch(ctx, opts)
	}
}

func (v *GRPCRouteV1alpha2) NewObject() runtime.Object {
	return &gatewayapi_v1alpha2.GRPCRoute{}
}
