package tftp

import (
	"io"

	tftpdata "beryju.io/gravity/tftp"
	"github.com/pin/tftp/v3"
	"go.uber.org/zap"
)

func (r *Role) readLogger(filename string, rf io.ReaderFrom) error {
	ot := rf.(tftp.OutgoingTransfer)
	r.log.Info("TFTP Read request", zap.String("filename", filename), zap.String("client", ot.RemoteAddr().IP.String()))
	return r.readHandler(filename, rf)
}

func (r *Role) readHandler(filename string, rf io.ReaderFrom) error {
	f, err := tftpdata.Root.Open(filename)
	if err != nil {
		return err
	}
	_, err = rf.ReadFrom(f)
	return err
}
