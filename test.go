package main

// import (
// 	"context"
// 	"flag"
// 	"fmt"
// 	"log"
// 	"os"
// 	"path/filepath"
// 	"strings"

// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// 	"k8s.io/client-go/kubernetes"
// 	"k8s.io/client-go/tools/clientcmd"
// )

// func main() {
// 	kubeconfig := flag.String("kubeconfig", filepath.Join(os.Getenv("HOME"), ".kube", "config"), "kubeconfig file")
// 	flag.Parse()

// 	// 初始化 Kubernetes 客户端
// 	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
// 	if err != nil {
// 		log.Fatalf("Failed to build config: %v", err)
// 	}
// 	clientset, err := kubernetes.NewForConfig(config)
// 	if err != nil {
// 		log.Fatalf("Failed to create client: %v", err)
// 	}

// 	// 定义 Pod 的模糊匹配关键词
// 	podNameKeyword := "kubeflow-dashboard" // 修改为你的模糊匹配关键词

// 	// 获取所有的 Pod 列表
// 	pods, err := clientset.CoreV1().Pods("infra").List(context.Background(), metav1.ListOptions{})
// 	if err != nil {
// 		log.Fatalf("Failed to get pods: %v", err)
// 	}

// 	// 打印匹配的 Pod 信息
// 	for _, pod := range pods.Items {
// 		if strings.Contains(pod.Name, podNameKeyword) {
// 			fmt.Printf("Pod Name: %s\n", pod.Name)
// 			fmt.Printf("Namespace: %s\n", pod.Namespace)
// 			fmt.Println(strings.Repeat("-", 20))
// 		}
// 	}
// }
