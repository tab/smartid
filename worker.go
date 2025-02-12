package smartid

import (
	"context"
	"sync"

	"github.com/tab/smartid/internal/config"
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
	sessionId string
	resultCh  chan Result
}

type Worker interface {
	Start(ctx context.Context)
	Stop()
	Process(ctx context.Context, sessionId string) <-chan Result
	WithConfig(cfg config.WorkerConfig) Worker
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

func (w *worker) WithConfig(cfg config.WorkerConfig) Worker {
	if cfg.Concurrency == 0 {
		cfg.Concurrency = DefaultConcurrency
	}

	if cfg.QueueSize == 0 {
		cfg.QueueSize = DefaultQueueSize
	}

	w.concurrency = cfg.Concurrency
	w.queue = make(chan Job, cfg.QueueSize)

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
	case w.queue <- Job{sessionId: sessionId, resultCh: resultCh}:
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

			person, err := w.client.FetchSession(ctx, j.sessionId)
			j.resultCh <- Result{Person: person, Err: err}

			close(j.resultCh)
		case <-ctx.Done():
			return
		}
	}
}
