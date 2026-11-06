---
name: hemanth-ssh-portfolio
description: Build a secure, read-only SSH-accessible terminal portfolio for Hemanth using a custom SSH application, not a normal shell account.
---

# Hemanth SSH Portfolio Skill

## Objective
Build a public, read-only SSH portfolio that visitors can access with a command such as:

`ssh -p 2222 visitor@ssh.hemanth.dev`

The service must display an interactive terminal UI for Hemanth's profile, projects, CTF results, roles, skills/focus areas, and public links.

## Non-negotiable security model
- Do NOT create a normal shared Linux user with an interactive shell for visitors.
- Do NOT expose `/bin/bash`, `/bin/sh`, SFTP, SCP, port forwarding, agent forwarding, X11 forwarding, or arbitrary command execution.
- Prefer a custom SSH application server using Go + Charmbracelet Wish + Bubble Tea (or an equivalent library) that authenticates visitors anonymously at the application layer and serves only the TUI.
- The SSH daemon/application must never pass visitor input to `exec`, a shell, subprocesses, SQL, or a filesystem path.
- The app should read only bundled, static portfolio content.
- Run as a dedicated unprivileged service account with no sudo privileges and minimal filesystem access.
- Bind only the required public SSH port. Keep administrative SSH separate and key-only.
- Add connection limits, per-IP rate limiting, maximum session duration, idle timeout, bounded concurrent sessions, and graceful disconnect handling.
- Keep server host keys persistent and protect their permissions.
- Use a firewall to expose only required ports.
- Do not store visitor input or IP addresses longer than operationally necessary. Do not collect personal information.
- Do not include secrets, credentials, private files, local paths, or personal contact details not approved for public release.
- No file upload/download features.

## Workflow
1. Inspect the repository and existing stack before making changes.
2. If the repo is a web portfolio, keep the SSH app as a separate service/subdirectory; do not destabilize the website.
3. Read `PORTFOLIO_DATA.md` and `ARCHITECTURE.md`.
4. Implement the TUI with a clear menu and keyboard navigation.
5. Implement the SSH transport with a custom app server (recommended: Go + Wish + Bubble Tea).
6. Add deployment assets: systemd unit, environment/config example, firewall notes, health-check command, and hardening guide.
7. Add tests for menu navigation, invalid input, long text, terminal resize, disconnected clients, and session limits.
8. Run formatting, linting, tests, and build. Do not claim a deployment was performed unless it was.
9. Report exact commands, files, security boundaries, and remaining setup.

## UX requirements
- Welcome banner and short intro.
- Menu: About, Focus Areas, Projects, CTF Achievements, Community & Roles, Profiles, Contact, Help/Exit.
- Keyboard navigation: number keys, arrow keys, Enter, Esc/q.
- Show the visitor's command to open the website if the web domain is confirmed.
- Make terminal output readable at 80 columns and adapt to terminal resize.
- Provide a plain-text fallback for limited terminals.
- Never rely on color alone; support no-color mode.
- Do not fake live system status, user counts, uptime, scan results, or activity.
- Provide a short privacy/security notice.

## Content integrity
Use `PORTFOLIO_DATA.md` as source of truth. Do not invent project URLs, achievements, certifications, employers, metrics, or dates. Treat uncertain details as editable/omittable. The GitHub username is H3M4N1H unless updated by repository evidence or user instruction.

## Completion
The result must be a working local SSH app, tested locally with an SSH client, with deployment instructions. Do not provision cloud resources, create accounts, alter DNS, open public firewall ports, or deploy without explicit instruction.
