---
name: hemanth-portfolio-builder
description: Build, refine, and validate Hemanth's personal cybersecurity portfolio using a hybrid of terminal-inspired cybersecurity brutalism and dark editorial minimalism. Use when creating or modifying this portfolio.
---

# Hemanth Portfolio Builder

## Mission

Build a distinctive, production-quality personal portfolio for Hemanth (P. Hemanth Gupta), a computer science student focused on cybersecurity, CTF competitions, security projects, and software development.

The visual direction is a deliberate hybrid:
- Terminal-inspired Cybersecurity Brutalism and dark SaaS UI.
- Editorial Minimalism and dark technical portfolio design.

The result must feel like the personal site of a serious emerging cybersecurity practitioner—not a generic developer template, gaming site, hacker cliché, or SaaS product landing page.

## Operating procedure

1. Inspect the repository before changing anything:
   - Identify framework, package manager, existing routes, assets, design system, and deployment configuration.
   - Read `README.md`, package manifests, and existing project instructions.
   - Preserve working project conventions unless there is a clear reason to improve them.
2. Make a concise implementation plan, then execute it without waiting for approval unless a destructive or genuinely blocking decision is required.
3. Build the complete responsive site, not just a hero mockup.
4. Use the supplied `PORTFOLIO_CONTENT.md` as the source of truth for personal content.
5. Use `DESIGN_SPEC.md` for the visual system and interaction direction.
6. Verify the implementation:
   - Run the project's available lint, typecheck, and build commands.
   - Fix errors caused by your changes.
   - Check all routes, links, responsive states, and console/build output where tools permit.
7. Finish with a concise report: what changed, important files, commands run, verification results, and any remaining manual steps.

## Content integrity rules

- Do not invent employers, internships, certifications, rankings, dates, awards, client names, testimonials, project metrics, vulnerability disclosures, or professional experience.
- Do not describe Hemanth as a professional penetration tester, security researcher, or engineer unless the supplied facts support that exact claim. Prefer accurate wording such as "Cybersecurity Enthusiast", "CTF Player", and "Computer Science & Engineering student".
- Treat technical topics as interests or areas of practice unless demonstrable project evidence supports stronger claims.
- Keep dates and achievements editable and easy to update.
- If a detail is uncertain, omit it or mark it as a clearly editable placeholder in the content file—not as a visible fake fact.
- Do not publish personal phone numbers, home addresses, private credentials, API keys, student IDs, or other sensitive data.
- Use the current GitHub identity `H3M4N1H` unless repository evidence or the user explicitly updates it. Do not silently use the older `Dreadonyx` identity.
- Use the provided public profile URLs only. Do not guess email addresses or social handles.

## Design principles

- Prioritize hierarchy, spacing, typography, and composition over decoration.
- Use near-black backgrounds, restrained surfaces, high-contrast typography, muted secondary text, and one restrained accent.
- Use terminal-inspired details as accents: command prompts, metadata labels, status indicators, section indices, monospace captions, and technical annotations.
- Avoid full-page terminal cosplay.
- Avoid Matrix rain, excessive neon, random code rain, skulls, hooded hacker stock photos, fake terminal output, gratuitous glitch effects, and generic AI-generated cyber imagery.
- Use real project screenshots or supplied assets when available. If unavailable, create clean CSS-based project covers or restrained abstract diagrams; never pretend a mock visual is a real screenshot.
- Motion must be purposeful, subtle, performant, and respect `prefers-reduced-motion`.
- The mobile experience is a first-class design target, not a compressed desktop layout.

## Recommended information architecture

1. Hero / introduction
2. About
3. Focus areas / technical toolkit
4. Featured projects
5. CTF achievements and competition experience
6. Roles and community involvement
7. Learning / current focus (optional, only if useful)
8. Contact and social links

Use anchor navigation and clear section hierarchy. Do not pad the site with empty sections.

## Technical quality

- Reuse the repository's existing framework and package manager where practical.
- If starting from scratch, prefer Next.js + TypeScript + Tailwind CSS unless project constraints suggest otherwise.
- Use semantic HTML, accessible labels, keyboard navigation, visible focus states, and adequate contrast.
- Optimize images, avoid unnecessary dependencies, and prevent layout shifts.
- Add page metadata, Open Graph metadata, canonical URL only if the actual domain is known, sitemap/robots where supported, and meaningful headings.
- Do not add analytics, tracking, external APIs, databases, authentication, or contact-form backends without explicit need and configuration.
- A contact form must not pretend to submit if no backend exists. Prefer direct email/social links unless a real submission path is configured.
- Use working links only. External links should be clearly identifiable and safe.
- Keep components maintainable and content data-driven where appropriate.

## Completion criteria

The work is complete only when:
- All core sections are implemented with credible content.
- Desktop, tablet, and mobile layouts are considered.
- Navigation and interactions work.
- No visible fake metrics or invented credentials remain.
- Lint/typecheck/build checks have been run where available.
- The final report accurately states what was and was not verified.
