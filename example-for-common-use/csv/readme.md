# csv相关操作

## 写入并读取csv文件(Demo1)

```
//os
func Create(name string) (*File, error)
```

```
//csv
func NewWriter(w io.Writer) *Writer    //file实现了writer、reader接口
func (w *Writer) Write(record []string) error  //写入数据
func (w *Writer) Flush() //刷入数据
```

```
//os
func Open(name string) (*File, error)
```

```
//csv
func NewReader(r io.Reader) *Reader  
func (r *Reader) ReadAll() (records [][]string, err error) //读取所有数据
func (w *Writer) Flush() //将数据刷新到缓冲区
```

## 将csv数据写入bytes缓冲区(Demo2)





