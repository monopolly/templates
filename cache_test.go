package templates

//testing

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	function, _, _, _ := runtime.Caller(0)
	fn := runtime.FuncForPC(function).Name()
	fmt.Println(fmt.Sprintf("%s\n", fn[strings.LastIndex(fn, ".Test")+5:]))
	a := assert.New(t)
	_ = a

	/* c := newcache()
	c.Add("nice1s", []byte("nice1s"), time.Second)
	c.Add("nice3s", []byte("nice3s"), 3*time.Second)
	c.Add("nice10s", []byte("nice10s"), 10*time.Second)

	a.Equal([]byte("nice1s"), c.Get("nice1s"))
	a.Equal([]byte("nice3s"), c.Get("nice3s"))
	a.Equal([]byte("nice10s"), c.Get("nice10s"))

	time.Sleep(time.Second * 3)

	a.Nil(c.Get("nice1s"))
	a.Nil(c.Get("nice3s"))
	a.NotNil(c.Get("nice10s")) */

}
