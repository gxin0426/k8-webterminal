package controllers

import (
	"context"
	"strings"

	"github.com/astaxie/beego"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type SearchPodController struct {
	beego.Controller
}

type PodInfo struct {
	Name      string
	Status    string
	Labels    map[string]string
	Node      string
	Restart   int32
	StartTime string
}

func (self *SearchPodController) Get() {

	config, err := buildConfigFromContextFlags("", beego.AppConfig.String("kubeconfig"))
	if err != nil {
		beego.Info("no config file  ", err)
		config, err = rest.InClusterConfig()
		if err != nil {
			beego.Error("get config err : ", err)
		}
	}

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		beego.Error("create clientset err: ", err)
	}
	namespace := self.GetString("namespace")
	podNameKeyword := self.GetString("pod")
	pods, err := clientset.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})

	var podList []PodInfo
	for _, pod := range pods.Items {
		if strings.Contains(pod.Name, podNameKeyword) {
			podInfo := PodInfo{
				Name:      pod.Name,
				Status:    string(pod.Status.Phase),
				Labels:    pod.Labels,
				Node:      pod.Spec.NodeName,
				StartTime: pod.Status.StartTime.String(),
				Restart:   pod.Status.ContainerStatuses[0].RestartCount,
			}
			podList = append(podList, podInfo)
		}

	}

	self.Data["Pods"] = podList
	self.TplName = "podinfo.html" // Pod 列表模板

}
