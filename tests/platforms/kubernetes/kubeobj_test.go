// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package kubernetes

import (
	"testing"

	"github.com/dapr/dapr/tests/utils"
	"github.com/stretchr/testify/assert"
	apiv1 "k8s.io/api/core/v1"
)

func TestBuildDeploymentObject(t *testing.T) {
	testApp := utils.AppDescription{
		AppName:        "testapp",
		DaprEnabled:    true,
		ImageName:      "helloworld",
		RegistryName:   "dariotest",
		Replicas:       1,
		IngressEnabled: true,
	}

	t.Run("Dapr Enabled", func(t *testing.T) {
		testApp.DaprEnabled = true

		// act
		obj := buildDeploymentObject("testNamespace", testApp)

		// assert
		assert.NotNil(t, obj)
		assert.Equal(t, "true", obj.Spec.Template.Annotations["dapr.io/enabled"])
	})

	t.Run("Dapr disabled", func(t *testing.T) {
		testApp.DaprEnabled = false

		// act
		obj := buildDeploymentObject("testNamespace", testApp)

		// assert
		assert.NotNil(t, obj)
		assert.Empty(t, obj.Spec.Template.Annotations)
	})
}

func TestBuildServiceObject(t *testing.T) {
	testApp := utils.AppDescription{
		AppName:        "testapp",
		DaprEnabled:    true,
		ImageName:      "helloworld",
		RegistryName:   "dariotest",
		Replicas:       1,
		IngressEnabled: true,
	}

	t.Run("Ingress is enabled", func(t *testing.T) {
		testApp.IngressEnabled = true

		// act
		obj := buildServiceObject("testNamespace", testApp)

		// assert
		assert.NotNil(t, obj)
		assert.Equal(t, apiv1.ServiceTypeLoadBalancer, obj.Spec.Type)
	})

	t.Run("Ingress is disabled", func(t *testing.T) {
		testApp.IngressEnabled = false

		// act
		obj := buildServiceObject("testNamespace", testApp)

		// assert
		assert.NotNil(t, obj)
		assert.Equal(t, apiv1.ServiceTypeClusterIP, obj.Spec.Type)
	})
}
