# Portfolio builds

This repository contains two projects:

- `personal-portfolio/` — static cybersecurity portfolio generated from editable JSON, CSS, and JavaScript.
- The repository root — an existing React/Vite Kage landing page.

## Personal portfolio

```sh
cd personal-portfolio
npm run lint
npm run build
npm start
```

The production output is `personal-portfolio/dist/`.

### Vercel deployment

The root `vercel.json` is configured for Vercel's free plan. Import this repository into Vercel, keep the project root at the repository root, and deploy. Vercel runs `npm --prefix personal-portfolio run build` and serves `personal-portfolio/dist/` as a static site.

The `ssh-portfolio/` service requires a persistent SSH process and cannot run on Vercel. Deploy it to a VPS or another long-running host instead.

## Kage landing page

The React/Vite app at the repository root can be run with:

```sh
npm install
npm run dev
```

Deploy it as a separate Vercel project if needed; the included Vercel configuration intentionally targets the static personal portfolio.

## SSH portfolio

See `ssh-portfolio/README.md`, `ssh-portfolio/DEPLOYMENT.md`, and `ssh-portfolio/SECURITY.md`.
