package liveMedia

import (
	. "groupsock"
)

const (
	// RTCP packet types:
	RTCP_PT_SR   = 200
	RTCP_PT_RR   = 201
	RTCP_PT_SDES = 202
	RTCP_PT_BYE  = 203
	RTCP_PT_APP  = 204

	// SDES tags:
	RTCP_SDES_END   = 0
	RTCP_SDES_CNAME = 1
	RTCP_SDES_NAME  = 2
	RTCP_SDES_EMAIL = 3
	RTCP_SDES_PHONE = 4
	RTCP_SDES_LOC   = 5
	RTCP_SDES_TOOL  = 6
	RTCP_SDES_NOTE  = 7
	RTCP_SDES_PRIV  = 8

	// overhead (bytes) of IP and UDP hdrs
	IP_UDP_HDR_SIZE = 28
)

type SDESItem struct {
	data []byte
}

// bytes, (1500, minus some allowance for IP, UDP, UMTP headers)
var maxRTCPPacketSize uint = 1450
var preferredPacketSize uint = 1000 // bytes

type RTCPInstance struct {
	typeOfEvent    int
	totSessionBW   uint
	CNAME          *SDESItem
	Sink           *RTPSink
	Source         *RTPSource
	outBuf         *OutPacketBuffer
	rtcpInterface  *RTPInterface
	ByeHandlerTask interface{}
	SRHandlerTask  interface{}
	RRHandlerTask  interface{}
}

func NewSDESItem(tag int, value string) *SDESItem {
	item := new(SDESItem)

	length := len(value)
	if length > 0xFF {
		length = 0xFF // maximum data length for a SDES item
	}

	//item.data[0] = tag
	//item.data[1] = length
	return item
}

func (this *SDESItem) totalSize() uint {
	return 2 + uint(this.data[1])
}

func NewRTCPInstance(rtcpGS *GroupSock, totSessionBW uint, cname string) *RTCPInstance {
	rtcp := new(RTCPInstance)
	rtcp.typeOfEvent = EVENT_REPORT
	rtcp.totSessionBW = totSessionBW
	rtcp.outBuf = NewOutPacketBuffer(preferredPacketSize, maxRTCPPacketSize)
	rtcp.CNAME = NewSDESItem(RTCP_SDES_CNAME, cname)

	rtcp.rtcpInterface = NewRTPInterface(rtcp, rtcpGS)
	rtcp.rtcpInterface.startNetworkReading()

	go rtcp.incomingReportHandler()
	//this.onExpire(rtcp)
	return rtcp
}

func (this *RTCPInstance) setSpecificRRHandler() {
}

func (this *RTCPInstance) SetByeHandler(handlerTask interface{}, clientData interface{}) {
	//this.byeHandlerTask = handlerTask
	//this.byeHandlerClientData = clientData
}

func (this *RTCPInstance) setSRHandler() {
}

func (this *RTCPInstance) setRRHandler() {
}

func (this *RTCPInstance) incomingReportHandler() {
}

func (this *RTCPInstance) onReceive() {
}

func (this *RTCPInstance) sendReport() {
	// Begin by including a SR and/or RR report:
	this.addReport()

	// Then, include a SDES:
	//this.addSDES()

	// Send the report:
	this.sendBuiltPacket()
}

func (this *RTCPInstance) sendBuiltPacket() {
	reportSize := this.outBuf.curPacketSize()
	this.rtcpInterface.sendPacket(this.outBuf.packet(), reportSize)
	this.outBuf.resetOffset()

	//this.lastSentSize = IP_UDP_HDR_SIZE + reportSize
	//this.haveJustSentPacket = true
	//this.lastPacketSentSize = reportSize
}

func (this *RTCPInstance) addReport() {
	/*
	   if this.sink != nil {
	       if this.sink.enableRTCPReports() {
	           return
	       }

	       if this.sink.nextTimestampHasBeenPreset() {
	           return
	       }

	       this.addSR()
	   } else if this.source != nil {
	       this.addRR()
	   }*/
}

func (this *RTCPInstance) addSR() {
}

func (this *RTCPInstance) addRR() {
}
