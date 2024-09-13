package partials

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/maddalax/mhtml/framework/h"
	"github.com/maddalax/mhtml/starter-template/news"
)

func NewsSheet(ctx *fiber.Ctx) *h.Partial {
	open := h.GetQueryParam(ctx, "open") == "true"
	if open {
		h.SessionIncr(ctx, "sheet-open-count")
	}
	return h.NewPartialWithHeaders(
		&map[string]string{
			"hx-trigger":  "sheetOpened",
			"hx-push-url": fmt.Sprintf("/news%s", h.Ternary(open, "?open=true", "")),
		},
		SheetWrapper(
			h.IfElseLazy(open, SheetOpen, SheetClosed),
			h.Swap(ctx, OpenSheetButton(open)),
			h.Swap(ctx, NewsSheetOpenCount(ctx).Root),
		),
	)
}

func NewsSheetOpenCount(ctx *fiber.Ctx) *h.Partial {
	rnd := h.SessionGet[int64](ctx, "sheet-open-count")
	if rnd == nil {
		rnd = new(int64)
	}

	open := h.GetQueryParam(ctx, "open") == "true"

	return h.NewPartial(h.Div(
		h.Id("sheet-open-count"),
		h.IfElse(open,
			h.Text(fmt.Sprintf("you opened sheet %d times", *rnd)),
			h.Text("sheet is not open")),
	),
	)
}

func SheetWrapper(children ...h.Renderable) h.Renderable {
	return h.Div(h.Id("sheet-partial"), h.Fragment(children...))
}

func SheetClosed() h.Renderable {
	return h.Div()
}

func SheetOpen() h.Renderable {
	return h.Fragment(h.Div(
		h.Class(`fixed top-0 right-0 h-full w-96 bg-gray-100 shadow-lg z-50`),
		h.Div(
			h.Class("p-4 overflow-y-auto h-full w-full flex flex-col gap-4"),
			h.P(h.Text("News Sheet"),
				h.Class("text-lg font-bold"),
			),
			h.P(h.Text("Here are the latest news stories."),
				h.Class("text-sm mt-2"),
			),
			ui.Button(ui.ButtonProps{
				Text:   "Close NewsSheet",
				Target: "#sheet-partial",
				Get:    h.GetPartialPathWithQs(NewsSheet, "open=false"),
			}),
			news.StoryList(),
		)))
}