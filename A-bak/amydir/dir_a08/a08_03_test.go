package dira08

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"runtime"
	"testgocode/amydir/dir_a08/json"
	"testing"
	"time"
	"unsafe"
)

type Serializer interface {
	Marshal(o interface{}) ([]byte, error)
	Unmarshal(d []byte, o interface{}) error
}

/*
type GotinySerializer struct {
	enc *gotiny.Encoder
	dec *gotiny.Decoder
}

func (g GotinySerializer) Marshal(o interface{}) ([]byte, error) {
	return g.enc.Encode(o), nil
}

func (g GotinySerializer) Unmarshal(d []byte, o interface{}) error {
	g.dec.Decode(d, o)
	return nil
}


func NewGotinySerializer(o interface{}) Serializer {
	ot := reflect.TypeOf(o)
	return GotinySerializer{
		enc: gotiny.NewEncoderWithType(ot),
		dec: gotiny.NewDecoderWithType(ot),
	}
}
*/

type GotinySerializer struct {
}

func (g GotinySerializer) Marshal(o interface{}) ([]byte, error) {
	return json.Marshal(o)
}

func (g GotinySerializer) Unmarshal(d []byte, o interface{}) error {
	return json.Unmarshal(d, o)
}
func NewGotinySerializer(o interface{}) Serializer {
	return GotinySerializer{}
}

func Benchmark_Gotiny_Marshal(b *testing.B) {
	benchMarshal(b, NewGotinySerializer(A{}))
}

func Benchmark_Gotiny_Unmarshal(b *testing.B) {
	benchUnmarshal(b, NewGotinySerializer(A{}))
}

//easyjson:json
type A struct {
	Name     string
	BirthDay time.Time
	Phone    string `gorm:"column:cost" json:"cost"`
	Siblings int
	Spouse   bool
	Money    float64
}

var (
	validate = os.Getenv("VALIDATE")
)

func randString(l int) string {
	buf := make([]byte, l)
	for i := 0; i < (l+1)/2; i++ {
		buf[i] = byte(rand.Intn(256))
	}
	return fmt.Sprintf("%x", buf)[:l]
}

func generate() []*A {
	a := make([]*A, 0, 1000)
	for i := 0; i < 1000; i++ {
		a = append(a, &A{
			Name:     randString(16),
			BirthDay: time.Now(),
			Phone:    randString(10),
			Siblings: rand.Intn(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
		})
	}
	return a
}

func benchMarshal(b *testing.B, s Serializer) {
	b.Helper()
	data := generate()
	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		o := data[rand.Intn(len(data))]
		bytes, err := s.Marshal(o)
		if err != nil {
			b.Fatalf("marshal error %s for %#v", err, o)
		}
		serialSize += len(bytes)
		if i == b.N-1 {
			// b2, _ := json.Marshal(o)
			// fmt.Println(string(b2))

			fmt.Println(string(bytes))
		}
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

func cmpTags(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if bv, ok := b[k]; !ok || bv != v {
			return false
		}
	}
	return true
}

func cmpAliases(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func benchUnmarshal(b *testing.B, s Serializer) {
	b.Helper()
	b.StopTimer()
	data := generate()
	ser := make([][]byte, len(data))
	var serialSize int
	for i, d := range data {
		o, err := s.Marshal(d)
		if err != nil {
			b.Fatal(err)
		}
		t := make([]byte, len(o))
		serialSize += copy(t, o)
		ser[i] = t
	}
	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		o := &A{}
		err := s.Unmarshal(ser[n], o)
		if err != nil {
			b.Fatalf("unmarshal error %s for %#x / %q", err, ser[n], ser[n])
		}
		// Validate unmarshalled data.
		if validate != "" {
			i := data[n]
			correct := o.Name == i.Name && o.Phone == i.Phone && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && o.BirthDay.Equal(i.BirthDay) //&& cmpTags(o.Tags, i.Tags) && cmpAliases(o.Aliases, i.Aliases)
			if !correct {
				b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
			}
		}
	}
}
func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
func TestTempDirInBenchmark(t *testing.T) {
	fArrs := []func(*testing.B){
		Benchmark_Gotiny_Marshal,
		Benchmark_Gotiny_Unmarshal,
	}
	for _, n := range fArrs {
		res := testing.Benchmark(n)
		fmt.Println(GetFunctionName(n),
			" NsPerOp:", uint64(res.NsPerOp()),
			" AllocsPerOp:", uint64(res.AllocsPerOp()),
			" AllocedBytesPerOp:", uint64(res.AllocedBytesPerOp()),
			"Seconds:", res.T.Seconds(),
			"Seconds runs:", float64(res.N)/res.T.Seconds(),
			" N:", res.N,
			" extra:", res.Extra,
		)
		// fmt.Println(res.String())
	}

}
func TestZxx(t *testing.T) {
	fmt.Println(unsafe.Sizeof(A{}))
	// panic("error 1")

	c := &A{}
	cType := reflect.TypeOf(c)
	cValue := reflect.ValueOf(c)
	structLen := cValue.Elem().NumField()
	fmt.Println(structLen)
	structMap := make(map[string]string, structLen)
	// structMap := make(map[string]string )  // 这样写也行

	for i := 0; i < structLen; i++ {
		field := cType.Elem().Field(i)
		jsonName := field.Tag.Get("json")
		if jsonName != "" {

			structMap[jsonName] = field.Name
		} else {
			structMap[field.Name] = field.Name

		}
	}
	fmt.Println(structMap)

}
