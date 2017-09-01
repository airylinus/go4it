package go4it

import (
	"sync"
	"testing"
)

//test
func TestNo1(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)
	//var jobs = make(chan *Resource)
	go InitDownloader(&wg)
	j1 := NewResource(
		"https://c2.staticflickr.com/4/3349/4569314574_8dc6edc83b_b.jpg",
		"/home/mongo/Downloads/demos/log-b",
	)
	j2 := NewResource(
		"http://www.absoluteability.com/wp-content/uploads/2014/07/beachgirls.jpg",
		"/home/mongo/Downloads/demos/log-a",
	)
	j3 := NewResource(
		"https://c1.staticflickr.com/5/4009/4568679755_652cfcf603_b.jpg",
		"/home/mongo/Downloads/demos/log-c",
	)
	Get(j1)
	Get(j2)
	Get(j3)
	wg.Wait()
	// for jb := range chanResult {
	// 	fmt.Println(string(jb.RawData))
	// }
}

//test
func TestDownload(t *testing.T) {
	j1 := NewResource(
		"https://c2.staticflickr.com/4/3349/4569314574_8dc6edc83b_b.jpg",
		"/home/mongo/Downloads/stars/log-b",
	)
	done, err := j1.download()
	if err != nil {
		t.Error("failed when http get : " + err.Error())
	}
	t.Log(done)
	if done <= 0 {
		t.Error("download failed ")
	}
}
