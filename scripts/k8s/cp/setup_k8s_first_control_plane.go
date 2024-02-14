package main

import (
	"fmt"
	"os"
	"os/exec"
)

func runCommand(command string) error {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	// Centos X
	//RED := "\033[31m"
	GREEN := "\033[32m"
	ENDCOLOR := "\033[0m"

	// Setup Kubernetes Control Plane
	fmt.Printf("%sSetup Kubernetes Control Plane%s\n", GREEN, ENDCOLOR)

	// Create kubernetes cluster
	fmt.Printf("%sCreate kubernetes cluster%s\n", GREEN, ENDCOLOR)
	// Execute only on the first control plane. Use join command to join to cluster.
	// You will find to join command at the result of the below command on the control plane server.
	err := runCommand(`sudo kubeadm init --pod-network-cidr=10.244.0.0/16 --upload-certs --kubernetes-version=v1.26.0 --control-plane-endpoint=$(hostname) --ignore-preflight-errors=all --cri-socket unix:///run/containerd/containerd.sock`)
	if err != nil {
		fmt.Println("Error during kubeadm init:", err)
		return
	}

	// Enable local user to access cluster info
	fmt.Printf("%sEnable local user to access cluster info%s\n", GREEN, ENDCOLOR)
	err = runCommand(`mkdir -p $HOME/.kube && sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config && sudo chown $(id -u):$(id -g) $HOME/.kube/config && export KUBECONFIG=/etc/kubernetes/admin.conf >> ~/.bashrc && source ~/.bashrc`)
	if err != nil {
		fmt.Println("Error during kubeconfig setup:", err)
		return
	}

	// Apply the CNI
	err = runCommand(`kubectl apply -f https://github.com/coreos/flannel/raw/master/Documentation/kube-flannel.yml`)
	if err != nil {
		fmt.Println("Error during CNI application:", err)
		return
	}

	// Disable pod schedule on control plane
	fmt.Printf("%sDisable pod schedule on master and control plane%s\n", GREEN, ENDCOLOR)
	err = runCommand(`kubectl taint node $(hostname) node-role.kubernetes.io/control-plane:NoSchedule-`)
	if err != nil {
		fmt.Println("Error during taint setup:", err)
		return
	}

	// Add control plane label for role
	err = runCommand(fmt.Sprintf(`kubectl label nodes $(hostname -s) "kubernetes.io/role=control-plane"`))
	if err != nil {
		fmt.Println("Error during label setup:", err)
		return
	}

	// Create join command and certificate key
	fmt.Printf("%sCreate join command and certificate key%s\n", GREEN, ENDCOLOR)
	JOINCOMMAND := "kubeadm token create --print-join-command"
	CERTIFICATEKEY := "kubeadm init phase upload-certs --upload-certs | grep -vw -e certificate -e Namespace"

	// Create control plane script to join the kubernetes cluster as control plane role.
	fmt.Printf("%sCreate control plane script to join the kubernetes cluster as control plane%s\n", GREEN, ENDCOLOR)
	setupControlPlaneScript := fmt.Sprintf(`#!/bin/bash
# Centos X
echo $(%s) --control-plane --certificate-key $(%s) \
  --node-name $(hostname -s) \
  --node-labels "kubernetes.io/role=control-plane"
`, JOINCOMMAND, CERTIFICATEKEY)
	err = runCommand(fmt.Sprintf(`echo '%s' | sudo tee setup_k8s_control_plane.sh`, setupControlPlaneScript))
	if err != nil {
		fmt.Println("Error creating setup_k8s_control_plane.sh script:", err)
		return
	}

	// Create worker script to join the kubernetes cluster as worker role
	fmt.Printf("%sCreate worker script to join the kubernetes cluster as worker role%s\n", GREEN, ENDCOLOR)
	setupWorkerScript := fmt.Sprintf(`#!/bin/bash
# Centos X
echo $(%s) \
  --node-name $(hostname -s) \
  --node-labels "kubernetes.io/role=worker"
`, JOINCOMMAND)
	err = runCommand(fmt.Sprintf(`echo '%s' | sudo tee setup_k8s_worker.sh`, setupWorkerScript))
	if err != nil {
		fmt.Println("Error creating setup_k8s_worker.sh script:", err)
		return
	}

	fmt.Println("Script execution completed successfully.")
}
