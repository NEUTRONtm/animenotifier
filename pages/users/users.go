package users

import (
	"sort"
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Active ...
func Active(ctx aero.Context) error {
	users := arn.FilterUsers(func(user *arn.User) bool {
		return user.HasAvatar() && user.HasNick() && user.IsActive()
	})

	arn.SortUsersFollowers(users)
	return ctx.HTML(components.Users(users, ctx.Path()))
}

// Pro ...
func Pro(ctx aero.Context) error {
	users := arn.FilterUsers(func(user *arn.User) bool {
		return user.IsPro()
	})

	sort.Slice(users, func(i, j int) bool {
		return users[i].Registered > users[j].Registered
	})

	return ctx.HTML(components.ProUsers(users, ctx.Path()))
}

// Editors ...
func Editors(ctx aero.Context) error {
	score := map[string]int{}
	users := []*arn.User{}

	for entry := range arn.StreamEditLogEntries() {
		entryScore := entry.EditorScore()

		if entryScore == 0 {
			continue
		}

		current, exists := score[entry.UserID]

		if !exists {
			users = append(users, entry.User())
		}

		score[entry.UserID] = current + entryScore
	}

	for ignore := range arn.StreamIgnoreAnimeDifferences() {
		score[ignore.CreatedBy] += arn.IgnoreAnimeDifferenceEditorScore
	}

	sort.Slice(users, func(i, j int) bool {
		scoreA := score[users[i].ID]
		scoreB := score[users[j].ID]

		if scoreA == scoreB {
			return users[i].Registered > users[j].Registered
		}

		return scoreA > scoreB
	})

	if len(users) > 10 {
		users = users[:10]
	}

	return ctx.HTML(components.EditorRankingList(users, score, ctx.Path()))
}

// ActiveNoAvatar ...
func ActiveNoAvatar(ctx aero.Context) error {
	users := arn.FilterUsers(func(user *arn.User) bool {
		return user.IsActive() && !user.HasAvatar()
	})

	followCount := arn.UserFollowerCountMap()

	sort.Slice(users, func(i, j int) bool {
		if users[i].HasAvatar() != users[j].HasAvatar() {
			return users[i].HasAvatar()
		}

		followersA := followCount[users[i].ID]
		followersB := followCount[users[j].ID]

		if followersA == followersB {
			return users[i].Nick < users[j].Nick
		}

		return followersA > followersB
	})

	return ctx.HTML(components.Users(users, ctx.Path()))
}

// Osu ...
func Osu(ctx aero.Context) error {
	users := arn.FilterUsers(func(user *arn.User) bool {
		return user.HasAvatar() && user.HasNick() && user.IsActive() && user.Accounts.Osu.PP > 0
	})

	// Sort by pp
	sort.Slice(users, func(i, j int) bool {
		return users[i].Accounts.Osu.PP > users[j].Accounts.Osu.PP
	})

	if len(users) > 10 {
		users = users[:10]
	}

	return ctx.HTML(components.OsuRankingList(users, ctx.Path()))
}

// Overwatch ...
func Overwatch(ctx aero.Context) error {
	users := arn.FilterUsers(func(user *arn.User) bool {
		return user.HasAvatar() && user.HasNick() && user.IsActive() && user.Accounts.Overwatch.SkillRating > 0
	})

	// Sort by Skill Ratings
	sort.Slice(users, func(i, j int) bool {
		return users[i].Accounts.Overwatch.SkillRating > users[j].Accounts.Overwatch.SkillRating
	})

	if len(users) > 10 {
		users = users[:10]
	}

	return ctx.HTML(components.OverwatchRankingList(users, ctx.Path()))
}

// FFXIV ...
func FFXIV(ctx aero.Context) error {
	users := arn.FilterUsers(func(user *arn.User) bool {
		return user.HasAvatar() && user.HasNick() && user.IsActive() && user.Accounts.FinalFantasyXIV.ItemLevel > 0
	})

	// Sort by item level
	sort.Slice(users, func(i, j int) bool {
		return users[i].Accounts.FinalFantasyXIV.ItemLevel > users[j].Accounts.FinalFantasyXIV.ItemLevel
	})

	if len(users) > 10 {
		users = users[:10]
	}

	return ctx.HTML(components.FinalFantasyXIVRankingList(users, ctx.Path()))
}

// ByCountry ...
func ByCountry(ctx aero.Context) error {
	countryName := ctx.Get("country")

	users := arn.FilterUsers(func(user *arn.User) bool {
		return strings.ToLower(user.Location.CountryName) == countryName && user.Settings().Privacy.ShowLocation && user.HasAvatar() && user.HasNick() && user.IsActive()
	})

	arn.SortUsersFollowers(users)
	return ctx.HTML(components.UsersByCountry(users, countryName))
}

// Staff ...
func Staff(ctx aero.Context) error {
	users := arn.FilterUsers(func(user *arn.User) bool {
		return user.HasAvatar() && user.HasNick() && user.IsActive() && user.Role != ""
	})

	// Make order deterministic
	sort.Slice(users, func(i, j int) bool {
		return users[i].Nick < users[j].Nick
	})

	admins := &utils.UserList{
		Name:  "Developer",
		Users: []*arn.User{},
	}

	contributors := &utils.UserList{
		Name:  "Contributors",
		Users: []*arn.User{},
	}

	// contributors.Users = append(contributors.Users, )

	editors := &utils.UserList{
		Name:  "Editors",
		Users: []*arn.User{},
	}

	contributorIDs := []string{
		"VJOK1ckvx", // Scott
		"SUQOAFFkR", // Allen
		"KQgtMWOiR", // Franksks
		"9NMYrAHiR", // Amatrelan
	}

	for _, user := range users {
		if user.Role == "admin" {
			admins.Users = append(admins.Users, user)
			continue
		}

		if arn.Contains(contributorIDs, user.ID) {
			contributors.Users = append(contributors.Users, user)
			continue
		}

		if user.Role == "editor" {
			editors.Users = append(editors.Users, user)
			continue
		}
	}

	userLists := []*utils.UserList{
		admins,
		contributors,
		editors,
	}

	return ctx.HTML(components.UserLists(userLists, ctx.Path()) + components.StaffRecruitment())
}
