# validate

## validateDemo1
校验结构体字段(字符串)最大最小长度
校验结构体字段(整数)最大最小值
```
`validate:"min=6,max=10"`
validate := validator.New()
validate.Struct(struct)
```
## validateDemo2
校验结构体字段 非空约束
```
`validate:"required"`
```
## validateDemo3


## quick reference
```
对于数值，则约束其值；
对于字符串，则约束其长度；
对于切片、数组和map，则约束其长度。
范围约束
len：等于参数值，例如len=10；
max：小于等于参数值，例如max=10；
min：大于等于参数值，例如min=10；
eq：等于参数值，注意与len不同。对于字符串，eq约束字符串本身的值，而len约束字符串长度。例如eq=10；
ne：不等于参数值，例如ne=10；
gt：大于参数值，例如gt=10；
gte：大于等于参数值，例如gte=10；
lt：小于参数值，例如lt=10；
lte：小于等于参数值，例如lte=10；
oneof：只能是列举出的值其中一个，这些值必须是数值或字符串，以空格分隔，如果字符串中有空格，将字符串用单引号包围，例如oneof=red green。

字符串
contains=：包含参数子串，例如contains=email；
containsany：包含参数中任意的 UNICODE 字符，例如containsany=abcd；
containsrune：包含参数表示的 rune 字符，例如containsrune=☻；
excludes：不包含参数子串，例如excludes=email；
excludesall：不包含参数中任意的 UNICODE 字符，例如excludesall=abcd；
excludesrune：不包含参数表示的 rune 字符，excludesrune=☻；
startswith：以参数子串为前缀，例如startswith=hello；
endswith：以参数子串为后缀，例如endswith=bye。

唯一性
使用unqiue来指定唯一性约束，对不同类型的处理如下：

对于数组和切片，unique约束没有重复的元素；
对于map，unique约束没有重复的值；
对于元素类型为结构体的切片，unique约束结构体对象的某个字段不重复，通过unqiue=field指定这个字段名。
```




