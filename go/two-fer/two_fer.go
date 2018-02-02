package twofer

func ShareWith(s string) string {
	person := s
	if person == "" {
		person = "you"
	}
	sentence := "One for "
	sentence += person
	sentence += ", one for me."
	return sentence
}
