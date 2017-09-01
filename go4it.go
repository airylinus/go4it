package go4it

/*--------------------------------*/

import (
	"errors"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

//SupportTypes defines all Image types download supported
var (
	SupportTypes = []string{"jpg", "jpeg"}
	timeout      = 7 * time.Second
	chanJobs     chan *Resource
	defaultJobs  = 20
)

//Resource is a wrape
type Resource struct {
	Meta string
	URL  string
	Path string
}

//NewResource init an Entity
func NewResource(url, path string) *Resource {
	e := Resource{URL: url, Path: path}
	return &e
}

//Get add resource to download queue
func Get(r *Resource) {
	chanJobs <- r
}

func init() {
	InitJobChan(defaultJobs)
}

//InitJobChan init job queue
func InitJobChan(batchSize int) {
	chanJobs = make(chan *Resource, batchSize)
}

//InitDownloader goroutine
func InitDownloader(wg *sync.WaitGroup) {
	for res := range chanJobs {
		go func(r *Resource) {
			if r.IsNeeded() {
				r.download()
			}
			wg.Done()
		}(res)
	}
}

//IsNeeded is
func (res *Resource) IsNeeded() bool {
	u := strings.Split(res.URL, ".")
	ext := u[len(u)-1]
	for _, v := range SupportTypes {
		if v == ext {
			res.Path += "." + ext
			return true
		}
	}
	return false
}

//download resource via URI
func (res *Resource) download() (int64, error) {

	t := time.Duration(timeout)
	client := http.Client{Timeout: t}
	request, err := http.NewRequest("GET", res.URL, nil)
	if err != nil {
		return 0, errors.New("error format of URL " + err.Error())
	}
	request.Header.Add("user-agent", GetRandomUserAgent())
	response, err := client.Do(request)
	if err != nil {
		return 0, errors.New("failed to fetch images" + err.Error())
	}
	length := response.ContentLength
	if length < 200 {
		return 0, errors.New("")
	}
	//panic(respHeader)
	f, err := os.Create(res.Path)
	if err != nil {
		//checkError(err)
		return 0, errors.New("failed create file : " + res.Path)
	}
	l, err := io.Copy(f, response.Body)
	if err != nil {
		//checkError(err)
		return 0, errors.New("failed copy from response body : " + err.Error())
	}
	if l < length {
		panic(response.Header)
	}
	return l, nil
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

var uaPool = [...]string{"Mozilla/5.0 (compatible, MSIE 10.0, Windows NT, DigExt)",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, 360SE)",
	"Mozilla/4.0 (compatible, MSIE 8.0, Windows NT 6.0, Trident/4.0)",
	"Mozilla/5.0 (compatible, MSIE 9.0, Windows NT 6.1, Trident/5.0,",
	"Opera/9.80 (Windows NT 6.1, U, en) Presto/2.8.131 Version/11.11",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, TencentTraveler 4.0)",
	"Mozilla/5.0 (Windows, U, Windows NT 6.1, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Macintosh, Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
	"Mozilla/5.0 (Macintosh, U, Intel Mac OS X 10_6_8, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	" Mozilla/5.0 (Linux, U, Android 3.0, en-us, Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
	"Mozilla/5.0 (iPad, U, CPU OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, Trident/4.0, SE 2.X MetaSr 1.0, SE 2.X MetaSr 1.0, .NET CLR 2.0.50727, SE 2.X MetaSr 1.0)",
	"Mozilla/5.0 (iPhone, U, CPU iPhone OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"MQQBrowser/26 Mozilla/5.0 (Linux, U, Android 2.3.7, zh-cn, MB200 Build/GRJ22, CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1"}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

//GetRandomUserAgent returns a UA string by random
func GetRandomUserAgent() string {
	return uaPool[r.Intn(len(uaPool))]
}
