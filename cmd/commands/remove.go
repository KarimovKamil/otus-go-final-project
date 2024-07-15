package commands

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/KarimovKamil/otus-go-final-project/internal/entity/request"
	"github.com/mailru/easyjson"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove network from list",
	Run: func(cmd *cobra.Command, args []string) {
		usedCommand := strings.Fields(cmd.CommandPath())[1]
		if len(args) != 2 {
			fmt.Println("Usage: abf " + usedCommand + " remove <ip> <mask>")
			return
		}
		if usedCommand == whiteList {
			removeFromList(args[0], args[1], whiteList)
			return
		} else if usedCommand == blackList {
			removeFromList(args[0], args[1], blackList)
			return
		}
		fmt.Println("Unknown command")
	},
}

func init() {
	removeCmdForWhiteList := *removeCmd
	removeCmdForBlackList := *removeCmd
	whiteListCmd.AddCommand(&removeCmdForWhiteList)
	blackListCmd.AddCommand(&removeCmdForBlackList)
}

func removeFromList(ip, mask, list string) {
	networkRequest := &request.NetworkRequest{Network: ip + "/" + mask}
	requestBody, _ := easyjson.Marshal(networkRequest)

	httpRequest, err := http.NewRequestWithContext(context.Background(), http.MethodDelete,
		serverAddress+"/api/"+list, bytes.NewBuffer(requestBody))
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(httpRequest.Body)
	if err != nil {
		fmt.Println(err)
	}

	response, err := http.DefaultClient.Do(httpRequest)
	if response == nil {
		fmt.Println("No response, check server address")
		return
	}
	defer response.Body.Close()
	if err != nil {
		fmt.Println(err)
	}

	switch response.StatusCode {
	case 200:
		fmt.Println("Successfully removed from " + list)
	case 400:
		fmt.Println("Invalid ip or mask")
	case 500:
		fmt.Println("Internal server error")
	default:
		fmt.Println("Unknown error")
	}
}
