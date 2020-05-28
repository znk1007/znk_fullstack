//go 兼容版本

// +build !go1.8

//Package ws websocket
package websocket

func (c *Conn) writeBufs(bufs ...[]byte) error {
	for _, buf := range bufs {
		if len(buf) > 0 {
			if _, err := c.conn.Write(buf); err != nil {
				return err
			}
		}
	}
	return nil
}
