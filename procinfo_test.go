package androidutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemoryInfo(t *testing.T) {
	testdata := []byte(`MemTotal:        3690332 kB
		MemFree:          106332 kB
		MemAvailable:    1582960 kB
		Buffers:           54648 kB
		Cached:          1589388 kB
		SwapCached:         1264 kB
		Active:          1592056 kB
		Inactive:        1296872 kB
		Active(anon):    1230568 kB
		Inactive(anon):    43976 kB
		Active(file):     361488 kB
		Inactive(file):  1252896 kB
		Unevictable:         256 kB
		Mlocked:             256 kB
		SwapTotal:       1572860 kB
		SwapFree:        1561000 kB
		Dirty:                 0 kB
		Writeback:             0 kB
		AnonPages:       1245236 kB
		Mapped:           480356 kB
		Shmem:             29672 kB
		Slab:             138480 kB
		SReclaimable:      44712 kB
		SUnreclaim:        93768 kB
		KernelStack:       33152 kB
		PageTables:        48000 kB
		NFS_Unstable:          0 kB
		Bounce:                0 kB
		WritebackTmp:          0 kB
		CommitLimit:     3418024 kB
		Committed_AS:   95129720 kB
		VmallocTotal:   258998208 kB
		VmallocUsed:      149436 kB
		VmallocChunk:   258718692 kB`)
	info, err := parseMemoryInfo(testdata)
	assert.NoError(t, err)
	assert.NotNil(t, info)
	assert.Equal(t, 106332, info["MemFree"])
	assert.Equal(t, 43976, info["Inactive(anon)"])
	assert.Equal(t, 0, info["NFS_Unstable"])
	assert.Equal(t, 258998208, info["VmallocTotal"])
}

func TestCpuInfo(t *testing.T) {
	testdata := []byte(`Processor       : AArch64 Processor rev 4 (aarch64)
	processor       : 0
	BogoMIPS        : 38.40
	Features        : pmull sha1 sha2 crc32
	CPU implementer : 0x41
	CPU architecture: 8
	CPU variant     : 0x0
	CPU part        : 0xd03
	CPU revision    : 4
	
	processor       : 1
	BogoMIPS        : 38.40
	Features        : fp aes
	CPU implementer : 0x41
	CPU architecture: 8
	CPU variant     : 0x0
	CPU part        : 0xd03
	CPU revision    : 4
	
	Hardware        : Qualcomm Technologies, Inc MSM8940`)
	hardware, processors, err := parseCpuInfo(testdata)
	assert.NoError(t, err)
	assert.Equal(t, "Qualcomm Technologies, Inc MSM8940", hardware)
	assert.Equal(t, 2, len(processors))
	assert.Equal(t, "38.40", processors[1].BogoMIPS)
	assert.Equal(t, []string{"fp", "aes"}, processors[1].Features)
	assert.Equal(t, []string{"pmull", "sha1", "sha2", "crc32"}, processors[0].Features)
}
