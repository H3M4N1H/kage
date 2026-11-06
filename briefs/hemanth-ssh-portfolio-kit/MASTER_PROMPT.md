# One-shot Codex prompt — Hemanth's SSH portfolio

Act as a senior Go engineer, terminal UX designer, and application security engineer. Build a secure, read-only SSH-accessible terminal portfolio for Hemanth in this repository.

Read and follow:
- `SKILL.md`
- `PORTFOLIO_DATA.md`
- `ARCHITECTURE.md`

## Goal

A visitor runs:

`ssh -p 2222 visitor@ssh.hemanth.dev`

They enter a custom interactive portfolio TUI, not a Linux shell. It should feel polished and consistent with Hemanth's website: dark technical editorial design, terminal-inspired details, restrained lime accent, clean typography, clear navigation.

## Critical architecture choice

Prefer a custom SSH application server using Go + Charmbracelet Wish + Bubble Tea (and Lip Gloss for styling). It must handle SSH sessions directly and render only the portfolio TUI. Do not expose an operating-system shell account.

If the repository already has a strong compatible architecture, explain the trade-off and preserve it where sensible. Do not use OpenSSH ForceCommand as the default for anonymous public visitors if it would require distributing a shared private key or password. Never create a shared visitor credential.

## Required menu

1. About
2. Focus areas / technical toolkit
3. Featured projects
4. CTF achievements
5. Community and roles
6. Public profiles
7. Contact
8. Help / exit

Navigation:
- Number keys and arrow keys.
- Enter to select.
- Esc/backspace to go back where appropriate.
- `q` to quit.
- Handle resize and narrow terminals.
- Provide a simple non-color fallback.

## Security requirements — mandatory

- Anonymous visitor access is allowed only inside the custom application protocol; no OS login shell.
- No shell, arbitrary command execution, subprocess invocation, file access chosen by the visitor, SFTP, SCP, port forwarding, agent forwarding, X11 forwarding, or upload/download.
- Never interpolate visitor input into shell commands, SQL, templates, or filesystem paths.
- Bundle static portfolio data or load a fixed read-only data file from a locked path.
- Run under a dedicated unprivileged Linux account with no sudo and minimal read permissions.
- Set max sessions, per-IP connection throttles, handshake timeout, idle timeout, max session duration, and output limits.
- Use persistent host keys with restrictive permissions and document host-key rotation.
- Avoid logging raw visitor input. Keep only minimal operational logs.
- Bind to the configured interface/port; do not expose an admin listener.
- Document firewall rules. Do not alter host firewall automatically.
- Do not hardcode host private keys, passwords, cloud credentials, or API tokens.
- Do not claim security guarantees that have not been tested.

## Engineering requirements

- Inspect the repo and toolchain first.
- Use Go modules and pinned dependency versions.
- Keep content separate from UI and transport.
- Make components testable.
- Add unit tests for content loading, menu routing, invalid selections, rate limiter/session cap, and text wrapping.
- Add integration tests where feasible.
- Add a `Makefile` or clear scripts for `fmt`, `test`, `build`, `run`, and `check`.
- Add `.gitignore` for generated binaries, local host keys, logs, and secrets.
- Add a sample config with safe defaults and no secrets.
- Add a systemd unit running as an unprivileged service user with `NoNewPrivileges`, `PrivateTmp`, `ProtectSystem`, `ProtectHome`, `ProtectKernelTunables`, `ProtectKernelModules`, `RestrictSUIDSGID`, `LockPersonality`, `MemoryDenyWriteExecute` where compatible, and a tight `ReadWritePaths` list for the host-key directory only. Test the service restrictions and document incompatible directives rather than silently weakening them.
- Avoid running as root. If binding port 22 is required, prefer a high port such as 2222 and use a reverse proxy/port forwarding only if securely configured and documented.
- Provide a local run path and test with an SSH client if available.
- Do not deploy, create a cloud account, change DNS, open ports, or commit/push.

## Visual and UX direction

- Near-black background, warm white text, muted grey metadata, restrained lime accent.
- Monospace labels, section numbers, fine separators, simple box-drawing header.
- Avoid excessive ANSI animation, blinking text, rainbow gradients, fake live telemetry, Matrix rain, and hacker clichés.
- Be polished and legible at 80x24; adapt for 120x40.
- Respect terminal color capability and provide plain output fallback.

## Content

Use only facts in `PORTFOLIO_DATA.md`. Do not invent links, email, achievements, dates, certifications, or professional claims. Omit unverified details rather than publishing placeholders as facts. Public profile links should be printed in a copy-friendly form.

## Deliverables

- Working custom SSH TUI application.
- Static portfolio data file.
- Unit/integration tests.
- Config example.
- systemd unit.
- Deployment guide for an Ubuntu LTS VPS or Oracle Cloud Always Free VM.
- Security threat model and hardening checklist.
- README with local run and client connection instructions.
- Optional `ssh://` or `ssh -p ...` snippet for the website.

## QA

Run formatting, tests, build, and local SSH connection test if environment permits. Verify no visitor can obtain a shell or access files. If security testing cannot be completed, list it explicitly as outstanding.

Finish with a report: architecture, files, commands/results, security controls, deployment steps, and remaining manual inputs. Now inspect and implement the complete service; do not stop at a plan.
