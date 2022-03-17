/*
   @Author: StudentCWZ
   @Description:
   @File: snoyflake
   @Software: GoLand
   @Project: GoWeb
   @Date: 2022/3/17 14:52
*/

package main

import (
	"fmt"
	"github.com/sony/sonyflake"
	"time"
)

var (
	sonyFlake     *sonyflake.Sonyflake
	sonyMachineID uint16
)

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

// Init 需传入当前的机器 ID
func Init(startTime string, machineId uint16) (err error) {
	sonyMachineID = machineId
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	settings := sonyflake.Settings{
		StartTime: st,
		MachineID: getMachineID,
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	return
}

// GenID 生成 id
func GenID() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("sony flake not inited")
		return
	}
	id, err = sonyFlake.NextID()
	return
}

func main() {
	if err := Init("2020-07-01", 1); err != nil {
		fmt.Printf("Init failed, err: %v\n", err)
		return
	}
	id, _ := GenID()
	fmt.Println(id)
}
