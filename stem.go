// Package stem is a Snowball stemmer cgo port. Provides word stem extraction functionality.
package stem

// #include "include/libstemmer.h"
// #include <stdlib.h>
import "C"
import (
    "sync"
    "fmt"
    "unsafe"
)

const (
    MaxAllowedStemLength = 80
    DefaultEncoding = "UTF_8"
)

// WordStemmer provides Snowball stemming functionality.
type WordStemmer struct {
    stemmer    *C.Stemmer  // Snowball struct needed for detection (actually a typedef for struct sb_stemmer)
    mutex      sync.Mutex  // Mutex used to guarantee thread safety for Snowball calls
    algorithm   string
    encoding   string
}

// Creates new word stemmer with specified algorithm and encoding identifiers. If it is successfully created, it
// must be closed as it needs to free native Snowball resources.
//
// NOTE: Check libstemmer/modules.txt for allowed names of algorithms and encodings.
func NewWordStemmer(algorithm string, encoding string) (*WordStemmer, error) {
    ws := new(WordStemmer)

    algStr := C.CString(algorithm)
    encStr := C.CString(encoding)

    defer C.free(unsafe.Pointer(algStr))
    defer C.free(unsafe.Pointer(encStr))

    ws.stemmer = C.sb_stemmer_new(algStr, encStr)

    if nil == ws.stemmer {
        return nil, fmt.Errorf("Cannot create word stemmer with algorithm: '%s' and encoding: '%s'", algorithm, encoding)
    }

    ws.algorithm = algorithm
    ws.encoding = encoding

    return ws, nil
}

// Stem extracts word's stem. Language/encoding of the word
// should match the algorithm/encoding of the created stemmer.
func (ws *WordStemmer) Stem(word []byte) ([]byte, error) {
    ws.mutex.Lock()
    defer ws.mutex.Unlock()

    wordCString := C.CString(string(word))
    defer C.free(unsafe.Pointer(wordCString))

    // Stem the word
    stemResult := C.sb_stemmer_stem(ws.stemmer, 
                                    (*C.sb_symbol)(unsafe.Pointer(wordCString)), 
                                    C.int(len(word)));
 
    if nil == stemResult {
        return nil, fmt.Errorf("Stemmer for (%s;%s) cannot extract stem for word: '%s'",
                              ws.algorithm,
                              ws.encoding,
                              word)
    }

    stemLen := C.sb_stemmer_length(ws.stemmer)

    if stemLen <= 0 || 
       stemLen > MaxAllowedStemLength {
        return nil, fmt.Errorf("Stemmer for (%s;%s) got incorrect stem length for word: '%s': len = '%d'",
                              ws.algorithm,
                              ws.encoding,
                              word,
                              stemLen)
    }

    return C.GoBytes(unsafe.Pointer(stemResult), stemLen), nil
}

// Close frees native C resources
func (ws *WordStemmer) Close() error {
    ws.mutex.Lock()
    defer ws.mutex.Unlock()

    if nil != ws.stemmer {
        C.sb_stemmer_delete(ws.stemmer)
        ws.stemmer = nil
        return nil
    }

    return fmt.Errorf("Cannot close: stemmer is not valid.")
}
