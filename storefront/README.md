# Storefront Block

Copy this package as source when an application needs the storefront prototype.

Copy:

- `props.go`
- `storefront.templ`
- `controls.templ`
- `helpers.go`
- `placeholders.go`

Do not copy generated files:

- `*_templ.go`

After copying:

1. Change the package name only if the destination package uses a different name.
2. Run `templ generate ./...`.
3. Run `.ui8px` checks for the destination app.

The package is intentionally self-contained and depends only on UI8Kit.
