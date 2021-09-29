/**
 * Copyright (2021, ) Institute of Software, Chinese Academy of Sciences
 */
package kubesys

import (
	"encoding/json"
	jsonObj "github.com/kubesys/kubernetes-client-go/pkg/json"
)

/**
 *      author: wuheng@iscas.ac.cn
 *      date  : 2021/9/30
 */
func ToJsonObject(bytes []byte) *jsonObj.JsonObject {
	json, err := jsonObj.ParseObject(string(bytes))
	if err != nil {
		return nil
	}
	return json
}

func ToGolangMap(bytes []byte) map[string]interface{} {
	values := make(map[string]interface{})
	json.Unmarshal([]byte(bytes), &values)
	return values
}