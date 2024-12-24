package utiles

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func PostReq(url string, headers map[string]string, data interface{}) (interface{}, error) {

	jsonStr, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	if headers != nil {
		for key, header := range headers {
			req.Header.Set(key, header)
		}

	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("intervalReq Error: ", err)
		return nil, err
	} else {
		defer resp.Body.Close()
	}

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := io.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
	return resp.Body, nil
}
