package smartid

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_NewWorker(t *testing.T) {
	c := NewClient()

	type result struct {
		concurrency int
		queueSize   int
	}

	tests := []struct {
		name     string
		before   func(w Worker)
		expected result
	}{
		{
			name: "Success",
			before: func(w Worker) {
				w.WithConcurrency(3).WithQueueSize(15)
			},
			expected: result{
				concurrency: 3,
				queueSize:   15,
			},
		},
		{
			name:   "Default values",
			before: func(w Worker) {},
			expected: result{
				concurrency: DefaultConcurrency,
				queueSize:   DefaultQueueSize,
			},
		},
		{
			name: "Zero values",
			before: func(w Worker) {
				w.WithConcurrency(0).WithQueueSize(0)
			},
			expected: result{
				concurrency: DefaultConcurrency,
				queueSize:   DefaultQueueSize,
			},
		},
		{
			name: "Without concurrency option",
			before: func(w Worker) {
				w.WithQueueSize(500)
			},
			expected: result{
				concurrency: DefaultConcurrency,
				queueSize:   500,
			},
		},
		{
			name: "Without queue size option",
			before: func(w Worker) {
				w.WithConcurrency(25)
			},
			expected: result{
				concurrency: 25,
				queueSize:   DefaultQueueSize,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewWorker(c)

			tt.before(w)

			workerImpl := w.(*worker)

			assert.Equal(t, tt.expected.concurrency, workerImpl.concurrency)
			assert.Equal(t, tt.expected.queueSize, cap(workerImpl.queue))
		})
	}
}

func Test_Worker_WithConcurrency(t *testing.T) {
	c := NewClient()

	type result struct {
		concurrency int
		queueSize   int
	}

	tests := []struct {
		name     string
		before   func(w Worker)
		expected result
	}{
		{
			name: "Success",
			before: func(w Worker) {
				w.WithConcurrency(3)
			},
			expected: result{
				concurrency: 3,
				queueSize:   DefaultQueueSize,
			},
		},
		{
			name:   "Default values",
			before: func(w Worker) {},
			expected: result{
				concurrency: DefaultConcurrency,
				queueSize:   DefaultQueueSize,
			},
		},
		{
			name: "Zero values",
			before: func(w Worker) {
				w.WithConcurrency(0)
			},
			expected: result{
				concurrency: DefaultConcurrency,
				queueSize:   DefaultQueueSize,
			},
		},
		{
			name: "Without concurrency option",
			before: func(w Worker) {
				w.WithQueueSize(500)
			},
			expected: result{
				concurrency: DefaultConcurrency,
				queueSize:   500,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewWorker(c)

			tt.before(w)

			workerImpl := w.(*worker)

			assert.Equal(t, tt.expected.concurrency, workerImpl.concurrency)
			assert.Equal(t, tt.expected.queueSize, cap(workerImpl.queue))
		})
	}
}

func Test_Worker_WithQueueSize(t *testing.T) {
	c := NewClient()

	type result struct {
		concurrency int
		queueSize   int
	}

	tests := []struct {
		name     string
		before   func(w Worker)
		expected result
	}{
		{
			name: "Success",
			before: func(w Worker) {
				w.WithQueueSize(15)
			},
			expected: result{
				concurrency: DefaultConcurrency,
				queueSize:   15,
			},
		},
		{
			name:   "Default values",
			before: func(w Worker) {},
			expected: result{
				concurrency: DefaultConcurrency,
				queueSize:   DefaultQueueSize,
			},
		},
		{
			name: "Zero values",
			before: func(w Worker) {
				w.WithQueueSize(0)
			},
			expected: result{
				concurrency: DefaultConcurrency,
				queueSize:   DefaultQueueSize,
			},
		},
		{
			name: "Without queue size option",
			before: func(w Worker) {
				w.WithConcurrency(500)
			},
			expected: result{
				concurrency: 500,
				queueSize:   DefaultQueueSize,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewWorker(c)

			tt.before(w)

			workerImpl := w.(*worker)

			assert.Equal(t, tt.expected.concurrency, workerImpl.concurrency)
			assert.Equal(t, tt.expected.queueSize, cap(workerImpl.queue))
		})
	}
}

func Test_Worker_Start(t *testing.T) {
	ctx := context.Background()
	c := NewClient()
	w := NewWorker(c).WithConcurrency(3).WithQueueSize(15)

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

			assert.NotNil(t, w)
		})
	}
}

func Test_Worker_Stop(t *testing.T) {
	ctx := context.Background()
	c := NewClient()
	w := NewWorker(c).WithConcurrency(3).WithQueueSize(15)

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

			assert.NoError(t, nil)
		})
	}
}

func Test_Worker_Process(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockClient := NewMockClient(ctrl)
	w := NewWorker(mockClient).WithConcurrency(3).WithQueueSize(15)
	w.Start(ctx)
	defer w.Stop()

	tests := []struct {
		name      string
		sessionId string
		before    func()
		expect    *Person
		err       error
	}{
		{
			name:      "Success",
			sessionId: "c2731f5e-9d63-4db7-b83c-db528d2f7021",
			before: func() {
				mockClient.EXPECT().
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
			err: nil,
		},
		{
			name:      "Error: USER_REFUSED",
			sessionId: "c2731f5e-9d63-4db7-b83c-db528d2f7021",
			before: func() {
				mockClient.EXPECT().
					FetchSession(ctx, "c2731f5e-9d63-4db7-b83c-db528d2f7021").
					Return(nil, &Error{Code: "USER_REFUSED"})
			},
			expect: nil,
			err:    &Error{Code: "USER_REFUSED"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.before()

			resultCh := w.Process(ctx, tt.sessionId)
			assert.NotNil(t, resultCh)

			result := <-resultCh
			if tt.err != nil {
				assert.Equal(t, tt.err, result.Err)
			} else {
				assert.NoError(t, result.Err)
				assert.Equal(t, tt.expect, result.Person)
			}
		})
	}
}
