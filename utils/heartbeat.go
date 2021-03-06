package utils

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/djherbis/times"
)

//go:generate counterfeiter -o ../fakes/fake_heartbeat.go . Heartbeat
type Heartbeat interface {
	Exists() (bool, error)
	Create() error
	Update() error
	GetLastUpdateTimestamp() (time.Time, error)
}

type heartbeat struct {
	log      *log.Logger
	filePath string
}

func NewHeartbeat(log *log.Logger, mountpoint string) Heartbeat {
	filePath := path.Join(mountpoint, ".ubiquity.lock")
	return &heartbeat{log: log, filePath: filePath}
}
func (l *heartbeat) Exists() (bool, error) {
	_, err := times.Stat(l.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
func (l *heartbeat) Create() error {
	fd, err := os.Create(l.filePath)
	if err != nil {
		return err
	}
	fd.Sync()
	fd.Close()
	return nil
}

func (l *heartbeat) Update() error {
	fd, err := os.Create(l.filePath)
	if err != nil {
		return err
	}
	fd.Sync()
	fd.Close()
	return nil
}

func (l *heartbeat) GetLastUpdateTimestamp() (time.Time, error) {
	fi, err := times.Stat(l.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return time.Time{}, err
		} else {
			return time.Time{}, fmt.Errorf("Failed to stat %s : %s\n", l.filePath, err.Error())
		}
	}
	return fi.ChangeTime(), nil
}
