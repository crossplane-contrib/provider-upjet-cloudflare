package config

import (
	"testing"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// The messages below are the generated Cloudflare client's own text, taken from
// cloudflare-go's parameter validation. A resource whose identifier argument is
// not matched here cannot be created at all: its pre-create Read fails, upjet
// treats the failure as fatal, and the create never runs.
func TestIsMissingIdentifierDiagnostic(t *testing.T) {
	for _, tc := range []struct {
		name string
		diag *tfprotov6.Diagnostic
		want bool
	}{
		{
			name: "identifier argument ending in _id",
			diag: &tfprotov6.Diagnostic{
				Summary: "failed to make http request",
				Detail:  "missing required dns_record_id parameter",
			},
			want: true,
		},
		{
			name: "identifier argument ending in _identifier",
			diag: &tfprotov6.Diagnostic{
				Summary: "failed to make http request",
				Detail:  "missing required rule_identifier parameter",
			},
			want: true,
		},
		{
			name: "multi-word identifier argument ending in _identifier",
			diag: &tfprotov6.Diagnostic{
				Summary: "failed to make http request",
				Detail:  "missing required destination_address_identifier parameter",
			},
			want: true,
		},
		{
			name: "scoping argument, matched deliberately",
			diag: &tfprotov6.Diagnostic{
				Summary: "failed to make http request",
				Detail:  "missing required zone_id parameter",
			},
			want: true,
		},
		{
			name: "unrelated API failure is not absence",
			diag: &tfprotov6.Diagnostic{
				Summary: "failed to make http request",
				Detail:  "Authentication error (10000)",
			},
			want: false,
		},
		{
			name: "missing non-identifier argument is not absence",
			diag: &tfprotov6.Diagnostic{
				Summary: "failed to make http request",
				Detail:  "missing required dataset parameter",
			},
			want: false,
		},
		{
			name: "nil diagnostic",
			diag: nil,
			want: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := isMissingIdentifierDiagnostic(tc.diag); got != tc.want {
				t.Errorf("isMissingIdentifierDiagnostic() = %v, want %v", got, tc.want)
			}
		})
	}
}

// A single non-absence error among the diagnostics must make the whole set
// non-absence. Reporting absence there would make upjet create a second
// external resource alongside the one it failed to read.
func TestPluginFrameworkExternalNameIsNotFoundDiagnosticFn(t *testing.T) {
	missing := &tfprotov6.Diagnostic{
		Severity: tfprotov6.DiagnosticSeverityError,
		Summary:  "failed to make http request",
		Detail:   "missing required rule_identifier parameter",
	}
	other := &tfprotov6.Diagnostic{
		Severity: tfprotov6.DiagnosticSeverityError,
		Summary:  "failed to make http request",
		Detail:   "Authentication error (10000)",
	}
	warning := &tfprotov6.Diagnostic{
		Severity: tfprotov6.DiagnosticSeverityWarning,
		Summary:  "Resource not found",
	}

	fn := pluginFrameworkExternalName(config.ExternalName{}).IsNotFoundDiagnosticFn
	for _, tc := range []struct {
		name  string
		diags []*tfprotov6.Diagnostic
		want  bool
	}{
		{name: "no diagnostics", diags: nil, want: false},
		{name: "only the missing identifier", diags: []*tfprotov6.Diagnostic{missing}, want: true},
		{name: "missing identifier alongside another error", diags: []*tfprotov6.Diagnostic{missing, other}, want: false},
		{name: "only another error", diags: []*tfprotov6.Diagnostic{other}, want: false},
		{name: "warnings are ignored", diags: []*tfprotov6.Diagnostic{warning}, want: false},
		{name: "warning alongside the missing identifier", diags: []*tfprotov6.Diagnostic{warning, missing}, want: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := fn(tc.diags); got != tc.want {
				t.Errorf("IsNotFoundDiagnosticFn() = %v, want %v", got, tc.want)
			}
		})
	}
}
