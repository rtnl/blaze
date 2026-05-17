package craft

type PacketKind string

const (
	PacketKindClientbound PacketKind = "clientbound"
	PacketKindServerbound PacketKind = "serverbound"
)

type Packet struct {
	State SessionState `json:"state"`
	Kind  PacketKind   `json:"kind"`
	Id    int32        `json:"id"`
	Name  string       `json:"name"`
}
