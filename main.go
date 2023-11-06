package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// create a directory
// if already exists
// delete the old and create a new one
// return -1 when error occurs
// return 0 if succeeds
func createDir(dir string) int {
	_, err := os.Stat(dir)

	if os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			return -1
		}
	} else if err == nil {
		err := os.RemoveAll(dir)
		if err != nil {
			return -1
		}

		err = os.Mkdir(dir, 0755)
		if err != nil {
			return -1
		}
	} else {
		return -1
	}
	return 0
}

// resize sequencially, require input image path
// and output image path
func sequencial() {
	for i := 0; i < 10000; i++ {
		resizeImage(
			fmt.Sprintf("test_%d.JPEG", i),
		)
	}
}

// one thread to push all file names
// close the queue when finished
func enqueueImage(bq *BoundedQueue) {
	for i := 0; i < 10000; i++ {
		bq.enqueue(fmt.Sprintf("test_%d.JPEG", i))
	}
	bq.close()
}

// one thread dequeue and resize
// stop when no image in the queue
func dequeueImage(bq *BoundedQueue) {
	for {
		fileName, err := bq.dequeue()
		if err != 0 {
			break
		} else {
			resizeImage(fileName)
		}
	}
}

var throughputPath = "./throughput/"
var enLatencyPath = "./latency/enqueue/"
var deLatencyPath = "./latency/dequeue/"

func testSeqThroughput() float64 {
	start := time.Now()
	sequencial()
	duration := time.Since(start)
	throughput := float64(10000) / duration.Seconds()
	// fmt.Printf("sequencial throughput: %.2f\n", throughput)
	return throughput
}

func testMultThroughput(threadNum int, capacity int) float64 {
	var bq BoundedQueue
	bq.init(capacity)

	var wg sync.WaitGroup

	start := time.Now()

	wg.Add(1)
	go func() {
		defer wg.Done()
		enqueueImage(&bq)
	}()

	for i := 0; i < threadNum-1; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			dequeueImage(&bq)
		}()
	}

	wg.Wait()

	duration := time.Since(start)
	throughput := float64(10000) / duration.Seconds()
	// fmt.Printf("thread: %d, capacity: %d, time: %.2f, throughput: %.2f\n", threadNum, capacity, duration.Seconds(), throughput)
	return throughput
}

func testThroughput() {
	threadNum := []int{2, 4, 8, 16}
	capacityNum := []int{10, 100, 1000}
	createDir("./output")
	createDir("./throughput")
	filePath := fmt.Sprintf("./throughput/throughput.txt")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	th := testSeqThroughput()
	file.Write([]byte(fmt.Sprintf("%.2f\n", th)))

	for _, i := range threadNum {
		for _, j := range capacityNum {
			th := testMultThroughput(i, j)
			file.Write([]byte(fmt.Sprintf("%d,%d,%.2f\n", i, j, th)))
		}
	}
	file.Close()
}

func testMultLatency(enNum int, deNum int, capacity int) {
	enFilePath := fmt.Sprintf("%senqueue_%d_%d_%d.txt", enLatencyPath, enNum, deNum, capacity)
	deFilePath := fmt.Sprintf("%sdequeue_%d_%d_%d.txt", deLatencyPath, enNum, deNum, capacity)
	enFile, err1 := os.Create(enFilePath)
	deFile, err2 := os.Create(deFilePath)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	var bq BoundedQueue
	bq.init(capacity)

	enLatencyChannel := make(chan int64)
	deLatencyChannle := make(chan int64)

	//set up enqueue thread
	for i := 0; i < enNum; i++ {
		k := i
		go func() {
			for j := (10000 / enNum) * k; j < (10000/enNum)*(k+1); j++ {
				fileName := fmt.Sprintf("test_%d.JPEG", j)
				start := time.Now()
				bq.enqueue(fileName)
				latency := time.Since(start)
				enLatencyChannel <- latency.Microseconds()
			}
		}()
	}

	// set up dequeue thread
	for i := 0; i < deNum; i++ {
		go func() {
			for {
				start := time.Now()
				fileName, err := bq.dequeue()
				latency := time.Since(start)
				if err != 0 {
					break
				} else {
					resizeImage(fileName)
				}
				deLatencyChannle <- latency.Microseconds()
			}
		}()
	}

	for i := 0; i < 10000; i++ {
		en := <-enLatencyChannel
		de := <-deLatencyChannle
		enFile.Write([]byte(fmt.Sprintf("%d\n", en)))
		deFile.Write([]byte(fmt.Sprintf("%d\n", de)))
	}
}

func testLatency() {
	createDir("output")
	createDir("latency")
	os.Mkdir("latency/enqueue", 0755)
	os.Mkdir("latency/dequeue", 0755)
	testMultLatency(1, 10, 100)
	testMultLatency(1, 10, 1000)
	testMultLatency(10, 10, 100)
	testMultLatency(10, 10, 1000)
}

func main() {
	testThroughput()
	testLatency()
}
