package ass

import (
	"errors"
	"fmt"
)

type CodepointSet map[rune]struct{}

type ContentInfo struct {
	LineNum    uint   // 行号
	RawContent string // 文本内容
}

// 新增：格式定义结构体
type FormatInfo struct {
	Fields []string // 字段名称列表
}

// 重新设计：样式信息结构体
type StyleInfo struct {
	Content    *ContentInfo      // 原始内容
	Fields     map[string]string // 字段名->值的映射
	FormatInfo *FormatInfo       // 格式定义
}

// 重新设计：对话信息结构体
type DialogueInfo struct {
	Content    *ContentInfo      // 原始内容
	Fields     map[string]string // 字段名->值的映射
	FormatInfo *FormatInfo       // 格式定义
}

type FontDesc struct {
	FontName string // 字体名称
	Bold     uint   // 字粗
	Italic   uint   // 是否启用斜体，0->不启用
}

// String 返回 FontDesc 的字符串表示，用于排序
func (fd *FontDesc) String() string {
	return fmt.Sprintf("%s_%d_%d", fd.FontName, fd.Bold, fd.Italic)
}

const (
	defaultFontName     = "Default" // 默认字体名称
	defaultFontSize     = 400       // 默认字体大小
	defaultBoldFontSize = 700       // 默认粗细大小
	defaultItalic       = 0         // 默认不斜体
	defaultItalicSlant  = 100       // 默认斜体倾斜度
)

var (
	ErrStyleParseFailed   = errors.New("failed to parse style") // 未找到 [V4 Styles] 等模块
	ErrInvalidStyleFormat = errors.New("invalid style format")  // Styles 格式解析失败
	ErrEventParseFailed   = errors.New("failed to parse event") // 未找到 [Events] 等模块
	ErrInvalidEventFormat = errors.New("invalid event format")  // Events 格式解析失败
	ErrInvalidBoldValue   = errors.New("invalid bold value")    // 不合法字重值
	ErrInvalidItalicValue = errors.New("invalid italic value")  // 不合法斜体值
	ErrMissingFormat      = errors.New("missing format line")   // 缺少格式定义行
)

// 新增：解析状态结构体
type parseState struct {
	inStyleSection bool // 是否在 [V4 Styles] 模块中
	inEventSection bool // 是否在 [Events] 模块中
	hasStyle       bool // 是否已找到 [V4 Styles] 模块
	hasEvent       bool // 是否已找到 [Events] 模块
}
