#!/usr/bin/env bash

trap refuse EXIT

timeout="${TIMEOUT:-1s}"

refuse() {
	while :; do
		echo "[e] $(date): I won't exit!!!"
		sleep $timeout
	done
}

while :; do
	echo "[i] $(date): Normal operation...."
	sleep $timeout
done

