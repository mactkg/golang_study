package lengthconv

import "fmt"

type Meter float64
type Feet float64

func (m Meter) String() string {
	return fmt.Sprintf("%gm", m)
}

func (ft Feet) String() string {
	return fmt.Sprintf("%gft", ft)
}
