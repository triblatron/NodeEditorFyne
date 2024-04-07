package nodeeditor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNodeModelHasEmptyPorts(t *testing.T) {
	sut := NewNodeModel(0, "test", CAT_NONE)
	assert.Equal(t, 0, len(sut.Ports))
}

func TestNewNodeModelCannotAddNilPort(t *testing.T) {
	sut := NewNodeModel(0, "test", CAT_NONE)
	sut.AddPort(nil)
	assert.Equal(t, 0, len(sut.Ports))
}

func TestNewNodeModelCanAddValidPort(t *testing.T) {
	sut := NewNodeModel(0, "test", CAT_NONE)
	port := NewPortModel(0, "test", TYPE_DOUBLE, DIR_OUT)
	sut.AddPort(port)
	assert.Equal(t, 1, len(sut.Ports))
}
