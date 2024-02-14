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

	// Setup Kubernetes
	fmt.Printf("%sSetup Kubernetes%s\n", GREEN, ENDCOLOR)

	// Create kubernetes repo
	fmt.Printf("%sCreate kubernetes repo%s\n", GREEN, ENDCOLOR)
	kubeRepoScript := `cat <<EOF | sudo tee /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-\$basearch
enabled=1
gpgcheck=1
gpgkey=https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
exclude=kubelet kubeadm kubectl
EOF`
	err := runCommand(kubeRepoScript)
	if err != nil {
		fmt.Println("Error creating kubernetes repo:", err)
		return
	}

	// Set SELinux in permissive mode (Disable effectively)
	fmt.Printf("%sSet SELinux in permissive mode (Disable effectively)%s\n", GREEN, ENDCOLOR)
	err = runCommand("sudo setenforce 0")
	if err != nil {
		fmt.Println("Error setting SELinux to permissive mode:", err)
		return
	}
	err = runCommand(`sudo sed -i 's/^SELINUX=enforcing$/SELINUX=permissive/' /etc/selinux/config`)
	if err != nil {
		fmt.Println("Error updating SELinux config:", err)
		return
	}

	// Install kubelet, kubeadm and kubectl
	fmt.Printf("%sInstall kubelet, kubeadm and kubectl%s\n", GREEN, ENDCOLOR)
	err = runCommand("sudo yum install -y kubelet-1.26.0-0 kubeadm-1.26.0-0 kubectl-1.26.0-0 --disableexcludes=kubernetes")
	if err != nil {
		fmt.Println("Error installing kubelet, kubeadm, kubectl:", err)
		return
	}

	// Memory swapoff
	fmt.Printf("%sMemory swapoff%s\n", GREEN, ENDCOLOR)
	err = runCommand(`sudo sed -i '/ swap / s/^\(.*\)$/#\1/g' /etc/fstab && sudo swapoff -a`)
	if err != nil {
		fmt.Println("Error disabling swap:", err)
		return
	}

	// Enable Kubelet
	fmt.Printf("%sEnable Kubelet%s\n", GREEN, ENDCOLOR)
	err = runCommand("sudo systemctl enable kubelet")
	if err != nil {
		fmt.Println("Error enabling kubelet:", err)
		return
	}

	// Configure the firewall rules on the ports
	fmt.Printf("%sConfigure the firewall rules on the ports%s\n", GREEN, ENDCOLOR)
	firewallCommands := []string{
		"firewall-cmd --permanent --add-port=6443/tcp",
		"firewall-cmd --permanent --add-port=2379-2380/tcp",
		"firewall-cmd --permanent --add-port=10250/tcp",
		"firewall-cmd --permanent --add-port=10251/tcp",
		"firewall-cmd --permanent --add-port=10252/tcp",
		"firewall-cmd --permanent --add-port=10255/tcp",
		"firewall-cmd --reload",
	}
	for _, cmd := range firewallCommands {
		err = runCommand(cmd)
		if err != nil {
			fmt.Printf("Error running firewall command '%s': %v\n", cmd, err)
			return
		}
	}

	// Add kernel modules
	fmt.Printf("%sAdd kernel modules%s\n", GREEN, ENDCOLOR)
	modprobeCommands := []string{
		"sudo modprobe overlay",
		"sudo modprobe br_netfilter",
		`sudo tee /etc/modules-load.d/containerd.conf <<EOF
overlay
br_netfilter
EOF`,
	}
	for _, cmd := range modprobeCommands {
		err = runCommand(cmd)
		if err != nil {
			fmt.Printf("Error running modprobe command '%s': %v\n", cmd, err)
			return
		}
	}

	// Set the bridged traffic for iptables
	fmt.Printf("%sSet the bridged traffic for iptables%s\n", GREEN, ENDCOLOR)
	iptablesCommands := []string{
		`sudo tee /etc/sysctl.d/kubernetes.conf<<EOF
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1
EOF`,
		"sudo sysctl --system",
	}
	for _, cmd := range iptablesCommands {
		err = runCommand(cmd)
		if err != nil {
			fmt.Printf("Error running iptables command '%s': %v\n", cmd, err)
			return
		}
	}

	// Check and install containerd
	fmt.Printf("%sCheck and install containerd%s\n", GREEN, ENDCOLOR)
	if isActive, _ := isSystemdActive("containerd"); isActive {
		err = runCommand(`sudo sed -i 's/disabled_plugins/#disabled_plugins/g' /etc/containerd/config.toml && sudo systemctl restart containerd`)
	} else {
		err = runCommand(`sudo yum install -y containerd.io && mkdir -p /etc/containerd && containerd config default>/etc/containerd/config.toml && sudo sed -i 's/disabled_plugins/#disabled_plugins/g' /etc/containerd/config.toml && sudo systemctl restart containerd && sudo systemctl enable containerd`)
	}
	if err != nil {
		fmt.Println("Error checking and installing containerd:", err)
		return
	}

	// Pull kubeadm config images
	fmt.Printf("%sPull kubeadm config images%s\n", GREEN, ENDCOLOR)
	err = runCommand("sudo kubeadm config images pull --cri-socket unix:///run/containerd/containerd.sock --kubernetes-version v1.26.0")
	if err != nil {
		fmt.Println("Error pulling kubeadm config images:", err)
		return
	}

	fmt.Println("Script execution completed successfully.")
}

func isSystemdActive(service string) (bool, error) {
	cmd := exec.Command("systemctl", "is-active", service)
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() == 3 {
			// Exit code 3 indicates that the service is inactive
			return false, nil
		}
		return false, err
	}
	return true, nil
}
