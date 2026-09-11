# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.3](https://github.com/matthewhartstonge/argon2/compare/v1.6.0...v1.0.3) (2026-09-11)


### ⚠ BREAKING CHANGES

* **deps:** bump golang.org/x/crypto from 0.13.0 to 0.16.0

### Features

* **cmd/argon2:** adds argon2 cli. ([f0efb16](https://github.com/matthewhartstonge/argon2/commit/f0efb16f936f6fadf3b5a6686b8ad952caa4f8f7))
* **cmd/argon2:** adds argon2 cli. ([32a22f9](https://github.com/matthewhartstonge/argon2/commit/32a22f986e28f574553af5a4540b5d74439f1492))
* **deps:** bump golang.org/x/crypto from 0.13.0 to 0.16.0 ([8dbc527](https://github.com/matthewhartstonge/argon2/commit/8dbc52707b213d2c69660edb118b56f915eec4b0))
* **deps:** upgrades to go@1.23.0. ([9c86600](https://github.com/matthewhartstonge/argon2/commit/9c86600313babe362ca9cbaae5d270931e72e8c2))
* **deps:** upgrades to go@1.24.0. ([#103](https://github.com/matthewhartstonge/argon2/issues/103)) ([c97401e](https://github.com/matthewhartstonge/argon2/commit/c97401ee22f98f87da22929df1228dd883606e43))
* **deps:** upgrades to go@1.25.0. ([#124](https://github.com/matthewhartstonge/argon2/issues/124)) ([f17ebfc](https://github.com/matthewhartstonge/argon2/commit/f17ebfc1e3b84f59b15d95ac9e3230aa65a1f58c))
* **deps:** upgrades to go@1.26.0. ([#150](https://github.com/matthewhartstonge/argon2/issues/150)) ([4edf827](https://github.com/matthewhartstonge/argon2/commit/4edf827ca357890a61c914489f88b61f6c5d5cac))
* **goreleaser:** adds configuration for automated builds and publishing. ([6c3e044](https://github.com/matthewhartstonge/argon2/commit/6c3e04444c101e95926bdb009dacdb11052d4f3b))
* **goreleaser:** adds configuration for automated builds. ([c8385f8](https://github.com/matthewhartstonge/argon2/commit/c8385f89c37ee265cb8afdfa34fdb56777daf9d1))
* return error on attempting to hash `argon2d`. ([71f8bcb](https://github.com/matthewhartstonge/argon2/commit/71f8bcb6797b19eeb4b142e7f7a6b7d56cee521c))


### Bug Fixes

* **.github/dependabot:** fixes branch targets in dependabot configuration. ([64ab947](https://github.com/matthewhartstonge/argon2/commit/64ab947d1921dc80d8352516cf9c711044b2ed8f))
* Bump golang.org/x/crypto from 0.36.0 to 0.37.0 ([#87](https://github.com/matthewhartstonge/argon2/issues/87)) ([3c08e2b](https://github.com/matthewhartstonge/argon2/commit/3c08e2bd89c9b90792e94198ce542440d018a6db))
* **cmd/argon2:** enable ldflags variable configuration. ([988c8fe](https://github.com/matthewhartstonge/argon2/commit/988c8fec2409ed091efe93ae9ab4d99da7b851f8))
* **cmd/argon2:** fixes cli parallelism print line. ([#126](https://github.com/matthewhartstonge/argon2/issues/126)) ([b08c0e9](https://github.com/matthewhartstonge/argon2/commit/b08c0e9f3fc9c265480d98849f49e7cf4f86eb6e))
* **deps:** bump golang.org/x/crypto from 0.37.0 to 0.38.0 ([#90](https://github.com/matthewhartstonge/argon2/issues/90)) ([4d70c13](https://github.com/matthewhartstonge/argon2/commit/4d70c13642edccdd2f8acc1255be8e8bb79dcafa))
* **deps:** bump golang.org/x/crypto from 0.38.0 to 0.39.0 ([#93](https://github.com/matthewhartstonge/argon2/issues/93)) ([313b810](https://github.com/matthewhartstonge/argon2/commit/313b810d8b4f4c6cf8076ad8f9614dd284b2dd62))
* **deps:** bump golang.org/x/crypto from 0.39.0 to 0.40.0 ([#95](https://github.com/matthewhartstonge/argon2/issues/95)) ([2b750ba](https://github.com/matthewhartstonge/argon2/commit/2b750bacbd0efe280224a2eaa9b5a58228621e55))
* **deps:** bump golang.org/x/crypto from 0.40.0 to 0.41.0 ([#98](https://github.com/matthewhartstonge/argon2/issues/98)) ([7856060](https://github.com/matthewhartstonge/argon2/commit/7856060a8ff5b50d179ddff805b92f512bf43fe5))
* **deps:** bump golang.org/x/crypto from 0.41.0 to 0.42.0 ([#102](https://github.com/matthewhartstonge/argon2/issues/102)) ([9585efb](https://github.com/matthewhartstonge/argon2/commit/9585efb963b3a021af0aadc7df478085b71a2ddf))
* **deps:** bump golang.org/x/crypto from 0.42.0 to 0.43.0 ([#105](https://github.com/matthewhartstonge/argon2/issues/105)) ([86719dd](https://github.com/matthewhartstonge/argon2/commit/86719dd20df1ebfb7f9f9c42385408ca7d3105fc))
* **deps:** bump golang.org/x/crypto from 0.43.0 to 0.44.0 ([#108](https://github.com/matthewhartstonge/argon2/issues/108)) ([f708b79](https://github.com/matthewhartstonge/argon2/commit/f708b799722682a561d83f7660f34eb927db6a3e))
* **deps:** bump golang.org/x/crypto from 0.44.0 to 0.45.0 mitigates CVE-2025-58181 ([#111](https://github.com/matthewhartstonge/argon2/issues/111)) ([96af449](https://github.com/matthewhartstonge/argon2/commit/96af449bc2d1dad3f802bffd431ec22ed1457059))
* **deps:** bump golang.org/x/crypto from 0.45.0 to 0.46.0 ([#116](https://github.com/matthewhartstonge/argon2/issues/116)) ([36de8a9](https://github.com/matthewhartstonge/argon2/commit/36de8a93a47ee25e39bfc2a3d30dff678c28c87a))
* **deps:** bump golang.org/x/crypto from 0.46.0 to 0.47.0 ([#118](https://github.com/matthewhartstonge/argon2/issues/118)) ([740f11f](https://github.com/matthewhartstonge/argon2/commit/740f11f3f60d3822847251ba26ed417e457f1aea))
* **deps:** bump golang.org/x/crypto from 0.47.0 to 0.48.0 ([#120](https://github.com/matthewhartstonge/argon2/issues/120)) ([ddd8f24](https://github.com/matthewhartstonge/argon2/commit/ddd8f24df8e6727ee97562a3871e4e968d6ec47b))
* **deps:** bump golang.org/x/crypto from 0.49.0 to 0.50.0 ([#131](https://github.com/matthewhartstonge/argon2/issues/131)) ([c6096ae](https://github.com/matthewhartstonge/argon2/commit/c6096ae2dcb7067e9526eb1523d0f205cb365386))
* **deps:** bump golang.org/x/crypto from 0.50.0 to 0.51.0 ([#134](https://github.com/matthewhartstonge/argon2/issues/134)) ([0e1c93a](https://github.com/matthewhartstonge/argon2/commit/0e1c93aff3555666f1e1c9acebb7f586337e7b21))
* **deps:** bump golang.org/x/crypto from 0.51.0 to 0.52.0 ([#137](https://github.com/matthewhartstonge/argon2/issues/137)) ([5d52290](https://github.com/matthewhartstonge/argon2/commit/5d52290393a9c321992c7df0146f4ba0b9dd12c3))
* **deps:** bump golang.org/x/crypto from 0.52.0 to 0.53.0 ([#139](https://github.com/matthewhartstonge/argon2/issues/139)) ([d3fcd9a](https://github.com/matthewhartstonge/argon2/commit/d3fcd9ab595ff546f39075184f67295789090ce0))
* **deps:** bump golang.org/x/crypto from 0.53.0 to 0.54.0 ([#143](https://github.com/matthewhartstonge/argon2/issues/143)) ([1f89e55](https://github.com/matthewhartstonge/argon2/commit/1f89e556d32547b821d1d3fa5f88bf752ae72227))
* **deps:** bump golang.org/x/crypto from 0.54.0 to 0.55.0 ([#146](https://github.com/matthewhartstonge/argon2/issues/146)) ([adaaf65](https://github.com/matthewhartstonge/argon2/commit/adaaf6572f95e9fe826888060503df8e9186e09c))
* **deps:** bump golang.org/x/crypto from 0.55.0 to 0.56.0 ([#148](https://github.com/matthewhartstonge/argon2/issues/148)) ([1dc64d8](https://github.com/matthewhartstonge/argon2/commit/1dc64d8aa900becd4c806cedb38e04d8eb341c27))
* **lint:** fixes gosec G115 - potential under/overflow on hash and salt length when calling `Decode`. Now returns `ErrDecodingFail`. ([7c659c8](https://github.com/matthewhartstonge/argon2/commit/7c659c8b35ed8ceecf97b8045c288c7dd32137c0))
* **lint:** fixes gosec reported G115 on Decode ([a4d601c](https://github.com/matthewhartstonge/argon2/commit/a4d601c8280d4d5a6481b563ba1a4c385636637b))
* **lint:** migrates if-else blocks to tagged switches (QF1003). ([04cdec5](https://github.com/matthewhartstonge/argon2/commit/04cdec5d85be929c70f569a0da4d128ba565014c))
* publish and notarise binaries. ([4f962ec](https://github.com/matthewhartstonge/argon2/commit/4f962ecce5751578a3946e8b75c84ac1f4af02e0))
* restore support for 32-bit platforms ([#128](https://github.com/matthewhartstonge/argon2/issues/128)) ([3dcd610](https://github.com/matthewhartstonge/argon2/commit/3dcd610ce2c358ec2b39b1377bd87ed6156ae1f1))


### Miscellaneous

* **.github/workflows/go:** adds go@v1.20 to the test matrix. ([55710d9](https://github.com/matthewhartstonge/argon2/commit/55710d9e7f6ece529d7084d16587a9d2f6d1e573))
* **.github/workflows:** adds support for release-please. ([0538a35](https://github.com/matthewhartstonge/argon2/commit/0538a352a7efe5c95c4597d84c8a33d33ec47f69))
* adds information on crypto upgrade. ([8819f0c](https://github.com/matthewhartstonge/argon2/commit/8819f0c11967ee4564ed17b65f203ce9781db364))
* bootstrap releases for path: . ([c111ef3](https://github.com/matthewhartstonge/argon2/commit/c111ef3c7cecaa02536ab82ff201c6a9e9f7cb54))
* Bump actions/setup-go from 4 to 5 ([#29](https://github.com/matthewhartstonge/argon2/issues/29)) ([b7fa053](https://github.com/matthewhartstonge/argon2/commit/b7fa05398980d2c3a320e7ad1552bd6a803e12c1))
* Bump golang.org/x/crypto from 0.16.0 to 0.17.0 ([#30](https://github.com/matthewhartstonge/argon2/issues/30)) ([32e3c85](https://github.com/matthewhartstonge/argon2/commit/32e3c85f78dd52f4cbd6567c8a340ae611007044))
* Bump golang.org/x/crypto from 0.17.0 to 0.18.0 ([#32](https://github.com/matthewhartstonge/argon2/issues/32)) ([7dd1a2a](https://github.com/matthewhartstonge/argon2/commit/7dd1a2ab245e8404512a84aeb71006584a40f098))
* Bump golang.org/x/crypto from 0.18.0 to 0.19.0 ([#33](https://github.com/matthewhartstonge/argon2/issues/33)) ([cd7c9a7](https://github.com/matthewhartstonge/argon2/commit/cd7c9a78058e4f0e15e6a6a74bb5a53706509cc4))
* Bump golang.org/x/crypto from 0.19.0 to 0.20.0 ([#36](https://github.com/matthewhartstonge/argon2/issues/36)) ([16c9336](https://github.com/matthewhartstonge/argon2/commit/16c93368c683deae1a4fa69e3f792db6ead4ce4a))
* Bump golang.org/x/crypto from 0.20.0 to 0.21.0 ([#38](https://github.com/matthewhartstonge/argon2/issues/38)) ([ad1845a](https://github.com/matthewhartstonge/argon2/commit/ad1845a9c56e489ee167bfd133d96fe0f53d1cbf))
* Bump golang.org/x/crypto from 0.21.0 to 0.22.0 ([#39](https://github.com/matthewhartstonge/argon2/issues/39)) ([6d705ff](https://github.com/matthewhartstonge/argon2/commit/6d705ff1b3bb6753e6e408501ad59223cb643848))
* Bump golang.org/x/crypto from 0.22.0 to 0.23.0 ([#43](https://github.com/matthewhartstonge/argon2/issues/43)) ([33f46e2](https://github.com/matthewhartstonge/argon2/commit/33f46e29c0c136071121a5afe88d9dd25eff798f))
* Bump golang.org/x/crypto from 0.23.0 to 0.24.0 ([#46](https://github.com/matthewhartstonge/argon2/issues/46)) ([ab2ac7e](https://github.com/matthewhartstonge/argon2/commit/ab2ac7e2489131b1f57b3c2aa1cb547b62fdd7a7))
* Bump golang.org/x/crypto from 0.24.0 to 0.25.0 ([#47](https://github.com/matthewhartstonge/argon2/issues/47)) ([39504a4](https://github.com/matthewhartstonge/argon2/commit/39504a4dd821c0b559a904ff10452a60a9d69c05))
* Bump golang.org/x/crypto from 0.25.0 to 0.26.0 ([#49](https://github.com/matthewhartstonge/argon2/issues/49)) ([62179c3](https://github.com/matthewhartstonge/argon2/commit/62179c3b7f7d2c64ab790cf240b2100095ae2097))
* Bump golang.org/x/crypto from 0.26.0 to 0.27.0 ([#50](https://github.com/matthewhartstonge/argon2/issues/50)) ([8517a11](https://github.com/matthewhartstonge/argon2/commit/8517a11380ebf6b9f98b878de811bf1140def7fd))
* Bump golang.org/x/crypto from 0.27.0 to 0.28.0 ([#56](https://github.com/matthewhartstonge/argon2/issues/56)) ([dbe7209](https://github.com/matthewhartstonge/argon2/commit/dbe72099829775cac5519e4afcd5b475200bebef))
* Bump golang.org/x/crypto from 0.28.0 to 0.29.0 ([#57](https://github.com/matthewhartstonge/argon2/issues/57)) ([7d61928](https://github.com/matthewhartstonge/argon2/commit/7d619289f4cafec245e657e6ea183d48f00e7cbd))
* Bump golang.org/x/crypto from 0.29.0 to 0.30.0 ([#59](https://github.com/matthewhartstonge/argon2/issues/59)) ([5f7242a](https://github.com/matthewhartstonge/argon2/commit/5f7242a0294cdf6ec59878681151ed095b96d2d4))
* Bump golang.org/x/crypto from 0.30.0 to 0.31.0 ([#60](https://github.com/matthewhartstonge/argon2/issues/60)) ([3bb40c5](https://github.com/matthewhartstonge/argon2/commit/3bb40c5e87345e039ea1f406b10fd07850f8120c))
* Bump golang.org/x/crypto from 0.31.0 to 0.32.0 ([6ecf48f](https://github.com/matthewhartstonge/argon2/commit/6ecf48faf21a36476e87f72ec19afd6ee6f45cc5))
* Bump golang.org/x/crypto from 0.32.0 to 0.33.0 ([9d1020e](https://github.com/matthewhartstonge/argon2/commit/9d1020e94da63dfcf120ac75675c8a36d19a64f9))
* Bump golang.org/x/crypto from 0.32.0 to 0.33.0 ([2c465a0](https://github.com/matthewhartstonge/argon2/commit/2c465a0b50b5bc4c22c97a692daabfc83d207968))
* Bump golang.org/x/crypto from 0.33.0 to 0.36.0 ([0283d44](https://github.com/matthewhartstonge/argon2/commit/0283d440f256384d38974605463ab83c33e15b3c))
* Bump golang.org/x/crypto from 0.33.0 to 0.36.0 ([d1e6eaf](https://github.com/matthewhartstonge/argon2/commit/d1e6eafb582e8b24d6b99502bb88d8834a6c7e7a))
* Bump golangci/golangci-lint-action from 3.7.0 to 3.7.1 ([#34](https://github.com/matthewhartstonge/argon2/issues/34)) ([177c9d6](https://github.com/matthewhartstonge/argon2/commit/177c9d690f4f2fc066b792edd8379020596e5875))
* Bump golangci/golangci-lint-action from 3.7.1 to 4.0.0 ([#35](https://github.com/matthewhartstonge/argon2/issues/35)) ([7ddea0e](https://github.com/matthewhartstonge/argon2/commit/7ddea0e312720d5462e051dd4ce9b0b1ac7336e4))
* Bump golangci/golangci-lint-action from 4.0.0 to 5.0.0 ([#40](https://github.com/matthewhartstonge/argon2/issues/40)) ([418f97f](https://github.com/matthewhartstonge/argon2/commit/418f97fccfec574be7d7545aec6ff30615ce7cd8))
* Bump golangci/golangci-lint-action from 5.0.0 to 5.1.0 ([#41](https://github.com/matthewhartstonge/argon2/issues/41)) ([5c3c1e6](https://github.com/matthewhartstonge/argon2/commit/5c3c1e62a4d344f3ab6beba7f314d39a296759c8))
* Bump golangci/golangci-lint-action from 5.1.0 to 5.3.0 ([#42](https://github.com/matthewhartstonge/argon2/issues/42)) ([c82c954](https://github.com/matthewhartstonge/argon2/commit/c82c95462787d725021a143428a501e4c78eb74b))
* Bump golangci/golangci-lint-action from 5.3.0 to 6.0.0 ([#44](https://github.com/matthewhartstonge/argon2/issues/44)) ([d6b2856](https://github.com/matthewhartstonge/argon2/commit/d6b2856b998ff981b0220ab1410cc54f6e86c6bd))
* Bump golangci/golangci-lint-action from 6.0.0 to 6.0.1 ([#45](https://github.com/matthewhartstonge/argon2/issues/45)) ([f58a919](https://github.com/matthewhartstonge/argon2/commit/f58a91979973d8bd92fdee4a77fa1f5c7b85f2cd))
* Bump golangci/golangci-lint-action from 6.0.1 to 6.1.0 ([#48](https://github.com/matthewhartstonge/argon2/issues/48)) ([7c1a9fe](https://github.com/matthewhartstonge/argon2/commit/7c1a9fe9d5ad67d378c2e8848fb68da1c066cdda))
* Bump golangci/golangci-lint-action from 6.1.0 to 6.1.1 ([#55](https://github.com/matthewhartstonge/argon2/issues/55)) ([383fe5d](https://github.com/matthewhartstonge/argon2/commit/383fe5d38c619d5bf773d4cbf537a2c7e3ba831d))
* Bump golangci/golangci-lint-action from 6.1.1 to 6.2.0 ([5831551](https://github.com/matthewhartstonge/argon2/commit/5831551ff9bf35e42f425a33b29a543c49edc8f5))
* Bump golangci/golangci-lint-action from 6.1.1 to 6.2.0 ([9910661](https://github.com/matthewhartstonge/argon2/commit/9910661967765098b90e3457384b0838d866712b))
* Bump golangci/golangci-lint-action from 6.2.0 to 6.3.0 ([802130e](https://github.com/matthewhartstonge/argon2/commit/802130e82fcb57175d71f46dadd78727db0e9625))
* Bump golangci/golangci-lint-action from 6.2.0 to 6.3.0 ([ce43c35](https://github.com/matthewhartstonge/argon2/commit/ce43c35578485c9dc949801711d6599ac8b2bab1))
* Bump golangci/golangci-lint-action from 6.3.0 to 6.3.2 ([e19c7f4](https://github.com/matthewhartstonge/argon2/commit/e19c7f488545a61153f25cdcbec3a1ef35e29c15))
* Bump golangci/golangci-lint-action from 6.3.0 to 6.3.2 ([221e817](https://github.com/matthewhartstonge/argon2/commit/221e8174444973f90ed1cd1f493c05e9264ee985))
* Bump golangci/golangci-lint-action from 6.3.2 to 6.4.0 ([6d2c6ab](https://github.com/matthewhartstonge/argon2/commit/6d2c6ab04fa3f4cd85e5e24fbbad825ecbc253c3))
* Bump golangci/golangci-lint-action from 6.4.0 to 6.4.1 ([e06c03d](https://github.com/matthewhartstonge/argon2/commit/e06c03dd4be91c219ff7a740ee5bd1bb8ccc0000))
* Bump golangci/golangci-lint-action from 6.4.1 to 6.5.0 ([b6b483c](https://github.com/matthewhartstonge/argon2/commit/b6b483c22fdd4dd65e2575e58859ddf34aa0d644))
* Bump golangci/golangci-lint-action from 6.5.0 to 6.5.1 ([cd17b23](https://github.com/matthewhartstonge/argon2/commit/cd17b23374ccc35f8e114292c3038521c67df4a7))
* Bump golangci/golangci-lint-action from 6.5.1 to 6.5.2 ([c3b63ba](https://github.com/matthewhartstonge/argon2/commit/c3b63ba1a3d180fd2be1f98c3bc36372bd28ed89))
* Bump golangci/golangci-lint-action from 6.5.2 to 7.0.0 ([f7f7312](https://github.com/matthewhartstonge/argon2/commit/f7f7312e5fb4a1e429dfff78ffdc47f50d92ab86))
* Bump golangci/golangci-lint-action from 6.5.2 to 7.0.0 ([b50f10b](https://github.com/matthewhartstonge/argon2/commit/b50f10b4becb803b4a2b4dcffe2986cf38168b6d))
* **cmd/argon2:** adds commit information. ([1fc1b93](https://github.com/matthewhartstonge/argon2/commit/1fc1b9329d98bfd31742209dedec7427d2447452))
* **cmd/argon2:** adds custom usage signature. ([29e7635](https://github.com/matthewhartstonge/argon2/commit/29e76358e66d407541a04aca724916903a97be1c))
* **deps:** updates dependencies ([efcbcce](https://github.com/matthewhartstonge/argon2/commit/efcbcce98406c8304d41b681330864af14e7aeb1))
* **lint:** fixes lint. ([74ac017](https://github.com/matthewhartstonge/argon2/commit/74ac017ab4af828cb0638bc0f0ae005853dd89b4))
* **main:** release 0.3.3 ([437acc0](https://github.com/matthewhartstonge/argon2/commit/437acc0fe33136b695561c2e8148275ded43925f))
* **main:** release 0.3.4 ([#23](https://github.com/matthewhartstonge/argon2/issues/23)) ([103bf62](https://github.com/matthewhartstonge/argon2/commit/103bf6263ca491f082b7549b63e6b94d5062b1a6))
* **main:** release 1.0.0 ([#28](https://github.com/matthewhartstonge/argon2/issues/28)) ([602a14d](https://github.com/matthewhartstonge/argon2/commit/602a14da583febc436dbbbcda71b59cc09bf2daf))
* **main:** release 1.0.1 ([#54](https://github.com/matthewhartstonge/argon2/issues/54)) ([5c69489](https://github.com/matthewhartstonge/argon2/commit/5c69489dbfedf0b51d44c8dc3c9afdf6ffc4476d))
* **main:** release 1.0.2 ([#58](https://github.com/matthewhartstonge/argon2/issues/58)) ([ead2fc9](https://github.com/matthewhartstonge/argon2/commit/ead2fc971e8cdcfbad7936800df11442c17ca213))
* **main:** release 1.0.3 ([a36e047](https://github.com/matthewhartstonge/argon2/commit/a36e047ced84c6250dff1fb62ab0aaf336caea03))
* **main:** release 1.0.3 ([19ae5a4](https://github.com/matthewhartstonge/argon2/commit/19ae5a4f3a5fc26e704b8b370fce18da4cc9ae13))
* **main:** release 1.1.0 ([b67b3b1](https://github.com/matthewhartstonge/argon2/commit/b67b3b14817664a839cbb51603a46f89e5ad4aa7))
* **main:** release 1.1.0 ([8f11089](https://github.com/matthewhartstonge/argon2/commit/8f110894155cf6258b73a88ff5cbe2d6d0be5963))
* **main:** release 1.1.1 ([8118fcd](https://github.com/matthewhartstonge/argon2/commit/8118fcd1712917cd71a90b4a523633dd61b61386))
* **main:** release 1.1.1 ([bdaf38b](https://github.com/matthewhartstonge/argon2/commit/bdaf38bf39e16276994c593100ddab224a6e5722))
* **main:** release 1.2.0 ([162c5c3](https://github.com/matthewhartstonge/argon2/commit/162c5c32a0c701e490d24558fd50dab52713f783))
* **main:** release 1.2.0 ([bdc2608](https://github.com/matthewhartstonge/argon2/commit/bdc26082e51524718d8fe05411a0096c95874e8e))
* **main:** release 1.2.1 ([#88](https://github.com/matthewhartstonge/argon2/issues/88)) ([456492a](https://github.com/matthewhartstonge/argon2/commit/456492a22cbacfc1e19218b63b37c4855c57b783))
* **main:** release 1.3.0 ([59778e3](https://github.com/matthewhartstonge/argon2/commit/59778e368bfc0228e58c7f97238a60a4bfd03591))
* **main:** release 1.3.0 ([eeb3737](https://github.com/matthewhartstonge/argon2/commit/eeb373729110f1571cae1fc401a38e44c387a13a))
* **main:** release 1.3.1 ([c5f6176](https://github.com/matthewhartstonge/argon2/commit/c5f61768ed91936fec702f97bb1de98af2200a3f))
* **main:** release 1.3.1 ([5589d19](https://github.com/matthewhartstonge/argon2/commit/5589d192cb5d0591cf9a1ab61430976515ebef04))
* **main:** release 1.3.2 ([#94](https://github.com/matthewhartstonge/argon2/issues/94)) ([3901986](https://github.com/matthewhartstonge/argon2/commit/390198654624593e9d78b3285af155fcc78e98ab))
* **main:** release 1.3.3 ([c0c3d50](https://github.com/matthewhartstonge/argon2/commit/c0c3d50f451481bbc1f2a6ca19af20c39525992f))
* **main:** release 1.3.4 ([#99](https://github.com/matthewhartstonge/argon2/issues/99)) ([a04e355](https://github.com/matthewhartstonge/argon2/commit/a04e355568b4f747d8c3ab54c9f6dcb75f1382c0))
* **main:** release 1.4.0 ([#104](https://github.com/matthewhartstonge/argon2/issues/104)) ([40a464b](https://github.com/matthewhartstonge/argon2/commit/40a464b1d9e6d2c5d1c4fdb4e80954e2046005d5))
* **main:** release 1.4.1 ([#106](https://github.com/matthewhartstonge/argon2/issues/106)) ([b7ef281](https://github.com/matthewhartstonge/argon2/commit/b7ef281749f7014caeb30d5905f7991b0da968bd))
* **main:** release 1.4.2 ([#109](https://github.com/matthewhartstonge/argon2/issues/109)) ([5e0359d](https://github.com/matthewhartstonge/argon2/commit/5e0359d420a3351102751f073b9a28a5285b51e8))
* **main:** release 1.4.3 ([#112](https://github.com/matthewhartstonge/argon2/issues/112)) ([46b2273](https://github.com/matthewhartstonge/argon2/commit/46b227336162bc15da28e81d7730767b7f007d1b))
* **main:** release 1.4.4 ([#117](https://github.com/matthewhartstonge/argon2/issues/117)) ([0ec6119](https://github.com/matthewhartstonge/argon2/commit/0ec61198e7473d2b7bc7fd27ee37145d3311c1d1))
* **main:** release 1.4.5 ([#119](https://github.com/matthewhartstonge/argon2/issues/119)) ([d04b999](https://github.com/matthewhartstonge/argon2/commit/d04b999282e8de713665d4c5dab36009be2eff69))
* **main:** release 1.4.6 ([#121](https://github.com/matthewhartstonge/argon2/issues/121)) ([2d735e7](https://github.com/matthewhartstonge/argon2/commit/2d735e7e9c5276d26fb1ec090332af3a7dad0a3a))
* **main:** release 1.5.0 ([#125](https://github.com/matthewhartstonge/argon2/issues/125)) ([bc3334a](https://github.com/matthewhartstonge/argon2/commit/bc3334a0b03115fadc82ffd1a6d49a3219fe714e))
* **main:** release 1.5.1 ([#130](https://github.com/matthewhartstonge/argon2/issues/130)) ([3c929c5](https://github.com/matthewhartstonge/argon2/commit/3c929c5a2956dd6afe22fbf36d21bb8860d2c47e))
* **main:** release 1.5.2 ([#132](https://github.com/matthewhartstonge/argon2/issues/132)) ([b0a069d](https://github.com/matthewhartstonge/argon2/commit/b0a069debe21e24d7e32f45faea1ca60bcde0193))
* **main:** release 1.5.3 ([#135](https://github.com/matthewhartstonge/argon2/issues/135)) ([45032e1](https://github.com/matthewhartstonge/argon2/commit/45032e1b025a2cd752afe90a4ca09f5b7733ccb8))
* **main:** release 1.5.4 ([#138](https://github.com/matthewhartstonge/argon2/issues/138)) ([166ea53](https://github.com/matthewhartstonge/argon2/commit/166ea53735f5b0f35a37c8f583bbd78c6f5f8c9c))
* **main:** release 1.5.5 ([#140](https://github.com/matthewhartstonge/argon2/issues/140)) ([b320a8e](https://github.com/matthewhartstonge/argon2/commit/b320a8e0ba321d124dfbed6f4116464fa9d3a268))
* **main:** release 1.5.6 ([#144](https://github.com/matthewhartstonge/argon2/issues/144)) ([563b703](https://github.com/matthewhartstonge/argon2/commit/563b703311319f4fc7ecb6e1754b3c635524a9c3))
* **main:** release 1.5.7 ([#147](https://github.com/matthewhartstonge/argon2/issues/147)) ([7d1f6f0](https://github.com/matthewhartstonge/argon2/commit/7d1f6f0a9ccc397f4d65e4ac8ad7ba6c899e093e))
* **main:** release 1.6.0 ([#151](https://github.com/matthewhartstonge/argon2/issues/151)) ([111cead](https://github.com/matthewhartstonge/argon2/commit/111cead24f5a38c1706a7d8bf82f2f6904b0cbed))
* release 1.0.0 ([fbe015c](https://github.com/matthewhartstonge/argon2/commit/fbe015cd6ebc9dc6890d00d5f09cd4b42583fc28))
* release 1.0.1 ([f68ab6f](https://github.com/matthewhartstonge/argon2/commit/f68ab6fc71041029ca749a899646046fb9f10ce3))
* release 1.0.1 ([#52](https://github.com/matthewhartstonge/argon2/issues/52)) ([9117a69](https://github.com/matthewhartstonge/argon2/commit/9117a69d61ea0ad70e3f8d92b9259a7f06030cd6))
* release 1.0.2 ([64bcf86](https://github.com/matthewhartstonge/argon2/commit/64bcf86837060e454fafb59acc3d12feb3a3322d))
* release 1.0.3 ([6221117](https://github.com/matthewhartstonge/argon2/commit/6221117dbb6b8eedb57d3b877d521bfc0ad4caff))
* release 1.0.3 ([601f492](https://github.com/matthewhartstonge/argon2/commit/601f49212436450d975dc2b4f9f30c6d0018ab8f))


### Build System

* **.goreleaser.yaml:** use short commit hash. ([ddf767a](https://github.com/matthewhartstonge/argon2/commit/ddf767affd93421db9a503612adb2831d892ae34))
* **deps:** bump actions/checkout from 3 to 4 ([#21](https://github.com/matthewhartstonge/argon2/issues/21)) ([782498e](https://github.com/matthewhartstonge/argon2/commit/782498e32e1904049215014c3e56ce9500a0ad7d))
* **deps:** bump actions/checkout from 4 to 5 ([#100](https://github.com/matthewhartstonge/argon2/issues/100)) ([b5a0969](https://github.com/matthewhartstonge/argon2/commit/b5a096968efe979c1a9ec2dad6c2179e550bd1b2))
* **deps:** bump actions/checkout from 5 to 6 ([#113](https://github.com/matthewhartstonge/argon2/issues/113)) ([5121e08](https://github.com/matthewhartstonge/argon2/commit/5121e08ab23f601a99a97974bb025e62efadfbda))
* **deps:** bump actions/checkout from 6 to 7 ([#141](https://github.com/matthewhartstonge/argon2/issues/141)) ([b895844](https://github.com/matthewhartstonge/argon2/commit/b8958444e31bca7ed7d8f412e6ac3ada4395ff08))
* **deps:** bump actions/setup-go from 5 to 6 ([#101](https://github.com/matthewhartstonge/argon2/issues/101)) ([8ee356f](https://github.com/matthewhartstonge/argon2/commit/8ee356fbf1bf626891a3e299794b51035a3e366f))
* **deps:** bump golang.org/x/crypto from 0.11.0 to 0.12.0 ([#19](https://github.com/matthewhartstonge/argon2/issues/19)) ([11f7788](https://github.com/matthewhartstonge/argon2/commit/11f7788ab0e5dcfd1d77fb0cd5aa79d1fc49a871))
* **deps:** bump golang.org/x/crypto from 0.12.0 to 0.13.0 ([#22](https://github.com/matthewhartstonge/argon2/issues/22)) ([7c46a68](https://github.com/matthewhartstonge/argon2/commit/7c46a681a9b1b1bf18ac5de8d8c7e0b2322837d2))
* **deps:** bump golang.org/x/crypto from 0.48.0 to 0.49.0 ([#123](https://github.com/matthewhartstonge/argon2/issues/123)) ([78c3433](https://github.com/matthewhartstonge/argon2/commit/78c34336ca1136e10ebc9fde1f3616ec91652e76))
* **deps:** bump golangci/golangci-lint-action from 3.6.0 to 3.7.0 ([#20](https://github.com/matthewhartstonge/argon2/issues/20)) ([6d8007c](https://github.com/matthewhartstonge/argon2/commit/6d8007c954b69068fa8b55c0166e813dd9558a25))
* **deps:** bump golangci/golangci-lint-action from 7.0.0 to 8.0.0 ([a37e4ad](https://github.com/matthewhartstonge/argon2/commit/a37e4adef7eb645d2e3c02bff080ef95baa201ef))
* **deps:** bump golangci/golangci-lint-action from 7.0.0 to 8.0.0 ([f98283b](https://github.com/matthewhartstonge/argon2/commit/f98283b14ee2de9fb8c4042f5f3ffbe0c3317577))
* **deps:** bump golangci/golangci-lint-action from 8.0.0 to 9.0.0 ([#107](https://github.com/matthewhartstonge/argon2/issues/107)) ([45899d5](https://github.com/matthewhartstonge/argon2/commit/45899d5d8602b0971271c299ac7f0b1a3d755d48))
* **deps:** bump golangci/golangci-lint-action from 9.0.0 to 9.1.0 ([#114](https://github.com/matthewhartstonge/argon2/issues/114)) ([cd8e9d8](https://github.com/matthewhartstonge/argon2/commit/cd8e9d8e5beb1e0467ea62e6910cceec77f87dab))
* **deps:** bump golangci/golangci-lint-action from 9.2.1 to 9.3.0 ([#142](https://github.com/matthewhartstonge/argon2/issues/142)) ([b36334f](https://github.com/matthewhartstonge/argon2/commit/b36334fa1421dd93c293305bf1eb17f6afa6971f))
* **deps:** bump google-github-actions/release-please-action from 3 to 4 ([#27](https://github.com/matthewhartstonge/argon2/issues/27)) ([7551bcc](https://github.com/matthewhartstonge/argon2/commit/7551bcc3dd3e60a9f0b06a1090ecec29f0b133d9))
* **deps:** bump googleapis/release-please-action from 4 to 5 ([#133](https://github.com/matthewhartstonge/argon2/issues/133)) ([86ded65](https://github.com/matthewhartstonge/argon2/commit/86ded65845985d79d45dd83c767b439defd33c2c))
* **deps:** bump goreleaser/goreleaser-action from 6 to 7 ([#122](https://github.com/matthewhartstonge/argon2/issues/122)) ([2de8356](https://github.com/matthewhartstonge/argon2/commit/2de83561f7ddf363a19feb837189adaca5919d49))
* **deps:** updates golang.org/x/crypto from 0.10.0 to 0.11.0 ([7c48950](https://github.com/matthewhartstonge/argon2/commit/7c48950b480813ed463a8ef246865489362ad3aa))
* **deps:** updates golang.org/x/crypto from 0.3.0 to 0.10.0 ([7189d50](https://github.com/matthewhartstonge/argon2/commit/7189d5018449b8e03213b648bec0308aac86c218))
* **deps:** updates golang.org/x/crypto from 0.3.0 to 0.10.0 ([d9c0208](https://github.com/matthewhartstonge/argon2/commit/d9c02085e54d01a0ee299c494e799ec84a401c51))
* **goreleaser:** migrates 'format' usages to `archives.formats` and `archives.format_overrides.formats`. ([903d599](https://github.com/matthewhartstonge/argon2/commit/903d599b2a139a8334aedaa53000f2b83ed3f281))
* **lint:** adds custom golangci-lint config to include gosec. ([ab8e690](https://github.com/matthewhartstonge/argon2/commit/ab8e6904f9074478ddd7b8efd4c2a50046041f2c))


### CI

* **.github/workflows/go:** tests against `go@1.22`. ([#37](https://github.com/matthewhartstonge/argon2/issues/37)) ([65bc263](https://github.com/matthewhartstonge/argon2/commit/65bc263dab1c8fcd8cab7ce6d1951e3b0e97ecdc))
* **.github/workflows/go:** use {n, n-1} for tests. ([9432b59](https://github.com/matthewhartstonge/argon2/commit/9432b59f617e2038c70c3b9404ef3b14c88944ad))
* **.github/workflows/release:** enable notarising macOS binaries. ([470fac9](https://github.com/matthewhartstonge/argon2/commit/470fac9e34c6c18f0c1874d7344ea7eeaefe54c4))
* **.github/workflows/release:** use n-1 for release builds. ([f37ca08](https://github.com/matthewhartstonge/argon2/commit/f37ca08c2f95ff256db895431d591789f16f19fe))
* **.github/workflows:** adds goreleaser publishing. ([f142dcd](https://github.com/matthewhartstonge/argon2/commit/f142dcd65bb927d03a34727a8a36f23216a6a49a))
* **.github/workflows:** consolidate release workflow. ([acd9d01](https://github.com/matthewhartstonge/argon2/commit/acd9d01fb1602fa9caa13d02cbdcb50bb941bbe5))
* adds go@1.23 to the testing matrix. ([#51](https://github.com/matthewhartstonge/argon2/issues/51)) ([876ead4](https://github.com/matthewhartstonge/argon2/commit/876ead45cc74199018e5bf2aa200921c40af197c))
* **cmd/argon2:** removes lower bounds checking for uint. ([61cb37f](https://github.com/matthewhartstonge/argon2/commit/61cb37fe12b64816bd5f3ced63930318e376c8d5))
* **deps:** bump golangci/golangci-lint-action from 9.2.0 to 9.2.1 ([#136](https://github.com/matthewhartstonge/argon2/issues/136)) ([90946e5](https://github.com/matthewhartstonge/argon2/commit/90946e5619b03680074ffea45dfb34131db9c598))
* **goreleaser:** enables release publishing. ([b0f7e35](https://github.com/matthewhartstonge/argon2/commit/b0f7e35dc815111ed133b88ffce5bc4b82bb5ade))
* **release-please-config.json:** automate binary version setting based on tagged release. ([0bf8bc5](https://github.com/matthewhartstonge/argon2/commit/0bf8bc5dbf63aec62c7d8e0c64eda3188f03d40e))
* **release-please:** configures changelog sections. ([9c48acc](https://github.com/matthewhartstonge/argon2/commit/9c48acc502ee9190ec03bbf8d7642719aa282846))
* **release-please:** moves to draft releases, to enable go-releaser publishing. ([267d498](https://github.com/matthewhartstonge/argon2/commit/267d4981385091e37e346abc2b564c7afb728d68))


### Tests

* **.github/workflows:** test for 32-bit platform support. ([#129](https://github.com/matthewhartstonge/argon2/issues/129)) ([d2602bc](https://github.com/matthewhartstonge/argon2/commit/d2602bcf52b5d89f958e0868978e07802c8ef022))
* **Decode:** check for expected errors. ([caee659](https://github.com/matthewhartstonge/argon2/commit/caee6592a2fab9f83b595c89ac645451fd829f27))


### Documentation

* add information on the `golang.org/x` versioning strategy. ([4ab11d6](https://github.com/matthewhartstonge/argon2/commit/4ab11d6f18454581fd34268749b40e231a600c76))
* adds information on requirement to upgrade to `go@1.23.0`. ([2ad21c9](https://github.com/matthewhartstonge/argon2/commit/2ad21c9beb1abf2d11dcb047e110ec3a6748f285))
* **argon2:** adds information on not being able to verify argon2d hashes. ([7174361](https://github.com/matthewhartstonge/argon2/commit/7174361e99cce3cf9df7d958a0925e83cbe23983))


### Refactors

* make decode strings self documenting. ([dcbb888](https://github.com/matthewhartstonge/argon2/commit/dcbb88823afa4e2f906639fef0557c8fe50756bb))

## [1.6.0](https://github.com/matthewhartstonge/argon2/compare/v1.5.7...v1.6.0) (2026-09-10)


### Features

* **deps:** upgrades to go@1.26.0. ([#150](https://github.com/matthewhartstonge/argon2/issues/150)) ([4edf827](https://github.com/matthewhartstonge/argon2/commit/4edf827ca357890a61c914489f88b61f6c5d5cac))


### Bug Fixes

* **deps:** bump golang.org/x/crypto from 0.55.0 to 0.56.0 ([#148](https://github.com/matthewhartstonge/argon2/issues/148)) ([1dc64d8](https://github.com/matthewhartstonge/argon2/commit/1dc64d8aa900becd4c806cedb38e04d8eb341c27))


### CI

* **goreleaser:** enables release publishing. ([b0f7e35](https://github.com/matthewhartstonge/argon2/commit/b0f7e35dc815111ed133b88ffce5bc4b82bb5ade))
* **release-please:** configures changelog sections. ([9c48acc](https://github.com/matthewhartstonge/argon2/commit/9c48acc502ee9190ec03bbf8d7642719aa282846))
* **release-please:** moves to draft releases, to enable go-releaser publishing. ([267d498](https://github.com/matthewhartstonge/argon2/commit/267d4981385091e37e346abc2b564c7afb728d68))

## [1.5.7](https://github.com/matthewhartstonge/argon2/compare/v1.5.6...v1.5.7) (2026-08-15)


### Bug Fixes

* **deps:** bump golang.org/x/crypto from 0.54.0 to 0.55.0 ([#146](https://github.com/matthewhartstonge/argon2/issues/146)) ([adaaf65](https://github.com/matthewhartstonge/argon2/commit/adaaf6572f95e9fe826888060503df8e9186e09c))

## [1.5.6](https://github.com/matthewhartstonge/argon2/compare/v1.5.5...v1.5.6) (2026-07-12)


### Bug Fixes

* **deps:** bump golang.org/x/crypto from 0.53.0 to 0.54.0 ([#143](https://github.com/matthewhartstonge/argon2/issues/143)) ([1f89e55](https://github.com/matthewhartstonge/argon2/commit/1f89e556d32547b821d1d3fa5f88bf752ae72227))

## [1.5.5](https://github.com/matthewhartstonge/argon2/compare/v1.5.4...v1.5.5) (2026-06-11)


### Bug Fixes

* **deps:** bump golang.org/x/crypto from 0.52.0 to 0.53.0 ([#139](https://github.com/matthewhartstonge/argon2/issues/139)) ([d3fcd9a](https://github.com/matthewhartstonge/argon2/commit/d3fcd9ab595ff546f39075184f67295789090ce0))

## [1.5.4](https://github.com/matthewhartstonge/argon2/compare/v1.5.3...v1.5.4) (2026-05-24)


### Bug Fixes

* **deps:** bump golang.org/x/crypto from 0.51.0 to 0.52.0 ([#137](https://github.com/matthewhartstonge/argon2/issues/137)) ([5d52290](https://github.com/matthewhartstonge/argon2/commit/5d52290393a9c321992c7df0146f4ba0b9dd12c3))

## [1.5.3](https://github.com/matthewhartstonge/argon2/compare/v1.5.2...v1.5.3) (2026-05-14)


### Bug Fixes

* **deps:** bump golang.org/x/crypto from 0.50.0 to 0.51.0 ([#134](https://github.com/matthewhartstonge/argon2/issues/134)) ([0e1c93a](https://github.com/matthewhartstonge/argon2/commit/0e1c93aff3555666f1e1c9acebb7f586337e7b21))

## [1.5.2](https://github.com/matthewhartstonge/argon2/compare/v1.5.1...v1.5.2) (2026-04-09)


### Bug Fixes

* **deps:** bump golang.org/x/crypto from 0.49.0 to 0.50.0 ([#131](https://github.com/matthewhartstonge/argon2/issues/131)) ([c6096ae](https://github.com/matthewhartstonge/argon2/commit/c6096ae2dcb7067e9526eb1523d0f205cb365386))

## [1.5.1](https://github.com/matthewhartstonge/argon2/compare/v1.5.0...v1.5.1) (2026-04-07)


### Bug Fixes

* restore support for 32-bit platforms ([#128](https://github.com/matthewhartstonge/argon2/issues/128)) ([3dcd610](https://github.com/matthewhartstonge/argon2/commit/3dcd610ce2c358ec2b39b1377bd87ed6156ae1f1))

## [1.5.0](https://github.com/matthewhartstonge/argon2/compare/v1.4.6...v1.5.0) (2026-03-17)


### Features

* **deps:** upgrades to go@1.25.0. ([#124](https://github.com/matthewhartstonge/argon2/issues/124)) ([f17ebfc](https://github.com/matthewhartstonge/argon2/commit/f17ebfc1e3b84f59b15d95ac9e3230aa65a1f58c))


### Bug Fixes

* **cmd/argon2:** fixes cli parallelism print line. ([#126](https://github.com/matthewhartstonge/argon2/issues/126)) ([b08c0e9](https://github.com/matthewhartstonge/argon2/commit/b08c0e9f3fc9c265480d98849f49e7cf4f86eb6e))
* **lint:** fixes gosec G115 - potential under/overflow on hash and salt length when calling `Decode`. Now returns `ErrDecodingFail`. ([7c659c8](https://github.com/matthewhartstonge/argon2/commit/7c659c8b35ed8ceecf97b8045c288c7dd32137c0))
* **lint:** fixes gosec reported G115 on Decode ([a4d601c](https://github.com/matthewhartstonge/argon2/commit/a4d601c8280d4d5a6481b563ba1a4c385636637b))
* **deps:** bump golang.org/x/crypto from 0.48.0 to 0.49.0 ([#123](https://github.com/matthewhartstonge/argon2/issues/123)) ([78c3433](https://github.com/matthewhartstonge/argon2/commit/78c34336ca1136e10ebc9fde1f3616ec91652e76))

## [1.4.6](https://github.com/matthewhartstonge/argon2/compare/v1.4.5...v1.4.6) (2026-02-13)


### Bug Fixes

* **deps:** bump golang.org/x/crypto from 0.47.0 to 0.48.0 ([#120](https://github.com/matthewhartstonge/argon2/issues/120)) ([ddd8f24](https://github.com/matthewhartstonge/argon2/commit/ddd8f24df8e6727ee97562a3871e4e968d6ec47b))

## [1.4.5](https://github.com/matthewhartstonge/argon2/compare/v1.4.4...v1.4.5) (2026-01-27)


### Bug Fixes

* **deps:** bump golang.org/x/crypto from 0.46.0 to 0.47.0 ([#118](https://github.com/matthewhartstonge/argon2/issues/118)) ([740f11f](https://github.com/matthewhartstonge/argon2/commit/740f11f3f60d3822847251ba26ed417e457f1aea))

## [1.4.4](https://github.com/matthewhartstonge/argon2/compare/v1.4.3...v1.4.4) (2025-12-22)


### Bug Fixes

* **deps:** bump golang.org/x/crypto from 0.45.0 to 0.46.0 ([#116](https://github.com/matthewhartstonge/argon2/issues/116)) ([36de8a9](https://github.com/matthewhartstonge/argon2/commit/36de8a93a47ee25e39bfc2a3d30dff678c28c87a))

## [1.4.3](https://github.com/matthewhartstonge/argon2/compare/v1.4.2...v1.4.3) (2025-11-20)


### Bug Fixes

* **deps:** bump golang.org/x/crypto from 0.44.0 to 0.45.0 mitigates CVE-2025-58181 ([#111](https://github.com/matthewhartstonge/argon2/issues/111)) ([96af449](https://github.com/matthewhartstonge/argon2/commit/96af449bc2d1dad3f802bffd431ec22ed1457059))

## [1.4.2](https://github.com/matthewhartstonge/argon2/compare/v1.4.1...v1.4.2) (2025-11-12)


### Bug Fixes

* **deps:** bump golang.org/x/crypto from 0.43.0 to 0.44.0 ([#108](https://github.com/matthewhartstonge/argon2/issues/108)) ([f708b79](https://github.com/matthewhartstonge/argon2/commit/f708b799722682a561d83f7660f34eb927db6a3e))

## [1.4.1](https://github.com/matthewhartstonge/argon2/compare/v1.4.0...v1.4.1) (2025-10-10)


### Bug Fixes

* **deps:** bump golang.org/x/crypto from 0.42.0 to 0.43.0 ([#105](https://github.com/matthewhartstonge/argon2/issues/105)) ([86719dd](https://github.com/matthewhartstonge/argon2/commit/86719dd20df1ebfb7f9f9c42385408ca7d3105fc))

## [1.4.0](https://github.com/matthewhartstonge/argon2/compare/v1.3.4...v1.4.0) (2025-09-17)


### Features

* **deps:** upgrades to go@1.24.0. ([#103](https://github.com/matthewhartstonge/argon2/issues/103)) ([c97401e](https://github.com/matthewhartstonge/argon2/commit/c97401ee22f98f87da22929df1228dd883606e43))


### Bug Fixes

* **deps:** bump golang.org/x/crypto from 0.41.0 to 0.42.0 ([#102](https://github.com/matthewhartstonge/argon2/issues/102)) ([9585efb](https://github.com/matthewhartstonge/argon2/commit/9585efb963b3a021af0aadc7df478085b71a2ddf))

## [1.3.4](https://github.com/matthewhartstonge/argon2/compare/v1.3.3...v1.3.4) (2025-08-08)


### Bug Fixes

* **deps:** bump golang.org/x/crypto from 0.40.0 to 0.41.0 ([#98](https://github.com/matthewhartstonge/argon2/issues/98)) ([7856060](https://github.com/matthewhartstonge/argon2/commit/7856060a8ff5b50d179ddff805b92f512bf43fe5))

## [1.3.3](https://github.com/matthewhartstonge/argon2/compare/v1.3.2...v1.3.3) (2025-07-15)


### Bug Fixes

* **deps:** bump golang.org/x/crypto from 0.39.0 to 0.40.0 ([#95](https://github.com/matthewhartstonge/argon2/issues/95)) ([2b750ba](https://github.com/matthewhartstonge/argon2/commit/2b750bacbd0efe280224a2eaa9b5a58228621e55))

## [1.3.2](https://github.com/matthewhartstonge/argon2/compare/v1.3.1...v1.3.2) (2025-06-05)


### Bug Fixes

* **deps:** bump golang.org/x/crypto from 0.38.0 to 0.39.0 ([#93](https://github.com/matthewhartstonge/argon2/issues/93)) ([313b810](https://github.com/matthewhartstonge/argon2/commit/313b810d8b4f4c6cf8076ad8f9614dd284b2dd62))

## [1.3.1](https://github.com/matthewhartstonge/argon2/compare/v1.3.0...v1.3.1) (2025-05-06)


### Bug Fixes

* **deps:** bump golang.org/x/crypto from 0.37.0 to 0.38.0 ([#90](https://github.com/matthewhartstonge/argon2/issues/90)) ([4d70c13](https://github.com/matthewhartstonge/argon2/commit/4d70c13642edccdd2f8acc1255be8e8bb79dcafa))

## [1.3.0](https://github.com/matthewhartstonge/argon2/compare/v1.2.1...v1.3.0) (2025-04-22)


### Features

* return error on attempting to hash `argon2d`. ([71f8bcb](https://github.com/matthewhartstonge/argon2/commit/71f8bcb6797b19eeb4b142e7f7a6b7d56cee521c))


### Bug Fixes

* **lint:** migrates if-else blocks to tagged switches (QF1003). ([04cdec5](https://github.com/matthewhartstonge/argon2/commit/04cdec5d85be929c70f569a0da4d128ba565014c))

## [1.2.1](https://github.com/matthewhartstonge/argon2/compare/v1.2.0...v1.2.1) (2025-04-08)


### Bug Fixes

* Bump golang.org/x/crypto from 0.36.0 to 0.37.0 ([#87](https://github.com/matthewhartstonge/argon2/issues/87)) ([3c08e2b](https://github.com/matthewhartstonge/argon2/commit/3c08e2bd89c9b90792e94198ce542440d018a6db))

## [1.2.0](https://github.com/matthewhartstonge/argon2/compare/v1.1.1...v1.2.0) (2025-03-06)

### Features

* **deps:** upgrades to go@1.23.0. ([9c86600](https://github.com/matthewhartstonge/argon2/commit/9c86600313babe362ca9cbaae5d270931e72e8c2))

note: `go@1.21` introduced a change in how the go directive works, now enforcing Minimal Version Selection (MVS). The best thing to do is upgrade your Go toolchain to `n-1` and set the `go` directive in your `go.mod` file to `go 1.(N-1).0` - notice the requirement for the patch set to `0`.

If you are a library maintainer, PLEASE make sure to only set this to `0` and remove the `toolchain` directive, otherwise you will make downstream users will cry (myself included).

Why? As part of an internal Go proposal, all golang.org/x libraries now only support n-1.

> I propose that each time that a new major Go release 1.N.0 is made, the go directive in all golang.org/x repos will be upgraded to go 1.(N-1).0. For example, when Go 1.28.0 is released, golang.org/x modules would have their go directive set to go 1.27.0.

Refer: https://go.googlesource.com/proposal/+/master/design/69095-x-repo-continuous-go.md

## [1.1.1](https://github.com/matthewhartstonge/argon2/compare/v1.1.0...v1.1.1) (2024-12-23)


### Bug Fixes

* **cmd/argon2:** enable ldflags variable configuration. ([988c8fe](https://github.com/matthewhartstonge/argon2/commit/988c8fec2409ed091efe93ae9ab4d99da7b851f8))

## [1.1.0](https://github.com/matthewhartstonge/argon2/compare/v1.0.3...v1.1.0) (2024-12-21)


### Features

* **cmd/argon2:** adds argon2 cli. ([32a22f9](https://github.com/matthewhartstonge/argon2/commit/32a22f986e28f574553af5a4540b5d74439f1492))
* **goreleaser:** adds configuration for automated builds. ([c8385f8](https://github.com/matthewhartstonge/argon2/commit/c8385f89c37ee265cb8afdfa34fdb56777daf9d1))

## [1.0.3](https://github.com/matthewhartstonge/argon2/compare/v1.0.2...v1.0.3) (2024-12-12)


### Miscellaneous Chores

* release 1.0.3 ([601f492](https://github.com/matthewhartstonge/argon2/commit/601f49212436450d975dc2b4f9f30c6d0018ab8f))
* chore: Bump golang.org/x/crypto from 0.30.0 to 0.31.0 ([3bb40c5](https://github.com/matthewhartstonge/argon2/commit/3bb40c5e87345e039ea1f406b10fd07850f8120c))

## [1.0.2](https://github.com/matthewhartstonge/argon2/compare/v1.0.1...v1.0.2) (2024-11-10)


### Miscellaneous Chores

* release 1.0.2 ([64bcf86](https://github.com/matthewhartstonge/argon2/commit/64bcf86837060e454fafb59acc3d12feb3a3322d))
* chore: Bump golang.org/x/crypto from 0.28.0 to 0.29.0 ([7d61928](https://github.com/matthewhartstonge/argon2/commit/7d619289f4cafec245e657e6ea183d48f00e7cbd))
* chore: Bump golang.org/x/crypto from 0.27.0 to 0.28.0 ([dbe7209](https://github.com/matthewhartstonge/argon2/commit/dbe72099829775cac5519e4afcd5b475200bebef))

## [1.0.1](https://github.com/matthewhartstonge/argon2/compare/v1.0.0...v1.0.1) (2024-09-12)


### Miscellaneous Chores

* release 1.0.1 ([f68ab6f](https://github.com/matthewhartstonge/argon2/commit/f68ab6fc71041029ca749a899646046fb9f10ce3))
* chore: Bump golang.org/x/crypto from 0.26.0 to 0.27.0 ([#50](https://github.com/matthewhartstonge/argon2/pull/50))

## [1.0.0](https://github.com/matthewhartstonge/argon2/compare/v0.3.4...v1.0.0) (2023-12-03)


### ⚠ BREAKING CHANGES

* **deps:** bump golang.org/x/crypto from 0.13.0 to 0.16.0

### Features

* **deps:** bump golang.org/x/crypto from 0.13.0 to 0.16.0 ([8dbc527](https://github.com/matthewhartstonge/argon2/commit/8dbc52707b213d2c69660edb118b56f915eec4b0))


### Miscellaneous Chores

* release 1.0.0 ([fbe015c](https://github.com/matthewhartstonge/argon2/commit/fbe015cd6ebc9dc6890d00d5f09cd4b42583fc28))

## [0.3.4](https://github.com/matthewhartstonge/argon2/compare/v0.3.3...v0.3.4) (2023-09-06)


### Miscellaneous Chores

* **deps:** updates dependencies ([efcbcce](https://github.com/matthewhartstonge/argon2/commit/efcbcce98406c8304d41b681330864af14e7aeb1))

## [0.3.3](https://github.com/matthewhartstonge/argon2/compare/v0.3.2...v0.3.3) (2023-07-04)


### Bug Fixes

* **.github/dependabot:** fixes branch targets in dependabot configuration. ([64ab947](https://github.com/matthewhartstonge/argon2/commit/64ab947d1921dc80d8352516cf9c711044b2ed8f))

## [Unreleased]
## [v0.3.2] - 2022-11-18
### Changed
- deps: updates to `golang.org/x/crypto@v0.3.0`.
- .github: adjusts dependabot to file version updates against development.

## [v0.3.1] - 2022-11-10
### Changed
- deps: updates to `golang.org/x/crypto@v0.2.0`.

## [v0.3.0] - 2022-11-05
### Added
- .github: adds support for github actions, codeowners and dependabot.

### Changed
- deps: updates to `go@1.18` and `golang.org/x/crypto@v0.1.0`.
- readme: update badges and remove references to travis ci.
- *: update project licensing.

### Fixed
- argon2: fixes grammar and comment line length.

### Removed
- dep: removes support for dep.
- travis: removes travis configuration.

## [v0.2.1] - 2022-02-22
GPG signed release.

### Added
- tests: adds benchmarks to compare the default recommendations.

### Changed
- travis: updates to test against go1.17.
- deps: updates `go.mod` to `go@1.17` and `x/crypto` to latest.
- explicitly ignores returned error values under benchmarks (errcheck).
- argon2: updates recommended defaults as described by RFC9106.

## [v0.2.0] - 2022-02-22
### Added
- tests: adds benchmarks to compare the default recommendations.

### Changed
- travis: updates to test against go1.17.
- deps: updates `go.mod` to `go@1.17` and `x/crypto` to latest.
- explicitly ignores returned error values under benchmarks (errcheck).
- argon2: updates recommended defaults as described by RFC9106.

## [v0.1.5] - 2021-08-06
### Changed
- travis: require go >= v1.9
- deps: bumps to the latest version of `golang.org/x/crypto`.
- Upgrades to mitigate users importing a vulnerable version of `golang.org/x/crypto/ssh` that contains CVE-2020-9283.
- Upgrades to mitigate users importing a vulnerable version of `golang.org/x/text` that contains CVE-2020-14040.
- travis: updated to test against go 1.12+, migrates to go mod for dependency management.
- travis: forces use of go modules under ci.
- travis: sets `go@v1.11.4` as the lowest supported `go mod` version due to a change in go build.

## [v0.1.4] - 2021-08-06
### Changed
- readme: updated.

## [v0.1.3] - 2021-08-06
### Added
- deps: adds support for go modules.

### Changed
- travis: updates to test against `go@{1.12, 1.13, 1.14}`
- deps: update `x/crypto@master` to point to latest commit.
  - This is mainly to mitigate users from CVEs in other `x/crypto` implementations, namely CVE-2020-7919 (cryptobyte) and CVE-2020-9283 (ssh).
- readme: adds a tl;dr, updates benchmarks.

### Fixed
- _example: fixes calling location of `VerifyEncoded()`

### Removed
- tests: removes benchmarking against native bindings to remove dependencies under go mod.

## [v0.1.2] - 2018-09-18
### Added
- tests: added benchmarks to compare against native argon2 bindings.

### Changed
- readme: updated to include travis build badge.
- argon2: Updates SecureZeroMemory to match upstream for better performance.
- deps: Updated to support dep v0.5.0
- readme: Update with new benchmark stats due to SecureZeroMemory performance tweak.

### Fixed
- readme: Fixes example pathing.

## [v0.1.1] - 2018-06-14
### Changed
- deps: unpins golang.org/x/crypto from a specific revision.

## [v0.1.0] - 2018-05-30
### Added
- Initial Commit

### Fixed
- git: Fixes repo github links
- readme: Fix example link

[Unreleased]: https://github.com/matthewhartstonge/argon2/tree/master
[v0.3.2]: https://github.com/matthewhartstonge/argon2/tree/v0.3.2
[v0.3.1]: https://github.com/matthewhartstonge/argon2/tree/v0.3.1
[v0.3.0]: https://github.com/matthewhartstonge/argon2/tree/v0.3.0
[v0.2.1]: https://github.com/matthewhartstonge/argon2/tree/v0.2.1
[v0.2.0]: https://github.com/matthewhartstonge/argon2/tree/v0.2.0
[v0.1.5]: https://github.com/matthewhartstonge/argon2/tree/v0.1.5
[v0.1.4]: https://github.com/matthewhartstonge/argon2/tree/v0.1.4
[v0.1.3]: https://github.com/matthewhartstonge/argon2/tree/v0.1.3
[v0.1.2]: https://github.com/matthewhartstonge/argon2/tree/v0.1.2
[v0.1.1]: https://github.com/matthewhartstonge/argon2/tree/v0.1.1
[v0.1.0]: https://github.com/matthewhartstonge/argon2/tree/v0.1.0
