# Design specification

## Creative direction

**Hybrid label:** Dark Cybersecurity Brutalism × Editorial Minimalism × Technical Portfolio UI.

**Reference balance:** approximately 65% Subadevan-style editorial minimalism and 35% CybercomCTF-style terminal/technical brutalism. This is a creative direction, not a request to clone either reference.

## Design tokens

Suggested CSS variables:
```css
:root {
  --bg: #08090B;
  --surface: #101216;
  --surface-raised: #15191F;
  --border: rgba(255,255,255,.10);
  --text: #F1F2EE;
  --muted: #969BA3;
  --accent: #B5F36B;
  --accent-ink: #11150B;
  --radius-card: 14px;
  --content-width: 1160px;
}
```

Use the accent for selected links, small status marks, active navigation, and a few key details—not for every heading or panel.

## Typography

- Display/headings: Space Grotesk, Geist, or a similar modern geometric sans.
- Body: Inter, Geist Sans, or system sans.
- Metadata/code: JetBrains Mono or a system monospace.
- Use fluid type sizing with `clamp()`.
- Keep paragraphs readable and avoid excessively wide text measures.
- Use uppercase micro-labels sparingly with tracking.

## Composition

- Strong hero with oversized name and a compact positioning statement.
- Editorial section introductions with small index labels (01, 02, 03).
- Asymmetric two-column or split layouts where useful.
- Project cards with a consistent information hierarchy and restrained visual treatment.
- Fine dividers, subtle borders, and deliberate whitespace.
- A compact navigation bar that remains usable on mobile.
- Footer should be useful, not decorative filler.

## Terminal / cyber motifs

Use a small vocabulary of motifs:
- `root@hemanth:~` style prompt only as a decorative label, never as a fake live terminal.
- `[ SYSTEM / PROFILE ]` or section metadata.
- Monospace timestamps only when real.
- Small status indicators only when meaningful.
- Technical diagrams or project-specific visual abstractions where useful.

Never simulate a real terminal command result, network scan, system status, uptime, or security alert.

## Motion

- 150–300ms transitions for hover/focus.
- Small reveal animations are acceptable.
- No mandatory loading intro.
- No scroll hijacking.
- Respect `prefers-reduced-motion`.
- Avoid heavy WebGL/Three.js unless the existing site already uses it and it materially improves the result.

## Responsive behavior

- Mobile navigation must be keyboard and touch accessible.
- Hero typography must wrap cleanly.
- Project cards should collapse into a single column on narrow screens.
- No horizontal scrolling.
- Avoid tiny monospace text.
- Keep buttons and tap targets comfortable.

## Image direction

Prefer no generic hero stock image. Use typography, a simple monogram, abstract geometry, or original CSS details. Project screenshots should be genuine. If no genuine screenshot exists, create a clearly illustrative visual—not a fabricated product screenshot.

## Quality bar

The site should feel intentionally art-directed. Avoid a template-like stack of identical rounded cards. Vary section composition while maintaining a coherent system. The content should remain the visual focus.
