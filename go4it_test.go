package go4it

import (
	"fmt"
	"sync"
	"testing"
)

//test
func TestNo1(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)
	var jobs = make(chan *Resource)
	go InitDownloader(jobs, &wg)
	j1 := NewResource(
		"https://img1.doubanio.com/icon/u1399511-7.jpg",
		"/home/mongo/Downloads/test/log-b",
	)
	j2 := NewResource(
		"https://img3.doubanio.com/icon/u2329213-10.jpg",
		"/home/mongo/Downloads/test/log-a",
	)
	j3 := NewResource(
		"https://img3.doubanio.com/icon/u13694634-3.jpg",
		"/home/mongo/Downloads/test/log-c",
	)
	jobs <- j2
	jobs <- j1
	jobs <- j3
	wg.Wait()
	// for jb := range chanResult {
	// 	fmt.Println(string(jb.RawData))
	// }
}

//test
func TestRemoteFetch(t *testing.T) {
	j1 := NewResource(
		"https://ss0.bdstatic.com/5aV1bjqh_Q23odCf/static/superman/img/logo/logo_white_fe6da1ec.png",
		"/home/mongo/Downloads/test/log-b",
	)
	response, err := remoteFetch(j1)
	if err != nil {
		t.Error("failed when http get")
	}
	if response == nil {
		t.Error("nil return")
	}
	if len(response) == 0 {
		t.Error("empty response")
	}
	fmt.Print(string(response))
}
