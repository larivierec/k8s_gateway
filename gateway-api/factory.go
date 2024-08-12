package gw_api

import "fmt"

func GetRouteVersion(routeType, version string) (RouteVersionInterface, error) {
	switch routeType {
	case "HTTPRoute":
		switch version {
		case "v1":
			return &HTTPRouteV1{}, nil
		}
	case "TLSRoute":
		switch version {
		case "v1alpha2":
			return &TLSRouteV1alpha2{}, nil
		}
	case "GRPCRoute":
		switch version {
		case "v1alpha2":
			return &GRPCRouteV1alpha2{}, nil
		case "v1":
			return &GRPCRouteV1{}, nil
		}
	}
	return nil, fmt.Errorf("unsupported route type or version: %s/%s", routeType, version)
}
