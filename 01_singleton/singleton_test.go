package singleton_test

import (
	singleton "github.com/junior/go-design-pattern/01_singleton"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInstance(t *testing.T) {
	assert.True(t, singleton.GetInstance() == singleton.GetInstance())
	//assert.False(t, singleton.GetInstance() == singleton.GetLazyInstance())
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singleton.GetInstance() != singleton.GetInstance() {
				b.Errorf("test fail")
			}
		}
	})
}
