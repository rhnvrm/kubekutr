package models

// Config represents the structure to hold configuration loaded from an external data source.
type Config struct {
	Deployments  []Deployment  `koanf:"deployments"`
	Services     []Service     `koanf:"services"`
	Ingresses    []Ingress     `koanf:"ingresses"`
	StatefulSets []StatefulSet `koanf:"statefulsets"`
}

// Deployment represents configuration options for the Deployment spec.
type Deployment struct {
	Name       string      `koanf:"name"`
	Replicas   string      `koanf:"replicas"`
	Containers []Container `koanf:"containers"`
	Labels     []Label     `koanf:"labels"`
}

// Container represents configuration options for the Container spec in a Pod definition.
type Container struct {
	Name          string   `koanf:"name"`
	Image         string   `koanf:"image"`
	EnvSecret     string   `koanf:"envSecret"`
	EnvVars       []EnvVar `koanf:"envVars"`
	Container     string   `koanf:"container"`
	PortInt       string   `koanf:"port"`
	Command       string   `koanf:"command"`
	Args          string   `koanf:"arg"`
	ConfigMapName string   `koanf:"configmap"`
}

// Service represents configuration options for Service spec.
type Service struct {
	Name       string `koanf:"name"`
	Port       int    `koanf:"port"`
	TargetPort int    `koanf:"targetPort"`
	Type       string `koanf:"type"`
}

// Ingress represents configuration options for Ingress spec.
type Ingress struct {
	Name        string        `koanf:"name"`
	Class       string        `koanf:"class"`
	Paths       []IngressPath `koanf:"ingressPaths"`
	Annotations []Annotation  `koanf:"annotations"`
}

// IngressPath represents the definition for `paths` specified in Ingress.
type IngressPath struct {
	Path    string `koanf:"path"`
	Service string `koanf:"service"`
	Port    string `koanf:"port"`
}

// StatefulSet represents configuration options for StatefulSet spec.
type StatefulSet struct {
	Name        string        `koanf:"name"`
	Class       string        `koanf:"class"`
	Paths       []IngressPath `koanf:"ingressPaths"`
	Annotations []Annotation  `koanf:"annotations"`
}

// Resource is a set of common actions performed on Resource Types.
type Resource interface {
	GetMetaData() ResourceMeta
}

// ResourceMeta contains metadata for preparing resource manifests.
type ResourceMeta struct {
	Name         string
	Config       map[string]interface{}
	TemplatePath string
	ManifestPath string
}

// Annotation represents the name of annotation value.
type Annotation struct {
	Name string `koanf:"name"`
}

// Label represents the kv pair for a label.
type Label struct {
	Label string `koanf:"label"`
}

// EnvVar represents the env variables to be used in Pod definition.
type EnvVar struct {
	Name  string `koanf:"name"`
	Value string `koanf:"value"`
}
