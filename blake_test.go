// Quick benchmarks for blake2b hashing as used in Filecoin protocol 1 address generation.
package blakebench

import (
	"crypto/rand"
	"hash"
	"testing"

	minio_blake "github.com/minio/blake2b-simd"
	xcrypto_blake "golang.org/x/crypto/blake2b"
)

const (
	payloadDataLen  = 65 // uncompressed secp256k1 public key
	payloadHashLen  = 20 // blake2b-160
	checksumDataLen = 21 // protocol byte + payload hash
	checksumHashLen = 4  // blake2b-32
)

var (
	testPayloadData  [payloadDataLen]byte
	testChecksumData [checksumDataLen]byte
)

func init() {
	_, _ = rand.Read(testPayloadData[:])
	_, _ = rand.Read(testChecksumData[:])
}

// generic hash sum benchmark helper to reduce benchmark boilerplate
func benchHashSum(b *testing.B, hasher hash.Hash, hashLen int, data []byte) {
	b.Helper()

	sum := make([]byte, 0, hashLen)
	for i := 0; i < b.N; i++ {
		hasher.Reset()
		hasher.Write(data)
		sum = sum[:0]
		hasher.Sum(sum)
	}
}

func BenchmarkXCryptoBlake(b *testing.B) {
	b.Run("fcaddr-payload", func(b *testing.B) {
		hasher, err := xcrypto_blake.New(payloadHashLen, nil)
		if err != nil {
			b.Fatal(err)
		}

		benchHashSum(b, hasher, payloadHashLen, testPayloadData[:])
	})

	b.Run("fcaddr-chksum", func(b *testing.B) {
		hasher, err := xcrypto_blake.New(checksumDataLen, nil)
		if err != nil {
			b.Fatal(err)
		}

		benchHashSum(b, hasher, checksumDataLen, testChecksumData[:])
	})
}

func BenchmarkMinioBlakeSIMD(b *testing.B) {
	b.Run("fcaddr-payload", func(b *testing.B) {
		hasher, err := minio_blake.New(&minio_blake.Config{Size: payloadHashLen})
		if err != nil {
			b.Fatal(err)
		}

		benchHashSum(b, hasher, payloadHashLen, testPayloadData[:])
	})

	b.Run("fcaddr-chksum", func(b *testing.B) {
		hasher, err := minio_blake.New(&minio_blake.Config{Size: checksumDataLen})
		if err != nil {
			b.Fatal(err)
		}

		benchHashSum(b, hasher, checksumDataLen, testChecksumData[:])
	})
}
