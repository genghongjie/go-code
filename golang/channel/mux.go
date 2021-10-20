package main

import (
	"log"
	"strconv"
	"sync"
	"time"
)

var count = 0

func main() {
	goCache := GoCache{
		Data: make(map[string]interface{}),
		Mux:  new(sync.Mutex),
	}

	done := time.Tick(2 * time.Second)
	for {
		select {
		case <-done:
			for i := 0; i < 5; i++ {

				go func() {
					log.Println(goCache.GetData("A"))
					go PutData(&goCache, "params")

				}()
			}
		}
	}

}

//查询数据
func PutData(gc *GoCache, p string) {
	count++
	str := "数据ID：" + strconv.Itoa(count)
	time.Sleep(3 * time.Second)
	gc.PutData("A", str)

}

type GoCache struct {
	Data map[string]interface{}
	Mux  *sync.Mutex
}

func (gc *GoCache) GetData(key string) interface{} {
	gc.Mux.Lock()
	defer gc.Mux.Unlock()

	if value, ok := gc.Data[key]; ok {
		return value
	} else {

		time.Sleep(3 * time.Second)
		//调用查询数据的 方法
		newValue := "我是初始化数据"

		gc.Data[key] = newValue
		return newValue
	}
}

func (gc *GoCache) PutData(key string, value interface{}) {
	gc.Mux.Lock()
	defer gc.Mux.Unlock()
	gc.Data[key] = value
}
