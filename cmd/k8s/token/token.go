package token

import (
	"context"
	"fmt"
	"log"

	k8s_cmd "github.com/sikalabs/slu/cmd/k8s"
	"github.com/sikalabs/slu/utils/k8s"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CmdFlagServiceAccount string
var CmdFlagNamespace string

var Cmd = &cobra.Command{
	Use:   "token",
	Short: "Bumb & commit version of (SL) Go application",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		clientset, defaultNamespace, _ := k8s.KubernetesClient()

		namespace := defaultNamespace
		if CmdFlagNamespace != "" {
			namespace = CmdFlagNamespace
		}

		saClient := clientset.CoreV1().ServiceAccounts(namespace)
		secretClient := clientset.CoreV1().Secrets(namespace)

		sa, err := saClient.Get(context.TODO(), CmdFlagServiceAccount, metav1.GetOptions{})
		if err != nil {
			log.Fatal(err)
		}
		secret, err := secretClient.Get(context.TODO(), sa.Secrets[0].Name, metav1.GetOptions{})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(secret.Data["token"]))
	},
}

func init() {
	k8s_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&CmdFlagServiceAccount,
		"service-account",
		"s",
		"default",
		"Kubernetes ServiceAccount",
	)
	Cmd.Flags().StringVarP(
		&CmdFlagNamespace,
		"namespace",
		"n",
		"",
		"Kubernetes Namespace",
	)
}
