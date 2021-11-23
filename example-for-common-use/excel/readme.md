# excel

## 将结构体写入excel文件并保存(Demo1)

```
func NewFile() *File
func (f *File) AddSheet(sheetName string) (*Sheet, error)
func (s *Sheet) AddRow() *Row
func (r *Row) AddCell() *Cell
func (c *Cell) SetValue(n interface{})

func (f *File) Save(path string) (err error)
```



## 将结构体以excel形式写入字节流缓冲区(Demo2)

```
func (f *File) Write(writer io.Writer) (err error)  //将excel数据写入某个实现了writer接口的数据区(文件、缓冲区)
```

