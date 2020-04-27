#!/usr/bin/env bash

echo "build binary file"
go build

model_struct=./example
model_file=user.go
model_destination=./output/model
module_name=user
module_source=./output/model
module_destination=./output/module
service_file=service.go
service_source=./output/module

echo "建立 model"
./gecko general model -s $model_struct -f $model_file -d $model_destination
echo "建立 repository"
./gecko general repo -s $module_source -f $model_file -d $module_destination -p $module_name
echo "建立 service"
./gecko general svc -s $module_source -f $model_file -d $module_destination -p $module_name
echo "建立 delivery"
./gecko general delivery http -s $module_destination/$module_name -f $service_file -d $module_destination -p $module_name