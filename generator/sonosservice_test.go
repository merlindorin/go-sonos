package generator_test

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/hoomy-official/go-sonos/generator"
	"github.com/hoomy-official/go-sonos/specs"
	"github.com/stretchr/testify/assert"
)

var (
	//go:embed testData/S6.json
	sonosServiceTestData []byte

	//go:embed testData/documentation.json
	sonosServiceDocumentationRaw []byte

	//go:embed testData/S6.go.test
	expectedService string
)

func TestSonosService(t *testing.T) {
	t.Run("sonosServices", func(t *testing.T) {
		var api specs.APISpec
		err := json.Unmarshal(sonosServiceTestData, &api)
		assert.Nil(t, err)

		var doc specs.Documentation
		err = json.Unmarshal(sonosServiceDocumentationRaw, &doc)
		assert.Nil(t, err)

		f := new(bytes.Buffer)
		err = generator.GenerateServices(f, api, doc)
		assert.Nil(t, err)

		assert.Equal(t, f.String(), expectedService)
	})
}
