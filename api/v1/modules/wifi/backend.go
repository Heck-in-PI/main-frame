package wifi

import (
	"encoding/json"
	"io"
	"log"
	wifi_common "mf-backend/api/v1/modules/wifi/common"
	v1_common "mf-backend/api/v1/v1Common"
	"net"
	"strings"

	"net/http"

	"github.com/bettercap/bettercap/network"
	"github.com/gorilla/mux"
	wifi "github.com/mdlayher/wifi"
	goWireless "github.com/theojulienne/go-wireless"
)

// interfaces handler
func interfacesHandler(resp http.ResponseWriter, req *http.Request) {

	defer req.Body.Close()

	if req.Method == "GET" {

		wifiClient, err := wifi.New()
		if err != nil {

			errorMessage := v1_common.ErrorMessage{
				Error: err.Error(),
			}

			v1_common.JsonResponceHandler(resp, http.StatusBadRequest, errorMessage)

			return
		}

		ifaces, err := wifiClient.Interfaces()
		if err != nil {

			errorMessage := v1_common.ErrorMessage{
				Error: err.Error(),
			}

			v1_common.JsonResponceHandler(resp, http.StatusBadRequest, errorMessage)

			return
		}

		var wirelessInterfaces []wifi_common.WirelessInterface
		for _, iface := range ifaces {
			wirelessInterface := wifi_common.WirelessInterface{
				Index:        iface.Index,
				Name:         iface.Name,
				HardwareAddr: iface.HardwareAddr.String(),
				PHY:          iface.PHY,
				Device:       iface.Device,
				Type:         iface.Type,
				Frequency:    iface.Frequency,
			}

			wirelessInterfaces = append(wirelessInterfaces, wirelessInterface)
		}

		v1_common.JsonResponceHandler(resp, http.StatusOK, wirelessInterfaces)
	} else {

		errorMessage := v1_common.ErrorMessage{
			Error: "Invalid Request",
		}

		v1_common.JsonResponceHandler(resp, http.StatusBadRequest, errorMessage)
	}
}

// scan access point handler
func scanApHandler(resp http.ResponseWriter, req *http.Request) {

	defer req.Body.Close()

	if req.Method == "GET" {

		muxVars := mux.Vars(req)
		interfaceName := muxVars["interfaceName"]
		if interfaceName == "" {
			errorMessage := v1_common.ErrorMessage{
				Error: "interface name must be specified in path",
			}

			v1_common.JsonResponceHandler(resp, http.StatusBadRequest, errorMessage)

			return
		}

		client, err := goWireless.NewClient(interfaceName)
		if err != nil {

			errorMessage := v1_common.ErrorMessage{
				Error: err.Error(),
			}

			v1_common.JsonResponceHandler(resp, http.StatusBadRequest, errorMessage)

			return
		}

		log.Println(client)
		defer client.Close()

		aps, err := client.Scan()
		if err != nil {

			errorMessage := v1_common.ErrorMessage{
				Error: err.Error(),
			}

			v1_common.JsonResponceHandler(resp, http.StatusBadRequest, errorMessage)

			return
		}

		v1_common.JsonResponceHandler(resp, http.StatusOK, aps)

	} else {
		resp.Write([]byte("{\"err\":\"invalid request\"}"))
	}
}

// death handler
func deauthHandler(resp http.ResponseWriter, req *http.Request) {

	defer req.Body.Close()

	if req.Method == "POST" {

		var deauther wifi_common.Deauther

		body, _ := io.ReadAll(req.Body)
		err := json.Unmarshal(body, &deauther)
		if err != nil {

			errorMessage := v1_common.ErrorMessage{
				Error: err.Error(),
			}

			v1_common.JsonResponceHandler(resp, http.StatusBadRequest, errorMessage)

			return
		}

		wifiModule, err := wifi_common.NewWiFiModule(deauther.InterfaceName)
		if err != nil {

			errorMessage := v1_common.ErrorMessage{
				Error: err.Error(),
			}

			v1_common.JsonResponceHandler(resp, http.StatusBadRequest, errorMessage)

			return
		}

		bssid, err := net.ParseMAC(deauther.ApMac)
		if err != nil {

			errorMessage := v1_common.ErrorMessage{
				Error: err.Error(),
			}

			v1_common.JsonResponceHandler(resp, http.StatusBadRequest, errorMessage)

			return
		}

		if deauther.ClientMac == "" || strings.ToLower(deauther.ClientMac) == "all" {

			clients := wifi_common.FindApClients(bssid, wifiModule.Handle)
			if len(clients) == 0 {
				errorMessage := v1_common.ErrorMessage{
					Error: "Can't find client related to ap " + deauther.ApMac,
				}

				v1_common.JsonResponceHandler(resp, http.StatusBadRequest, errorMessage)

				return
			}

			// filter out safe clients
			if len(deauther.SafeClients) != 0 {

				var filteredClient []*network.Station
				for _, client := range clients {
					for _, safeClient := range deauther.SafeClients {
						if client.Endpoint.HwAddress == safeClient {
							continue
						}

						filteredClient = append(filteredClient, client)
					}
				}

				clients = filteredClient
				if len(clients) == 0 {
					errorMessage := v1_common.ErrorMessage{
						Error: "Can't find other client related to ap " + deauther.ApMac,
					}

					v1_common.JsonResponceHandler(resp, http.StatusBadRequest, errorMessage)

					return
				}
			}

			for _, client := range clients {
				log.Println("kicking out from:", bssid, ", client: ", client.Endpoint.HW)
				wifiModule.SendDeauthPacket(bssid, client.Endpoint.HW)
			}
		} else {

			client, err := net.ParseMAC(deauther.ClientMac)
			if err != nil {

				errorMessage := v1_common.ErrorMessage{
					Error: err.Error(),
				}

				v1_common.JsonResponceHandler(resp, http.StatusBadRequest, errorMessage)

				return
			}

			log.Println("kicking out from:", bssid, ", client: ", client)
			wifiModule.SendDeauthPacket(bssid, client)

		}
		resp.WriteHeader(http.StatusOK)
	} else {
		resp.Write([]byte("{\"err\":\"invalid request\"}"))
	}
}
