package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	assert.Equal(t, q.Len(), 0)
	assert.Nil(t, q.Dequeue())
	assert.Nil(t, q.Dequeue())

	q.Enqueue("hola")
	assert.Equal(t, q.Len(), 1)
	assert.Equal(t, q.Dequeue(), "hola")

	for i := 0; i < 10; i++ {
		quantity := 1000
		offset := i * quantity
		for nu := offset; nu < quantity+offset; nu++ {
			q.Enqueue(nu)
		}
		assert.Equal(t, q.Len(), quantity)

		for nu := offset; nu < quantity+offset; nu++ {
			assert.Equal(t, q.Dequeue(), nu)
		}
		assert.Nil(t, q.Dequeue())
		assert.Equal(t, q.Len(), 0)
	}

}
