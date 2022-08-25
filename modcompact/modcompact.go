/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License. You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package modcompact

// By Rensmt
import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/bashery/linethrift"
)

type TMoreCompactProtocol struct {
	__a        []byte
	__b        [][]string
	__c        []byte
	__d        []byte
	__e        []string
	__f        []byte
	__h        []byte
	__last_fid int
	__last_pos int
	__last_sid int
	data       []byte
	res        interface{}
	err        bool
}

type ExceptionMod struct {
	Code   int64  `json:"1"`
	Reason string `json:"2"`
}

func TMoreCompactProtocolGoods(decodedByte []byte) *TMoreCompactProtocol {
	A := new(TMoreCompactProtocol)
	A.func_A()
	A.funcD(decodedByte)
	return A
}

func (p *TMoreCompactProtocol) readVarint(data []byte) []int {
	var (
		result = 0
		shift  = 0
		i      = 0
	)
	for {
		byteh := int(data[i])
		i++
		result |= (byteh & 0x7f) << shift
		if byteh>>7 == 0 {
			return []int{result, i}
		}
		shift += 7
	}
}

func (p *TMoreCompactProtocol) readCollectionBegin(x []byte) (typea int, size int, length int) {
	size_type := int(x[0])
	size = size_type >> 4
	typea = size_type & 0x0f
	length = 0
	if size == 15 {
		owsh := p.readVarint(x[1:])
		size, length = owsh[0], owsh[1]
	}
	return typea, size, length + 1
}

func (p *TMoreCompactProtocol) funcB() int {
	var (
		i2 = 0
		i3 = 0
	)
	for {
		//l2, _ := strconv.Atoi(string(p.data[p.__last_pos]))
		l2 := int(p.data[p.__last_pos])
		p.__last_pos++
		i2 |= (l2 & 127) << i3
		if (l2 & 128) != 128 {
			return i2
		}
		i3 += 7
	}
}

func (p *TMoreCompactProtocol) funcC(pX int, i2 int) []byte {
	if i2 == 0 {
		return []byte{}
	}
	bArr := p.data[pX : pX+i2]
	return bArr
}

func (p *TMoreCompactProtocol) funcD(d []byte) {
	p.data = d
	p.funcT()
}

func (p *TMoreCompactProtocol) funcE() {
	var a interface{}
	fid := p.funcY()
	if fid == 0 {
		return
	} else if fid == 1 {
		a = p.funcG(p.funcW())
		p.res = a
	} else if fid == 2 {
		a = p.funcG(p.funcW())
		p.res = a
		p.err = true
	} else if fid == 6 { //exception
		return
	} else { //EOF
		return
	}
}

func (p *TMoreCompactProtocol) funcF(n int) int {
	return (n >> 1) ^ -(n & 1)
}

func (p *TMoreCompactProtocol) T2() bool {
	b := p.funcB()
	if b == 0 {
		return false
	}
	return true
}

func (p *TMoreCompactProtocol) funcG(t int) interface{} {
	var (
		a interface{}
		b = 0
		c = 0
	)
	if t == 2 {
		b = p.funcB()
		if b == 0 {
			a = false
		} else {
			a = true
		}
	} else if t == 3 {
		bx := bytes.NewReader(p.data[p.__last_pos:])
		x, _ := bx.ReadByte()
		a = int(x)
		p.__last_pos++
	} else if t == 4 {
		p.__last_pos += 8
		a = ""
	} else if t == 8 {
		a = p.funcF(p.funcX(p.data[p.__last_pos:]))
	} else if t == 10 {
		a = p.funcF(p.funcB())
	} else if t == 11 {
		a = p.funcS()
		return a
	} else if t == 12 {
		an := make(map[string]interface{})
		b := p.funcB()
		c := p.funcN(b)
		for _, d := range c { //d = int of operation type
			an[strconv.Itoa(d)] = p.funcG(p.funcW())
		}
		a = an
	} else if t == 13 {
		var an = make(map[string]string)
		c := p.funcB()
		if c != 0 {
			d := p.funcY()
			t1, t2 := p.funcQ(int(d))
			for i := 0; i < c; i++ {
				k := fmt.Sprint(p.funcG(t1))
				v := fmt.Sprint(p.funcG(t2))
				an[k] = v
			}
		}
		a = an
	} else if t == 14 || t == 15 {
		var an []interface{}
		ftype, count, offset := p.readCollectionBegin(p.data[p.__last_pos:])
		p.__last_pos += offset
		for i := 0; i < count; i++ {
			b := p.funcG(p.func_D(ftype)) // GODD!!! THIS OPERATION
			an = append(an, b)
		}
		a = an
	} else if t == 16 {
		b = p.funcB()
		c = -(b & 1) ^ p.func_E(b, 1)
		p.__last_sid += c
		a = strconv.Itoa(p.__last_sid)
	} else if t == 17 {
		b = p.funcB()
		if len(p.__e) > b {
			a = p.__e[b]
		}
	}
	return a
}

func (p *TMoreCompactProtocol) funcM() {
	a := p.funcB()
	for _a := 0; _a < a; _a++ {
		bArr := []byte{p.data[p.__last_pos]}
		bArr = append(bArr, p.func_C(p.data[p.__last_pos+1:p.__last_pos+17])...)
		p.__e = append(p.__e, string(bArr))
		p.__last_pos += 17
	}
	p.funcE()
}

