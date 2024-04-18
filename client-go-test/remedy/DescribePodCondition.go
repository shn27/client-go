package remedy

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func DescribePodConditon(podName string, namespace string, clientset *kubernetes.Clientset) {
	pod, err := clientset.CoreV1().Pods(namespace).Get(context.Background(), podName, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
	conditions := pod.Status.Conditions
	if err != nil {
		panic(err.Error())
	}
	for _, condition := range conditions {
		fmt.Printf("Type: %s, Status: %s, LastTransitionTime: %s, Reason: %s, Message: %s\n",
			condition.Type, condition.Status, condition.LastTransitionTime, condition.Reason, condition.Message)
	}
}
