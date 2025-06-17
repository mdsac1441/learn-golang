/******************************************************************************
Welcome to GDB Online.
GDB online is an online compiler and debugger tool for C, C++, Python, Java, PHP, Ruby, Perl,
C#, OCaml, VB, Swift, Pascal, Fortran, Haskell, Objective-C, Assembly, HTML, CSS, JS, SQLite, Prolog.
Code, Compile, Run and Debug online from anywhere in world.

*******************************************************************************/
package main
import "fmt"

    type Student struct{
     id uint
     name string
     age uint
    }
    
    func (s *Student) setStudentDetails( id uint,name string, age uint) {
        if id != s.id{
        s.id=id;
        s.name=name;
        s.age=age;
        }else{
            fmt.Printf("This Id already exists for Old Student");
        }

    }
    
    func(s Student) getStudentDetails( id uint){
        if s.id == id {
        fmt.Printf("id:%d\t",s.id);
        fmt.Printf("name:%s\t",s.name);
        fmt.Printf("age:%d",s.age);
        }else{
            fmt.Printf("Didn't match Id");
        }
    }


func main() {
    s1 :=Student{id:3,name:"a",age:9};
    s2 :=Student{id:3,name:"b",age:4};
    s1.getStudentDetails(3);
    fmt.Println();
    fmt.Println("(Before Change):", s1.name);
    (&s1).setStudentDetails(1,"ahmed",25);
    s1.getStudentDetails(1);
    fmt.Println();
    fmt.Println("(After Change):", s1.name);
    (&s2).setStudentDetails(1,"sharique",21);
    s2.getStudentDetails(1);
    fmt.Println();
}
    
