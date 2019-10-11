package basicpersonaldata

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Fetch : PDからデータを取得する
func (b *BasicPersonalData) Fetch() (*BasicPersonalData, error) {

	req, err := http.NewRequest("GET", "http://pd:8080/api/basics", nil)
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

	var r io.Reader = resp.Body
	// r = io.TeeReader(r, os.Stderr)

	replyData := new(BasicPersonalData)

	if err := json.NewDecoder(r).Decode(replyData); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
	}

	// fmt.Println(replyData)
	return replyData, nil
}

// Update : 基本データの更新
func (b *BasicPersonalData) Update(value *UpdateBasicPersonalData) error {

	switch value.Column {
	case ID:
		b.ID = value.Value
	case Name:
		b.Name = value.Value
	case Birthday:
		b.Birthday = value.Value // TODO : 型変換とか値の変換はここでやる
	case Uncategorized:
		return errors.New("update columns undefined")
	}

	jsonBytes, err := json.Marshal(b)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return err
	}

	req, err := http.NewRequest("POST", "http://pd:8080/api/basics", bytes.NewBuffer(jsonBytes))
	if err != nil {
		fmt.Printf("pd error, cannot create http request")
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("pd error! cannot exec http request")
		return err
	}
	defer resp.Body.Close()

	return nil
}
