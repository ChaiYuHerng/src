package ike

import (
	"free5gc/src/n3iwf/ike/handler"
	ike_message "free5gc/src/n3iwf/ike/message"
	"free5gc/src/n3iwf/logger"
	"net"

	"github.com/sirupsen/logrus"
)

var ikeLog *logrus.Entry

func init() {
	ikeLog = logger.IKELog
}

func Dispatch(udpConn *net.UDPConn, localAddr, remoteAddr *net.UDPAddr, msg []byte) {
	// As specified in RFC 7296 section 3.1, the IKE message send from/to UDP port 4500
	// should prepend a 4 bytes zero
	if localAddr.Port == 4500 {
		for i := 0; i < 4; i++ {
			if msg[i] != 0 {
				ikeLog.Warn(
					"[IKE] Received an IKE packet that does not prepend 4 bytes zero from UDP port 4500," +
						" this packet may be the UDP encapsulated ESP. The packet will not be handled.")
				return
			}
		}
		msg = msg[4:]
	}

	ikeMessage, err := ike_message.Decode(msg)
	if err != nil {
		ikeLog.Error(err)
		return
	}

	switch ikeMessage.ExchangeType {
	case ike_message.IKE_SA_INIT:
		handler.HandleIKESAINIT(udpConn, localAddr, remoteAddr, ikeMessage)
	case ike_message.IKE_AUTH:
		handler.HandleIKEAUTH(udpConn, localAddr, remoteAddr, ikeMessage)
	case ike_message.CREATE_CHILD_SA:
		handler.HandleCREATECHILDSA(udpConn, localAddr, remoteAddr, ikeMessage)
	default:
		ikeLog.Warnf("Unimplemented IKE message type, exchange type: %d", ikeMessage.ExchangeType)
	}

}
