package talos

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("talos_machine_secrets", func(r *config.Resource) {
        //r.ShortGroup = "machineSecrets"
    })
    p.AddResourceConfigurator("talos_machine_configuration_data_source", func(r *config.Resource) {
        //r.ShortGroup = "machineConfiguration"
    })
    p.AddResourceConfigurator("talos_machine_configuration_apply", func(r *config.Resource) {
        //r.ShortGroup = "machineConfigurationApply"

        // r.References["machineSecrets"] = config.Reference{
        //     Type: "github.com/thevenincloud/provider-talos/apis/talos/v1alpha1.MachineSecrets",
        // }
        // r.References["machineConfiguration"] = config.Reference{
        //     Type: "github.com/thevenincloud/provider-talos/apis/talos/v1alpha1.MachineConfiguration",
        // }
    })
    p.AddResourceConfigurator("talos_machine_bootstrap", func(r *config.Resource) {
        //r.ShortGroup = "machineBootstrap"
    })
    p.AddResourceConfigurator("talos_image_factory_schematic", func(r *config.Resource) {
        //r.ShortGroup = "imageFactorySchematic"
    })
    p.AddResourceConfigurator("talos_cluster_kubeconfig", func(r *config.Resource) {
        //r.ShortGroup = "clusterKubeconfig"
    })
}