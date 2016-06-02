package spamc_test

import (
	"testing"
	"comentarismo-spam/spamc"
)

func TestTidy(t *testing.T) {
	test_string := "fjalsdfj $()*#()#*@)&(*&(*^@#*&)!fajs`ldkfj 23"

	if s_out := spamc.Tidy(test_string); s_out != "fjalsdfj fajs ldkfj 23" {
		t.Errorf("Tidy failed:\n expected: fjalsdfj fajsldkfj 23\n result:%s\n", s_out)
	}
}

func TestOccurances(t *testing.T) {
	words := []string{"fjalsdfj", "23", "fjalsdfj", "23", "ok"}
	res := spamc.Occurances(words)
	expected_res := map[string]uint{
		"23":       2,
		"fjalsdfj": 2,
		"ok":       1,
	}

	for k, v := range expected_res {
		if res[k] != v {
			t.Errorf("Occurances failed: %s", expected_res)
		}
	}
}

func TestFlush(t *testing.T) {
	spamc.Train("good", "sunshine drugs love sex lobster sloth")
	spamc.Flush()

	exists := spamc.RedisClient.Exists(spamc.Redis_prefix + "good")
	if exists.Val() {
		t.Errorf("Flush failed")
	}
}

func TestClassify(t *testing.T) {
	spamc.Train("good", "sunshine drugs love sex lobster sloth")
	spamc.Train("bad", "fear death horror government zombie god")

	class := spamc.Classify("sloths are so cute i love them")
	if class != "good" {
		t.Errorf("Classify failed, should be good, result: %s", class)
	}

	class = spamc.Classify("i fear god and love the government")
	if class != "bad" {
		t.Errorf("Classify failed, should be bad, result: %s", class)
	}
}

func TestUntrain(t *testing.T) {
	spamc.Flush()
	spamc.Train("good", "sunshine drugs love sex lobster sloth")
	spamc.Untrain("good", "sunshine drugs love sex lobster sloth")

	exists := spamc.RedisClient.Exists(spamc.Redis_prefix + "good")
	if exists.Val() {
		t.Errorf("TestUntrain failed")
	}
}

func init() {
	spamc.Flush()
}
