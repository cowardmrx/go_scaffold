package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type Model struct {
	UserModel *UserModel
}

type JsonTime struct {
	time.Time
}

type Timestamp struct {
	CreatedAt JsonTime `json:"created_at"`
	UpdatedAt JsonTime `json:"updated_at"`
}

//	@method MarshalJSON
//	@description: 为JsonTime 重写 MarshaJSON 方法，在此方法中实现自定义格式的转换；
//	@receiver jt
//	@return []byte
//	@return error
func (jt JsonTime) MarshalJSON() ([]byte, error) {
	if jt.IsZero() {
		return []byte(`""`), nil
	}
	output := fmt.Sprintf("\"%s\"", jt.Format("2006-01-02 15:04:05"))
	return []byte(output), nil
}

//	@method Value
//	@description: 为 JsonTime 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库；
//	@receiver jt
//	@return driver.Value
//	@return error
func (jt JsonTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if jt.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return jt.Time, nil
}

//	@method Scan
//	@description: 为 JsonTime 实现 Scan 方法，读取数据库时会调用该方法将时间数据转换成自定义时间类型；
//	@receiver jt
//	@param v interface{}
//	@return error
func (jt *JsonTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*jt = JsonTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// JSON json
type JSON json.RawMessage

//	@method Scan
//	@description: 实现 sql.Scanner 接口，Scan 将 value 扫描至 Json
//	@receiver j
//	@param value interface{}
//	@return error
func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Failed to unmarshal JSONB value:", value)
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

//	@method Value
//	@description: 实现 driver.Valuer 接口，Value 返回 json value
//	@receiver j
//	@return driver.Value
//	@return error
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}
