package axtransform

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
 __    _           ___
|  |  |_|_____ ___|_  |
|  |__| |     | .'|  _|
|_____|_|_|_|_|__,|___|
zed (21.03.2024)
*/

type testStructA struct {
	A int
}

type testStructB struct {
	B int
}

func TestAxTransform_Transform(t *testing.T) {
	transform := NewAxTransform[*testStructA, *testStructB]().
		WithMiddlewares(
			func(c *TransformContext[*testStructA, *testStructB]) {
				c.Next()
				t.Logf("From: %d To: %d", c.From.A, c.To.B)
			},
			func(c *TransformContext[*testStructA, *testStructB]) {
				c.To = &testStructB{}
				c.Next()
			},
			func(c *TransformContext[*testStructA, *testStructB]) {
				assert.NotNil(t, c.From)
				assert.NotNil(t, c.To)
				c.To.B = c.From.A * 2
				c.Next()
			},
		).
		Build()
	a := &testStructA{A: 100}
	b, err := transform.Transform(a)
	assert.Nil(t, err)
	assert.Equal(t, b.B, 200)
}
