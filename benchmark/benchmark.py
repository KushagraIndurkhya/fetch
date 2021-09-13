URL="https://saimei.ftp.acc.umu.se/debian-cd/current/amd64/iso-cd/debian-11.0.0-amd64-netinst.iso"
REPEATS=10

fetch_cmd=f'''fetch download {URL} debian.iso --threads=20 > /dev/null'''
curl_cmd=f"curl --silent -o deb.iso {URL} > /dev/null"
wget_cmd=f"wget -q {URL} > /dev/null"

import os
os.system(
f'''python3 -c '
from time_me import benchmark
print("fetch",benchmark("{fetch_cmd}",{REPEATS}))
' &
python3 -c '
from time_me import benchmark
print("curl",benchmark("{curl_cmd}",{REPEATS}))
' &
python3 -c '
from time_me import benchmark
print("wget",benchmark("{wget_cmd}",{REPEATS}))
' &''')