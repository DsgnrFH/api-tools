package id

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"sync"
	"time"
)

type UUIDv7Generator struct {
	mu           sync.Mutex
	lastMS       int64
	sequenceCounter uint16
}

var defaultGen = &UUIDv7Generator{}

func GenerateUUIDv7() string {
	return defaultGen.Next()
}

func (g *UUIDv7Generator) Next() string {
	g.mu.Lock()
	defer g.mu.Unlock()

	nowMS := time.Now().UnixMilli()

	// Handle clock drift or multiple generations in the same millisecond slot
	if nowMS <= g.lastMS {
		nowMS = g.lastMS
		g.sequenceCounter++
	} else {
		g.lastMS = nowMS
		g.sequenceCounter = 0
	}

	var uuid [16]byte

	binary.BigEndian.PutUint64(uuid[0:8], uint64(nowMS))
	copy(uuid[0:6], uuid[2:8])

	// UUIDv7 requires bits 4-7 of byte 6 to be the version (0111 -> 0x70)
	binary.BigEndian.PutUint16(uuid[6:8], g.sequenceCounter)
	uuid[6] = (uuid[6] & 0x0F) | 0x70 

	_, _ = rand.Read(uuid[8:16])
	// UUID variant 1 requires bits 6-7 of byte 8 to be 10xxxxxx (0x80)
	uuid[8] = (uuid[8] & 0x3F) | 0x80

	// Format into standard 8-4-4-4-12 hex string layout
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}
