package smartid

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/tab/smartid/internal/config"
	"github.com/tab/smartid/internal/models"
)

func Test_NewWorker(t *testing.T) {
	ctx := context.Background()
	client := NewClient()

	worker := NewWorker(ctx, client)

	tests := []struct {
		name     string
		before   func()
		expected *Worker
	}{
		{
			name: "Success",
			before: func() {
				worker.WithConfig(config.WorkerConfig{
					Concurrency: 3,
					QueueSize:   15,
				})
			},
			expected: &Worker{
				ctx:         ctx,
				provider:    client,
				queue:       make(chan Job, 15),
				concurrency: 3,
			},
		},
		{
			name: "Default values",
			before: func() {
				worker.WithConfig(config.WorkerConfig{})
			},
			expected: &Worker{
				ctx:         ctx,
				provider:    client,
				queue:       make(chan Job, 100),
				concurrency: 10,
			},
		},
		{
			name: "Zero values",
			before: func() {
				worker.WithConfig(config.WorkerConfig{
					Concurrency: 0,
					QueueSize:   0,
				})
			},
			expected: &Worker{
				ctx:         ctx,
				provider:    client,
				queue:       make(chan Job, 100),
				concurrency: 10,
			},
		},
		{
			name: "Without concurrency option",
			before: func() {
				worker.WithConfig(config.WorkerConfig{
					QueueSize: 500,
				})
			},
			expected: &Worker{
				ctx:         ctx,
				provider:    client,
				queue:       make(chan Job, 500),
				concurrency: 10,
			},
		},
		{
			name: "Without queue size option",
			before: func() {
				worker.WithConfig(config.WorkerConfig{
					Concurrency: 25,
				})
			},
			expected: &Worker{
				ctx:         ctx,
				provider:    client,
				queue:       make(chan Job, 100),
				concurrency: 25,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before()

			assert.Equal(t, tt.expected.concurrency, worker.concurrency)
			assert.Equal(t, cap(tt.expected.queue), cap(worker.queue))
		})
	}
}

func Test_Worker_Start(t *testing.T) {
	client := NewClient()
	worker := NewWorker(context.Background(), client).WithConfig(config.WorkerConfig{
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
			worker.Start()
			worker.Stop()

			assert.NotNil(t, worker)
		})
	}
}

func Test_Worker_Stop(t *testing.T) {
	client := NewClient()
	worker := NewWorker(context.Background(), client).WithConfig(config.WorkerConfig{
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
			worker.Start()
			worker.Stop()
			assert.Nil(t, worker.queue)
		})
	}
}

func Test_Worker_Process(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	client := NewMockProvider(ctrl)
	worker := NewWorker(ctx, client).WithConfig(config.WorkerConfig{
		Concurrency: 3,
		QueueSize:   15,
	})

	worker.Start()
	defer worker.Stop()

	tests := []struct {
		name      string
		sessionId string
		before    func()
		expect    *models.Person
		error     error
	}{
		{
			name:      "Success",
			sessionId: "c2731f5e-9d63-4db7-b83c-db528d2f7021",
			before: func() {
				client.EXPECT().
					FetchSession(ctx, "c2731f5e-9d63-4db7-b83c-db528d2f7021").
					Return(&models.Person{
						IdentityNumber: "PNOEE-30303039914",
						PersonalCode:   "30303039914",
						FirstName:      "TESTNUMBER",
						LastName:       "OK",
					}, nil)
			},
			expect: &models.Person{
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

			resultCh := worker.Process(tt.sessionId)
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
