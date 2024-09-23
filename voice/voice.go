package voice

import (
	"encoding/json"
	"fmt"
)

type VoiceResponse struct {
	Code int `json:"Code"`
	Data struct {
		SliceDetails []struct {
			EndTime   int    `json:"EndTime"`
			Labels    string `json:"Labels"`
			StartTime int    `json:"StartTime"`
			Text      string `json:"Text"`
			Url       string `json:"Url"`
		} `json:"SliceDetails"`
	} `json:"Data"`
	Message   string `json:"Message"`
	RequestId string `json:"RequestId"`
}

func ParseResponse() {
	jsonStr := `{
    "Code": 200,
    "Data": {
        "SliceDetails": [
            {
                "EndTime": 4065,
                "Labels": "political_content,xxxx",
                "StartTime": 0,
                "Text": "恶心的",
                "Url": "https://aliyundoc.com"
            }
        ]
    },
    "Message": "OK",
    "RequestId": "AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****"
}`
	var response VoiceResponse
	err := json.Unmarshal([]byte(jsonStr), &response)
	if err != nil {
		panic(err)
	}

	var results []map[string]interface{}
	for _, v := range response.Data.SliceDetails {
		result := make(map[string]interface{})
		result["start_time"] = v.StartTime
		result["end_time"] = v.EndTime
		result["text"] = v.Text
		result["labels"] = v.Labels
		result["url"] = v.Url
		results = append(results, result)
	}

	marshal, err := json.Marshal(results)
	if err != nil {
		return
	}

	fmt.Println("scan result:", string(marshal))
}
