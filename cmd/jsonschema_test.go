package cmd_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/linuxsuren/api-testing/cmd"
	"github.com/linuxsuren/api-testing/pkg/server"
	fakeruntime "github.com/linuxsuren/go-fake-runtime"
	"github.com/stretchr/testify/assert"
)

func TestJSONSchemaCmd(t *testing.T) {
	c := cmd.NewRootCmd(fakeruntime.FakeExecer{ExpectOS: "linux"},
		cmd.NewFakeGRPCServer(), server.NewFakeHTTPServer())

	buf := new(bytes.Buffer)
	c.SetOut(buf)

	c.SetArgs([]string{"json"})
	err := c.Execute()
	assert.Nil(t, err)
	assert.True(t, strings.Contains(buf.String(), "schema"))
}
