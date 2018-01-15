package main

func checkCI(ci string) (output bool) {

	if ci == "Travis CI" {

		return true

	} else {

		return false
	}
}

func doSomethingUntested(msg string) {
	println(msg)
}
