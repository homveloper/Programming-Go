package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

type CServer struct {
	m_mapRoom  map[string]*SRoom
	m_Commands chan SCommand
}

func NewServer() *CServer {
	return &CServer{
		m_mapRoom:  make(map[string]*SRoom),
		m_Commands: make(chan SCommand),
	}
}

func (Server *CServer) Run() {
	for Cmd := range Server.m_Commands {
		switch Cmd.m_ID {
		case CMD_NICK:
			Server.Nick(Cmd.m_Client, Cmd.m_Args)
		case CMD_JOIN:
			Server.Join(Cmd.m_Client, Cmd.m_Args)
		case CMD_ROOMS:
			Server.ListRoom(Cmd.m_Client, Cmd.m_Args)
		case CMD_MSG:
			Server.Msg(Cmd.m_Client, Cmd.m_Args)
		case CMD_QUIT:
			Server.Quit(Cmd.m_Client, Cmd.m_Args)
		}
	}
}

func (s *CServer) NewClient(Conn net.Conn) {
	log.Printf("New Client Has Connected : %s", Conn.RemoteAddr().String())

	Client := &CClient{
		m_Conn:     Conn,
		m_Nick:     "Anonymous",
		m_Commands: s.m_Commands,
	}

	Client.ReadInput()
}

func (Server *CServer) Nick(Client *CClient, Args []string) {
	Client.m_Nick = Args[1]
	Client.Msg(fmt.Sprintf("Now Your %s", Client.m_Nick))
}
func (Server *CServer) Join(Client *CClient, Args []string) {
	RoomName := Args[1]

	Room, bHas := Server.m_mapRoom[RoomName]
	if !bHas {
		Room = &SRoom{
			m_Name:      RoomName,
			m_mapMember: make(map[net.Addr]*CClient),
		}
		Server.m_mapRoom[RoomName] = Room
	}

	Room.m_mapMember[Client.m_Conn.RemoteAddr()] = Client
	Server.QuitRoom(Client)
	Client.m_Room = Room

	Room.BroadCast(Client, fmt.Sprintf("%s Has Joined The Room", Client.m_Nick))
	Client.Msg(fmt.Sprintf("Welcom to %s", Room.m_Name))
}
func (Server *CServer) ListRoom(Client *CClient, Args []string) {
	var arrayRoomName []string

	for Name := range Server.m_mapRoom {
		arrayRoomName = append(arrayRoomName, Name)
	}

	Client.Msg(fmt.Sprintf("Available Room Are : %s", strings.Join(arrayRoomName, ", ")))
}
func (Server *CServer) Msg(Client *CClient, Args []string) {
	if Client.m_Room == nil {
		Client.Err(errors.New("you must join the room first"))
		return
	}

	Client.m_Room.BroadCast(Client, Client.m_Nick+": "+strings.Join(Args[1:], " "))
}
func (Server *CServer) Quit(Client *CClient, Args []string) {
	log.Printf("Client Has Disconnected: %s", Client.m_Conn.RemoteAddr().String())
	Server.QuitRoom(Client)

	Client.Msg("Sad To See You Go :(")
	Client.m_Conn.Close()
}
func (Server *CServer) QuitRoom(Client *CClient) {
	if Client.m_Room != nil {
		delete(Client.m_Room.m_mapMember, Client.m_Conn.RemoteAddr())
		Client.m_Room.BroadCast(Client, fmt.Sprintf("%s Has Left The Room", Client.m_Nick))
	}
}
