package main

import (
	"net/url"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/openservicemesh/osm-health/pkg/connectivity"
	"github.com/openservicemesh/osm-health/pkg/kuberneteshelper"
)

const connectivityPodToURLDesc = `
Checks connectivity between Kubernetes pods
	(add more descriptive description)
`

const connectivityPodToURLExample = `
Example:
	(add example)
`

func newConnectivityPodToURLCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "pod-to-url SOURCE_POD DESTINATION_URL",
		Short:   "Checks connectivity between a Kubernetes pod and a given URL",
		Example: connectivityPodToURLExample,
		Long:    connectivityPodToURLDesc,
		Args:    cobra.ExactArgs(2),
		RunE: func(_ *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.Errorf("provide both SOURCE_POD and DESTINATION_URL")
			}

			fromPod, err := kuberneteshelper.PodFromString(args[0])
			if err != nil {
				return errors.New("invalid SOURCE_POD")
			}

			toURL, err := url.Parse(args[1])
			if err != nil {
				return errors.New("invalid DESTINATION_URL")
			}

			connectivity.PodToURL(fromPod, toURL)
			return nil
		},
	}
}
