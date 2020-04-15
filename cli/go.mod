module cli

go 1.13

require (
	github.com/Netzlink/pangolin/pangolin-operator v0.0.0-20200415123807-07fc34a637f7
	k8s.io/apimachinery v0.0.0
	k8s.io/client-go v12.0.0+incompatible
)

replace (
	github.com/docker/docker => github.com/docker/docker v1.12.7-0.20170113203305-175f18293774
	github.com/openshift/api => github.com/openshift/api v0.0.0-20190923092516-169848dd8137
	k8s.io/api => k8s.io/kubernetes/staging/src/k8s.io/api v0.0.0-20190623232353-8c3b7d7679cc
	k8s.io/apiextensions-apiserver => k8s.io/kubernetes/staging/src/k8s.io/apiextensions-apiserver v0.0.0-20190623232353-8c3b7d7679cc
	k8s.io/apimachinery => k8s.io/kubernetes/staging/src/k8s.io/apimachinery v0.0.0-20190623232353-8c3b7d7679cc
	k8s.io/apiserver => k8s.io/kubernetes/staging/src/k8s.io/apiserver v0.0.0-20190623232353-8c3b7d7679cc
	k8s.io/cli-runtime => k8s.io/kubernetes/staging/src/k8s.io/cli-runtime v0.0.0-20190623232353-8c3b7d7679cc
	k8s.io/client-go => k8s.io/kubernetes/staging/src/k8s.io/client-go v0.0.0-20190623232353-8c3b7d7679cc
	k8s.io/cloud-provider => k8s.io/kubernetes/staging/src/k8s.io/cloud-provider v0.0.0-20190623232353-8c3b7d7679cc
	k8s.io/cluster-bootstrap => k8s.io/kubernetes/staging/src/k8s.io/cluster-bootstrap v0.0.0-20190623232353-8c3b7d7679cc
	k8s.io/code-generator => k8s.io/kubernetes/staging/src/k8s.io/code-generator v0.0.0-20190623232353-8c3b7d7679cc
	k8s.io/component-base => k8s.io/kubernetes/staging/src/k8s.io/component-base v0.0.0-20190623232353-8c3b7d7679cc
	k8s.io/cri-api => k8s.io/kubernetes/staging/src/k8s.io/cri-api v0.0.0-20190623232353-8c3b7d7679cc
	k8s.io/csi-translation-lib => k8s.io/kubernetes/staging/src/k8s.io/csi-translation-lib v0.0.0-20190623232353-8c3b7d7679cc
	k8s.io/kube-aggregator => k8s.io/kubernetes/staging/src/k8s.io/kube-aggregator v0.0.0-20190623232353-8c3b7d7679cc
	k8s.io/kube-controller-manager => k8s.io/kubernetes/staging/src/k8s.io/kube-controller-manager v0.0.0-20190623232353-8c3b7d7679cc
	k8s.io/kube-proxy => k8s.io/kubernetes/staging/src/k8s.io/kube-proxy v0.0.0-20190623232353-8c3b7d7679cc
	k8s.io/kube-scheduler => k8s.io/kubernetes/staging/src/k8s.io/kube-scheduler v0.0.0-20190623232353-8c3b7d7679cc
	k8s.io/kubectl => k8s.io/kubernetes/staging/src/k8s.io/kubectl v0.0.0-20190623232353-8c3b7d7679cc
	k8s.io/kubelet => k8s.io/kubernetes/staging/src/k8s.io/kubelet v0.0.0-20190623232353-8c3b7d7679cc
	k8s.io/legacy-cloud-providers => k8s.io/kubernetes/staging/src/k8s.io/legacy-cloud-providers v0.0.0-20190623232353-8c3b7d7679cc
	k8s.io/metrics => k8s.io/kubernetes/staging/src/k8s.io/metrics v0.0.0-20190623232353-8c3b7d7679cc
	k8s.io/node-api => k8s.io/kubernetes/staging/src/k8s.io/node-api v0.0.0-20190623232353-8c3b7d7679cc
	k8s.io/sample-apiserver => k8s.io/kubernetes/staging/src/k8s.io/sample-apiserver v0.0.0-20190623232353-8c3b7d7679cc
	k8s.io/sample-cli-plugin => k8s.io/kubernetes/staging/src/k8s.io/sample-cli-plugin v0.0.0-20190623232353-8c3b7d7679cc
	k8s.io/sample-controller => k8s.io/kubernetes/staging/src/k8s.io/sample-controller v0.0.0-20190623232353-8c3b7d7679cc
)
