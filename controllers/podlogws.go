package controllers

import (
	"bufio"
	"context"
	"net/http"

	"github.com/astaxie/beego"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
)

type PodLogSockjs struct {
	conn      sockjs.Session
	sizeChan  chan *remotecommand.TerminalSize
	context   string
	namespace string
	pod       string
	container string
}

func (self PodLogSockjs) Write(p []byte) (int, error) {
	err := self.conn.Send(string(p))
	return len(p), err
}

func HandlerLog(t *PodLogSockjs) error {
	config, err := buildConfigFromContextFlags(t.context, beego.AppConfig.String("kubeconfig"))
	if err != nil {
		beego.Info("no config file  ", err)
		config, err = rest.InClusterConfig()
		if err != nil {
			return err
		}
	}
	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		return err
	}

	req := clientset.CoreV1().Pods(t.namespace).GetLogs(t.pod, &v1.PodLogOptions{})
	// req := clientset.CoreV1().Pods(t.namespace).GetLogs(t.pod, &v1.PodLogOptions{Follow: true})
	podLogs, err := req.Stream(context.Background())
	if err != nil {
		return err
	}
	defer podLogs.Close()

	scanner := bufio.NewScanner(podLogs)
	for scanner.Scan() {
		line := scanner.Text()
		// 发送日志行到 WebSocket 连接
		_, err := t.Write([]byte(line))
		if err != nil {
			// 处理发送错误
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	// buf := new(bytes.Buffer)
	// _, err = io.Copy(buf, podLogs)
	// if err != nil {
	// 	return err
	// }
	// _, err = t.Write(buf.Bytes())
	// if err != nil {
	// 	return err
	// }

	// 	str := buf.String()
	// go func() {
	// 	buffer := make([]byte, 4096)
	// 	for {
	// 		n, err := podLogs.Read(buffer)
	// 		if err != nil {
	// 			beego.Error("Error reading pod logs:", err)
	// 			return
	// 		}

	// 		_, err = t.Write(buffer[:n])
	// 		if err != nil {
	// 			beego.Error("Error sending logs via WebSocket:", err)
	// 			return
	// 		}
	// 	}
	// }()
	return nil

}

// 实现http.handler 接口获取入参
func (self PodLogSockjs) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := r.FormValue("context")
	namespace := r.FormValue("namespace")
	pod := r.FormValue("pod")
	container := r.FormValue("container")
	Sockjshandler := func(session sockjs.Session) {
		t := &PodLogSockjs{session, make(chan *remotecommand.TerminalSize),
			context, namespace, pod, container}
		if err := HandlerLog(t); err != nil {
			beego.Error(err)
			beego.Error(HandlerLog(t))
		}
	}
	sockjs.NewHandler("/podlog/ws", sockjs.DefaultOptions, Sockjshandler).ServeHTTP(w, r)
}
