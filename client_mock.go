// Code generated by MockGen. DO NOT EDIT.
// Source: client.go
//
// Generated by this command:
//
//	mockgen -source=client.go -destination=client_mock.go -package=smartid
//

// Package smartid is a generated GoMock package.
package smartid

import (
	context "context"
	tls "crypto/tls"
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
	isgomock struct{}
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// CreateSession mocks base method.
func (m *MockClient) CreateSession(ctx context.Context, nationalIdentityNumber string) (*Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", ctx, nationalIdentityNumber)
	ret0, _ := ret[0].(*Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockClientMockRecorder) CreateSession(ctx, nationalIdentityNumber any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockClient)(nil).CreateSession), ctx, nationalIdentityNumber)
}

// FetchSession mocks base method.
func (m *MockClient) FetchSession(ctx context.Context, sessionId string) (*Person, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchSession", ctx, sessionId)
	ret0, _ := ret[0].(*Person)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchSession indicates an expected call of FetchSession.
func (mr *MockClientMockRecorder) FetchSession(ctx, sessionId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchSession", reflect.TypeOf((*MockClient)(nil).FetchSession), ctx, sessionId)
}

// Validate mocks base method.
func (m *MockClient) Validate() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockClientMockRecorder) Validate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockClient)(nil).Validate))
}

// WithCertificateLevel mocks base method.
func (m *MockClient) WithCertificateLevel(level string) Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithCertificateLevel", level)
	ret0, _ := ret[0].(Client)
	return ret0
}

// WithCertificateLevel indicates an expected call of WithCertificateLevel.
func (mr *MockClientMockRecorder) WithCertificateLevel(level any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithCertificateLevel", reflect.TypeOf((*MockClient)(nil).WithCertificateLevel), level)
}

// WithDisplayText200 mocks base method.
func (m *MockClient) WithDisplayText200(text string) Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithDisplayText200", text)
	ret0, _ := ret[0].(Client)
	return ret0
}

// WithDisplayText200 indicates an expected call of WithDisplayText200.
func (mr *MockClientMockRecorder) WithDisplayText200(text any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithDisplayText200", reflect.TypeOf((*MockClient)(nil).WithDisplayText200), text)
}

// WithDisplayText60 mocks base method.
func (m *MockClient) WithDisplayText60(text string) Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithDisplayText60", text)
	ret0, _ := ret[0].(Client)
	return ret0
}

// WithDisplayText60 indicates an expected call of WithDisplayText60.
func (mr *MockClientMockRecorder) WithDisplayText60(text any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithDisplayText60", reflect.TypeOf((*MockClient)(nil).WithDisplayText60), text)
}

// WithHashType mocks base method.
func (m *MockClient) WithHashType(hashType string) Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithHashType", hashType)
	ret0, _ := ret[0].(Client)
	return ret0
}

// WithHashType indicates an expected call of WithHashType.
func (mr *MockClientMockRecorder) WithHashType(hashType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithHashType", reflect.TypeOf((*MockClient)(nil).WithHashType), hashType)
}

// WithInteractionType mocks base method.
func (m *MockClient) WithInteractionType(interactionType string) Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithInteractionType", interactionType)
	ret0, _ := ret[0].(Client)
	return ret0
}

// WithInteractionType indicates an expected call of WithInteractionType.
func (mr *MockClientMockRecorder) WithInteractionType(interactionType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithInteractionType", reflect.TypeOf((*MockClient)(nil).WithInteractionType), interactionType)
}

// WithNonce mocks base method.
func (m *MockClient) WithNonce(nonce string) Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithNonce", nonce)
	ret0, _ := ret[0].(Client)
	return ret0
}

// WithNonce indicates an expected call of WithNonce.
func (mr *MockClientMockRecorder) WithNonce(nonce any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithNonce", reflect.TypeOf((*MockClient)(nil).WithNonce), nonce)
}

// WithRelyingPartyName mocks base method.
func (m *MockClient) WithRelyingPartyName(name string) Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithRelyingPartyName", name)
	ret0, _ := ret[0].(Client)
	return ret0
}

// WithRelyingPartyName indicates an expected call of WithRelyingPartyName.
func (mr *MockClientMockRecorder) WithRelyingPartyName(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithRelyingPartyName", reflect.TypeOf((*MockClient)(nil).WithRelyingPartyName), name)
}

// WithRelyingPartyUUID mocks base method.
func (m *MockClient) WithRelyingPartyUUID(id string) Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithRelyingPartyUUID", id)
	ret0, _ := ret[0].(Client)
	return ret0
}

// WithRelyingPartyUUID indicates an expected call of WithRelyingPartyUUID.
func (mr *MockClientMockRecorder) WithRelyingPartyUUID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithRelyingPartyUUID", reflect.TypeOf((*MockClient)(nil).WithRelyingPartyUUID), id)
}

// WithTLSConfig mocks base method.
func (m *MockClient) WithTLSConfig(tlsConfig *tls.Config) Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTLSConfig", tlsConfig)
	ret0, _ := ret[0].(Client)
	return ret0
}

// WithTLSConfig indicates an expected call of WithTLSConfig.
func (mr *MockClientMockRecorder) WithTLSConfig(tlsConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTLSConfig", reflect.TypeOf((*MockClient)(nil).WithTLSConfig), tlsConfig)
}

// WithTimeout mocks base method.
func (m *MockClient) WithTimeout(timeout time.Duration) Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTimeout", timeout)
	ret0, _ := ret[0].(Client)
	return ret0
}

// WithTimeout indicates an expected call of WithTimeout.
func (mr *MockClientMockRecorder) WithTimeout(timeout any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTimeout", reflect.TypeOf((*MockClient)(nil).WithTimeout), timeout)
}

// WithURL mocks base method.
func (m *MockClient) WithURL(url string) Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithURL", url)
	ret0, _ := ret[0].(Client)
	return ret0
}

// WithURL indicates an expected call of WithURL.
func (mr *MockClientMockRecorder) WithURL(url any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithURL", reflect.TypeOf((*MockClient)(nil).WithURL), url)
}
