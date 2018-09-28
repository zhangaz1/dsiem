package event

import (
	"dsiem/internal/dsiem/pkg/asset"
	"encoding/json"
)

// NormalizedEvent represents data received from logstash
type NormalizedEvent struct {
	ConnID       uint64 `json:"conn_id,omitempty"`
	EventID      string `json:"event_id"`
	Timestamp    string `json:"timestamp"`
	Sensor       string `json:"sensor"`
	PluginID     int    `json:"plugin_id"`
	PluginSID    int    `json:"plugin_sid"`
	Product      string `json:"product"`
	Category     string `json:"category"`
	SubCategory  string `json:"subcategory,omitempty"`
	SrcIP        string `json:"src_ip"`
	SrcPort      int    `json:"src_port"`
	DstIP        string `json:"dst_ip"`
	DstPort      int    `json:"dst_port"`
	Protocol     string `json:"protocol"`
	CustomData1  string `json:"custom_data1,omitempty"`
	CustomLabel1 string `json:"custom_label1,omitempty"`
	CustomData2  string `json:"custom_data2,omitempty"`
	CustomLabel2 string `json:"custom_label2,omitempty"`
	CustomData3  string `json:"custom_data3,omitempty"`
	CustomLabel3 string `json:"custom_label3,omitempty"`
}

// Valid check if event contains valid content for required fields
func (e *NormalizedEvent) Valid() bool {
	// fmt.Println(e.Timestamp, ":", e.Sensor, ":", e.EventID, ":", e.SrcIP, ":", e.DstIP, ":", e.PluginID, ":", e.PluginSID)
	if e.Timestamp == "" || e.Sensor == "" || e.EventID == "" || e.SrcIP == "" || e.DstIP == "" {
		return false
	}

	if e.PluginID == 0 || e.PluginSID == 0 {
		if e.Product == "" || e.Category == "" {
			return false
		}
	}
	return true
}

// FromBytes initialize NormalizedEvent
func (e *NormalizedEvent) FromBytes(b []byte) error {
	err := json.Unmarshal(b, &e)
	return err
}

// SrcIPInHomeNet check if event SrcIP is is HOME_NET
func (e *NormalizedEvent) SrcIPInHomeNet() bool {
	res, _ := asset.IsInHomeNet(e.SrcIP)
	return res
}

// DstIPInHomeNet check if event DstIP is is HOME_NET
func (e *NormalizedEvent) DstIPInHomeNet() bool {
	res, _ := asset.IsInHomeNet(e.DstIP)
	return res
}
