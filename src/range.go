package main

//using cuss words is inappropriate and violent. No more cuss word GUYSSS!

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, word := range msg{
		for _, badWord := range badWords{
			if word == badWord{
				return i
			}
		}
	}
	return -1
}