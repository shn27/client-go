package remedy

import "k8s.io/client-go/kubernetes"

func AlertInstanceDown(podName string, namespace string, clientset *kubernetes.Clientset) {
	DescribePodConditon(podName, namespace, clientset)
}
