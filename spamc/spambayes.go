package spamc

import (
	redis "gopkg.in/redis.v3"
	"github.com/kylelemons/go-gypsy/yaml"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
	"os"
	"comentarismo-spam/lang"
)

var (
	English_ignore_words_map = make(map[string]int)
	Portuguese_ignore_words_map = make(map[string]int)
	Spanish_ignore_words_map = make(map[string]int)
	Italian_ignore_words_map = make(map[string]int)
	French_ignore_words_map = make(map[string]int)

	RedisClient                        *redis.Client
	Redis_prefix = "bayes:"
	correction = 0.1
)

var REDIS_HOST = os.Getenv("REDIS_HOST")
var REDIS_PORT = os.Getenv("REDIS_PORT")
var REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
var SPAM_DEBUG = os.Getenv("SPAM_DEBUG");

var LEARNSPAM = os.Getenv("LEARNSPAM")

// replace \_.,<>:;~+|\[\]?`"!@#$%^&*()\s chars with whitespace
// re.sub(r'[\_.,<>:;~+|\[\]?`"!@#$%^&*()\s]', ' '
func Tidy(s string) (safe string) {
	reg, err := regexp.Compile("[\\_.,:;~+|\\[\\]?`\"!@#$%^&*()\\s]+")
	if err != nil {
		Debug("Error: Tidy, ", err)
		return
	}

	text_in_lower := strings.ToLower(s)
	safe = reg.ReplaceAllLiteralString(text_in_lower, " ")
	return
}

// tidy the input text, ignore those text composed with less than 2 chars
func Tokenizer(s string, ignore_words_map map[string]int) (res []string) {
	words := strings.Fields(Tidy(s))
	// this slice's length should be initialized to 0
	// otherwise, the first element will be the whitespace(empty string)
	res = make([]string, 0)

	for _, word := range words {
		strings.TrimSpace(word)
		_, omit := ignore_words_map[word]
		if omit || len(word) <= 2 {
			continue
		}
		res = append(res, word)
	}

	return
}


// compute word occurances
func Occurances(words []string) (counts map[string]uint) {
	counts = make(map[string]uint)
	for _, word := range words {
		if _, ok := counts[word]; ok {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}
	return
}

func Flush() {
	reply := RedisClient.SMembers(Redis_prefix + "categories")

	for _, key := range reply.Val() {
		RedisClient.Del(Redis_prefix + string(key))
	}

	RedisClient.Del(Redis_prefix + "categories")
}

func Train(categories, text string) {
	RedisClient.SAdd(Redis_prefix + "categories", categories)

	token_occur := Occurances(Tokenizer(text, English_ignore_words_map))
	for word, count := range token_occur {
		Debug("Train, ", word, count)
		RedisClient.HIncrBy(Redis_prefix + categories, word, int64(count))
	}
}

func Untrain(categories, text string) {
	token_occur := Occurances(Tokenizer(text, English_ignore_words_map))
	for word, count := range token_occur {
		reply := RedisClient.HGet(Redis_prefix + categories, word)

		cur, _ := strconv.ParseUint(string(reply.Val()), 10, 0)
		if cur != 0 {
			inew := cur - uint64(count)
			if inew > 0 {
				RedisClient.HSet(Redis_prefix + categories, word, strconv.Itoa(int(inew)))
			} else {
				RedisClient.HDel(Redis_prefix + categories, word)
			}
		}
	}

	if Tally(categories) == 0 {
		RedisClient.Del(Redis_prefix + categories)
		RedisClient.SRem(Redis_prefix + "categories", categories)
	}
}

func Classify(text string) (key string) {
	scores := Score(text)
	max := 0.0
	if scores != nil {
		for k, v := range scores {
			if v <= max {
				max = v
				key = k
			}
		}
		return
	}
	key = "I dont know"
	return
}

func Score(text string) (res map[string]float64) {

	detectedLang, err := lang.Guess(text)
	if err != nil {
		detectedLang = "en"
		return
	}

	token_occur := GetOccurances(detectedLang, text)

	//token_occur := Occurances(English_tokenizer(text))
	Debug("Score token_occur, ", token_occur)
	res = make(map[string]float64)

	reply := RedisClient.SMembers(Redis_prefix + "categories")

	Debug("Score reply, ", reply)
	for v1, category := range reply.Val() {
		Debug("Score range reply.Val() ", v1, category)
		tally := Tally(category)
		Debug("Score tally, ", tally)
		if tally == 0 {
			continue
		}

		res[category] = 0.0
		for word, v := range token_occur {
			Debug("Score, range token_occur, ", word, v)

			Debug("Score, RedisClient.HGet ", Redis_prefix + category, word)
			score := RedisClient.HGet(Redis_prefix + category, word)
			Debug("Score, RedisClient.HGet result ", score.Val())

			if score == nil {
				continue
			}

			targetVal := score.Val()
			if targetVal == "" {
				continue
			}

			iVal, err := strconv.ParseFloat(targetVal, 64)
			if err != nil {
				Debug("Error: Score, ", err)
				return nil
			}

			Debug("Score, ival ", iVal)

			if iVal == 0.0 {
				iVal = correction
			}
			res[category] += math.Log(iVal / float64(tally))
			Debug("Score, res[category], ", res[category])
		}
	}

	return res
}

var supportedLang []string = []string{
	"pt",
	"pt_PT",
	"pt_BR",
	"fr",
	"it",
	"es",
}

func GetOccurances(lang, text string) (counts map[string]uint) {

	//check if lang is supported
	supported := false
	for _, v := range supportedLang {
		if v == lang {
			supported = true
		}
	}

	if !supported {
		Debug("WARN: Will use default lang EN ", lang, supportedLang)
		counts = Occurances(Tokenizer(text, English_ignore_words_map))
		return
	}

	//identify what will be the tokenizer used
	if strings.ContainsAny(lang, "pt") {
		counts = Occurances(Tokenizer(text, Portuguese_ignore_words_map))
	}

	//identify what will be the tokenizer used
	if lang == "es" {
		counts = Occurances(Tokenizer(text, Spanish_ignore_words_map))
	}

	if lang == "it" {
		counts = Occurances(Tokenizer(text, Italian_ignore_words_map))
	}

	if lang == "fr" {
		counts = Occurances(Tokenizer(text, French_ignore_words_map))
	}

	return
}

func Tally(category string) (sum uint64) {
	vals := RedisClient.HVals(Redis_prefix + category)

	for _, val := range vals.Val() {
		iVal, err := strconv.ParseUint(string(val), 10, 0)
		if err != nil {
			Debug("Error: Tally, ", err)
			return
		}

		sum += iVal
	}

	return sum
}

// init function, load the configs
// fill english_ignore_words_map
func init() {
	if REDIS_HOST == "" {
		REDIS_HOST = "g7-box"
	}
	if REDIS_PORT == "" {
		REDIS_PORT = "6379"
	}
	if REDIS_PASSWORD == "" {
	}

	StartLanguageIgnore()


	// get redis connection info
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     REDIS_HOST + ":" + REDIS_PORT,
		Password: REDIS_PASSWORD, // no password set
		DB:       0, // use default DB
	})

	pong, err := RedisClient.Ping().Result()
	if err != nil {
		Debug("Error: init, Can't connect to Redis Server", err)
		panic("Can't connect to Redis Server")
	}
	Debug(pong)

	//train with world know spam words
	if LEARNSPAM == "true" {
		StartLanguageSpam("config_spamwords_en.yaml", "english_spam")
	}
}

