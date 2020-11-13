package main 
   
import ( 
    "fmt"
    "encoding/json"
) 
   
type Employee struct{ 
       
    Name string 
    Address string 
    Age int
} 
   
 
func main() { 
       
    
    var emp Employee 

    Data := []byte(`{ 
        "Name": "Deeksha",   
        "Address": "Hyderabad", 
        "Age": 21 
    }`) 
       
    err := json.Unmarshal(Data, &emp) 
       
    if err != nil { 
           
            fmt.Println(err) 
    } 
       
    fmt.Println("Struct is:", emp) 
    fmt.Printf("%s lives in %s.\n", emp.Name, emp.Address) 
       
    var emp2 []Employee 
       
    Data2 := []byte(` 
    [ 
        {"Name": "Vani", "Address": "Delhi", "Age": 21}, 
        {"Name": "Rashi", "Address": "Noida", "Age": 24}, 
        {"Name": "Rohit", "Address": "Pune", "Age": 25} 
    ]`) 
       
    err2 := json.Unmarshal(Data2, &emp2) 
       
        if err2 != nil { 
       
            fmt.Println(err2) 
        } 
       
    for i := range emp2{ 
       
        fmt.Println(emp2[i]) 
    } 
} 