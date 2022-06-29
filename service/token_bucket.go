package service

import (
	"context"
	"sync/atomic"
	"time"
)

type Bucket struct {
	tokens chan struct{}
	burst  uint
	limit  uint // token 生成速率，1秒多少个
	n      int32
	close  chan struct{}
}

//type Bucket chan struct{}

func NewBucket(burst, limit uint) *Bucket {
	b := &Bucket{
		tokens: make(chan struct{}, burst),
		burst:  burst,
		limit:  limit,
		close:  make(chan struct{}),
	}
	go b.Drop()
	return b
}

func (b *Bucket) Drop() {
	interval := DropInterval(b.limit)
	// keep drop
	for {
		select {
		case <-b.close:
			return
		case b.tokens <- struct{}{}:
			atomic.AddInt32(&b.n, 1)
			time.Sleep(interval * time.Nanosecond)
		}
	}
}

func (b *Bucket) Token() int32 {
	return b.n
}

func (b *Bucket) Pick() bool {
	select {
	case <-b.tokens:
		atomic.AddInt32(&b.n, -1)
		return true
	default:
		return false
	}
}

func (b *Bucket) Wait(ctx context.Context) bool {
	select {
	case <-b.tokens:
		atomic.AddInt32(&b.n, -1)
		return true
	case <-ctx.Done():
		return false
	}
}

func (b *Bucket) Close() {
	close(b.close)
}

// DropInterval interval of Nanosecond to drop a token
func DropInterval(limit uint) time.Duration {
	return (time.Second / time.Nanosecond) / time.Duration(limit)
}
