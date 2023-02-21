
<a name="Bootstrap Go Mod Changelog"></a>
## Bootstrap Module (in Go)
[Github repository](https://github.com/edgexfoundry/go-mod-bootstrap)

## Change Logs for EdgeX Dependencies

- [go-mod-core-contracts](https://github.com/edgexfoundry/go-mod-core-contracts/blob/main/CHANGELOG.md)
- [go-mod-messaging](https://github.com/edgexfoundry/go-mod-messaging/blob/main/CHANGELOG.md)
- [go-mod-registry](https://github.com/edgexfoundry/go-mod-registry/blob/main/CHANGELOG.md)
- [go-mod-secrets](https://github.com/edgexfoundry/go-mod-secrets/blob/main/CHANGELOG.md)
- [go-mod-configuration](https://github.com/edgexfoundry/go-mod-configuration/blob/main/CHANGELOG.md) 

## [v2.3.0] - 2022-11-09

### Features ✨

- Add capability to use messaging based Command Client ([#384](https://github.com/edgexfoundry/go-mod-bootstrap/issues/384)) ([#9ad12a8](https://github.com/edgexfoundry/go-mod-bootstrap/commits/9ad12a8))
- Add Consul security metrics ([#383](https://github.com/edgexfoundry/go-mod-bootstrap/issues/383)) ([#a43e448](https://github.com/edgexfoundry/go-mod-bootstrap/commits/a43e448))
- Add service metrics for Secrets requested and stored ([#376](https://github.com/edgexfoundry/go-mod-bootstrap/issues/376)) ([#42c52e2](https://github.com/edgexfoundry/go-mod-bootstrap/commits/42c52e2))
- Added SecretUpdated  API ([#373](https://github.com/edgexfoundry/go-mod-bootstrap/issues/373)) ([#f58aa0b](https://github.com/edgexfoundry/go-mod-bootstrap/commits/f58aa0b))
- Redact logging of insecure secrets env override ([#367](https://github.com/edgexfoundry/go-mod-bootstrap/issues/367)) ([#9565883](https://github.com/edgexfoundry/go-mod-bootstrap/commits/9565883))
- Added HasSecret API ([#364](https://github.com/edgexfoundry/go-mod-bootstrap/issues/364)) ([#61f5503](https://github.com/edgexfoundry/go-mod-bootstrap/commits/61f5503))
- Add new 'Topics' field and external MQTT BootstrapHandler ([#365](https://github.com/edgexfoundry/go-mod-bootstrap/issues/365)) ([#6dab13b](https://github.com/edgexfoundry/go-mod-bootstrap/commits/6dab13b))
- Add common Messaging bootstrap handler ([#360](https://github.com/edgexfoundry/go-mod-bootstrap/issues/360)) ([#aaf2123](https://github.com/edgexfoundry/go-mod-bootstrap/commits/aaf2123))
- Add Histogram to supported metric types ([#346](https://github.com/edgexfoundry/go-mod-bootstrap/issues/346)) ([#57130b2](https://github.com/edgexfoundry/go-mod-bootstrap/commits/57130b2))
- Put CA cert into MessageBusInfo for all AuthModes ([#324](https://github.com/edgexfoundry/go-mod-bootstrap/issues/324)) ([#4dbfa01](https://github.com/edgexfoundry/go-mod-bootstrap/commits/4dbfa01))

### Bug Fixes 🐛

- Add capability to override config provider settings with "none" ([#381](https://github.com/edgexfoundry/go-mod-bootstrap/issues/381)) ([#3493ca4](https://github.com/edgexfoundry/go-mod-bootstrap/commits/3493ca4))
- Run WatchForChange in a new thread ([#362](https://github.com/edgexfoundry/go-mod-bootstrap/issues/362)) ([#9c98e1c](https://github.com/edgexfoundry/go-mod-bootstrap/commits/9c98e1c))
- Ensure exit with non-zero code when error occurs ([#358](https://github.com/edgexfoundry/go-mod-bootstrap/issues/358)) ([#816d4c9](https://github.com/edgexfoundry/go-mod-bootstrap/commits/816d4c9))

### Build 👷

- Upgrade to Go 1.18 ([#1361f04](https://github.com/edgexfoundry/go-mod-bootstrap/commit/1361f04))

## [v2.2.0] - 2022-05-11

### Features ✨

- Add RequestLimitMiddleware for Service.MaxRequestSize config ([#321](https://github.com/edgexfoundry/go-mod-bootstrap/issues/321)) ([#42b690d](https://github.com/edgexfoundry/go-mod-bootstrap/commits/42b690d))
- Implement service metrics bootstrap and common capability ([#313](https://github.com/edgexfoundry/go-mod-bootstrap/issues/313)) ([#8132711](https://github.com/edgexfoundry/go-mod-bootstrap/commits/8132711))
- Location of client service obtained from the registry ([#305](https://github.com/edgexfoundry/go-mod-bootstrap/issues/305)) ([#78c5fc9](https://github.com/edgexfoundry/go-mod-bootstrap/commits/78c5fc9))
- **security:** Use go-mod-secrets version that includes the capability of using non_delayedstart go build tags ([#317](https://github.com/edgexfoundry/go-mod-bootstrap/issues/317)) ([#2a6ac6a](https://github.com/edgexfoundry/go-mod-bootstrap/commits/2a6ac6a))
- **security:** Integrate runtime spiffe token provider client from go-mod-secrets ([#4bf6376](https://github.com/edgexfoundry/go-mod-bootstrap/commits/4bf6376))

### Bug Fixes 🐛

- Generate proper Consul basepath on Windows ([#0cfe34c](https://github.com/edgexfoundry/go-mod-bootstrap/commits/0cfe34c))
- **config:** ignore first change notification in ListenForCustomConfigChanges ([#315](https://github.com/edgexfoundry/go-mod-bootstrap/issues/315)) ([#6332299](https://github.com/edgexfoundry/go-mod-bootstrap/commits/6332299))

### Build 👷

- Added "make lint" target and added to "make test" target  ([#302](https://github.com/edgexfoundry/go-mod-bootstrap/issues/302)) ([#d813076](https://github.com/edgexfoundry/go-mod-bootstrap/commits/d813076))

<a name="v2.1.0"></a>
## [v2.1.0] - 2021-11-17

### Features ✨

- Use Http Request timeout handler ([#267](https://github.com/edgexfoundry/go-mod-bootstrap/issues/267)) ([#4da2238](https://github.com/edgexfoundry/go-mod-bootstrap/commits/4da2238))
- **security:** Add Access Token callback Vault token reload ([#285](https://github.com/edgexfoundry/go-mod-bootstrap/issues/285)) ([#64217dd](https://github.com/edgexfoundry/go-mod-bootstrap/commits/64217dd))
- **security:** Add optional capability to seed service secrets ([#276](https://github.com/edgexfoundry/go-mod-bootstrap/issues/276)) ([#a4676a4](https://github.com/edgexfoundry/go-mod-bootstrap/commits/a4676a4))
- **security:** Add func to process CORS ([#288](https://github.com/edgexfoundry/go-mod-bootstrap/issues/288)) ([#c292656](https://github.com/edgexfoundry/go-mod-bootstrap/commits/c292656))
- **security:** Create CORS related config struct ([#286](https://github.com/edgexfoundry/go-mod-bootstrap/issues/286)) ([#4ec4738](https://github.com/edgexfoundry/go-mod-bootstrap/commits/4ec4738))

### Bug Fixes 🐛

- Use correct name when logging EDGEX_CONF_DIR override ([#266](https://github.com/edgexfoundry/go-mod-bootstrap/issues/266)) ([#2a375e7](https://github.com/edgexfoundry/go-mod-bootstrap/commits/2a375e7))

## [v2.0.0] - 2021-06-30
### Features ✨
- **v2:** Add Subscribe config to MessageQueue config ([#240](https://github.com/edgexfoundry/go-mod-bootstrap/issues/240)) ([#ac14ba0](https://github.com/edgexfoundry/go-mod-bootstrap/commits/ac14ba0))
- **v2:** Add bootstrap handler to create Messaging Client with secure options ([#225](https://github.com/edgexfoundry/go-mod-bootstrap/issues/225)) ([#ae196fc](https://github.com/edgexfoundry/go-mod-bootstrap/commits/ae196fc))
- **v2:** Use SecretProvider to get Config/Registry access tokens ([#202](https://github.com/edgexfoundry/go-mod-bootstrap/issues/202)) ([#5d19aa5](https://github.com/edgexfoundry/go-mod-bootstrap/commits/5d19aa5))
- **v2:** Enable use of Registry & Config client access token ([#195](https://github.com/edgexfoundry/go-mod-bootstrap/issues/195)) ([#f9d06ec](https://github.com/edgexfoundry/go-mod-bootstrap/commits/f9d06ec))
- **v2:** Add overwrite capability for custom configuration ([#185](https://github.com/edgexfoundry/go-mod-bootstrap/issues/185)) ([#90b8a51](https://github.com/edgexfoundry/go-mod-bootstrap/commits/90b8a51))
- **v2:** Add support for load/listen custom configuration ([#180](https://github.com/edgexfoundry/go-mod-bootstrap/issues/180)) ([#f277873](https://github.com/edgexfoundry/go-mod-bootstrap/commits/f277873))
- **v2:** Add config client in DIC ([#178](https://github.com/edgexfoundry/go-mod-bootstrap/issues/178)) ([#ecde49d](https://github.com/edgexfoundry/go-mod-bootstrap/commits/ecde49d))
- **v2:** Add helper to query DIC and returns the DeviceServiceCommandClient instance ([#162](https://github.com/edgexfoundry/go-mod-bootstrap/issues/162)) ([#c087e44](https://github.com/edgexfoundry/go-mod-bootstrap/commits/c087e44))
- **v2:** Create Helper functions to retrieve client library instances through DIC ([#158](https://github.com/edgexfoundry/go-mod-bootstrap/issues/158)) ([#3d89601](https://github.com/edgexfoundry/go-mod-bootstrap/commits/3d89601))
### Bug Fixes 🐛
- Use /api/v2/ping for Registry healthchecks ([#196](https://github.com/edgexfoundry/go-mod-bootstrap/issues/196)) ([#7d55b1a](https://github.com/edgexfoundry/go-mod-bootstrap/commits/7d55b1a))
- Add conditional for error message and return false on error ([#f4390fe](https://github.com/edgexfoundry/go-mod-bootstrap/commits/f4390fe))
- Replace hyphen with underscore in override names ([#216](https://github.com/edgexfoundry/go-mod-bootstrap/issues/216)) ([#9f3edfd](https://github.com/edgexfoundry/go-mod-bootstrap/commits/9f3edfd))
    ```
    BREAKING CHANGE:
    Overrides that have hyphens will not longer work and must be updated replace hyphens with underscores.
    ```
- Remove messaging handler to avoid implicit ZMQ dependency ([#235](https://github.com/edgexfoundry/go-mod-bootstrap/issues/235)) ([#9df977d](https://github.com/edgexfoundry/go-mod-bootstrap/commits/9df977d))
- Fix Secure MessageBus Secret validation for non-secure mode ([#233](https://github.com/edgexfoundry/go-mod-bootstrap/issues/233)) ([#f6c98ef](https://github.com/edgexfoundry/go-mod-bootstrap/commits/f6c98ef))
- Generate mock for latest SecretProvider interface ([#206](https://github.com/edgexfoundry/go-mod-bootstrap/issues/206)) ([#359809f](https://github.com/edgexfoundry/go-mod-bootstrap/commits/359809f))
- Use V2 Ping for health check ([#5bb40c1](https://github.com/edgexfoundry/go-mod-bootstrap/commits/5bb40c1))
- **secuirty:** remove retry config items from SecretStore config ([#248](https://github.com/edgexfoundry/go-mod-bootstrap/issues/248)) ([#6002097](https://github.com/edgexfoundry/go-mod-bootstrap/commits/6002097))
### Code Refactoring ♻
- Update ServiceInfo struct to be used by all services and add MaxRequestSize ([#9e3af34](https://github.com/edgexfoundry/go-mod-bootstrap/commits/9e3af34))
    ```
    BREAKING CHANGE:
    Service configuration has changed for all services
    ```
- Update calling GenerateConsulToken ([#212](https://github.com/edgexfoundry/go-mod-bootstrap/issues/212)) ([#e295a6e](https://github.com/edgexfoundry/go-mod-bootstrap/commits/e295a6e))
- Replace use of BurntSushi/toml with pelletier/go-toml ([#6c8f2b4](https://github.com/edgexfoundry/go-mod-bootstrap/commits/6c8f2b4))
- Expose ConfigVersion so services can use if needed ([#204](https://github.com/edgexfoundry/go-mod-bootstrap/issues/204)) ([#e966ad5](https://github.com/edgexfoundry/go-mod-bootstrap/commits/e966ad5))
- Set the Config Version when creating Config Client ([#201](https://github.com/edgexfoundry/go-mod-bootstrap/issues/201)) ([#615e600](https://github.com/edgexfoundry/go-mod-bootstrap/commits/615e600))
    ```
    BREAKING CHANGE:
    Configuration in Consul now under the `/2.0/` path
    ```
- Refactor ListenForCustomConfigChanges to avoid use of channel ([#187](https://github.com/edgexfoundry/go-mod-bootstrap/issues/187)) ([#cffb2fe](https://github.com/edgexfoundry/go-mod-bootstrap/commits/cffb2fe))
- Updated go.mod for tagged go-mod-secrets and fixed unittest ([#05db8a1](https://github.com/edgexfoundry/go-mod-bootstrap/commits/05db8a1))
- Add comment for new Type setting. ([#d2e6caa](https://github.com/edgexfoundry/go-mod-bootstrap/commits/d2e6caa))

<a name="v0.0.68"></a>
## [v0.0.68] - 2021-01-04
### Features ✨
- Enhance Timer to be used for timed loops beyond bootstrapping ([#141](https://github.com/edgexfoundry/go-mod-bootstrap/issues/141)) ([#ff8e38c](https://github.com/edgexfoundry/go-mod-bootstrap/commits/ff8e38c))

<a name="v0.0.67"></a>
## [v0.0.67] - 2021-01-04
### Bug Fixes 🐛
- Add setting the configured LogLevel ([#143](https://github.com/edgexfoundry/go-mod-bootstrap/issues/143)) ([#9cbc3d8](https://github.com/edgexfoundry/go-mod-bootstrap/commits/9cbc3d8))

<a name="v0.0.66"></a>
## [v0.0.66] - 2020-12-30
### Code Refactoring ♻
- Remove backward compatibility code ([#139](https://github.com/edgexfoundry/go-mod-bootstrap/issues/139)) ([#c10d266](https://github.com/edgexfoundry/go-mod-bootstrap/commits/c10d266))

<a name="v0.0.65"></a>
## [v0.0.65] - 2020-12-29
### Code Refactoring ♻
- Refactor to remove remote and file logging ([#138](https://github.com/edgexfoundry/go-mod-bootstrap/issues/138)) ([#d92118e](https://github.com/edgexfoundry/go-mod-bootstrap/commits/d92118e))

<a name="v0.0.62"></a>
## [v0.0.62] - 2020-12-20
### Code Refactoring ♻
- Secret Provider for all services ([#134](https://github.com/edgexfoundry/go-mod-bootstrap/issues/134)) ([#6cb9329](https://github.com/edgexfoundry/go-mod-bootstrap/commits/6cb9329))

<a name="v0.0.59"></a>
## [v0.0.59] - 2020-11-25
### Bug Fixes 🐛
- LoggingClientFrom handle nil case properly ([#c95d24f](https://github.com/edgexfoundry/go-mod-bootstrap/commits/c95d24f))

<a name="v0.0.58"></a>
## [v0.0.58] - 2020-11-19
### Features ✨
- Allow service to pass in initial logging client ([#3651de7](https://github.com/edgexfoundry/go-mod-bootstrap/commits/3651de7))

<a name="v0.0.57"></a>
## [v0.0.57] - 2020-10-28
### Bug Fixes 🐛
- Accept argument lists with a -r substring ([#dc0e6ea](https://github.com/edgexfoundry/go-mod-bootstrap/commits/dc0e6ea))

<a name="v0.0.50"></a>
## [v0.0.50] - 2020-10-14
### Bug Fixes 🐛
- Handle env override values which have the '=' character ([#4846fb7](https://github.com/edgexfoundry/go-mod-bootstrap/commits/4846fb7))

<a name="v0.0.41"></a>
## [v0.0.41] - 2020-09-29
### Bug Fixes 🐛
- Increase default startup duration to 60 seconds ([#0761e33](https://github.com/edgexfoundry/go-mod-bootstrap/commits/0761e33))

<a name="v0.0.37"></a>
## [v0.0.37] - 2020-07-30
### Bug Fixes 🐛
- Startup Duration and Interval never updated from default values ([#c35f13c](https://github.com/edgexfoundry/go-mod-bootstrap/commits/c35f13c))

<a name="v0.0.36"></a>
## [v0.0.36] - 2020-07-13
### Bug Fixes 🐛
- Configurable ip address for ListenAndServe, fixes [#83](https://github.com/edgexfoundry/go-mod-bootstrap/issues/83) ([#ec63238](https://github.com/edgexfoundry/go-mod-bootstrap/commits/ec63238))

<a name="v0.0.35"></a>
## [v0.0.35] - 2020-07-07
### Code Refactoring ♻
- **config:** Remove ClientMonitor from the ServiceInfo struct ([#efe9cb9](https://github.com/edgexfoundry/go-mod-bootstrap/commits/efe9cb9))

<a name="v0.0.33"></a>
## [v0.0.33] - 2020-06-01
### Bug Fixes 🐛
- Changed from using blank hostname to 0.0.0.0 ([#38f87ec](https://github.com/edgexfoundry/go-mod-bootstrap/commits/38f87ec))
- Don't use hostname for webserver ListenAndServe ([#6dbe24f](https://github.com/edgexfoundry/go-mod-bootstrap/commits/6dbe24f))

<a name="v0.0.32"></a>
## [v0.0.32] - 2020-05-29
### Bug Fixes 🐛
- Allow overrides that have empty/blank value ([#5497010](https://github.com/edgexfoundry/go-mod-bootstrap/commits/5497010))

<a name="v0.0.31"></a>
## [v0.0.31] - 2020-04-29
### Bug Fixes 🐛
- **config:** Ignore first config changes notification on start-up ([#2834834](https://github.com/edgexfoundry/go-mod-bootstrap/commits/2834834))

<a name="v0.0.30"></a>
## [v0.0.30] - 2020-04-21
### Features ✨
- **environment:** Perform case insensitive comparision for override names ([#3d7becb](https://github.com/edgexfoundry/go-mod-bootstrap/commits/3d7becb))

<a name="v0.0.28"></a>
## [v0.0.28] - 2020-04-14
### Bug Fixes 🐛
- **config:** Change UpdatedStream to be defined as `chan struct{}` ([#6d2e43b](https://github.com/edgexfoundry/go-mod-bootstrap/commits/6d2e43b))

<a name="v0.0.26"></a>
## [v0.0.26] - 2020-03-31
### Bug Fixes 🐛
- **logging:** Logger not configured properly ([#017c944](https://github.com/edgexfoundry/go-mod-bootstrap/commits/017c944))

<a name="v0.0.25"></a>
## [v0.0.25] - 2020-03-30
### Features ✨
- Add Self seeding, env var overrides, cmd-line options per ADR 0005-Service-Self-Config.md ([#59](https://github.com/edgexfoundry/go-mod-bootstrap/issues/59)) ([#e56334c](https://github.com/edgexfoundry/go-mod-bootstrap/commits/e56334c))

<a name="v0.0.24"></a>
## [v0.0.24] - 2020-03-26
### Bug Fixes 🐛
- Add retry loop for secret client if initial token is invalid ([#60](https://github.com/edgexfoundry/go-mod-bootstrap/issues/60)) ([#ecac4d1](https://github.com/edgexfoundry/go-mod-bootstrap/commits/ecac4d1))

<a name="v0.0.13"></a>
## [v0.0.13] - 2020-02-04
### Bug
- **config:** Embedded types do not work with package we use to pull from Consul ([#38](https://github.com/edgexfoundry/go-mod-bootstrap/issues/38)) ([#2d9fcd4](https://github.com/edgexfoundry/go-mod-bootstrap/commits/2d9fcd4))

<a name="v0.0.12"></a>
## [v0.0.12] - 2020-01-31
### Code Refactoring ♻
- **registry:** Integrate new Configuration & Registry clients ([#915c058](https://github.com/edgexfoundry/go-mod-bootstrap/commits/915c058))

