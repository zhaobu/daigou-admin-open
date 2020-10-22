package utils

import (
	"daigou-admin/global"
	"daigou-admin/utils/zaplog"
	"fmt"
	"os"

	"github.com/gogf/gf/util/gutil"
	"github.com/sony/sonyflake"
)

var flake *sonyflake.Sonyflake

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

func MakeDirs(paths []string) error {
	for _, v := range paths {
		err := os.MkdirAll(v, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func GenUniqueID() uint64 {
	if flake == nil {
		flake = sonyflake.NewSonyflake(sonyflake.Settings{
			MachineID: func() (uint16, error) {
				return 0, nil
			},
		})
	}
	id, err := flake.NextID()
	if err != nil {
		zaplog.Errorf("flake.NextID() failed with %s", err)
	}
	return id
}

func Dump(v interface{}) string {
	if global.G_Config.System.Mode == "debug" {
		return gutil.Export(v)
	}
	return fmt.Sprintf("%v", v)
}
