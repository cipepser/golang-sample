package main
import(
	"fmt"
)


type Task struct {
	ID int
	Detail string
	done bool
}

// func Finish(task *Task) {
// 	task.done = true
// }

// コンストラクタ
func NewTask(id int, detail string) *Task {
	task := &Task {
		ID: id,
		Detail: detail,
		done: false,
	}
	return task
}

// メソッド
// func (task Task) String() string {
//     str := fmt.Sprintf("%d) %s\n", task.ID, task.Detail)
//     return str
// }

func (task *Task) Finish() {
	task.done = true
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
	task := NewTask(1, "buy the milk")
	task.Finish()
	fmt.Printf("%s", task)
	
}