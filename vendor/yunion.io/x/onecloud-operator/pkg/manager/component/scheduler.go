// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package component

import (
	apps "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	"yunion.io/x/onecloud/pkg/scheduler/options"

	"yunion.io/x/onecloud-operator/pkg/apis/constants"
	"yunion.io/x/onecloud-operator/pkg/apis/onecloud/v1alpha1"
	"yunion.io/x/onecloud-operator/pkg/controller"
	"yunion.io/x/onecloud-operator/pkg/manager"
)

type schedulerManager struct {
	*ComponentManager
}

func newSchedulerManager(man *ComponentManager) manager.Manager {
	return &schedulerManager{man}
}

func (m *schedulerManager) Sync(oc *v1alpha1.OnecloudCluster) error {
	return syncComponent(m, oc, oc.Spec.Scheduler.Disable)
}

func (m *schedulerManager) getPhaseControl(man controller.ComponentManager) controller.PhaseControl {
	return controller.NewRegisterEndpointComponent(man, v1alpha1.SchedulerComponentType,
		constants.ServiceNameScheduler, constants.ServiceTypeScheduler, constants.SchedulerPort, "")
}

func (m *schedulerManager) getConfigMap(oc *v1alpha1.OnecloudCluster, cfg *v1alpha1.OnecloudClusterConfig) (*corev1.ConfigMap, error) {
	opt := options.GetOptions()
	if err := SetOptionsDefault(opt, constants.ServiceTypeScheduler); err != nil {
		return nil, err
	}
	// scheduler use region config directly
	config := cfg.RegionServer
	SetDBOptions(&opt.DBOptions, oc.Spec.Mysql, config.DB)
	SetOptionsServiceTLS(&opt.BaseOptions)
	SetServiceCommonOptions(&opt.CommonOptions, oc, config.ServiceDBCommonOptions.ServiceCommonOptions)

	opt.SchedulerPort = constants.SchedulerPort
	return m.newServiceConfigMap(v1alpha1.SchedulerComponentType, oc, opt), nil
}

func (m *schedulerManager) getService(oc *v1alpha1.OnecloudCluster) *corev1.Service {
	return m.newSingleNodePortService(v1alpha1.SchedulerComponentType, oc, constants.SchedulerPort)
}

func (m *schedulerManager) getDeployment(oc *v1alpha1.OnecloudCluster, cfg *v1alpha1.OnecloudClusterConfig) (*apps.Deployment, error) {
	return m.newCloudServiceSinglePortDeployment(v1alpha1.SchedulerComponentType, oc, oc.Spec.Scheduler, constants.SchedulerPort, false)
}

func (m *schedulerManager) getDeploymentStatus(oc *v1alpha1.OnecloudCluster) *v1alpha1.DeploymentStatus {
	return &oc.Status.Scheduler
}
