
## 观察者模式
```
package main

import "fmt"

func main() {
	var sub Subject = &ViperChangeNotice{}
	customer := Customer{1}
	sub.Register(&customer)
	sub.Register(&Customer{2})
	sub.DeRegister(&customer)
	sub.NotifyAll()
}

// 观察者
type ObServer interface {
	Update(data interface{})
	Action() // 具体执行
	GetId() interface{}
}

// 具体观察者
type Customer struct {
	id int
}

func (c Customer) Update(data interface{}) {

}
func (c *Customer) Action() {
	fmt.Println(c.id, "hello")
}
func (c *Customer) GetId() interface{} {
	return c.id
}

// 主题
type Subject interface {
	Register(server ObServer)
	DeRegister(server ObServer)
	NotifyAll()
}

// 具体主体
type ViperChangeNotice struct {
	Observers []ObServer
}

func (v *ViperChangeNotice) Register(server ObServer) {
	v.Observers = append(v.Observers, server)
}
func (v *ViperChangeNotice) DeRegister(server ObServer) {
	for i, item := range v.Observers {
		if item.GetId() == server.GetId() {
			v.Observers = append(v.Observers[:i], v.Observers[i+1:]...)
			return
		}
	}
}
func (v *ViperChangeNotice) NotifyAll() {
	for _, item := range v.Observers {
		item.Action()
	}
}

```
