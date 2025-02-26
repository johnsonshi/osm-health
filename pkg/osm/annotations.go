package osm

// SupportedAnnotations maintains a mapping of OSM version to supported annotations.
var SupportedAnnotations = map[ControllerVersion][]Annotation{
	"v0.5": {
		"openservicemesh.io/monitored-by",
		"openservicemesh.io/sidecar-injection",
		"openservicemesh.io/metrics",
	},
	"v0.6": {
		"openservicemesh.io/monitored-by",
		"openservicemesh.io/sidecar-injection",
		"openservicemesh.io/metrics",
	},
	"v0.7": {
		"openservicemesh.io/monitored-by",
		"openservicemesh.io/sidecar-injection",
		"openservicemesh.io/metrics",
	},
	"v0.8": {
		"openservicemesh.io/monitored-by",
		"openservicemesh.io/sidecar-injection",
		"openservicemesh.io/metrics",
	},
	"v0.9": {
		"openservicemesh.io/monitored-by",
		"openservicemesh.io/sidecar-injection",
		"openservicemesh.io/metrics",
		"openservicemesh.io/outbound-port-exclusion-list",
		"openservicemesh.io/inbound-port-exclusion-list",
	},
}
