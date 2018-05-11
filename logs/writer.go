package main

import (
	"io"
	"sync"
	"time"
)

type logWriter struct {
	sync.Mutex
	writer io.Writer
}

func newLogWriter(wr io.Writer) *logWriter {
	return &logWriter{writer: wr}
}

func (lg *logWriter) write(when time.Time, msg string) {
	lg.Lock()
	defer lg.Unlock()
	h := append(append([]byte{'['}, formatTimeShort(when)...), ']')
	lg.writer.Write(append(append(append(h, ' '), msg...), '\n'))
}

const (
	y1  = `0123456789`
	y2  = `0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789`
	y3  = `0000000000111111111122222222223333333333444444444455555555556666666666777777777788888888889999999999`
	y4  = `0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789`
	mo1 = `000000000111`
	mo2 = `123456789012`
	d1  = `0000000001111111111222222222233`
	d2  = `1234567890123456789012345678901`
	h1  = `000000000011111111112222`
	h2  = `012345678901234567890123`
	mi1 = `000000000011111111112222222222333333333344444444445555555555`
	mi2 = `012345678901234567890123456789012345678901234567890123456789`
	s1  = `000000000011111111112222222222333333333344444444445555555555`
	s2  = `012345678901234567890123456789012345678901234567890123456789`
	ns1 = `0123456789`
)

func formatTimeBase(when time.Time) []byte {
	y, mo, d := when.Date()
	h, mi, s := when.Clock()
	ns := when.Nanosecond() / 1000000
	//len("20180511180445001")==17
	var buf [17]byte

	buf[0] = y1[y/1000%10]
	buf[1] = y2[y/100]
	buf[2] = y3[y-y/100*100]
	buf[3] = y4[y-y/100*100]
	buf[4] = mo1[mo-1]
	buf[5] = mo2[mo-1]
	buf[6] = d1[d-1]
	buf[7] = d2[d-1]
	buf[8] = h1[h]
	buf[9] = h2[h]
	buf[10] = mi1[mi]
	buf[11] = mi2[mi]
	buf[12] = s1[s]
	buf[13] = s2[s]
	buf[14] = ns1[ns/100]
	buf[15] = ns1[ns%100/10]
	buf[16] = ns1[ns%10]

	return buf[0:]
}

func formatTimeLong(when time.Time) []byte {
	base := formatTimeBase(when)
	//len("01/02/2018 15:04:05.123")==23
	var buf [23]byte

	buf[0] = base[4] //mo1
	buf[1] = base[5] //mo2
	buf[2] = '/'
	buf[3] = base[6] //d1
	buf[4] = base[7] //d2
	buf[5] = '/'
	buf[6] = base[0] //y1
	buf[7] = base[1] //y2
	buf[8] = base[2] //y3
	buf[9] = base[3] //y4
	buf[10] = ' '
	buf[11] = base[8] //h1
	buf[12] = base[9] //h2
	buf[13] = ':'
	buf[14] = base[10] //mi1
	buf[15] = base[11] //mi2
	buf[16] = ':'
	buf[17] = base[12] //s1
	buf[18] = base[13] //s2
	buf[19] = '.'
	buf[20] = base[14] //ns1
	buf[21] = base[15] //ns2
	buf[22] = base[16] //ns3

	return buf[0:]
}

func formatTimeShort(when time.Time) []byte {
	base := formatTimeBase(when)
	//len("05/11/2018 18:14:11")==19
	var buf [19]byte

	buf[0] = base[4] //mo1
	buf[1] = base[5] //mo2
	buf[2] = '/'
	buf[3] = base[6] //d1
	buf[4] = base[7] //d2
	buf[5] = '/'
	buf[6] = base[0] //y1
	buf[7] = base[1] //y2
	buf[8] = base[2] //y3
	buf[9] = base[3] //y4
	buf[10] = ' '
	buf[11] = base[8] //h1
	buf[12] = base[9] //h2
	buf[13] = ':'
	buf[14] = base[10] //mi1
	buf[15] = base[11] //mi2
	buf[16] = ':'
	buf[17] = base[12] //s1
	buf[18] = base[13] //s2

	return buf[0:]
}
