package ddnsman

import (
	"context"
	"errors"
	"log/slog"
	"net/netip"
	"testing"

	"github.com/libdns/libdns"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockProvider struct {
	mock.Mock
}

func (m *MockProvider) GetRecords(ctx context.Context, zone string) ([]libdns.Record, error) {
	args := m.Called(ctx, zone)
	return args.Get(0).([]libdns.Record), args.Error(1)
}

func (m *MockProvider) SetRecords(ctx context.Context, zone string, recs []libdns.Record) ([]libdns.Record, error) {
	args := m.Called(ctx, zone, recs)
	return args.Get(0).([]libdns.Record), args.Error(1)
}

func TestNew(t *testing.T) {
	t.Parallel()
	upd, err := New(&Configuration{})
	assert.NotNil(t, upd)
	assert.NoError(t, err)
}

func TestUpdater_checkRecord(t *testing.T) {
	t.Parallel()
	slog.SetDefault(slog.New(slog.DiscardHandler))
	tt := []struct {
		name        string
		externalIP  netip.Addr
		setting     Setting
		expectedErr string
		expectFunc  func(m *MockProvider)
	}{
		{
			name:       "success",
			externalIP: netip.MustParseAddr("10.0.0.2"),
			setting: Setting{
				Domain:  "example.com.",
				Records: []string{"sub1", "sub2"},
			},
			expectFunc: func(m *MockProvider) {
				m.On("GetRecords", context.Background(), "example.com.").Return(
					[]libdns.Record{
						libdns.Address{Name: "sub1.example.com.", IP: netip.MustParseAddr("10.0.0.2")},
						libdns.Address{Name: "sub2.example.com.", IP: netip.MustParseAddr("10.0.0.3")},
					}, nil)
				m.On("SetRecords", context.Background(), "example.com.",
					[]libdns.Record{
						libdns.Address{Name: "sub2.example.com.", IP: netip.MustParseAddr("10.0.0.2")},
					}).Return([]libdns.Record{}, nil)
			},
		},
		{
			name:       "success noop",
			externalIP: netip.MustParseAddr("10.0.0.2"),
			setting: Setting{
				Domain:  "example.com.",
				Records: []string{"sub1", "sub2"},
			},
			expectFunc: func(m *MockProvider) {
				m.On("GetRecords", context.Background(), "example.com.").Return(
					[]libdns.Record{
						libdns.Address{Name: "sub1.example.com.", IP: netip.MustParseAddr("10.0.0.2")},
						libdns.Address{Name: "sub2.example.com.", IP: netip.MustParseAddr("10.0.0.2")},
					}, nil)
			},
		},
		{
			name:       "error get records",
			externalIP: netip.MustParseAddr("10.0.0.2"),
			setting: Setting{
				Domain:  "example.com.",
				Records: []string{"sub1", "sub2"},
			},
			expectedErr: `getting records for zone "example.com.": some error`,
			expectFunc: func(m *MockProvider) {
				m.On("GetRecords", context.Background(), "example.com.").Return([]libdns.Record{}, errors.New("some error"))
			},
		},
		{
			name:       "error set records",
			externalIP: netip.MustParseAddr("10.0.0.2"),
			setting: Setting{
				Domain:  "example.com.",
				Records: []string{"sub1", "sub2"},
			},
			expectedErr: `setting records: some error`,
			expectFunc: func(m *MockProvider) {
				m.On("GetRecords", context.Background(), "example.com.").Return(
					[]libdns.Record{
						libdns.Address{Name: "sub1.example.com.", IP: netip.MustParseAddr("10.0.0.2")},
						libdns.Address{Name: "sub2.example.com.", IP: netip.MustParseAddr("10.0.0.3")},
					}, nil)
				m.On("SetRecords", context.Background(), "example.com.",
					[]libdns.Record{
						libdns.Address{Name: "sub2.example.com.", IP: netip.MustParseAddr("10.0.0.2")},
					}).Return([]libdns.Record{}, errors.New("some error"))
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			m := new(MockProvider)
			defer m.AssertExpectations(t)
			tc.expectFunc(m)
			tc.setting.provider = m
			u, err := New(&Configuration{})
			require.NoError(t, err)
			assertError(t,
				u.checkRecord(context.Background(), tc.externalIP, tc.setting),
				tc.expectedErr,
			)
		})
	}
}
