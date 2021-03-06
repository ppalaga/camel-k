/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package context

import (
	"github.com/apache/camel-k/pkg/apis/camel/v1alpha1"
	"github.com/apache/camel-k/pkg/platform"
	"github.com/apache/camel-k/pkg/util/digest"
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/sirupsen/logrus"
)

// NewInitializeAction creates a new initialization handling action for the context
func NewInitializeAction() Action {
	return &initializeAction{}
}

type initializeAction struct {
}

func (action *initializeAction) Name() string {
	return "initialize"
}

func (action *initializeAction) CanHandle(context *v1alpha1.IntegrationContext) bool {
	return context.Status.Phase == ""
}

func (action *initializeAction) Handle(context *v1alpha1.IntegrationContext) error {
	// The integration platform needs to be initialized before starting to create contexts
	if _, err := platform.GetCurrentPlatform(context.Namespace); err != nil {
		logrus.Info("Waiting for a integration platform to be initialized")
		return nil
	}

	target := context.DeepCopy()

	// execute custom initialization
	//if err := trait.apply(nil, context); err != nil {
	//	return err
	//}

	// update the status
	logrus.Info("Context ", target.Name, " transitioning to state ", v1alpha1.IntegrationContextPhaseBuilding)
	target.Status.Phase = v1alpha1.IntegrationContextPhaseBuilding
	dgst, err := digest.ComputeForIntegrationContext(context)
	if err != nil {
		return err
	}
	target.Status.Digest = dgst

	return sdk.Update(target)
}
