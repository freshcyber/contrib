package uuid

import (
	"fmt"
	flake "github.com/zheng-ji/goSnowFlake"
)

// GetFlakeID GetFlakeID
func GetFlakeID() int64 {
	// Params: Given the workerId, 0 < workerId < 1024
	var id int64
	iw, err := flake.NewIdWorker(1)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	id, err = iw.NextId()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return id
}

// IwFlake IwFlake
var IwFlake *flake.IdWorker

// Flake Flake
type Flake struct {
	iw *flake.IdWorker
}

// Init Init
func (f *Flake) Init() error {
	_iw, err := flake.NewIdWorker(512)
	if err != nil {
		fmt.Println(err)
		return err
	}
	f.iw = _iw
	return nil
}

// NextID NextID
func (f *Flake) NextID() (int64, error) {
	_id, err := f.iw.NextId()
	if err != nil {
		fmt.Println("NextID error:", err)
		return -1, err
	}
	return _id, nil
}
