import { plural } from "./Utils"

const oneSecond = 1000
const oneMinute = 60 * oneSecond
const oneHour = 60 * oneMinute
const oneDay = 24 * oneHour
const oneWeek = 7 * oneDay
const oneMonth = 30 * oneDay
const oneYear = 365.25 * oneDay

export var monthNames = [
	"January", "February", "March",
	"April", "May", "June", "July",
	"August", "September", "October",
	"November", "December"
]

export var dayNames = [
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday"
]

function getRemainingTime(remaining: number): string {
	let remainingAbs = Math.abs(remaining)

	if(remainingAbs >= oneYear) {
		return plural(Math.round(remaining / oneYear), "year")
	}

	if(remainingAbs >= oneMonth) {
		return plural(Math.round(remaining / oneMonth), "month")
	}

	if(remainingAbs >= oneWeek) {
		return plural(Math.round(remaining / oneWeek), "week")
	}

	if(remainingAbs >= oneDay) {
		return plural(Math.round(remaining / oneDay), "day")
	}

	if(remainingAbs >= oneHour) {
		return plural(Math.round(remaining / oneHour), " hour")
	}

	if(remainingAbs >= oneMinute) {
		return plural(Math.round(remaining / oneMinute), " minute")
	}

	if(remainingAbs >= oneSecond) {
		return plural(Math.round(remaining / oneSecond), " second")
	}

	return "Just now"
}

export function displayAiringDate(element: HTMLElement, now: Date) {
	if(!element.dataset.startDate || !element.dataset.endDate) {
		element.textContent = ""
		return
	}

	let startDate = new Date(element.dataset.startDate)
	let endDate = new Date(element.dataset.endDate)

	let h = startDate.getHours()
	let m = startDate.getMinutes()
	let startTime = (h <= 9 ? "0" + h : h) + ":" + (m <= 9 ? "0" + m : m)

	h = endDate.getHours()
	m = endDate.getMinutes()
	let endTime = (h <= 9 ? "0" + h : h) + ":" + (m <= 9 ? "0" + m : m)

	let airingVerb = "will be airing"

	let remaining = startDate.getTime() - now.getTime()
	let remainingString = getRemainingTime(remaining)

	// Add "ago" if the date is in the past
	if(remainingString.startsWith("-")) {
		remainingString = remainingString.substring(1) + " ago"
	}

	element.textContent = remainingString

	if(remaining < 0) {
		airingVerb = "aired"
	}

	let tooltip = "Episode " + element.dataset.episodeNumber + " " + airingVerb + " " + dayNames[startDate.getDay()] + " from " + startTime + " - " + endTime

	if(element.classList.contains("no-tip")) {
		element.title = tooltip
	} else {
		element.setAttribute("aria-label", tooltip)
		element.classList.add("tip")
	}
}

export function displayDate(element: HTMLElement, now: Date) {
	if(!element.dataset.date) {
		element.textContent = ""
		return
	}

	let startDate = new Date(element.dataset.date)

	let h = startDate.getHours()
	let m = startDate.getMinutes()
	let startTime = (h <= 9 ? "0" + h : h) + ":" + (m <= 9 ? "0" + m : m)

	let remaining = startDate.getTime() - now.getTime()
	let remainingString = getRemainingTime(remaining)

	// Add "ago" if the date is in the past
	if(remainingString.startsWith("-")) {
		remainingString = remainingString.substring(1) + " ago"
	}

	element.textContent = remainingString
	let tooltip = dayNames[startDate.getDay()] + " " + startTime

	if(element.classList.contains("no-tip")) {
		element.title = tooltip
	} else {
		element.setAttribute("aria-label", tooltip)
		element.classList.add("tip")
	}
}

export function displayTime(element: HTMLElement) {
	if(!element.dataset.date) {
		element.textContent = ""
		return
	}

	let startDate = new Date(element.dataset.date)

	let h = startDate.getHours()
	let m = startDate.getMinutes()
	let startTime = (h <= 9 ? "0" + h : h) + ":" + (m <= 9 ? "0" + m : m)

	element.textContent = startTime
}