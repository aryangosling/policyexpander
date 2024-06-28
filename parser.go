package policyverse

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"reflect"
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
	jsonFile, err := os.Open("data/aws/iam-dataset.json")
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

func _expandPolicy(policy string, filteredPoliciesMap map[string]struct{}) []string {
	result_policy := strings.TrimSuffix(policy, "*")
	for _, str := range actionsNames {
		// fmt.Println(str, policy)
		if strings.HasPrefix(str, result_policy) {
			filteredPoliciesMap[str] = struct{}{}
		}
	}
	var filteredPolicies []string

	for str := range filteredPoliciesMap {
		filteredPolicies = append(filteredPolicies, str)
	}

	sort.Strings(filteredPolicies)
	return filteredPolicies
}

func isArray(obj interface{}) bool {
	return reflect.TypeOf(obj).Kind() == reflect.Array || reflect.TypeOf(obj).Kind() == reflect.Slice
}

// ensureArray ensures that the given object is an array,
// by creating a slice and adding it as a single element if it's not already an array or slice.
func ensureArray(obj interface{}) []interface{} {
	if isArray(obj) {
		return obj.([]interface{})
	}
	return []interface{}{obj}
}

func ExpandPolicy(policy map[string]interface{}) []string {
	filteredPoliciesMap := make(map[string]struct{})

	if statements, ok := policy["Statement"].([]interface{}); ok {
		for _, statement := range statements {
			// Check if each statement is a map
			if stmtMap, ok := statement.(map[string]interface{}); ok {
				// Check if the "Action" key exists and is a slice
				if actions, ok := stmtMap["Action"]; ok {
					actions := ensureArray(actions)
					for _, action := range actions {
						action := action.(string)
						// Print each action

						if strings.HasSuffix(action, "*") {
							_expandPolicy(action, filteredPoliciesMap)
						} else {
							filteredPoliciesMap[action] = struct{}{}
						}

					}
				}
			}
		}
	}

	var filteredPolicies []string
	for str := range filteredPoliciesMap {
		filteredPolicies = append(filteredPolicies, str)
	}

	sort.Strings(filteredPolicies)
	return filteredPolicies

}
