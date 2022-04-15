package main

import (
	"GoParser/test/semantic"
	"GoParser/test/utils"
	"fmt"
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
	utils.WalkDir("/Users/mickey/git/mis-backend/", walk_mss_be_go_file, "/Users/mickey/git/mis-backend/vendor/")
	wg.Done()

	fmt.Printf("\n\n总耗时：%.3f seconds\n", time.Since(start).Seconds())
}

var goroutineCounter int32 = 0
var wg = sync.WaitGroup{}

func walk_mss_be_go_file(filePath string) {
	atomic.AddInt32(&goroutineCounter, 1)
	wg.Add(1)
	f := func() {
		if strings.HasSuffix(filePath, ".go") {
			////println("=======================================================================================")
			//println("开始处理：", filePath)
			print(".")
			semantic.WaklTest(filePath)
			//println("=======================================================================================")
		}
		atomic.AddInt32(&goroutineCounter, -1)
		wg.Done()
	}
	if goroutineCounter < 20 {
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
