package helper

import "testing"

func TestEqualsIgnoreSpaces(t *testing.T) {

	t.Run("Should validate two equals strings without space", func(t *testing.T) {
		if false == EqualsIgnoreSpaces("bubble", "bubble") {
			t.Errorf("comparation is wrong, 'bubble' and 'bubble' are equals")
		}
	})

	t.Run("Should validate two different strings without space", func(t *testing.T) {
		if true == EqualsIgnoreSpaces("rock", "water") {
			t.Errorf("comparation is wrong, 'rock' and 'water' are differents")
		}
	})

	t.Run("Should validate two equals strings with same number of spaces", func(t *testing.T) {
		if false == EqualsIgnoreSpaces("  bubble", "  bubble") {
			t.Errorf("comparation is wrong, '  bubble' and '  bubble' are equals")
		}
	})

	t.Run("Should validate two equals strings with different number of spaces", func(t *testing.T) {
		if false == EqualsIgnoreSpaces("     bubble", "            bubble") {
			t.Errorf("comparation is wrong, '  bubble' and '  bubble' are equals")
		}
	})

	t.Run("Should validate two equals strings but only the first one contains left spaces", func(t *testing.T) {
		if false == EqualsIgnoreSpaces("  bubble", "bubble") {
			t.Errorf("comparation is wrong, '  bubble' and 'bubble' are equals")
		}
	})

	t.Run("Should validate two equals strings but only the second one contains left spaces", func(t *testing.T) {
		if false == EqualsIgnoreSpaces("bubble", "  bubble") {
			t.Errorf("comparation is wrong, '  bubble' and 'bubble' are equals")
		}
	})

	t.Run("Should validate two equals strings but only the first one contains right spaces", func(t *testing.T) {
		if false == EqualsIgnoreSpaces("bubble  ", "bubble") {
			t.Errorf("comparation is wrong, 'bubble  ' and 'bubble' are equals")
		}
	})

	t.Run("Should validate two equals strings but only the second one contains right spaces", func(t *testing.T) {
		if false == EqualsIgnoreSpaces("bubble", "bubble  ") {
			t.Errorf("comparation is wrong, 'bubble' and 'bubble  ' are equals")
		}
	})

	t.Run("Should validate two equals strings with character [NBSP]", func(t *testing.T) {
		if false == EqualsIgnoreSpaces(" bubble", " bubble") {
			t.Errorf("comparation is wrong, ' bubble' and ' bubble' are equals")
		}
	})

	t.Run("Should validate two equals strings but only the first one contains character [NBSP] on left", func(t *testing.T) {
		if false == EqualsIgnoreSpaces(" bubble", "bubble") {
			t.Errorf("comparation is wrong, ' bubble' and 'bubble' are equals")
		}
	})

	t.Run("Should validate two equals strings but only the second one contains character [NBSP] on left", func(t *testing.T) {
		if false == EqualsIgnoreSpaces("bubble", " bubble") {
			t.Errorf("comparation is wrong, 'bubble' and ' bubble' are equals")
		}
	})

	t.Run("Should validate two equals strings but only the first one contains character [NBSP] on right", func(t *testing.T) {
		if false == EqualsIgnoreSpaces("bubble ", "bubble") {
			t.Errorf("comparation is wrong, 'bubble ' and 'bubble' are equals")
		}
	})

	t.Run("Should validate two equals strings but only the second one contains character [NBSP] on right", func(t *testing.T) {
		if false == EqualsIgnoreSpaces("bubble", "bubble ") {
			t.Errorf("comparation is wrong, 'bubble' and 'bubble ' are equals")
		}
	})

	t.Run("Should validate two equals strings with different number of characters [NBSP]", func(t *testing.T) {
		if false == EqualsIgnoreSpaces("   bubble ", "     bubble") {
			t.Errorf("comparation is wrong, '   bubble ' and '     bubble' are equals")
		}
	})

	t.Run("Should validate two equals strings but the first one contains normal blank spaces and the second one contains character [NBSP]", func(t *testing.T) {
		if false == EqualsIgnoreSpaces("bubble ", "bubble ") {
			t.Errorf("comparation is wrong, 'bubble ' and 'bubble ' are equals")
		}
	})
}
