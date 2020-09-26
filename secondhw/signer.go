package main

import (
	"sort"
	"strconv"
	"sync"
	"strings"
)

func worker(function job, in, out chan interface{}, wg *sync.WaitGroup) {
	function(in, out)
	close(out)
	wg.Done()
}

var mutex sync.Mutex
func checkMd5(out chan string, data string) {
	mutex.Lock()
	out <- DataSignerMd5(data)
	mutex.Unlock()
}

func checkCrc32(out chan string, in chan string) {
	out <- DataSignerCrc32(<-in)
}

type CrcChans struct {
	firstCrc32  chan string
	secondCrc32 chan string
}

func SingleHash(in, out chan interface{}) {
	var chans []CrcChans = make([]CrcChans, 0)

	for data := range in {
		firstCrc32 := make(chan string, 1)
		md5 := make(chan string, 1)
		secondCrc32 := make(chan string, 1)

		chans = append(chans, CrcChans{
			firstCrc32,
			secondCrc32,
		})

		tmp := make(chan string, 1)
		tmp <- strconv.Itoa(data.(int))

		go checkCrc32(firstCrc32, tmp)
		go checkMd5(md5, strconv.Itoa(data.(int)))
		go checkCrc32(secondCrc32, md5)
	}

	for _, value := range chans {
		data := (<-value.firstCrc32 + "~" + <-value.secondCrc32)
		out <- data
	}
}

func MultiHash(in, out chan interface{}) {
	results := make([][]chan string, 0)

	for data := range in {
		var preResult = make([]chan string, 6, 6)
		for i := 0; i < 6; i++ {
			preResult[i] = make(chan string, 1)

			tmp := make(chan string, 1)
			tmp <- strconv.Itoa(i) + data.(string)

			go checkCrc32(preResult[i], tmp)
		}

		results = append(results, preResult)
	}

	for _, ch := range results {
		var result string
		for i := 0; i < 6; i++ {
			result += <-ch[i]
		}
		out <- result
	}
}

func CombineResults(in, out chan interface{}) {
	var data []string
	for itemFromIn := range in {
		value, err := itemFromIn.(string)
		if !err {
			panic("Invalid channel data")
		}
		data = append(data, value)
	}
	sort.Strings(data)

	out <- strings.Join(data, "_");
}

func ExecutePipeline(args ...job) {
	var tmpChanIn chan interface{}
	var tmpChanOut chan interface{}

	wg := sync.WaitGroup{}

	for num, arg := range args {
		wg.Add(1)
		if num % 2 == 0 {
			tmpChanOut = make(chan interface{})
			go worker(arg, tmpChanIn, tmpChanOut, &wg)
		} else {
			tmpChanIn = make(chan interface{})
			go worker(arg, tmpChanOut, tmpChanIn, &wg)
		}
	}

	wg.Wait()
}

func main() {
}
