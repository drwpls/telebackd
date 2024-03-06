# telebackd

## Building

1. Creating workspace: `mkdir workspace && cd workspace`
2. Download SDK: [Openwrt image files](https://downloads.openwrt.org/releases/21.02.3/targets/ramips/mt7621)  
    File: `openwrt-sdk-21.02.3-ramips-mt7621_gcc-8.4.0_musl.Linux-x86_64.tar.xz`, sha256sum `163e71b91c66af24c61a924964f8deeb816babde3d554306e53fd82364f46632`
3. Extract image: `tar xvf ./openwrt-sdk-21.02.3-ramips-mt7621_gcc-8.4.0_musl.Linux-x86_64.tar.xz`  
   There will be a directory named `openwrt-sdk-21.02.3-ramips-mt7621_gcc-8.4.0_musl.Linux-x86_64`
4. Get to `package` folder and `cd openwrt-sdk-21.02.3-ramips-mt7621_gcc-8.4.0_musl.Linux-x86_64/package/`
5. Clone the project into telebackd-src folder: `git clone https://github.com/drwpls/telebackd.git telebackd-src`  
    The `package` folder structure:
```bash
    $ package tree -L 3
    .
    ├── kernel
    │   └── linux
    │       ├── files
    │       ├── Makefile
    │       └── modules
    ├── Makefile
    ├── telebackd-src
    │   ├── package
    │   │   └── telebackd
    │   └── README.md
    └── toolchain
        ├── glibc-files
        │   └── etc
        └── Makefile

    11 directories, 4 files
```
6. We need to move `telebackd` folder into `package` and remove other files (VCS,...)  
   `cp -r ./telebackd-src/package/telebackd . && rm ./telebackd-src -rf`
7. Edit `Makefile` to set proper Arch: `https://go.dev/wiki/GoMips`  
    I use AC2100 Router with `mipsel_24kc` CPU:
```bash
    export GOARCH=mipsle
    export GOMIPS=softfloat
```
8. Enable `telebackd` in `make` context: 
   Change dir to `workspace/openwrt-sdk-21.02.3-ramips-mt7621_gcc-8.4.0_musl.Linux-x86_64` then run `make menuconfig`  
   Navigate to `Utilities` and enable `telebackd`
9. Compile the package `make package/telebackd/compile V=s`, the compiled bundle will be in `workspace/openwrt-sdk-21.02.3-ramips-mt7621_gcc-8.4.0_musl.Linux-x86_64/bin/packages/mipsel_24kc/base/telebackd_1.0.0-1_mipsel_24kc.ipk`

## Installing
10. Upload the `ipk` file into server and install using `opkg install telebackd_1.0.0-1_mipsel_24kc.ipk`
11. Config the `admin_id` and `bot_token`:  
    `uci set telebackd.telebackd[0].admin_id=<your_admin_id>`  
    `uci set telebackd.telebackd[0].bot_token=<your:bottoken>`
12. Restart telebackd:  
    `service telebackd restart`
    