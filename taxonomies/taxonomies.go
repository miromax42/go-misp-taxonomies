package taxonomies

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	re      = regexp.MustCompile("[:=]") //nolint:gochecknoglobals  //global
	format3 = "%s:%s=%q"                 //nolint:gochecknoglobals  //global
	format2 = "%s:%s"                    //nolint:gochecknoglobals	//global
)

type Taxonomy struct {
	Namespace string
	Predicate string
	Value     string

	vType valueType
}

func (t Taxonomy) String() string {
	if t.Value == "" {
		return fmt.Sprintf(format2, t.Namespace, t.Predicate)
	}

	return fmt.Sprintf(format3, t.Namespace, t.Predicate, t.Value)
}

func fromString(s string) (t Taxonomy, err error) {
	split := re.Split(s, -1)
	fmt.Println(len(split))

	switch len(split) {
	case 2:
		t.Namespace = split[0]
		t.Predicate = split[1]
	case 3:
		t.Namespace = split[0]
		t.Predicate = split[1]
		t.Value = strings.Trim(split[2], `\"`)
	default:
		return t, ErrWrongFormat
	}

	return
}

func (t Taxonomy) ValueFloat32() (f float64, err error) {
	f, err = strconv.ParseFloat(t.Value, 32)

	return
}

func (t Taxonomy) ValueFloat64() (f float64, err error) {
	f, err = strconv.ParseFloat(t.Value, 64)

	return
}

func (t Taxonomy) ValueInt() (i int, err error) {
	i, err = strconv.Atoi(t.Value)
	return
}
