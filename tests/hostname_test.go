package tests

import (
	"github.com/premsvmm/Webinar/controllers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGet(t *testing.T) {
	res := controllers.Hostname
	assert.NotNil(t,res)
}
