package basicpersonaldata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Fetch : PDからデータを取得する
func (b *BasicPersonalData) Fetch() (*BasicPersonalData, error) {

	req, err := http.NewRequest("GET", "https://pd:8080/api/basics", nil)
	if err != nil {
		fmt.Printf("pd error, cannot create http request")
		return b, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("pd error! cannot exec http request")
		return b, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("pd error! cannot read response body")
		return b, err
	}

	fmt.Println(body)

	jsonBytes := ([]byte)(body)
	replyData := new(BasicPersonalData)

	if err := json.Unmarshal(jsonBytes, replyData); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
	}

	return replyData, nil
}

// Update : 基本データの更新
func (b *BasicPersonalData) Update() (*BasicPersonalData, error) {

	jsonBytes, err := json.Marshal(b)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return b, err
	}

	req, err := http.NewRequest("POST", "https://pd:8080/api/basics", bytes.NewBuffer(jsonBytes))
	if err != nil {
		fmt.Printf("pd error, cannot create http request")
		return b, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("pd error! cannot exec http request")
		return b, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("pd error! cannot read response body")
		return b, err
	}

	fmt.Println(body)

	jsonBytes = ([]byte)(body)
	replyData := new(BasicPersonalData)

	if err := json.Unmarshal(jsonBytes, replyData); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
	}

	return replyData, nil
}
