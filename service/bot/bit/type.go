package bit

type Bit struct {
	Uid  int64
	Hash float64
	Bit  float64
	Food float64
}

type RedPacket struct {
	Status bool    `redis:"status"`
	Uid    int64   `redis:"uid"`
	Total  float64 `redis:"total"`
	Share  int64   `redis:"share"`
}
