package spamc_test

import (
	"testing"
	"comentarismo-spam/spamc"
)

func TestClassifySpamEn(t *testing.T) {
	spamc.Train("good", "sunshine drugs love sex lobster sloth")
	spamc.Train("bad", "fear death horror government zombie god")

	targetWord := "sloths are so cute i love them"
	class := spamc.Classify(targetWord)
	if class != "good" {
		t.Errorf("Classify failed, word (%s) should be good, result: %s", targetWord, class)
	}

	targetWord = "i fear god and love the government"
	class = spamc.Classify(targetWord)
	if class != "bad" {
		t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
	}

	targetWord = "Fantastic deal"
	class = spamc.Classify(targetWord)
	if class != "bad" {
		t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
	}else {
		t.Logf("SPAM DETECTED, word (%s), result: %s", targetWord, class)
	}

	targetWord = "Get paid Now"
	class = spamc.Classify(targetWord)
	if class != "bad" {
		t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
	}else {
		t.Logf("SPAM DETECTED, word (%s), result: %s", targetWord, class)
	}

	targetWord = "Cancel at any time with Full refund"
	class = spamc.Classify(targetWord)
	if class != "bad" {
		t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
	}else {
		t.Logf("SPAM DETECTED, word (%s), result: %s", targetWord, class)
	}

	targetWord = "Easy terms its Full refund"
	class = spamc.Classify(targetWord)
	if class != "bad" {
		t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
	}else {
		t.Logf("SPAM DETECTED, word (%s), result: %s", targetWord, class)
	}

}

func init() {
	spamc.Flush()
}
