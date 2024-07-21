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
	serviceTestData []byte

	//go:embed testData/documentation.json
	serviceDocumentationRaw []byte

	//go:embed testData/alarmClock.go.test
	expectedActions string
)

func TestService(t *testing.T) {
	t.Run("alarmService", func(t *testing.T) {
		var o specs.APISpec
		err := json.Unmarshal(serviceTestData, &o)
		assert.Nil(t, err)

		var d specs.Documentation
		err = json.Unmarshal(serviceDocumentationRaw, &d)
		assert.Nil(t, err)

		f := new(bytes.Buffer)
		err = generator.GenerateService(f, o, o.Services[0], d)
		assert.Nil(t, err)

		assert.Equal(t, f.String(), expectedActions)
	})
}
