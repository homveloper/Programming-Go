package main

type CommandID int

const (
	CMD_NICK CommandID = iota
	CMD_JOIN
	CMD_ROOMS
	CMD_MSG
	CMD_QUIT
)

type SCommand struct {
	m_ID     CommandID
	m_Client *CClient
	m_Args   []string
}
