package flag

// Value接口代表了存储在标志内的值
type Value interface {
	String() string   // 用于格式化标志对应的值，可用于输出命令行帮助信息。有了该方法，每个flag.Value其实也是fmt.Stringer
	Set(string) error // 解析了传入的字符串参数并更新标志值。可以认为set方法是string方法的逆操作
}
