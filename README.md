# FastyGo Blocks

**Official repository:** [github.com/fastygo/blocks](https://github.com/fastygo/blocks)

## Why Blocks exists

[**UI8Kit**](https://github.com/fastygo/ui8kit) gives **atoms** and neutral layout. Blocks provide copy-first page organisms without keeping a runtime dependency on a sibling workspace.

**Blocks** is where **organisms** live: coherent sections of a page with a **stable, named API** (`Hero`, `Pricing`, `Shell region`, …), composed directly from **UI8Kit** primitives and design tokens.

So: **Blocks exist so product pages can copy whole sections** without each app re-implementing the same wiring, markup, and primitive imports.

---

## Layering (where Blocks sits)

| Layer | Responsibility |
|-------|----------------|
| **UI8Kit** | Atoms + layout grammar; **Go** + minimal `ui8kit` JS. |
| **Elements** | Copy-first widgets + extra APG JS; **no** page scenes. |
| **Blocks** (this repo) | **Page organisms** — compose UI8Kit primitives; scene-specific props and layout. |
| **App** | Brand, routes, data, campaigns, one-off views. |

---

## Standards (defaults)

- **Scenes and composition.** Blocks wire **sections** of the UI: headers, heroes, sidebars, multi-column regions. Brand-only gradients, campaign art, and product-specific chrome do not belong here.
- **Copy-first dependencies:** Blocks may import **UI8Kit** only. Do not import Elements, other Blocks, apps, or workspace-local packages. See [`.cursor/rules/blocks-layer.mdc`](.cursor/rules/blocks-layer.mdc).
- **Direct UI8Kit use is allowed** for primitives in block templates (no rule forcing every button through Elements).
- **Growth path:** duplicate first; extract only when the copy contract remains UI8Kit-only and portable.
- **Go / templ:** use only explicit Tailwind utility classes that pass `.ui8px` policy. Do not ship named CSS classes, custom CSS, brand classes, or part-class override APIs from Blocks.
- **Quality:** **Integration** checks (composition, focus between regions, block-level fixtures) fit here or in apps; isolated widget a11y stays primarily in **Elements**. Cross-stack guidance: [Elements `.project/VALIDATION-AND-TOOLING.md`](https://github.com/fastygo/elements/blob/main/.project/VALIDATION-AND-TOOLING.md) (same ideas apply when you add a `Blocks/.project` later).

---

## UI8Kit in one sentence

**Go:** primitives you assemble in block templates.  
**JS:** small behaviors already bundled with the kit — blocks rely on it indirectly when using kit markup conventions.

## Elements in one sentence

**Go + JS:** drop-in widgets (menus, trees, … over time) so blocks stay **declarative** and do not re-implement APG keyboard models.

## Blocks in one sentence

**Go/templ:** named **page sections** with a clear props model — the reusable “organism” layer between design systems and the shipping application.

## Packages

- `marketing`: reusable marketing sections such as landing heroes, feature grids, and simple page footers.
- `storefront`: copy-first commerce prototype sections.

---

## FastyGo stack

| Repository | Role |
|------------|------|
| [**framework**](https://github.com/fastygo/framework) | Core framework and examples |
| [**ui8kit**](https://github.com/fastygo/ui8kit) | Go/templ primitives + minimal `ui8kit` JS |
| [**elements**](https://github.com/fastygo/elements) | Reusable widgets + APG-oriented JS |
| [**blocks**](https://github.com/fastygo/blocks) | Page-level organisms (this repo) |

---

## Local development

Use the cross-repo workspace when changing UI8Kit, Elements, Blocks, and Framework examples together:

```bash
cd e:/_@Go/.WorkSpace-Framework
go work sync
```

This module uses a pseudo-zero UI8Kit requirement with a local replace:

```text
replace github.com/fastygo/ui8kit => ../@UI8Kit
```

Run checks from this directory. `ui8px` is intentionally not installed as a dependency; call it through `npx`:

```bash
go test ./...
npx ui8px@latest lint ./...
npx ui8px@latest lint ./... --learn
npx ui8px@latest validate patterns ./...
```

Replace pseudo-zero versions with tagged releases only when publishing stable module versions.

---

## License

This project is licensed under the **MIT License** — see [`LICENSE`](LICENSE).
