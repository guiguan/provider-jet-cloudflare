package worker

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_worker_script", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
	})
	p.AddResourceConfigurator("cloudflare_worker_route", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.References["script_name"] = config.Reference{
			Type: "Script",
		}
	})
}
