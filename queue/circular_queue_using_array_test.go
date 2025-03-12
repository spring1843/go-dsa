package queue

import "testing"

/*
TestCircularQueue tests solution(s) with the following signature and problem description:

	func (queue *CircularQueue) enqueue(n int) error
	func (queue *CircularQueue) dequeue() (int, error)

Given a size, implement a circular queue using a slice.

A circular queue also called a ring buffer is different from a normal queue in that
the last element is connected to the first element.

For example given a circular array of size 4 if we enqueue integers {1,2,3,4,5,6} and then dequeue 2 times
we should get 3 and 4.

	        , - ~ - ,                 , - ~ - ,                 , - ~ - ,                 , - ~ - ,
	    , '    ...    ' ,         , '    ...    ' ,         , '    ...    ' ,         , '    ...    ' ,
	  ,        ...        ,     ,        ...        ,     ,        ...        ,     ,        ...        ,
	 ,    1    ...         ,   ,    1    ...   2     ,   ,    1    ...   2     ,   ,    1    ...   2     ,
	,          ...          , ,          ...          , ,          ...          , ,          ...          ,
	,  ...................  , ,  ...................  , ,  ...................  , ,  ...................  ,
	,  ...................  , ,  ...................  , ,  ...................  , ,  ...................  ,
	 ,         ...         ,   ,         ...         ,   ,         ...         ,   ,         ...         ,
	  ,        ...        ,     ,        ...        ,     ,        ...   3    ,     ,   4    ...   3    ,
	    ,      ...     , '        ,      ...     , '        ,      ...     , '        ,      ...     , '
	      ' - ,...,  '              ' - ,...,  '              ' - ,...,  '              ' - ,...,  '

	        , - ~ - ,                 , - ~ - ,                 , - ~ - ,                 , - ~ - ,
	    , '    ...    ' ,         , '    ...    ' ,         , '    ...    ' ,         , '    ...    ' ,
	  ,        ...        ,     ,        ...        ,     ,        ...        ,     ,        ...        ,
	 ,     5   ...   2     ,   ,     5   ...   6     ,   ,     5   ...   6     ,   ,    5    ...    6    ,
	,          ...          , ,          ...          , ,          ...          , ,          ...          ,
	,  ...................  , ,  ...................  , ,  ...................  , ,  ...................  ,
	,  ...................  , ,  ...................  , ,  ...................  , ,  ...................  ,
	 ,         ...         ,   ,         ...         ,   ,         ...         ,   ,         ...         ,
	  ,    4   ...   3    ,     ,    4   ...   3    ,     ,    4   ...        ,     ,        ...        ,
	    ,      ...     , '        ,      ...     , '        ,      ...     , '        ,      ...     , '
	      ' - ,...,  '              ' - ,...,  '              ' - ,...,  '              ' - ,...,  '

As shown in the above example the circular queue does not run out of capacity and still allows FIFO operations.
*/

func TestCircularQueue(t *testing.T) {

	type testRound struct {
		enqueueStart int
		enqueueEnd   int
		dequeueStart int
		dequeueEnd   int
		expectedErr  error
	}

	tests := []struct {
		size       int
		testRounds []testRound
	}{
		{1, []testRound{{1, 6, 6, 6, nil}}},
		{2, []testRound{{1, 6, 5, 5, nil}}},
		{4, []testRound{{1, 6, 3, 6, nil}}},
		{4, []testRound{{1, 6, 3, 6, nil}, {1, 6, 3, 6, nil}}},
		{4, []testRound{{1, 6, 3, 7, ErrQueueEmpty}}},
	}

	for i, test := range tests {
		queue := NewCircularQueue(test.size)

		for j, testRound := range test.testRounds {
			for i := testRound.enqueueStart; i <= testRound.enqueueEnd; i++ {
				queue.enqueue(i)
			}

			for want := testRound.dequeueStart; want <= testRound.dequeueEnd; want++ {
				got, err := queue.dequeue()
				if err != nil {
					if err != testRound.expectedErr {
						t.Fatalf("Failed test case #%d round #%d. Unexpected error %s", i, j, err)
					}
					break
				}

				if got != want {
					t.Fatalf("Failed test case #%d round #%d. Want %d, got %d", i, j, want, got)
				}
			}
		}
	}
}
