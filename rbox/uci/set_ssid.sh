#!/bin/sh

SSID="$@"

source /lib/functions.sh

do_wireless() {
	local wifi_iface="$1"
	local ssid="$2"

	uci set wireless."$wifi_iface".ssid="$ssid"
}

config_load wireless

config_foreach do_wireless wifi-iface "$SSID"

uci commit wireless

wifi reload

rm -rf /tmp/set_ssid.sh
