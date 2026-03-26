package cmd

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/knoxctl/pkg/kube"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var getNamespacesCmd = &cobra.Command{
	Use:     "ns",
	Aliases: []string{"namespaces", "namespace"},
	Short:   "List all namespaces",
	Long:    `Display all namespaces in the Kubernetes cluster with their status and age.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := kube.GetClient(kubeconfig)
		if err != nil {
			return fmt.Errorf("error creating kubernetes client: %w", err)
		}

		namespaces, err := client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return fmt.Errorf("error listing namespaces: %w", err)
		}

		if len(namespaces.Items) == 0 {
			fmt.Println("No namespaces found.")
			return nil
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
		fmt.Fprintln(w, "NAME\tSTATUS\tAGE")

		for _, ns := range namespaces.Items {
			age := formatAge(ns.CreationTimestamp.Time)
			fmt.Fprintf(w, "%s\t%s\t%s\n",
				ns.Name, string(ns.Status.Phase), age)
		}

		w.Flush()
		return nil
	},
}

func init() {
	getCmd.AddCommand(getNamespacesCmd)
}
