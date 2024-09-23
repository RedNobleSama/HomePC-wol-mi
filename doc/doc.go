package doc

import (
	"encoding/json"
	"fmt"
)

type DocResponse struct {
	Code int `json:"Code"`
	Data struct {
		DataId     string `json:"DataId"`
		PageResult []struct {
			ImageResult []struct {
				Description string `json:"Description"`
				LabelResult []struct {
					Label      string  `json:"label,omitempty"`
					Confidence float64 `json:"Confidence,omitempty"`
					Label1     string  `json:"Label,omitempty"`
				} `json:"LabelResult"`
				Service string `json:"Service"`
			} `json:"ImageResult"`
			ImageUrl   string `json:"ImageUrl"`
			PageNum    int    `json:"PageNum"`
			TextResult []struct {
				Description string `json:"Description"`
				Labels      string `json:"Labels"`
				RiskTips    string `json:"RiskTips"`
				RiskWords   string `json:"RiskWords"`
				Service     string `json:"Service"`
				Text        string `json:"Text"`
			} `json:"TextResult"`
		} `json:"PageResult"`
		Url string `json:"Url"`
	} `json:"Data"`
	Message   string `json:"Message"`
	RequestId string `json:"RequestId"`
}

func ParseResponse() {
	jsonStr := `{
    "Code": 200,
    "Data": {
        "DataId": "fileId*****",
        "PageResult": [
            {
                "ImageResult": [
                    {
                        "Description": "对文档页面的图像内容审核",
                        "LabelResult": [
                            {
                                "label": "nonLabel"
                            }
                        ],
                        "Service": "baselineCheck"
                    }
                ],
                "ImageUrl": "http://oss.aliyundoc.com/a.png",
                "PageNum": 1,
                "TextResult": [
                    {
                        "Description": "对文档页面的文字内容审核",
                        "Labels": "",
                        "RiskTips": "",
                        "RiskWords": "",
                        "Service": "pgc_detection",
                        "Text": "内容安全产品测试用例a"
                    }
                ]
            },
            {
                "ImageResult": [
                    {
                        "Description": "对文档页面的图像内容审核",
                        "LabelResult": [
                            {
                                "Confidence": 89.01,
                                "Label": "pornographic_adultContent_tii"
                            }
                        ],
                        "Service": "baselineCheck"
                    }
                ],
                "ImageUrl": "http://oss.aliyundoc.com/b.png",
                "PageNum": 10,
                "TextResult": [
                    {
                        "Description": "对文档页面的文字内容审核",
                        "Labels": "contraband,sexual_content",
                        "RiskTips": "违禁_违禁商品,色情_影视资源,色情_低俗",
                        "RiskWords": "风险词A,风险词B",
                        "Service": "ad_compliance_detection",
                        "Text": "内容安全产品测试用例b"
                    }
                ]
            }
        ],
        "Url": "http://www.aliyundoc.com/a.docx"
    },
    "Message": "SUCCESS",
    "RequestId": "1D0854A7-AAAAA-BBBBBBB-CC8292AE5"
}`
	var docResponse DocResponse
	err := json.Unmarshal([]byte(jsonStr), &docResponse)
	if err != nil {
		return
	}

	var results []map[string]interface{}
	for _, page := range docResponse.Data.PageResult {
		var pageMap = make(map[string]interface{})
		pageMap["pageNum"] = page.PageNum

		var imageResults []map[string]interface{}
		var textResults []map[string]interface{}
		for _, image := range page.ImageResult {
			var imageMap = make(map[string]interface{})
			imageMap["result"] = image.LabelResult

			imageResults = append(imageResults, imageMap)
		}

		for _, text := range page.TextResult {
			var textMap = make(map[string]interface{})
			textMap["text"] = text.Text
			textMap["riskWords"] = text.RiskWords
			textMap["riskTips"] = text.RiskTips

			textResults = append(textResults, textMap)
		}

		pageMap["imageResults"] = imageResults
		pageMap["textResults"] = textResults

		results = append(results, pageMap)
	}

	marshal, err := json.Marshal(results)
	if err != nil {
		return
	}

	fmt.Println("scan result:", string(marshal))
}
