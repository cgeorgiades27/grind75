package validanagram

/*
#include <string.h>
#include <stdbool.h>

#define NUM_ASCII_CHARS 128

bool valid_anagram(char *s, char *t)
{
    int len_s = strlen(s);
    int len_t = strlen(t);

    if (len_s != len_t)
        return false;

    // assign all entries to 0
    int arr[NUM_ASCII_CHARS] = {0};
    for (int i = 0; i < len_s; ++i)
    {
        arr[s[i]]++;
        arr[t[i]]--;
    }

    for (int i = 0; i < NUM_ASCII_CHARS; ++i)
    {
        if (arr[i] != 0)
            return false;
    }

    return true;
}
*/
import "C"

const MaxAsciiChar = 128

func ValidAnagramCG(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	slc := make([]uint32, MaxAsciiChar)
	for i := range s {
		slc[s[i]]++
		slc[t[i]]--
	}

	for _, elem := range slc {
		if elem != 0 {
			return false
		}
	}

	return true
}

func ValidAnagramC(s, t string) bool {
	cs := C.CString(s)
	ct := C.CString(t)
	return bool(C.valid_anagram(cs, ct))
}
