import AnimeNotifier from "../AnimeNotifier"
import { findAll } from "scripts/Utils"

// Filter anime on explore page
export function filterAnime(arn: AnimeNotifier, _: HTMLInputElement) {
	let root = document.getElementById("filter-root") as HTMLElement

	let elementYear = document.getElementById("filter-year") as HTMLSelectElement
	let elementSeason = document.getElementById("filter-season") as HTMLSelectElement
	let elementStatus = document.getElementById("filter-status") as HTMLSelectElement
	let elementType = document.getElementById("filter-type") as HTMLSelectElement

	for(let element of findAll("anime-grid-image")) {
		let img = element as HTMLImageElement
		img.src = arn.emptyPixel()
		img.classList.remove("element-found")
		img.classList.remove("element-color-preview")
	}

	let year = elementYear.value || "any"
	let season = elementSeason.value || "any"
	let status = elementStatus.value || "any"
	let type = elementType.value || "any"

	arn.diff(`${root.dataset.url}/${year}/${season}/${status}/${type}`)
}

// Toggle hiding added anime.
export function toggleHideAddedAnime() {
	hideAddedAnime()
}

// Hides anime that are already in your list.
export function hideAddedAnime() {
	for(let anime of findAll("anime-grid-cell")) {
		if(anime.dataset.added !== "true") {
			continue
		}

		anime.classList.toggle("anime-grid-cell-hide")
	}
}

// Hides anime that are not in your list.
export async function calendarShowAddedAnimeOnly(arn: AnimeNotifier, element: HTMLInputElement) {
	let calendar = document.getElementById("calendar")

	if(!calendar || calendar.dataset.showAddedAnimeOnly === undefined) {
		return
	}

	// Toggling the switch will trigger the CSS rules
	if(calendar.dataset.showAddedAnimeOnly === "true") {
		calendar.dataset.showAddedAnimeOnly = "false"
	} else {
		calendar.dataset.showAddedAnimeOnly = "true"
	}

	// Save the state in the database
	let showAddedAnimeOnly = calendar.dataset.showAddedAnimeOnly === "true"
	let apiEndpoint = arn.findAPIEndpoint(element)

	try {
		await arn.post(apiEndpoint, {
			"Calendar.ShowAddedAnimeOnly": showAddedAnimeOnly
		})
	} catch(err) {
		arn.statusMessage.showError(err)
	}
}