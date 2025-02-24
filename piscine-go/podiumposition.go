package piscine

func PodiumPosition(podium [][]string) [][]string {
	for k, j := 0, len(podium)-1; k < j; k, j = k+1, j-1 {
		podium[k], podium[j] = podium[j], podium[k]
	}
	return podium
}
