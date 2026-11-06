# Security boundary and threat model

## Assets and trust

The public assets are the embedded profile text and TUI. The private asset is the persistent SSH host key. The operating-system account, filesystem, and administrator SSH service must remain inaccessible to visitors. Visitor usernames, commands, terminal metadata, channel types, and all SSH messages are untrusted.

Only an SSH `session` channel is registered. Each connection gets at most one session. A channel request filter allows only `shell` (which starts the application, not a system shell), bounded `pty-req`, and bounded `window-change`. It explicitly rejects exec, subsystem/SFTP/SCP, agent forwarding, X11, and environment requests. No forwarding channel or global request handlers are registered. There is no shell/subprocess execution or visitor-controlled file access in application code. PTYs are emulated by the SSH library; no OS login is created.

## Resource limits

| Control | Default |
| --- | --- |
| Concurrent connections, including handshakes | 32 |
| Concurrent connections per source IP | 2 |
| Connection attempts per source IP | 6 per fixed minute window |
| Rate-limit address entries | 4096 maximum; expired inactive entries removed on admission |
| Initial handshake / session-channel opening | 10 seconds |
| Idle inbound network activity | 2 minutes |
| Absolute connection lifetime | 10 minutes |
| Network input / output per connection | 256 KiB / 2 MiB |
| Write stall deadline | 10 seconds |
| Session requests per connection | 2048 |
| Session request payload | 4096 bytes |
| Accepted terminal dimensions | 1–500 columns, 1–200 rows |
| Rendered terminal dimensions | At most 160 × 60 |
| TUI rendering | 20 frames per second |

Byte counts include SSH framing and negotiation. Hitting a budget disconnects the client. Rate limiting is a fixed-window guard; it permits boundary bursts and is not a substitute for provider-level DDoS mitigation. NAT users share an IP budget. Distributed connections, TCP accept floods, upstream bandwidth saturation, and vulnerabilities in dependencies/OS remain risks.

Host-key files must be regular files with no group/other permissions. The key directory must be service-owned and non-writable by visitors or other users. File-path configuration is administrator-controlled. The service intentionally does not support dynamic content or plugins.

## Privacy

Application logs contain only startup, shutdown/fatal configuration errors. No visitor input, usernames, or IPs are logged. IP addresses are retained in bounded process memory for rate limiting; expired inactive entries are purged on subsequent admissions (and all disappear at process exit). Host/provider logs may have independent policies.

## Validation and remaining deployment checks

Automated tests exercise content validation, all menu choices, invalid selection, wrapping, resize, persistent/restrictive host keys, IP/global admission limits, output budget, disconnect cleanup, plain and interactive SSH, timeout handling, and denied exec, subsystem, environment, agent/X11, TCP channels and global forwarding.

This is targeted testing, not a security audit or proof against every SSH protocol attack. Before public deployment:

- Run the tests and a dependency vulnerability scan with the deployment toolchain.
- Test the systemd service on the target Ubuntu kernel, including `MemoryDenyWriteExecute`, permissions, graceful stop/restart, and key persistence. The unit is supplied but has not been installed or exercised as a system service here.
- Independently attempt OpenSSH exec, SFTP, SCP, local/remote forwarding, agent/X11 forwarding; verify denial with the deployed endpoint.
- Set provider ingress rules and a host firewall. Keep administrative SSH key-only and separate.
- Test resource limits under controlled load, and inspect memory/CPU/file-descriptor use.
- Keep Go, Wish, Bubble Tea, x/crypto and the OS patched; review dependency changes before deploying.
- Publish the host fingerprint on the actual HTTPS website. Never ask visitors to disable host-key checking.
