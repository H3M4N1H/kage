# Hemanth Portfolio — Codex handoff kit

This kit contains the project brief and reusable instructions for building the portfolio.

## Files

- `SKILL.md` — reusable agent skill/instructions for this portfolio.
- `MASTER_PROMPT.md` — one-shot implementation prompt to paste into Codex.
- `PORTFOLIO_CONTENT.md` — factual profile, projects, achievements, roles, and public links.
- `DESIGN_SPEC.md` — design system and visual direction.

## How to use

1. Open the existing portfolio repository in Codex.
2. Add these four Markdown files to the repository root (or place them in a project knowledge folder that Codex can read).
3. If your Codex environment supports project skills, register `SKILL.md` as a project skill. Otherwise, keep it in the repository and tell Codex to follow it.
4. Paste the full contents of `MASTER_PROMPT.md` into the Codex task.
5. Let the agent inspect, implement, and verify the project.
6. Review the factual confirmation items before deploying.

## Important

The content file includes facts from prior conversations. Confirm current role titles, project status, event naming, and public URLs before publishing. The prompt explicitly instructs the agent not to invent missing information.

The agent should not commit, push, deploy, purchase a domain, or change external services unless separately instructed.
