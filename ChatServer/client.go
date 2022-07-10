package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type CClient struct {
	m_Conn     net.Conn
	m_Nick     string
	m_Room     *SRoom
	m_Commands chan<- SCommand
}

func (Client *CClient) ReadInput() {
	for {
		RecvMsg, Err := bufio.NewReader(Client.m_Conn).ReadString('\n')

		if Err != nil {
			return
		}

		RecvMsg = strings.Trim(RecvMsg, "\r\n")
		Args := strings.Split(RecvMsg, " ")

		Cmd := strings.TrimSpace(Args[0])

		switch Cmd {
		case "/nick":
			Client.m_Commands <- SCommand{
				m_ID:     CMD_NICK,
				m_Client: Client,
				m_Args:   Args,
			}
		case "/join":
			Client.m_Commands <- SCommand{
				m_ID:     CMD_JOIN,
				m_Client: Client,
				m_Args:   Args,
			}
		case "/rooms":
			Client.m_Commands <- SCommand{
				m_ID:     CMD_ROOMS,
				m_Client: Client,
				m_Args:   Args,
			}
		case "/msg":
			Client.m_Commands <- SCommand{
				m_ID:     CMD_MSG,
				m_Client: Client,
				m_Args:   Args,
			}
		case "/quit":
			Client.m_Commands <- SCommand{
				m_ID:     CMD_QUIT,
				m_Client: Client,
				m_Args:   Args,
			}
		default:
			Client.Err(fmt.Errorf("unknown command: %s", Cmd))
		}
	}
}

func (c *CClient) Err(Err error) {
	c.m_Conn.Write([]byte("Err: " + Err.Error() + "\n"))
}

func (c *CClient) Msg(Msg string) {
	c.m_Conn.Write([]byte("> " + Msg + "\n"))
}
