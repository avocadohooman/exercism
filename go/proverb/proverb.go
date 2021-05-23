package proverb

func Proverb(rhyme []string) []string {
	//First I check if the received paramater is empty. If yes, I return an empty string array
	if len(rhyme) == 0 {
		return []string{};
	}
	// Here I allocate memory for a string array depending on the length of the received string array "rhyme". This is necessary, so I can create 
	// a dynimcally allocated array which can adjust to whatever length the proverb should be
	proverb := make([]string, len(rhyme))
	for i := 0; i < len(rhyme); i++ {
		if (i < len(rhyme) && rhyme[i + 1]) {
			proverb[i] = "For want of a " + rhyme[i] + " the " + rhyme[i + 1] + " was lost."
		} else {
			proverb[i] = "And all for the want of a " + rhyme[0] + ".";
		}
	}
	return proverb
}
