linux
```bash
vcpkg install --triplet=x64-linux-release
export PKG_CONFIG_PATH="$PWD/vcpkg_installed/x64-linux-release/lib/pkgconfig"

```

macOS
```zsh
vcpkg install --triplet=arm64-osx-release
export PKG_CONFIG_PATH="$PWD/vcpkg_installed/arm64-osx-release/lib/pkgconfig"
```

//#cgo LDFLAGS: -static -lstdc++ -lm -lpthread
#cgo darwin LDFLAGS: -framework CoreFoundation