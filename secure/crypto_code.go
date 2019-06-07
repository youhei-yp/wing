// Copyright (c) 2018-2019 Dunyu All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2019/05/22   yangping       New version
// -------------------------------------------------------------------

package secure

import (
	"fmt"
	"github.com/youhei-yp/wing/invar"
	"math/rand"
	"time"
)

/*
 * Useage as follows:
 *
 * [CODE:]
 * coder := mvc.NewSoleCoder(nil)
 * code, _ := coder.Gen(6)
 * logger.I("6 chars code:", code)
 *
 * code, _ = coder.Gen(8)
 * logger.I("8 chars code:", code)
 *
 * code, _ := coder.Gen(6, 5)
 * logger.I("max retry 5 times, 6 chars code:", code)
 *
 * code, _ = coder.Gen(8, 5)
 * logger.I("max retry 5 times, 8 chars code:", code)
 * [CODE]
 */

// SoleCoder random coder to generate unique number code
type SoleCoder struct {
	codes map[string]bool
}

// NewSoleCoder create SoleCoder and init with exist codes
func NewSoleCoder(data []string) *SoleCoder {
	coder := &SoleCoder{
		codes: make(map[string]bool),
	}
	if data != nil && len(data) > 0 {
		for _, code := range data {
			coder.codes[code] = true
		}
	}
	return coder
}

// Gen generate a given length number code, it may throw a error
// when over the retry times
func (c *SoleCoder) Gen(codelen int, times ...int) (string, error) {
	if c.codes == nil {
		c.codes = make(map[string]bool)
	}

	radix := 1
	for i := 0; i < codelen; i++ {
		radix *= 10
	}

	if len(times) > 0 {
		for i := 0; i < times[0]; i++ {
			code, err := c.innerGenerate(codelen, radix)
			if err != nil {
				continue
			}
			return code, nil
		}
		return "", invar.ErrOverTimes
	}
	return c.innerGenerate(codelen, radix)
}

// Remove remove used sole code outof cache
func (c *SoleCoder) Remove(code string) {
	if c.codes != nil {
		if ok, _ := c.codes[code]; ok {
			c.codes[code] = false
		}
	}
}

// innerGenerate generate a give length number code
func (c *SoleCoder) innerGenerate(codelen, radix int) (string, error) {
	rand.Seed(time.Now().UnixNano())
	format := "%0" + fmt.Sprintf("%d", codelen) + "d"
	code := fmt.Sprintf(format, rand.Intn(radix))

	// check generated code if it unique
	if ok, cc := c.codes[code]; ok && cc {
		return "", invar.ErrDupData
	}
	c.codes[code] = true
	return code, nil
}
