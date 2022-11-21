package main


func TestPalindrome(t *testing.T) {
	if !lsPalindrome("detartrated") {
	t .Error(`IsPalindrome("detartrated") = false`)
	>
	if !lsPalindrome("kayak") {
	t.Error(`IsPalindrome("kayak") = false`)
	}
	}
	func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
	t.Error(`IsPalindrome("palindrome") = true`)
	}
	}