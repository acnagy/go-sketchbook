package main

func checkCI(ci string) (output bool) {
	if ci == "Travis CI" {
		return true
	}

	return
}
