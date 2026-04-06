# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## 0.3.1 (2026-04-06)
### Patch Changes
- chore: Publishing tags compatible with go modules

## 0.3.0 (2026-04-06)
### Minor Changes
- feat: Patch CreateNewSite to return correct API response ([#13](https://github.com/Tohaker/omada-go-sdk/pull/13) by @Tohaker)

## 0.2.1 (2026-04-05)
### Patch Changes
- ci: Adding changeset management process and using makefile to build project ([#9](https://github.com/Tohaker/omada-go-sdk/pull/9) by @Tohaker)

## 0.2.0 (2026-04-03)
### Added
- Documentation is now available at https://tohaker.github.io/omada-go-sdk

### Changed
- SDK is now generated with `openapi-generator-cli` v7.21.0
- Multipart requests now defer the closing of files when calling the `addFile` method.

## 0.1.0 (2026-03-29)
### Added
- Initial generation of Omada SDK from source
- Created Auth spec to fill the gaps in the original source spec
