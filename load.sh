if [[ ! -e /Volumes/RPI-RP2 ]]; then
	echo "Device not found. Did you start it in BOOTLOADER mode?"
	exit 1
fi

cd $1
tinygo flash -target=pico ./
cd -
