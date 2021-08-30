package rbox

import (
	"os/exec"
	"strings"
)

type RemoteBox struct {
	boxname string
}

func (b *RemoteBox) Scp(script string) error {
	scp := []string{
		"/home/rbox/uci/" + script,
		b.boxname + ":/tmp",
	}
	_, err := exec.Command("scp", scp...).Output()
	return err
}

func (b *RemoteBox) Ssh(script, param string) error {
	ssh := []string{
		b.boxname,
		"/tmp/" + script,
		param,
	}
	_, err := exec.Command("ssh", ssh...).Output()
	return err
}

func (b *RemoteBox) GetConfig() ([]byte, error) {
	cmd := exec.Command("ssh", b.boxname, "uci show")
	config, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (b *RemoteBox) GetSSIDs() ([]string, error) {
	SSIDs := []string{}

	commands := []string{
		"uci get wireless.@wifi-iface[0].ssid",
		"uci get wireless.@wifi-iface[1].ssid",
	}
	for _, command := range commands {
		cmd := exec.Command("ssh", b.boxname, command)
		SSID, err := cmd.Output()
		if err != nil {
			return nil, err
		}
		SSIDs = append(SSIDs, strings.TrimSpace(string(SSID)))
	}

	return SSIDs, nil
}

func (b *RemoteBox) SetSSIDs(SSID string) error {
	if err := b.Scp("set_ssid.sh"); err != nil {
		return err
	}

	return b.Ssh("set_ssid.sh", SSID)
}

func (b *RemoteBox) GetMACs() ([]string, error) {
	cmd := exec.Command("ssh", b.boxname, "uci get firewall.macs.entry")
	cmdMACs, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	MACs := strings.Split(strings.Trim(string(cmdMACs), "\n"), " ")
	return MACs, nil
}

func (b *RemoteBox) SetMACs(MACs string) error {
	if err := b.Scp("set_ipset_macs.sh"); err != nil {
		return err
	}

	return b.Ssh("set_ipset_macs.sh", MACs)
}

func NewRBox(boxname string) RemoteBox {
	return RemoteBox{
		boxname: boxname,
	}
}