func (p *TMoreCompactProtocol) funcN(d int) []int {
	var (
		a = []int{}
		i = 0
	)
	for {
		b := 1 << i
		if b > d {
			break
		} else if d&b != 0 {
			a = append(a, i)
		}
		i++
	}
	return a
}

func (p *TMoreCompactProtocol) funcQ(d int) (int, int) {
	return p.func_D(d >> 4), p.func_D(d & 15)
}

func (p *TMoreCompactProtocol) funcS() string {
	var (
		a = p.funcB()
		b = p.data[p.__last_pos : p.__last_pos+a]
	)
	p.__last_pos += a
	return string(b)
}

func (p *TMoreCompactProtocol) funcT() {
	p.__last_pos = 3
	if len(p.data) == 4 {
		fmt.Println("Data (error: 20):", p.data)
	} else {
		var (
			a = p.funcB()
			b = p.funcC(p.__last_pos, a)
			d = 0
			f = 0
			g = 0
		)
		p.__d = make([]byte, a<<1)
		for _, h := range b {
			var (
				_a = 0
				_b = 128
			)
			for _a < 8 {
				//hInt, err := strconv.Atoi(string(h))
				//fmt.Println(err)
				//fmt.Println("H =>", hInt)
				hInt := int(h)
				if hInt&_b == 0 {
					d = (g << 1) + 1
				} else {
					d = (g << 1) + 2
				}
				if p.__a[d] != 0 {
					if f >= len(p.__d) {
						cD := make([]byte, 4)
						for z := 0; z < 4; z++ {
							cD[z] = byte(len(p.__d))
						}
						p.__d = append(p.__d, cD...)
					}
					p.__d[f] = p.__a[d]
					f++
					g = 0
				} else {
					g = d
				}
				_b >>= 1
				_a++
			}
		}
		p.__last_pos += a
		p.funcM()
	}
}

func (p *TMoreCompactProtocol) funcW() int {
	a := p.__d[p.__last_fid]
	p.__last_fid++
	return int(a)
}

func (p *TMoreCompactProtocol) funcX(a []byte) int {
	var (
		c = 0
		d = 0
		i = 0
	)
	for {
		e := a[i]
		i++
		//ee, _ := strconv.Atoi(string(e))
		ee := int(e)
		c |= (ee & 0x7f) << d
		if e>>7 == 0 {
			p.__last_pos += i
			return c
		}
		d += 7
	}
}

func (p *TMoreCompactProtocol) funcY() byte {
	a := p.data[p.__last_pos]
	p.__last_pos++
	return a
}

func (p *TMoreCompactProtocol) funcZ() bool {
	if len(p.data) > p.__last_pos {
		return true
	}
	return false
}

func (p *TMoreCompactProtocol) func_A() {

	p.__a = make([]byte, 512)
	p.__b = make([][]string, 18)
	p.func_B([]string{"1", "0", "1", "1"}, 2)
	p.func_B([]string{"1", "0", "1", "0", "1", "0", "0", "1"}, 3)
	p.func_B([]string{"1", "0", "1", "0", "1", "0", "0", "0"}, 4)
	p.func_B([]string{"1", "0", "1", "0", "1", "1", "1"}, 6)
	p.func_B([]string{"0", "1"}, 8)
	p.func_B([]string{"0", "0"}, 10)
	p.func_B([]string{"1", "0", "1", "0", "0"}, 11)
	p.func_B([]string{"1", "1", "0", "1"}, 12)
	p.func_B([]string{"1", "0", "1", "0", "1", "1", "0"}, 13)
	p.func_B([]string{"1", "0", "1", "0", "1", "0", "1"}, 14)
	p.func_B([]string{"1", "1", "0", "0"}, 15)
	p.func_B([]string{"1", "1", "1"}, 16)
	p.func_B([]string{"1", "0", "0"}, 17)
}

func (p *TMoreCompactProtocol) func_B(cArr []string, b2 int) {
	p.__b[b2] = cArr
	i2 := 0
	for _, c2 := range cArr {
		if c2 == "0" {
			i2 = (i2 << 1) + 1
		} else if c2 == "1" {
			i2 = (i2 << 1) + 2
		}
	}
	p.__a[i2] = byte(b2)
}

func (p *TMoreCompactProtocol) func_C(val []byte) []byte {
	x := hex.EncodeToString(val)
	return []byte(x)
}

func (p *TMoreCompactProtocol) func_D(val int) int {
	if val == 0 {
		return 0
	}
	if val == 1 || val == 2 {
		return 2
	}
	if val == 3 {
		return 3
	}
	if val == 4 {
		return 6
	}
	if val == 5 {
		return 8
	}
	if val == 6 {
		return 10
	}
	if val == 7 {
		return 4
	}
	if val == 8 {
		return 11
	}
	if val == 9 {
		return 15
	}
	if val == 10 {
		return 14
	}
	if val == 11 {
		return 13
	}
	if val == 12 {
		return 12
	}
	return 404
}

func (p *TMoreCompactProtocol) func_E(val int, n int) int {
	if val >= 0 {
		val >>= n
	} else {
		val = ((val + 0x1000000000000000) >> n)
	}
	return val
}

func (p *TMoreCompactProtocol) GETOPS() (res []*linethrift.Operation, err *ExceptionMod) {
	if p.err == true {
		jsonStr, _ := json.Marshal(p.res)
		json.Unmarshal(jsonStr, &err)
	} else {
		jsonStr, _ := json.Marshal(p.res)
		json.Unmarshal(jsonStr, &res)
	}
	return res, err
}
