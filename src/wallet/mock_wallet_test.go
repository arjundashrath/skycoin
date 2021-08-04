// Code generated by mockery v1.0.0. DO NOT EDIT.

package wallet

import (
	cipher "github.com/skycoin/skycoin/src/cipher"
	bip44 "github.com/skycoin/skycoin/src/cipher/bip44"

	crypto "github.com/skycoin/skycoin/src/cipher/crypto"

	mock "github.com/stretchr/testify/mock"
)

// MockWallet is an autogenerated mock type for the Wallet type
type MockWallet struct {
	mock.Mock
}

// Accounts provides a mock function with given fields:
func (_m *MockWallet) Accounts() []Bip44Account {
	ret := _m.Called()

	var r0 []Bip44Account
	if rf, ok := ret.Get(0).(func() []Bip44Account); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Bip44Account)
		}
	}

	return r0
}

// Bip44Coin provides a mock function with given fields:
func (_m *MockWallet) Bip44Coin() *bip44.CoinType {
	ret := _m.Called()

	var r0 *bip44.CoinType
	if rf, ok := ret.Get(0).(func() *bip44.CoinType); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bip44.CoinType)
		}
	}

	return r0
}

// Clone provides a mock function with given fields:
func (_m *MockWallet) Clone() Wallet {
	ret := _m.Called()

	var r0 Wallet
	if rf, ok := ret.Get(0).(func() Wallet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Wallet)
		}
	}

	return r0
}

// Coin provides a mock function with given fields:
func (_m *MockWallet) Coin() CoinType {
	ret := _m.Called()

	var r0 CoinType
	if rf, ok := ret.Get(0).(func() CoinType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(CoinType)
	}

	return r0
}

// CopyFromRef provides a mock function with given fields: src
func (_m *MockWallet) CopyFromRef(src Wallet) {
	_m.Called(src)
}

// CryptoType provides a mock function with given fields:
func (_m *MockWallet) CryptoType() crypto.CryptoType {
	ret := _m.Called()

	var r0 crypto.CryptoType
	if rf, ok := ret.Get(0).(func() crypto.CryptoType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(crypto.CryptoType)
	}

	return r0
}

