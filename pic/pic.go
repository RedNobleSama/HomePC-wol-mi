package pic

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Label       string  `json:"Label"`
	Confidence  float64 `json:"Confidence"`
	Description string  `json:"Description"`
}

type Data struct {
	Result    []Result `json:"Result"`
	RiskLevel string   `json:"RiskLevel"`
}

type Response struct {
	Data      Data   `json:"Data"`
	RequestId string `json:"RequestId"`
}

// 定一个方法解析结构体
func ParseResponse() {
	jsonStr := `{
   "Msg": "success",
   "Code": 200,
   "Data": {
       "DataId": "img123****",
       "Result": [
           {
               "Label": "pornographic_adultContent",
               "Confidence": 81,
               "Description": "成人色情"
           },
           {
               "Label": "sexual_partialNudity",
               "Confidence": 98,
               "Description": "肢体裸露或性感"
           },
           {
               "Label": "violent_explosion",
               "Confidence": 70,
               "Description": "烟火类内容"
           },
           {
               "Label": "violent_explosion_lib",
               "Confidence": 81,
               "Description": "烟火类内容_命中自定义库"
           }
       ],
       "RiskLevel": "high",
       "Frame": "[{\"Result\":[{\"Confidence\":98.18,\"Label\":\"contraband_gamble\"},{\"Confidence\":96.39,\"Label\":\"pornographic_adultContent\"},{\"Confidence\":95.27,\"Label\":\"violent_explosion\"}],\"TempUrl\":\"http://www.aliyundoc.com/test1.jpg\"},{\"Result\":[{\"Confidence\":91.18,\"Label\":\"violent_explosion_lib\"}],\"TempUrl\":\"http://www.aliyundoc.com/test2.jpg\"}]",
       "FrameNum": 2
   },
   "RequestId": "ABCD1234-1234-1234-1234-123****"
}`
	var resp Response
	err := json.Unmarshal([]byte(jsonStr), &resp)
	if err != nil {
		return
	}

	fmt.Println("riskLevel:", resp.Data.RiskLevel)

	var results []map[string]interface{}
	for _, v := range resp.Data.Result {
		var resultMap = make(map[string]interface{})
		resultMap["标签"] = v.Label
		resultMap["置信度"] = v.Confidence
		resultMap["描述"] = v.Description
		results = append(results, resultMap)
	}

	marshal, err := json.Marshal(results)
	if err != nil {
		return
	}

	fmt.Println("scan result:", string(marshal))
	return

}
