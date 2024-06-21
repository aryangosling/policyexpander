package policyverse

import "testing"

func TestPolicyParser(t *testing.T) {
	expected_reponse := []string{
		"secretsmanager:GetRandomPassword",
		"secretsmanager:GetResourcePolicy",
		"secretsmanager:GetSecretValue",
	}

	expanded_policies := expandPolicy("secretsmanager:Get*")

	if len(expanded_policies) != len(expected_reponse) {
		t.Error("Expected value doesnt match expexted response for", "secretsmanager:Get*")
	}
	for i := range expanded_policies {
		if expanded_policies[i] != expected_reponse[i] {
			t.Error("Expected value doesnt match expexted response")
		}
	}
}
