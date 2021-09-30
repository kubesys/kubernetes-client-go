/**
 * Copyright (2021, ) Institute of Software, Chinese Academy of Sciences
 */
package kubesys

import (
	"encoding/json"
	"strings"
)

/**
 *      author: wuheng@iscas.ac.cn
 *      date  : 2021/9/30
 */
func extract(client KubernetesClient, registry *Registry) {
	registryRequest, _ := client.CreateRequest("GET", client.Url, nil)
	registryStringValues, _ := client.RequestResource(registryRequest)

	registryValues := make(map[string]interface{})
	json.Unmarshal([]byte(registryStringValues), &registryValues)

	for _, v := range registryValues["paths"].([]interface{}) {
		path := v.(string)
		if strings.HasPrefix(path, "/api") && (len(strings.Split(path, "/")) == 4 || strings.EqualFold(path, "/api/v1")) {
			resourceRequest, _ := client.CreateRequest("GET", client.Url+path, nil)
			resourceStringValues, _ := client.RequestResource(resourceRequest)

			resourceValues := make(map[string]interface{})
			json.Unmarshal([]byte(resourceStringValues), &resourceValues)

			apiVersion := resourceValues["groupVersion"].(string)
			for _, w := range resourceValues["resources"].([]interface{}) {
				Register(client.Url + path, registry.RuleBase, w.(map[string]interface{}), apiVersion)
			}
		}
	}
}
