package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

const configuredHost = "127.0.0.1:6420"

var allAPISetsEnabled = map[string]struct{}{
	EndpointsRead:                  struct{}{},
	EndpointsStatus:                struct{}{},
	EndpointsWallet:                struct{}{},
	EndpointsWalletSeed:            struct{}{},
	EndpointsDeprecatedWalletSpend: struct{}{},
}

func defaultMuxConfig() muxConfig {
	return muxConfig{
		host:           configuredHost,
		appLoc:         ".",
		disableCSP:     true,
		enabledAPISets: allAPISetsEnabled,
	}
}

var endpoints = []string{
	"/address_uxouts",
	"/addresscount",
	"/balance",
	"/block",
	"/blockchain/metadata",
	"/blockchain/progress",
	"/blocks",
	"/coinSupply",
	"/explorer/address",
	"/health",
	"/injectTransaction",
	"/last_blocks",
	"/version",
	"/network/connection",
	"/network/connections",
	"/network/connections/exchange",
	"/network/connections/trust",
	"/network/defaultConnections",
	"/outputs",
	"/pendingTxs",
	"/rawtx",
	"/richlist",
	"/resendUnconfirmedTxns",
	"/transaction",
	"/transactions",
	"/uxout",
	"/wallet",
	"/wallet/balance",
	"/wallet/create",
	"/wallet/newAddress",
	"/wallet/newSeed",
	"/wallet/seed",
	"/wallet/spend",
	"/wallet/transaction",
	"/wallet/transactions",
	"/wallet/unload",
	"/wallet/update",
	"/wallets",
	"/wallets/folderName",
	"/webrpc",

	"/api/v1/address_uxouts",
	"/api/v1/addresscount",
	"/api/v1/balance",
	"/api/v1/block",
	"/api/v1/blockchain/metadata",
	"/api/v1/blockchain/progress",
	"/api/v1/blocks",
	"/api/v1/coinSupply",
	"/api/v1/explorer/address",
	"/api/v1/health",
	"/api/v1/injectTransaction",
	"/api/v1/last_blocks",
	"/api/v1/version",
	"/api/v1/network/connection",
	"/api/v1/network/connections",
	"/api/v1/network/connections/exchange",
	"/api/v1/network/connections/trust",
	"/api/v1/network/defaultConnections",
	"/api/v1/outputs",
	"/api/v1/pendingTxs",
	"/api/v1/rawtx",
	"/api/v1/richlist",
	"/api/v1/resendUnconfirmedTxns",
	"/api/v1/transaction",
	"/api/v1/transactions",
	"/api/v1/uxout",
	"/api/v1/wallet",
	"/api/v1/wallet/balance",
	"/api/v1/wallet/create",
	"/api/v1/wallet/newAddress",
	"/api/v1/wallet/newSeed",
	"/api/v1/wallet/seed",
	"/api/v1/wallet/spend",
	"/api/v1/wallet/transaction",
	"/api/v1/wallet/transactions",
	"/api/v1/wallet/unload",
	"/api/v1/wallet/update",
	"/api/v1/wallets",
	"/api/v1/wallets/folderName",
	"/api/v1/webrpc",

	"/api/v2/transaction/verify",
	"/api/v2/address/verify",
}

// TestEnableGUI tests enable gui option, EnableGUI isn't part of Gateway API,
// we can't control the output by mocking the Gateway like other tests. Instead,
// we create a full webserver for each test case.
func TestEnableGUI(t *testing.T) {
	tt := []struct {
		name       string
		enableGUI  bool
		endpoint   string
		appLoc     string
		expectCode int
		expectBody string
	}{
		{
			name:       "disable gui GET /",
			enableGUI:  false,
			endpoint:   "/",
			appLoc:     "",
			expectCode: http.StatusNotFound,
			expectBody: "404 Not Found\n",
		},
		{
			name:       "disable gui GET /invalid-path",
			enableGUI:  false,
			endpoint:   "/invalid-path",
			appLoc:     "",
			expectCode: http.StatusNotFound,
			expectBody: "404 Not Found\n",
		},
		{
			name:       "enable gui GET /",
			enableGUI:  true,
			endpoint:   "/",
			appLoc:     "../gui/static",
			expectCode: http.StatusOK,
			expectBody: "",
		},
		{
			name:       "enable gui GET /invalid-path",
			enableGUI:  true,
			endpoint:   "/invalid-path",
			appLoc:     "../gui/static",
			expectCode: http.StatusNotFound,
			expectBody: "404 Not Found\n",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, tc.endpoint, nil)
			require.NoError(t, err)

			gateway := &MockGatewayer{}

			rr := httptest.NewRecorder()
			handler := newServerMux(muxConfig{
				host:       configuredHost,
				appLoc:     tc.appLoc,
				disableCSP: true,
			}, gateway, &CSRFStore{}, nil)
			handler.ServeHTTP(rr, req)

			c := Config{
				EnableGUI:   tc.enableGUI,
				DisableCSRF: true,
				StaticDir:   tc.appLoc,
			}

			host := "127.0.0.1:6423"
			s, err := Create(host, c, gateway)
			require.NoError(t, err)

			wg := sync.WaitGroup{}
			wg.Add(1)
			go func() {
				defer wg.Done()
				err := s.Serve()
				if err != nil && err.Error() != fmt.Sprintf("accept tcp %s: use of closed network connection", host) {
					require.NoError(t, err)
				}
			}()

			defer func() {
				s.listener.Close()
				wg.Wait()
			}()

			url := fmt.Sprintf("http://%s/%s", host, tc.endpoint)
			rsp, err := http.Get(url)
			require.NoError(t, err)

			defer rsp.Body.Close()
			require.Equal(t, tc.expectCode, rsp.StatusCode)

			body, err := ioutil.ReadAll(rr.Body)
			require.NoError(t, err)

			if rsp.StatusCode != http.StatusOK {
				require.Equal(t, tc.expectBody, string(body))
			}
		})
	}
}

func TestAPISetDisabled(t *testing.T) {
	for _, e := range append(endpoints, []string{"/csrf", "/api/v1/csrf"}...) {
		switch e {
		case "/webrpc", "/api/v1/webrpc":
			continue
		}

		t.Run(e, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, e, nil)
			require.NoError(t, err)

			cfg := defaultMuxConfig()
			cfg.enableUnversionedAPI = true
			cfg.enableJSON20RPC = false
			cfg.enabledAPISets = map[string]struct{}{} // disable all API sets

			handler := newServerMux(cfg, &MockGatewayer{}, &CSRFStore{
				Enabled: true,
			}, nil)

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			switch e {
			case "/csrf", "/api/v1/csrf", "/version", "/api/v1/version": // always enabled
				require.Equal(t, http.StatusOK, rr.Code)
			default:
				require.Equal(t, http.StatusForbidden, rr.Code)
				require.Equal(t, "403 Forbidden - Endpoint is disabled", strings.TrimSpace(rr.Body.String()))
			}
		})
	}
}
