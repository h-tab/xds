package server

import (
	"log"

	xdspb2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	//corev2 "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	//core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	discoverypb "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/gogo/protobuf/proto"
)

const (
	// TODO: mapping
	v2ldsURL = "type.googleapis.com/envoy.api.v2.Listener"
	v2rdsURL = "type.googleapis.com/envoy.api.v2.RouteConfiguration"
	v2cdsURL = "type.googleapis.com/envoy.api.v2.Cluster"
	v2edsURL = "type.googleapis.com/envoy.api.v2.ClusterLoadAssignment"
	ldsURL   = "type.googleapis.com/envoy.config.listener.v3.Listener"
	rdsURL   = "type.googleapis.com/envoy.config.route.v3.RouteConfiguration"
	cdsURL   = "type.googleapis.com/envoy.config.cluster.v3.Cluster"
	edsURL   = "type.googleapis.com/envoy.config.endpoint.v3.ClusterLoadAssignment"
)

// DiscoveryResponseToV2 converts a v3 proto struct to a v2 one.
func DiscoveryResponseToV2(r *discoverypb.DiscoveryResponse) *xdspb2.DiscoveryResponse {
	b := proto.NewBuffer(nil)
	b.SetDeterministic(true)
	err := b.Marshal(r)

	err = err
	x := &xdspb2.DiscoveryResponse{}
	if err := proto.Unmarshal(b.Bytes(), x); err != nil {
		log.Fatalln("Failed to parse DiscoveryResponse:", err)
	}

	x.TypeUrl = v2edsURL
	for i := range x.GetResources() {
		x.Resources[i].TypeUrl = v2edsURL
	}
	log.Printf("RESPONSE TO V2 %v", x)

	return x
}

// DiscoveryRequestToV3 converts a v2 proto struct to a v3 one.
func DiscoveryRequestToV3(r *xdspb2.DiscoveryRequest) *discoverypb.DiscoveryRequest {
	b := proto.NewBuffer(nil)
	b.SetDeterministic(true)
	err := b.Marshal(r)

	err = err
	x := &discoverypb.DiscoveryRequest{}
	if err := proto.Unmarshal(b.Bytes(), x); err != nil {
		log.Fatalln("Failed to parse DiscoveryRequest:", err)
	}
	x.TypeUrl = edsURL
	log.Printf("REQUEST TO V3 %v", x)

	return x
}
