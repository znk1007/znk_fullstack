package ws

import "crypto/tls"

type joinReader struct {
	// c *Conn
}

func doHandshake(tlsConn *tls.Conn, cfg *tls.Config) error {
	if err := tlsConn.Handshake(); err != nil {
		return err
	}
	if !cfg.InsecureSkipVerify {
		if err := tlsConn.VerifyHostname(cfg.ServerName); err != nil {
			return err
		}
	}
	return nil
}
