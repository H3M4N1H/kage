# Hemanth portfolio builds

Both Downloads kits have been implemented here as separate local projects. Original briefs are in `briefs/`. Nothing has been deployed, committed, or pushed.

## Personal portfolio

```sh
cd personal-portfolio
npm run lint
npm run build
npm start
```

Open http://127.0.0.1:4173. The production output is `personal-portfolio/dist/`; it can be served by any static host. No npm dependencies are required. The build generates indexable HTML from editable JSON, CSS, and a small project-filter script. All core content remains visible without JavaScript.

Sections: introduction, About with a short education line, skills, 27 projects, Events with 13 placements and expandable CTF history, experience, and contact. The first nine projects appear initially; filters or Show all reveal the rest. The website restores the preferred neon dashboard: three-column hero, uppercase name, read-only terminal profile, interactive 3D globe, skill tabs, and thumbnail project cards. Decorative project visuals are CSS/SVG illustrations. See `docs/DESIGN.md` for asset sources and design details.

Main files: `personal-portfolio/src/content.json`, `src/style.css`, `src/site.js`, and `scripts/build.mjs`. Metadata, favicon and robots are included. A canonical URL, absolute social preview URL and sitemap are deliberately deferred until the real domain is known. There is no TypeScript configuration, backend, analytics, or fake contact form.

### Vercel deployment

The repository includes a root `vercel.json` configured for the free Vercel plan. Import this repository into Vercel, keep the project root at the repository root, and deploy. Vercel runs `npm --prefix personal-portfolio run build` and serves `personal-portfolio/dist/` as a static site. No server, database, or paid Vercel feature is required.

The `ssh-portfolio` service is separate and cannot run on Vercel because it requires a persistent SSH process; deploy that service to a VPS or another long-running host instead.

## SSH portfolio

See [local instructions](ssh-portfolio/README.md), [deployment guide](ssh-portfolio/DEPLOYMENT.md), and [security controls](ssh-portfolio/SECURITY.md).

The built binary is `ssh-portfolio/bin/ssh-portfolio` (Linux amd64). Run it from the `ssh-portfolio` directory, then connect with `./connect.sh` after checking the locally generated host-key fingerprint. The default bind is loopback only.

## Review

See [content requiring confirmation](docs/CONTENT_REVIEW.md) and [verification results](docs/VERIFICATION.md). Deployment and systemd activation are manual follow-up work; no public endpoint was created.
