package zone

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "zone"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_zone", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Zone"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zone_setting", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZoneSetting"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		// The "value" field is a dynamic pseudo-type (DynamicPseudoType) which
		// gets rendered in the CRD as x-kubernetes-preserve-unknown-fields
		// without a concrete type. CEL validation rules using has() cannot
		// compile against such fields. Mark it as Optional to prevent the
		// generation of an invalid CEL required-field validation rule.
		if s, ok := r.TerraformResource.Schema["value"]; ok {
			s.Optional = true
		}
	})

	p.AddResourceConfigurator("cloudflare_zone_dnssec", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZoneDNSSEC"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_zone_lockdown", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZoneLockdown"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_zone_hold", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZoneHold"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_zone_cache_reserve", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZoneCacheReserve"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_zone_cache_variants", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZoneCacheVariants"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})
}
