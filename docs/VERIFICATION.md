# Verification — 2026-11-03

## Website

- `npm run lint`: passed (JavaScript syntax and content-schema checks).
- `npm run build`: passed; static output in `personal-portfolio/dist`.
- Playwright with local Chromium: passed at 360, 768, and 1440 pixels wide.
- Tested HTTP response, horizontal overflow, project-filter counts, navigation anchors, reduced-motion preference, no-JavaScript content, and JavaScript page errors.
- Desktop and mobile full-page screenshots inspected visually. Screenshots saved in `/tmp/portfolio-{360,768,1440}.png` during the build session.
- Regression script: `personal-portfolio/tests/browser.py`. Requires Python Playwright and Chromium; use `CHROMIUM_PATH` to select an existing browser. Start the local website before running.
- No TypeScript configuration; type checking is not applicable. Live external profile destinations were taken verbatim from the brief, not independently audited.

## SSH service

- Go 1.27.3 toolchain downloaded from go.dev; archive SHA-256 verified against release metadata.
- `go fmt ./...`, `go mod tidy`, `go vet ./...`: passed.
- `go test -race -buildvcs=false ./...`: passed, including local SSH integration tests.
- `go build -buildvcs=false -trimpath -o bin/ssh-portfolio .`: passed (Linux amd64).
- Real OpenSSH client with a known-hosts entry derived from the local server key: plain output passed, remote exec and SFTP requests rejected.
- Regression client script: `ssh-portfolio/tests/openssh.py`; run against the local service after its key has been generated.
- `systemd-analyze verify` passed for a temporary copy whose ExecStart points to the actual local binary. The shipped unit intentionally points to `/opt/ssh-portfolio/ssh-portfolio`, which does not exist until installation.
- Unit was not installed or runtime-tested under systemd. Target-host hardening and deployment smoke tests remain necessary.

The initial mobile overflow and SSH idle deadline failures were fixed and the relevant checks rerun successfully. No public deployment, firewall/DNS modification, cloud provisioning, Git commit, or push occurred.

## Neon reference redesign

The website was rebuilt to match the supplied neon cybersecurity dashboard reference. Lint and production build passed. Chromium checks passed at 360, 768, 1440 and 1672 pixels: no overflow or JavaScript errors, working skills tabs and arrow-key navigation, mobile navigation with Escape support, project filters, project deep links (including previously hidden projects), anchors, reduced motion and readable no-JavaScript content. Desktop and mobile screenshots were inspected; globe/fonts are served locally. The SSH service was unchanged.

## Terminal / Swiss UI and interactive globe

Lint and production build passed. Chromium checks at 360, 768, 1440, and 1672px cover globe dragging, keyboard rotation, zoom/reset, explicit rotation/pause, static fallback on data-load failure, navigation-prompt commands and unknown-command feedback, filters, repeated project deep links, mobile navigation, skills tabs, reduced motion, no-JavaScript content and layout overflow. No JavaScript page errors occurred. Screenshots were inspected and tablet skill cards expanded to two columns for readability. No SSH code was changed.

## Reference-inspired editorial update

Live visual inspection of Cybercom and Subadevan informed layout changes. Existing globe and navigation coverage is retained; the browser checks additionally cover Ctrl+K, Escape, opening the jump menu from its button, and selecting a section. Production build and JavaScript/content validation pass. Screenshots are captured at 360, 768, 1440 and 1672px. Sources and palette decisions are recorded in DESIGN.md.

## Restored preferred dashboard

The user selected the earlier compact neon dashboard. Restored the matching wording, uppercase hero, static profile window and thumbnail project layout; retained the interactive 3D globe. Removed the later command prompt/menu and editorial layout. Browser checks passed at 360, 768, 1440 and 1672px for globe controls/fallback, filters, deep links, tabs, mobile navigation, anchors, reduced motion, no-JavaScript content, and absence of horizontal overflow/page errors. The restored desktop screenshot was inspected. The globe's default visual scale was subsequently increased within its existing canvas to better match the preferred reference; lint and build passed again.

Window-padding correction: verified widths 320, 360, 768, 1024, 1180, 1280, 1366, 1440, and 1920px. Shared content-edge alignment, minimum gutters, zero horizontal overflow, hero/panel separation, and laptop hero/summary visibility passed. Content uses a centered 1440px maximum container with 32/24/16px responsive gutters.

Full-width alignment and About/Skills correction: tested 320, 360, 390, 600, 700, 768, 900, 1024, 1180, 1280, 1366, 1440, 1920 and 2560px. All major sections have matching left/right gutters; no tested horizontal overflow or text clipping; sequential sections do not overlap; navigation anchors remain below the sticky header. About and Skills now have distinct scroll positions at every width. Filters, skill tabs, and page-error checks passed.


