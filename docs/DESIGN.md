# Current direction: restored neon dashboard

The user explicitly preferred the earlier compact neon dashboard over the subsequent Swiss/editorial revisions. The latest build restores its uppercase hero, three-column composition, read-only profile window with traffic-light chrome, bordered summary strip, skills cards, and project thumbnails. The interactive 3D globe and accessibility/responsive fixes are retained. The full-lime band, enlarged editorial project rows, jump menu, and navigation prompt are removed. Accent color is a stronger lime `#b6ff18` against dark charcoal.

The sections below document earlier design iterations and asset sources; this current-direction note takes precedence.

# Neon dashboard design

The visual direction now follows the user-supplied `Neon Cybersecurity Portfolio Dashboard.png` reference. This supersedes the original kit's spacious editorial composition.

- Near-black `#050808`, panel `#0b0f10`, off-white `#e6e9eb`, muted `#a1a8ac`, lime `#b1ff24`, borders `#343a3d`.
- Locally hosted Space Grotesk for strong display typography; JetBrains Mono for navigation, content and terminal-style metadata.
- Desktop: compact top navigation; three-column introduction, globe and profile panel; framed summary strip; three-column about, skills and featured-project dashboard.
- Mobile: accessible collapsed navigation and a single-column flow. Keyboard-accessible skill tabs, real project filtering and project deep links.
- Static dotted globe generated from Natural Earth land polygons, centered on Asia with a Chennai marker. No WebGL, tracking, or runtime data fetches.
- Profile panel is explicitly static/read-only; no running command simulation or invented live telemetry.
- Project illustrations are decorative CSS/SVG abstractions, not application screenshots. No resume button is shown because no resume was supplied.

## Asset sources

- Natural Earth 1:110m land polygons: public domain. Source: https://github.com/nvkelso/natural-earth-vector/blob/master/geojson/ne_110m_land.geojson . Local source in `personal-portfolio/scripts/data/land.geojson`; regenerate via `python3 personal-portfolio/scripts/globe.py`.
- Space Grotesk and JetBrains Mono: Google Fonts distribution under SIL Open Font License. Font files and their licenses are bundled in `personal-portfolio/public/fonts/`.
- Interface SVG icons and project visuals are defined locally in the build source.

Content remains grounded in the supplied kit. The reference's CGPA, placement count, event registration metric, seniority, and availability were not treated as verified personal facts.

## Terminal / Swiss refinement

The latest direction uses a strict grid and an 8-pixel spacing scale. Desktop viewport gutters are 32px, tablet gutters 24px, and mobile gutters 16px. Window bodies use 24px padding (16px mobile); title bars use 16px horizontal padding. Square borders, flat charcoal surfaces, sans-serif editorial headings, and monospace controls replace gradients and decorative chrome.

`src/globe.js` is a self-contained custom element rendering a perspective-projected 3D geographic point cloud on canvas. Its source points come from the existing Natural Earth dataset. Dragging, arrow keys, zoom controls, reset and user-triggered rotation are supported. It is static by default, stops animation offscreen/in background tabs, caps pixel density at 2, and keeps the SVG fallback if canvas or data loading fails. Reduced-motion changes stop rotation; visitors can explicitly start it themselves. No WebGL or third-party runtime dependency is required.

The profile window now has a bounded navigation prompt. Supported commands navigate only to local sections (`help`, `home`, `about`, `skills`, `projects`, `ctf`, `education`, `contact`, `clear`). Nothing is sent to a server or shell. Input/output are limited, and unknown commands receive useful feedback.

## Live reference study and bolder editorial direction

Inspected https://www.cybercomctf.com/ and https://subadevan.tech/ in Chromium. Cybercom informed the numbered navigation blocks, sharp section boundaries, and larger content stages. Subadevan informed the oversized personal name, asymmetric introduction, fine grid, and keyboard-accessible jump menu. No reference-site text, personal claims, photos, background imagery, or source code was copied.

The palette is this portfolio's own: neutral black `#050505`, white `#f8f9fb`, electric lime `#c1ff00`, and a restrained blue `#4966ff` in illustrative project covers. The lime is used as a full-width profile-summary band and active controls, rather than only tiny accents. The hero is now two columns with an open globe view; the navigation terminal has its own section. Projects use wide editorial rows, larger typography and illustrative graphics. `Ctrl/Cmd+K` opens a native dialog for navigation, with Escape, keyboard focus management and a visible button.

Latest adjustment: the globe rotates automatically on load. Pause and drag controls remain available; reduced-motion preferences suppress automatic rotation.

Globe refinement: three inclined green orbital paths with depth-aware front/back rendering and small markers surround the rotating globe. Denser land dots and a dim 7,000-point global surface preserve the ocean-facing silhouette. Asia and Pacific views were inspected in Chromium; rotation and mobile overflow checks passed.

Latest layout correction: removed the 1440px content cap. All main section edges and header/footer now share fluid width with 32px desktop, 24px tablet and 16px mobile gutters. About has its own background/narrative section and distinct navigation position; Skills is a separate technical grid. This supersedes earlier centered-container notes.


Current page spacing: centered 1440px maximum outer container shared by the header, main content, and footer. Desktop inner gutters scale from 32px to 64px; tablet gutters are 24px and phone gutters 16px. This replaces the previous uncapped full-window layout at the user’s request.
