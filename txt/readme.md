# txt

## Demo1
```
生成一个按格式的txt文件
file, error := os.Create("./barCode.txt")
file.Write([]byte(barCode))
func (f *File) WriteString(s string) (n int, err error)
file.Close()
```