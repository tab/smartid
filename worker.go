package smartid

import (
	"context"
	"sync"
)

const (
	DefaultConcurrency = 10
	DefaultQueueSize   = 100
)

type Result struct {
	Person *Person
	Err    error
}

type Job struct {
	SessionId string
	ResultCh  chan Result
}

type Worker interface {
	Start(ctx context.Context)
	Stop()
	Process(ctx context.Context, sessionId string) <-chan Result

	WithConcurrency(concurrency int) Worker
	WithQueueSize(size int) Worker
}

type worker struct {
	client      Client
	queue       chan Job
	concurrency int
	wg          sync.WaitGroup
}

func NewWorker(client Client) Worker {
	return &worker{
		client:      client,
		queue:       make(chan Job, DefaultQueueSize),
		concurrency: DefaultConcurrency,
	}
}

func (w *worker) WithConcurrency(concurrency int) Worker {
	if concurrency <= 0 {
		concurrency = DefaultConcurrency
	}

	w.concurrency = concurrency
	return w
}

func (w *worker) WithQueueSize(size int) Worker {
	if size <= 0 {
		size = DefaultQueueSize
	}

	w.queue = make(chan Job, size)
	return w
}

func (w *worker) Start(ctx context.Context) {
	for i := 0; i < w.concurrency; i++ {
		w.wg.Add(1)
		go w.perform(ctx)
	}
}

func (w *worker) Stop() {
	close(w.queue)
	w.wg.Wait()
}

func (w *worker) Process(ctx context.Context, sessionId string) <-chan Result {
	resultCh := make(chan Result, 1)

	select {
	case <-ctx.Done():
		resultCh <- Result{Err: ctx.Err()}
		close(resultCh)
	case w.queue <- Job{SessionId: sessionId, ResultCh: resultCh}:
	}

	return resultCh
}

func (w *worker) perform(ctx context.Context) {
	defer w.wg.Done()

	for {
		select {
		case j, ok := <-w.queue:
			if !ok {
				return
			}

			person, err := w.client.FetchSession(ctx, j.SessionId)
			j.ResultCh <- Result{Person: person, Err: err}

			close(j.ResultCh)
		case <-ctx.Done():
			return
		}
	}
}
