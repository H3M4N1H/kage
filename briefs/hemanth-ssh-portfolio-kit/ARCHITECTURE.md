# Architecture and deployment notes

## Recommended design

Use a custom SSH application server (Go + Charmbracelet Wish + Bubble Tea). Visitors authenticate only into the application session and are shown the TUI. There is no underlying Linux user shell exposed to the public.

```text
Visitor terminal
   |  SSH (TCP 2222)
   v
Custom SSH application (Wish)
   |-- session limits / rate limiter
   |-- Bubble Tea terminal UI
   |-- read-only static content
   v
No shell, no SFTP, no arbitrary filesystem access
```

## Why not a regular Linux account?

A normal SSH account can expose a shell and filesystem. A shared password/key is difficult to distribute safely. A custom SSH application provides a purpose-built public experience without granting OS access.

## Suggested deployment

- Provider: Oracle Cloud Infrastructure Always Free eligible Ubuntu VM, if available in the selected home region.
- Service port: TCP 2222.
- DNS: `ssh.<your-domain>` A record pointing to the VM public IPv4 address.
- Website: keep separately hosted (e.g. Vercel/Cloudflare Pages); add a link/command to SSH experience.
- Service: systemd, dedicated unprivileged account.
- Admin SSH: separate account, key-only authentication, restricted source IP where practical.
- Firewall: allow TCP 2222 for visitors and admin SSH only as required. Never disable the firewall wholesale.

## Cost and provider caveats

Oracle documents Always Free compute, including up to two AMD micro VMs and an Ampere A1 allocation equivalent to 2 OCPUs and 12 GB RAM for Always Free tenancies. Capacity can be unavailable in a home region. Oracle also documents possible reclamation of idle compute instances under specified utilization conditions. Confirm the current console labels and account limits before creating resources. Do not create paid resources by accident.

Official references:
- https://docs.oracle.com/en-us/iaas/Content/FreeTier/freetier_topic-Always_Free_Resources.htm
- https://www.oracle.com/cloud/free/faq/

## Operational checklist

- Keep OS and Go dependencies patched.
- Persist SSH host keys; do not regenerate on every restart.
- Add host-key fingerprint to the website so visitors can verify the server.
- Limit concurrent sessions and session duration.
- Rate-limit connection attempts.
- Set memory/CPU limits via systemd where supported.
- Use journald with log rotation; avoid logging visitor-entered text.
- Monitor service health and disk usage.
- Back up only the static content/config required; never back up or publish private keys in Git.
- Document how to stop/disable the service and rotate host keys.

## Host key trust

On first connection, SSH clients display a host-key fingerprint prompt. Publish the fingerprint on the HTTPS portfolio website over a separate trusted channel. Do not tell visitors to blindly accept unknown keys.

## Deployment caution

The coding agent may generate deployment files and commands, but must not provision the VM, change DNS, modify cloud networking, or deploy without explicit user approval.
