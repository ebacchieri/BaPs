package protocol

import (
	"errors"
	"fmt"
	"./protocol/mx"

	"github.com/bytedance/sonic"
	"./protocol/cmd"
	"./protocol/proto"
)

type NetworkProtocolResponse struct {
	Packet   string `json:"packet"`   // 包数据
	Protocol string `json:"protocol"` // 协议名称
}

// UnmarshalRequest 解码req数据包
func UnmarshalRequest(b []byte) (mx.Message, *proto.BasePacket, error) {
	base := new(proto.BasePacket)
	err := sonic.Unmarshal(b, base)
	if err != nil {
		return nil, nil, err
	}
	packet := cmd.Get().GetRequestPacketByCmdId(base.Protocol)
	if packet == nil {
		return nil, nil, errors.New(fmt.Sprintf("request unknown cmd id: %v", base.Protocol))
	}
	err = sonic.Unmarshal(b, packet)
	if err != nil {
		return nil, nil, err
	}

	return packet, base, nil
}

// MarshalRequest 编码req数据包
func MarshalRequest(m mx.Message) ([]byte, error) {
	return sonic.Marshal(m)
}

// UnmarshalResponse 解码rsp数据包
func UnmarshalResponse(bin []byte) (mx.Message, proto.Protocol, error) {
	ojb1 := new(NetworkProtocolResponse)
	err := sonic.Unmarshal(bin, ojb1)
	if err != nil {
		return nil, 0, err
	}
	rspCmd := proto.Protocol_Common_Cheat.Value(ojb1.Protocol)
	ojb2 := cmd.Get().GetResponsePacketByCmdId(rspCmd)
	if ojb2 == nil {
		return nil, 0, errors.New(fmt.Sprintf("response unknown cmd id: %v", ojb1.Protocol))
	}
	err = sonic.UnmarshalString(ojb1.Packet, ojb2)
	return ojb2, rspCmd, nil
}

// MarshalResponse 编码rsp数据包
func MarshalResponse(m mx.Message) (*NetworkProtocolResponse, error) {
	if m == nil {
		return nil, errors.New("message nil")
	}
	jsonData, err := sonic.MarshalString(m)
	if err != nil {
		return nil, err
	}
	v := &NetworkProtocolResponse{
		Packet:   jsonData,
		Protocol: cmd.Get().GetCmdIdByProtoObj(m).String(),
	}

	return v, nil
}
