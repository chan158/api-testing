/**
MIT License

Copyright (c) 2023 API Testing Authors.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package generator

import (
	"testing"

	_ "embed"

	atest "github.com/linuxsuren/api-testing/pkg/testing"
	"github.com/stretchr/testify/assert"
)

func TestJmeterConvert(t *testing.T) {
	t.Run("common", func(t *testing.T) {
		jmeterConvert := GetTestSuiteConverter("jmeter")
		assert.NotNil(t, jmeterConvert)

		converters := GetTestSuiteConverters()
		assert.Equal(t, 1, len(converters))
	})

	testSuite := &atest.TestSuite{
		Name: "API Testing",
		API:  "http://localhost:8080",
		Items: []atest.TestCase{{
			Name: "hello-jmeter",
			Request: atest.Request{
				Method: "POST",
				API:    "/GetSuites",
			},
		}},
	}

	converter := &jmeterConverter{}
	output, err := converter.Convert(testSuite)
	assert.NoError(t, err)

	assert.Equal(t, expectedJmeter, output, output)
}

//go:embed testdata/expected_jmeter.jmx
var expectedJmeter string
