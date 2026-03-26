package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	namespace     string
	allNamespaces bool
	kubeconfig    string
)

var rootCmd = &cobra.Command{
	Use:   "knoxctl",
	Short: "knoxctl - A simplified kubectl-like CLI tool",
	Long: `
██╗  ██╗███╗   ██╗ ██████╗ ██╗  ██╗ ██████╗████████╗██╗     
██║ ██╔╝████╗  ██║██╔═══██╗╚██╗██╔╝██╔════╝╚══██╔══╝██║     
█████╔╝ ██╔██╗ ██║██║   ██║ ╚███╔╝ ██║        ██║   ██║     
██╔═██╗ ██║╚██╗██║██║   ██║ ██╔██╗ ██║        ██║   ██║     
██║  ██╗██║ ╚████║╚██████╔╝██╔╝ ██╗╚██████╗   ██║   ███████╗
╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚═╝  ╚═╝ ╚═════╝   ╚═╝   ╚══════╝

knoxctl is a lightweight CLI tool for interacting with Kubernetes clusters.
It provides quick access to common operations like listing pods, deployments, 
services, nodes, and namespaces with a clean, formatted output.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "default", "The namespace to use for the request")
	rootCmd.PersistentFlags().BoolVarP(&allNamespaces, "all-namespaces", "A", false, "List resources across all namespaces")
	rootCmd.PersistentFlags().StringVar(&kubeconfig, "kubeconfig", "", "Path to the kubeconfig file (default: ~/.kube/config)")
}
