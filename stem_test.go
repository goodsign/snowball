package snowball

import (
    "os"
    "io"
    "fmt"
    "bufio"
    "testing"
    "regexp"
)

var (
    WordStemLineRx *regexp.Regexp = regexp.MustCompile(`\A(?P<word>\S+)\s*(?P<stem>\S+).*\z`)
)

type wordStem struct {
    word          string
    expectedStem  string 
}


func testLanguage(t *testing.T, language string, encoding string, filename string) {
    // Create stemmer
    stemmer, err := NewWordStemmer(language, encoding)

    if nil != err {
        t.Fatalf("Cannot create stemmer: %s", err)
    }
    defer stemmer.Close()

    // Open test data file
    f, err := os.Open(filename)

    if nil != err {
        t.Fatalf("Cannot open word stems file: %s", err)
    }

    // Open reader on it
    r := bufio.NewReader(f)

    var e error

    // Read while any error occurs
    for {
        var line []byte 

        // Read line
        line, _, e = r.ReadLine()

        if nil != e {
            break
        }

        // Extract data using regexp pattern
        matches := WordStemLineRx.FindStringSubmatch(string(line))

        if 3 != len(matches) {
            e = fmt.Errorf("Incorrect line in file '%s':'%s'", filename, line)
            break
        }

        // 0 is the whole line, 1 is 'word', 2 is 'stem'
        testItem := wordStem{matches[1], matches[2]}

        // Stem the word using stemmer
        stm, err := stemmer.Stem([]byte(testItem.word))

        if nil != err {
            t.Error(err)
            continue
        }

        // Compare stemmer result and expected result from file.
        if string(stm) != testItem.expectedStem {
            t.Errorf("Language: '%s' (%s) Word: '%s' Expected stem: '%s' Got stem: '%s'", 
                     language,
                     DefaultEncoding,
                     testItem.word, 
                     testItem.expectedStem, 
                     stm)
        }
    }

    if nil != e && io.EOF != e {
        t.Fatal(e)
    }
}

func TestUTF8(t *testing.T) {
    testLanguage(t, "russian",      DefaultEncoding, "test/rus_test.txt")
    testLanguage(t, "danish",       DefaultEncoding, "test/danish_test.txt")
    testLanguage(t, "dutch",        DefaultEncoding, "test/dutch_test.txt")
    testLanguage(t, "english",      DefaultEncoding, "test/english_test.txt")
    testLanguage(t, "finnish",      DefaultEncoding, "test/finnish_test.txt")
    testLanguage(t, "french",       DefaultEncoding, "test/french_test.txt")
    testLanguage(t, "german",       DefaultEncoding, "test/german_test.txt")
    testLanguage(t, "hungarian",    DefaultEncoding, "test/hungarian_test.txt")
    testLanguage(t, "italian",      DefaultEncoding, "test/italian_test.txt")
    testLanguage(t, "norwegian",    DefaultEncoding, "test/norwegian_test.txt")
    testLanguage(t, "portuguese",   DefaultEncoding, "test/portuguese_test.txt")
    testLanguage(t, "romanian",     DefaultEncoding, "test/romanian_test.txt")
    testLanguage(t, "spanish",      DefaultEncoding, "test/spanish_test.txt")
    testLanguage(t, "swedish",      DefaultEncoding, "test/swedish_test.txt")
}

func TestKOI8R(t *testing.T) {
    testLanguage(t, "russian", "KOI8_R", "test/rus_koi8r_test.txt")
}