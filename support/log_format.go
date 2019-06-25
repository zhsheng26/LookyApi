package support

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"sort"
	"strings"
	"time"
)

type Formatter struct {
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	levelColor := getColorByLevel(entry.Level)
	// output buffer
	b := &bytes.Buffer{}
	var log strings.Builder
	// write level
	level := strings.ToUpper(entry.Level.String())
	_, _ = fmt.Fprintf(b, "\x1b[%dm", levelColor)
	log.WriteString("[")
	log.WriteString(level[:4])
	log.WriteString("] ")
	// write time
	log.WriteString(FormatDate(time.Now(), YYYY_MM_DD_HH_MM_SS_EN))
	// write message
	msg := fmt.Sprintf(" [Message]: %s", strings.TrimSpace(entry.Message))
	log.WriteString(msg)
	// write caller
	caller := ""
	if entry.HasCaller() {
		funcVal := fmt.Sprintf(" [Func]: %s()", entry.Caller.Function)
		fileVal := fmt.Sprintf(" [File]: %s:%d", entry.Caller.File, entry.Caller.Line)
		if fileVal == "" {
			caller = funcVal
		} else if funcVal == "" {
			caller = fileVal
		} else {
			caller = fileVal + " " + funcVal
		}
	}

	log.WriteString(caller)
	// write fields-----------------
	if len(entry.Data) != 0 {
		log.WriteByte('\n')
		log.WriteString("       -> Field: ")
		fields := make([]string, 0, len(entry.Data))
		for field := range entry.Data {
			fields = append(fields, field)
		}

		sort.Strings(fields)
		for _, field := range fields {
			s := fmt.Sprintf("%s = %v; ", field, entry.Data[field])
			log.WriteString(s)
		}
	}

	log.WriteByte('\n')
	b.WriteString(log.String())
	return b.Bytes(), nil
}

func (f *Formatter) writeFields(b *bytes.Buffer, entry *logrus.Entry) {
	if len(entry.Data) != 0 {
		fields := make([]string, 0, len(entry.Data))
		for field := range entry.Data {
			fields = append(fields, field)
		}

		sort.Strings(fields)

		for _, field := range fields {
			f.writeField(b, entry, field)
		}
	}
}

func (f *Formatter) writeField(b *bytes.Buffer, entry *logrus.Entry, field string) {
	_, _ = fmt.Fprintf(b, "{%s:%v} ", field, entry.Data[field])
}

const (
	colorRed    = 31
	colorYellow = 33
	colorBlue   = 36
	colorGray   = 37
)

func getColorByLevel(level logrus.Level) int {
	switch level {
	case logrus.DebugLevel:
		return colorGray
	case logrus.WarnLevel:
		return colorYellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		return colorRed
	default:
		return colorBlue
	}
}
