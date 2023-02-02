package nostr

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Connection struct {
	socket *websocket.Conn
	mutex  sync.Mutex
}

func NewConnection(socket *websocket.Conn) *Connection {
	return &Connection{
		socket: socket,
	}
}

func (c *Connection) WriteJSON(v interface{}) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	err := c.socket.WriteJSON(v)
	if err != nil {
		err2 := c.Close()
		if err2 != nil {
			return err2
		}
		return err
	}
	return nil
}

func (c *Connection) WriteMessage(messageType int, data []byte) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	err := c.socket.WriteMessage(messageType, data)
	if err != nil {
		err2 := c.Close()
		if err2 != nil {
			return err2
		}
		return err
	}
	return nil
}

func (c *Connection) Close() error {
	err := c.socket.Close()
	if err != nil {
		err := c.socket.Close()
		if err != nil {
			err := c.socket.UnderlyingConn().Close()
			if err != nil {
				return err
			}
		}
		return err
	}
	return nil
}
