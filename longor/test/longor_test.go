package test

import (
  "flag"
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
)

var imgName = flag.String("imgName", "", "image name")

func TestTerraformVersion(t *testing.T) {

  // Build image
	buildOptions := &docker.BuildOptions{
		Tags: []string{*imgName},
    BuildArgs: []string{"-f", "Dockerfile"},
	}

	docker.Build(t, "../", buildOptions)

	// Make sure Terraform version is 1.3.5
	opts := &docker.RunOptions{Command: []string{"terraform", "-version"}}
	output := docker.Run(t, *imgName, opts)
	assert.Contains(t, output, "Terraform v1.3.5\non linux_amd64")
}
