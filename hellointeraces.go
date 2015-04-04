package main
import "fmt"
type programmer struct{
   name string
}

type developer struct{
   programmer
   language string
   linesPerHour int
}

func (d developer) canWrite(languageName string)bool{
   if d.language == languageName{
      return true
   }
   
   return false
}

func main(){
   fmt.Println("Hello, World!")
   languagesArr := []string{"java", "python"}
   developersArr:= []developer{developer{programmer {"John"}, "java", 100}, developer{programmer{"culey"}, "python", 500}} 
   developersArr = append(developersArr, developer{programmer{"Hiren"}, "go", 500})
   fmt.Println("Languages: ")
   for n := range languagesArr{
      fmt.Println(languagesArr[n])
   }
   
   fmt.Println("Developers: ")
   for _, d := range developersArr{
      fmt.Println(d.name)
   }
   
   
   for _, d:= range developersArr{
      for _,l:= range languagesArr{
         if d.canWrite(l){
            fmt.Printf("%s can write %s code \n", d.name, l)
         }else{
             fmt.Printf("%s can't write %s code \n", d.name, l)

         }
      }
   }
}

