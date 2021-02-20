```go
type Datetime time.Time

func (t Datetime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

```


## 重写 MarshalJson
https://colobu.com/2020/03/19/Custom-JSON-Marshalling-in-Go/


