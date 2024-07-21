package generate

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/hoomy-official/go-sonos/generator"
	"github.com/hoomy-official/go-sonos/specs"
	"github.com/vanyda-official/go-shared/pkg/cmd"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type Generate struct {
	OutputDirectory string `default:"./api" help:"output directory"`
	SpecGlob        string `default:"./specs/sonos-*.json" help:"spec glob"`
	DocFilename     string `default:"./specs/documentation.json" help:"root api folder glob"`

	CleanAll bool `default:"false" help:"Clean root api content"`
	MkdirAll bool `default:"true" help:"create necessary directories"`
	Stdout   bool `default:"false" help:"write output in stdout instead of file"`
}

func (g Generate) Run(common *cmd.Commons) error {
	logger, err := common.Logger()
	if err != nil {
		return err
	}

	abs, err := filepath.Abs(g.SpecGlob)
	if err != nil {
		logger.Error("cannot get specs glob absolute path", zap.Error(err), zap.String("SpecGlob", g.SpecGlob))
		return err
	}

	specsFilenames, err := filepath.Glob(abs)
	if err != nil {
		logger.Error("cannot glob specs", zap.Error(err), zap.String("specGlob", abs))
		return err
	}

	if g.CleanAll && !g.Stdout {
		err = os.RemoveAll(g.OutputDirectory)
		if err != nil {
			logger.Error("cannot clean output directory", zap.Error(err), zap.String("outputDirectory", g.OutputDirectory))
			return err
		}
	}

	docFilename, err := filepath.Abs(g.DocFilename)
	if err != nil {
		logger.Error("cannot get doc absolute path", zap.Error(err), zap.String("DocFilename", g.DocFilename))
		return err
	}

	docFile, err := os.ReadFile(docFilename)
	if err != nil {
		logger.Error("cannot read doc", zap.Error(err), zap.String("DocFilename", g.DocFilename))
		return err
	}

	var doc specs.Documentation
	err = json.Unmarshal(docFile, &doc)
	if err != nil {
		logger.Error("cannot unmarshal doc", zap.Error(err), zap.String("DocFilename", g.DocFilename))
		return err
	}

	grp := errgroup.Group{}

	for _, s := range specsFilenames {
		var spec specs.APISpec

		specFile, er := os.ReadFile(s)
		if er != nil {
			logger.Error("cannot read spec", zap.Error(er), zap.String("SpecGlob", s))
			return er
		}

		er = json.Unmarshal(specFile, &spec)
		if er != nil {
			logger.Error("cannot unmarshal spec", zap.Error(er), zap.String("SpecGlob", s))
			return er
		}

		version := fmt.Sprintf("v%d", spec.SoftwareGeneration)
		fileWriter := NewFileWriter(g.OutputDirectory, version, spec.Model, g.Stdout, g.MkdirAll)

		grp.Go(generateTypes(fileWriter, spec))
		grp.Go(generateAllSpecs(fileWriter, spec, doc))
	}

	return grp.Wait()
}

type FileWriter struct {
	path     string
	stdout   bool
	mkdirAll bool
}

func NewFileWriter(root string, generation string, model string, stdout bool, mkdirAll bool) *FileWriter {
	return &FileWriter{
		path:     filepath.Join(root, generation, model),
		stdout:   stdout,
		mkdirAll: mkdirAll,
	}
}

func (f FileWriter) BuildWriter(filename string) (io.Writer, error) {
	if f.stdout {
		return os.Stdout, nil
	}

	if f.mkdirAll {
		if err := os.MkdirAll(f.path, 0777); err != nil {
			return nil, err
		}
	}

	return os.Create(filepath.Join(f.path, filename))
}

func generateTypes(f *FileWriter, spec specs.APISpec) func() error {
	return func() error {
		typeWriter, err := f.BuildWriter("types.go")
		if err != nil {
			return err
		}

		err = generator.GenerateTypes(typeWriter, spec)
		if err != nil {
			return err
		}

		return nil
	}
}

func generateAllSpecs(f *FileWriter, spec specs.APISpec, doc specs.Documentation) func() error {
	return func() error {
		servicesWriter, err := f.BuildWriter(fmt.Sprintf("%s.go", spec.Model))
		if err != nil {
			return err
		}

		err = generator.GenerateServices(servicesWriter, spec, doc)
		if err != nil {
			return err
		}

		for _, service := range spec.Services {
			serviceWriter, er := f.BuildWriter(fmt.Sprintf("%s.go", lowerCamelCase(service.Name)))
			if er != nil {
				return er
			}

			if er = generator.GenerateService(serviceWriter, spec, service, doc); er != nil {
				return er
			}
		}

		return nil
	}
}

func lowerCamelCase(s string) string {
	// let's assume that is there is multiple upper case it is intentional
	if !(len(s) >= 2 && strings.ToUpper(s[:2]) == s[:2]) {
		s = strings.ToLower(s[:1]) + s[1:]
	}

	if s == "type" {
		return "t"
	}

	return s
}
