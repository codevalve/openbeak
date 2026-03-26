# SIEM Integration Handbook

This document outlines the requirements and blueprints for integrating OpenBeak findings with Enterprise SIEM/SOAR platforms.

## Logging Formats

### 1. CEF (Common Event Format)
Used by ArcSight and many other SIEMs.
**Structure:** `CEF:Version|Device Vendor|Device Product|Device Version|Signature ID|Name|Severity|[Extension]`

### 2. LEEF (Log Event Extended Format)
Used by IBM QRadar.
**Structure:** `LEEF:1.0|Vendor|Product|Version|EventID|`

### 3. GELF (Graylog Extended Log Format)
JSON-based, used by Graylog and ELK stacks.

## v1.1.0 Goals
- [ ] Implement `CEFInk` for `internal/tentacles`.
- [ ] Implement `SyslogInk` (RFC 5424).
- [ ] Support custom TCP/UDP sinks for live streaming.

## Alerting mapping
| OpenBeak Severity | SIEM Severity (0-10) |
|-------------------|----------------------|
| High              | 9-10                 |
| Medium            | 5-6                  |
| Low               | 2-3                  |
