import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'm Abir.\n A golang learner.")
}

//	type Product struct {
//		// id made private
//		id int
//		// building alias, to pass custom case
//		Title       string `json:"custom-title"`
//		Description string
//		Price       float64
//		ImgUrl      string
//	}
type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	allowCors(w)
	if r.Method != http.MethodGet {
		http.Error(w, "Please, give me GET request", 400)
		return
	}

	sendData(w, productList, 200)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	allowCors(w)
	allowOptions(w, r)
	if r.Method != http.MethodPost {
		http.Error(w, "Please, give me Post request", 400)
		return
	}
	// r.Body => description, imageUrl, price, title =>
	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz, give me valid json", 400)
		return
	}
	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)
	sendData(w, productList, 201)
}

func allowCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

}

func allowOptions(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(200)
	}
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)

	mux.HandleFunc("/about", aboutHandler)

	mux.HandleFunc("/api/products", getProducts)
	mux.HandleFunc("/api/create-product", createProduct)

	fmt.Println("Server running on :5000")

	err := http.ListenAndServe(":5000", mux)

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is orange, I love orange",
		Price:       100,
		ImgUrl:      "https://imgs.search.brave.com/pteG9T1-Pgd47tAC6az-lh1OjLB1aoxqpWUhlU37z38/rs:fit:500:0:1:0/g:ce/aHR0cHM6Ly93d3cu/cmQuY29tL3dwLWNv/bnRlbnQvdXBsb2Fk/cy8yMDE3LzEyLzAx/X29yYW5nZXNfRmlu/YWxseSVFMiU4MCU5/NEhlcmUlRTIlODAl/OTlzLVdoaWNoLSVF/MiU4MCU5Q09yYW5n/ZSVFMiU4MCU5RC1D/YW1lLUZpcnN0LXRo/ZS1Db2xvci1vci10/aGUtRnJ1aXRfNjkx/MDY0MzUzX0x1Y2t5/LUJ1c2luZXNzLmpw/Zz9maXQ9NjQwLDQy/Nw",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is green",
		Price:       100,
		ImgUrl:      "https://imgs.search.brave.com/k64M2mdcB9vnknWz3p_Ovw1zKux-E81TzEBt-UVDJXs/rs:fit:500:0:1:0/g:ce/aHR0cHM6Ly9tZWRp/YS5pc3RvY2twaG90/by5jb20vaWQvMTE1/MjA2Nzc3Mi9waG90/by9kZWxpY2lvdXMt/cmVkLWFwcGxlcy1v/bi1yZXRhaWwtZGlz/cGxheS1hdC1zdXBl/cm1hcmtldC5qcGc_/cz02MTJ4NjEyJnc9/MCZrPTIwJmM9enZB/ckJKVmZtM1lyQlhO/ZHN4YWFXU2VMVEJU/RmVhM0VTOTg1bVhR/QXFtaz0",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is Green",
		Price:       100,
		ImgUrl:      "https://imgs.search.brave.com/Bh203dGWHpOJQYtYaG77ohn8ZCpH8xt0veS9QVN1FMg/rs:fit:500:0:1:0/g:ce/aHR0cHM6Ly9tZWRp/YS5nZXR0eWltYWdl/cy5jb20vaWQvMjE2/ODc1MjgxNi9waG90/by9jbG9zZS11cC1v/Zi1iYW5hbmEtdHJl/ZS5qcGc_cz02MTJ4/NjEyJnc9MCZrPTIw/JmM9aXBHNWRNLUxk/R0ZEY0hta1ZOSjlJ/LUp2X2ZvdDNKY2Vz/Q0V6MlZnUjVRTT0",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Guava",
		Description: "Guava is green",
		Price:       100,
		ImgUrl:      "https://imgs.search.brave.com/VosOYMkoCA43ivdH_eD_232M3utY3zTMhDCCUHXR-hU/rs:fit:0:180:1:0/g:ce/aHR0cHM6Ly9jZG4u/bW9zLmNtcy5mdXR1/cmVjZG4ubmV0L1FX/SlplaWo3cGlwbnNp/YTc1ZEdOdlgtMjMw/LTgwLmpwZw",
	}
	prd5 := Product{
		ID:          5,
		Title:       "Pomegranate",
		Description: "Pomegranate is red",
		Price:       100,
		ImgUrl:      "https://imgs.search.brave.com/lr54-BpcmvudejK69bknqjndUkfJQ0VOImgtugxbems/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly93d3cu/bnV0cml0aW9uYWR2/YW5jZS5jb20vd3At/Y29udGVudC91cGxv/YWRzLzIwMjMvMDgv/Y3V0LXBvbWVncmFu/YXRlLXNob3dpbmct/cmVkLXNlZWRzLmpw/Zw",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	productList = append(productList, prd4)
	productList = append(productList, prd5)
}
