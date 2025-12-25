package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var n int = len(birdsPerDay)
    var totalBirds int = 0
    for i := 0; i < n; i++ {
        totalBirds += birdsPerDay[i]
    }
    return totalBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var startingDate int = 7 * (week - 1)
    var endingDate int = startingDate + 7
    var birdCount int = 0
    for i := startingDate ; i < endingDate; i++ {
        birdCount += birdsPerDay[i]
    }
    return birdCount
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	var n int = len(birdsPerDay)
    for i := 0; i < n; i += 2 {
        birdsPerDay[i] += 1
    }
    return birdsPerDay
}
