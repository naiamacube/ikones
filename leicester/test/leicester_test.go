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
	opts := &docker.RunOptions{Command: []string{"vault", "-version"}}
	output := docker.Run(t, *imgName, opts)
	assert.Equal(t, "Vault v1.12.2 (415e1fe3118eebd5df6cb60d13defdc01aa17b03), built 2022-11-23T12:53:46Z", output)
}
