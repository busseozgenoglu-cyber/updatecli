package githubaction

import (
	"strings"
	"testing"
)

func TestWorkflowManifestGitHubTemplateUsesDefaultYAMLEngine(t *testing.T) {
	if strings.Contains(workflowManifestGitHubTemplate, "engine: 'yamlpath'") {
		t.Fatal("GitHub Action autodiscovery targets must use the default YAML engine")
	}

	if got := strings.Count(workflowManifestGitHubTemplate, "kind: 'yaml'"); got != 2 {
		t.Fatalf("expected two quoted YAML targets, got %d", got)
	}
	if got := strings.Count(workflowManifestGitHubTemplate, "    kind: yaml\n"); got != 1 {
		t.Fatalf("expected one unquoted YAML target, got %d", got)
	}
}
