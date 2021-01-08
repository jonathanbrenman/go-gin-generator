package factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildPingController(t *testing.T) {
	assert := assert.New(t)
	factoryCtrl := NewCtrlFactory()
	pingCtrl := factoryCtrl.BuildPingController()
	assert.NotNil(pingCtrl)
}