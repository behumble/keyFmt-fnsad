# (비공식) fnsad용 keyFmt CLI

## 개요

- `fnsad` 를 포함한 Cosmos-SDK의 node binary는 unarmored hex 형식의 private key를 import하는 기능이 없다
- DOSI Vault의 `개인 키 보기` 기능은 unarmored hex를 노출한다
- "unarmored hex 형식의 private key" -> `fnsad`가 import 가능한 private key export 형식으로 변환툴 제공한다

## 가정

- 샘플용 개인키(private key) : `0000000000000000000000000000000000000000000000000000000000000001`
  - **주의: 이 private key는 공개된 값(`0x01`)이며 에 실제 자금을 넣으면 안됨**
- [go](https://go.dev/) 1.2 이상 설치
- `fnsad`는 설치되어 있다고 가정

## 빌드 & 실행

```bash
$ go build  # ./keyfmt 실행파일 생성
go: downloading github.com/Finschia/finschia-sdk v0.47.2
go: downloading github.com/regen-network/protobuf v1.3.3-alpha.regen.1
go: downloading github.com/Finschia/ostracon v1.1.0
go: downloading github.com/hdevalence/ed25519consensus v0.1.0
go: downloading github.com/tendermint/crypto v0.0.0-20191022145703-50d29ede1e15
go: downloading github.com/bgentry/speakeasy v0.1.0
go: downloading github.com/tendermint/go-amino v0.16.0
go: downloading github.com/mattn/go-isatty v0.0.17
go: downloading github.com/btcsuite/btcd v0.22.1
go: downloading gopkg.in/yaml.v2 v2.4.0
go: downloading golang.org/x/crypto v0.8.0
go: downloading github.com/tendermint/tendermint v0.34.20
go: downloading github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d
go: downloading github.com/mailru/easyjson v0.7.7
go: downloading github.com/pkg/errors v0.9.1
go: downloading github.com/tendermint/tm-db v0.6.7
go: downloading golang.org/x/exp v0.0.0-20220722155223-a9213eeb770e
go: downloading filippo.io/edwards25519 v1.0.0
go: downloading golang.org/x/sys v0.7.0
go: downloading github.com/ChainSafe/go-schnorrkel v0.0.0-20200405005733-88cbf1b4c40d
go: downloading github.com/confio/ics23/go v0.9.0
go: downloading github.com/cosmos/btcutil v1.0.5
go: downloading github.com/spf13/cobra v1.7.0
go: downloading golang.org/x/net v0.9.0
go: downloading google.golang.org/grpc v1.54.0
go: downloading github.com/go-kit/log v0.2.1
go: downloading github.com/go-logfmt/logfmt v0.6.0
go: downloading github.com/rs/zerolog v1.29.1
go: downloading gopkg.in/natefinch/lumberjack.v2 v2.2.1
go: downloading github.com/josharian/intern v1.0.0
go: downloading github.com/davecgh/go-spew v1.1.1
go: downloading github.com/r2ishiguro/vrf v0.0.0-20180716233122-192de52975eb
go: downloading github.com/armon/go-metrics v0.4.1
go: downloading github.com/prometheus/client_golang v1.15.0
go: downloading github.com/prometheus/common v0.42.0
go: downloading github.com/google/btree v1.0.0
go: downloading github.com/syndtr/goleveldb v1.0.1-0.20200815110645-5c35d600f0ca
go: downloading github.com/spf13/viper v1.15.0
go: downloading github.com/spf13/pflag v1.0.5
go: downloading github.com/cosmos/go-bip39 v1.0.0
go: downloading github.com/gtank/merlin v0.1.1
go: downloading github.com/gtank/ristretto255 v0.1.2
go: downloading google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f
go: downloading github.com/go-kit/kit v0.12.0
go: downloading github.com/mattn/go-colorable v0.1.13
go: downloading github.com/hashicorp/go-immutable-radix v1.3.1
go: downloading github.com/golang/protobuf v1.5.3
go: downloading github.com/matttproud/golang_protobuf_extensions v1.0.4
go: downloading github.com/prometheus/client_model v0.3.0
go: downloading github.com/fsnotify/fsnotify v1.6.0
go: downloading github.com/mitchellh/mapstructure v1.5.0
go: downloading github.com/spf13/afero v1.9.3
go: downloading github.com/spf13/cast v1.5.0
go: downloading github.com/spf13/jwalterweatherman v1.1.0
go: downloading github.com/beorn7/perks v1.0.1
go: downloading github.com/cespare/xxhash/v2 v2.2.0
go: downloading github.com/prometheus/procfs v0.9.0
go: downloading github.com/cespare/xxhash v1.1.0
go: downloading google.golang.org/protobuf v1.30.0
go: downloading github.com/libp2p/go-buffer-pool v0.1.0
go: downloading github.com/mimoo/StrobeGo v0.0.0-20181016162300-f8f6d4d2b643
go: downloading github.com/subosito/gotenv v1.4.2
go: downloading github.com/hashicorp/hcl v1.0.0
go: downloading gopkg.in/ini.v1 v1.67.0
go: downloading github.com/magiconair/properties v1.8.7
go: downloading github.com/pelletier/go-toml/v2 v2.0.6
go: downloading golang.org/x/text v0.9.0
go: downloading gopkg.in/yaml.v3 v3.0.1
go: downloading github.com/golang/snappy v0.0.4

$ fnsad keys list   # 현재 keyring에 관리중인 key가 없다
[]

$ ./keyfmt > 1n.asc # console에서 hex private key를 입력받아 `1n.asc` 파일에 armored 형식으로 저장
Enter 64 character raw hex private secp256k1 key:   # console에 위의 샘플 개인키(unarmored hex) 입력
Pick a password, at least 8 chars:  # 8자리 이상의 password를 정하여 입력한다. 나중에 import할때 다시 입력 필요

$ cat 1n.asc    # armored private key 확인
-----BEGIN OSTRACON PRIVATE KEY-----
kdf: bcrypt
salt: 71CB74A283E35CCC09FEDEE8870E306A

lzwnNWvRfyKI/kQi9/S42mBPAtpFwq6/IIvcKC5RkdZ4/fHwnXO0ah4UAa4OYSz0
fnvQBTKN2unxxawc9VZytom97vzTPuG7hXuSuYo=
=fLzy
-----END OSTRACON PRIVATE KEY-----

$ fnsad keys import 1n 1n.asc   # 1n 이라는 이름으로 keyring에 import
Enter passphrase to decrypt your key:   # 위에서 입력한 8자리 이상의 password

$ fnsad keys list   # keyring에 추가됨을 확인
- name: 1n
  type: local
  address: link1w508d6qejxtdg4y5r3zarvary0c5xw7k0nhc5s
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"Anm+Zn753LusVaBilc6HCwcCm/zbLc4o2VnygVsW+BeY"}'
  mnemonic: ""
```

## 참고자료

- [Validator Guide - Agoric/agoric-sdk Wiki](https://github.com/Agoric/agoric-sdk/wiki/Validator-Guide#private-key-only-no-mnemonic-recovery-phrase)
- [import unarmored-hex private key in keychain - issue #4024 - agoric-sdk](https://github.com/Agoric/agoric-sdk/issues/4024)
- [Finschia/finschia-sdk](https://github.com/Finschia/finschia-sdk)

## LICENSE

Apache 2.0
