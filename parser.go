package policyverse

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

var actionsNames []string

func init() {
	actionsNames = getPolicies()
	// answer := expandPolicy("secretsmanager:Get*")
	// print(answer)
}

type Action struct {
	EffectiveActionNames []string `json:"effective_action_names"`
}

type Data struct {
	Actions []Action `json:"policies"`
}

func getPolicies() []string {
	// Open the JSON file
	jsonFile, err := os.Open("data.json")
	if err != nil {
		log.Fatalf("Failed to open JSON File : %v", err)
	}
	defer jsonFile.Close()
	// Read the file content
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	// Unmarshal the JSON content into the Data struct
	var data Data
	if err := json.Unmarshal(byteValue, &data); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	// Extract the effective_action_names
	var effectiveActionNames []string
	for _, action := range data.Actions {
		// log.Printf(action.EffectiveActionNames[0])
		effectiveActionNames = append(effectiveActionNames, action.EffectiveActionNames...)
	}

	return effectiveActionNames

}

func expandPolicy(policy string) []string {
	result_policy := strings.TrimSuffix(policy, "*")
	filteredPoliciesMap := make(map[string]struct{})

	fmt.Println(result_policy)
	for _, str := range actionsNames {
		if strings.HasPrefix(str, result_policy) {
			filteredPoliciesMap[str] = struct{}{}
		}
	}
	var filteredPolicies []string
	fmt.Println(len(filteredPoliciesMap))
	for str := range filteredPoliciesMap {
		filteredPolicies = append(filteredPolicies, str)
	}

	sort.Strings(filteredPolicies)
	return filteredPolicies

}
