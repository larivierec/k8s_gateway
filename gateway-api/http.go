package gw_api

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
	gatewayapi_v1 "sigs.k8s.io/gateway-api/apis/v1"
	gatewayClient "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned"
)

type HTTPRouteV1 struct{}

func (v *HTTPRouteV1) NewLister(ctx context.Context, client gatewayClient.Interface, ns string) cache.ListFunc {
	return func(opts metav1.ListOptions) (runtime.Object, error) {
		return client.GatewayV1().HTTPRoutes(ns).List(ctx, opts)
	}
}

func (v *HTTPRouteV1) NewWatcher(ctx context.Context, client gatewayClient.Interface, ns string) cache.WatchFunc {
	return func(opts metav1.ListOptions) (watch.Interface, error) {
		return client.GatewayV1().HTTPRoutes(ns).Watch(ctx, opts)
	}
}

func (v *HTTPRouteV1) NewObject() runtime.Object {
	return &gatewayapi_v1.HTTPRoute{}
}
