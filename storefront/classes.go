package storefront

import (
	"strings"

	"github.com/fastygo/ui8kit/utils"
)

func pageClass(props StorefrontPageProps) string {
	return utils.Cn("w-full bg-background text-foreground", props.Classes.Page)
}

func mainGridClass(props StorefrontPageProps) string {
	return utils.Cn("grid grid-cols-1 gap-6 xl:grid-cols-4", props.Classes.Main)
}

func contentClass(props StorefrontPageProps) string {
	return utils.Cn("min-w-0 xl:col-span-3", props.Classes.Content)
}

func railClass(props StorefrontPageProps) string {
	return utils.Cn("grid gap-6 xl:sticky xl:top-6 xl:self-start", props.Classes.Rail)
}

func softCardClass(extra ...string) string {
	return utils.Cn(append([]string{"overflow-hidden rounded-lg border border-border bg-card shadow-sm"}, extra...)...)
}

func mutedPanelClass(extra ...string) string {
	return utils.Cn(append([]string{"rounded-lg border border-border bg-muted/40"}, extra...)...)
}

func sectionHeaderClass() string {
	return "flex items-center justify-between gap-4"
}

func mediaClass(extra ...string) string {
	return utils.Cn(append([]string{"h-full w-full rounded-lg object-cover"}, extra...)...)
}

func iconCircleClass() string {
	return "inline-flex h-10 w-10 items-center justify-center rounded-full border border-border bg-background"
}

func labelOrFallback(value string, fallback string) string {
	if strings.TrimSpace(value) != "" {
		return value
	}
	return fallback
}

func hasAction(action Action) bool {
	return strings.TrimSpace(action.Label) != "" && strings.TrimSpace(action.Href) != ""
}

func actionLabel(action Action) string {
	if strings.TrimSpace(action.AriaLabel) != "" {
		return action.AriaLabel
	}
	return action.Label
}
