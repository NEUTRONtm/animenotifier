package userroutes

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/pages/animelist"
	"github.com/animenotifier/notify.moe/pages/animelistitem"
	"github.com/animenotifier/notify.moe/pages/compare"
	"github.com/animenotifier/notify.moe/pages/explore/explorerelations"
	"github.com/animenotifier/notify.moe/pages/notifications"
	"github.com/animenotifier/notify.moe/pages/profile"
	"github.com/animenotifier/notify.moe/pages/profile/profilecharacters"
	"github.com/animenotifier/notify.moe/pages/recommended"
	"github.com/animenotifier/notify.moe/pages/user"
	"github.com/animenotifier/notify.moe/utils/page"
)

// Register registers the page routes.
func Register(app *aero.Application) {
	// User profiles
	page.Get(app, "/user", user.Get)
	page.Get(app, "/user/:nick", profile.Get)
	page.Get(app, "/user/:nick/characters/liked", profilecharacters.Liked)
	// page.Get(app, "/user/:nick/forum/threads", profile.GetThreadsByUser)
	// page.Get(app, "/user/:nick/forum/posts", profile.GetPostsByUser)
	// page.Get(app, "/user/:nick/soundtracks/added", profiletracks.Added)
	// page.Get(app, "/user/:nick/soundtracks/added/from/:index", profiletracks.Added)
	// page.Get(app, "/user/:nick/soundtracks/liked", profiletracks.Liked)
	// page.Get(app, "/user/:nick/soundtracks/liked/from/:index", profiletracks.Liked)
	// page.Get(app, "/user/:nick/quotes/added", profilequotes.Added)
	// page.Get(app, "/user/:nick/quotes/added/from/:index", profilequotes.Added)
	// page.Get(app, "/user/:nick/quotes/liked", profilequotes.Liked)
	// page.Get(app, "/user/:nick/quotes/liked/from/:index", profilequotes.Liked)
	// page.Get(app, "/user/:nick/stats", profile.GetStatsByUser)
	// page.Get(app, "/user/:nick/followers", profile.GetFollowers)
	page.Get(app, "/user/:nick/animelist/anime/:id", animelistitem.Get)
	page.Get(app, "/user/:nick/anime/recommended", recommended.Anime)
	page.Get(app, "/user/:nick/anime/sequels", explorerelations.Sequels)
	page.Get(app, "/user/:nick/notifications", notifications.ByUser)
	page.Get(app, "/user/:nick/edit", user.Edit)

	// Anime list
	page.Get(app, "/user/:nick/animelist/watching", animelist.FilterByStatus(arn.AnimeListStatusWatching))
	page.Get(app, "/user/:nick/animelist/completed", animelist.FilterByStatus(arn.AnimeListStatusCompleted))
	page.Get(app, "/user/:nick/animelist/planned", animelist.FilterByStatus(arn.AnimeListStatusPlanned))
	page.Get(app, "/user/:nick/animelist/hold", animelist.FilterByStatus(arn.AnimeListStatusHold))
	page.Get(app, "/user/:nick/animelist/dropped", animelist.FilterByStatus(arn.AnimeListStatusDropped))

	page.Get(app, "/user/:nick/animelist/watching/from/:index", animelist.FilterByStatus(arn.AnimeListStatusWatching))
	page.Get(app, "/user/:nick/animelist/completed/from/:index", animelist.FilterByStatus(arn.AnimeListStatusCompleted))
	page.Get(app, "/user/:nick/animelist/planned/from/:index", animelist.FilterByStatus(arn.AnimeListStatusPlanned))
	page.Get(app, "/user/:nick/animelist/hold/from/:index", animelist.FilterByStatus(arn.AnimeListStatusHold))
	page.Get(app, "/user/:nick/animelist/dropped/from/:index", animelist.FilterByStatus(arn.AnimeListStatusDropped))

	// Redirects
	page.Get(app, "/animelist/watching", animelist.Redirect)
	page.Get(app, "/animelist/completed", animelist.Redirect)
	page.Get(app, "/animelist/planned", animelist.Redirect)
	page.Get(app, "/animelist/hold", animelist.Redirect)
	page.Get(app, "/animelist/dropped", animelist.Redirect)

	// Compare
	page.Get(app, "/compare/animelist/:nick-1/:nick-2", compare.AnimeList)

	// Notifications
	page.Get(app, "/notifications", notifications.ByUser)
	page.Get(app, "/notifications/all", notifications.All)
}
