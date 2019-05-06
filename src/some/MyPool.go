package some

import (
	"sync"
	"sync/atomic"
	"time"
)

type Mystruct struct {
	config interface{}
	mypool *MyPool
}

type MyPool struct {
	config    interface{}
	queue     chan *MyConn
	curcount  int64
	maxcount  int64
	initcount int64
	ctimeout  time.Duration
	rtimeout  time.Duration
	wtimeout  time.Duration
	close     int
}

type MyConn struct {
	config interface{}
	conn   interface{}
	//config change
	//pool *MyPool{}
}

var mystruct *Mystruct
var oncelock sync.Once

func GetMystruct() *Mystruct {
	if mystruct != nil {
		return mystruct
	}
	oncelock.Do(func() {
		mystruct := &Mystruct{
			config: nil,
		}
		mystruct.Open()
	})
	return mystruct
}

func maininit() {
	mystruct := &Mystruct{
		config: nil,
	}
	mystruct.Open()
}
func (ms *Mystruct) Open() {
	ms.mypool = &MyPool{
		config: nil,
		queue:  make(chan *MyConn, ms.mypool.maxcount),
	}
	ms.mypool.Init()
}
func (ms *Mystruct) Do(cmd string, args ...interface{}) {
	conn := ms.mypool.Get()
	//conn.Do()
	ms.mypool.Put(conn)
}

func (mp *MyPool) Init() {
	for i := int64(0); i < mp.initcount; i++ {
		mp.queue <- &MyConn{}
	}
}
func (mp *MyPool) Get() *MyConn {

	select {
	case conn := <-mp.queue:
		return conn
	default:
	}

	if atomic.AddInt64(&mp.curcount, 1) > mp.maxcount {
		atomic.AddInt64(&mp.curcount, -1)
		timec := time.NewTimer(mp.ctimeout)
		select {
		case conn := <-mp.queue:
			return conn
		case <-timec.C:
			return nil
		}
	}
	conn := &MyConn{}
	return conn
}
func (mp *MyPool) Put(t *MyConn) {

	//不是这个pool的，直接关todo

	if mp.close == 1 {
		//conn close
		atomic.AddInt64(&mp.curcount, -1)
	}
	select {
	case mp.queue <- t:
		return
	default:
		//conn close
		atomic.AddInt64(&mp.curcount, -1)
	}
}
func (mp *MyPool) Close() {
	mp.close = 1
	atomic.StoreInt64(&mp.curcount, mp.maxcount+1)
	for len(mp.queue) > 0 {
		select {
		case conn := <-mp.queue:
			//conn close
		default:
			break
		}
	}
}
