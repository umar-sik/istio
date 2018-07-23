//  Copyright 2018 Istio Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package driver

import (
	"testing"

	"istio.io/istio/pkg/test/dependency"
	"istio.io/istio/pkg/test/errors"
)

const (
	// EnvKube is the kubernetes environment
	EnvKube = "kube"

	// EnvLocal is the local environment
	EnvLocal = "local"
)

// Args is the set of arguments to the test driver.
type Args struct {
	// Environment to run the tests in. By default, a local environment will be used.
	Environment string

	// Path to kube config file. Required if the environment is kubernetes.
	KubeConfig string

	// TestID is the id of the test suite. This should supplied by the author once, and must be immutable.
	TestID string

	// Do not cleanup the resources after the test run.
	NoCleanup bool

	// M is testing.M
	M *testing.M

	// SuiteDependencies is the set of dependencies the suite needs.
	SuiteDependencies []dependency.Instance

	// Local working directory root for creating temporary directories / files in. If left empty,
	// os.TempDir() will be used.
	WorkDir string

	// Hub environment variable
	Hub string

	// Tag environment variable
	Tag string
}

// DefaultArgs returns the default set of arguments.
func DefaultArgs() *Args {
	return &Args{
		Environment: EnvLocal,
	}
}

// Validate the arguments.
func (a *Args) Validate() error {
	switch a.Environment {
	case EnvLocal, EnvKube:

	default:
		return errors.UnrecognizedEnvironment(a.Environment)
	}

	if a.Environment == EnvKube && a.KubeConfig == "" {
		return errors.MissingKubeConfigForEnvironment(EnvKube)
	}

	if a.TestID == "" || len(a.TestID) > maxTestIDLength {
		return errors.InvalidTestID(maxTestIDLength)
	}

	return nil
}
