package cluster

import (
	"github.com/weaveworks/flux"
	"github.com/weaveworks/flux/policy"
	"github.com/weaveworks/flux/resource"
)

// Doubles as a cluster.Cluster and cluster.Manifests implementation
type Mock struct {
	AllServicesFunc         func(maybeNamespace string) ([]Service, error)
	SomeServicesFunc        func([]flux.ServiceID) ([]Service, error)
	PingFunc                func() error
	ExportFunc              func() ([]byte, error)
	SyncFunc                func(SyncDef) error
	FindDefinedServicesFunc func(path string) (map[flux.ServiceID][]string, error)
	UpdateDefinitionFunc    func(def []byte, newImageID flux.ImageID) ([]byte, error)
	LoadManifestsFunc       func(paths ...string) (map[string]resource.Resource, error)
	ParseManifestsFunc      func([]byte) (map[string]resource.Resource, error)
	UpdateManifestFunc      func(path, resourceID string, f func(def []byte) ([]byte, error)) error
	UpdatePoliciesFunc      func([]byte, policy.Update) ([]byte, error)
	ServicesWithPolicyFunc  func(path string, p policy.Policy) (flux.ServiceIDSet, error)
}

func (m *Mock) AllServices(maybeNamespace string) ([]Service, error) {
	return m.AllServicesFunc(maybeNamespace)
}

func (m *Mock) SomeServices(s []flux.ServiceID) ([]Service, error) {
	return m.SomeServicesFunc(s)
}

func (m *Mock) Ping() error {
	return m.PingFunc()
}

func (m *Mock) Export() ([]byte, error) {
	return m.ExportFunc()
}

func (m *Mock) Sync(c SyncDef) error {
	return m.SyncFunc(c)
}

func (m *Mock) FindDefinedServices(path string) (map[flux.ServiceID][]string, error) {
	return m.FindDefinedServicesFunc(path)
}

func (m *Mock) UpdateDefinition(def []byte, newImageID flux.ImageID) ([]byte, error) {
	return m.UpdateDefinitionFunc(def, newImageID)
}

func (m *Mock) LoadManifests(paths ...string) (map[string]resource.Resource, error) {
	return m.LoadManifestsFunc(paths...)
}

func (m *Mock) ParseManifests(def []byte) (map[string]resource.Resource, error) {
	return m.ParseManifestsFunc(def)
}

func (m *Mock) UpdateManifest(path string, resourceID string, f func(def []byte) ([]byte, error)) error {
	return m.UpdateManifestFunc(path, resourceID, f)
}

func (m *Mock) UpdatePolicies(def []byte, p policy.Update) ([]byte, error) {
	return m.UpdatePoliciesFunc(def, p)
}

func (m *Mock) ServicesWithPolicy(path string, p policy.Policy) (flux.ServiceIDSet, error) {
	return m.ServicesWithPolicyFunc(path, p)
}