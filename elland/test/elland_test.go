package test

import (
  "flag"
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
)

var imgName = flag.String("imgName", "", "image name")

func TestPythonPackagesVersions(t *testing.T) {

  // Build image
	buildOptions := &docker.BuildOptions{
		Tags: []string{*imgName},
    BuildArgs: []string{"-f", "Dockerfile"},
	}

	docker.Build(t, "../", buildOptions)

	// Make sure Python version is: 3.9.12
	pythonOpts := &docker.RunOptions{Command: []string{"python", "--version"}}
	pythonOutput := docker.Run(t, *imgName, pythonOpts)
	assert.Equal(t, "Python 3.9.12", pythonOutput)

	// Make sure PyTorch version is: 1.13.0
  torchOpts := &docker.RunOptions{Command: []string{"python", "-u", "-c", "import torch; print(torch.__version__)"}}
  torchOutput := docker.Run(t, *imgName, torchOpts)
  assert.Equal(t, "1.13.0", torchOutput)

	// Make sure PyTorch Lightning version is: 1.7.7
  lightningOpts := &docker.RunOptions{Command: []string{"python", "-u", "-c", "import pytorch_lightning; print(pytorch_lightning.__version__)"}}
  lightningOutput := docker.Run(t, *imgName, lightningOpts)
  assert.Equal(t, "1.7.7", lightningOutput)
}
