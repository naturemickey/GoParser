package main

import (
	"GoParser/test/utils"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//go func() {
	//	for {
	//		traceMemStats()
	//		time.Sleep(2 * time.Second)
	//	}
	//}()
	start := time.Now()
	utils.WalkDir("/Users/mickey/git/mis-backend/", asdfadsfasdfasdfdf, "/Users/mickey/git/mis-backend/vendor/")
	_wg.Wait()

	fmt.Printf("\n\n文件数：%d\n", 文件数)
	fmt.Printf("行数(包括空行)：%d\n", 行数_包括空行)
	fmt.Printf("行数(不包括空行)：%d\n", 行数_不包括空行)
	fmt.Printf("\n\n总耗时：%.3f seconds\n", time.Since(start).Seconds())
}

var _goroutineCounter int32 = 0
var _wg = sync.WaitGroup{}
var 文件数 int32
var 行数_包括空行 int32
var 行数_不包括空行 int32

func asdfadsfasdfasdfdf(filePath string) {
	atomic.AddInt32(&_goroutineCounter, 1)
	_wg.Add(1)
	f := func() {
		if strings.HasSuffix(filePath, ".go") {
			print(".")
			atomic.AddInt32(&文件数, 1)
			file, err := ioutil.ReadFile(filePath)
			if err != nil {
				panic(err.Error())
			}
			code := string(file)

			lines := strings.Split(code, "\n")

			atomic.AddInt32(&行数_包括空行, int32(len(lines)))

			for _, line := range lines {
				line = strings.TrimSpace(line)
				if len(line) > 0 {
					atomic.AddInt32(&行数_不包括空行, 1)
				}
			}
		}
		atomic.AddInt32(&_goroutineCounter, -1)
		_wg.Done()
	}
	if _goroutineCounter < 16 {
		go f()
	} else {
		f()
	}
}

//func traceMemStats() {
//	var ms runtime.MemStats
//	runtime.ReadMemStats(&ms)
//	log.Printf("Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
//}
