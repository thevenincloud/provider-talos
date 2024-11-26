/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"talos_client_configuration_data_source": config.IdentifierFromProvider,
	"talos_cluster_health_data_source": config.IdentifierFromProvider,
	"talos_cluster_kubeconfig_data_source": config.IdentifierFromProvider,
	"talos_image_factory_extensions_versions_data_source": config.IdentifierFromProvider,
	"talos_image_factory_overlays_versions_data_source": config.IdentifierFromProvider,
	"talos_image_factory_urls_data_source": config.IdentifierFromProvider,
	"talos_image_factory_versions_data_source": config.IdentifierFromProvider,
	"talos_machine_configuration_data_source": config.IdentifierFromProvider,
	"talos_machine_disks_data_source": config.IdentifierFromProvider,
	"talos_cluster_kubeconfig": config.IdentifierFromProvider,
	"talos_image_factory_schematic": config.IdentifierFromProvider,
	"talos_machine_bootstrap": config.IdentifierFromProvider,
	"talos_machine_configuration_apply": config.IdentifierFromProvider,
	"talos_machine_secrets": config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
