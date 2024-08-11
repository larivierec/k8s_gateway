package gw_api

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/cache"
	gatewayClient "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned"
)

type RouteVersionInterface interface {
	NewLister(ctx context.Context, client gatewayClient.Interface, ns string) cache.ListFunc
	NewWatcher(ctx context.Context, client gatewayClient.Interface, ns string) cache.WatchFunc
	NewObject() runtime.Object
}
