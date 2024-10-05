package eacnet

import (
	"crypto/sha256"
	"encoding/binary"
	"time"
)

type Salt []byte

var (
	KonasuteClientSalt  Salt = []byte("d4BK3JFREkH5WuyTVEJQ2jbS9h2-df4D")
	KonasuteServerSalt  Salt = []byte("tGhCtgLuTjV7cZ2phWuCpQ8iwSypVn4W")
	InfinitasClientSalt Salt = []byte("fAwHp6G2FLPHN_ZGBhREJG5flt3hNu")
	InfinitasServerSalt Salt = []byte("NH_P-urkCV9npxR90kaAR7YnqDTRL-")
)

func (s Salt) Sign(data []byte) []byte {
	return s.SignWithTime(data, time.Now())
}

func (s Salt) SignWithTime(data []byte, t time.Time) []byte {
	hash := sha256.New()
	hash.Write(data)
	hash.Write(binary.BigEndian.AppendUint64(nil, uint64(t.Unix()/60)))
	hash.Write(s)
	return hash.Sum(nil)
}
