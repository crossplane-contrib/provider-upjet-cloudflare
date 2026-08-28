package config

import (
	"regexp"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// PluginFrameworkResources are the Terraform resource names served through
// upjet's Terraform Plugin Framework runtime rather than its Terraform CLI
// runtime.
//
// terraform-provider-cloudflare v5 is built entirely on the plugin framework,
// and its Read implementations return an error-severity diagnostic rather than
// empty state when the resource identifier is absent. Under the CLI runtime
// upjet writes an empty "id" into the Terraform state before the first create
// and then refreshes, so Read is always called without an identifier and every
// create fails during the pre-create observe with a message of the form
// "missing required <argument> parameter".
//
// The plugin framework runtime is the only one that exposes
// IsNotFoundDiagnosticFn, which lets those diagnostics be interpreted as "the
// external resource does not exist" so that the create can proceed. Resources
// are moved onto it individually as they are verified against the live API.
var PluginFrameworkResources = []string{
	"cloudflare_dns_record",
	"cloudflare_email_routing_address",
	"cloudflare_email_routing_catch_all",
	"cloudflare_email_routing_dns",
	"cloudflare_email_routing_rule",
	"cloudflare_email_routing_settings",
	"cloudflare_zone_setting",
}

// isPluginFrameworkResource reports whether the named Terraform resource is
// served through the plugin framework runtime.
func isPluginFrameworkResource(name string) bool {
	for _, n := range PluginFrameworkResources {
		if n == name {
			return true
		}
	}
	return false
}

// PluginFrameworkIncludeList returns PluginFrameworkResources as exact-match
// regular expressions, for WithTerraformPluginFrameworkIncludeList.
func PluginFrameworkIncludeList() []string {
	l := make([]string, len(PluginFrameworkResources))
	for i, name := range PluginFrameworkResources {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
	}
	return l
}

// missingIdentifierMessage matches the generated Cloudflare client's message
// for an empty path parameter. The argument name varies per resource and ends
// in either "_id" or "_identifier": "dns_record_id", "setting_id",
// "rule_identifier", "destination_address_identifier".
//
// Matching both suffixes is what makes cloudflare_email_routing_rule and
// cloudflare_email_routing_address work. Their identifier arguments are
// rule_identifier and destination_address_identifier, so a pattern written for
// "_id parameter" alone does not match them and their Read failures stay fatal,
// which blocks every create.
//
// Scoping arguments such as zone_id and account_id match this pattern too. They
// are populated from the spec, so an empty one is a malformed resource rather
// than a missing external one; the create that follows the "not found" then
// fails with the same message. The error surfaces either way, one step later.
var missingIdentifierMessage = regexp.MustCompile(`missing required [a-z0-9_]*(?:_id|_identifier) parameter`)

// isMissingIdentifierDiagnostic reports whether a diagnostic returned by a
// Read call means the external resource has no identifier yet, and so does not
// exist.
//
// The Cloudflare provider surfaces this as a request-construction failure
// rather than a 404, because it never reaches the API: the generated client
// validates that the path parameter is non-empty and fails first.
func isMissingIdentifierDiagnostic(d *tfprotov6.Diagnostic) bool {
	if d == nil {
		return false
	}
	for _, text := range []string{d.Summary, d.Detail} {
		if missingIdentifierMessage.MatchString(text) {
			return true
		}
	}
	return false
}

// pluginFrameworkExternalName returns ext with an IsNotFoundDiagnosticFn that
// treats a missing-identifier Read failure as a non-existent external resource.
func pluginFrameworkExternalName(ext config.ExternalName) config.ExternalName {
	ext.IsNotFoundDiagnosticFn = func(diags []*tfprotov6.Diagnostic) bool {
		found := false
		for _, d := range diags {
			if d == nil || d.Severity != tfprotov6.DiagnosticSeverityError {
				continue
			}
			if !isMissingIdentifierDiagnostic(d) {
				// Some other error. Surface it rather than reporting the
				// resource as absent, which would make upjet create a
				// duplicate.
				return false
			}
			found = true
		}
		return found
	}
	return ext
}
