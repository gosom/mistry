# Changelog

Breaking changes are prefixed with a "[BREAKING]" label.

## master (unreleased)

### Added

- [server] `mistryd --help` now displays the available filesystem adapters [[628ff12](https://github.com/skroutz/mistry/commit/628ff120062599ddb5bb0f2d41cc4d2ae47890ab)]

### Changed

- [server] Removed debug logs coming from the web view [[28e9743](https://github.com/skroutz/mistry/commit/28e97433293fdddbf62089c1514bb15c7efbd829)]


### Fixed

- [server] We would erroneously consider failed builds as successful, which resulted in some builds starting with cold caches instead of being incremental [[ab5ba18](https://github.com/skroutz/mistry/commit/ab5ba18b59ffd579834abd69e83c756263e4c858)]
- [client] The client now accepts dynamic arguments in the form of `--foo bar` (in addition to `--foo=bar`). Previously, it would panic [[f209061](https://github.com/skroutz/mistry/commit/f209061cd16274e4a198ec7d3c8be05718874b93)]
- [client] If the path passed to `--target` did not exist, it was erroneously created as a file [[1bfdeb4](https://github.com/skroutz/mistry/commit/1bfdeb4fccab06910be760d90d8bdef246fb4a3f)]





## 0.1.0 (2018-10-01)

### Added

- Support for opaque parameters [[#97](https://github.com/skroutz/mistry/pull/97)]
- server: Version flag `--version/-v` [[5c20927](https://github.com/skroutz/mistry/commit/5c209278bd6bf1032a1958eb252098b9e1ae228a)]


### Fixed

- server: Synchronize filesystem operations when symlinking [[502a42b](https://github.com/skroutz/mistry/commit/502a42b)]
- server: Errors on the build bootstrap phase would not abort the build [[828eddc](https://github.com/skroutz/mistry/commit/828eddc)]
- server: Socket FDs to docker were never closed [[b079128](b079128c018f145f013a5a2f2e3a51cfe37926e3)]
- webview: improve render performance [[#76](https://github.com/skroutz/mistry/issues/76)]

### Changed

- server: build info contains information about build errors [[7a3427](https://github.com/skroutz/mistry/commit/7a3427)]
- server: build info contains information about build cache usage [[93fd733](https://github.com/skroutz/mistry/commit/93fd733)]
- server: build info contains information about group [[5ff4cb1](https://github.com/skroutz/mistry/commit/5ff4cb1)]
- server: build info contains information about build time [[65b3ef2](https://github.com/skroutz/mistry/commit/65b3ef2)]





## 0.0.2 (2018-05-15)

### Added

- client: Output container stderr on non-zero exit code [[#85](https://github.com/skroutz/mistry/pull/85)]
- client: Add a `--timeout` option to specify maximum time to wait for a job [[#81](https://github.com/skroutz/mistry/pull/70)]
- server: Introduced a configuration option to limit the number of concurrent builds [[73c44ec](https://github.com/skroutz/mistry/commit/73c44ecc924260ccf61bad220eb26cd51a1f30d6)]
- server: Add `--rebuild` option to rebuild the docker images of a selection of projects ignoring the image cache [[#70](https://github.com/skroutz/mistry/pull/70)]
- client: Add `--rebuild` option to rebuild the docker image ignoring the image cache [[#70](https://github.com/skroutz/mistry/pull/70)]
- client: Add `--clear-target` option to clear target path before fetching
  artifacts [[#63](https://github.com/skroutz/mistry/pull/63)]
- client: Build logs are now displayed when in verbose mode [[#65](https://github.com/skroutz/mistry/pull/65)]
- Asynchronous job scheduling [[#61](https://github.com/skroutz/mistry/pull/61)]
- Web view [[#17](https://github.com/skroutz/mistry/pull/17)]

### Changed

- **[BREAKING]** server: failed image builds are now always visible as ready [[#75](https://github.com/skroutz/mistry/issues/75)]
- server: Job parameters are not logged, making the logs less verbose
- **[BREAKING]** Failed build results are no longer cached [[#62](https://github.com/skroutz/mistry/pull/62)]
- **[BREAKING]** client/server: Client and server binaries are renamed to "mistryd" and "mistry" respectively.
  Also project is now go-gettable. [[abbfb58](https://github.com/skroutz/mistry/commit/abbfb58d5a2aaf3eaebf9408d81ec7d459326416)]
- client: default host is now 0.0.0.0

### Fixed

- Don't delete build results on docker image build failure [[#75](https://github.com/skroutz/mistry/issues/75)]
- If a container with the same name exists, we remove it so that the new container
  can run [[#20](https://github.com/skroutz/mistry/issues/20)]
- Streaming log output in web view might occassionally hang [[7c07ca1](7c07ca177639cd6be7f9a860fb39c01370f35779)]

## 0.0.1 (2018-04-12)

First release!
