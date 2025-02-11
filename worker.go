package smartid

import (
	"context"
	"sync"

	"github.com/tab/smartid/internal/config"
	"github.com/tab/smartid/internal/models"
)

const (
	Concurrency = 10
	QueueSize   = 100
)

type Result struct {
	Person *models.Person
	Err    error
}

type Job struct {
	ctx       context.Context
	sessionId string
	resultCh  chan Result
}

type BackgroundWorker interface {
	Start()
	Stop()
	Process(ctx context.Context, sessionId string) <-chan Result
	WithConfig(cfg config.WorkerConfig) Worker
}

type Worker struct {
	ctx         context.Context
	provider    Provider
	queue       chan Job
	concurrency int
	wg          sync.WaitGroup
}

func NewWorker(ctx context.Context, provider Provider) *Worker {
	return &Worker{
		ctx:         ctx,
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

func (w *Worker) Start() {
	for i := 0; i < w.concurrency; i++ {
		w.wg.Add(1)
		go w.perform()
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

func (w *Worker) perform() {
	defer w.wg.Done()

	for {
		select {
		case j, ok := <-w.queue:
			if !ok {
				return
			}

			person, err := w.provider.FetchSession(w.ctx, j.sessionId)
			j.resultCh <- Result{Person: person, Err: err}
			close(j.resultCh)
		case <-w.ctx.Done():
			return
		}
	}
}
