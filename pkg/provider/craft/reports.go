package craft

import (
	"encoding/json"
	"os"

	"github.com/samber/lo"
)

type PacketsReport struct {
	Handshake     PacketsReportGroup `json:"handshake"`
	Status        PacketsReportGroup `json:"status"`
	Login         PacketsReportGroup `json:"login"`
	Configuration PacketsReportGroup `json:"configuration"`
	Play          PacketsReportGroup `json:"play"`
}

type PacketsReportGroup struct {
	Clientbound map[string]PacketDescription `json:"clientbound"`
	Serverbound map[string]PacketDescription `json:"serverbound"`
}

type PacketDescription struct {
	ProtocolId int32 `json:"protocol_id"`
}

func ReadPacketsReport(path string) (report *PacketsReport, err error) {
	var (
		file *os.File
	)

	file, err = os.Open(path)
	if err != nil {
		return
	}

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&report)
	if err != nil {
		return
	}

	return
}

func (g *PacketsReportGroup) GeneratePackets(state SessionState) (result []Packet) {
	return lo.Concat(
		lo.MapToSlice(g.Clientbound, func(key string, input PacketDescription) Packet {
			return Packet{
				State: state,
				Kind:  PacketKindClientbound,
				Id:    input.ProtocolId,
				Name:  key,
			}
		}),
		lo.MapToSlice(g.Serverbound, func(key string, input PacketDescription) Packet {
			return Packet{
				State: state,
				Kind:  PacketKindServerbound,
				Id:    input.ProtocolId,
				Name:  key,
			}
		}),
	)
}

func (r *PacketsReport) GeneratePackets() (result []Packet) {
	result = make([]Packet, 0)

	result = append(result, r.Handshake.GeneratePackets(SessionStateHandshake)...)
	result = append(result, r.Status.GeneratePackets(SessionStateStatus)...)
	result = append(result, r.Login.GeneratePackets(SessionStateLogin)...)
	result = append(result, r.Configuration.GeneratePackets(SessionStateConfiguration)...)
	result = append(result, r.Play.GeneratePackets(SessionStatePlay)...)

	return
}
