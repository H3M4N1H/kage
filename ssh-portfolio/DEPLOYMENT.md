# Deployment guide (manual; not performed)

Use an Ubuntu LTS VPS or an existing Oracle Cloud VM. Confirm the provider's current capacity, billing labels, resource quotas, and network rules before creating any resources. No claim of free capacity is made here. Keep the static website separately hosted.

## Install the service

Build for the target architecture as an ordinary user. The supplied binary is Linux amd64; for an ARM64 VM, build with `GOOS=linux GOARCH=arm64 CGO_ENABLED=0 make build`.

On the target, create a dedicated account and install the compiled binary and supplied unit:

```sh
sudo useradd --system --home /var/lib/ssh-portfolio --shell /usr/sbin/nologin ssh-portfolio
sudo install -d -o root -g root -m 0755 /opt/ssh-portfolio
sudo install -o root -g root -m 0755 bin/ssh-portfolio /opt/ssh-portfolio/ssh-portfolio
sudo install -m 0644 deploy/ssh-portfolio.service /etc/systemd/system/ssh-portfolio.service
sudo install -m 0644 deploy/portfolio.env.example /etc/ssh-portfolio.env
sudo systemd-analyze verify /etc/systemd/system/ssh-portfolio.service
sudo systemctl daemon-reload
sudo systemctl enable --now ssh-portfolio
sudo systemctl status ssh-portfolio
sudo journalctl -u ssh-portfolio --since today
```

Start on loopback, as supplied. Compare the fingerprint and connect locally. Test forbidden commands and transfers before exposing the service. If a sandbox directive fails on the target kernel, diagnose it; do not silently remove hardening. This build did not run these privileged installation commands.

The unit lets systemd create `/var/lib/ssh-portfolio` with owner `ssh-portfolio`, mode 0700. Only that directory is writable apart from the private temporary directory. The host key persists across restarts.

## Public networking and host trust

After local verification, explicitly set `PORTFOLIO_ADDR=0.0.0.0:2222` in `/etc/ssh-portfolio.env`, restart, and allow inbound TCP 2222 in both provider and host firewalls. Configure IPv6 deliberately if needed. Keep administrative SSH on its own account with key-only authentication and restricted source addresses where practical. Do not disable the firewall or share admin keys with visitors.

On a host already managed with UFW, review `sudo ufw status verbose` and the existing admin rule before adding `sudo ufw allow 2222/tcp`. Do not enable or replace firewall policy blindly over an SSH session.

Point a real `ssh.<your-domain>` A/AAAA record at the verified public address. The kit's `ssh.hemanth.dev` is only an example; no ownership was established. Once the actual domain is known, add the connection command and host-key fingerprint to the HTTPS website.

Inspect the server key locally:

```sh
sudo ssh-keygen -lf /var/lib/ssh-portfolio/host_ed25519
```

Publish the fingerprint, never the private key. Clients should compare the first-connection fingerprint using that trusted HTTPS channel. For unattended health checks use a pinned known-hosts file and strict checking:

```sh
ssh -T -p 2222 -o BatchMode=yes -o ConnectTimeout=5 -o StrictHostKeyChecking=yes visitor@YOUR_CONFIRMED_HOST
```

No password is required by the application; the host key still must be checked. Do not use `StrictHostKeyChecking=no`.

## Operations

- Stop: `sudo systemctl stop ssh-portfolio`.
- Disable on boot: `sudo systemctl disable ssh-portfolio`.
- Restart after a binary/config update: `sudo systemctl restart ssh-portfolio`.
- Inspect: `sudo journalctl -u ssh-portfolio --since today`; configure journald storage/retention for the host as appropriate.
- Back up the private host key only into access-controlled encrypted storage if key continuity is required; never commit it or place it in a public artifact.
- Rotation: stop the service, archive the old key securely, remove it from the state path, restart to generate a new key, publish and independently verify the new fingerprint. Explain the change to visitors before they update their known-hosts entry. For a compromised key, retire it immediately and announce the incident through a trusted channel.
