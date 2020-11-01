package setting

import (
	"time"
)

type ServerSettingS struct {
	RunMode			string
	HttpPort		string
	ReadTimeOut		time.Duration
	WriteTimeOut 	time.Duration
}

type AppSettingS struct {
	DefaultPageSize	int
	MaxPageSize		int
	LogSavePath		string
	LogFileName		string
	LogFileExt		string
}

type DataBaseSettingS struct {
	DBType			string
	UserName		string
	Password		string
	Host			string
	DBName			string
	TablePrefix		string
	Charset			string
	ParseTime		bool
	MaxIdleConns	int
	MaxOpenConns	int
}

//unmarshal configs from FILE to STRUCT
func (s *Setting)ReadSection(k string,v interface{})error{
	err := s.vp.UnmarshalKey(k,v)
	if err != nil{
		return err
	}
	return nil
}

