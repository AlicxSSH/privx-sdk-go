//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package licensemanager

// License license definition
type License struct {
	LicenseStatus           string   `json:"license_status,omitempty"`
	Version                 string   `json:"version,omitempty"`
	CreationDate            string   `json:"creation_date,omitempty"`
	ExpiryDate              string   `json:"expiry_date,omitempty"`
	LastRefreshedDate       string   `json:"last_refreshed_date,omitempty"`
	Customer                string   `json:"customer,omitempty"`
	SerialNumber            string   `json:"serial_number,omitempty"`
	Product                 string   `json:"product,omitempty"`
	LicenseCode             string   `json:"license_code,omitempty"`
	Status                  int      `json:"status,omitempty"`
	Message                 int      `json:"message,omitempty"`
	MaxHosts                int      `json:"max_hosts,omitempty"`
	MaxAuditedHosts         int      `json:"max_audited_hosts,omitempty"`
	MaxConcurrentSSHConns   int      `json:"max_concurrent_ssh_conns,omitempty"`
	MaxConcurrentRDPConns   int      `json:"max_concurrent_rdp_conns,omitempty"`
	MaxConcurrentHTTPSConns int      `json:"max_concurrent_https_conns,omitempty"`
	AnalyticsEnabled        bool     `json:"analytics_enabled,omitempty"`
	IsOffline               bool     `json:"isoffline,omitempty"`
	Optin                   bool     `json:"optin,omitempty"`
	Features                []string `json:"features,omitempty"`
}
