package policyverse

import (
	"encoding/json"
	"testing"
)

func TestPolicyParser(t *testing.T) {
	expected_reponse := []string{
		"secretsmanager:GetRandomPassword",
		"secretsmanager:GetResourcePolicy",
		"secretsmanager:GetSecretValue",
	}
	jsonData := `{
        "Statement": [{
            "Action": ["secretsmanager:Get*"],
            "Resource": "*",
            "Effect": "Allow"
        }]
    }`
	var data map[string]interface{}
	json.Unmarshal([]byte(jsonData), &data)
	expanded_policies := expandPolicy(data)

	if len(expanded_policies) != len(expected_reponse) {
		t.Error("Expected value doesnt match expexted response for", "secretsmanager:Get*", expanded_policies)

	}
	for i := range expanded_policies {
		if expanded_policies[i] != expected_reponse[i] {
			t.Error("Expected value doesnt match expexted response")
		}
	}
}
