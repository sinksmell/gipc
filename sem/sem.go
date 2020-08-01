package sem

/*
#include "sem.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

var (
	ObtainErr   = errors.New("obtain a semaphore err")
	SetValueErr = errors.New("set semaphore value err")
	DestroyErr  = errors.New("destroy semaphore err")
	PErr        = errors.New("sem p err")
	VErr        = errors.New("sem v err")
)

const (
	okCode   = 0
	failCode = -1
)

// Semaphore is a abstract interface
type Semaphore interface {
	Destroy() error
	P() error
	V() error
}

// SystemVSem is systemV semaphore
type SystemVSem struct {
	ID int
}

// P ...
func (sem *SystemVSem) P() error {
	var buf C.struct_sembuf
	buf.sem_num = 0
	buf.sem_op = -1
	buf.sem_flg = C.SEM_UNDO
	code := C.semop(C.int(sem.ID), (*C.struct_sembuf)(unsafe.Pointer(&buf)), C.size_t(1))
	if int(code) != okCode {
		return fmt.Errorf("%w code=%d", PErr, int(code))
	}
	return nil
}

// V ...
func (sem *SystemVSem) V() error {
	var buf C.struct_sembuf
	buf.sem_num = 0
	buf.sem_op = 1
	buf.sem_flg = C.SEM_UNDO
	code := C.semop(C.int(sem.ID), (*C.struct_sembuf)(unsafe.Pointer(&buf)), C.size_t(1))
	if int(code) != okCode {
		return fmt.Errorf("%w code=%d", VErr, int(code))
	}
	return nil
}

// Get Or Create a system V semaphore
func GetSysVSem(key int) (*SystemVSem, error) {
	mod := 0666 | int(C.IPC_CREAT)
	id := C.semget(C.key_t(key), C.int(1), C.int(mod))
	if int64(id) == failCode {
		return nil, fmt.Errorf("%w id=%d", ObtainErr, int(id))
	}
	return &SystemVSem{ID: int(id)}, nil
}

// SetValue init
func (sem *SystemVSem) SetValue(v int64) error {
	var s C.union_semun
	*(*C.int)(unsafe.Pointer(&(s))) = C.int(v)
	code := C.cgo_semctl(C.int(sem.ID), C.int(0), C.SETVAL, s)
	if int(code) != okCode {
		return fmt.Errorf("%w code=%d", SetValueErr, int(code))
	}
	return nil
}

// Destroy delete a semaphore
func (sem *SystemVSem) Destroy() error {
	var s C.union_semun
	code := C.cgo_semctl(C.int(sem.ID), 0, C.IPC_RMID, s)
	if int(code) != okCode {
		return fmt.Errorf("%w code=%d", DestroyErr, int(code))
	}
	return nil
}
