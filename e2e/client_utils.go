// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package e2e

import (
	"os"
	"path"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/client-go/discovery/cached/memory"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func clientConfig() (*rest.Config, error) {
	return clientcmd.BuildConfigFromFlags("", path.Join(os.Getenv("HOME"), ".kube/config"))
}

func newClientSet(config *rest.Config) (*kubernetes.Clientset, error) {
	return kubernetes.NewForConfig(config)
}

func newRESTMapper(clientSet *kubernetes.Clientset) (meta.ResettableRESTMapper, error) {
	discoveryClient := memory.NewMemCacheClient(clientSet.Discovery())
	mapper := restmapper.NewDeferredDiscoveryRESTMapper(discoveryClient)
	return mapper, nil
}

func newKubeClient(config *rest.Config, mapper meta.RESTMapper) (client.Client, error) {
	return client.New(config, client.Options{
		Mapper: mapper,
	})
}
