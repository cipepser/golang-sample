package main
import(
	"fmt"
)

// type Task struct {
// 	ID int
// 	Detail string
// 	done bool
// }

// func Finish(task *Task) {
// 	task.done = true
// }

// コンストラクタ
// func NewTask(id int, detail string) *Task {
// 	task := &Task {
// 		ID: id,
// 		Detail: detail,
// 		done: false,
// 	}
// 	return task
// }

// メソッド
// func (task Task) String() string {
//     str := fmt.Sprintf("%d) %s\n", task.ID, task.Detail)
//     return str
// }
// 
// func (task *Task) Finish() {
// 	task.done = true
// }

// インタフェース
// type Stringer interface {
// 	String() string
// }
// 
// func Print(stringer Stringer) {
// 	fmt.Println(stringer.String())
// }


// 型の埋め込み
type User struct {
	FirstName string
	LastName string
}

func (u *User) FullName() string {
	fullname := fmt.Sprintf("%s %s", u.FirstName, u.LastName)
	return fullname
}

func NewUser(firstName, lastName string) *User {
	return &User {
		FirstName: firstName,
		LastName: lastName,
	}
}

type Task struct {
	ID int
	Detail string
	done bool
	*User
}

func NewTask(id int, detail, firstName, lastName string) *Task {
	task := &Task {
		ID: id,
		Detail: detail,
		done: false,
		User: NewUser(firstName, lastName),
	}
	return task
}

func main() {
	// var task Task = Task {
	// 	ID: 1,
	// 	Detail: "buy the milk",
	// 	done: true,
	// }
	
	// var task Task = Task {
	// 1, "buy the milk", true, 
	// }
	
	// task := Task{}
	// fmt.Println(task.ID)
	// fmt.Println(task.Detail)
	// fmt.Println(task.done)

	// ポインタ型
	// var task *Task = &Task{done: false}
	// Finish(task)
	// fmt.Println(task.done)
	
	// new()
	// var task *Task = new(Task)
	// task.ID = 1
	// task.Detail = "buy the milk"
	// fmt.Println(task.done)
	
	// コンストラクタ
	// NewHogemethodとするのが通例
	// task := NewTask(1, "buy the milk")
	// fmt.Printf("%+v\n", task)
	
	// メソッド
	// task := NewTask(1, "buy the milk")
	// task.Finish()
	// fmt.Printf("%s", task)

	// インタフェース
	// task := NewTask(1, "buy the milk")
	// Print(task)

	// 型の埋め込み
	task := NewTask(1, "buy the milk", "Jxck", "Daniel")
	fmt.Println(task.FirstName)
	fmt.Println(task.LastName)
	fmt.Println(task.FullName())
	fmt.Println(task.User)

}