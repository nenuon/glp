package conv

import "fmt"

// 温度
type Celsius float64
type Fahrenheit float64

// 長さ
type Feet float64
type Metre float64

// 重さ
type Pound float64
type Kilogramme float64

func (c Celsius) String() string    { return fmt.Sprintf("%g℃", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func (f Feet) String() string  { return fmt.Sprintf("%gft", f) }
func (m Metre) String() string { return fmt.Sprintf("%gm", m) }

func FToM(f Feet) Metre { return Metre(f / 3.28) }
func MToF(m Metre) Feet { return Feet(m * 3.28) }

func (p Pound) String() string      { return fmt.Sprintf("%glb", p) }
func (k Kilogramme) String() string { return fmt.Sprintf("%gkg", k) }

func PToK(p Pound) Kilogramme { return Kilogramme(p * 0.453) }
func KToP(k Kilogramme) Pound { return Pound(k / 0.453) }
