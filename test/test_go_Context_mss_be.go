package main

import (
	"GoParser/test/semantic"
	"GoParser/test/utils"
	"strings"
	"sync/atomic"
)

func main() {
	//go func() {
	//	for {
	//		traceMemStats()
	//		time.Sleep(2 * time.Second)
	//	}
	//}()
	utils.WalkDir("/Users/mickey/git/mis-backend/", walk_mss_be_go_file)
}

var goroutineCounter int32 = 0

func walk_mss_be_go_file(filePath string) {
	if strings.HasPrefix(filePath, "/Users/mickey/git/mis-backend/vendor") {
		return
	}
	atomic.AddInt32(&goroutineCounter, 1)
	f := func() {
		if strings.HasSuffix(filePath, ".go") {
			////println("=======================================================================================")
			//println("开始处理：", filePath)
			print(".")
			semantic.WaklTest(filePath)
			//println("=======================================================================================")
		}
		atomic.AddInt32(&goroutineCounter, -1)
	}
	//if goroutineCounter < 50 {
	//	go f()
	//} else {
	f()
	//}
}

//func traceMemStats() {
//	var ms runtime.MemStats
//	runtime.ReadMemStats(&ms)
//	log.Printf("Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
//}
