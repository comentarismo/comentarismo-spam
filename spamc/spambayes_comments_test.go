package spamc_test

import (
	"testing"
	"comentarismo-spam/spamc"
	"log"
)

func TestClassifySpamEn(t *testing.T) {
	log.Println("Will start server on learning mode, default to English. ")
	targetFile := spamc.GetPWD("/spamc/config_spamwords_en.yaml")

	spamc.StartLanguageSpam(targetFile, "english_spam","en")

	lang := "en"
	spamc.Train("good", "sunshine drugs love sex lobster sloth",lang)
	spamc.Train("bad", "fear death horror government zombie god",lang)

	targetWord := "sloths are so cute i love them"
	class := spamc.Classify(targetWord,lang)
	if class != "good" {
		t.Errorf("Classify failed, word (%s) should be good, result: %s", targetWord, class)
	}

	targetWord = "i fear god and love the government"
	class = spamc.Classify(targetWord,lang)
	if class != "bad" {
		t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
	}

	targetWord = "Fantastic deal"
	class = spamc.Classify(targetWord,lang)
	if class != "bad" {
		t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
	}else {
		t.Logf("SPAM DETECTED, word (%s), result: %s", targetWord, class)
	}

	targetWord = "Get paid Now"
	class = spamc.Classify(targetWord,lang)
	if class != "bad" {
		t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
	}else {
		t.Logf("SPAM DETECTED, word (%s), result: %s", targetWord, class)
	}

	targetWord = "Cancel at any time with Full refund"
	class = spamc.Classify(targetWord,lang)
	if class != "bad" {
		t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
	}else {
		t.Logf("SPAM DETECTED, word (%s), result: %s", targetWord, class)
	}

	targetWord = "Easy terms its Full refund"
	class = spamc.Classify(targetWord,lang)
	if class != "bad" {
		t.Errorf("Classify failed, word (%s) should be bad, result: %s", targetWord, class)
	}else {
		t.Logf("SPAM DETECTED, word (%s), result: %s", targetWord, class)
	}

}

func init() {
	spamc.Flush()
}
