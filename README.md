# go4it

## example

```golang
import "github.com/airylinus/go4it"

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
        //uncomment this line will set size of job queue to 10, default is 20
        //go4it.InitJobChan(10)
	go go4it.InitDownloader(&wg)
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
	go4it.Get(j1)
	go4it.Get(j2)
	go4it.Get(j3)
	wg.Wait()
}
```
