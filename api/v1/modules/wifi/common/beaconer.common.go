package wifi_common

import (
	"errors"
	"log"
	"net"
	"time"

	"github.com/bettercap/bettercap/packets"
)

type Beaconer struct {
	ApName       string
	ApMac        string
	ApChannel    int
	ApEncryption bool
}

func (mod *WiFiModule) ApSettings(beaconer Beaconer) error {

	bssid, err := net.ParseMAC(beaconer.ApMac)
	if err != nil {

		return err
	}

	mod.apConfig.SSID = beaconer.ApName
	mod.apConfig.BSSID = bssid
	mod.apConfig.Channel = beaconer.ApChannel
	mod.apConfig.Encryption = beaconer.ApEncryption

	return nil
}

func (mod *WiFiModule) StartAp() error {

	if mod.apRunning {
		return errors.New(mod.apConfig.SSID + " is running")
	}

	go func() {

		mod.apRunning = true
		defer func() {
			mod.apRunning = false
		}()

		for seqn := uint16(0); mod.Running(); seqn++ {
			mod.writes.Add(1)
			defer mod.writes.Done()

			select {
			case <-BeaconerChanel:
				return
			default:
			}

			if err, pkt := packets.NewDot11Beacon(mod.apConfig, seqn); err != nil {
				log.Printf("could not create beacon packet: %s\n", err)
			} else {
				mod.injectPacket(pkt)
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()

	return nil
}
