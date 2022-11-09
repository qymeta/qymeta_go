package util

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"sort"
)

var host = "http://qymeta_api.newmin.cn/api/"

func GenerateApiSign(maps map[string]string) string {
	var keys []string
	for k, _ := range maps {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	newsMap := make(map[string]string)
	for _, k := range keys {
		newsMap[k] = maps[k]
	}

	fmt.Println(newsMap)
	marshal, _ := json.Marshal(newsMap)
	fmt.Println("marshal", string(marshal)+newsMap["secret"])
	newSig := md5.Sum([]byte(fmt.Sprintf("%x", md5.Sum([]byte(string(marshal)+newsMap["secret"])))))
	fmt.Println("newSig", fmt.Sprintf("%x", newSig))
	return fmt.Sprintf("%x", newSig)

}

func PostJsonRequest(uri string, params map[string]string) interface{} {
	path := host + uri
	bytesData := &bytes.Buffer{}
	writer := multipart.NewWriter(bytesData)
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	_ = writer.Close()
	req, err := http.NewRequest("POST", path, bytesData)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	if err != nil {
		return err.Error()
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)

}
