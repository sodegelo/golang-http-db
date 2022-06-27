package main
 
import ( 
	"database/sql"
	"fmt"
	"time" 
	"math"
	"math/rand"
	"net/http" 
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Name        string `json:"nome"`
	Description string `json:"description"`
}

var products []Product

func main() {
	execThread();
	//routes()
}  
 
func webServerWorker(workerId int, msg chan int )  {
	for res := range msg {
		res_2 := math.Sqrt(math.Sqrt(rand.Float64()))
		fmt.Println("Workerid:",workerId, " Mensagem processada: ", res," resultado:",res_2)
		//time.Sleep(time.Millisecond * 100)
	}
}

func execThread()  {
	msg := make(chan int)

	 
	for i := 0; i < 10000; i++ {
		go webServerWorker(i,msg)
	}

	for i := 0; i < 1000000; i++ {
		msg <- i
	}
	time.Sleep(time.Second * 10)
}
func routes()  {
	generateProducts()
	e := echo.New()
 
	e.GET("/list", listAll)
    e.POST("/product", createProduct)
	e.Logger.Fatal(e.Start(":8080"))
}

func generateProducts()  {
	p1 := Product{Name: "Julio", Description: "Julio"}
	p2 := Product{Name: "Julio2", Description: "Julio2"}
	products = append(products,p1,p2)
}


 
func createProduct( c echo.Context) error{
	product := Product{}
	c.Bind(&product)
	err := persistProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil);
	}
	return c.JSON(http.StatusCreated, product);
}
func listAll(c echo.Context) error{
	return c.JSON(200, products)
}

 
func persistProduct(product Product) error {
	db, err := sql.Open("sqlite3", "production.db")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO products(name, description) values($1,$2)")

	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(product.Name, product.Description)

	if err != nil{
		return err
	}

	return nil;
}
