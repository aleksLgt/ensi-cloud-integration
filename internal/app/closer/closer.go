package closer

import (
	"context"
	"fmt"
	"strings"
	"sync"
)

type Closer struct {
	mu    sync.Mutex
	funcs []Func
}

type Func func(ctx context.Context) error

func (c *Closer) Add(f Func) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.funcs = append(c.funcs, f)
}

func (c *Closer) Close(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	var (
		msgs     = make([]string, 0, len(c.funcs))
		complete = make(chan struct{}, 1) // TODO
	)

	go func() {
		for _, fn := range c.funcs {
			if err := fn(ctx); err != nil {
				msgs = append(msgs, fmt.Sprintf("[!] %v", err))
			}
		}

		complete <- struct{}{}
	}() // TODO погуглить про синтактис вызова функций в горутинах (последние круглые скобки - зачем?)

	select {
	case <-complete:
		break
	case <-ctx.Done():
		return fmt.Errorf("shutdown cancelled: %v", ctx.Err())
	}

	if len(msgs) > 0 {
		return fmt.Errorf(
			"shutdown finished with error(s): \n%s",
			strings.Join(msgs, "\n"),
		)
	}

	return nil
}
