package strings

import "math/rand"

// RandStr generates a pseudo random string with the provided lenght
func RandStr(seed int64, length int) string {
	r := rand.New(rand.NewSource(seed))

	ran_str := make([]byte, 0)

	for i := 0; i < length; i++ {
		ran_str = append(ran_str, byte(65+r.Intn(25)))
	}

	return string(ran_str)
}
