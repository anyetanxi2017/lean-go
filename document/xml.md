Golang解析xml数据
https://blog.csdn.net/weixin_30532837/article/details/95505385?utm_medium=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromMachineLearnPai2-1.control&dist_request_id=1328602.1432.16148704768215097&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromMachineLearnPai2-1.control



```

package utils

import (
	"encoding/xml"
	"fmt"
	"gin-vue-admin/global"
	"log"
)

const (
	username = "MBC161"
	password = "QDIY1a4A"
	num      = "852582"
)

type FetchUtil struct {
}

// 单发
func (f FetchUtil) SendMsg(callee, content string) *SendMsgResp {
	randomInt := NumUtil{}.RandomInt(10000, 99999)
	caller := fmt.Sprint(num, randomInt)
	url := fmt.Sprintf("http://127.0.0.1:8138/14.dox?UserName=%s&PassWord=%s&Caller=%s&Callee=%s&CallerAddrTon=1&CallerAddrNpi=1&CalleeAddrTon=1&CalleeAddrNpi=1&CharSet=0&DCS=8&Text=%s",
		username, password, caller, "86"+callee, content)
	data, err := HttpUtil{}.Get(url, nil, nil)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil
	}
	var xmlRes SendMsgResp
	if err = xml.Unmarshal([]byte(data), &xmlRes); err != nil {
		log.Println(err)
	}
	return &xmlRes
}

//<?xml version="1.0" encoding="UTF-8"?>
//<Message>
//<Head>
//<MessageID>SubmitMsgAck</MessageID>
//<Status>2</Status>
//</Head>
//<Body/>
//</Message>
type SendMsgResp struct {
	// 这里要注意xml.Name 这个tag，它表示后面的数据的父元素是什么，如果没有填写这个信息，在数据解析的时候可能会获取不到数据。
	MessageName xml.Name `xml:"Message"`
	Head        struct {
		XMLName   xml.Name `xml:"Head"`
		MessageID string   `xml:"MessageID"`
		Status    int      `xml:"Status"`
	}
}

```
