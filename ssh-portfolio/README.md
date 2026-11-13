# Hemanth SSH portfolio

A custom, read-only SSH application using Wish, Bubble Tea, and Lip Gloss. Anonymous SSH authentication grants access to bundled portfolio content only. It does not create a Linux visitor account or invoke an OS shell.

## Local use

Requires the Go version in `go.mod` (1.27.3 or newer), Make, and an SSH client. Dependencies are pinned in `go.mod` / `go.sum`.

```sh
make check
make run
```

The compiled Linux amd64 binary can run without Go. The default bind is `127.0.0.1:2222`; the process refuses to run as root. The first run creates `.state/host_ed25519` with mode 0600. Keep it private and persistent.

In another terminal, inspect the **local** server's fingerprint, then compare it with the SSH prompt:

```sh
ssh-keygen -lf .state/host_ed25519
./connect.sh
```

Use 1–8 or arrows and Enter; Esc/backspace returns to the menu; arrows scroll section content; q/Ctrl+C exits. Resize is supported; very small terminals show a resize message. `ssh -T -p 2222 visitor@127.0.0.1` prints plain, copy-friendly text and disconnects. `TERM=dumb ssh ...` also selects plain text. Set the server's `NO_COLOR=1` for a colorless interactive TUI.

Edit `content.json` and rebuild to change the embedded content. There is no visitor-selectable file path and no runtime content directory. Current supplied roles and results are included; remaining optional links and hosting details are documented in `../docs/CONTENT_REVIEW.md`.

## Files and architecture

- `content.go`, `content.json`: embedded, validated static content.
- `ui.go`: menu routing, wrapping, scrolling, terminal resize and styling.
- `server.go`: persistent host key, SSH request allowlist, per-session UI.
- `limits.go`: admission control, per-IP rate limit, wire byte budgets and deadlines.
- `app_test.go`: unit and local SSH integration tests.
- `deploy/`: example environment file and hardened systemd unit.

The transport is based on [Wish](https://github.com/charmbracelet/wish); application-specific filters narrow its request handling to this portfolio.

See [deployment](DEPLOYMENT.md) and [threat model](SECURITY.md). No cloud resources, DNS, firewall rules, or services were changed by the build.

The main menu shows a user-supplied bold coral-red `h3m4n1h` block banner when the terminal is wide enough and at least 23 rows high. Smaller windows use a compact title so navigation remains visible. Section pages prioritize readable content. Plain SSH output includes an uncolored banner.

## Client connection reuse

Use `./connect.sh` for the local interactive preview. It disables SSH multiplexing for this connection and forces terminal allocation. The portfolio accepts one session per SSH connection, so clients that reuse a ControlMaster socket can report “Session open refused by peer.” The launcher bypasses that socket without changing your SSH settings or deleting it.

Equivalent command:

```sh
ssh -tt -o ControlPath=none -o ControlMaster=no -o ControlPersist=no -p 2222 visitor@127.0.0.1
```

For the colored local preview, start the server with `./run.sh`, then connect using `./connect.sh`. `PORTFOLIO_COLOR=always` overrides inherited `NO_COLOR` for remote sessions; `PORTFOLIO_COLOR=never` disables styling. `TERM=dumb` and non-interactive output remain plain.
