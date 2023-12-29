package src

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContextBackground(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)

}

// context with value

func TestContextVal(t *testing.T) {
	//this parent context
	ctxA := context.Background()

	//this child context
	contextB := context.WithValue(ctxA, "b", "B")
	contextC := context.WithValue(ctxA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
}

// context with cancel
/*
Selain menambahkan value ke context, kita juga bisa menambahkan sinyal cancel ke context
Kapan sinyal cancel diperlukan dalam context?
Biasanya ketika kita butuh menjalankan proses lain, dan kita ingin bisa memberi sinyal cancel ke
proses tersebut
Biasanya proses ini berupa goroutine yang berbeda, sehingga dengan mudah jika kita ingin
membatalkan eksekusi goroutine, kita bisa mengirim sinyal cancel ke context nya
Namun ingat, goroutine yang menggunakan context, tetap harus melakukan pengecekan terhadap
context nya, jika tidak, tidak ada gunanya
Untuk membuat context dengan cancel signal, kita bisa menggunakan function
context.WithCancel(parent)
*/

func TestContextCancel(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("counter", n)
		if n == 10 {
			break
		}
	}

	cancel()
	fmt.Println(runtime.NumGoroutine())
}

//context with timeout
/*
Selain menambahkan value ke context, dan juga sinyal cancel, kita juga bisa menambahkan sinyal
cancel ke context secara otomatis dengan menggunakan pengaturan timeout
Dengan menggunakan pengaturan timeout, kita tidak perlu melakukan eksekusi cancel secara
manual, cancel akan otomatis di eksekusi jika waktu timeout sudah terlewati
Penggunaan context dengan timeout sangat cocok ketika misal kita melakukan query ke database
atau http api, namun ingin menentukan batas maksimal timeout nya
Untuk membuat context dengan cancel signal secara otomatis menggunakan timeout, kita bisa
menggunakan function context.WithTimeout(parent, duration)
*/

func TestContextTimeOut(t *testing.T) {

	fmt.Println(runtime.NumGoroutine())

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)

	defer cancel()

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("counter: ", n)
	}

	fmt.Println(runtime.NumGoroutine())

}

// context with deadline

/*
Selain menggunakan timeout untuk melakukan cancel secara otomatis, kita juga bisa menggunakan
deadline
Pengaturan deadline sedikit berbeda dengan timeout, jika timeout kita beri waktu dari sekarang,
kalo deadline ditentukan kapan waktu timeout nya, misal jam 12 siang hari ini
Untuk membuat context dengan cancel signal secara otomatis menggunakan deadline, kita bisa
menggunakan function context.WithDeadline(parent, time)
*/

func TestContextDeadLine(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())

	ctx := context.Background()

	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("counter :", n)
	}

	fmt.Println(runtime.NumGoroutine())
}
