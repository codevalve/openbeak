# OpenBeak Inking (OPSEC Levels) 🐙🛡️

The `openbeak scan` engine uses a "Stealth-First" philosophy. To manage the signal-to-noise ratio during a reconnaissance operation, OpenBeak categorizes findings into severity levels and filters them before they reach the operator's display or log files.

## Inking Levels

You control the verbosity using the `--inking` (or `-i`) flag:

| Mode | Minimum Severity | Hunting Logic | Operator Outcome |
| :--- | :--- | :--- | :--- |
| **`stealth`** (Default) | **High** | Only probes confirmed OpenClaw signatures (e.g., specific version headers). | **Maximum OPSEC:** Only guaranteed hits for "The Claw" are inked. |
| **`tactical`** | **Medium** | Probes for exposed OpenClaw endpoints or configurations that act as perimeter indicators. | **Breached Perimeter:** Finds OpenClaw + bad configs. |
| **`verbose`** | **Low** | Probes everything (Generic Nginx/Apache servers, 401/403 endpoints). | **Complete Recon:** High noise, but full context on the target's attack surface. |

## How Findings are Scored
Hunters internally score their findings from `Low` to `High`. For example, the `http_discovery` tentacle behaves as follows:
- **High:** Confirmed `X-OpenClaw-Version` header detected.
- **Medium:** Specific sensitive endpoints (e.g., `/skills/list`) are accessible (200 OK).
- **Low:** Generic web services (e.g., 200 OK on `/`) or protected endpoints (401/403). OpenBeak extracts the `Server` header for these to aid in generalized reconnaissance.

## Configurable Output Files
By default, OpenBeak outputs to the current directory. You can configure the output paths dynamically for automated ingestions:

```bash
openbeak scan --inking verbose --out-json /tmp/results.json --out-log /tmp/activity.log
```
