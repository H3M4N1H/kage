# One-shot build prompt — Hemanth's personal portfolio

You are the lead product designer, senior frontend engineer, accessibility reviewer, and QA engineer for this task.

Build a complete, polished, responsive personal portfolio for **P. Hemanth Gupta** in the current repository. Follow `SKILL.md`, `PORTFOLIO_CONTENT.md`, and `DESIGN_SPEC.md` as the project brief and source of truth.

## Main objective

Create a portfolio that combines:

1. **CybercomCTF-inspired direction:** Terminal-inspired Cybersecurity Brutalism + Dark SaaS UI.
2. **Subadevan-inspired direction:** Editorial Minimalism + Dark Technical Portfolio UI.

Do not copy either website. Extract the design principles and create an original identity for Hemanth.

The visual balance should be approximately:
- 65% editorial minimalism: typography, whitespace, grid, storytelling, restrained composition.
- 35% technical brutalism: terminal details, monospace labels, system metadata, data panels, technical motifs.

The website should look credible to cybersecurity professionals, CTF organizers, recruiters, developers, and technical communities.

## Before coding

- Inspect the existing repository and identify the framework, package manager, scripts, routes, assets, and conventions.
- Preserve existing working functionality and configuration.
- Do not assume the project is empty.
- Do not ask for confirmation for routine design or implementation decisions. Make sensible choices and proceed.
- If an essential value such as the final domain or contact email is unknown, omit it or use a clearly documented editable placeholder. Do not fabricate it.

## Content requirements

Use `PORTFOLIO_CONTENT.md` exactly as the factual source. Build the site around:
- Cybersecurity and CTF identity.
- Computer Science & Engineering education.
- Hands-on projects and technical interests.
- CTF placements and competition participation.
- Community, club, design, and founder roles.
- Public GitHub, LinkedIn, TryHackMe, and CTFtime profiles.

Write concise, confident, natural copy. Avoid exaggerated claims, buzzword stuffing, and fake corporate language. Do not invent metrics, certifications, internships, employment, client outcomes, or project capabilities. If a project status is uncertain, describe it as a project or prototype without claiming production adoption.

## Required experience

### Hero
- Strong typographic name treatment.
- Clear positioning line, e.g. "Cybersecurity enthusiast. CTF player. Builder."
- One short introduction that says what Hemanth explores and builds.
- Primary CTA to selected work and secondary CTA to contact/social profiles.
- Small technical metadata row such as location at city/country level only if desired, discipline, or availability—but do not invent availability.
- A restrained terminal/system motif, not a giant fake terminal.

### About
- A short first-person or third-person introduction.
- Explain the intersection of cybersecurity, CTF problem-solving, and building useful tools.
- Mention current CSE (Cyber Security) education accurately.

### Focus areas / toolkit
- Group topics into sensible categories (e.g. Security, Development, Systems & Tools).
- Distinguish skills from tools and interests.
- Do not imply expert-level proficiency in every listed tool.
- Avoid arbitrary skill percentage bars.

### Projects
Create visually rich project cards for relevant projects in `PORTFOLIO_CONTENT.md`, prioritizing:
- CTF-AI
- SafeKey
- SpecFit
- GHOST
- Yugma / other projects only if they add value and the content is sufficiently clear.

Each project card should include:
- Name
- One-line purpose
- Short explanation
- Relevant stack/tags, only when supported
- Status only when known
- A real repository/demo link only if supplied or discoverable in the repository

Never invent repository URLs or screenshots. Use tasteful CSS project covers or diagrams when real media is unavailable.

### Achievements / CTF
- Feature verified placements, including CryptClash CTF third place and VIT Cyberquest third place, subject to the verification note in the content file.
- Present competition achievements in a clean timeline, result grid, or editorial list.
- Do not invent dates, team sizes, score counts, or rankings beyond supplied facts.

### Roles / community
- Include relevant roles in Cybercom, Cyventura, and NodeKraft based on the supplied content.
- Make the distinctions clear: club/community role, design contribution, and agency founder role.
- Do not imply that a community role is paid employment.

### Contact
- Use only the supplied public profile links.
- Include a clear contact CTA.
- If an email address is not provided, do not make one up.
- Do not create a fake contact form.

## Visual system

Use the design tokens in `DESIGN_SPEC.md`. The intended visual language:
- Background: near-black, not pure #000 everywhere.
- Surfaces: subtle charcoal panels with restrained borders.
- Text: warm white and muted grey.
- Accent: muted lime/acid green, used sparingly.
- Typography: confident geometric sans for headings, readable sans for body, monospace for metadata.
- Layout: editorial grids, asymmetric composition, deliberate whitespace, strong section rhythm.
- Details: fine rules, small section numbers, labels, compact metadata, occasional terminal prompt motifs.

Do not use:
- Matrix rain or animated code walls.
- Neon overload.
- Stock photos of hooded hackers.
- Random skulls, binary backgrounds, or meaningless code.
- Excessive glassmorphism.
- Fake dashboards, fake terminal commands, fake statistics, or fake activity indicators.
- Huge 3D scenes that harm performance or distract from content.

## Interaction and motion

- Smooth anchor navigation with accessible behavior.
- Subtle entrance transitions and hover feedback.
- Functional mobile navigation.
- Optional small interactive detail (e.g. copy profile link) only if it genuinely improves usability.
- Respect reduced-motion preferences.
- Avoid scroll hijacking, excessive parallax, and animation that blocks reading.

## Responsive and accessibility requirements

- Design mobile-first and verify narrow screens.
- No horizontal overflow.
- Navigation must work with keyboard and touch.
- Semantic landmarks and heading order.
- Visible focus indicators.
- Accessible names for icon-only controls.
- Sufficient contrast.
- Reduced-motion support.
- Images must have meaningful alt text or be decorative appropriately.

## SEO and metadata

- Add a concise page title and description.
- Add Open Graph and social preview metadata.
- Add favicon/app icon only if a suitable asset exists or can be created from a simple original monogram.
- Add sitemap and robots configuration if supported by the framework.
- Do not invent a canonical domain. Use the real configured domain only if found in the repository; otherwise leave canonical configuration out and report it.
- Keep content indexable; do not hide core content behind client-only interactions.

## Implementation rules

- Use the existing stack and package manager if the repository already has them.
- Prefer reusable components and data-driven content.
- Avoid adding dependencies for effects that CSS can handle.
- Do not replace the entire project configuration without a clear technical reason.
- Do not commit, push, deploy, purchase domains, or modify external services.
- Do not expose secrets or add credentials.
- Ensure all visible buttons and links do something meaningful.
- If there is no backend, do not fake backend behavior.

## QA / final verification

After implementation:
1. Run the relevant lint command.
2. Run type checking if configured.
3. Run the production build.
4. Fix all errors introduced by your work.
5. Check for broken internal links and placeholder content.
6. Check responsive behavior at mobile, tablet, and desktop widths if browser tooling is available.
7. If browser tooling is unavailable, state that visual browser QA was not performed; do not claim it was.

## Final response format

Report:
- What was built.
- Main routes/sections.
- Key design decisions.
- Important files changed.
- Commands executed and their results.
- Any factual content requiring Hemanth's confirmation.
- Any remaining manual steps (e.g. add final domain, email, project URLs, or real screenshots).

Now inspect the repository and build the complete portfolio. Do not stop after producing a plan or a mockup.
