package closer

import (
	"log"
	"os"
	"os/signal"
	"sync"
)

var globalCloser = New()

// Add to add closer func
func Add(f ...func() error) {
	globalCloser.Add(f...)
}

// Wait wait to close funcs
func Wait() {
	globalCloser.Wait()
}

// CloseAll funcs
func CloseAll() {
	globalCloser.CloseAll()
}

// Closer for close funcs
type Closer struct {
	mu    sync.Mutex
	once  sync.Once
	done  chan struct{}
	funcs []func() error
}

// New construct Closer
func New(sig ...os.Signal) *Closer {
	c := &Closer{done: make(chan struct{})}
	if len(sig) > 0 {
		go func() {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, sig...)
			<-ch
			signal.Stop(ch)
		}()
	}

	return c
}

// Add func
func (c *Closer) Add(f ...func() error) {
	c.mu.Lock()
	c.funcs = append(c.funcs, f...)
	c.mu.Unlock()
}

// Wait channel done
func (c *Closer) Wait() {
	<-c.done
}

// CloseAll run funcs
func (c *Closer) CloseAll() {
	c.once.Do(func() {
		defer close(c.done)

		c.mu.Lock()
		funcs := c.funcs
		c.funcs = nil
		c.mu.Unlock()

		errs := make(chan error, len(funcs))
		for _, f := range funcs {
			go func(f func() error) {
				errs <- f()
			}(f)
		}

		for i := 0; i < cap(errs); i++ {
			if err := <-errs; err != nil {
				log.Printf("error returned from closer:%v", err.Error())
			}
		}
	})
}