Resume update (November 3, 2026): content validation and build passed; responsive alignment passed at 14 widths (320–2560px); browser navigation, filtering, skills and globe controls passed. Both résumé download buttons returned bytes identical to the supplied PDF. SSH tests passed after updating the featured-project expectation to Aegis.

Events update (November 4, 2026): build and content validation passed. All seven résumé placements, user-confirmed attendance notes, team-result labels, and internal anchors verified. Responsive alignment, header navigation, text clipping, project filters, and skill tabs passed at 14 widths from 320 to 2560 pixels.

Expanded profile update (November 4, 2026): 27 projects, 13 competition placements, and 32 CTF history entries. Content lint and build passed; responsive checks passed at 14 widths, including all projects and expanded histories. Browser checks passed for show-all, filters, links, navigation, skills, globe controls, and no-JavaScript content. SSH race tests and binary build passed. Confirmed both third-place achievements and the latest Cyventura / HiRes Lossless title corrections.

Globe interaction update: removed toolbar and profile summary strip. Automatic rotation pauses on click or drag and resumes on the next click. Keyboard controls remain available (Space/Enter toggle). Browser interaction checks and responsive alignment checks at 14 widths passed. Reduced-motion preference is respected.

Interactive terminal: content lint/build passed, terminal interaction tests passed at six widths (commands, history, completion, clear, safe text, links, PDF download, repeated project links after filtering, fixed-height scrolling, and fallback). Overall layout checks passed at 14 widths.

User requested reverting only the interactive terminal. Restored the previous static profile box, removed its command component and generated profile data, and retained the globe interactions and all earlier content/layout updates.

Completion pass — November 6, 2026: web content lint/build and full browser interaction checks passed after restoring the static profile box. SSH race/integration tests, go vet, and binary build passed. Added bold red h3m4n1h ASCII menu banner with compact fallback, tested menu selection and width/height bounds from 24x10 to 160x40. Local website and SSH preview servers started; public hosting remains unconfigured. Refreshed both built kits.

Centered padding update: content lint/build and alignment checks passed at 14 widths (320–2560px). Header, footer, and all sections share centered edges; no horizontal overflow or clipped text was found with projects and event history expanded.

Mobile hero correction: added 24px effective side insets, consistent vertical spacing, 48px actions, stacked buttons at 380px and below, and responsive globe sizing. Visually inspected at 390px. Alignment, text clipping, action sizing, and section edge checks passed across 14 widths.

Mobile opening-screen fix: grouped introduction and globe into a minimum viewport-height region beneath the header, with profile box following below. Seven mobile viewport checks confirm the box starts at least 20px below the first screen; 14-width alignment regression passed. Visually checked the 390x844 opening screen.

SSH client multiplexing diagnosis: successfully connected to the running local server with OpenSSH using -tt, ControlPath=none, ControlMaster=no, and ControlPersist=no; menu and ASCII banner rendered and q exited cleanly. Added ssh-portfolio/connect.sh to bypass reused client sockets without changing user settings.

November 13, 2026: added restrained one-time 380ms fade/rise reveals. Motion tests passed (no repeated reveals, reduced-motion cancellation, mobile first-screen layout); full browser interaction suite passed at four widths. Added the exact user-supplied block banner and remote color override; SSH race tests, vet, build, and responsive banner tests passed. Verified a real SSH connection emitted bold red (ANSI256 203) banner and green (155) highlights; q exited normally. run.sh starts a colored preview despite inherited NO_COLOR.

November 13, 2026 update: selected Aegis, SafeKey, SpecterNet and ranq with public repository links. GitHub REST API snapshot: H3M4N1H, 38 public repositories, 51 followers, 68 following. Current profile handles use h3m4n1h; historical competition aliases remain unchanged. Added NodeKraft (nodekraft.tech) and HiRes Lossless (hireslossless-team.in) role links. CTFtime blocked fresh requests; existing November 4 team results retained without claiming a new verification date. Section navigation scrolls and moves focus without adding URL fragments.

Directional motion refinement: replaced the subtle 12px/380ms reveals with 700ms left/right/up entrances, 20px travel on phones and 32px on desktop, plus 80ms stagger on paired project/experience cards. Inspired by the HiRes Lossless website’s 700ms scroll reveals. Motion and seven-size mobile hero checks passed; reduced motion, one-time playback, and clean URL navigation retained.

Motion polish: removed the opacity reset that could hide already-painted content when a scroll observer fired late. Entrances begin 96px before viewport entry and use transform-only 900ms slides with gentler easing. Unrelated focus/click events no longer cancel every entrance. Regression checks sample rendered frames for constant opacity and overflow and verify that unrelated focus does not cancel a running slide; reduced-motion behavior remains supported.
