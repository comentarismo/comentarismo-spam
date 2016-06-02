package spamc_test

import (
	"testing"
	"comentarismo-spam/spamc"
	"github.com/stretchr/testify/assert"
)

func TestPortuguese_tokenizer(t *testing.T) {
	test_string := "escola exemplo estudante $(.;)*#()#*@)&(*&(*^@#*&)!fajs`ldkfj 23"
	expected_res := []string{"escola","estudante", "fajs", "ldkfj"}

	words := spamc.Tokenizer(test_string,spamc.Portuguese_ignore_words_map)
	assert.True(t, len(words)>0, "expected to find some words?! ")
	for i, word := range expected_res {
		if words[i] != word {
			t.Errorf("tokenizer failed, expected: %s", expected_res)
			t.Errorf("tokenizer failed, actually: %s, len:%d", words, len(words))
		}
	}
}

func TestEnglish_tokenizer(t *testing.T) {
	test_string := "love again fjalsdfj $(.;)*#()#*@)&(*&(*^@#*&)!fajs`ldkfj 23"
	expected_res := []string{"love","fjalsdfj", "fajs", "ldkfj"}

	words := spamc.Tokenizer(test_string,spamc.English_ignore_words_map)
	assert.True(t, len(words)>0, "expected to find some words?! ")
	for i, word := range expected_res {
		if words[i] != word {
			t.Errorf("tokenizer failed, expected: %s", expected_res)
			t.Errorf("tokenizer failed, actually: %s, len:%d", words, len(words))
		}
	}
}


func TestSpanish_tokenizer(t *testing.T) {
	test_string := "amor parecer fjalsdfj $(.;)*#()#*@)&(*&(*^@#*&)!fajs`ldkfj 23"
	expected_res := []string{"amor","fjalsdfj", "fajs", "ldkfj"}

	words := spamc.Tokenizer(test_string,spamc.Spanish_ignore_words_map)
	assert.True(t, len(words)>0, "expected to find some words?! ")
	for i, word := range expected_res {
		if words[i] != word {
			t.Errorf("tokenizer failed, expected: %s", expected_res)
			t.Errorf("tokenizer failed, actually: %s, len:%d", words, len(words))
		}
	}
}

func TestItalian_tokenizer(t *testing.T) {
	test_string := "pizza fianco fjalsdfj $(.;)*#()#*@)&(*&(*^@#*&)!fajs`ldkfj 23"
	expected_res := []string{"pizza","fjalsdfj", "fajs", "ldkfj"}

	words := spamc.Tokenizer(test_string,spamc.Italian_ignore_words_map)
	assert.True(t, len(words)>0, "expected to find some words?! ")
	for i, word := range expected_res {
		if words[i] != word {
			t.Errorf("tokenizer failed, expected: %s", expected_res)
			t.Errorf("tokenizer failed, actually: %s, len:%d", words, len(words))
		}
	}
}

func TestFrench_tokenizer(t *testing.T) {
	test_string := "creme plusieurs fjalsdfj $(.;)*#()#*@)&(*&(*^@#*&)!fajs`ldkfj 23"
	expected_res := []string{"creme","fjalsdfj", "fajs", "ldkfj"}

	words := spamc.Tokenizer(test_string,spamc.French_ignore_words_map)
	assert.True(t, len(words)>0, "expected to find some words?! ")
	for i, word := range expected_res {
		if words[i] != word {
			t.Errorf("tokenizer failed, expected: %s", expected_res)
			t.Errorf("tokenizer failed, actually: %s, len:%d", words, len(words))
		}
	}
}