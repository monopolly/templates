package templates

import (
	"strconv"
	"sync"
	"time"

	"github.com/monopolly/encode/xxhash"
)

const (
	shards = 32000
	format = "0601021504" //YYMMDDHHMM
)

type cache struct {
	sync.RWMutex
	cache []map[uint32][]byte //id = body
	times map[uint32][]uint32 //time = id
}

func newcache() (a *cache) {
	a = &cache{
		cache: make([]map[uint32][]byte, shards),
		times: make(map[uint32][]uint32),
	}
	go a.timer()
	return
}

//время жизни в минутах
func (a *cache) Add(key string, value []byte, ttl ...time.Duration) {
	id := hash(key)
	a.Lock()
	if a.cache[id%shards] == nil {
		a.cache[id%shards] = make(map[uint32][]byte)
	}
	a.cache[id%shards][id] = value
	if len(ttl) > 0 {
		t := times(
			time.Now().Add(ttl[0]).Format(format),
		)
		a.times[t] = append(a.times[t], id)
	}
	a.Unlock()
}

func (a *cache) Get(key string) (r []byte) {
	id := hash(key)
	a.RLock()
	r = a.cache[id%shards][id]
	a.RUnlock()
	return
}

func (a *cache) timer() {
	for {
		t := times(time.Now().Add(-1 * time.Minute).Format(format))
		if len(a.times[t]) > 0 {
			a.Lock()
			for _, id := range a.times[t] {
				delete(a.cache[id%shards], id)
			}
			delete(a.times, t)
			a.Unlock()
		}
		time.Sleep(1 * time.Minute)
	}
}

func hash(v string) uint32 {
	return xxhash.Checksum32([]byte(v))
}

func times(s string) uint32 {
	num, _ := strconv.Atoi(s)
	return uint32(num)
}
