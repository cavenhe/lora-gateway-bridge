// Code generated by "stringer -type=PacketType"; DO NOT EDIT

package gateway

import "fmt"

const _PacketType_name = "PushDataPushACKPullDataPullRespPullACK"

var _PacketType_index = [...]uint8{0, 8, 15, 23, 31, 38}

func (i PacketType) String() string {
	if i >= PacketType(len(_PacketType_index)-1) {
		return fmt.Sprintf("PacketType(%d)", i)
	}
	return _PacketType_name[_PacketType_index[i]:_PacketType_index[i+1]]
}
