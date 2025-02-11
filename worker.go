package smartid

import (
	"context"
	"sync"

	"github.com/tab/smartid/internal/config"
)

const (
	Concurrency = 10
	QueueSize   = 100
)

type Result struct {
	Person *Person
	Err    error
}

type Job struct {
	ctx       context.Context
	sessionId string
	resultCh  chan Result
}

type BackgroundWorker interface {
	Start(ctx context.Context)
	Stop()
	Process(ctx context.Context, sessionId string) <-chan Result
	WithConfig(cfg config.WorkerConfig) *Worker
}

type Worker struct {
	provider    Provider
	queue       chan Job
	concurrency int
	wg          sync.WaitGroup
}

func NewWorker(provider Provider) *Worker {
	return &Worker{
		provider:    provider,
		queue:       make(chan Job, QueueSize),
		concurrency: Concurrency,
	}
}

func (w *Worker) WithConfig(cfg config.WorkerConfig) *Worker {
	if cfg.Concurrency == 0 {
		cfg.Concurrency = Concurrency
	}

	if cfg.QueueSize == 0 {
		cfg.QueueSize = QueueSize
	}

	w.concurrency = cfg.Concurrency
	w.queue = make(chan Job, cfg.QueueSize)

	return w
}

func (w *Worker) Start(ctx context.Context) {
	for i := 0; i < w.concurrency; i++ {
		w.wg.Add(1)
		go w.perform(ctx)
	}
}

func (w *Worker) Stop() {
	close(w.queue)
	w.wg.Wait()
	w.queue = nil
}

func (w *Worker) Process(ctx context.Context, sessionId string) <-chan Result {
	resultCh := make(chan Result, 1)

	select {
	case <-ctx.Done():
		resultCh <- Result{Err: ctx.Err()}
		close(resultCh)
	case w.queue <- Job{ctx: ctx, sessionId: sessionId, resultCh: resultCh}:
	}

	return resultCh
}

func (w *Worker) perform(ctx context.Context) {
	defer w.wg.Done()

	for {
		select {
		case j, ok := <-w.queue:
			if !ok {
				return
			}

			person, err := w.provider.FetchSession(ctx, j.sessionId)
			j.resultCh <- Result{Person: person, Err: err}
			close(j.resultCh)
		case <-ctx.Done():
			return
		}
	}
}
