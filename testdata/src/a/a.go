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
