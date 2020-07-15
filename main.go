package main

var err error

func main() {
	// Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	// if err != nil {
	// 	fmt.Println("Status:", err)
	// }
	// defer Config.DB.Close()
	// Config.DB.DropTableIfExists(&Model.User{}, &Model.Account{})
	// Config.DB.AutoMigrate(&Model.User{}, &Model.Account{})
	// r := Routes.SetUpRouter()
	
	// r.Run()
	//select case for go concurrency
// 	c1 :=make(chan string);
// 	c2 :=make(chan string);
//    go func(){
// 	   for{
// 	   c1 <- "Every 500ms";
// 	   time.Sleep(time.Millisecond*500)
// 	}
//    }()
//    go func(){
// 	for {  c2 <-"Every 2 sec";
// 	   time.Sleep(time.Second*2);
// 	}
//    }()
// for{
// 	select{
// 	case msg1 := <-c1:
// 	fmt.Println(msg1);
// 	case msg2 := <-c2:
// 		fmt.Println(msg2);
// 	}
	

// }





}

// func Count(name string,c chan string){
// 	for i:=1; i<=5; i++{
// 		c <- name
// 		time.Sleep(time.Second*2)
// 	}
// 	close(c)
// }

func worker(jobs <-chan int,result <-chan int)  {
	for n:=range jobs{
	   result <-fibo(n)
	}
	
}

func fibo(n int) int  {
	if n<=1{
		return n;
	}
	return fibo(n-1)+fibo((n-2))
}