// Deserialize provides a mock function with given fields: data
func (_m *MockWallet) Deserialize(data []byte) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EntriesLen provides a mock function with given fields: options
func (_m *MockWallet) EntriesLen(options ...Option) (int, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int
	if rf, ok := ret.Get(0).(func(...Option) int); ok {
		r0 = rf(options...)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...Option) error); ok {
		r1 = rf(options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Erase provides a mock function with given fields:
func (_m *MockWallet) Erase() {
	_m.Called()
}

// Filename provides a mock function with given fields:
func (_m *MockWallet) Filename() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Fingerprint provides a mock function with given fields:
func (_m *MockWallet) Fingerprint() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GenerateAddresses provides a mock function with given fields: options
func (_m *MockWallet) GenerateAddresses(options ...Option) ([]cipher.Addresser, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []cipher.Addresser
	if rf, ok := ret.Get(0).(func(...Option) []cipher.Addresser); ok {
		r0 = rf(options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cipher.Addresser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...Option) error); ok {
		r1 = rf(options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddresses provides a mock function with given fields: options
func (_m *MockWallet) GetAddresses(options ...Option) ([]cipher.Addresser, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []cipher.Addresser
	if rf, ok := ret.Get(0).(func(...Option) []cipher.Addresser); ok {
		r0 = rf(options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cipher.Addresser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...Option) error); ok {
		r1 = rf(options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEntries provides a mock function with given fields: options
func (_m *MockWallet) GetEntries(options ...Option) (Entries, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 Entries
	if rf, ok := ret.Get(0).(func(...Option) Entries); ok {
		r0 = rf(options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Entries)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...Option) error); ok {
		r1 = rf(options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEntry provides a mock function with given fields: addr, options
func (_m *MockWallet) GetEntry(addr cipher.Addresser, options ...Option) (Entry, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, addr)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 Entry
	if rf, ok := ret.Get(0).(func(cipher.Addresser, ...Option) Entry); ok {
		r0 = rf(addr, options...)
	} else {
		r0 = ret.Get(0).(Entry)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(cipher.Addresser, ...Option) error); ok {
		r1 = rf(addr, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEntryAt provides a mock function with given fields: i, options
func (_m *MockWallet) GetEntryAt(i int, options ...Option) (Entry, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, i)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 Entry
	if rf, ok := ret.Get(0).(func(int, ...Option) Entry); ok {
		r0 = rf(i, options...)
	} else {
		r0 = ret.Get(0).(Entry)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, ...Option) error); ok {
		r1 = rf(i, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasEntry provides a mock function with given fields: addr, options
func (_m *MockWallet) HasEntry(addr cipher.Addresser, options ...Option) (bool, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, addr)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 bool
	if rf, ok := ret.Get(0).(func(cipher.Addresser, ...Option) bool); ok {
		r0 = rf(addr, options...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(cipher.Addresser, ...Option) error); ok {
		r1 = rf(addr, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsEncrypted provides a mock function with given fields:
func (_m *MockWallet) IsEncrypted() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsTemp provides a mock function with given fields:
func (_m *MockWallet) IsTemp() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Label provides a mock function with given fields:
func (_m *MockWallet) Label() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// LastSeed provides a mock function with given fields:
func (_m *MockWallet) LastSeed() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Lock provides a mock function with given fields: password
func (_m *MockWallet) Lock(password []byte) error {
	ret := _m.Called(password)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ScanAddresses provides a mock function with given fields: scanN, tf
func (_m *MockWallet) ScanAddresses(scanN uint64, tf TransactionsFinder) ([]cipher.Addresser, error) {
	ret := _m.Called(scanN, tf)

	var r0 []cipher.Addresser
	if rf, ok := ret.Get(0).(func(uint64, TransactionsFinder) []cipher.Addresser); ok {
		r0 = rf(scanN, tf)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cipher.Addresser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, TransactionsFinder) error); ok {
		r1 = rf(scanN, tf)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Secrets provides a mock function with given fields:
func (_m *MockWallet) Secrets() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Seed provides a mock function with given fields:
func (_m *MockWallet) Seed() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// SeedPassphrase provides a mock function with given fields:
func (_m *MockWallet) SeedPassphrase() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Serialize provides a mock function with given fields:
func (_m *MockWallet) Serialize() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
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

// SetBip44Coin provides a mock function with given fields: ct
func (_m *MockWallet) SetBip44Coin(ct bip44.CoinType) {
	_m.Called(ct)
}

// SetCoin provides a mock function with given fields: coinType
func (_m *MockWallet) SetCoin(coinType CoinType) {
	_m.Called(coinType)
}

// SetCryptoType provides a mock function with given fields: ct
func (_m *MockWallet) SetCryptoType(ct crypto.CryptoType) {
	_m.Called(ct)
}

// SetDecoder provides a mock function with given fields: d
func (_m *MockWallet) SetDecoder(d Decoder) {
	_m.Called(d)
}

// SetFilename provides a mock function with given fields: _a0
func (_m *MockWallet) SetFilename(_a0 string) {
	_m.Called(_a0)
}

// SetLabel provides a mock function with given fields: _a0
func (_m *MockWallet) SetLabel(_a0 string) {
	_m.Called(_a0)
}

// SetTemp provides a mock function with given fields: temp
func (_m *MockWallet) SetTemp(temp bool) {
	_m.Called(temp)
}

// SetTimestamp provides a mock function with given fields: _a0
func (_m *MockWallet) SetTimestamp(_a0 int64) {
	_m.Called(_a0)
}

// Timestamp provides a mock function with given fields:
func (_m *MockWallet) Timestamp() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Type provides a mock function with given fields:
func (_m *MockWallet) Type() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Unlock provides a mock function with given fields: password
func (_m *MockWallet) Unlock(password []byte) (Wallet, error) {
	ret := _m.Called(password)

	var r0 Wallet
	if rf, ok := ret.Get(0).(func([]byte) Wallet); ok {
		r0 = rf(password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Version provides a mock function with given fields:
func (_m *MockWallet) Version() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// XPub provides a mock function with given fields:
func (_m *MockWallet) XPub() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
