# Changelog

## [1.4.6](https://github.com/Mrflatt/netconf/compare/v1.4.5...v1.4.6) (2026-08-19)


### Miscellaneous Chores

* bump the minor-dependencies group with 3 updates ([#84](https://github.com/Mrflatt/netconf/issues/84)) ([d419c93](https://github.com/Mrflatt/netconf/commit/d419c93f176db093cba121e0fd942f3f9c2266c9))

## [1.4.5](https://github.com/Mrflatt/netconf/compare/v1.4.4...v1.4.5) (2026-08-13)


### Bug Fixes

* stop Get hanging on vendor XML with illegal control chars ([#82](https://github.com/Mrflatt/netconf/issues/82)) ([0dbdb83](https://github.com/Mrflatt/netconf/commit/0dbdb83329b6023d8fdb667e78389f08bdd44202))


### Miscellaneous Chores

* bump actions/checkout from 6 to 7 in the actions group ([#79](https://github.com/Mrflatt/netconf/issues/79)) ([f0ae792](https://github.com/Mrflatt/netconf/commit/f0ae792598c95b1576eed866947bab234f167f4c))
* bump the minor-dependencies group across 1 directory with 2 updates ([#81](https://github.com/Mrflatt/netconf/issues/81)) ([9301d81](https://github.com/Mrflatt/netconf/commit/9301d81578ceccd2d43cff1f6e6fb5c65c7bbe5a))

## [1.4.4](https://github.com/Mrflatt/netconf/compare/v1.4.3...v1.4.4) (2026-06-17)


### Bug Fixes

* register reply channel before writing message ([a4453ba](https://github.com/Mrflatt/netconf/commit/a4453ba5de48557f68c959f30d0a09753b982b44))


### Miscellaneous Chores

* add MIT license to the project ([7953940](https://github.com/Mrflatt/netconf/commit/7953940ec07e75538b97261443e80837c4220712))
* bump golang.org/x/crypto in the minor-dependencies group ([#74](https://github.com/Mrflatt/netconf/issues/74)) ([d7d3184](https://github.com/Mrflatt/netconf/commit/d7d3184ad03a5b745e3bc27a3c6bd100a47a6692))
* bump the actions group across 1 directory with 2 updates ([#77](https://github.com/Mrflatt/netconf/issues/77)) ([8459f25](https://github.com/Mrflatt/netconf/commit/8459f25e9be69a58cd9e22ef5c80ab2a55601d3f))
* bump the actions group across 1 directory with 4 updates ([#71](https://github.com/Mrflatt/netconf/issues/71)) ([e1e8c6c](https://github.com/Mrflatt/netconf/commit/e1e8c6cdd0ba448212d10bc0c4aab6ea5cd1f6ff))
* bump the minor-dependencies group with 2 updates ([#69](https://github.com/Mrflatt/netconf/issues/69)) ([4a2689f](https://github.com/Mrflatt/netconf/commit/4a2689f601d85980455883133ada085f54e8cbf8))
* bump the minor-dependencies group with 2 updates ([#78](https://github.com/Mrflatt/netconf/issues/78)) ([2c069d0](https://github.com/Mrflatt/netconf/commit/2c069d0097ab139224600a61b129e47c86b5adfc))
* bump WillAbides/setup-go-faster in the actions group ([#73](https://github.com/Mrflatt/netconf/issues/73)) ([d5ed810](https://github.com/Mrflatt/netconf/commit/d5ed810f105f54724dd9591853693d9a118f9f1d))
* migrate module from networkguild to Mrflatt ([29f7863](https://github.com/Mrflatt/netconf/commit/29f78636a8e8bc6d6ea515218bfacd31dec6f91c))

## [1.4.3](https://github.com/networkguild/netconf/compare/v1.4.2...v1.4.3) (2026-04-29)


### Bug Fixes

* add XML declaration and vendor namespace support for RPC messages ([#66](https://github.com/networkguild/netconf/issues/66)) ([c823a6d](https://github.com/networkguild/netconf/commit/c823a6db875c19471da86af0f335ba39505fb64c))

## [1.4.2](https://github.com/networkguild/netconf/compare/v1.4.1...v1.4.2) (2026-04-29)


### Bug Fixes

* **session:** set CharsetReader to handle non-UTF-8 XML encodings ([a1e6611](https://github.com/networkguild/netconf/commit/a1e6611592734c5ca5cd363d55725ce844400939))

## [1.4.1](https://github.com/networkguild/netconf/compare/v1.4.0...v1.4.1) (2025-07-18)


### Bug Fixes

* **rpc:** rename rpc and rpc reply with all caps ([68e817e](https://github.com/networkguild/netconf/commit/68e817e11afd3df96f0205a474a16794eb652eb0))
* **rpc:** rename rpc and rpc reply with all caps ([2af061d](https://github.com/networkguild/netconf/commit/2af061d299cd98bb62c229a772af32ca9ee5f49e))
* use github auto-merge for dependabot ([c86c03f](https://github.com/networkguild/netconf/commit/c86c03f48e13dc69a0c81eb968a00f05bfefd12b))


### Miscellaneous Chores

* bump github.com/go-viper/mapstructure/v2 in the go_modules group ([#53](https://github.com/networkguild/netconf/issues/53)) ([f75ea5c](https://github.com/networkguild/netconf/commit/f75ea5cd433e3f8b800fe95f8297424345c9e197))
* bump golangci/golangci-lint-action in the actions group ([#50](https://github.com/networkguild/netconf/issues/50)) ([9b94cb1](https://github.com/networkguild/netconf/commit/9b94cb11efea60d6527b494815eeb671fa414ac7))
* bump the minor-dependencies group across 1 directory with 2 updates ([#52](https://github.com/networkguild/netconf/issues/52)) ([174f122](https://github.com/networkguild/netconf/commit/174f1221c3647f986b2c258bc4740247d15c011a))

## [1.4.0](https://github.com/networkguild/netconf/compare/v1.3.1...v1.4.0) (2025-04-29)


### Features

* add capability parsing and validation to ops ([f48732d](https://github.com/networkguild/netconf/commit/f48732d9adeb71497000440d894598a6bafa5c86))
* add golangci config and workflow check ([db54684](https://github.com/networkguild/netconf/commit/db546841a0941eca17d46577a6bf0f123a5bcb09))
* add release-please for automatic release creation ([bd7a135](https://github.com/networkguild/netconf/commit/bd7a13537be10f55443c37d0d8593359f64c592e))


### Bug Fixes

* change expected test assert ([27ca3aa](https://github.com/networkguild/netconf/commit/27ca3aaaa04afda5cd51a7a97d114805ba3daf6b))
* remove xml tags from messages not to fail serializing ([a8f6d2c](https://github.com/networkguild/netconf/commit/a8f6d2ccc96bbf6981043cd0f7497f38096deb36))


### Miscellaneous Chores

* Add another release-please job ([#46](https://github.com/networkguild/netconf/issues/46)) ([0604de7](https://github.com/networkguild/netconf/commit/0604de7f3bf30464b2d5a2c480e2e0e7bcf0ce2e))
* add release-please-config ([#47](https://github.com/networkguild/netconf/issues/47)) ([c890037](https://github.com/networkguild/netconf/commit/c89003775d6fc1ec11e3b6ce63e661e1b7283bab))
