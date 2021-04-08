// Copyright (2021, ) Institute of Software, Chinese Academy of Sciences
package kubesys

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

//  author: wuheng@iscas.ac.cn
//  date: 2021/4/8
type KubernetesClient struct {
	Url        string
	Token      string
	Http       *http.Client
}

func NewKubernetesClient(url string, token string) *KubernetesClient {
	client := new(KubernetesClient)
	client.Url = url
	client.Token = token
	client.Http = &http.Client {Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}
	return client
}

func (client *KubernetesClient) RequestResource(request *http.Request) (map[string]interface{}, error) {
	res, err := client.Http.Do(request)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var result = make(map[string]interface{})
	json.Unmarshal(body, &result)
	return result, nil
}

func (client *KubernetesClient) CreateRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer " + client.Token)
	return req, nil
}

func GetMapFromMap(values map[string]interface{}, key string) map[string]interface{} {
	return values[key].(map[string]interface{})
}

func GetArrayFromMap(values map[string]interface{}, key string) []interface{} {
	return values[key].([]interface{})
}

