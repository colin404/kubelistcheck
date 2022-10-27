<div align="center">

# kubelistcheck

</div>

---

`kubelistcheck` is a golang analyzer that checks if code get/list kubernetes resources from kube-apiserver cache

### Installation

```shell
go install github.com/colin404/kubelistcheck/cmd/kubelistcheck@latest
```

### Usage

```
kubelistcheck [-flag] [package]

Flags:
  -get
       enable check for GetOptions 
```

### Example

```go
// File: a.go
package a

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
)

func ListPods(cli clientset.Interface) (*v1.PodList, error) {
	return cli.CoreV1().Pods(metav1.NamespaceDefault).List(context.Background(), metav1.ListOptions{})
}

func GetPod(cli clientset.Interface, name string) (*v1.Pod, error) {
	return cli.CoreV1().Pods(metav1.NamespaceDefault).Get(context.Background(), name, metav1.GetOptions{})
}
```

Run `kubelistcheck`:
```bash
kubelistcheck a.go
```

or you can check recursively:

```
kubelistcheck ./...
```

output is:

```
/tmp/a.go:12:79: ResourceVersion is not setted in ListOptions
```
