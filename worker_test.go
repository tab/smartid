package smartid

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/tab/smartid/internal/config"
)

func Test_NewWorker(t *testing.T) {
	client := NewClient()

	w := NewWorker(client)

	tests := []struct {
		name     string
		before   func()
		expected *Worker
	}{
		{
			name: "Success",
			before: func() {
				w.WithConfig(config.WorkerConfig{
					Concurrency: 3,
					QueueSize:   15,
				})
			},
			expected: &Worker{
				provider:    client,
				queue:       make(chan Job, 15),
				concurrency: 3,
			},
		},
		{
			name: "Default values",
			before: func() {
				w.WithConfig(config.WorkerConfig{})
			},
			expected: &Worker{
				provider:    client,
				queue:       make(chan Job, 100),
				concurrency: 10,
			},
		},
		{
			name: "Zero values",
			before: func() {
				w.WithConfig(config.WorkerConfig{
					Concurrency: 0,
					QueueSize:   0,
				})
			},
			expected: &Worker{
				provider:    client,
				queue:       make(chan Job, 100),
				concurrency: 10,
			},
		},
		{
			name: "Without concurrency option",
			before: func() {
				w.WithConfig(config.WorkerConfig{
					QueueSize: 500,
				})
			},
			expected: &Worker{
				provider:    client,
				queue:       make(chan Job, 500),
				concurrency: 10,
			},
		},
		{
			name: "Without queue size option",
			before: func() {
				w.WithConfig(config.WorkerConfig{
					Concurrency: 25,
				})
			},
			expected: &Worker{
				provider:    client,
				queue:       make(chan Job, 100),
				concurrency: 25,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before()

			assert.Equal(t, tt.expected.concurrency, w.concurrency)
			assert.Equal(t, cap(tt.expected.queue), cap(w.queue))
		})
	}
}

func Test_Worker_Start(t *testing.T) {
	ctx := context.Background()
	client := NewClient()
	worker := NewWorker(client).WithConfig(config.WorkerConfig{
		Concurrency: 3,
		QueueSize:   15,
	})

	tests := []struct {
		name string
	}{
		{
			name: "Success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			worker.Start(ctx)
			worker.Stop()

			assert.NotNil(t, worker)
		})
	}
}

func Test_Worker_Stop(t *testing.T) {
	ctx := context.Background()
	client := NewClient()
	w := NewWorker(client).WithConfig(config.WorkerConfig{
		Concurrency: 3,
		QueueSize:   15,
	})

	tests := []struct {
		name string
	}{
		{
			name: "Success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w.Start(ctx)
			w.Stop()
			assert.Nil(t, w.queue)
		})
	}
}

func Test_Worker_Process(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	client := NewMockProvider(ctrl)
	w := NewWorker(client).WithConfig(config.WorkerConfig{
		Concurrency: 3,
		QueueSize:   15,
	})

	w.Start(ctx)
	defer w.Stop()

	tests := []struct {
		name      string
		sessionId string
		before    func()
		expect    *Person
		error     error
	}{
		{
			name:      "Success",
			sessionId: "c2731f5e-9d63-4db7-b83c-db528d2f7021",
			before: func() {
				client.EXPECT().
					FetchSession(ctx, "c2731f5e-9d63-4db7-b83c-db528d2f7021").
					Return(&Person{
						IdentityNumber: "PNOEE-30303039914",
						PersonalCode:   "30303039914",
						FirstName:      "TESTNUMBER",
						LastName:       "OK",
					}, nil)
			},
			expect: &Person{
				IdentityNumber: "PNOEE-30303039914",
				PersonalCode:   "30303039914",
				FirstName:      "TESTNUMBER",
				LastName:       "OK",
			},
			error: nil,
		},
		{
			name:      "Error: USER_REFUSED",
			sessionId: "c2731f5e-9d63-4db7-b83c-db528d2f7021",
			before: func() {
				client.EXPECT().
					FetchSession(ctx, "c2731f5e-9d63-4db7-b83c-db528d2f7021").
					Return(nil, &Error{Code: "USER_REFUSED"})
			},
			expect: nil,
			error:  &Error{Code: "USER_REFUSED"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before()

			resultCh := w.Process(ctx, tt.sessionId)
			assert.NotNil(t, resultCh)

			result := <-resultCh

			if tt.error != nil {
				assert.Equal(t, tt.error, result.Err)
			} else {
				assert.NoError(t, result.Err)
				assert.Equal(t, tt.expect, result.Person)
			}
		})
	}
}