func StartLanguageSpam(cfg_filename string, targetIgnore string){
	Debug("StartLanguageSpam init")
	config, err := yaml.ReadFile(cfg_filename)
	if err != nil {
		log.Fatalf("readfile(%s): %s", cfg_filename, err)
		Debug("Error: init, Can't readfile ", cfg_filename, err)
		panic("Error: init, Can't readfile ")
	}
	to_ignore, err := config.Get(targetIgnore)
	if err != nil {
		Debug("Error: init, Can't parse " + targetIgnore, to_ignore, err)
		panic("Error: init, Can't parse " + targetIgnore)
	}

	// get each separated words
	ignore_words_list := strings.Fields(to_ignore)
	for _, word := range ignore_words_list {
		word = strings.TrimSpace(word)
		Debug("StartLanguageSpam Train, ",word)
		Train("bad",word);
	}
	Debug("StartLanguageSpam end")

}

func StartLanguageIgnore() {

	/** BEGIN ENGLISH **/
	SetConfigs("config_en.yaml", "english_ignore", English_ignore_words_map)
	/** END ENGLISH **/

	/** BEGIN PORTUGUESE **/
	SetConfigs("config_pt.yaml", "portuguese_ignore", Portuguese_ignore_words_map)
	/** END PORTUGUESE **/

	/** BEGIN SPANISH **/
	SetConfigs("config_es.yaml", "spanish_ignore", Spanish_ignore_words_map)
	/** END SPANISH **/

	/** BEGIN ITALIAN **/
	SetConfigs("config_it.yaml", "italian_ignore", Italian_ignore_words_map)
	/** END ITALIAN **/

	/** BEGIN FRENCH **/
	SetConfigs("config_fr.yaml", "french_ignore", French_ignore_words_map)
	/** END FRENCH **/

}

func SetConfigs(cfg_filename string, targetIgnore string, ignore_words_map map[string]int) {
	config, err := yaml.ReadFile(cfg_filename)
	if err != nil {
		log.Fatalf("readfile(%s): %s", cfg_filename, err)
		Debug("Error: init, Can't readfile ", cfg_filename, err)
		panic("Error: init, Can't readfile ")
	}
	to_ignore, err := config.Get(targetIgnore)
	if err != nil {
		Debug("Error: init, Can't parse " + targetIgnore, to_ignore, err)
		panic("Error: init, Can't parse " + targetIgnore)
	}

	// get each separated words
	ignore_words_list := strings.Fields(to_ignore)
	for _, word := range ignore_words_list {
		word = strings.TrimSpace(word)
		ignore_words_map[word] = 1
	}
}

func Debug(v ...interface{}) {
	if SPAM_DEBUG == "true" {
		log.Println(v)
	}
}