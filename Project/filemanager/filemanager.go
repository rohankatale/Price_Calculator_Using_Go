package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath string
	OutputFilePath string
}


func(fm FileManager)ReadLines()([]string,error){
	file,err := os.Open(fm.InputFilePath)
	if err!=nil{
		return nil,errors.New("failed to open the file")
	}
	
	scanner := bufio.NewScanner(file)
	var data [] string
	for scanner.Scan(){
		data = append(data, scanner.Text())
	}
	defer file.Close()
	err = scanner.Err()
	if err!=nil{
		//file.Close()
		return nil,errors.New("failed to read line in the file")
	}

	 //file.Close()
	 return data,nil
}

func (fm FileManager)WriteResult(data interface{})error{
	file,err :=os.Create(fm.OutputFilePath)
	if err!=nil{
		return errors.New("failed to create a file")
	}
	time.Sleep(3*time.Second)

	encoder :=json.NewEncoder(file)
	err = encoder.Encode(data)
	if err!=nil{
		return errors.New("failed to convert data to json")
	}
	defer file.Close() 
	return nil
}

func New(inpath,outpath string)FileManager{

	return FileManager{
		InputFilePath: inpath, 
		OutputFilePath: outpath,
	}

}