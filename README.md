# Golang JSON

Sumber Tutorial:
[Udemy](https://www.udemy.com/course/pemrograman-go-lang-pemula-sampai-mahir/learn/lecture/28064300#overview) |
[Slide](https://docs.google.com/presentation/d/1mGVdO7Khmiw-9lDmkCWqd9l5_UqohKRkpHbjabRD--U/edit)


## Pengenalan Package JSON
---


### Pengenalan JSON

- JSON singkatan dari JavaScript Object Notation, merupakan struktur format data yang bentuknya seperti Object di JavaScript
- JSON merupakan struktur format data yang paling banyak digunakan saat kita membuat RESTful API
- Dan pate kelas ini kita akan menggunakan JSON juga
- https://www.json.org/json-en.html 


### Kode: Contoh JSON

```json
{
    "language": "go",
    "isFavorite" : true,
    "usedHours": 69
}
```


### Package JSON

- Go-Lang sudah menyediakan package json, dimana kita bisa menggunakan package ini untuk melakukan konversi data ke JSON (encode) atau sebaliknya (decode)
- https://pkg.go.dev/encoding/json 


## Encode JSON
---

- Go-Lang telah menyediakan function untuk melakukan konversi data ke JSON, yaitu menggunakan function `json.Marshal(interface{})`
- Karena parameter nya adalah `interface{}`, maka kita bisa masukan tipe data apapun ke dalam function Marshal


### Kode: Encode JSON

```go
func logJSON(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJSONEncode(t *testing.T) {
	logJSON("Hello, World!")
	logJSON(123)
	logJSON(false)
	logJSON([]string{"a", "b", "c"})
	logJSON(map[string]string{"hello": "world"})
}
```


## JSON Object
---

- Pada materi sebelumnya kita melakukan encode data seperti string, number, boolean, dan tipe data primitif lainnya
- Walaupun memang bisa dilakukan, karena sesuai dengan kontrak `interface{}`, namun tidak sesuai dengan kontrak JSON
- Jika mengikuti kontrak json.org, data JSON bentuknya adalah `Object` dan `Array`
- Sedangkan value nya baru berupa 


### Kode: Contoh JSON Object

```json
{
    "language": "go",
    "isFavorite" : true,
    "usedHours": 69
}
```


### Struct

- JSON Object di Go-Lang direpresentasikan dengan tipe data Struct
- Dimana tiap attribute di JSON Object merupakan attribute di Struct


### Kode: Encode Struct ke JSON

```go
type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "John",
		MiddleName: "Doe",
		LastName:   "Smith",
		Age:        25,
		Married:    true,
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	t.Log(string(bytes))
}
```


## Decode JSON
---

- Sekarang kita sudah tahu bagaimana caranya melakukan encode dari tipe data di Go-Lang ke JSON
- Namun bagaimana jika kebalikannya?
- Untuk melakukan konversi dari JSON ke tipe data di Go-Lang (Decode), kita bisa menggunakan function `json.Unmarshal(byte[], interface{})`
- Dimana `byte[]` adalah data JSON nya, sedangkan `interface{}` adalah tempat menyimpan hasil konversi, biasa berupa pointer


### Kode: Decode JSON

```go
func TestJSONDecode(t *testing.T) {
	data := `{"FirstName":"John","MiddleName":"Doe","LastName":"Smith","Age":25,"Married":true}`
	bytes := []byte(data)

	var customer Customer
	err := json.Unmarshal(bytes, &customer)
	if err != nil {
		t.Errorf("Error unmarshalling JSON: %s", err)
		panic(err)
	}

	t.Log(customer)
}
```


## JSON Array
---

- Selain tipe dalam bentuk Object, biasanya dalam JSON, kita kadang menggunakan tipe data Array
- Array di JSON mirip dengan Array di JavaScript, dia bisa berisikan tipe data primitif, atau tipe data kompleks (Object atau Array)
- Di Go-Lang, JSON Array direpresentasikan dalam bentuk slice
- Konversi dari JSON atau ke JSON dilakukan secara otomatis oleh package json menggunakan tipe data slice


### Kode: Customer 

```go
type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
}
```


### Kode: Address

```go
type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}
```


### Kode: JSON Array Encode

```go
func TestJSONArrayEncode(t *testing.T) {
	customer := Customer{
		FirstName:  "John",
		MiddleName: "Doe",
		LastName:   "Smith",
		Hobbies:    []string{"Skiing", "Snowboarding", "Go"},
		Addresses: []Address{
			{
				Street: "123 Main St",
				City:   "Anytown",
				State:  "CA",
				Zip:    "12345",
			},
			{
				Street: "456 Elm St",
				City:   "Anytown",
				State:  "CA",
				Zip:    "12345",
			},
		},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		t.Error(err)
		panic(err)
	}

	t.Log(string(bytes))
}
```


### Kode: JSON Array Decode

```go
func TestJSONArrayDecode(t *testing.T) {
	customer := Customer{}
	err := json.Unmarshal([]byte(`{
		"firstName": "John",
		"middleName": "Doe",
		"lastName": "Smith",
		"hobbies": ["Skiing", "Snowboarding", "Go"],
		"addresses": [
			{
				"street": "123 Main St",
				"city": "Anytown",
				"state": "CA",
				"zip": "12345"
			},
			{
				"street": "456 Elm St",
				"city": "Anytown",
				"state": "CA",
				"zip": "12345"
			}
		]
	}`), &customer)
	if err != nil {
		t.Error(err)
		panic(err)
	}

	t.Log(customer)
	t.Log(customer.Addresses)
}
```


### Kode: Encode JSON Array

```go
func TestJSONArrayDecodeDirect(t *testing.T) {
	data := `[
			{
				"street": "123 Main St",
				"city": "Anytown",
				"state": "CA",
				"zip": "12345"
			},
			{
				"street": "456 Elm St",
				"city": "Anytown",
				"state": "CA",
				"zip": "12345"
			}
		]`
	bytes := []byte(data)
	addresses := []Address{}
	err := json.Unmarshal(bytes, &addresses)
	if err != nil {
		t.Error(err)
		panic(err)
	}

	t.Log(addresses)
}
```


## JSON Tag
---

- Secara default atribut yang terdapat di Struct dan JSON akan di mapping sesuai dengan nama atribut  yang sama (case sensitive)
- Kadang ada style yang berbeda antara penamaan atribute di Struct dan di JSON, misal di JSON kita ingin menggunakan snake_case, tapi di Struct, kita ingin menggunakan PascalCase
- Untungnya, package json mendukun Tag Reflection
- Kita bisa menambahkan tag reflection dengan nama json, lalu diikuti dengan atribut yang kita inginkan ketika konversi dari atau ke JSON


### Kode: JSON Tag

```go
type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    string `json:"price"`
	ImageURL string `json:"image_url"`
}
```


## Map
---

- Saat menggunakan JSON, kadang mungkin kita menemukan kasus data JSON nya dynamic
- Artinya atribut nya tidak menentu, bisa bertambah, bisa berkurang, dan tidak tetap
- Pada kasus seperti itu, menggunakan Struct akan menyulitkan, karena pada Struct, kita harus menentukan semua atribut nya
- Untuk kasus seperti ini, kita bisa menggunakan tipe data `map[string]interface{}`
- Secara otomatis, atribut akan menjadi key di map, dan value menjadi value di map
- Namun karena value berupa `interface{}`, maka kita harus lakukan konversi secara manual jika ingin mengambil value nya
- Dan tipe data Map tidak mendukung JSON Tag lagi


### Kode: Map

```go
func TestJSONMap(t *testing.T) {
	data := `{"a":1,"b":2,"c":3}`
	bytes := []byte(data)

	var result map[string]interface{}
	err := json.Unmarshal(bytes, &result)
	if err != nil {
		t.Error(err)
	}

	t.Log(result)
}
```


## Streaming Decoder
---

- Sebelumnya kita belajar package json dengan melakukan konversi data JSON yang sudah dalam bentuk variable dan data `string` atau `[]byte`
- Pada kenyataanya, kadang data JSON nya berasal dari Input berupa `io.Reader (File, Network, Request Body)`
- Kita bisa saja membaca semua datanya terlebih dahulu, lalu simpan di variable, baru lakukan konversi dari JSON, namun hal ini sebenarnya tidak perlu dilakukan, karena package json memiliki fitur untuk membaca data dari Stream


### `json.Decoder`

- Untuk membuat json Decoder, kita bisa menggunakan function `json.NewDecoder(reader)`
- Selanjutnya untuk membaca isi input reader dan konversikan secara langsung ke data di Go-Lang, cukup gunakan function `Decode(interface{})`


### Kode: Streaming Decoder

```go
func TestJSONStreamDecoder(t *testing.T) {
	reader, _ := os.Open("./resources/input.json")
	decoder := json.NewDecoder(reader)

	product := &Product{}
	err := decoder.Decode(product)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
		panic(err)
	}

	t.Log(product)
}
```


## Streaming Encoder
---

- Selain decoder, package json juga mendukung membuat Encoder yang bisa digunakan untuk menulis langsung JSON nya ke `io.Writer`
- Dengan begitu, kita tidak perlu menyimpan JSON datanya terlebih dahulu ke dalam variable string atau `[]byte`, kita bisa langsung tulis ke `io.Writer`


### `json.Encoder`

- Untuk membuat Encoder, kita bisa menggunakan function `json.NewEncoder(writer)`
- Dan untuk menulis data sebagai JSON langsung ke writer, kita bisa gunakan function Encode(`interface{}`)


### Kode: Streaming Encoder

```go
func TestJSONStreamEncoder(t *testing.T) {
	product := &Product{
		Id:       "MBP-1",
		Name:     "MacBook Pro",
		Price:    123123,
		ImageURL: "mekbuk.com",
	}

	writer, _ := os.Create("./resources/output.json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(product)

	if err != nil {
		t.Errorf("Error encoding JSON: %s", err)
		panic(err)
	}
}
```