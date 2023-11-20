package wifi_test

import (
	"errors"
	myWifi "example_mock/internal/wifi"
	"fmt"
	"net"
	"testing"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

//go:generate mockery --all --testonly --quiet --outpkg wifi_test --output .

type rowTestSysInfoGetAddresses struct {
	addrs       []string
	errExpected error
}

type rowTestSysInfoGetNames struct {
	names       []string
	errExpected error
}

var testTableGetAddresses = []rowTestSysInfoGetAddresses{
	{
		addrs: []string{"00:11:22:33:44:55", "aa:bb:cc:dd:ee:ff"},
	},
	{
		errExpected: errors.New("ExpectedError"),
	},
}

var testTableGetNames = []rowTestSysInfoGetNames{
	{
		names: []string{"name1", "name2"},
	},
	{
		errExpected: errors.New("ExpectedError"),
	},
}

func TestGetAddresses(t *testing.T) {
	mockWifi := NewWiFi(t)
	wifiService := myWifi.Service{WiFi: mockWifi}
	for i, row := range testTableGetAddresses {
		mockWifi.On("Interfaces").Unset()
		mockWifi.On("Interfaces").Return(mockIfacesAdresses(row.addrs),
			row.errExpected)
		actualAddrs, err := wifiService.GetAddresses()
		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error:  %w, actual error: %w", i,
				row.errExpected, err)
			continue
		}
		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, parseMACs(row.addrs), actualAddrs, "row: %d, expected addrs: %s, actual addrs: %s", i,
			parseMACs(row.addrs), actualAddrs)
	}
}

func TestGetNames(t *testing.T) {
	mockWifi := NewWiFi(t)
	wifiService := myWifi.Service{WiFi: mockWifi}
	for i, row := range testTableGetNames {
		mockWifi.On("Interfaces").Unset()
		mockWifi.On("Interfaces").Return(mockIfacesNames(row.names),
			row.errExpected)
		actualNames, err := wifiService.GetNames()
		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error:  %w, actual error: %w", i,
				row.errExpected, err)
			continue
		}
		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.names, actualNames, "row: %d, expected addrs: %s, actual addrs: %s", i,
			parseMACs(row.names), actualNames)
	}
}

func mockIfacesAdresses(addrs []string) []*wifi.Interface {
	var interfaces []*wifi.Interface
	for i, addrStr := range addrs {
		hwAddr := parseMAC(addrStr)
		if hwAddr == nil {
			continue
		}
		iface := &wifi.Interface{
			Index:        i + 1,
			Name:         fmt.Sprintf("eth%d", i+1),
			HardwareAddr: hwAddr,
			PHY:          1,
			Device:       1,
			Type:         wifi.InterfaceTypeAPVLAN,
			Frequency:    0,
		}
		interfaces = append(interfaces, iface)
	}
	return interfaces
}

func mockIfacesNames(names []string) []*wifi.Interface {
	var interfaces []*wifi.Interface
	for i, nameStr := range names {
		iface := &wifi.Interface{
			Index:        i + 1,
			Name:         nameStr,
			HardwareAddr: nil,
			PHY:          1,
			Device:       1,
			Type:         wifi.InterfaceTypeAPVLAN,
			Frequency:    0,
		}
		interfaces = append(interfaces, iface)
	}
	return interfaces
}

func parseMACs(macStr []string) []net.HardwareAddr {
	var addrs []net.HardwareAddr
	for _, addr := range macStr {
		addrs = append(addrs, parseMAC(addr))
	}
	return addrs
}

func parseMAC(macStr string) net.HardwareAddr {
	hwAddr, err := net.ParseMAC(macStr)
	if err != nil {
		return nil
	}
	return hwAddr
}
