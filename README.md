# Provider Cloudflare

`provider-upjet-cloudflare` is a [Crossplane](https://crossplane.io/) provider for
Cloudflare that is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the Cloudflare
API.

It is generated from the [Cloudflare Terraform
provider](https://github.com/cloudflare/terraform-provider-cloudflare), so the
set of managed resources tracks that provider's schema. The version currently
generated against is pinned as `TERRAFORM_PROVIDER_VERSION` in the `Makefile`.

## Installing

Provider packages are published to the GitHub Container Registry of the
repository owner, as `ghcr.io/<owner>/provider-upjet-cloudflare`. Install a
released version with a `Provider`:

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-upjet-cloudflare
spec:
  package: ghcr.io/crossplane-contrib/provider-upjet-cloudflare:v0.1.0
```

Substitute the owner if you are installing a package built from a fork — see
[Building and publishing a package](#building-and-publishing-a-package).

Installing the provider takes a while: the package contains a CRD for every
Cloudflare resource the Terraform provider exposes, and the API server must
accept all of them before the provider reports `Healthy`. Watch it with:

```console
kubectl get providers.pkg.crossplane.io provider-upjet-cloudflare -w
```

`Installed=False` with a registry error in the condition message means the
package reference could not be pulled, not that the provider failed to start.

### Credentials

The provider reads credentials from a Kubernetes `Secret` holding a **flat JSON
object**, not a bare token. The recognised keys mirror the Terraform provider's
own arguments — `api_token`, `api_key`, `api_user_service_key`, `email`,
`base_url`, and `user_agent_operator_suffix`. An API token (rather than the
legacy global API key plus email) is the recommended form:

```console
kubectl create secret generic cloudflare-credentials \
  --namespace crossplane-system \
  --from-literal=credentials='{"api_token":"<CLOUDFLARE_API_TOKEN>"}'
```

Scope the token to only the resources you intend to manage. For DNS records and
email routing that is: Zone → Zone: Read; Zone → DNS: Edit; Zone → Email Routing
Rules: Edit.

### ProviderConfig

Crossplane v2 supports both namespaced and cluster-scoped configuration. Managed
resources in the namespaced API groups (`*.cloudflare.m.cloudflare.com`, such as
`dns.cloudflare.m.cloudflare.com`) reference a namespaced `ProviderConfig` or a
`ClusterProviderConfig`, both served under `cloudflare.m.cloudflare.com`:

```yaml
apiVersion: cloudflare.m.cloudflare.com/v1beta1
kind: ProviderConfig
metadata:
  name: default
  namespace: crossplane-system
spec:
  credentials:
    source: Secret
    secretRef:
      name: cloudflare-credentials
      namespace: crossplane-system
      key: credentials
```

The cluster-scoped managed resources live under `*.cloudflare.crossplane.io` and
take a cluster-scoped `ProviderConfig` from `cloudflare.cloudflare.com/v1beta1`.

The group names are not guessable from the project name, and they differ between
the namespaced and cluster-scoped trees. They are defined in
`apis/namespaced/v1beta1/register.go` and `apis/cluster/v1beta1/register.go`; the
authoritative list once installed is `kubectl get crd | grep cloudflare`.
Further examples are under [`examples/`](examples/).

## Building and publishing a package

The package is generated code, so it only needs rebuilding when one of its
inputs changes — a new Cloudflare Terraform provider release, a new
upjet/crossplane-runtime version, or a dependency security fix. Cloudflare API
changes alone do not require a rebuild; they reach this provider only once they
appear in a Terraform provider release.

Publishing is therefore a manually triggered, two-step process.

**1. Tag the release.**

```console
git tag -a v0.1.0 -m "Release v0.1.0"
git push origin v0.1.0
```

> The repository also has a `Tag` workflow, but it currently cannot run: it
> calls `crossplane-contrib/provider-workflows/.github/workflows/tag.yml`, which
> declares only an `on.workflow_dispatch` trigger and so is not callable as a
> reusable workflow. Dispatching it fails with `workflow is not reusable as it
> is missing a 'on.workflow_call' trigger`.

**2. Build and publish the package.** Run the `Publish Provider Package`
workflow **against the tag created above**, so the build is reproducible from a
fixed ref:

```console
gh workflow run publish-provider-package.yml --ref v0.1.0 -f version=v0.1.0
```

It can also be dispatched from the Actions tab in the GitHub web UI ("Run
workflow"), selecting the tag as the ref. The workflow accepts an optional
`go-version` input if the build needs a Go toolchain other than the default.

The publish workflow pushes to `ghcr.io/${{ github.repository_owner }}`, so a
fork publishes to its own owner's registry with no changes. It authenticates
with the run's own `GITHUB_TOKEN`; no personal access token or registry secret
is required. Mirroring to a secondary registry is skipped unless the
`XPKG_MIRROR_ACCESS_ID` and `XPKG_MIRROR_TOKEN` secrets are set.

Crossplane pulls anonymously, so the published package must be public. A package
created by a workflow in a public repository inherits that visibility, but if the
`Provider` fails to install with a registry authorization error, check it under
the owner's Packages → the package → Package settings → Change visibility.

Note that `ghcr.io` returns `401` to an unauthenticated manifest request even for
public packages, so a bare `curl` is not a visibility test. Fetch an anonymous
token first:

```console
tok=$(curl -s "https://ghcr.io/token?scope=repository:<owner>/provider-upjet-cloudflare:pull&service=ghcr.io" | jq -r .token)
curl -sS -o /dev/null -w '%{http_code}\n' -H "Authorization: Bearer $tok" \
  -H "Accept: application/vnd.oci.image.index.v1+json" \
  "https://ghcr.io/v2/<owner>/provider-upjet-cloudflare/manifests/v0.1.0"
```

For a genuinely private package it is the token request that fails, not the
manifest request.

### Upgrading the Cloudflare Terraform provider

Bump `TERRAFORM_PROVIDER_VERSION` in the `Makefile`, then regenerate and open a
pull request. CI reports schema and CRD breaking changes on the pull request, so
review that output before tagging a release.

## Developing

Run the code-generation pipeline:

```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build binary:

```console
make build
```

Build and publish locally, overriding the registry organisation:

```console
make all XPKG_REG_ORGS=ghcr.io/<owner> XPKG_REG_ORGS_NO_PROMOTE=ghcr.io/<owner>
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/crossplane-contrib/provider-upjet-cloudflare/issues).
