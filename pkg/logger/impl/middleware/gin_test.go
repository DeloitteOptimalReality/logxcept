package middleware

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestGinZapMiddleware(t *testing.T) {
	a := GinZapMiddleware()

	if a == nil {
		t.Fail()
	}
}

func testTraceHandler(t *testing.T) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		zl := RetrieveLogger(ctx)
		if zl == nil {
			t.Fail()
		}
		tr := ctx.Value(RequestLevelTrace).(string)
		c.Writer.Write([]byte(tr))
	}
}

func TestGinZapMiddlewareWithTrace(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	a := GinZapMiddlewareWithTrace()
	if a == nil {
		t.Fail()
	}
	r := gin.Default()
	r.GET("/trace", a, testTraceHandler(t))
	r.Use(a)
	go func() {
		defer wg.Done()
		if err := http.ListenAndServe("localhost:8080", r); err != nil {
			t.Error(err)
		}
	}()

	time.Sleep(100 * time.Millisecond)

	go func() {
		defer wg.Done()
		if !queryTestServer(t) {
			t.Fail()
		}
	}()

	wg.Wait()
}

func queryTestServer(t *testing.T) bool {
	a, e := http.Get("http://0.0.0.0:8080/trace")
	if e != nil || a == nil {
		t.Error(e)
		return false
	}
	defer a.Body.Close()

	if a.StatusCode != http.StatusOK {
		t.Errorf("error getting correct status code for query")
		return false
	}

	body, err := ioutil.ReadAll(a.Body)
	if err != nil {
		t.Error(err)
		return false
	}
	if len(string(body)) != 16 {
		t.Errorf("length of trace incorrect")
	}
	return true

}
