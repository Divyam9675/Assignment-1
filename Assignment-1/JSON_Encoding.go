package main 
   
import ( 
    "fmt"
    "encoding/json"
) 
   
 
type Employee struct{ 
       
   
    Name string 
    Age int
    Address string 
} 
   

func main() { 
       
   
    emp := Employee{"Ankit", 23, "New Delhi"} 
       
  
    emp_in, err := json.Marshal(emp) 
       
    if err != nil { 
           
       
        fmt.Println(err) 
    } 
       
   
    fmt.Println(string(emp_in)) 
    
  
    emp2 := []Employee{ 
        {Name: "Rahul", Age: 23, Address: "New Delhi"}, 
        {Name: "Priyanshi", Age: 20, Address: "Pune"}, 
        {Name: "Shivam", Age: 24, Address: "Bangalore"}, 
    } 
       
    
    emp2_in, err := json.Marshal(emp2) 
        
        if err != nil { 
       
            fmt.Println(err) 
        } 
           
    
    fmt.Println() 
        fmt.Println(string(emp2_in)) 
}