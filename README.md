# Provider Cloudflare

`provider-upjet-cloudflare` is a [Crossplane](https://crossplane.io/) provider
cloudflare that is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the Cloudflare
API.

## Getting Started

This cloudflare serves as a starting point for generating a new [Crossplane Provider](https://docs.crossplane.io/latest/packages/providers/) using the [`upjet`](https://github.com/crossplane/upjet) tooling. Please follow the guide linked below to generate a new Provider:

https://github.com/crossplane/upjet/blob/main/docs/generating-a-provider.md

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

### Which runtime serves a resource

terraform-provider-cloudflare v5 is built on the Terraform plugin framework,
and its Read implementations return an error diagnostic rather than empty state
when the resource identifier is absent. upjet's Terraform CLI runtime writes an
empty `id` into the Terraform state before the first create and then refreshes,
so Read is always called without an identifier and the create fails during the
pre-create observe:

```
observe failed: cannot run refresh: refresh failed: failed to make http request:
missing required dns_record_id parameter
```

The plugin framework runtime is the only one that exposes
`IsNotFoundDiagnosticFn`, which lets those diagnostics be read as "the external
resource does not exist" so the create can proceed. Resources listed in
`PluginFrameworkResources` (`config/plugin_framework.go`) are served through it;
everything else still uses the CLI runtime. Resources are moved across
individually as they are verified against the live API.

If a resource fails to create with a `missing required <resource>_id parameter`
message, it needs to be added to that list.

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/crossplane-contrib/provider-upjet-cloudflare/issues).
