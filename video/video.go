package video

import (
	"encoding/json"
	"fmt"
)

type VideoResponse struct {
	Code      int    `json:"Code"`
	RequestId string `json:"RequestId"`
	Data      struct {
		AudioResult struct {
			AudioSummarys []struct {
				Label    string `json:"Label"`
				LabelSum int    `json:"LabelSum"`
			} `json:"AudioSummarys"`
		} `json:"AudioResult"`
		FrameResult struct {
			FrameNum      int `json:"FrameNum"`
			FrameSummarys []struct {
				Label    string `json:"Label"`
				LabelSum int    `json:"LabelSum"`
			} `json:"FrameSummarys"`
			Frames []struct {
				Offset  int `json:"Offset"`
				Results []struct {
					Result []struct {
						Label      string  `json:"Label"`
						Confidence float64 `json:"Confidence,omitempty"`
					} `json:"Result"`
					Service string `json:"Service"`
				} `json:"Results"`
				TempUrl string `json:"TempUrl"`
			} `json:"Frames"`
		} `json:"FrameResult"`
	} `json:"Data"`
}

func ParseResponse() (*VideoResponse, error) {
	jsonStr := `{
    "Code": 200,
    "RequestId": "25106421-XXXX-XXXX-XXXX-15DA5AAAC546",
    "Message": "success finished",
    "Data": {
        "DataId": "dc16c28f-xxxx-xxxx-xxxx-51efe0131080",
        "TaskId": "AAAAA-BBBBB-2024-0307-0728",
        "AudioResult": {
            "AudioSummarys": [
                {
                    "Label": "sexual_sounds",
                    "LabelSum": 3
                }
            ],
            "SliceDetails": [
                {
                    "EndTime": 60,
                    "EndTimestamp": 1698912813192,
                    "Labels": "",
                    "StartTime": 30,
                    "StartTimestamp": 1698912783192,
                    "Text": "内容安全",
                    "Url": "http://abc.oss-cn-shanghai.aliyuncs.com/test.wav"
                },
                {
                    "EndTime": 30,
                    "EndTimestamp": 1698912813192,
                    "Extend": "{\"customizedWords\":\"服务\",\"customizedLibs\":\"test\"}",
                    "Labels": "C_customized",
                    "StartTime": 0,
                    "StartTimestamp": 1698912783192,
                    "Text": "欢迎使用阿里云内容安全服务",
                    "Url": "http://abc.oss-cn-shanghai.aliyuncs.com/test.wav"
                }
            ]
        },
        "FrameResult": {
            "FrameNum": 2,
            "FrameSummarys": [
                {
                    "Label": "violent_explosion",
                    "LabelSum": 8
                },
                {
                    "Label": "sexual_cleavage",
                    "LabelSum": 5
                }
            ],
            "Frames": [
                {
                    "Offset": 1,
                    "Results": [
                        {
                            "Result": [
                                {
                                    "Label": "nonLabel"
                                }
                            ],
                            "Service": "baselineCheck_global"
                        }
                    ],
                    "TempUrl": "http://abc.oss-ap-southeast-1.aliyuncs.com/test1.jpg"
                },
                {
                    "Offset": 2,
                    "Results": [
                        {
                            "Result": [
                                {
                                    "Confidence": 1,
                                    "Label": "sexual_cleavage"
                                },
                                {
                                    "Confidence": 74.1,
                                    "Label": "violent_explosion"
                                }
                            ],
                            "Service": "baselineCheck_global"
                        }
                    ],
                    "TempUrl": "http://abc.oss-ap-southeast-1.aliyuncs.com/test2.jpg"
                }
            ]
        }
    }
}`

	var resp VideoResponse
	err := json.Unmarshal([]byte(jsonStr), &resp)
	if err != nil {
		return nil, err
	}

	if len(resp.Data.AudioResult.AudioSummarys) > 0 || len(resp.Data.FrameResult.FrameSummarys) > 0 {
		fmt.Println("risklevel:", "high")
	}

	var results []map[string]interface{}
	for _, v := range resp.Data.AudioResult.AudioSummarys {
		var resultMap = make(map[string]interface{})
		resultMap["标签"] = v.Label
		resultMap["出现次数"] = v.LabelSum
		results = append(results, resultMap)
	}

	for _, v := range resp.Data.FrameResult.FrameSummarys {
		var resultMap = make(map[string]interface{})
		resultMap["标签"] = v.Label
		resultMap["出现次数"] = v.LabelSum
		results = append(results, resultMap)
	}

	marshal, err := json.Marshal(results)
	if err != nil {
		return nil, err
	}

	fmt.Println("scan result:", string(marshal))
	return &resp, nil

}
