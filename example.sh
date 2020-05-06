#!/usr/bin/env bash

#echo "get gecko"
#go get -u github.com/codingXiang/gecko
#
model_struct=./example
model_file=user.go
model_destination=./output/model
module_name=user
module_source=./output/model
module_destination=./output/module
service_file=service.go
service_source=./output/module
go build
echo "Create model"
./gecko general model -s $model_struct -f $model_file -d $model_destination
echo "Create repository"
./gecko general repo -s $module_source -f $model_file -d $module_destination -p $module_name
echo "Create service"
./gecko general svc -s $module_source -f $model_file -d $module_destination -p $module_name
echo "Create delivery"
./gecko general delivery http -s $module_destination/$module_name -f $service_file -d $module_destination -p $module_name