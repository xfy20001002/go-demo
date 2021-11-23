# oop

## 继承(inherit)
superman继承了human的方法
重写了Eat方法
新增了Fly方法
```
type SuperMan struct {
	Human //SuperMan类继承了Human类的方法
	level int
}
```



## 多态

通过interface实现多态

interface本质上是一个指针(**传递给接口时需要传递指针**)

具体的类实现interface中具有的方法(**可以新增在接口的基础上新增方法 但必须实现接口所有的方法**)



## 万能类型和类型推断(interfaceInfer)
interface{}是一个万能类型，所有类型内部都实现了它的接口，所以可以用interface类 型指向所有类型
要知道传入的参数到底为什么类型，可以使用类型推断机制(推断正确则ok为true 否则为false)
```
value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value = ", value)

		fmt.Printf("value type is %T\n", value)
	}
```