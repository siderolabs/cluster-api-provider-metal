module github.com/talos-systems/cluster-api-provider-metal

go 1.13

require (
	github.com/go-logr/logr v0.1.0
	github.com/onsi/ginkgo v1.8.0
	github.com/onsi/gomega v1.5.0
	github.com/vmware/goipmi v0.0.0-20181114221114-2333cd82d702
	k8s.io/apimachinery v0.0.0-20190817020851-f2f3a405f61d
	k8s.io/client-go v0.0.0-20190918200256-06eb1244587a
	sigs.k8s.io/cluster-api v0.2.9
	sigs.k8s.io/controller-runtime v0.3.0
)
