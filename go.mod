module github.com/openservicemesh/osm-health

go 1.16

require (
	github.com/deckarep/golang-set v1.7.1
	github.com/envoyproxy/go-control-plane v0.10.2-0.20220325020618-49ff273808a1
	github.com/fatih/color v1.13.0
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.2.0
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/openservicemesh/osm v0.8.2-0.20210823171715-5fba2d3f778b
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.23.0
	github.com/servicemeshinterface/smi-sdk-go v0.5.0
	github.com/spf13/cobra v1.5.0
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.8.0
	google.golang.org/protobuf v1.28.0
	helm.sh/helm/v3 v3.10.3
	k8s.io/api v0.25.2
	k8s.io/apimachinery v0.25.2
	k8s.io/cli-runtime v0.25.2
	k8s.io/client-go v0.25.2
)

replace (
	github.com/docker/distribution => github.com/docker/distribution v0.0.0-20191216044856-a8371794149d
	github.com/docker/docker => github.com/moby/moby v17.12.0-ce-rc1.0.20200618181300-9dc6525e6118+incompatible
)
