#!/bin/bash - 
set -o nounset                              # Treat unset variables as an error

export PULUMI_SKIP_UPDATE_CHECK=true

vagrant up

machines=$(vagrant status --machine-readable | grep metadata | cut -f 2 -d ',')
for a in $machines; do
  ip=$(vagrant ssh $a --command "ip -4 -o addr show eth1 | awk '{print \$4}' | cut -d '/' -f 1" | tr -d "\r")
  echo "Set IP ${ip} for machine ${a} into pulumi config"
  pulumi config set ${a}IP $ip
done


