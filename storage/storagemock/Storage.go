package storagemock

import mock "github.com/stretchr/testify/mock"
import model "github.com/nghialv/promviz/model"
import storage "github.com/nghialv/promviz/storage"

// Storage is an autogenerated mock type for the Storage type
type Storage struct {
	mock.Mock
}

// Add provides a mock function with given fields: _a0
func (_m *Storage) Add(_a0 *model.Snapshot) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Snapshot) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *Storage) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetChunk provides a mock function with given fields: _a0
func (_m *Storage) GetChunk(_a0 int64) (storage.Chunk, error) {
	ret := _m.Called(_a0)

	var r0 storage.Chunk
	if rf, ok := ret.Get(0).(func(int64) storage.Chunk); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(storage.Chunk)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestSnapshot provides a mock function with given fields:
func (_m *Storage) GetLatestSnapshot() (*model.Snapshot, error) {
	ret := _m.Called()

	var r0 *model.Snapshot
	if rf, ok := ret.Get(0).(func() *model.Snapshot); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Snapshot)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

var _ storage.Storage = (*Storage)(nil)