#!/bin/sh

MACS="$@"

uci del firewall.macs.entry

for MAC in $MACS
do
    uci add_list firewall.macs.entry="$MAC"
done

uci commit firewall

/etc/init.d/firewall restart

rm -rf /tmp/set_ipset_macs.sh
