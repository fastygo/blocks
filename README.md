# FastyGo Blocks

**Official repository:** [github.com/fastygo/blocks](https://github.com/fastygo/blocks)

## Why Blocks exists

[**UI8Kit**](https://github.com/fastygo/ui8kit) gives **atoms** and neutral layout. [**Elements**](https://github.com/fastygo/elements) gives **reusable widgets** and APG-oriented JS where the kit stops. Neither layer should carry **page-level scenes** (marketing hero, pricing grid, dashboard shell as a whole) — that would either bloat the kit or scatter the same composition across every application.

**Blocks** is where **organisms** live: coherent sections of a page with a **stable, named API** (`Hero`, `Pricing`, `Shell region`, …), composed from **UI8Kit** and **Elements**, optionally with **block-scoped** CSS (`bl-*`) for layout roles that are not generic tokens.

So: **Blocks exist so product pages reuse whole sections** without each app re-implementing the same wiring, markup, and primitive imports.

---

## Layering (where Blocks sits)

| Layer | Responsibility |
|-------|----------------|
| **UI8Kit** | Atoms + layout grammar; **Go** + minimal `ui8kit` JS. |
| **Elements** | Complex widgets + extra APG JS; **no** page scenes. |
| **Blocks** (this repo) | **Page organisms** — compose primitives + widgets; scene-specific props and layout. |
| **App** | Brand, routes, data, campaigns, one-off views. |

---

## Standards (defaults)

- **Scenes and composition.** Blocks wire **sections** of the UI: headers, heroes, sidebars, multi-column regions. Keep **brand-only** gradients, campaign art, and product-specific chrome in the **app** unless a pattern becomes a **reusable family** of blocks.
- **Dependencies:** Blocks may import **UI8Kit** and **Elements**. Lower layers **must not** import Blocks. See [`.cursor/rules/blocks-layer.mdc`](.cursor/rules/blocks-layer.mdc).
- **Direct UI8Kit use is allowed** for primitives in block templates (no rule forcing every button through Elements).
- **Growth path** (avoid premature Elements): (1) repeat inside one block → `internal/`; (2) shared across blocks in this repo → `pkg/` or a shared subpackage; (3) stable cross-product **widget** → **Elements**.
- **Go / templ:** Same discipline as upstream — variants and props from UI8Kit/Elements; avoid raw Tailwind utility strings in library templates; use `bl-*` or app CSS for scene-level styling.
- **Quality:** **Integration** checks (composition, focus between regions, block-level fixtures) fit here or in apps; isolated widget a11y stays primarily in **Elements**. Cross-stack guidance: [Elements `.project/VALIDATION-AND-TOOLING.md`](https://github.com/fastygo/elements/blob/main/.project/VALIDATION-AND-TOOLING.md) (same ideas apply when you add a `Blocks/.project` later).

---

## UI8Kit in one sentence

**Go:** primitives you assemble in block templates.  
**JS:** small behaviors already bundled with the kit — blocks rely on it indirectly when using kit markup conventions.

## Elements in one sentence

**Go + JS:** drop-in widgets (menus, trees, … over time) so blocks stay **declarative** and do not re-implement APG keyboard models.

## Blocks in one sentence

**Go/templ:** named **page sections** with a clear props model — the reusable “organism” layer between design systems and the shipping application.

---

## FastyGo stack

| Repository | Role |
|------------|------|
| [**framework**](https://github.com/fastygo/framework) | Core framework and examples |
| [**ui8kit**](https://github.com/fastygo/ui8kit) | Go/templ primitives + minimal `ui8kit` JS |
| [**elements**](https://github.com/fastygo/elements) | Reusable widgets + APG-oriented JS |
| [**blocks**](https://github.com/fastygo/blocks) | Page-level organisms (this repo) |

---

## License

This project is licensed under the **MIT License** — see [`LICENSE`](LICENSE).
