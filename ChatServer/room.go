package main

import "net"

type SRoom struct {
	m_Name      string
	m_mapMember map[net.Addr]*CClient
}

func (Room *SRoom) BroadCast(Sender *CClient, Msg string) {
	for Addr, Member := range Room.m_mapMember {
		if Addr != Sender.m_Conn.RemoteAddr() {
			Member.Msg(Msg)
		}
	}
}
