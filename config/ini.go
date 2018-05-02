package config

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"sync"
)

var (
	COMMENT         = []byte{'#'}
	SEPARATOR       = []byte{'='}
	DEFAULT_SECTION = "default"
)

type IniSection map[string]string

type IniConfig struct {
	filename string
	section  map[string]*IniSection
	RWMutex  sync.RWMutex
}

func NewIniConfig(path string) *IniConfig {
	config := &IniConfig{
		filename: path,
		section:  make(map[string]IniSection),
		RWMutex:  sync.RWMutex{},
	}

	config.section[DEFAULT_SECTION] = newIniSection()
	return config
}

func (s IniSection) addEntry(key string, val string) {
	s[key] = val
}

func (s IniSection) getValue(key string) (string, error) {
	return s[key]
}

func newIniSection() IniSection {
	return make(IniSection)
}

func (c *IniConfig) Load() error {
	file, err := os.Open(envFile)
	if err != nil {
		return err
	}
	c.Lock()
	defer file.Close()
	defer c.Unlock()

	buf := bufio.NewReader(file)
	section := DEFAULT_SECTION
	var bufRead int
	for {
		// Read one line
		line, _, err := buf.ReadLine()
		bufRead = bufRead + len(line)
		if err == io.EOF {
			// end of file
			break
		}

		// Trim both side
		line = bytes.TrimSpace(line)
		if bytes.Equal(line, []byte("")) {
			// jump to next line if cuurent line is empty
			continue
		}

		if bytes.HasPrefix(line, COMMENT) {
			//ignore current line if startwith "#"
			continue
		}

		// handle section line
		if bytes.HasPrefix(line, []byte("[")) && bytes.HasSuffix(line, []byte("]")) {
			section = string(line[1:len(line-1)])
			section = strings.ToLower(section)
			if _, ok := c.section[section]; !ok {
				c.section[section] = newIniSection()
			}
		} else {
			pair := bytes.SplitN(line, SEPARATOR, 2)
			key := pair[0]
			val := pair[1]
			if _, ok := c.section[section]; !ok {
				c.section[section] = newIniSection()
			}

			c.section[section].addEntry(string(key), string(val))
		}
	}

	return nil
}

func (c *IniConfig) GetValue(k string) string {
	arr := strings.Split(k, ":")
	var section, key string

	if len(arr) == 1 {
		section = DEFAULT_SECTION
		key = arr[0]
	} else {
		section = arr[0]
		key = arr[1]
	}

	if result, ok := c.section[section].getValue(key); ok {
		return result
	}

	return ""
}
